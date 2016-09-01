#!/bin/bash

# define suffix-less file name
FILE_NAME=3_4

# build executable-program
go build ${FILE_NAME}.go

# start server
./${FILE_NAME} &

# get PID of server
PID_SERVER=${!}

# use cygstart instead of xdg-open command on Cygwin-env.
cygstart http://localhost:8000

# terminate server after 10 seconds
sleep 5s
kill ${PID_SERVER}

#----- [EOF] -----#

