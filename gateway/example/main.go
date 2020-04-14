package main

import (
	"net/http"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/micro/v2/gateway/cmd"
	"github.com/micro/micro/v2/gateway/router"
)

func main() {
	cmd.Init(
		// router services filter
		router.WithOption(
			router.WithFilter(func(req *http.Request) selector.Filter {
				return selector.FilterLabel("key", "val")
			}),
		),
		// micro option
	)
}
