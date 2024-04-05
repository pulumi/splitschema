Manages an EC2 Transit Gateway Route Table.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.RouteTable("example", {transitGatewayId: aws_ec2_transit_gateway.example.id});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.RouteTable("example", transit_gateway_id=aws_ec2_transit_gateway["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.RouteTable("example", new()
    {
        TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
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
		_, err := ec2transitgateway.NewRouteTable(ctx, "example", &ec2transitgateway.RouteTableArgs{
			TransitGatewayId: pulumi.Any(aws_ec2_transit_gateway.Example.Id),
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
import com.pulumi.aws.ec2transitgateway.RouteTable;
import com.pulumi.aws.ec2transitgateway.RouteTableArgs;
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
        var example = new RouteTable("example", RouteTableArgs.builder()        
            .transitGatewayId(aws_ec2_transit_gateway.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:RouteTable
    properties:
      transitGatewayId: ${aws_ec2_transit_gateway.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_route_table` using the EC2 Transit Gateway Route Table identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/routeTable:RouteTable example tgw-rtb-12345678
```
 