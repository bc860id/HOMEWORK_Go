#!/bin/bash

# define suffix-less file name
FILE_NAME=3_6_ave_${1}

# build executable-program
go build ${FILE_NAME}.go

# create png-file
./${FILE_NAME} > ${FILE_NAME}.png

# open created png-file
cygstart ${FILE_NAME}.png

#----- [EOF] -----#

