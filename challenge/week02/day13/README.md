# Day 13 of #66DaysOfGo

_Last update:  Jul 25, 2023_.

---

Today, I've started a new series related to AWS development with Go.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6
- npm: 6.14.6
- serverless: 3.33.0
- aws-cli: 2.0.38
- jq: 1.6

## Setup

### Requirements

- npm configured
- aws cli installed
- aws credentials with the right permissions
  - I've selected the region us-east-1

### Configure

Probably one of the easiest ways to deploy applications in AWS is using the Serverless Framework. You won't have to deal with complex Cloudformation templates, but with a simple yml manifest.

### Install serverless with npm

```bash
$ npm install -g serverless

# shortened for brevity...
Serverless Framework successfully installed!

To start your first project run “serverless”.

Turn on automatic updates by running “serverless config --autoupdate”.
+ serverless@3.33.0
added 407 packages from 332 contributors in 49.112s


   ╭───────────────────────────────────────────────────────────────╮
   │                                                               │
   │      New major version of npm available! 6.14.6 → 9.8.1       │
   │   Changelog: https://github.com/npm/cli/releases/tag/v9.8.1   │
   │               Run npm install -g npm to update!               │
   │                                                               │
   ╰───────────────────────────────────────────────────────────────╯

```

### Create a simple template for Golang development

There are several templates to choose from. Let's pick "`aws-go`".

```bash
# The following will create a serverless project in the folder "test-aws-go"
$ serverless create --template aws-go --path test-aws-go
✔ Project successfully created in "test-aws-go" from "aws-go" template (5s)
```

Within the "test-aws-go", there's a .gitignore file. Copy all the contents to the root .gitignore file. Also add "node_modules/" to it.

```bash
cd test-aws-go
```

There are two .go files. By default they are using the AWS SDK V1.

Create a go.mod file, and retrieve the AWS dependencies.

```bash
$ go mod init example.com
go: creating new go.mod: module example.com
go: to add module requirements and sums:
        go mod tidy
```

```bash
$ go get -v all
go: added github.com/aws/aws-lambda-go v1.41.0
```

## Build and deploy the project

In the serverless manifest (serverless.yml) make a few changes:

- Comment out the line `- '!./**'` (you might face the error "[Serverless error - No file matches include / exclude patterns](https://stackoverflow.com/questions/66001173/serverless-error-no-file-matches-include-exclude-patterns)")
- Change the API type from HTTP to Rest API by modifying the lines `httpApi` to `http`.

```yaml
package:
  patterns:
    # - '!./**'
    - ./bin/**
````

You'll end up with:

```yaml
functions:
  hello:
    handler: bin/hello
    events:
      # - httpApi: # http
      - http: # rest
          path: /hello
          method: get
  world:
    handler: bin/world
    events:
      # - httpApi: # http
      - http: # rest
          path: /world
          method: get
```

Try to build the project. If there's no errors, you can safely deploy.

```bash
# build
$ make build
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
```

```bash
# deploy
$ make deploy
rm -rf ./bin
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/hello hello/main.go
env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/world world/main.go
sls deploy --verbose # shortcut for "serverless"

Deploying test-aws-go to stage dev (us-east-1)

Packaging
Excluding development dependencies for service package
Retrieving CloudFormation stack
Creating CloudFormation stack
Creating new change set
Waiting for new change set to be created
Change Set did not reach desired state, retrying
Executing created change set

# shortened for brevity...

  CREATE_IN_PROGRESS - AWS::ApiGateway::Deployment - ApiGatewayDeployment1690323972100
  CREATE_IN_PROGRESS - AWS::ApiGateway::Deployment - ApiGatewayDeployment1690323972100
  CREATE_COMPLETE - AWS::ApiGateway::Deployment - ApiGatewayDeployment1690323972100
  UPDATE_COMPLETE_CLEANUP_IN_PROGRESS - AWS::CloudFormation::Stack - test-aws-go-dev
  UPDATE_COMPLETE - AWS::CloudFormation::Stack - test-aws-go-dev
Retrieving CloudFormation stack
Removing old service artifacts from S3

✔ Service deployed to stack test-aws-go-dev (116s)

endpoints:
  GET - https://k7azu6and9.execute-api.us-east-1.amazonaws.com/dev/hello
  GET - https://k7azu6and9.execute-api.us-east-1.amazonaws.com/dev/world
functions:
  hello: test-aws-go-dev-hello (5.9 MB)
  world: test-aws-go-dev-world (5.9 MB)

Stack Outputs:
  HelloLambdaFunctionQualifiedArn: arn:aws:lambda:us-east-1:141430882746:function:test-aws-go-dev-hello:5
  WorldLambdaFunctionQualifiedArn: arn:aws:lambda:us-east-1:141430882746:function:test-aws-go-dev-world:4
  ServiceEndpoint: https://k7azu6and9.execute-api.us-east-1.amazonaws.com/dev
  ServerlessDeploymentBucketName: test-aws-go-dev-serverlessdeploymentbucket-1bglenh970q9e

Need a faster logging experience than CloudWatch? Try our Dev Mode in Console: run "serverless dev"
```

The final output will show you two new endpoints deployed, behind a Rest API, as AWS Lambda functions.

> Note: After the deployment, a (hidden) `.serverless` folder was created. In there, you can find the Cloudformation templates that are managed, under the hood, by the Serverless framework.

## Test out the service

```bash
curl -s -X GET https://k7azu6and9.execute-api.us-east-1.amazonaws.com/dev/hello | jq .
```bash

```json
{
  "message": "Go Serverless v1.0! Your function executed successfully!"
}
```

```bash
curl -s -X GET https://k7azu6and9.execute-api.us-east-1.amazonaws.com/dev/world | jq .
```

```json
{
  "message": "Okay so your other function also executed successfully!"
}
```

## Cleanup

```bash
$ serverless remove
Removing test-aws-go from stage dev (us-east-1)

✔ Service test-aws-go has been successfully removed (35s)
```

---

## Golang code

The following two functions were created automatically by the Serverless framework template.

### src/hello/main.go

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
    var buf bytes.Buffer

    body, err := json.Marshal(map[string]interface{}{
        "message": "Go Serverless v1.0! Your function executed successfully!",
    })
    if err != nil {
        return Response{StatusCode: 404}, err
    }
    json.HTMLEscape(&buf, body)

    resp := Response{
        StatusCode:      200,
        IsBase64Encoded: false,
        Body:            buf.String(),
        Headers: map[string]string{
            "Content-Type":           "application/json",
            "X-MyCompany-Func-Reply": "hello-handler",
        },
    }

    return resp, nil
}

func main() {
    lambda.Start(Handler)
}
```

### src/world/main.go

```go
package main

import (
    "bytes"
    "context"
    "encoding/json"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
    var buf bytes.Buffer

    body, err := json.Marshal(map[string]interface{}{
        "message": "Okay so your other function also executed successfully!",
    })
    if err != nil {
        return Response{StatusCode: 404}, err
    }
    json.HTMLEscape(&buf, body)

    resp := Response{
        StatusCode:      200,
        IsBase64Encoded: false,
        Body:            buf.String(),
        Headers: map[string]string{
            "Content-Type":           "application/json",
            "X-MyCompany-Func-Reply": "world-handler",
        },
    }

    return resp, nil
}

func main() {
    lambda.Start(Handler)
}
```

---

## References

- [Setting Up Serverless Framework With AWS](https://www.serverless.com/framework/docs/getting-started)
- [AWS SDK for Go](https://aws.amazon.com/sdk-for-go/)
