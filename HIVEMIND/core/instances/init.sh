#!/bin/bash

set -e

cat  ./pbs/header.proto ./pbs/*.txt > merged.proto # merged.proto will be over-written.

protoc -I=. \
	    --go_out . --go_opt paths=source_relative \
	    --go-grpc_out . --go-grpc_opt paths=source_relative \
	    ./merged.proto