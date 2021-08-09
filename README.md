# product-pack-test-go

## Deployment
To execute a deployment excute the following commands
1. `npm install` to install all dependencies required for cdk.
2. `cdk bootstrap`
3. `cdk deploy`

### Useful commands

 * `npm run build`   compile typescript to js
 * `npm run watch`   watch for changes and compile
 * `npm run test`    perform the jest unit tests
 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template


## Test Details

We have a Golang based API in the api directory of this repo. It is deployed to AWS with the Cloud Development Kit (CDK). 
The application itself is deployed to AWS ECS.

There is a swagger file in the repo with details of the request and the response.

There are some tests for the API in the app_test.go file. The scenarios set out in the test have been covered.

API is available on `http://GSTec-Farga-1BLYX3MS2MBWI-878009822.eu-west-1.elb.amazonaws.com/products/e8864473-91a0-4f4b-9ce6-903d15acce4f/packs?quantity={X}`. The product itself is not persisted the only product we have is statically configured in code and has ID `e8864473-91a0-4f4b-9ce6-903d15acce4f`.

The pack sizes are not currently configurable. But they are a property of the Product class. Which if it was persisted it would mean
that the pack sizes could be stored in the DB and therefore wouldn't need code changes to update.


### How to test Manually

Test product ID is e8864473-91a0-4f4b-9ce6-903d15acce4f. Products are manually configured in the getProductPacksHandler.
#### Product Not found 404
curl http://{HOSTNAME}:{PORT}/products/xxx/packs?quantity={Quantity}

#### Product Pack Configuration
curl http://{HOSTNAME}:{PORT}/products/e8864473-91a0-4f4b-9ce6-903d15acce4f/packs?quantity={Quantity}

### Running Locally
After running `make start` the API should be available on port 8080.

### Running tests
1. cd into the `api` directory.
2. Execute `go test`

## Resources
1. Deployment from https://medium.com/tysonworks/deploy-go-applications-to-ecs-using-aws-cdk-1a97d85bb4cb
2. Project structure https://sourcegraph.com/github.com/katzien/go-structure-examples@master/-/blob/layered/main.go?L30 https://about.sourcegraph.com/go/gophercon-2018-how-do-you-structure-your-go-apps/
3. Creating a REST API GO https://tutorialedge.net/golang/creating-restful-api-with-golang/
