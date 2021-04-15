#!/usr/bin/env bash

GOPATH="$(go env GOPATH)"
AVALANCHE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd ) # Directory above this script
AVALANCHEGO_VER="v1.3.2" # Current version of AvalancheGo to use
CORETH_VER="v0.4.0-rc.7" # Current version of Coreth to use
PREV_AVALANCHEGO_VER="plugin_mode" # Last release of AvalancheGo before the database upgrade
CORETH_PATH="$GOPATH/pkg/mod/github.com/ava-labs/coreth@$CORETH_VER"

BUILD_DIR=$AVALANCHE_PATH/build/avalanchego-$AVALANCHEGO_VER # Where AvalancheGo binary goes
PLUGIN_DIR="$BUILD_DIR/plugins" # Where plugin binaries (namely coreth) go
EVM_PATH="$PLUGIN_DIR/evm"
AVALANCHEGO_INNER_PATH="$BUILD_DIR/avalanchego-inner"
BINARY_MANAGER_PATH="$AVALANCHE_PATH/build/avalanchego"

PREV_BUILD_DIR="$AVALANCHE_PATH/build/avalanchego-$PREV_AVALANCHEGO_VER" # Where AvalancheGo binary goes
PREV_PLUGIN_DIR="$PREV_BUILD_DIR/plugins"
PREV_AVALANCHEGO_INNER_PATH="$PREV_BUILD_DIR/avalanchego-inner"
