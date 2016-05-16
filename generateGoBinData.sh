#!/bin/bash

# Create bindata.go like this:
$GOPATH/bin/go-bindata -prefix "template/" template/...
