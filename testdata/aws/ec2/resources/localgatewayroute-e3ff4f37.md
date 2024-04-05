Manages an EC2 Local Gateway Route. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.LocalGatewayRoute("example", {
    destinationCidrBlock: "172.16.0.0/16",
    localGatewayRouteTableId: data.aws_ec2_local_gateway_route_table.example.id,
    localGatewayVirtualInterfaceGroupId: data.aws_ec2_local_gateway_virtual_interface_group.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.LocalGatewayRoute("example",
    destination_cidr_block="172.16.0.0/16",
    local_gateway_route_table_id=data["aws_ec2_local_gateway_route_table"]["example"]["id"],
    local_gateway_virtual_interface_group_id=data["aws_ec2_local_gateway_virtual_interface_group"]["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.LocalGatewayRoute("example", new()
    {
        DestinationCidrBlock = "172.16.0.0/16",
        LocalGatewayRouteTableId = data.Aws_ec2_local_gateway_route_table.Example.Id,
        LocalGatewayVirtualInterfaceGroupId = data.Aws_ec2_local_gateway_virtual_interface_group.Example.Id,
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
		_, err := ec2.NewLocalGatewayRoute(ctx, "example", &ec2.LocalGatewayRouteArgs{
			DestinationCidrBlock:                pulumi.String("172.16.0.0/16"),
			LocalGatewayRouteTableId:            pulumi.Any(data.Aws_ec2_local_gateway_route_table.Example.Id),
			LocalGatewayVirtualInterfaceGroupId: pulumi.Any(data.Aws_ec2_local_gateway_virtual_interface_group.Example.Id),
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
import com.pulumi.aws.ec2.LocalGatewayRoute;
import com.pulumi.aws.ec2.LocalGatewayRouteArgs;
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
        var example = new LocalGatewayRoute("example", LocalGatewayRouteArgs.builder()        
            .destinationCidrBlock("172.16.0.0/16")
            .localGatewayRouteTableId(data.aws_ec2_local_gateway_route_table().example().id())
            .localGatewayVirtualInterfaceGroupId(data.aws_ec2_local_gateway_virtual_interface_group().example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:LocalGatewayRoute
    properties:
      destinationCidrBlock: 172.16.0.0/16
      localGatewayRouteTableId: ${data.aws_ec2_local_gateway_route_table.example.id}
      localGatewayVirtualInterfaceGroupId: ${data.aws_ec2_local_gateway_virtual_interface_group.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_local_gateway_route` using the EC2 Local Gateway Route Table identifier and destination CIDR block separated by underscores (`_`). For example:

```sh
 $ pulumi import aws:ec2/localGatewayRoute:LocalGatewayRoute example lgw-rtb-12345678_172.16.0.0/16
```
 