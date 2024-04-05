Creates an association between a Route53 Health Check and a Shield Advanced protected resource.
This association uses the health of your applications to improve responsiveness and accuracy in attack detection and mitigation.

Blog post: [AWS Shield Advanced now supports Health Based Detection](https://aws.amazon.com/about-aws/whats-new/2020/02/aws-shield-advanced-now-supports-health-based-detection/)

{{% examples %}}
## Example Usage
{{% example %}}
### Create an association between a protected EIP and a Route53 Health Check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const exampleEip = new aws.ec2.Eip("exampleEip", {
    domain: "vpc",
    tags: {
        Name: "example",
    },
});
const exampleProtection = new aws.shield.Protection("exampleProtection", {resourceArn: pulumi.all([currentPartition, currentRegion, currentCallerIdentity, exampleEip.id]).apply(([currentPartition, currentRegion, currentCallerIdentity, id]) => `arn:${currentPartition.partition}:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${id}`)});
const exampleHealthCheck = new aws.route53.HealthCheck("exampleHealthCheck", {
    ipAddress: exampleEip.publicIp,
    port: 80,
    type: "HTTP",
    resourcePath: "/ready",
    failureThreshold: 3,
    requestInterval: 30,
    tags: {
        Name: "tf-example-health-check",
    },
});
const exampleProtectionHealthCheckAssociation = new aws.shield.ProtectionHealthCheckAssociation("exampleProtectionHealthCheckAssociation", {
    healthCheckArn: exampleHealthCheck.arn,
    shieldProtectionId: exampleProtection.id,
});
```
```python
import pulumi
import pulumi_aws as aws

current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
current_partition = aws.get_partition()
example_eip = aws.ec2.Eip("exampleEip",
    domain="vpc",
    tags={
        "Name": "example",
    })
example_protection = aws.shield.Protection("exampleProtection", resource_arn=example_eip.id.apply(lambda id: f"arn:{current_partition.partition}:ec2:{current_region.name}:{current_caller_identity.account_id}:eip-allocation/{id}"))
example_health_check = aws.route53.HealthCheck("exampleHealthCheck",
    ip_address=example_eip.public_ip,
    port=80,
    type="HTTP",
    resource_path="/ready",
    failure_threshold=3,
    request_interval=30,
    tags={
        "Name": "tf-example-health-check",
    })
example_protection_health_check_association = aws.shield.ProtectionHealthCheckAssociation("exampleProtectionHealthCheckAssociation",
    health_check_arn=example_health_check.arn,
    shield_protection_id=example_protection.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var exampleEip = new Aws.Ec2.Eip("exampleEip", new()
    {
        Domain = "vpc",
        Tags = 
        {
            { "Name", "example" },
        },
    });

    var exampleProtection = new Aws.Shield.Protection("exampleProtection", new()
    {
        ResourceArn = Output.Tuple(currentPartition, currentRegion, currentCallerIdentity, exampleEip.Id).Apply(values =>
        {
            var currentPartition = values.Item1;
            var currentRegion = values.Item2;
            var currentCallerIdentity = values.Item3;
            var id = values.Item4;
            return $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:ec2:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:eip-allocation/{id}";
        }),
    });

    var exampleHealthCheck = new Aws.Route53.HealthCheck("exampleHealthCheck", new()
    {
        IpAddress = exampleEip.PublicIp,
        Port = 80,
        Type = "HTTP",
        ResourcePath = "/ready",
        FailureThreshold = 3,
        RequestInterval = 30,
        Tags = 
        {
            { "Name", "tf-example-health-check" },
        },
    });

    var exampleProtectionHealthCheckAssociation = new Aws.Shield.ProtectionHealthCheckAssociation("exampleProtectionHealthCheckAssociation", new()
    {
        HealthCheckArn = exampleHealthCheck.Arn,
        ShieldProtectionId = exampleProtection.Id,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleEip, err := ec2.NewEip(ctx, "exampleEip", &ec2.EipArgs{
			Domain: pulumi.String("vpc"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
			},
		})
		if err != nil {
			return err
		}
		exampleProtection, err := shield.NewProtection(ctx, "exampleProtection", &shield.ProtectionArgs{
			ResourceArn: exampleEip.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("arn:%v:ec2:%v:%v:eip-allocation/%v", currentPartition.Partition, currentRegion.Name, currentCallerIdentity.AccountId, id), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		exampleHealthCheck, err := route53.NewHealthCheck(ctx, "exampleHealthCheck", &route53.HealthCheckArgs{
			IpAddress:        exampleEip.PublicIp,
			Port:             pulumi.Int(80),
			Type:             pulumi.String("HTTP"),
			ResourcePath:     pulumi.String("/ready"),
			FailureThreshold: pulumi.Int(3),
			RequestInterval:  pulumi.Int(30),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-example-health-check"),
			},
		})
		if err != nil {
			return err
		}
		_, err = shield.NewProtectionHealthCheckAssociation(ctx, "exampleProtectionHealthCheckAssociation", &shield.ProtectionHealthCheckAssociationArgs{
			HealthCheckArn:     exampleHealthCheck.Arn,
			ShieldProtectionId: exampleProtection.ID(),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.ec2.Eip;
import com.pulumi.aws.ec2.EipArgs;
import com.pulumi.aws.shield.Protection;
import com.pulumi.aws.shield.ProtectionArgs;
import com.pulumi.aws.route53.HealthCheck;
import com.pulumi.aws.route53.HealthCheckArgs;
import com.pulumi.aws.shield.ProtectionHealthCheckAssociation;
import com.pulumi.aws.shield.ProtectionHealthCheckAssociationArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentPartition = AwsFunctions.getPartition();

        var exampleEip = new Eip("exampleEip", EipArgs.builder()        
            .domain("vpc")
            .tags(Map.of("Name", "example"))
            .build());

        var exampleProtection = new Protection("exampleProtection", ProtectionArgs.builder()        
            .resourceArn(exampleEip.id().applyValue(id -> String.format("arn:%s:ec2:%s:%s:eip-allocation/%s", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),id)))
            .build());

        var exampleHealthCheck = new HealthCheck("exampleHealthCheck", HealthCheckArgs.builder()        
            .ipAddress(exampleEip.publicIp())
            .port(80)
            .type("HTTP")
            .resourcePath("/ready")
            .failureThreshold("3")
            .requestInterval("30")
            .tags(Map.of("Name", "tf-example-health-check"))
            .build());

        var exampleProtectionHealthCheckAssociation = new ProtectionHealthCheckAssociation("exampleProtectionHealthCheckAssociation", ProtectionHealthCheckAssociationArgs.builder()        
            .healthCheckArn(exampleHealthCheck.arn())
            .shieldProtectionId(exampleProtection.id())
            .build());

    }
}
```
```yaml
resources:
  exampleEip:
    type: aws:ec2:Eip
    properties:
      domain: vpc
      tags:
        Name: example
  exampleProtection:
    type: aws:shield:Protection
    properties:
      resourceArn: arn:${currentPartition.partition}:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${exampleEip.id}
  exampleHealthCheck:
    type: aws:route53:HealthCheck
    properties:
      ipAddress: ${exampleEip.publicIp}
      port: 80
      type: HTTP
      resourcePath: /ready
      failureThreshold: '3'
      requestInterval: '30'
      tags:
        Name: tf-example-health-check
  exampleProtectionHealthCheckAssociation:
    type: aws:shield:ProtectionHealthCheckAssociation
    properties:
      healthCheckArn: ${exampleHealthCheck.arn}
      shieldProtectionId: ${exampleProtection.id}
variables:
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Shield protection health check association resources using the `shield_protection_id` and `health_check_arn`. For example:

```sh
 $ pulumi import aws:shield/protectionHealthCheckAssociation:ProtectionHealthCheckAssociation example ff9592dc-22f3-4e88-afa1-7b29fde9669a+arn:aws:route53:::healthcheck/3742b175-edb9-46bc-9359-f53e3b794b1b
```
 