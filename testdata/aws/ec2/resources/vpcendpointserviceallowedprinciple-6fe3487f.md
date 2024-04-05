Provides a resource to allow a principal to discover a VPC endpoint service.

> **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
both a standalone VPC Endpoint Service Allowed Principal resource
and a VPC Endpoint Service resource with an `allowed_principals` attribute. Do not use the same principal ARN in both
a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
and will overwrite the association.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const allowMeToFoo = new aws.ec2.VpcEndpointServiceAllowedPrinciple("allowMeToFoo", {
    vpcEndpointServiceId: aws_vpc_endpoint_service.foo.id,
    principalArn: current.then(current => current.arn),
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
allow_me_to_foo = aws.ec2.VpcEndpointServiceAllowedPrinciple("allowMeToFoo",
    vpc_endpoint_service_id=aws_vpc_endpoint_service["foo"]["id"],
    principal_arn=current.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var allowMeToFoo = new Aws.Ec2.VpcEndpointServiceAllowedPrinciple("allowMeToFoo", new()
    {
        VpcEndpointServiceId = aws_vpc_endpoint_service.Foo.Id,
        PrincipalArn = current.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointServiceAllowedPrinciple(ctx, "allowMeToFoo", &ec2.VpcEndpointServiceAllowedPrincipleArgs{
			VpcEndpointServiceId: pulumi.Any(aws_vpc_endpoint_service.Foo.Id),
			PrincipalArn:         *pulumi.String(current.Arn),
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
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.VpcEndpointServiceAllowedPrinciple;
import com.pulumi.aws.ec2.VpcEndpointServiceAllowedPrincipleArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var allowMeToFoo = new VpcEndpointServiceAllowedPrinciple("allowMeToFoo", VpcEndpointServiceAllowedPrincipleArgs.builder()        
            .vpcEndpointServiceId(aws_vpc_endpoint_service.foo().id())
            .principalArn(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn()))
            .build());

    }
}
```
```yaml
resources:
  allowMeToFoo:
    type: aws:ec2:VpcEndpointServiceAllowedPrinciple
    properties:
      vpcEndpointServiceId: ${aws_vpc_endpoint_service.foo.id}
      principalArn: ${current.arn}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}