#!/bin/bash

# define suffix-less file name
FILE_NAME=4_5

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} red orange red green yellow blue pink
./${FILE_NAME} red orange red red green yellow blue pink pink pink
./${FILE_NAME} red orange red red green green green yellow blue pink pink pink
./${FILE_NAME} red orange red red green green green yellow blue pink
./${FILE_NAME} red red red orange orange red red green green green yellow blue pink
./${FILE_NAME} red red red red red

#----- [EOF] -----#

