#!/usr/bin/env bash
set -ex

go get github.com/aws/aws-sdk-go

go get github.com/gin-gonic/gin

go get github.com/joho/godotenv

go get github.com/razorpay/razorpay-go

go build -o build/application .