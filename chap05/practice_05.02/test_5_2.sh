#!/bin/bash
go build P018_fetch.go
go build 5_2.go
./P018_fetch https://golang.org | ./5_2

