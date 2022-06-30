package applicationv1

import (
	"fmt"
	"os"

	"github.com/tamama/tamama-cloud-deployer-machine/resource"
	"github.com/urfave/cli/v2"
)

var (
	flag_version *bool
)

func (_self *ApplicationV1) _NewApp() (app *cli.App) {
	app = &cli.App{
		Name:        "tamama-cloud-deployer-machine",
		Usage:       "Tamama Cloud Deployer Machine",
		Description: "Tamama Cloud Deployer Machine",
		Authors: []*cli.Author{
			{
				Name:  "Tama MA",
				Email: "tamama@tamama.io",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Value:       false,
				Usage:       "print version",
				Destination: flag_version,
			},
		},
		Action: _self._RunAction,
	}

	app.Commands = []*cli.Command{
		_self._NewCommandVagrant(),
	}

	return app
}

func (_self *ApplicationV1) _RunAction(ctx *cli.Context) (err error) {
	switch {
	case ctx.Bool("version"):
		err = _self._RunAction_Version(ctx)
	default:
		cli.ShowAppHelp(ctx)
	}
	os.Exit(0)
	return err
}

func (_self *ApplicationV1) _RunAction_Version(ctx *cli.Context) (err error) {
	fmt.Printf("tamama-cloud-deployer-machine version %s\n", resource.Version)
	return nil
}
