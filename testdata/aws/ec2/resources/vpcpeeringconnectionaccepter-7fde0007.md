Provides a resource to manage the accepter's side of a VPC Peering Connection.

When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
accepter's account.
The requester can use the `aws.ec2.VpcPeeringConnection` resource to manage its side of the connection
and the accepter can use the `aws.ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
connection into management.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const peer = new aws.Provider("peer", {region: "us-west-2"});
// Accepter's credentials.
const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const peerVpc = new aws.ec2.Vpc("peerVpc", {cidrBlock: "10.1.0.0/16"}, {
    provider: aws.peer,
});
const peerCallerIdentity = aws.getCallerIdentity({});
// Requester's side of the connection.
const peerVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection", {
    vpcId: main.id,
    peerVpcId: peerVpc.id,
    peerOwnerId: peerCallerIdentity.then(peerCallerIdentity => peerCallerIdentity.accountId),
    peerRegion: "us-west-2",
    autoAccept: false,
    tags: {
        Side: "Requester",
    },
});
// Accepter's side of the connection.
const peerVpcPeeringConnectionAccepter = new aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", {
    vpcPeeringConnectionId: peerVpcPeeringConnection.id,
    autoAccept: true,
    tags: {
        Side: "Accepter",
    },
}, {
    provider: aws.peer,
});
```
```python
import pulumi
import pulumi_aws as aws

peer = aws.Provider("peer", region="us-west-2")
# Accepter's credentials.
main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
peer_vpc = aws.ec2.Vpc("peerVpc", cidr_block="10.1.0.0/16",
opts=pulumi.ResourceOptions(provider=aws["peer"]))
peer_caller_identity = aws.get_caller_identity()
# Requester's side of the connection.
peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
    vpc_id=main.id,
    peer_vpc_id=peer_vpc.id,
    peer_owner_id=peer_caller_identity.account_id,
    peer_region="us-west-2",
    auto_accept=False,
    tags={
        "Side": "Requester",
    })
# Accepter's side of the connection.
peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
    vpc_peering_connection_id=peer_vpc_peering_connection.id,
    auto_accept=True,
    tags={
        "Side": "Accepter",
    },
    opts=pulumi.ResourceOptions(provider=aws["peer"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var peer = new Aws.Provider("peer", new()
    {
        Region = "us-west-2",
    });

    // Accepter's credentials.
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var peerVpc = new Aws.Ec2.Vpc("peerVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = aws.Peer,
    });

    var peerCallerIdentity = Aws.GetCallerIdentity.Invoke();

    // Requester's side of the connection.
    var peerVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("peerVpcPeeringConnection", new()
    {
        VpcId = main.Id,
        PeerVpcId = peerVpc.Id,
        PeerOwnerId = peerCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        PeerRegion = "us-west-2",
        AutoAccept = false,
        Tags = 
        {
            { "Side", "Requester" },
        },
    });

    // Accepter's side of the connection.
    var peerVpcPeeringConnectionAccepter = new Aws.Ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", new()
    {
        VpcPeeringConnectionId = peerVpcPeeringConnection.Id,
        AutoAccept = true,
        Tags = 
        {
            { "Side", "Accepter" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Peer,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "peer", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		peerVpc, err := ec2.NewVpc(ctx, "peerVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		}, pulumi.Provider(aws.Peer))
		if err != nil {
			return err
		}
		peerCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		peerVpcPeeringConnection, err := ec2.NewVpcPeeringConnection(ctx, "peerVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
			VpcId:       main.ID(),
			PeerVpcId:   peerVpc.ID(),
			PeerOwnerId: *pulumi.String(peerCallerIdentity.AccountId),
			PeerRegion:  pulumi.String("us-west-2"),
			AutoAccept:  pulumi.Bool(false),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Requester"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcPeeringConnectionAccepter(ctx, "peerVpcPeeringConnectionAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
			VpcPeeringConnectionId: peerVpcPeeringConnection.ID(),
			AutoAccept:             pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Accepter"),
			},
		}, pulumi.Provider(aws.Peer))
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepter;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepterArgs;
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
        var peer = new Provider("peer", ProviderArgs.builder()        
            .region("us-west-2")
            .build());

        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var peerVpc = new Vpc("peerVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(aws.peer())
                .build());

        final var peerCallerIdentity = AwsFunctions.getCallerIdentity();

        var peerVpcPeeringConnection = new VpcPeeringConnection("peerVpcPeeringConnection", VpcPeeringConnectionArgs.builder()        
            .vpcId(main.id())
            .peerVpcId(peerVpc.id())
            .peerOwnerId(peerCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .peerRegion("us-west-2")
            .autoAccept(false)
            .tags(Map.of("Side", "Requester"))
            .build());

        var peerVpcPeeringConnectionAccepter = new VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", VpcPeeringConnectionAccepterArgs.builder()        
            .vpcPeeringConnectionId(peerVpcPeeringConnection.id())
            .autoAccept(true)
            .tags(Map.of("Side", "Accepter"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.peer())
                .build());

    }
}
```
```yaml
resources:
  peer:
    type: pulumi:providers:aws
    properties:
      region: us-west-2
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  peerVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
    options:
      provider: ${aws.peer}
  # Requester's side of the connection.
  peerVpcPeeringConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      vpcId: ${main.id}
      peerVpcId: ${peerVpc.id}
      peerOwnerId: ${peerCallerIdentity.accountId}
      peerRegion: us-west-2
      autoAccept: false
      tags:
        Side: Requester
  # Accepter's side of the connection.
  peerVpcPeeringConnectionAccepter:
    type: aws:ec2:VpcPeeringConnectionAccepter
    properties:
      vpcPeeringConnectionId: ${peerVpcPeeringConnection.id}
      autoAccept: true
      tags:
        Side: Accepter
    options:
      provider: ${aws.peer}
variables:
  peerCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Peering Connection Accepters using the Peering Connection ID. For example:

```sh
 $ pulumi import aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter example pcx-12345678
```
 Certain resource arguments, like `auto_accept`, do not have an EC2 API method for reading the information after peering connection creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

