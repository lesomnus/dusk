package dusk

type Descriptor interface {
	Name() string
	Version() string
}

type descriptor struct {
	name    string
	version string
}

func (d *descriptor) Name() string {
	return d.name
}

func (d *descriptor) Version() string {
	return d.version
}
