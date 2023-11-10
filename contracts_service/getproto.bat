@echo off
set version=%1

pushd %~dp0
go install github.com/techvify-oliver/protoserver@v%version%
go get github.com/techvify-oliver/protoserver@v%version%
go mod tidy
popd