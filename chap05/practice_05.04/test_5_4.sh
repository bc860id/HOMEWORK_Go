#!/bin/bash
FILE_NAME=5_4
go build P018_fetch.go
go build ${FILE_NAME}.go

if [ $? -ne 0 ]; then
	echo build failed.
	exit
fi

./P018_fetch https://golang.org | ./${FILE_NAME}

