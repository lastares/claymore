package concurrencyutil

import (
	"errors"

	"golang.org/x/sync/errgroup"
)

func NewWg(funcs []func() error, ops ...Option) error {
	if len(funcs) == 0 {
		return nil
	}
	options := Options{}
	for _, op := range ops {
		op(&options)
	}
	g := &errgroup.Group{}
	if options.Context != nil {
		g, _ = errgroup.WithContext(options.Context)
	}
	if options.Limit > 0 {
		g.SetLimit(options.Limit)
	}
	for _, f := range funcs {
		if f == nil {
			return errors.New("func is nil")
		}
		g.Go(f)
	}
	return g.Wait()
}
