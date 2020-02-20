#!/bin/sh

SCRIPT_DIR=$(readlink -e $(dirname $0))

protoc -I $SCRIPT_DIR/  $SCRIPT_DIR/*.proto --go_out=plugins=grpc:.