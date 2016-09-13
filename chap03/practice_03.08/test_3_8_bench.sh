#!/bin/bash

#set -x

# set variable
FILE_NAME_BASE="3_8_"
APPEND="_test"
OPTION="-bench=ForMeasure -benchmem"
REDIRECT="/dev/null"

# benchmark each program
go test ${FILE_NAME_BASE}complex64${APPEND}.go ${OPTION} 2> ${REDIRECT}
go test ${FILE_NAME_BASE}complex128${APPEND}.go ${OPTION} 2> ${REDIRECT}
go test ${FILE_NAME_BASE}bigFloat${APPEND}.go ${OPTION} 2> ${REDIRECT}
go test ${FILE_NAME_BASE}bigRat${APPEND}.go ${OPTION} 2> ${REDIRECT}

#----- [EOF] -----#

