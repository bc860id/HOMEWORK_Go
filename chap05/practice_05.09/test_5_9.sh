#!/bin/bash

# define suffix-less file name
FILE_NAME=5_9

# build executable-program
go build ${FILE_NAME}.go

# execute program
${FILE_NAME} abc \$135 bnnxy \$o1p2q3rs

#----- [EOF] -----#

