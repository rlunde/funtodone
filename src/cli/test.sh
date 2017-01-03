#!/bin/bash

# use ad-hoc testing for now -- later, maybe use aruba:
# https://github.com/cucumber/aruba

# all commands can use:
# --config string   config file (default is $HOME/.funtodone-cli.yaml)

../funtodone-cli create user --email test@example.com --name "test user" --password guessWhat
