#!/bin/bash

# define suffix-less file name
FILE_NAME=4_1

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} X
./${FILE_NAME} x
./${FILE_NAME} Ab12CxYz

#----- [EOF] -----#

