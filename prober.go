package main

// Prober is satisfied by http.prober and tcp.prober
type Prober interface {
	Probe(address string) error
}
