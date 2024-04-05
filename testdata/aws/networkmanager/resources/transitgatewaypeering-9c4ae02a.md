Creates a peering connection between an AWS Cloud WAN core network and an AWS Transit Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.TransitGatewayPeering("example", {
    coreNetworkId: awscc_networkmanager_core_network.example.id,
    transitGatewayArn: aws_ec2_transit_gateway.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.TransitGatewayPeering("example",
    core_network_id=awscc_networkmanager_core_network["example"]["id"],
    transit_gateway_arn=aws_ec2_transit_gateway["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.TransitGatewayPeering("example", new()
    {
        CoreNetworkId = awscc_networkmanager_core_network.Example.Id,
        TransitGatewayArn = aws_ec2_transit_gateway.Example.Arn,
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
		_, err := networkmanager.NewTransitGatewayPeering(ctx, "example", &networkmanager.TransitGatewayPeeringArgs{
			CoreNetworkId:     pulumi.Any(awscc_networkmanager_core_network.Example.Id),
			TransitGatewayArn: pulumi.Any(aws_ec2_transit_gateway.Example.Arn),
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
import com.pulumi.aws.networkmanager.TransitGatewayPeering;
import com.pulumi.aws.networkmanager.TransitGatewayPeeringArgs;
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
        var example = new TransitGatewayPeering("example", TransitGatewayPeeringArgs.builder()        
            .coreNetworkId(awscc_networkmanager_core_network.example().id())
            .transitGatewayArn(aws_ec2_transit_gateway.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:TransitGatewayPeering
    properties:
      coreNetworkId: ${awscc_networkmanager_core_network.example.id}
      transitGatewayArn: ${aws_ec2_transit_gateway.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_transit_gateway_peering` using the peering ID. For example:

```sh
 $ pulumi import aws:networkmanager/transitGatewayPeering:TransitGatewayPeering example peering-444555aaabbb11223
```
 