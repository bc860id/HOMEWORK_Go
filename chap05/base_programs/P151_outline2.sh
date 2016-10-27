#!/bin/bash

# define suffix-less file name
FILE_NAME=P151_outline2
FILE_FETCH=P018_fetch

# build executable-program
go build ${FILE_NAME}.go
go build ${FILE_FETCH}.go

# execute program
${FILE_FETCH} http://gopl.io | ${FILE_NAME}

#----- [EOF] -----#

