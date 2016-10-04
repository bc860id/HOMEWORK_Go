#!/bin/bash

# define suffix-less file name
FILE_NAME=4_9

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} ${FILE_NAME}_sample.txt > ${FILE_NAME}.txt

#----- [EOF] -----#

