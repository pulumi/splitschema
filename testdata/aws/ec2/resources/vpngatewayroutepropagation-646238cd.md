Requests automatic route propagation between a VPN gateway and a route table.

> **Note:** This resource should not be used with a route table that has
the `propagating_vgws` argument set. If that argument is set, any route
propagation not explicitly listed in its value will be removed.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpnGatewayRoutePropagation("example", {
    vpnGatewayId: aws_vpn_gateway.example.id,
    routeTableId: aws_route_table.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.VpnGatewayRoutePropagation("example",
    vpn_gateway_id=aws_vpn_gateway["example"]["id"],
    route_table_id=aws_route_table["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.VpnGatewayRoutePropagation("example", new()
    {
        VpnGatewayId = aws_vpn_gateway.Example.Id,
        RouteTableId = aws_route_table.Example.Id,
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
		_, err := ec2.NewVpnGatewayRoutePropagation(ctx, "example", &ec2.VpnGatewayRoutePropagationArgs{
			VpnGatewayId: pulumi.Any(aws_vpn_gateway.Example.Id),
			RouteTableId: pulumi.Any(aws_route_table.Example.Id),
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
import com.pulumi.aws.ec2.VpnGatewayRoutePropagation;
import com.pulumi.aws.ec2.VpnGatewayRoutePropagationArgs;
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
        var example = new VpnGatewayRoutePropagation("example", VpnGatewayRoutePropagationArgs.builder()        
            .vpnGatewayId(aws_vpn_gateway.example().id())
            .routeTableId(aws_route_table.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpnGatewayRoutePropagation
    properties:
      vpnGatewayId: ${aws_vpn_gateway.example.id}
      routeTableId: ${aws_route_table.example.id}
```
{{% /example %}}
{{% /examples %}}