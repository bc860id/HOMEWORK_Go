#!/bin/bash

# define suffix-less file name
FILE_NAME=P096_reverse

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME}

#----- [EOF] -----#

