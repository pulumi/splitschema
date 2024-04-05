Manages an EC2 Transit Gateway Route.

{{% examples %}}
## Example Usage
{{% example %}}
### Standard usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.Route("example", {
    destinationCidrBlock: "0.0.0.0/0",
    transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment.example.id,
    transitGatewayRouteTableId: aws_ec2_transit_gateway.example.association_default_route_table_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.Route("example",
    destination_cidr_block="0.0.0.0/0",
    transit_gateway_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"],
    transit_gateway_route_table_id=aws_ec2_transit_gateway["example"]["association_default_route_table_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.Route("example", new()
    {
        DestinationCidrBlock = "0.0.0.0/0",
        TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
        TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewRoute(ctx, "example", &ec2transitgateway.RouteArgs{
			DestinationCidrBlock:       pulumi.String("0.0.0.0/0"),
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
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
import com.pulumi.aws.ec2transitgateway.Route;
import com.pulumi.aws.ec2transitgateway.RouteArgs;
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
        var example = new Route("example", RouteArgs.builder()        
            .destinationCidrBlock("0.0.0.0/0")
            .transitGatewayAttachmentId(aws_ec2_transit_gateway_vpc_attachment.example().id())
            .transitGatewayRouteTableId(aws_ec2_transit_gateway.example().association_default_route_table_id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:Route
    properties:
      destinationCidrBlock: 0.0.0.0/0
      transitGatewayAttachmentId: ${aws_ec2_transit_gateway_vpc_attachment.example.id}
      transitGatewayRouteTableId: ${aws_ec2_transit_gateway.example.association_default_route_table_id}
```
{{% /example %}}
{{% example %}}
### Blackhole route

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.Route("example", {
    destinationCidrBlock: "0.0.0.0/0",
    blackhole: true,
    transitGatewayRouteTableId: aws_ec2_transit_gateway.example.association_default_route_table_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.Route("example",
    destination_cidr_block="0.0.0.0/0",
    blackhole=True,
    transit_gateway_route_table_id=aws_ec2_transit_gateway["example"]["association_default_route_table_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.Route("example", new()
    {
        DestinationCidrBlock = "0.0.0.0/0",
        Blackhole = true,
        TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.NewRoute(ctx, "example", &ec2transitgateway.RouteArgs{
			DestinationCidrBlock:       pulumi.String("0.0.0.0/0"),
			Blackhole:                  pulumi.Bool(true),
			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
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
import com.pulumi.aws.ec2transitgateway.Route;
import com.pulumi.aws.ec2transitgateway.RouteArgs;
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
        var example = new Route("example", RouteArgs.builder()        
            .destinationCidrBlock("0.0.0.0/0")
            .blackhole(true)
            .transitGatewayRouteTableId(aws_ec2_transit_gateway.example().association_default_route_table_id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:Route
    properties:
      destinationCidrBlock: 0.0.0.0/0
      blackhole: true
      transitGatewayRouteTableId: ${aws_ec2_transit_gateway.example.association_default_route_table_id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_route` using the EC2 Transit Gateway Route Table, an underscore, and the destination. For example:

```sh
 $ pulumi import aws:ec2transitgateway/route:Route example tgw-rtb-12345678_0.0.0.0/0
```
 