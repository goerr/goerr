#!/bin/sh

cd ../../goerr/
go build
if [ $? -ne 0 ];then exit 1 ;fi
cd ../test/err0/
../../goerr/goerr -f main.go -e err.go merge
rm err.go
go fmt main.go
cat -n main.go
go build && ./err0
git checkout  main.go err.go
