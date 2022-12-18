#!/bin/bash

goyacc -o parse.go vic.y
gofmt -w *.go
