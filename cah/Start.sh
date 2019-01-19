#!/bin/sh
cd app/frontend
yarn
yarn build
cd ..
go get -t
nohup go run . > log.log &
tail log.log
