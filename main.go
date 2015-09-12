package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/golang/glog"

	"github.com/karlkfi/probe/http"
	"github.com/karlkfi/probe/tcp"
	"time"
)

func main() {
	flagSet := flag.CommandLine
	c := parseFlags(flagSet)

	defer glog.Flush()
	glog.V(1).Info("Executing: ", strings.Join(os.Args, " "))

	// non-flag args
	args := flagSet.Args()
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, "Error: No address specified\n")
		flagSet.Usage()
		os.Exit(2)
	}

	if len(args) > 1 {
		fmt.Fprint(os.Stderr, "Error: Too many arguments specified\n")
		flagSet.Usage()
		os.Exit(2)
	}

	addrArg := args[0]

	addrURL, err := url.ParseRequestURI(addrArg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid address %q - Expected valid URL\n", addrArg)
		os.Exit(2)
	}

	if *c.maxAttempts == 0 {
		fmt.Fprintf(os.Stderr, "Error: Invalid attempts %q - Expected an int > 0 or exactly -1 (unlimited)\n", addrArg)
		os.Exit(2)
	}

	if *c.retryDelay < 0 {
		fmt.Fprintf(os.Stderr, "Error: Invalid delay %q - Expected a duration >= 0\n", addrArg)
		os.Exit(2)
	}

	var prober Prober
	var address string

	switch addrURL.Scheme {
	case SchemeTCP:
		dialer := tcp.NewDialer()

		if *c.timeout >= 0 {
			dialer.Timeout = *c.timeout
		}

		prober = tcp.NewProber(dialer)
		address = addrURL.Host
	case SchemeHTTP:
		fallthrough
	case SchemeHTTPS:
		client := http.NewInsecureClient()

		if *c.timeout >= 0 {
			client.Timeout = *c.timeout
		}

		prober = http.NewProber(client)
		address = addrURL.String()
	default:
		fmt.Fprint(os.Stderr, "Error: No probable address scheme specified - Expected \"tcp\", \"http\", or \"https\"\n")
		flagSet.Usage()
		os.Exit(2)
	}

	err = makeAttempts(prober, address, *c.maxAttempts, *c.retryDelay)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func makeAttempts(prober Prober, address string, maxAttempts int, retryDelay time.Duration) error {
	probeErr := prober.Probe(address)
	attemptsMade := 1
	glog.V(2).Infof("Attempt %d Failed: %v", attemptsMade, probeErr)

	for probeErr != nil && (maxAttempts < 0 || maxAttempts > attemptsMade) {
		time.Sleep(retryDelay)
		probeErr = prober.Probe(address)
		attemptsMade++
		glog.V(2).Infof("Attempt %d Failed: %v", attemptsMade, probeErr)
	}

	if probeErr != nil {
		return probeErr
	}
	return nil
}
