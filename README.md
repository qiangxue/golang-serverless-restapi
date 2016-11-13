# Serverless REST API in Go

This project demonstrates how to develop a serverless REST API using Go and [apex](http://apex.run/).

## Installation

Install apex and the Go packages as instructed below:

```shell
# install apex if you don't have it yet
curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sudo sh

# download this package
go get github.com/qiangxue/golang-serverless-restapi

# download the go-apex package
go get github.com/apex/go-apex
```

## Deploying Lambda

First, follow the [apex documentation](http://apex.run/#aws-credentials) to set up the AWS credentials needed for deploying Lambdas using apex.

Then, modify the `project.json` file by inserting the correct AWS `role` and `profile` values.

Run the following command to build and deploy the Lambda:
```shell
apex deploy --region us-east-1
```

And use the AWS console or the following command to verify that the Lambda is successfully deployed:
```shell
apex invoke apis
```

## Configuring AWS API Gateway

1. Log into the AWS console and switch to the API Gateway page. On that page, choose "Create new API" and enter the API name as "hello" (or any other name you prefer).
2. In the "Resources" tab of the "hello" API, click on the "Actions" dropdown button and select "Create Resource". In the "New Child Resource" page, select "Configure as proxy resource", and then click on the "Create Resource" button.
3. In the "/{proxy+} - ANY - Setup" page, choose "Lambda Region" as `us-east-1` (or the actual AWS region that you used to deploy the Lambda function) and enter the "Lambda Function" name as `rest_apis`. Click the "Save" button to complete the proxy resource setup.
4. Click on the "Actions" dropdown button and select "Deploy API". In the popup window, choose `[New Stage]` in the "Deployment stage" dropdown list, and enter `prod` in the "Stage name" input field. Click on the "Deploy" button to complete the deployment.
5. At this point, you should be redirected to a page showing the deployment information about the "hello" API. You should see an "Invoke URL" on the page. The URL may look like "https://bm7empvth7.execute-api.us-east-1.amazonaws.com/prod". This is the base URL that everyone can use to hit our APIs.

## Trying it Out

Run the following commands to verify the API is accessible and working as expected:

```shell
# replace InvokeURL with the actual Invoke URL found in Step 5
> curl InvokeURL/foo
hello

> curl InvokeURL/bar?hello=world
GET /bar?hello=world
```

In the AWS console, locate the Lambda function named `rest_apis` and check its monitoring result to verify that the Lambda was invoked.

## What's Next

In the above demo, we used the standard HTTP ServerMux to wire up the HTTP handlers for handling different API endpoints. You may replace it with your favorite third-party HTTP routers (e.g. [ozzo-routing](https://github.com/go-ozzo/ozzo-routing), [echo](https://github.com/labstack/echo), [gin](https://github.com/gin-gonic/gin)). This can be done easily by modifying the `functions/apis/main.go` file.

Because we are using AWS API Gateway as a proxy to invoke the Lambda, we may implement more API endpoints without the need of reconfiguring the Gateway. That is, each time we make changes to our REST API project, we only need to run `apex deploy` to deploy it to AWS. We do not need to reconfigure AWS API Gateway. 
  