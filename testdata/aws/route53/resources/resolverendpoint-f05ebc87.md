Provides a Route 53 Resolver endpoint resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.route53.ResolverEndpoint("foo", {
    direction: "INBOUND",
    securityGroupIds: [
        aws_security_group.sg1.id,
        aws_security_group.sg2.id,
    ],
    ipAddresses: [
        {
            subnetId: aws_subnet.sn1.id,
        },
        {
            subnetId: aws_subnet.sn2.id,
            ip: "10.0.64.4",
        },
    ],
    tags: {
        Environment: "Prod",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.route53.ResolverEndpoint("foo",
    direction="INBOUND",
    security_group_ids=[
        aws_security_group["sg1"]["id"],
        aws_security_group["sg2"]["id"],
    ],
    ip_addresses=[
        aws.route53.ResolverEndpointIpAddressArgs(
            subnet_id=aws_subnet["sn1"]["id"],
        ),
        aws.route53.ResolverEndpointIpAddressArgs(
            subnet_id=aws_subnet["sn2"]["id"],
            ip="10.0.64.4",
        ),
    ],
    tags={
        "Environment": "Prod",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Route53.ResolverEndpoint("foo", new()
    {
        Direction = "INBOUND",
        SecurityGroupIds = new[]
        {
            aws_security_group.Sg1.Id,
            aws_security_group.Sg2.Id,
        },
        IpAddresses = new[]
        {
            new Aws.Route53.Inputs.ResolverEndpointIpAddressArgs
            {
                SubnetId = aws_subnet.Sn1.Id,
            },
            new Aws.Route53.Inputs.ResolverEndpointIpAddressArgs
            {
                SubnetId = aws_subnet.Sn2.Id,
                Ip = "10.0.64.4",
            },
        },
        Tags = 
        {
            { "Environment", "Prod" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewResolverEndpoint(ctx, "foo", &route53.ResolverEndpointArgs{
			Direction: pulumi.String("INBOUND"),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Sg1.Id,
				aws_security_group.Sg2.Id,
			},
			IpAddresses: route53.ResolverEndpointIpAddressArray{
				&route53.ResolverEndpointIpAddressArgs{
					SubnetId: pulumi.Any(aws_subnet.Sn1.Id),
				},
				&route53.ResolverEndpointIpAddressArgs{
					SubnetId: pulumi.Any(aws_subnet.Sn2.Id),
					Ip:       pulumi.String("10.0.64.4"),
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
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
import com.pulumi.aws.route53.ResolverEndpoint;
import com.pulumi.aws.route53.ResolverEndpointArgs;
import com.pulumi.aws.route53.inputs.ResolverEndpointIpAddressArgs;
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
        var foo = new ResolverEndpoint("foo", ResolverEndpointArgs.builder()        
            .direction("INBOUND")
            .securityGroupIds(            
                aws_security_group.sg1().id(),
                aws_security_group.sg2().id())
            .ipAddresses(            
                ResolverEndpointIpAddressArgs.builder()
                    .subnetId(aws_subnet.sn1().id())
                    .build(),
                ResolverEndpointIpAddressArgs.builder()
                    .subnetId(aws_subnet.sn2().id())
                    .ip("10.0.64.4")
                    .build())
            .tags(Map.of("Environment", "Prod"))
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:route53:ResolverEndpoint
    properties:
      direction: INBOUND
      securityGroupIds:
        - ${aws_security_group.sg1.id}
        - ${aws_security_group.sg2.id}
      ipAddresses:
        - subnetId: ${aws_subnet.sn1.id}
        - subnetId: ${aws_subnet.sn2.id}
          ip: 10.0.64.4
      tags:
        Environment: Prod
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import

Route 53 Resolver endpoints using the Route 53 Resolver endpoint ID. For example:

```sh
 $ pulumi import aws:route53/resolverEndpoint:ResolverEndpoint foo rslvr-in-abcdef01234567890
```
 