#!/bin/bash
VERSION=$(cat VERSION)
go run cmd/webapi/webapi.go -v $VERSION