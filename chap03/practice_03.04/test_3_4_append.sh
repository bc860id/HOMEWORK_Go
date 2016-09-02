#!/bin/bash

# define suffix-less file name
FILE_NAME=3_4_append

# build executable-program
go build ${FILE_NAME}.go

# start server
./${FILE_NAME} &

# get PID of server
PID_SERVER=${!}

# use cygstart instead of xdg-open command on Cygwin-env.
cygstart http://localhost:8000
cygstart http://localhost:8000/?line=black\&width=1200\&height=640
cygstart http://localhost:8000/?width=300\&line=lime
cygstart http://localhost:8000
cygstart http://localhost:8000/?cells=300\&line=gray
cygstart http://localhost:8000/?xyrange=10

# terminate server after 10 seconds
sleep 10s
kill ${PID_SERVER}

#----- [EOF] -----#

