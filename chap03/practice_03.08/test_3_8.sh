#!/bin/bash

# define suffix-less file name
FILE_NAME_BASE=3_8_
#ZOOM=1
ZOOM=4

# build executable-program
go build ${FILE_NAME_BASE}complex64.go
go build ${FILE_NAME_BASE}complex128.go
go build ${FILE_NAME_BASE}bigFloat.go
go build ${FILE_NAME_BASE}bigRat.go

# create png-file
./${FILE_NAME_BASE}complex64 ${ZOOM} > ${FILE_NAME_BASE}complex64.png
./${FILE_NAME_BASE}complex128 ${ZOOM} > ${FILE_NAME_BASE}complex128.png
./${FILE_NAME_BASE}bigFloat ${ZOOM} > ${FILE_NAME_BASE}bigFloat.png
#./${FILE_NAME_BASE}bigRat ${ZOOM} > ${FILE_NAME_BASE}bigRat.png
./${FILE_NAME_BASE}bigRat 1 > ${FILE_NAME_BASE}bigRat.png

# open created png-file
cygstart ${FILE_NAME_BASE}complex64.png
cygstart ${FILE_NAME_BASE}complex128.png
cygstart ${FILE_NAME_BASE}bigFloat.png
cygstart ${FILE_NAME_BASE}bigRat.png

#----- [EOF] -----#

