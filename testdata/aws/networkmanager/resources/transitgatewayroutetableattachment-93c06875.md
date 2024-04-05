Creates a transit gateway route table attachment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.TransitGatewayRouteTableAttachment("example", {
    peeringId: aws_networkmanager_transit_gateway_peering.example.id,
    transitGatewayRouteTableArn: aws_ec2_transit_gateway_route_table.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.TransitGatewayRouteTableAttachment("example",
    peering_id=aws_networkmanager_transit_gateway_peering["example"]["id"],
    transit_gateway_route_table_arn=aws_ec2_transit_gateway_route_table["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.TransitGatewayRouteTableAttachment("example", new()
    {
        PeeringId = aws_networkmanager_transit_gateway_peering.Example.Id,
        TransitGatewayRouteTableArn = aws_ec2_transit_gateway_route_table.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkmanager.NewTransitGatewayRouteTableAttachment(ctx, "example", &networkmanager.TransitGatewayRouteTableAttachmentArgs{
			PeeringId:                   pulumi.Any(aws_networkmanager_transit_gateway_peering.Example.Id),
			TransitGatewayRouteTableArn: pulumi.Any(aws_ec2_transit_gateway_route_table.Example.Arn),
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
import com.pulumi.aws.networkmanager.TransitGatewayRouteTableAttachment;
import com.pulumi.aws.networkmanager.TransitGatewayRouteTableAttachmentArgs;
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
        var example = new TransitGatewayRouteTableAttachment("example", TransitGatewayRouteTableAttachmentArgs.builder()        
            .peeringId(aws_networkmanager_transit_gateway_peering.example().id())
            .transitGatewayRouteTableArn(aws_ec2_transit_gateway_route_table.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:TransitGatewayRouteTableAttachment
    properties:
      peeringId: ${aws_networkmanager_transit_gateway_peering.example.id}
      transitGatewayRouteTableArn: ${aws_ec2_transit_gateway_route_table.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_transit_gateway_route_table_attachment` using the attachment ID. For example:

```sh
 $ pulumi import aws:networkmanager/transitGatewayRouteTableAttachment:TransitGatewayRouteTableAttachment example attachment-0f8fa60d2238d1bd8
```
 