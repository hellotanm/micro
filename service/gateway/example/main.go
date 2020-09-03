package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/micro/micro/v3/cmd"

	"github.com/micro/micro/v3/service/gateway"
	gwRouter "github.com/micro/micro/v3/service/gateway/router"
)

func main() {
	cmd.Register(
		gateway.Commands(
			gwRouter.WithFilter(func(req *http.Request) gwRouter.ServiceFilter {
				return gwRouter.FilterLabel("key", "val")
			}),
		),
	)

	if err := cmd.DefaultCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
