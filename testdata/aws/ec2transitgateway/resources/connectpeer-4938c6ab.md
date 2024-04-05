Manages an EC2 Transit Gateway Connect Peer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConnect = new aws.ec2transitgateway.Connect("exampleConnect", {
    transportAttachmentId: aws_ec2_transit_gateway_vpc_attachment.example.id,
    transitGatewayId: aws_ec2_transit_gateway.example.id,
});
const exampleConnectPeer = new aws.ec2transitgateway.ConnectPeer("exampleConnectPeer", {
    peerAddress: "10.1.2.3",
    insideCidrBlocks: ["169.254.100.0/29"],
    transitGatewayAttachmentId: exampleConnect.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_connect = aws.ec2transitgateway.Connect("exampleConnect",
    transport_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"],
    transit_gateway_id=aws_ec2_transit_gateway["example"]["id"])
example_connect_peer = aws.ec2transitgateway.ConnectPeer("exampleConnectPeer",
    peer_address="10.1.2.3",
    inside_cidr_blocks=["169.254.100.0/29"],
    transit_gateway_attachment_id=example_connect.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConnect = new Aws.Ec2TransitGateway.Connect("exampleConnect", new()
    {
        TransportAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
        TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
    });

    var exampleConnectPeer = new Aws.Ec2TransitGateway.ConnectPeer("exampleConnectPeer", new()
    {
        PeerAddress = "10.1.2.3",
        InsideCidrBlocks = new[]
        {
            "169.254.100.0/29",
        },
        TransitGatewayAttachmentId = exampleConnect.Id,
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
		exampleConnect, err := ec2transitgateway.NewConnect(ctx, "exampleConnect", &ec2transitgateway.ConnectArgs{
			TransportAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
			TransitGatewayId:      pulumi.Any(aws_ec2_transit_gateway.Example.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewConnectPeer(ctx, "exampleConnectPeer", &ec2transitgateway.ConnectPeerArgs{
			PeerAddress: pulumi.String("10.1.2.3"),
			InsideCidrBlocks: pulumi.StringArray{
				pulumi.String("169.254.100.0/29"),
			},
			TransitGatewayAttachmentId: exampleConnect.ID(),
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
import com.pulumi.aws.ec2transitgateway.Connect;
import com.pulumi.aws.ec2transitgateway.ConnectArgs;
import com.pulumi.aws.ec2transitgateway.ConnectPeer;
import com.pulumi.aws.ec2transitgateway.ConnectPeerArgs;
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
        var exampleConnect = new Connect("exampleConnect", ConnectArgs.builder()        
            .transportAttachmentId(aws_ec2_transit_gateway_vpc_attachment.example().id())
            .transitGatewayId(aws_ec2_transit_gateway.example().id())
            .build());

        var exampleConnectPeer = new ConnectPeer("exampleConnectPeer", ConnectPeerArgs.builder()        
            .peerAddress("10.1.2.3")
            .insideCidrBlocks("169.254.100.0/29")
            .transitGatewayAttachmentId(exampleConnect.id())
            .build());

    }
}
```
```yaml
resources:
  exampleConnect:
    type: aws:ec2transitgateway:Connect
    properties:
      transportAttachmentId: ${aws_ec2_transit_gateway_vpc_attachment.example.id}
      transitGatewayId: ${aws_ec2_transit_gateway.example.id}
  exampleConnectPeer:
    type: aws:ec2transitgateway:ConnectPeer
    properties:
      peerAddress: 10.1.2.3
      insideCidrBlocks:
        - 169.254.100.0/29
      transitGatewayAttachmentId: ${exampleConnect.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_connect_peer` using the EC2 Transit Gateway Connect Peer identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/connectPeer:ConnectPeer example tgw-connect-peer-12345678
```
 