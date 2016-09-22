#!/bin/bash

# define suffix-less file name
FILE_NAME=4_3

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME}

#----- [EOF] -----#

