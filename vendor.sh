#!/usr/bin/env bash
# vendor init
rm -rf $GOPATH/src/github.com/Loopring/relay/vendor
govendor init

# vendor add external libraries
govendor add +external