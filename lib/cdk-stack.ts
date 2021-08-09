import * as cdk from "@aws-cdk/core";
import * as ec2 from "@aws-cdk/aws-ec2";
import * as ecs from '@aws-cdk/aws-ecs';
import * as ecs_patterns from '@aws-cdk/aws-ecs-patterns';
import { config } from "dotenv";
config();

export class CdkStack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps){
        super(scope, id, props);

        const vpc = new ec2.Vpc(this, 'vpc', { maxAzs: 2 });

        const cluster = new ecs.Cluster(this, 'Cluster', { vpc });

        const fargateService = new ecs_patterns.ApplicationLoadBalancedFargateService(this, "FargateService", {
            cluster,
            taskImageOptions: {
                image: ecs.ContainerImage.fromAsset(`${__dirname}/../api/`),
                containerPort: 8080,
                environment: {
                    DEPLOYED_DATE: Date.now().toLocaleString()
                }
            },
            desiredCount: 1,
            minHealthyPercent: 100 // Percentage of the desired count that need to be available.
        });

        new cdk.CfnOutput(this, 'LoadBalancerDNS', { value: fargateService.loadBalancer.loadBalancerDnsName });
    }
}