#!/bin/sh
cd server/frontend
yarn
yarn build
cd ..
go get
nohup go run . > log.log &
tail log.log
