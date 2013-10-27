#!/bin/bash

BASE_DIR=$(dirname "$(readlink -f $0)")
SRC_DIR="$BASE_DIR/src"

cd "$SRC_DIR"
go fmt && go build -o ../parseUnihan
exit $?
