#!/bin/sh

cd ../../goerr/
go build
if [ $? -ne 0 ];then exit 1 ;fi
cd ../test/err1/
../../goerr/goerr -f main.go -e err.go merge
rm err.go
go fmt main.go
cat -n main.go
# go build && ./err1
git checkout  main.go err.go
