#!/bin/bash

if [ ! -d "gen" ]; then
  mkdir gen gen/models gen/restapi
fi

swagger generate server -t gen --exclude-main -A GOProject # you need to change that
swagger generate client -t gen -A GOProject # you need to change that
