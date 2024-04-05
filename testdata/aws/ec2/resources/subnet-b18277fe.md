Provides an VPC subnet resource.

> **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), subnets associated with Lambda Functions can take up to 45 minutes to successfully delete.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Subnet("main", {
    vpcId: aws_vpc.main.id,
    cidrBlock: "10.0.1.0/24",
    tags: {
        Name: "Main",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Subnet("main",
    vpc_id=aws_vpc["main"]["id"],
    cidr_block="10.0.1.0/24",
    tags={
        "Name": "Main",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Subnet("main", new()
    {
        VpcId = aws_vpc.Main.Id,
        CidrBlock = "10.0.1.0/24",
        Tags = 
        {
            { "Name", "Main" },
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
		_, err := ec2.NewSubnet(ctx, "main", &ec2.SubnetArgs{
			VpcId:     pulumi.Any(aws_vpc.Main.Id),
			CidrBlock: pulumi.String("10.0.1.0/24"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Main"),
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
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
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
        var main = new Subnet("main", SubnetArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .cidrBlock("10.0.1.0/24")
            .tags(Map.of("Name", "Main"))
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${aws_vpc.main.id}
      cidrBlock: 10.0.1.0/24
      tags:
        Name: Main
```
{{% /example %}}
{{% example %}}
### Subnets In Secondary VPC CIDR Blocks

When managing subnets in one of a VPC's secondary CIDR blocks created using a `aws.ec2.VpcIpv4CidrBlockAssociation`
resource, it is recommended to reference that resource's `vpc_id` attribute to ensure correct dependency ordering.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const secondaryCidr = new aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", {
    vpcId: aws_vpc.main.id,
    cidrBlock: "172.20.0.0/16",
});
const inSecondaryCidr = new aws.ec2.Subnet("inSecondaryCidr", {
    vpcId: secondaryCidr.vpcId,
    cidrBlock: "172.20.0.0/24",
});
```
```python
import pulumi
import pulumi_aws as aws

secondary_cidr = aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr",
    vpc_id=aws_vpc["main"]["id"],
    cidr_block="172.20.0.0/16")
in_secondary_cidr = aws.ec2.Subnet("inSecondaryCidr",
    vpc_id=secondary_cidr.vpc_id,
    cidr_block="172.20.0.0/24")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var secondaryCidr = new Aws.Ec2.VpcIpv4CidrBlockAssociation("secondaryCidr", new()
    {
        VpcId = aws_vpc.Main.Id,
        CidrBlock = "172.20.0.0/16",
    });

    var inSecondaryCidr = new Aws.Ec2.Subnet("inSecondaryCidr", new()
    {
        VpcId = secondaryCidr.VpcId,
        CidrBlock = "172.20.0.0/24",
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
		secondaryCidr, err := ec2.NewVpcIpv4CidrBlockAssociation(ctx, "secondaryCidr", &ec2.VpcIpv4CidrBlockAssociationArgs{
			VpcId:     pulumi.Any(aws_vpc.Main.Id),
			CidrBlock: pulumi.String("172.20.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "inSecondaryCidr", &ec2.SubnetArgs{
			VpcId:     secondaryCidr.VpcId,
			CidrBlock: pulumi.String("172.20.0.0/24"),
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
import com.pulumi.aws.ec2.VpcIpv4CidrBlockAssociation;
import com.pulumi.aws.ec2.VpcIpv4CidrBlockAssociationArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
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
        var secondaryCidr = new VpcIpv4CidrBlockAssociation("secondaryCidr", VpcIpv4CidrBlockAssociationArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .cidrBlock("172.20.0.0/16")
            .build());

        var inSecondaryCidr = new Subnet("inSecondaryCidr", SubnetArgs.builder()        
            .vpcId(secondaryCidr.vpcId())
            .cidrBlock("172.20.0.0/24")
            .build());

    }
}
```
```yaml
resources:
  secondaryCidr:
    type: aws:ec2:VpcIpv4CidrBlockAssociation
    properties:
      vpcId: ${aws_vpc.main.id}
      cidrBlock: 172.20.0.0/16
  inSecondaryCidr:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${secondaryCidr.vpcId}
      cidrBlock: 172.20.0.0/24
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import subnets using the subnet `id`. For example:

```sh
 $ pulumi import aws:ec2/subnet:Subnet public_subnet subnet-9d4a7b6c
```
 