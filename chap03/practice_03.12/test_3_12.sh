#!/bin/bash

# define suffix-less file name
FILE_NAME=3_12

# build executable-program
go build ${FILE_NAME}.go

# test
./${FILE_NAME} 12345 51423
./${FILE_NAME} 12345 51428
./${FILE_NAME} 123457 51428
./${FILE_NAME} 12��345 5142��3
./${FILE_NAME} 12345�� ��51423
./${FILE_NAME} 12345�� ��51423
./${FILE_NAME} 12345�� 51423��
./${FILE_NAME} �� ��
./${FILE_NAME} �~�� ���~

#----- [EOF] -----#

