## Instructions

---

```bash
# init project
go mod init go-serverless

# AWS SDK - https://github.com/aws/aws-sdk-go
go get github.com/aws/aws-sdk-go
# Update package
go get -u github.com/aws/aws-sdk-go

# AWS Lambda
go get github.com/aws/aws-lambda-go

# ?
go mod tidy

# check problems && build
GOOS=linux go build -o build/main cmd/main.go

# zip the build
zip -jrm build/main.zip build/main
```

---

## AWS Console

---

### Create Lambda

1. Lambda > `Create Function`
2. Runtime: `Go 1.x`
3. Change default execution role: `Create a new role from AWS policy templates`
4. Policy Templates: `Simple Microservice Permission (DynamoDB)`
5. `Create`

### Change handler on function details:

1. Runtime Settings: `Edit`
2. Handler: `main`
3. `Save`

### Upload code:

1. Code Source: `Upload from` > `a .zip file`
2. File: `main.zip`
3. `Save`

### Create DynamoDB Table

1. DynamoDB > `Create Table`
2. Table name: `go-serverless`
3. Partition key: `email`
4. `Create Table`

### Create API Gateway and Deploy API

1. API Gateway > `Build >REST API`
2. Protocol: `REST`
3. new API: `NEW API`
4. API name: `go-serverless`
5. `Create API`
6. Actions > `Create Method` > `ANY`
7. Integration Type: `Lambda Function`
8. Use Lambda Proxy Integration: `checked`
9. Lambda function: `go-serverless`
10. Use default timeout: `checked`
11. Actions > `Deploy API`
12. Deployment state: `[New Stage]`
13. Stage name: `staging`
14. `Deploy`
