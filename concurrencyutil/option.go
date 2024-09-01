package concurrencyutil

import "context"

type Options struct {
	Context context.Context
	Limit   int
}

type Option func(o *Options)

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		if ctx == nil {
			o.Context = context.Background()
		}
		o.Context = ctx
	}
}

func WithLimit(n int) Option {
	return func(o *Options) {
		o.Limit = n
	}
}
