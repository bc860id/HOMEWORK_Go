#!/bin/bash

# define suffix-less file name
FILE_NAME=3_12

# build executable-program
go build ${FILE_NAME}.go

# test
./${FILE_NAME} 12345 51423
./${FILE_NAME} 12345 51428
./${FILE_NAME} 123457 51428
./${FILE_NAME} 12è„345 5142è„3
./${FILE_NAME} 12345Å® Å©51423
./${FILE_NAME} 12345â∫ â∫51423
./${FILE_NAME} 12345Å® 51423Å®
./${FILE_NAME} Åù Åù
./${FILE_NAME} Å~Åù ÅùÅ~

#----- [EOF] -----#

