Authorizes a VPC in a different account to be associated with a local Route53 Hosted Zone.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const alternate = new aws.Provider("alternate", {});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
    cidrBlock: "10.6.0.0/16",
    enableDnsHostnames: true,
    enableDnsSupport: true,
});
const exampleZone = new aws.route53.Zone("exampleZone", {vpcs: [{
    vpcId: exampleVpc.id,
}]});
const alternateVpc = new aws.ec2.Vpc("alternateVpc", {
    cidrBlock: "10.7.0.0/16",
    enableDnsHostnames: true,
    enableDnsSupport: true,
}, {
    provider: aws.alternate,
});
const exampleVpcAssociationAuthorization = new aws.route53.VpcAssociationAuthorization("exampleVpcAssociationAuthorization", {
    vpcId: alternateVpc.id,
    zoneId: exampleZone.id,
});
const exampleZoneAssociation = new aws.route53.ZoneAssociation("exampleZoneAssociation", {
    vpcId: exampleVpcAssociationAuthorization.vpcId,
    zoneId: exampleVpcAssociationAuthorization.zoneId,
}, {
    provider: aws.alternate,
});
```
```python
import pulumi
import pulumi_aws as aws

alternate = aws.Provider("alternate")
example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="10.6.0.0/16",
    enable_dns_hostnames=True,
    enable_dns_support=True)
example_zone = aws.route53.Zone("exampleZone", vpcs=[aws.route53.ZoneVpcArgs(
    vpc_id=example_vpc.id,
)])
alternate_vpc = aws.ec2.Vpc("alternateVpc",
    cidr_block="10.7.0.0/16",
    enable_dns_hostnames=True,
    enable_dns_support=True,
    opts=pulumi.ResourceOptions(provider=aws["alternate"]))
example_vpc_association_authorization = aws.route53.VpcAssociationAuthorization("exampleVpcAssociationAuthorization",
    vpc_id=alternate_vpc.id,
    zone_id=example_zone.id)
example_zone_association = aws.route53.ZoneAssociation("exampleZoneAssociation",
    vpc_id=example_vpc_association_authorization.vpc_id,
    zone_id=example_vpc_association_authorization.zone_id,
    opts=pulumi.ResourceOptions(provider=aws["alternate"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var alternate = new Aws.Provider("alternate");

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.6.0.0/16",
        EnableDnsHostnames = true,
        EnableDnsSupport = true,
    });

    var exampleZone = new Aws.Route53.Zone("exampleZone", new()
    {
        Vpcs = new[]
        {
            new Aws.Route53.Inputs.ZoneVpcArgs
            {
                VpcId = exampleVpc.Id,
            },
        },
    });

    var alternateVpc = new Aws.Ec2.Vpc("alternateVpc", new()
    {
        CidrBlock = "10.7.0.0/16",
        EnableDnsHostnames = true,
        EnableDnsSupport = true,
    }, new CustomResourceOptions
    {
        Provider = aws.Alternate,
    });

    var exampleVpcAssociationAuthorization = new Aws.Route53.VpcAssociationAuthorization("exampleVpcAssociationAuthorization", new()
    {
        VpcId = alternateVpc.Id,
        ZoneId = exampleZone.Id,
    });

    var exampleZoneAssociation = new Aws.Route53.ZoneAssociation("exampleZoneAssociation", new()
    {
        VpcId = exampleVpcAssociationAuthorization.VpcId,
        ZoneId = exampleVpcAssociationAuthorization.ZoneId,
    }, new CustomResourceOptions
    {
        Provider = aws.Alternate,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "alternate", nil)
		if err != nil {
			return err
		}
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.6.0.0/16"),
			EnableDnsHostnames: pulumi.Bool(true),
			EnableDnsSupport:   pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		exampleZone, err := route53.NewZone(ctx, "exampleZone", &route53.ZoneArgs{
			Vpcs: route53.ZoneVpcArray{
				&route53.ZoneVpcArgs{
					VpcId: exampleVpc.ID(),
				},
			},
		})
		if err != nil {
			return err
		}
		alternateVpc, err := ec2.NewVpc(ctx, "alternateVpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.7.0.0/16"),
			EnableDnsHostnames: pulumi.Bool(true),
			EnableDnsSupport:   pulumi.Bool(true),
		}, pulumi.Provider(aws.Alternate))
		if err != nil {
			return err
		}
		exampleVpcAssociationAuthorization, err := route53.NewVpcAssociationAuthorization(ctx, "exampleVpcAssociationAuthorization", &route53.VpcAssociationAuthorizationArgs{
			VpcId:  alternateVpc.ID(),
			ZoneId: exampleZone.ID(),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewZoneAssociation(ctx, "exampleZoneAssociation", &route53.ZoneAssociationArgs{
			VpcId:  exampleVpcAssociationAuthorization.VpcId,
			ZoneId: exampleVpcAssociationAuthorization.ZoneId,
		}, pulumi.Provider(aws.Alternate))
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.route53.Zone;
import com.pulumi.aws.route53.ZoneArgs;
import com.pulumi.aws.route53.inputs.ZoneVpcArgs;
import com.pulumi.aws.route53.VpcAssociationAuthorization;
import com.pulumi.aws.route53.VpcAssociationAuthorizationArgs;
import com.pulumi.aws.route53.ZoneAssociation;
import com.pulumi.aws.route53.ZoneAssociationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var alternate = new Provider("alternate");

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.6.0.0/16")
            .enableDnsHostnames(true)
            .enableDnsSupport(true)
            .build());

        var exampleZone = new Zone("exampleZone", ZoneArgs.builder()        
            .vpcs(ZoneVpcArgs.builder()
                .vpcId(exampleVpc.id())
                .build())
            .build());

        var alternateVpc = new Vpc("alternateVpc", VpcArgs.builder()        
            .cidrBlock("10.7.0.0/16")
            .enableDnsHostnames(true)
            .enableDnsSupport(true)
            .build(), CustomResourceOptions.builder()
                .provider(aws.alternate())
                .build());

        var exampleVpcAssociationAuthorization = new VpcAssociationAuthorization("exampleVpcAssociationAuthorization", VpcAssociationAuthorizationArgs.builder()        
            .vpcId(alternateVpc.id())
            .zoneId(exampleZone.id())
            .build());

        var exampleZoneAssociation = new ZoneAssociation("exampleZoneAssociation", ZoneAssociationArgs.builder()        
            .vpcId(exampleVpcAssociationAuthorization.vpcId())
            .zoneId(exampleVpcAssociationAuthorization.zoneId())
            .build(), CustomResourceOptions.builder()
                .provider(aws.alternate())
                .build());

    }
}
```
```yaml
resources:
  alternate:
    type: pulumi:providers:aws
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.6.0.0/16
      enableDnsHostnames: true
      enableDnsSupport: true
  exampleZone:
    type: aws:route53:Zone
    properties:
      vpcs:
        - vpcId: ${exampleVpc.id}
  alternateVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.7.0.0/16
      enableDnsHostnames: true
      enableDnsSupport: true
    options:
      provider: ${aws.alternate}
  exampleVpcAssociationAuthorization:
    type: aws:route53:VpcAssociationAuthorization
    properties:
      vpcId: ${alternateVpc.id}
      zoneId: ${exampleZone.id}
  exampleZoneAssociation:
    type: aws:route53:ZoneAssociation
    properties:
      vpcId: ${exampleVpcAssociationAuthorization.vpcId}
      zoneId: ${exampleVpcAssociationAuthorization.zoneId}
    options:
      provider: ${aws.alternate}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route 53 VPC Association Authorizations using the Hosted Zone ID and VPC ID, separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization example Z123456ABCDEFG:vpc-12345678
```
 