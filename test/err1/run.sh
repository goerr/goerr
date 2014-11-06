#!/bin/sh

../../goerr/goerr -f main.go -e err.go merge
rm err.go
go build && ./err1
git checkout  main.go err.go
