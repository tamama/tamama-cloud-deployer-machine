package applicationv1

import (
	"fmt"
	"os"

	"github.com/tamama/tamama-cloud-deployer-machine/resource"
	"github.com/urfave/cli/v2"
)

var (
	flag_vagrant_init_single_node_ubuntu_2204 *bool
)

func (_self *ApplicationV1) _NewCommandVagrant() (command *cli.Command) {
	command = &cli.Command{
		Name:        "vagrant",
		Aliases:     []string{},
		Usage:       "Selects vagrant-machine mode",
		Description: "Selects vagrant-machine mode",
		//Hidden:      true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "init-single-node-ubuntu-2204",
				Aliases:     []string{"init"},
				Value:       false,
				Usage:       "initializes single-node vagrant-machine w/ Ubuntu-2204",
				Destination: flag_vagrant_init_single_node_ubuntu_2204,
			},
		},
		Action: _self._RunCommandActionVagrant,
	}
	return command
}

func (_self *ApplicationV1) _RunCommandActionVagrant(ctx *cli.Context) (err error) {
	switch {
	case ctx.Bool("init-single-node-ubuntu-2204"):
		err = _self._RunCommandActionVagrantInitSingleNodeUbuntu2204(ctx)
	default:
		cli.ShowSubcommandHelp(ctx)
	}
	os.Exit(0)
	return err
}

func (_self *ApplicationV1) _RunCommandActionVagrantInitSingleNodeUbuntu2204(ctx *cli.Context) error {
	fmt.Printf("%s", string(resource.ProfileVagrantSingleNodeUbuntu2204Vagrantfile))
	
	os.Exit(0)
	return nil
}
