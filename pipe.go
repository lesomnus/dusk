package dusk

import (
	"context"
	"path/filepath"
)

type Step interface {
	Do(ctx context.Context, reg Registry, dst string) error
}

func Seq(steps ...Step) Step {
	return StepFunc(func(ctx context.Context, reg Registry, dst string) error {
		for _, step := range steps {
			if err := step.Do(ctx, reg, dst); err != nil {
				return err
			}
		}

		return nil
	})
}

type StepFunc func(ctx context.Context, reg Registry, dst string) error

func (f StepFunc) Do(ctx context.Context, reg Registry, dst string) error {
	return f(ctx, reg, dst)
}

type Sink interface {
	Do(desc Descriptor, dst string) ([]Exec, error)
}

type SinkFunc func(desc Descriptor, dst string) ([]Exec, error)

func (f SinkFunc) Do(desc Descriptor, dst string) ([]Exec, error) {
	return f(desc, dst)
}

type Piped struct {
	desc descriptor
	step Step
}

func Pipe(name, version string, steps ...Step) *Piped {
	return &Piped{
		desc: descriptor{name, version},
		step: Seq(steps...),
	}
}

func (p *Piped) Sink(sink SinkFunc) Package {
	return &piped{
		descriptor: p.desc,

		step: p.step,
		sink: sink,
	}
}

func (p *Piped) Pick(rpath string) Package {
	return p.Sink(SinkFunc(func(desc Descriptor, dst string) ([]Exec, error) {
		p := filepath.Join(dst, rpath)
		f := Static(desc.Name(), desc.Version(), p)
		return []Exec{f}, nil
	}))
}

type piped struct {
	descriptor
	step Step
	sink Sink
}

func (p *piped) Unpack(ctx context.Context, reg Registry, dst string) ([]Exec, error) {
	if err := p.step.Do(ctx, reg, dst); err != nil {
		return nil, err
	}

	return p.sink.Do(&p.descriptor, dst)
}
