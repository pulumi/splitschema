The VPC Peering Connection data source provides details about
a specific VPC peering connection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pc = aws.ec2.getVpcPeeringConnection({
    vpcId: aws_vpc.foo.id,
    peerCidrBlock: "10.0.1.0/22",
});
// Create a route table
const rt = new aws.ec2.RouteTable("rt", {vpcId: aws_vpc.foo.id});
// Create a route
const route = new aws.ec2.Route("route", {
    routeTableId: rt.id,
    destinationCidrBlock: pc.then(pc => pc.peerCidrBlock),
    vpcPeeringConnectionId: pc.then(pc => pc.id),
});
```
```python
import pulumi
import pulumi_aws as aws

pc = aws.ec2.get_vpc_peering_connection(vpc_id=aws_vpc["foo"]["id"],
    peer_cidr_block="10.0.1.0/22")
# Create a route table
rt = aws.ec2.RouteTable("rt", vpc_id=aws_vpc["foo"]["id"])
# Create a route
route = aws.ec2.Route("route",
    route_table_id=rt.id,
    destination_cidr_block=pc.peer_cidr_block,
    vpc_peering_connection_id=pc.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pc = Aws.Ec2.GetVpcPeeringConnection.Invoke(new()
    {
        VpcId = aws_vpc.Foo.Id,
        PeerCidrBlock = "10.0.1.0/22",
    });

    // Create a route table
    var rt = new Aws.Ec2.RouteTable("rt", new()
    {
        VpcId = aws_vpc.Foo.Id,
    });

    // Create a route
    var route = new Aws.Ec2.Route("route", new()
    {
        RouteTableId = rt.Id,
        DestinationCidrBlock = pc.Apply(getVpcPeeringConnectionResult => getVpcPeeringConnectionResult.PeerCidrBlock),
        VpcPeeringConnectionId = pc.Apply(getVpcPeeringConnectionResult => getVpcPeeringConnectionResult.Id),
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
		pc, err := ec2.LookupVpcPeeringConnection(ctx, &ec2.LookupVpcPeeringConnectionArgs{
			VpcId:         pulumi.StringRef(aws_vpc.Foo.Id),
			PeerCidrBlock: pulumi.StringRef("10.0.1.0/22"),
		}, nil)
		if err != nil {
			return err
		}
		rt, err := ec2.NewRouteTable(ctx, "rt", &ec2.RouteTableArgs{
			VpcId: pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			RouteTableId:           rt.ID(),
			DestinationCidrBlock:   *pulumi.String(pc.PeerCidrBlock),
			VpcPeeringConnectionId: *pulumi.String(pc.Id),
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
import com.pulumi.aws.ec2.inputs.GetVpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.RouteTable;
import com.pulumi.aws.ec2.RouteTableArgs;
import com.pulumi.aws.ec2.Route;
import com.pulumi.aws.ec2.RouteArgs;
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
        final var pc = Ec2Functions.getVpcPeeringConnection(GetVpcPeeringConnectionArgs.builder()
            .vpcId(aws_vpc.foo().id())
            .peerCidrBlock("10.0.1.0/22")
            .build());

        var rt = new RouteTable("rt", RouteTableArgs.builder()        
            .vpcId(aws_vpc.foo().id())
            .build());

        var route = new Route("route", RouteArgs.builder()        
            .routeTableId(rt.id())
            .destinationCidrBlock(pc.applyValue(getVpcPeeringConnectionResult -> getVpcPeeringConnectionResult.peerCidrBlock()))
            .vpcPeeringConnectionId(pc.applyValue(getVpcPeeringConnectionResult -> getVpcPeeringConnectionResult.id()))
            .build());

    }
}
```
```yaml
resources:
  # Create a route table
  rt:
    type: aws:ec2:RouteTable
    properties:
      vpcId: ${aws_vpc.foo.id}
  # Create a route
  route:
    type: aws:ec2:Route
    properties:
      routeTableId: ${rt.id}
      destinationCidrBlock: ${pc.peerCidrBlock}
      vpcPeeringConnectionId: ${pc.id}
variables:
  pc:
    fn::invoke:
      Function: aws:ec2:getVpcPeeringConnection
      Arguments:
        vpcId: ${aws_vpc.foo.id}
        peerCidrBlock: 10.0.1.0/22
```
{{% /example %}}
{{% /examples %}}