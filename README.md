# Tamama Cloud Deployer Machine

## Vagrant

### Prerequisites

These host specifications work for my setup:
1. Golang - 1.18
2. Vagrant - 2.2.19
3. VirtualBox - 6.1.30

### Getting started


Download the installer on your host:
```
go install github.com/tamama/tamama-cloud-deployer-machine@v0.2.1
```

Navigate to your folder of choice:
```
mkdir -p ~/workspace/var/vagrant/tamama-machine
cd ~/workspace/var/vagrant/tamama-machine
```

Initialize our Vagrantfile:  
```
tamama-cloud-deployer-machine vagrant --init
```
(Note that this will NOT overwrite your existing Vagrantfile)

Start our VM:
```
vagrant up
```

