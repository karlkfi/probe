package http

type prober struct {
	getter Getter
}

func NewProber(getter Getter) *prober {
	return &prober{
		getter: getter,
	}
}

func (pr *prober) Probe(address string) error {
	return Probe(address, pr.getter)
}
