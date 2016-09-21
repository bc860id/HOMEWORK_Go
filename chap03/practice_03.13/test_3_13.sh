#!/bin/bash

# define suffix-less file name
FILE_NAME=3_13

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME}

#----- [EOF] -----#

