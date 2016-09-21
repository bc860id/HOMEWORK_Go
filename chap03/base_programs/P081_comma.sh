#!/bin/bash

# define suffix-less file name
FILE_NAME=P081_comma

# build executable-program
go build ${FILE_NAME}.go

# test
./${FILE_NAME} 12345678ABCDE

#----- [EOF] -----#

