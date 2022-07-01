package applicationv1

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/tamama/tamama-cloud-deployer-machine/resource"
	"github.com/urfave/cli/v2"
)

var (
	flag_vagrant_init_single_node_rhel_9      *bool
	flag_vagrant_init_single_node_ubuntu_2204 *bool

	flag_vagrant_force *bool
)

func (_self *ApplicationV1) _NewCommandVagrant() (command *cli.Command) {
	command = &cli.Command{
		Name:        "vagrant",
		Aliases:     []string{},
		Usage:       "Selects vagrant-machine mode",
		Description: "Selects vagrant-machine mode",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "init-single-node-rhel-9",
				Aliases:     []string{"init-redhat"},
				Value:       false,
				Usage:       "initializes single-node vagrant-machine w/ Redhat Enterprise Linux 9",
				Destination: flag_vagrant_init_single_node_rhel_9,
			},
			&cli.BoolFlag{
				Name:        "init-single-node-ubuntu-2204",
				Aliases:     []string{"init-ubuntu", "init"},
				Value:       false,
				Usage:       "initializes single-node vagrant-machine w/ Ubuntu-2204",
				Destination: flag_vagrant_init_single_node_ubuntu_2204,
			},
			&cli.BoolFlag{
				Name:        "force",
				Aliases:     []string{"f"},
				Value:       false,
				Usage:       "force override",
				Destination: flag_vagrant_force,
			},
		},
		Action: _self._RunCommandAction_Vagrant,
	}
	return command
}

func (_self *ApplicationV1) _RunCommandAction_Vagrant(ctx *cli.Context) (err error) {
	switch {
	case ctx.Bool("init-single-node-rhel-9"):
		err = _self._RunCommandAction_Vagrant_InitSingleNodeRhel9(ctx)
	case ctx.Bool("init-single-node-ubuntu-2204"):
		err = _self._RunCommandAction_Vagrant_InitSingleNodeUbuntu2204(ctx)
	default:
		cli.ShowSubcommandHelp(ctx)
	}
	os.Exit(0)
	return err
}

func (_self *ApplicationV1) _RunCommandAction_Vagrant_InitSingleNodeRhel9(ctx *cli.Context) error {
	filename_of_target_vagrantfile := fmt.Sprintf("Vagrantfile_single-node-rhel-9_%s", resource.Version)
	filename_of_symlink_vagrantfile := "Vagrantfile"

	if _, err := os.Stat(filename_of_target_vagrantfile); err == nil {
		if !ctx.Bool("force") {
			log.Printf("Already exists %+v in this folder. Please archive/rename existing file and try again.", filename_of_target_vagrantfile)
			os.Exit(1)
			return nil
		}
	}

	err := ioutil.WriteFile(filename_of_target_vagrantfile, resource.ProfileVagrantSingleNodeRhel9Vagrantfile, 0644)
	if err != nil {
		return err
	}

	if _, err := os.Lstat(filename_of_symlink_vagrantfile); err == nil {
		os.Remove(filename_of_symlink_vagrantfile)
	}

	os.Symlink(filename_of_target_vagrantfile, filename_of_symlink_vagrantfile)

	os.Exit(0)
	return nil
}

func (_self *ApplicationV1) _RunCommandAction_Vagrant_InitSingleNodeUbuntu2204(ctx *cli.Context) error {
	filename_of_target_vagrantfile := fmt.Sprintf("Vagrantfile_single-node-ubuntu-2204_%s", resource.Version)
	filename_of_symlink_vagrantfile := "Vagrantfile"

	if _, err := os.Stat(filename_of_target_vagrantfile); err == nil {
		if !ctx.Bool("force") {
			log.Printf("Already exists %+v in this folder. Please archive/rename existing file and try again.", filename_of_target_vagrantfile)
			os.Exit(1)
			return nil
		}
	}

	err := ioutil.WriteFile(filename_of_target_vagrantfile, resource.ProfileVagrantSingleNodeUbuntu2204Vagrantfile, 0644)
	if err != nil {
		return err
	}

	if _, err := os.Lstat(filename_of_symlink_vagrantfile); err == nil {
		os.Remove(filename_of_symlink_vagrantfile)
	}

	os.Symlink(filename_of_target_vagrantfile, filename_of_symlink_vagrantfile)

	os.Exit(0)
	return nil
}
