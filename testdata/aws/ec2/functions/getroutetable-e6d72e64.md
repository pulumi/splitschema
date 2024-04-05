`aws.ec2.RouteTable` provides details about a specific Route Table.

This resource can prove useful when a module accepts a Subnet ID as an input variable and needs to, for example, add a route in the Route Table.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how one might accept a Route Table ID as a variable and use this data source to obtain the data necessary to create a route.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const subnetId = config.requireObject("subnetId");
const selected = aws.ec2.getRouteTable({
    subnetId: subnetId,
});
const route = new aws.ec2.Route("route", {
    routeTableId: selected.then(selected => selected.id),
    destinationCidrBlock: "10.0.1.0/22",
    vpcPeeringConnectionId: "pcx-45ff3dc1",
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
subnet_id = config.require_object("subnetId")
selected = aws.ec2.get_route_table(subnet_id=subnet_id)
route = aws.ec2.Route("route",
    route_table_id=selected.id,
    destination_cidr_block="10.0.1.0/22",
    vpc_peering_connection_id="pcx-45ff3dc1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var subnetId = config.RequireObject<dynamic>("subnetId");
    var selected = Aws.Ec2.GetRouteTable.Invoke(new()
    {
        SubnetId = subnetId,
    });

    var route = new Aws.Ec2.Route("route", new()
    {
        RouteTableId = selected.Apply(getRouteTableResult => getRouteTableResult.Id),
        DestinationCidrBlock = "10.0.1.0/22",
        VpcPeeringConnectionId = "pcx-45ff3dc1",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		subnetId := cfg.RequireObject("subnetId")
		selected, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
			SubnetId: pulumi.StringRef(subnetId),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			RouteTableId:           *pulumi.String(selected.Id),
			DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
			VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
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
import com.pulumi.aws.ec2.inputs.GetRouteTableArgs;
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
        final var config = ctx.config();
        final var subnetId = config.get("subnetId");
        final var selected = Ec2Functions.getRouteTable(GetRouteTableArgs.builder()
            .subnetId(subnetId)
            .build());

        var route = new Route("route", RouteArgs.builder()        
            .routeTableId(selected.applyValue(getRouteTableResult -> getRouteTableResult.id()))
            .destinationCidrBlock("10.0.1.0/22")
            .vpcPeeringConnectionId("pcx-45ff3dc1")
            .build());

    }
}
```
```yaml
configuration:
  subnetId:
    type: dynamic
resources:
  route:
    type: aws:ec2:Route
    properties:
      routeTableId: ${selected.id}
      destinationCidrBlock: 10.0.1.0/22
      vpcPeeringConnectionId: pcx-45ff3dc1
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getRouteTable
      Arguments:
        subnetId: ${subnetId}
```
{{% /example %}}
{{% /examples %}}