Manages an EC2 Transit Gateway Peering Attachment.
For examples of custom route table association and propagation, see the [EC2 Transit Gateway Networking Examples Guide](https://docs.aws.amazon.com/vpc/latest/tgw/TGW_Scenarios.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const local = new aws.Provider("local", {region: "us-east-1"});
const peer = new aws.Provider("peer", {region: "us-west-2"});
const peerRegion = aws.getRegion({});
const localTransitGateway = new aws.ec2transitgateway.TransitGateway("localTransitGateway", {tags: {
    Name: "Local TGW",
}}, {
    provider: aws.local,
});
const peerTransitGateway = new aws.ec2transitgateway.TransitGateway("peerTransitGateway", {tags: {
    Name: "Peer TGW",
}}, {
    provider: aws.peer,
});
const example = new aws.ec2transitgateway.PeeringAttachment("example", {
    peerAccountId: peerTransitGateway.ownerId,
    peerRegion: peerRegion.then(peerRegion => peerRegion.name),
    peerTransitGatewayId: peerTransitGateway.id,
    transitGatewayId: localTransitGateway.id,
    tags: {
        Name: "TGW Peering Requestor",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

local = aws.Provider("local", region="us-east-1")
peer = aws.Provider("peer", region="us-west-2")
peer_region = aws.get_region()
local_transit_gateway = aws.ec2transitgateway.TransitGateway("localTransitGateway", tags={
    "Name": "Local TGW",
},
opts=pulumi.ResourceOptions(provider=aws["local"]))
peer_transit_gateway = aws.ec2transitgateway.TransitGateway("peerTransitGateway", tags={
    "Name": "Peer TGW",
},
opts=pulumi.ResourceOptions(provider=aws["peer"]))
example = aws.ec2transitgateway.PeeringAttachment("example",
    peer_account_id=peer_transit_gateway.owner_id,
    peer_region=peer_region.name,
    peer_transit_gateway_id=peer_transit_gateway.id,
    transit_gateway_id=local_transit_gateway.id,
    tags={
        "Name": "TGW Peering Requestor",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var local = new Aws.Provider("local", new()
    {
        Region = "us-east-1",
    });

    var peer = new Aws.Provider("peer", new()
    {
        Region = "us-west-2",
    });

    var peerRegion = Aws.GetRegion.Invoke();

    var localTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("localTransitGateway", new()
    {
        Tags = 
        {
            { "Name", "Local TGW" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Local,
    });

    var peerTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("peerTransitGateway", new()
    {
        Tags = 
        {
            { "Name", "Peer TGW" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Peer,
    });

    var example = new Aws.Ec2TransitGateway.PeeringAttachment("example", new()
    {
        PeerAccountId = peerTransitGateway.OwnerId,
        PeerRegion = peerRegion.Apply(getRegionResult => getRegionResult.Name),
        PeerTransitGatewayId = peerTransitGateway.Id,
        TransitGatewayId = localTransitGateway.Id,
        Tags = 
        {
            { "Name", "TGW Peering Requestor" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "local", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "peer", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		peerRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		localTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "localTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Local TGW"),
			},
		}, pulumi.Provider(aws.Local))
		if err != nil {
			return err
		}
		peerTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "peerTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Peer TGW"),
			},
		}, pulumi.Provider(aws.Peer))
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewPeeringAttachment(ctx, "example", &ec2transitgateway.PeeringAttachmentArgs{
			PeerAccountId:        peerTransitGateway.OwnerId,
			PeerRegion:           *pulumi.String(peerRegion.Name),
			PeerTransitGatewayId: peerTransitGateway.ID(),
			TransitGatewayId:     localTransitGateway.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("TGW Peering Requestor"),
			},
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
import com.pulumi.aws.ec2transitgateway.PeeringAttachment;
import com.pulumi.aws.ec2transitgateway.PeeringAttachmentArgs;
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
        var local = new Provider("local", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var peer = new Provider("peer", ProviderArgs.builder()        
            .region("us-west-2")
            .build());

        final var peerRegion = AwsFunctions.getRegion();

        var localTransitGateway = new TransitGateway("localTransitGateway", TransitGatewayArgs.builder()        
            .tags(Map.of("Name", "Local TGW"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.local())
                .build());

        var peerTransitGateway = new TransitGateway("peerTransitGateway", TransitGatewayArgs.builder()        
            .tags(Map.of("Name", "Peer TGW"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.peer())
                .build());

        var example = new PeeringAttachment("example", PeeringAttachmentArgs.builder()        
            .peerAccountId(peerTransitGateway.ownerId())
            .peerRegion(peerRegion.applyValue(getRegionResult -> getRegionResult.name()))
            .peerTransitGatewayId(peerTransitGateway.id())
            .transitGatewayId(localTransitGateway.id())
            .tags(Map.of("Name", "TGW Peering Requestor"))
            .build());

    }
}
```
```yaml
resources:
  local:
    type: pulumi:providers:aws
    properties:
      region: us-east-1
  peer:
    type: pulumi:providers:aws
    properties:
      region: us-west-2
  localTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
    properties:
      tags:
        Name: Local TGW
    options:
      provider: ${aws.local}
  peerTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
    properties:
      tags:
        Name: Peer TGW
    options:
      provider: ${aws.peer}
  example:
    type: aws:ec2transitgateway:PeeringAttachment
    properties:
      peerAccountId: ${peerTransitGateway.ownerId}
      peerRegion: ${peerRegion.name}
      peerTransitGatewayId: ${peerTransitGateway.id}
      transitGatewayId: ${localTransitGateway.id}
      tags:
        Name: TGW Peering Requestor
variables:
  peerRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_peering_attachment` using the EC2 Transit Gateway Attachment identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/peeringAttachment:PeeringAttachment example tgw-attach-12345678
```
 