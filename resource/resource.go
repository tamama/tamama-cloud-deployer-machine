package resource

import (
	_ "embed"
)

var (
	//go:embed profile/vagrant/single-node/rhel-9/Vagrantfile
	ProfileVagrantSingleNodeRhel9Vagrantfile []byte

	//go:embed profile/vagrant/single-node/rhel-9-test/Vagrantfile
	ProfileVagrantSingleNodeRhel9TestVagrantfile []byte

	//go:embed profile/vagrant/single-node/ubuntu-2204/Vagrantfile
	ProfileVagrantSingleNodeUbuntu2204Vagrantfile []byte

	//go:embed version/VERSION
	Version []byte
)
