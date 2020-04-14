package registry

import (
	"net/http"

	"github.com/micro/go-micro/v2/api/router"
	"github.com/micro/go-micro/v2/client/selector"
)

type Options struct {
	router.Options
	routerOpts []router.Option

	Filters []Filter
}

type Option func(o *Options)

type Filter func(req *http.Request) selector.Filter

func NewOptions(opts ...Option) Options {
	options := Options{}

	for _, o := range opts {
		if o == nil {
			continue
		}
		o(&options)
	}

	options.Options = router.NewOptions(options.routerOpts...)

	return options
}

// Router options, override by NewRouter()'s router.Option
func WithRouterOption(opt ...router.Option) Option {
	return func(o *Options) {
		o.routerOpts = append(o.routerOpts, opt...)
	}
}

func WithFilter(f ...Filter) Option {
	return func(o *Options) {
		o.Filters = append(o.Filters, f...)
	}
}

func WithOption(opts ...Option) Option {
	return func(o *Options) {
		for _, opt := range opts {
			if opt == nil {
				continue
			}
			opt(o)
		}
	}
}
