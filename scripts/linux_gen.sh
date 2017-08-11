#!/bin/sh
echo "*******************"
echo "Builiding for linux"
echo "*******************"

cd ../
pwd
export GOOS=linux
#set GOOS=linux # in command manuly do this - in bash scripts use above export
go env
go install
go build -o ./build/sun_linux