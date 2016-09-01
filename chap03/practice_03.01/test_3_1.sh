#!/bin/bash

# define suffix-less file name
FILE_NAME=3_1

# build executable-program
go build ${FILE_NAME}.go

# create html-file
./${FILE_NAME} > ${FILE_NAME}.html

# use cygstart instead of xdg-open command on Cygwin-env.
cygstart ${FILE_NAME}.html

#----- [EOF] -----#

