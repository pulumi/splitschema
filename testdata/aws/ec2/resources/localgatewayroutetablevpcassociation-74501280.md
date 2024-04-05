Manages an EC2 Local Gateway Route Table VPC Association. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-local-gateways.html#vpc-associations).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLocalGatewayRouteTable = aws.ec2.getLocalGatewayRouteTable({
    outpostArn: "arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef",
});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const exampleLocalGatewayRouteTableVpcAssociation = new aws.ec2.LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation", {
    localGatewayRouteTableId: exampleLocalGatewayRouteTable.then(exampleLocalGatewayRouteTable => exampleLocalGatewayRouteTable.id),
    vpcId: exampleVpc.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_local_gateway_route_table = aws.ec2.get_local_gateway_route_table(outpost_arn="arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_local_gateway_route_table_vpc_association = aws.ec2.LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation",
    local_gateway_route_table_id=example_local_gateway_route_table.id,
    vpc_id=example_vpc.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLocalGatewayRouteTable = Aws.Ec2.GetLocalGatewayRouteTable.Invoke(new()
    {
        OutpostArn = "arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef",
    });

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var exampleLocalGatewayRouteTableVpcAssociation = new Aws.Ec2.LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation", new()
    {
        LocalGatewayRouteTableId = exampleLocalGatewayRouteTable.Apply(getLocalGatewayRouteTableResult => getLocalGatewayRouteTableResult.Id),
        VpcId = exampleVpc.Id,
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
		exampleLocalGatewayRouteTable, err := ec2.GetLocalGatewayRouteTable(ctx, &ec2.GetLocalGatewayRouteTableArgs{
			OutpostArn: pulumi.StringRef("arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef"),
		}, nil)
		if err != nil {
			return err
		}
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewLocalGatewayRouteTableVpcAssociation(ctx, "exampleLocalGatewayRouteTableVpcAssociation", &ec2.LocalGatewayRouteTableVpcAssociationArgs{
			LocalGatewayRouteTableId: *pulumi.String(exampleLocalGatewayRouteTable.Id),
			VpcId:                    exampleVpc.ID(),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetLocalGatewayRouteTableArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.LocalGatewayRouteTableVpcAssociation;
import com.pulumi.aws.ec2.LocalGatewayRouteTableVpcAssociationArgs;
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
        final var exampleLocalGatewayRouteTable = Ec2Functions.getLocalGatewayRouteTable(GetLocalGatewayRouteTableArgs.builder()
            .outpostArn("arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef")
            .build());

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var exampleLocalGatewayRouteTableVpcAssociation = new LocalGatewayRouteTableVpcAssociation("exampleLocalGatewayRouteTableVpcAssociation", LocalGatewayRouteTableVpcAssociationArgs.builder()        
            .localGatewayRouteTableId(exampleLocalGatewayRouteTable.applyValue(getLocalGatewayRouteTableResult -> getLocalGatewayRouteTableResult.id()))
            .vpcId(exampleVpc.id())
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
  exampleLocalGatewayRouteTableVpcAssociation:
    type: aws:ec2:LocalGatewayRouteTableVpcAssociation
    properties:
      localGatewayRouteTableId: ${exampleLocalGatewayRouteTable.id}
      vpcId: ${exampleVpc.id}
variables:
  exampleLocalGatewayRouteTable:
    fn::invoke:
      Function: aws:ec2:getLocalGatewayRouteTable
      Arguments:
        outpostArn: arn:aws:outposts:us-west-2:123456789012:outpost/op-1234567890abcdef
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_local_gateway_route_table_vpc_association` using the Local Gateway Route Table VPC Association identifier. For example:

```sh
 $ pulumi import aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation example lgw-vpc-assoc-1234567890abcdef
```
 