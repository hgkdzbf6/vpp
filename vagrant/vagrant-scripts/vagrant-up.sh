#!/bin/bash

set -eo pipefail

export K8S_NODE_OS=${K8S_NODE_OS:-ubuntu}
export K8S_NODES=${K8S_NODES:-1}
export VAGRANT_DEFAULT_PROVIDER=${VAGRANT_DEFAULT_PROVIDER:-virtualbox}

# Default values for environment deploymen
export VAGRANT_VAGRANTFILE=../Vagrantfile-prod
DEV_ENV="false"
TEST_ENV="true"

# Override defaults from arguments
while [ "$1" != "" ]; do
    case $1 in
        -d | --dev-env )
            DEV_ENV="true"
            echo "Using development environment"
            ;;
        -t | --test-env )
            echo "Using testing environment"
            ;;
        * )
            echo "Invalid parameter: "$1
            exit 1
    esac
    shift
done

if [ "${DEV_ENV}" == "true" ]
then
    export VAGRANT_VAGRANTFILE=../Vagrantfile-dev; vagrant up
fi
  
if [ "${TEST_ENV}" == "true" ]
then
    echo "Using testing environment"
    vagrant up 
fi
