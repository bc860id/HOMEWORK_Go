#!/bin/bash
go build P018_fetch.go
go build 5_3.go

if [ $? -ne 0 ]; then
	echo build failed.
	exit
fi

./P018_fetch https://golang.org | ./5_3

