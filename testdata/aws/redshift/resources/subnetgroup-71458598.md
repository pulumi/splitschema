Creates a new Amazon Redshift subnet group. You must provide a list of one or more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when creating Amazon Redshift subnet group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"});
const fooSubnet = new aws.ec2.Subnet("fooSubnet", {
    cidrBlock: "10.1.1.0/24",
    availabilityZone: "us-west-2a",
    vpcId: fooVpc.id,
    tags: {
        Name: "tf-dbsubnet-test-1",
    },
});
const bar = new aws.ec2.Subnet("bar", {
    cidrBlock: "10.1.2.0/24",
    availabilityZone: "us-west-2b",
    vpcId: fooVpc.id,
    tags: {
        Name: "tf-dbsubnet-test-2",
    },
});
const fooSubnetGroup = new aws.redshift.SubnetGroup("fooSubnetGroup", {
    subnetIds: [
        fooSubnet.id,
        bar.id,
    ],
    tags: {
        environment: "Production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
foo_subnet = aws.ec2.Subnet("fooSubnet",
    cidr_block="10.1.1.0/24",
    availability_zone="us-west-2a",
    vpc_id=foo_vpc.id,
    tags={
        "Name": "tf-dbsubnet-test-1",
    })
bar = aws.ec2.Subnet("bar",
    cidr_block="10.1.2.0/24",
    availability_zone="us-west-2b",
    vpc_id=foo_vpc.id,
    tags={
        "Name": "tf-dbsubnet-test-2",
    })
foo_subnet_group = aws.redshift.SubnetGroup("fooSubnetGroup",
    subnet_ids=[
        foo_subnet.id,
        bar.id,
    ],
    tags={
        "environment": "Production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooVpc = new Aws.Ec2.Vpc("fooVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new()
    {
        CidrBlock = "10.1.1.0/24",
        AvailabilityZone = "us-west-2a",
        VpcId = fooVpc.Id,
        Tags = 
        {
            { "Name", "tf-dbsubnet-test-1" },
        },
    });

    var bar = new Aws.Ec2.Subnet("bar", new()
    {
        CidrBlock = "10.1.2.0/24",
        AvailabilityZone = "us-west-2b",
        VpcId = fooVpc.Id,
        Tags = 
        {
            { "Name", "tf-dbsubnet-test-2" },
        },
    });

    var fooSubnetGroup = new Aws.RedShift.SubnetGroup("fooSubnetGroup", new()
    {
        SubnetIds = new[]
        {
            fooSubnet.Id,
            bar.Id,
        },
        Tags = 
        {
            { "environment", "Production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
			CidrBlock:        pulumi.String("10.1.1.0/24"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			VpcId:            fooVpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-dbsubnet-test-1"),
			},
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewSubnet(ctx, "bar", &ec2.SubnetArgs{
			CidrBlock:        pulumi.String("10.1.2.0/24"),
			AvailabilityZone: pulumi.String("us-west-2b"),
			VpcId:            fooVpc.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-dbsubnet-test-2"),
			},
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewSubnetGroup(ctx, "fooSubnetGroup", &redshift.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				fooSubnet.ID(),
				bar.ID(),
			},
			Tags: pulumi.StringMap{
				"environment": pulumi.String("Production"),
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
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.redshift.SubnetGroup;
import com.pulumi.aws.redshift.SubnetGroupArgs;
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
        var fooVpc = new Vpc("fooVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        var fooSubnet = new Subnet("fooSubnet", SubnetArgs.builder()        
            .cidrBlock("10.1.1.0/24")
            .availabilityZone("us-west-2a")
            .vpcId(fooVpc.id())
            .tags(Map.of("Name", "tf-dbsubnet-test-1"))
            .build());

        var bar = new Subnet("bar", SubnetArgs.builder()        
            .cidrBlock("10.1.2.0/24")
            .availabilityZone("us-west-2b")
            .vpcId(fooVpc.id())
            .tags(Map.of("Name", "tf-dbsubnet-test-2"))
            .build());

        var fooSubnetGroup = new SubnetGroup("fooSubnetGroup", SubnetGroupArgs.builder()        
            .subnetIds(            
                fooSubnet.id(),
                bar.id())
            .tags(Map.of("environment", "Production"))
            .build());

    }
}
```
```yaml
resources:
  fooVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  fooSubnet:
    type: aws:ec2:Subnet
    properties:
      cidrBlock: 10.1.1.0/24
      availabilityZone: us-west-2a
      vpcId: ${fooVpc.id}
      tags:
        Name: tf-dbsubnet-test-1
  bar:
    type: aws:ec2:Subnet
    properties:
      cidrBlock: 10.1.2.0/24
      availabilityZone: us-west-2b
      vpcId: ${fooVpc.id}
      tags:
        Name: tf-dbsubnet-test-2
  fooSubnetGroup:
    type: aws:redshift:SubnetGroup
    properties:
      subnetIds:
        - ${fooSubnet.id}
        - ${bar.id}
      tags:
        environment: Production
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift subnet groups using the `name`. For example:

```sh
 $ pulumi import aws:redshift/subnetGroup:SubnetGroup testgroup1 test-cluster-subnet-group
```
 