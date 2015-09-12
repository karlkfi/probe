package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

const (
	validTimeUnits = "Valid time units: ns, us (or µs), ms, s, m, h."

	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeTCP   = "tcp"
)

type config struct {
	timeout  *time.Duration
	maxAttempts *int
	retryDelay *time.Duration
}

func (c *config) addflags(s *flag.FlagSet) {
	timeout := s.Duration("timeout", -1, "Timeout duration. "+validTimeUnits)
	s.DurationVar(timeout, "t", -1, "Shortcut for --timeout")
	c.timeout = timeout

	maxAttempts := s.Int("max-attempts", 1, "Maximum number of attempts to make (default=1, unlimitted=-1)")
	s.IntVar(maxAttempts, "a", 1, "Shortcut for --max-attempts")
	c.maxAttempts = maxAttempts

	delay := s.Duration("retry-delay", 1*time.Second, "Delay between attempts. "+validTimeUnits)
	s.DurationVar(delay, "d", 1*time.Second, "Shortcut for --retry-delay")
	c.retryDelay = delay
}

func usage(s *flag.FlagSet) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <address>\n", path.Base(os.Args[0]))
		s.PrintDefaults()
	}
}

func parseFlags(s *flag.FlagSet) *config {
	c := &config{}
	c.addflags(s)
	s.Usage = usage(s)
	s.Parse(os.Args[1:])
	return c
}
