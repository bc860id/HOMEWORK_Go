#!/bin/bash

# define suffix-less file name
FILE_NAME=4_1

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} X x
./${FILE_NAME} x x
./${FILE_NAME} Ab12CxYz ab
./${FILE_NAME} Ab12CxYz Ab12CxYz

#----- [EOF] -----#

