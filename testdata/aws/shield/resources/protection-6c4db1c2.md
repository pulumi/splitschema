Enables AWS Shield Advanced for a specific AWS resource.
The resource can be an Amazon CloudFront distribution, Elastic Load Balancing load balancer, AWS Global Accelerator accelerator, Elastic IP Address, or an Amazon Route 53 hosted zone.

{{% examples %}}
## Example Usage
{{% example %}}
### Create protection

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({});
const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const exampleEip = new aws.ec2.Eip("exampleEip", {domain: "vpc"});
const exampleProtection = new aws.shield.Protection("exampleProtection", {
    resourceArn: pulumi.all([currentRegion, currentCallerIdentity, exampleEip.id]).apply(([currentRegion, currentCallerIdentity, id]) => `arn:aws:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${id}`),
    tags: {
        Environment: "Dev",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones()
current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
example_eip = aws.ec2.Eip("exampleEip", domain="vpc")
example_protection = aws.shield.Protection("exampleProtection",
    resource_arn=example_eip.id.apply(lambda id: f"arn:aws:ec2:{current_region.name}:{current_caller_identity.account_id}:eip-allocation/{id}"),
    tags={
        "Environment": "Dev",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke();

    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var exampleEip = new Aws.Ec2.Eip("exampleEip", new()
    {
        Domain = "vpc",
    });

    var exampleProtection = new Aws.Shield.Protection("exampleProtection", new()
    {
        ResourceArn = Output.Tuple(currentRegion, currentCallerIdentity, exampleEip.Id).Apply(values =>
        {
            var currentRegion = values.Item1;
            var currentCallerIdentity = values.Item2;
            var id = values.Item3;
            return $"arn:aws:ec2:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:eip-allocation/{id}";
        }),
        Tags = 
        {
            { "Environment", "Dev" },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetAvailabilityZones(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleEip, err := ec2.NewEip(ctx, "exampleEip", &ec2.EipArgs{
			Domain: pulumi.String("vpc"),
		})
		if err != nil {
			return err
		}
		_, err = shield.NewProtection(ctx, "exampleProtection", &shield.ProtectionArgs{
			ResourceArn: exampleEip.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("arn:aws:ec2:%v:%v:eip-allocation/%v", currentRegion.Name, currentCallerIdentity.AccountId, id), nil
			}).(pulumi.StringOutput),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
			},
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
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.Eip;
import com.pulumi.aws.ec2.EipArgs;
import com.pulumi.aws.shield.Protection;
import com.pulumi.aws.shield.ProtectionArgs;
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
        final var available = AwsFunctions.getAvailabilityZones();

        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        var exampleEip = new Eip("exampleEip", EipArgs.builder()        
            .domain("vpc")
            .build());

        var exampleProtection = new Protection("exampleProtection", ProtectionArgs.builder()        
            .resourceArn(exampleEip.id().applyValue(id -> String.format("arn:aws:ec2:%s:%s:eip-allocation/%s", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),id)))
            .tags(Map.of("Environment", "Dev"))
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
  exampleProtection:
    type: aws:shield:Protection
    properties:
      resourceArn: arn:aws:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${exampleEip.id}
      tags:
        Environment: Dev
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments: {}
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Shield protection resources using specifying their ID. For example:

```sh
 $ pulumi import aws:shield/protection:Protection example ff9592dc-22f3-4e88-afa1-7b29fde9669a
```
 