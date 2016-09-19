#!/bin/bash

# define suffix-less file name
FILE_NAME=3_9

# build executable-program
go build ${FILE_NAME}.go

# start server
./${FILE_NAME} &

# get PID of server
PID_SERVER=${!}

# use cygstart instead of xdg-open command on Cygwin-env.
#cygstart http://localhost:8000/?x=1024\&y=512
cygstart http://localhost:8000/?x=1\&y=-1
cygstart http://localhost:8000
cygstart http://localhost:8000/?zoom=2\&x=-0.5\&y=0.3
#cygstart http://localhost:8000/?zoom=2\&x=1024\&y=512
cygstart http://localhost:8000
cygstart http://localhost:8000/?zoom=4
#cygstart http://localhost:8000/?xxx=32\&zoom=10
#cygstart http://localhost:8000/?terminate

# terminate server after 10 seconds
sleep 30s
kill ${PID_SERVER}

#----- [EOF] -----#

