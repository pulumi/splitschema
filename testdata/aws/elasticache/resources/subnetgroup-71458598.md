Provides an ElastiCache Subnet Group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("fooVpc", {
    cidrBlock: "10.0.0.0/16",
    tags: {
        Name: "tf-test",
    },
});
const fooSubnet = new aws.ec2.Subnet("fooSubnet", {
    vpcId: fooVpc.id,
    cidrBlock: "10.0.0.0/24",
    availabilityZone: "us-west-2a",
    tags: {
        Name: "tf-test",
    },
});
const bar = new aws.elasticache.SubnetGroup("bar", {subnetIds: [fooSubnet.id]});
```
```python
import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc",
    cidr_block="10.0.0.0/16",
    tags={
        "Name": "tf-test",
    })
foo_subnet = aws.ec2.Subnet("fooSubnet",
    vpc_id=foo_vpc.id,
    cidr_block="10.0.0.0/24",
    availability_zone="us-west-2a",
    tags={
        "Name": "tf-test",
    })
bar = aws.elasticache.SubnetGroup("bar", subnet_ids=[foo_subnet.id])
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
        CidrBlock = "10.0.0.0/16",
        Tags = 
        {
            { "Name", "tf-test" },
        },
    });

    var fooSubnet = new Aws.Ec2.Subnet("fooSubnet", new()
    {
        VpcId = fooVpc.Id,
        CidrBlock = "10.0.0.0/24",
        AvailabilityZone = "us-west-2a",
        Tags = 
        {
            { "Name", "tf-test" },
        },
    });

    var bar = new Aws.ElastiCache.SubnetGroup("bar", new()
    {
        SubnetIds = new[]
        {
            fooSubnet.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test"),
			},
		})
		if err != nil {
			return err
		}
		fooSubnet, err := ec2.NewSubnet(ctx, "fooSubnet", &ec2.SubnetArgs{
			VpcId:            fooVpc.ID(),
			CidrBlock:        pulumi.String("10.0.0.0/24"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test"),
			},
		})
		if err != nil {
			return err
		}
		_, err = elasticache.NewSubnetGroup(ctx, "bar", &elasticache.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				fooSubnet.ID(),
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
import com.pulumi.aws.elasticache.SubnetGroup;
import com.pulumi.aws.elasticache.SubnetGroupArgs;
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
            .cidrBlock("10.0.0.0/16")
            .tags(Map.of("Name", "tf-test"))
            .build());

        var fooSubnet = new Subnet("fooSubnet", SubnetArgs.builder()        
            .vpcId(fooVpc.id())
            .cidrBlock("10.0.0.0/24")
            .availabilityZone("us-west-2a")
            .tags(Map.of("Name", "tf-test"))
            .build());

        var bar = new SubnetGroup("bar", SubnetGroupArgs.builder()        
            .subnetIds(fooSubnet.id())
            .build());

    }
}
```
```yaml
resources:
  fooVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
      tags:
        Name: tf-test
  fooSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${fooVpc.id}
      cidrBlock: 10.0.0.0/24
      availabilityZone: us-west-2a
      tags:
        Name: tf-test
  bar:
    type: aws:elasticache:SubnetGroup
    properties:
      subnetIds:
        - ${fooSubnet.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache Subnet Groups using the `name`. For example:

```sh
 $ pulumi import aws:elasticache/subnetGroup:SubnetGroup bar tf-test-cache-subnet
```
 