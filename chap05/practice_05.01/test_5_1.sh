#!/bin/bash
go build P018_fetch.go
go build 5_1.go
./P018_fetch https://golang.org | ./5_1

