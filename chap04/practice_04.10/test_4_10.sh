#!/bin/bash

# define suffix-less file name
FILE_NAME=4_10

# build executable-program
go build ${FILE_NAME}.go

# execute
#./${FILE_NAME} repo:golang/go is:open json
./${FILE_NAME} repo:golang/go is:open json decoder

#----- [EOF] -----#

