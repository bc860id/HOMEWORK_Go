#!/bin/bash

# define suffix-less file name
FILE_NAME=P068_mandelbrot

# build executable-program
go build ${FILE_NAME}.go

# create png-file
./${FILE_NAME} > ${FILE_NAME}.png

# open created png-file
cygstart ${FILE_NAME}.png

#----- [EOF] -----#

