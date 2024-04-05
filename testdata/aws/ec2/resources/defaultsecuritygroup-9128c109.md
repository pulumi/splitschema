Provides a resource to manage a default security group. This resource can manage the default security group of the default or a non-default VPC.

> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `aws.ec2.DefaultSecurityGroup` resource behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management.

When the provider first begins managing the default security group, it **immediately removes all ingress and egress rules in the Security Group**. It then creates any rules specified in the configuration. This way only the rules specified in the configuration are created.

This resource treats its inline rules as absolute; only the rules defined inline are created, and any additions/removals external to this resource will result in diff shown. For these reasons, this resource is incompatible with the `aws.ec2.SecurityGroupRule` resource.

For more information about default security groups, see the AWS documentation on [Default Security Groups][aws-default-security-groups]. To manage normal security groups, see the `aws.ec2.SecurityGroup` resource.

{{% examples %}}
## Example Usage
{{% example %}}

The following config gives the default security group the same rules that AWS provides by default but under management by this provider. This means that any ingress or egress rules added or changed will be detected as drift.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainvpc = new aws.ec2.Vpc("mainvpc", {cidrBlock: "10.1.0.0/16"});
const _default = new aws.ec2.DefaultSecurityGroup("default", {
    vpcId: mainvpc.id,
    ingress: [{
        protocol: "-1",
        self: true,
        fromPort: 0,
        toPort: 0,
    }],
    egress: [{
        fromPort: 0,
        toPort: 0,
        protocol: "-1",
        cidrBlocks: ["0.0.0.0/0"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

mainvpc = aws.ec2.Vpc("mainvpc", cidr_block="10.1.0.0/16")
default = aws.ec2.DefaultSecurityGroup("default",
    vpc_id=mainvpc.id,
    ingress=[aws.ec2.DefaultSecurityGroupIngressArgs(
        protocol="-1",
        self=True,
        from_port=0,
        to_port=0,
    )],
    egress=[aws.ec2.DefaultSecurityGroupEgressArgs(
        from_port=0,
        to_port=0,
        protocol="-1",
        cidr_blocks=["0.0.0.0/0"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mainvpc = new Aws.Ec2.Vpc("mainvpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var @default = new Aws.Ec2.DefaultSecurityGroup("default", new()
    {
        VpcId = mainvpc.Id,
        Ingress = new[]
        {
            new Aws.Ec2.Inputs.DefaultSecurityGroupIngressArgs
            {
                Protocol = "-1",
                Self = true,
                FromPort = 0,
                ToPort = 0,
            },
        },
        Egress = new[]
        {
            new Aws.Ec2.Inputs.DefaultSecurityGroupEgressArgs
            {
                FromPort = 0,
                ToPort = 0,
                Protocol = "-1",
                CidrBlocks = new[]
                {
                    "0.0.0.0/0",
                },
            },
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
		mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewDefaultSecurityGroup(ctx, "default", &ec2.DefaultSecurityGroupArgs{
			VpcId: mainvpc.ID(),
			Ingress: ec2.DefaultSecurityGroupIngressArray{
				&ec2.DefaultSecurityGroupIngressArgs{
					Protocol: pulumi.String("-1"),
					Self:     pulumi.Bool(true),
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
				},
			},
			Egress: ec2.DefaultSecurityGroupEgressArray{
				&ec2.DefaultSecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /example %}}
{{% example %}}
### Example Config To Deny All Egress Traffic, Allowing Ingress

The following denies all Egress traffic by omitting any `egress` rules, while including the default `ingress` rule to allow all traffic.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainvpc = new aws.ec2.Vpc("mainvpc", {cidrBlock: "10.1.0.0/16"});
const _default = new aws.ec2.DefaultSecurityGroup("default", {
    vpcId: mainvpc.id,
    ingress: [{
        protocol: "-1",
        self: true,
        fromPort: 0,
        toPort: 0,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

mainvpc = aws.ec2.Vpc("mainvpc", cidr_block="10.1.0.0/16")
default = aws.ec2.DefaultSecurityGroup("default",
    vpc_id=mainvpc.id,
    ingress=[aws.ec2.DefaultSecurityGroupIngressArgs(
        protocol="-1",
        self=True,
        from_port=0,
        to_port=0,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mainvpc = new Aws.Ec2.Vpc("mainvpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var @default = new Aws.Ec2.DefaultSecurityGroup("default", new()
    {
        VpcId = mainvpc.Id,
        Ingress = new[]
        {
            new Aws.Ec2.Inputs.DefaultSecurityGroupIngressArgs
            {
                Protocol = "-1",
                Self = true,
                FromPort = 0,
                ToPort = 0,
            },
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
		mainvpc, err := ec2.NewVpc(ctx, "mainvpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewDefaultSecurityGroup(ctx, "default", &ec2.DefaultSecurityGroupArgs{
			VpcId: mainvpc.ID(),
			Ingress: ec2.DefaultSecurityGroupIngressArray{
				&ec2.DefaultSecurityGroupIngressArgs{
					Protocol: pulumi.String("-1"),
					Self:     pulumi.Bool(true),
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /example %}}
### Removing `aws.ec2.DefaultSecurityGroup` From Your Configuration

Removing this resource from your configuration will remove it from your statefile and management, but will not destroy the Security Group. All ingress or egress rules will be left as they are at the time of removal. You can resume managing them via the AWS Console.
{{% /examples %}}

## Import

Using `pulumi import`, import Security Groups using the security group `id`. For example:

```sh
 $ pulumi import aws:ec2/defaultSecurityGroup:DefaultSecurityGroup default_sg sg-903004f8
```
 