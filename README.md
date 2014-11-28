# Canary

Canary is a dynamic DNS update utility written in Node.js and using the Sails.js framework.

![Canary](http://f870d4d1e656b386e423-644dd79aec9558152263f90d87ffdef9.r61.cf1.rackcdn.com/canary-sm.png "Canary")

### Project Dependencies

The following tools need to be in place for a working development environment.

* VirtualBox
* Vagrant
* Ansible


### Development Environment

A Vagrant configuration file has been provided to set up a development instance in VirtualBox. An Ansible playbook has also been provided to aid in the initial setup of the instance.

Clone the repository and navigate to the project directory.

First, launch the instance.

    $ vagrant up

Next, configure the Vagrant instance for development using Ansible. Be patient, the process could take a while to finish.

    $ ansible-playbook deploy.yml

Once configured, connect to the vagrant instance and get to work. The application source is mounted in the vagrant user's home directory.

    $ vagrant ssh
    $ cd /usr/src/canary

Run the application in development mode.

    $ sails lift


### Testing

Run the application tests.

    $ sails run tests

Generate the application coverage report.

    $ sails run coverage


### Platforms

Development has occurred on the following platforms:

* Ubuntu 14.04 LTS (Trusty)
