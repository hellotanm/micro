package main

import (
	"net/http"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/micro/v2/gateway/cmd"
	"github.com/micro/micro/v2/gateway/router"
)

func main() {
	// Router services filter
	opt := router.WithOption(
		router.WithFilter(func(req *http.Request) selector.Filter {
			return selector.FilterLabel("key", "val")
		}),
	)

	cmd.Init(
		opt,
		// micro option
	)
}
