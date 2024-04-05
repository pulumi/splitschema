Provides a resource to create a routing table entry (a route) in a VPC routing table.

> **NOTE on Route Tables and Routes:** This provider currently provides both a standalone Route resource and a Route Table resource with routes defined in-line. At this time you cannot use a Route Table with in-line routes in conjunction with any Route resources. Doing so will cause a conflict of rule settings and will overwrite rules.

> **NOTE on `gateway_id` attribute:** The AWS API is very forgiving with the resource ID passed in the `gateway_id` attribute. For example an `aws.ec2.Route` resource can be created with an `aws.ec2.NatGateway` or `aws.ec2.EgressOnlyInternetGateway` ID specified for the `gateway_id` attribute. Specifying anything other than an `aws.ec2.InternetGateway` or `aws.ec2.VpnGateway` ID will lead to this provider reporting a permanent diff between your configuration and recorded state, as the AWS API returns the more-specific attribute. If you are experiencing constant diffs with an `aws.ec2.Route` resource, the first thing to check is that the correct attribute is being specified.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const route = new aws.ec2.Route("route", {
    routeTableId: "rtb-4fbb3ac4",
    destinationCidrBlock: "10.0.1.0/22",
    vpcPeeringConnectionId: "pcx-45ff3dc1",
}, {
    dependsOn: [aws_route_table.testing],
});
```
```python
import pulumi
import pulumi_aws as aws

route = aws.ec2.Route("route",
    route_table_id="rtb-4fbb3ac4",
    destination_cidr_block="10.0.1.0/22",
    vpc_peering_connection_id="pcx-45ff3dc1",
    opts=pulumi.ResourceOptions(depends_on=[aws_route_table["testing"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var route = new Aws.Ec2.Route("route", new()
    {
        RouteTableId = "rtb-4fbb3ac4",
        DestinationCidrBlock = "10.0.1.0/22",
        VpcPeeringConnectionId = "pcx-45ff3dc1",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_route_table.Testing,
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
		_, err := ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			RouteTableId:           pulumi.String("rtb-4fbb3ac4"),
			DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
			VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_route_table.Testing,
		}))
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
import com.pulumi.aws.ec2.Route;
import com.pulumi.aws.ec2.RouteArgs;
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
        var route = new Route("route", RouteArgs.builder()        
            .routeTableId("rtb-4fbb3ac4")
            .destinationCidrBlock("10.0.1.0/22")
            .vpcPeeringConnectionId("pcx-45ff3dc1")
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_route_table.testing())
                .build());

    }
}
```
```yaml
resources:
  route:
    type: aws:ec2:Route
    properties:
      routeTableId: rtb-4fbb3ac4
      destinationCidrBlock: 10.0.1.0/22
      vpcPeeringConnectionId: pcx-45ff3dc1
    options:
      dependson:
        - ${aws_route_table.testing}
```
{{% /example %}}
{{% /examples %}}
## Example IPv6 Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpc = new aws.ec2.Vpc("vpc", {
    cidrBlock: "10.1.0.0/16",
    assignGeneratedIpv6CidrBlock: true,
});
const egress = new aws.ec2.EgressOnlyInternetGateway("egress", {vpcId: vpc.id});
const route = new aws.ec2.Route("route", {
    routeTableId: "rtb-4fbb3ac4",
    destinationIpv6CidrBlock: "::/0",
    egressOnlyGatewayId: egress.id,
});
```
```python
import pulumi
import pulumi_aws as aws

vpc = aws.ec2.Vpc("vpc",
    cidr_block="10.1.0.0/16",
    assign_generated_ipv6_cidr_block=True)
egress = aws.ec2.EgressOnlyInternetGateway("egress", vpc_id=vpc.id)
route = aws.ec2.Route("route",
    route_table_id="rtb-4fbb3ac4",
    destination_ipv6_cidr_block="::/0",
    egress_only_gateway_id=egress.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var vpc = new Aws.Ec2.Vpc("vpc", new()
    {
        CidrBlock = "10.1.0.0/16",
        AssignGeneratedIpv6CidrBlock = true,
    });

    var egress = new Aws.Ec2.EgressOnlyInternetGateway("egress", new()
    {
        VpcId = vpc.Id,
    });

    var route = new Aws.Ec2.Route("route", new()
    {
        RouteTableId = "rtb-4fbb3ac4",
        DestinationIpv6CidrBlock = "::/0",
        EgressOnlyGatewayId = egress.Id,
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
		vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
			CidrBlock:                    pulumi.String("10.1.0.0/16"),
			AssignGeneratedIpv6CidrBlock: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		egress, err := ec2.NewEgressOnlyInternetGateway(ctx, "egress", &ec2.EgressOnlyInternetGatewayArgs{
			VpcId: vpc.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
			RouteTableId:             pulumi.String("rtb-4fbb3ac4"),
			DestinationIpv6CidrBlock: pulumi.String("::/0"),
			EgressOnlyGatewayId:      egress.ID(),
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
        var vpc = new Vpc("vpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .assignGeneratedIpv6CidrBlock(true)
            .build());

        var egress = new EgressOnlyInternetGateway("egress", EgressOnlyInternetGatewayArgs.builder()        
            .vpcId(vpc.id())
            .build());

        var route = new Route("route", RouteArgs.builder()        
            .routeTableId("rtb-4fbb3ac4")
            .destinationIpv6CidrBlock("::/0")
            .egressOnlyGatewayId(egress.id())
            .build());

    }
}
```
```yaml
resources:
  vpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
      assignGeneratedIpv6CidrBlock: true
  egress:
    type: aws:ec2:EgressOnlyInternetGateway
    properties:
      vpcId: ${vpc.id}
  route:
    type: aws:ec2:Route
    properties:
      routeTableId: rtb-4fbb3ac4
      destinationIpv6CidrBlock: ::/0
      egressOnlyGatewayId: ${egress.id}
```


## Import

Import a route in route table `rtb-656C65616E6F72` with an IPv6 destination CIDR of `2620:0:2d0:200::8/125`:

Import a route in route table `rtb-656C65616E6F72` with a managed prefix list destination of `pl-0570a1d2d725c16be`:

__Using `pulumi import` to import__ individual routes using `ROUTETABLEID_DESTINATION`. Import [local routes](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html#RouteTables) using the VPC's IPv4 or IPv6 CIDR blocks. For example:

Import a route in route table `rtb-656C65616E6F72` with an IPv4 destination CIDR of `10.42.0.0/16`:

```sh
 $ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_10.42.0.0/16
```
 Import a route in route table `rtb-656C65616E6F72` with an IPv6 destination CIDR of `2620:0:2d0:200::8/125`:

```sh
 $ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_2620:0:2d0:200::8/125
```
 Import a route in route table `rtb-656C65616E6F72` with a managed prefix list destination of `pl-0570a1d2d725c16be`:

```sh
 $ pulumi import aws:ec2/route:Route my_route rtb-656C65616E6F72_pl-0570a1d2d725c16be
```
 