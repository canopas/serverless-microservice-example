# Serverless microservice example

This repository guides you to implement serverless microservice by using simple lambda function.

Perform following steps to configure the microservice example.

1. Run `GOOS=linux GOARCH=amd64 go build -o main main.go`
 This command will generate binary file from `main.go`.
 
2. Run `zip lambda.zip main`.
This will the compress generated binary file and will make zip file named main.zip. 
