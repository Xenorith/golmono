#!/bin/bash
set -ex

REPO_HOME="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"
cd $REPO_HOME

# assume all architectures are amd64
OSARCH="linux-amd64"
if [[ $(uname -s) == "Darwin" ]]; then
    OSARCH="darwin-amd64"
fi

# copies binaries built using ./godlew build into a specific product's bin/ directory
# first argument is name of product folder to copy into
PRODUCT_NAME=$1
rm -rf $PRODUCT_NAME/bin
mkdir -p $PRODUCT_NAME/bin
# subsequent arguments are names of binaries to copy
for BINARY_NAME in ${@:2}; do
    cp out/build/$BINARY_NAME/unspecified/$OSARCH/$BINARY_NAME $PRODUCT_NAME/bin/$BINARY_NAME
done
