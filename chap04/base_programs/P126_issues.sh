#!/bin/bash

# define suffix-less file name
FILE_NAME=P126_issues

# build executable-program
go build ${FILE_NAME}.go

# execute
./${FILE_NAME} repo:golang/go is:open json decoder

#----- [EOF] -----#

