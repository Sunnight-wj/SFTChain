#!/bin/bash


export DAEMON_NAME=junod
export DAEMON_HOME=$HOME/.juno

export APP_HOME=/home/test/SFTChain
export TMP_COSMOVISOR_DIR=$HOME/tmp-cosmovisor

export DAEMON_ALLOW_DOWNLOAD_BINARIES=false
export DAEMON_RESTART_AFTER_UPGRADE=true

cp $GOPATH/bin/$DAEMON_NAME $TMP_COSMOVISOR_DIR/genesis/bin/
cd $APP_HOME

cp -R $TMP_COSMOVISOR_DIR $DAEMON_HOME



