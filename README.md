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
tamama-cloud-deployer-machine vagrant --init --force
```

Start our VM:
```
vagrant up
```

## (TL;DR) Notes for developers

### Option 1 - The quick and dirty way ...

Navigate to the source folder of your choice:
```
mkdir -p ~/workspace/src/github.com/tamama/tamama-cloud-deployer-machine
cd ~/workspace/src/github.com/tamama/tamama-cloud-deployer-machine
```
Git clone repo:
```
git clone https://github.com/tamama/tamama-cloud-deployer-machine .
```
Switch to `develop` branch:
```
git checkout develop
```
Build from source:
```
make init
make build
```  
Smoke-check - version:
```
/tmp/tamama-cloud-deployer-machine --version
```

### Option 2 - The professional way ...

Fork this Git repo under your namespace, say:
```
https://github.com/mamata/tamama-cloud-deployer-machine
```
Navigate to the source folder of your choice:
```
mkdir -p ~/workspace/src/github.com/mamata/tamama-cloud-deployer-machine
cd ~/workspace/src/github.com/mamata/tamama-cloud-deployer-machine
```
Git clone repo:
```
git clone https://github.com/mamata/tamama-cloud-deployer-machine .
```
Add the source Git repo to your upstream:
```
git remote add upstream https://github.com/tamama/tamama-cloud-deployer-machine
```
Switch to `develop` branch:
```
git checkout develop
```
Fetch latest changes (upstream Git repo => your Git repo)
```
git fetch upstream
git pull upstream develop
git push origin develop
```
Always develop from your feature branch:
```
git checkout feat/jira-xxxx
git rebase develop
```
(Hint: It is always a good idea to sync up with upstream repo. Otherwise, conflicts would kill you eventually.)

