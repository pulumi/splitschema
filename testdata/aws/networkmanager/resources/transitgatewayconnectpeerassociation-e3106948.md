Associates a transit gateway Connect peer with a device, and optionally, with a link.
If you specify a link, it must be associated with the specified device.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.TransitGatewayConnectPeerAssociation("example", {
    globalNetworkId: aws_networkmanager_global_network.example.id,
    deviceId: aws_networkmanager_device.example.id,
    transitGatewayConnectPeerArn: aws_ec2_transit_gateway_connect_peer.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.TransitGatewayConnectPeerAssociation("example",
    global_network_id=aws_networkmanager_global_network["example"]["id"],
    device_id=aws_networkmanager_device["example"]["id"],
    transit_gateway_connect_peer_arn=aws_ec2_transit_gateway_connect_peer["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.TransitGatewayConnectPeerAssociation("example", new()
    {
        GlobalNetworkId = aws_networkmanager_global_network.Example.Id,
        DeviceId = aws_networkmanager_device.Example.Id,
        TransitGatewayConnectPeerArn = aws_ec2_transit_gateway_connect_peer.Example.Arn,
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
		_, err := networkmanager.NewTransitGatewayConnectPeerAssociation(ctx, "example", &networkmanager.TransitGatewayConnectPeerAssociationArgs{
			GlobalNetworkId:              pulumi.Any(aws_networkmanager_global_network.Example.Id),
			DeviceId:                     pulumi.Any(aws_networkmanager_device.Example.Id),
			TransitGatewayConnectPeerArn: pulumi.Any(aws_ec2_transit_gateway_connect_peer.Example.Arn),
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
import com.pulumi.aws.networkmanager.TransitGatewayConnectPeerAssociation;
import com.pulumi.aws.networkmanager.TransitGatewayConnectPeerAssociationArgs;
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
        var example = new TransitGatewayConnectPeerAssociation("example", TransitGatewayConnectPeerAssociationArgs.builder()        
            .globalNetworkId(aws_networkmanager_global_network.example().id())
            .deviceId(aws_networkmanager_device.example().id())
            .transitGatewayConnectPeerArn(aws_ec2_transit_gateway_connect_peer.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:TransitGatewayConnectPeerAssociation
    properties:
      globalNetworkId: ${aws_networkmanager_global_network.example.id}
      deviceId: ${aws_networkmanager_device.example.id}
      transitGatewayConnectPeerArn: ${aws_ec2_transit_gateway_connect_peer.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_transit_gateway_connect_peer_association` using the global network ID and customer gateway ARN. For example:

```sh
 $ pulumi import aws:networkmanager/transitGatewayConnectPeerAssociation:TransitGatewayConnectPeerAssociation example global-network-0d47f6t230mz46dy4,arn:aws:ec2:us-west-2:123456789012:transit-gateway-connect-peer/tgw-connect-peer-12345678
```
 