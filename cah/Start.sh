#!/bin/sh
cd server/public/react
yarn
yarn build
cd ../..
go get
nohup go run . > log.log &
