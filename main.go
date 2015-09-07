package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"
	"flag"

	"github.com/golang/glog"

	"github.com/karlkfi/probe/http"
)

const validTimeUnits = "Valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\"."

var (
	flagHTTP    = flag.Bool("http", false, "Use HTTP probing")
	flagTimeout = flag.Duration("timeout", -1, "Timeout duration. "+validTimeUnits)
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

	if !*flagHTTP {
		fmt.Fprint(os.Stderr, "Error: No probe type specified - Expected \"--http\"\n")
		flag.Usage()
		os.Exit(2)
	}

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

	address := args[0]

	addrURL, err := url.ParseRequestURI(address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid address %q - Expected valid URL\n", address)
		os.Exit(2)
	}

	client := http.NewInsecureClient()

	if *flagTimeout >= 0 {
		client.Timeout = *flagTimeout
	}

	prober := http.NewProber(client)

	probeErr := prober.Probe(addrURL)
	if probeErr != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", probeErr.Error())
		os.Exit(1)
	}
}
