Provides a MemoryDB Subnet Group.

More information about subnet groups can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/subnetgroups.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const exampleSubnet = new aws.ec2.Subnet("exampleSubnet", {
    vpcId: exampleVpc.id,
    cidrBlock: "10.0.0.0/24",
    availabilityZone: "us-west-2a",
});
const exampleSubnetGroup = new aws.memorydb.SubnetGroup("exampleSubnetGroup", {subnetIds: [exampleSubnet.id]});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_subnet = aws.ec2.Subnet("exampleSubnet",
    vpc_id=example_vpc.id,
    cidr_block="10.0.0.0/24",
    availability_zone="us-west-2a")
example_subnet_group = aws.memorydb.SubnetGroup("exampleSubnetGroup", subnet_ids=[example_subnet.id])
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
        CidrBlock = "10.0.0.0/16",
    });

    var exampleSubnet = new Aws.Ec2.Subnet("exampleSubnet", new()
    {
        VpcId = exampleVpc.Id,
        CidrBlock = "10.0.0.0/24",
        AvailabilityZone = "us-west-2a",
    });

    var exampleSubnetGroup = new Aws.MemoryDb.SubnetGroup("exampleSubnetGroup", new()
    {
        SubnetIds = new[]
        {
            exampleSubnet.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/memorydb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		exampleSubnet, err := ec2.NewSubnet(ctx, "exampleSubnet", &ec2.SubnetArgs{
			VpcId:            exampleVpc.ID(),
			CidrBlock:        pulumi.String("10.0.0.0/24"),
			AvailabilityZone: pulumi.String("us-west-2a"),
		})
		if err != nil {
			return err
		}
		_, err = memorydb.NewSubnetGroup(ctx, "exampleSubnetGroup", &memorydb.SubnetGroupArgs{
			SubnetIds: pulumi.StringArray{
				exampleSubnet.ID(),
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
import com.pulumi.aws.memorydb.SubnetGroup;
import com.pulumi.aws.memorydb.SubnetGroupArgs;
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
            .cidrBlock("10.0.0.0/16")
            .build());

        var exampleSubnet = new Subnet("exampleSubnet", SubnetArgs.builder()        
            .vpcId(exampleVpc.id())
            .cidrBlock("10.0.0.0/24")
            .availabilityZone("us-west-2a")
            .build());

        var exampleSubnetGroup = new SubnetGroup("exampleSubnetGroup", SubnetGroupArgs.builder()        
            .subnetIds(exampleSubnet.id())
            .build());

    }
}
```
```yaml
resources:
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  exampleSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${exampleVpc.id}
      cidrBlock: 10.0.0.0/24
      availabilityZone: us-west-2a
  exampleSubnetGroup:
    type: aws:memorydb:SubnetGroup
    properties:
      subnetIds:
        - ${exampleSubnet.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a subnet group using its `name`. For example:

```sh
 $ pulumi import aws:memorydb/subnetGroup:SubnetGroup example my-subnet-group
```
 