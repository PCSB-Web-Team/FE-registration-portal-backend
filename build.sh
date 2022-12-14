#!/usr/bin/env bash
set -xe

# get all of the dependencies needed
go get github.com/aws/aws-sdk-go
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get github.com/razorpay/razorpay-go

# create the application binary that EB uses
go build -o bin/application .