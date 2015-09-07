package tcp

type prober struct {
	dialer Dialer
}

func NewProber(dialer Dialer) *prober {
	return &prober{
		dialer: dialer,
	}
}

func (pr *prober) Probe(address string) error {
	return Probe(address, pr.dialer)
}
