package main

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"expvar"
	"flag"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	stdlog "log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/linfn/camo"
	"github.com/linfn/camo/internal/envflag"
)

var camoDir = getCamoDir()

var (
	help      = flag.Bool("help", false, "help")
	password  = envflag.String("password", "CAMO_PASSWORD", "", "Set a password. It is recommended to use the environment variable CAMO_PASSWORD to set the password.")
	tun4      = envflag.Bool("4", "CAMO_ENABLE_IP4", false, "tunneling for IPv4 only")
	tun6      = envflag.Bool("6", "CAMO_ENABLE_IP6", false, "tunneling for IPv6 only")
	resolve   = envflag.String("resolve", "CAMO_RESOLVE", "", "provide a custom address for a specific host and port pair")
	resolve4  = envflag.Bool("resolve4", "CAMO_RESOLVE4", false, "resolve host name to IPv4 addresses only")
	resolve6  = envflag.Bool("resolve6", "CAMO_RESOLVE6", false, "resolve host name to IPv6 addresses only")
	mtu       = envflag.Int("mtu", "CAMO_MTU", camo.DefaultMTU, "mtu")
	reGateway = envflag.Bool("redirect-gateway", "CAMO_REDIRECT_GATEWAY", true, "redirect the gateway")
	logLevel  = envflag.String("log-level", "CAMO_LOG_LEVEL", camo.LogLevelTexts[camo.LogLevelInfo], "log level")
	useH2C    = envflag.Bool("h2c", "CAMO_H2C", false, "use h2c (for debug)")
	debugHTTP = envflag.String("debug-http", "CAMO_DEBUG_HTTP", "", "debug http server listen address")
)

var (
	log        *camo.LevelLogger
	host       string
	cid        string
	remoteAddr atomic.Value
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] host\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if *help {
		return
	}

	initLog()

	host = flag.Arg(0)
	if host == "" {
		host = os.Getenv("CAMO_HOST")
		if host == "" {
			log.Fatal("missing host")
		}
	}

	if !*tun4 && !*tun6 {
		*tun4 = true
		*tun6 = true
	}

	if *resolve4 && *resolve6 {
		log.Fatal("can not use -resolve4 and -resolve6 at the same time")
	}

	if *resolve != "" {
		addr, err := camo.GetHostPortAddr(*resolve, "443")
		if err != nil {
			log.Fatalf("resolve addr %s error: %v", *resolve, err)
		}
		*resolve = addr
	}

	if *password == "" {
		log.Fatal("missing password")
	}
	// hidden the password to expvar and pprof package
	for i := range os.Args {
		if os.Args[i] == "-password" || os.Args[i] == "--password" {
			os.Args[i+1] = "*"
		}
	}

	cid = ensureCID(host)
}

func main() {
	if *help {
		flag.Usage()
		return
	}

	iface, err := camo.NewTun(*mtu)
	if err != nil {
		log.Fatalf("failed to create tun device: %v", err)
	}
	defer iface.Close()

	ctx, cancel := context.WithCancel(context.Background())

	c := &camo.Client{
		MTU:  *mtu,
		CID:  cid,
		Host: host,
		Dial: func(network, addr string) (net.Conn, error) {
			if *resolve4 {
				network = "tcp4"
			} else if *resolve6 {
				network = "tcp6"
			}
			if *resolve != "" {
				addr = *resolve
			}
			var d net.Dialer
			conn, err := d.DialContext(ctx, network, addr)
			if err == nil {
				remoteAddr.Store(conn.RemoteAddr())
				log.Infof("connection succeeded. remote: %s", conn.RemoteAddr())
			}
			return conn, err
		},
		Auth:   func(r *http.Request) { camo.SetAuth(r, *password) },
		Logger: log,
		UseH2C: *useH2C,
		Noise:  getNoise(),
	}

	expvar.Publish("camo", c.Metrics())

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		log.Debugf("receive signal %s", <-c)
		cancel()
	}()

	if *debugHTTP != "" {
		go debugHTTPServer()
	}

	runClient(ctx, c, iface)
}

func initLog() {
	logLevel, ok := camo.LogLevelValues[strings.ToUpper(*logLevel)]
	if !ok {
		stdlog.Fatal("invalid log level")
	}
	log = camo.NewLogger(stdlog.New(os.Stderr, "", stdlog.LstdFlags|stdlog.Llongfile), logLevel)
}

func getCamoDir() string {
	dir, err := os.UserCacheDir()
	if err == nil {
		return dir + "/camo"
	}
	return ".camo"
}

func ensureCamoDir() {
	err := os.MkdirAll(camoDir, os.ModePerm)
	if err != nil {
		log.Panicf("failed to create camo dir. path: %s, error: %v", camoDir, err)
	}
}

func ensureCID(host string) string {
	id, err := machineid.ProtectedID("camo@" + host)
	if err == nil {
		return id
	}
	log.Warnf("failed to get protected machineid: %v", err)

	cidFile := camoDir + "/cid"

	b, err := ioutil.ReadFile(cidFile)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("failed to read cid file. path: %s, error: %v", cidFile, err)
	}
	if len(b) == 0 {
		b = make([]byte, 32)
		if _, err = rand.Read(b); err != nil {
			log.Fatalf("failed to generate rand: %v", err)
		}
		ensureCamoDir()
		err = ioutil.WriteFile(cidFile, b, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to save cid file. path: %s, error: %v", cidFile, err)
		}
		log.Debugf("cid file saved. path: %s", cidFile)
	}

	log.Debugf("load cid from %s", cidFile)

	mac := hmac.New(sha256.New, b)
	mac.Write([]byte("camo@" + host))
	return hex.EncodeToString(mac.Sum(nil))
}

func getNoise() int {
	return int(crc32.ChecksumIEEE([]byte(cid)))
}

func setupTunHandler(c *camo.Client, iface *camo.Iface) func(net.IP, net.IPMask, net.IP) (func(), error) {
	return func(tunIP net.IP, mask net.IPMask, gateway net.IP) (reset func(), err error) {
		var rollback camo.Rollback
		defer func() {
			if err != nil {
				rollback.Do()
			}
		}()

		var (
			cidr     = camo.ToCIDR(tunIP, mask)
			tunIPVer int
		)
		if tunIP.To4() != nil {
			tunIPVer = 4
			if err = iface.SetIPv4(cidr, gateway); err != nil {
				return nil, err
			}
			rollback.Add(func() { iface.SetIPv4("", nil) })
		} else {
			tunIPVer = 6
			if err = iface.SetIPv6(cidr); err != nil {
				return nil, err
			}
			rollback.Add(func() { iface.SetIPv6("") })
		}
		log.Infof("%s(%s) up", iface.Name(), cidr)

		if *reGateway {
			// bypass tun for server ip
			srvAddr, _ := remoteAddr.Load().(net.Addr)
			if srvAddr == nil {
				return nil, errors.New("failed to get server address")
			}
			srvIP, _, err := net.SplitHostPort(srvAddr.String())
			if err != nil {
				return nil, fmt.Errorf("failed to get server ip: %v (%s)", err, srvAddr)
			}
			srvIPVer := 4
			if !camo.IsIPv4(srvIP) {
				srvIPVer = 6
			}
			if srvIPVer == tunIPVer {
				oldGateway, oldDev, err := camo.GetRoute(srvIP)
				if err != nil {
					return nil, err
				}
				err = camo.AddRoute(srvIP, oldGateway, oldDev)
				if err != nil {
					return nil, err
				}
				rollback.Add(func() { camo.DelRoute(srvIP, oldGateway, oldDev) })
			}

			resetGateway, err := camo.RedirectGateway(iface.Name(), gateway.String())
			if err != nil {
				return nil, err
			}
			rollback.Add(resetGateway)
		}

		return rollback.Do, nil
	}
}

func runClient(ctx context.Context, c *camo.Client, iface *camo.Iface) {
	openTunnel := func(ctx context.Context, ipVersion int) (func(context.Context) error, error) {
		var err error

		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)

		var res *camo.IPResult
		if ipVersion == 4 {
			res, err = c.RequestIPv4(ctx)
		} else {
			res, err = c.RequestIPv6(ctx)
		}
		if err != nil {
			cancel()
			return nil, err
		}

		log.Infof("client get %s", res)

		var (
			ip   = res.IP
			mask = res.Mask
			gw   = res.Gateway
		)

		tunnel, err := c.OpenTunnel(ctx, ip)
		if err != nil {
			cancel()
			return nil, err
		}

		cancel()

		reset, err := setupTunHandler(c, iface)(ip, mask, gw)
		if err != nil {
			tunnel(ctx) // use a canceled ctx to terminate the tunnel
			return nil, fmt.Errorf("setup tunnel error: %v", err)
		}

		return func(ctx context.Context) error {
			defer reset()
			return tunnel(ctx)
		}, nil
	}

	tunneld := func(ctx context.Context, ipVersion int) {
		firstRound := true
		for {
			tunnel, err := openTunnel(ctx, ipVersion)
			if ctx.Err() != nil {
				break
			}
			if err != nil {
				log.Errorf("failed to open IPv%d tunnel: %v", ipVersion, err)
				if ae, ok := err.(*camo.ClientAPIError); ok {
					if !firstRound || ae.Temporary() {
						goto RETRY
					}
				}
				break
			}

			log.Infof("IPv%d tunnel opened", ipVersion)

			err = tunnel(ctx)
			if ctx.Err() != nil {
				log.Infof("IPv%d tunnel closed", ipVersion)
				break
			}
			log.Errorf("IPv%d tunnel closed: %v", ipVersion, err)

			firstRound = false

		RETRY:
			if ctx.Err() != nil {
				break
			}
			// TODO exponential backoff
			time.Sleep(1 * time.Second)
		}

		if ctx.Err() == nil {
			log.Errorf("IPv%d tunnel thread exited", ipVersion)
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	exit := cancel

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := c.ServeIface(ctx, iface)
		if ctx.Err() != nil {
			return
		}
		log.Errorf("serve iface exited: %v", err)
		exit()
	}()

	var tunWG sync.WaitGroup

	if *tun4 {
		tunWG.Add(1)
		go func() {
			tunneld(ctx, 4)
			tunWG.Done()
		}()
	}

	if *tun6 {
		tunWG.Add(1)
		go func() {
			tunneld(ctx, 6)
			tunWG.Done()
		}()
	}

	tunWG.Wait()
	if ctx.Err() == nil {
		exit()
	}

	wg.Wait()
	return
}

func debugHTTPServer() {
	err := http.ListenAndServe(*debugHTTP, nil)
	if err != http.ErrServerClosed {
		log.Errorf("debug http server exited: %v", err)
	}
}
