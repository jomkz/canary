# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
	config.vm.box = "ubuntu/trusty64"
	config.vm.hostname = "canary"
	config.vm.network "private_network", ip: "192.168.9.10"
	config.vm.synced_folder "canary", "/usr/src/canary"

	config.vm.provider "virtualbox" do |vbx|
		vbx.customize ["modifyvm", :id, "--memory", "1024"]
		vbx.customize ["modifyvm", :id, "--nic2", "hostonly"]
	end
end
