#!/bin/bash

# define suffix-less file name
FILE_NAME=4_11

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} create abc
./${FILE_NAME} update abc 1 open

#----- [EOF] -----#

