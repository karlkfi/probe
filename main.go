package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/golang/glog"

	"github.com/karlkfi/probe/http"
	"github.com/karlkfi/probe/tcp"
)

const (
	validTimeUnits = "Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."

	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeTCP   = "tcp"
)

var (
	flagTimeout         = flag.Duration("timeout", -1, "Timeout duration. "+validTimeUnits)
	flagTimeoutShortcut = flag.Duration("t", -1, "Shortcut for --timeout")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [flags] <address>\n", path.Base(os.Args[0]))
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	defer glog.Flush()
	glog.V(1).Info("Executing: ", strings.Join(os.Args, " "))

	// non-flag args
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "Error: No address specified\n")
		flag.Usage()
		os.Exit(2)
	}

	if len(args) > 1 {
		fmt.Fprint(os.Stderr, "Error: Too many arguments specified\n")
		flag.Usage()
		os.Exit(2)
	}

	addrArg := args[0]

	addrURL, err := url.ParseRequestURI(addrArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid address %q - Expected valid URL\n", addrArg)
		os.Exit(2)
	}

	if *flagTimeout >= 0 && *flagTimeoutShortcut >= 0 {
		fmt.Fprint(os.Stderr, "Error: Both \"--timeout\" and its shortcut \"-t\" specified - Expected at most one\n")
		flag.Usage()
		os.Exit(2)
	}

	var prober Prober
	var address string

	switch addrURL.Scheme {
	case SchemeTCP:
		dialer := tcp.NewDialer()

		if *flagTimeout >= 0 {
			dialer.Timeout = *flagTimeout
		} else if *flagTimeoutShortcut >= 0 {
			dialer.Timeout = *flagTimeout
		}

		prober = tcp.NewProber(dialer)
		address = addrURL.Host
	case SchemeHTTP:
		fallthrough
	case SchemeHTTPS:
		client := http.NewInsecureClient()

		if *flagTimeout >= 0 {
			client.Timeout = *flagTimeout
		} else if *flagTimeoutShortcut >= 0 {
			client.Timeout = *flagTimeout
		}

		prober = http.NewProber(client)
		address = addrURL.String()
	default:
		fmt.Fprint(os.Stderr, "Error: No probable address scheme specified - Expected \"tcp\", \"http\", or \"https\"\n")
		flag.Usage()
		os.Exit(2)
	}

	probeErr := prober.Probe(address)
	if probeErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", probeErr.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
