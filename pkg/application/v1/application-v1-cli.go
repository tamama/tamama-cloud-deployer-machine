package applicationv1

import (
	"os"

	"github.com/urfave/cli/v2"
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
		Flags: []cli.Flag{},
		Action: _self._RunAction,
	}

	app.Commands = []*cli.Command{
		_self._NewCommandVagrant(),
	}

	return app
}

func (_self *ApplicationV1) _RunAction(ctx *cli.Context) (err error) {
	switch {
	default:
		cli.ShowAppHelp(ctx)
	}
	os.Exit(0)
	return err
}
