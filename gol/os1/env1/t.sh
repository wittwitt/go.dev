#!/bin/bash

export GODEV_TEST_VAR1="v1value"

# not env
GODEV_TEST_VAR2="v2value"

./env1


# as env
GODEV_TEST_VAR2="v2value22222" ./env1