package cmd

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/cmd"
	mcmd "github.com/micro/micro/v2/cmd"
	api "github.com/micro/micro/v2/gateway"
	"github.com/micro/micro/v2/gateway/router/registry"
)

var (
	GitCommit string
	GitTag    string
	BuildDate string

	name        = "micro"
	description = "A microservice runtime"
	version     = "latest"
)

func buildVersion() string {
	microVersion := version

	if GitTag != "" {
		microVersion = GitTag
	}

	if GitCommit != "" {
		microVersion += fmt.Sprintf("-%s", GitCommit)
	}

	if BuildDate != "" {
		microVersion += fmt.Sprintf("-%s", BuildDate)
	}

	return microVersion
}

// Init initialised the command line
func Init(opt registry.Option, options ...micro.Option) {
	app := cmd.App()
	app.Commands = append(app.Commands, api.Commands(opt, options...)...)

	mcmd.Setup(app, options...)

	cmd.Init(
		cmd.Name(name),
		cmd.Description(description),
		cmd.Version(buildVersion()),
	)
}
