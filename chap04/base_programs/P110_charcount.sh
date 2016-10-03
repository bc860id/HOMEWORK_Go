#!/bin/bash

# define suffix-less file name
FILE_NAME=P110_charcount

# build executable-program
go build ${FILE_NAME}.go

# execute
cat ${FILE_NAME}_sample.txt | ./${FILE_NAME} > ${FILE_NAME}_result.txt

#----- [EOF] -----#

