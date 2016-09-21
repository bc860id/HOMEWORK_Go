#!/bin/bash

# define suffix-less file name
FILE_NAME=3_10

# build executable-program
go build ${FILE_NAME}.go

# test
./${FILE_NAME} 12345678ABCDEFGHI

#----- [EOF] -----#

