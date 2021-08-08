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


### How to test Manually

Test product ID is e8864473-91a0-4f4b-9ce6-903d15acce4f. Products are manually configured in the getProductPacksHandler.
#### Product Not found 404
curl http://localhost:8080/products/xxx/packs?quantity={Quantity}

#### Product Pack Configuration
curl http://localhost:8080/products/e8864473-91a0-4f4b-9ce6-903d15acce4f/packs?quantity={Quantity}

## Resources
1. Deployment from https://medium.com/tysonworks/deploy-go-applications-to-ecs-using-aws-cdk-1a97d85bb4cb
2. Project structure https://sourcegraph.com/github.com/katzien/go-structure-examples@master/-/blob/layered/main.go?L30 https://about.sourcegraph.com/go/gophercon-2018-how-do-you-structure-your-go-apps/
3. Creating a REST API GO https://tutorialedge.net/golang/creating-restful-api-with-golang/
