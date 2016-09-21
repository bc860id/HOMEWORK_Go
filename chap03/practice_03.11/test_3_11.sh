#!/bin/bash

# define suffix-less file name
FILE_NAME=3_11

# build executable-program
go build ${FILE_NAME}.go

# test
./${FILE_NAME} 12345678.012345679
./${FILE_NAME} 12345678
./${FILE_NAME} -12345678.012345679
./${FILE_NAME} 0.012345679
./${FILE_NAME} -0.012345679
./${FILE_NAME} 0
./${FILE_NAME} -0

#----- [EOF] -----#

