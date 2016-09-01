#!/bin/bash

case ${1} in
	eggcase | mogle | saddle) ;;
	*) echo "	usage: ${0} eggcase|mogle|saddle"; exit;
esac

# define suffix-less file name
FILE_NAME=3_2_${1}

# build executable-program
go build ${FILE_NAME}.go

# create html-file
./${FILE_NAME} > ${FILE_NAME}.html

# use cygstart instead of xdg-open command on Cygwin-env.
cygstart ${FILE_NAME}.html

#----- [EOF] -----#

