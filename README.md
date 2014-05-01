# Canary

Canary is a dynamic DNS update utility written using Sails.js.


### Project Dependencies

The following tools need to be in place for a working development environment.

* VirtualBox 4.3.10
* Vagrant 1.5.4
* Ansible 1.5.5


### Development Environment

A Vagrant configuration file has been provided to set up a development instance in VirtualBox. An Ansible playbook has also been provided to aid in the initial setup of the instance.

Clone the repository and navigate to the project directory.

First, ensure that the permissions on the Vagrant private key are set correctly and launch the instance.

    $ chmod 600 Vagrant.pem
    $ vagrant up

Next, configure the Vagrant instance for development using Ansible. Be patient, the process could take a while to finish.

    $ ansible-playbook site.yml

Once configured, connect to the vagrant instance and get to work. The application source is mounted in the vagrant user's home directory.

    $ vagrant ssh
    $ cd canary

Run the application in development mode.

    $ sails lift


### Testing

Run the application tests.

    $ sails run tests

Generate the application coverage report.

    $ sails run coverage


### Platforms

Development has occurred on the following platforms:

* Ubuntu 12.04 LTS (Precise)
