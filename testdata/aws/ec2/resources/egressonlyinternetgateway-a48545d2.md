[IPv6 only] Creates an egress-only Internet gateway for your VPC.
An egress-only Internet gateway is used to enable outbound communication
over IPv6 from instances in your VPC to the Internet, and prevents hosts
outside of your VPC from initiating an IPv6 connection with your instance.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
    cidrBlock: "10.1.0.0/16",
    assignGeneratedIpv6CidrBlock: true,
});
const exampleEgressOnlyInternetGateway = new aws.ec2.EgressOnlyInternetGateway("exampleEgressOnlyInternetGateway", {
    vpcId: exampleVpc.id,
    tags: {
        Name: "main",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="10.1.0.0/16",
    assign_generated_ipv6_cidr_block=True)
example_egress_only_internet_gateway = aws.ec2.EgressOnlyInternetGateway("exampleEgressOnlyInternetGateway",
    vpc_id=example_vpc.id,
    tags={
        "Name": "main",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
        AssignGeneratedIpv6CidrBlock = true,
    });

    var exampleEgressOnlyInternetGateway = new Aws.Ec2.EgressOnlyInternetGateway("exampleEgressOnlyInternetGateway", new()
    {
        VpcId = exampleVpc.Id,
        Tags = 
        {
            { "Name", "main" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock:                    pulumi.String("10.1.0.0/16"),
			AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewEgressOnlyInternetGateway(ctx, "exampleEgressOnlyInternetGateway", &ec2.EgressOnlyInternetGatewayArgs{
			VpcId: exampleVpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.EgressOnlyInternetGateway;
import com.pulumi.aws.ec2.EgressOnlyInternetGatewayArgs;
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
        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .assignGeneratedIpv6CidrBlock(true)
            .build());

        var exampleEgressOnlyInternetGateway = new EgressOnlyInternetGateway("exampleEgressOnlyInternetGateway", EgressOnlyInternetGatewayArgs.builder()        
            .vpcId(exampleVpc.id())
            .tags(Map.of("Name", "main"))
            .build());

    }
}
```
```yaml
resources:
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
      assignGeneratedIpv6CidrBlock: true
  exampleEgressOnlyInternetGateway:
    type: aws:ec2:EgressOnlyInternetGateway
    properties:
      vpcId: ${exampleVpc.id}
      tags:
        Name: main
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Egress-only Internet gateways using the `id`. For example:

```sh
 $ pulumi import aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway example eigw-015e0e244e24dfe8a
```
 