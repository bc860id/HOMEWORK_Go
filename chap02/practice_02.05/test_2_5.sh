#!/bin/bash

go test -bench=PopCount1
go test -bench=PopCount2
go test -bench=PopCount3
go test -bench=PopCount4
go test -bench=.

