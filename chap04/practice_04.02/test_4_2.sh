#!/bin/bash

# define suffix-less file name
FILE_NAME=4_2

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} X
./${FILE_NAME} -s=256 X
./${FILE_NAME} -s=384 X
./${FILE_NAME} -s=512 X
./${FILE_NAME} -s=355 X

#----- [EOF] -----#

