#!/bin/bash

# define suffix-less file name
FILE_NAME=4_8

# build executable-program
go build ${FILE_NAME}.go

# execute
cat ${FILE_NAME}_sample.txt | ./${FILE_NAME} > ${FILE_NAME}.txt

#----- [EOF] -----#

