#!/bin/bash

# upload.sh folder-name name threads



curl http://localhost:8080/upload --data "{\"environment\": {\"ENVVAR\":\"value\" }, \"name\": \"$2\", \"threads\": $3, \"tarball\": \"`tar -cvf - -C $1 . | base64 -w 0`\"}"
