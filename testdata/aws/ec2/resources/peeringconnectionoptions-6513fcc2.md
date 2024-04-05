Provides a resource to manage VPC peering connection options.

> **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
both a standalone VPC Peering Connection Options and a VPC Peering Connection
resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
Doing so will cause a conflict of options and will overwrite the options.
Using a VPC Peering Connection Options resource decouples management of the connection options from
management of the VPC Peering Connection and allows options to be set correctly in cross-region and
cross-account scenarios.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.0.0.0/16"});
const bar = new aws.ec2.Vpc("bar", {cidrBlock: "10.1.0.0/16"});
const fooVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection", {
    vpcId: fooVpc.id,
    peerVpcId: bar.id,
    autoAccept: true,
});
const fooPeeringConnectionOptions = new aws.ec2.PeeringConnectionOptions("fooPeeringConnectionOptions", {
    vpcPeeringConnectionId: fooVpcPeeringConnection.id,
    accepter: {
        allowRemoteVpcDnsResolution: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.0.0.0/16")
bar = aws.ec2.Vpc("bar", cidr_block="10.1.0.0/16")
foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
    vpc_id=foo_vpc.id,
    peer_vpc_id=bar.id,
    auto_accept=True)
foo_peering_connection_options = aws.ec2.PeeringConnectionOptions("fooPeeringConnectionOptions",
    vpc_peering_connection_id=foo_vpc_peering_connection.id,
    accepter=aws.ec2.PeeringConnectionOptionsAccepterArgs(
        allow_remote_vpc_dns_resolution=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooVpc = new Aws.Ec2.Vpc("fooVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var bar = new Aws.Ec2.Vpc("bar", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var fooVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("fooVpcPeeringConnection", new()
    {
        VpcId = fooVpc.Id,
        PeerVpcId = bar.Id,
        AutoAccept = true,
    });

    var fooPeeringConnectionOptions = new Aws.Ec2.PeeringConnectionOptions("fooPeeringConnectionOptions", new()
    {
        VpcPeeringConnectionId = fooVpcPeeringConnection.Id,
        Accepter = new Aws.Ec2.Inputs.PeeringConnectionOptionsAccepterArgs
        {
            AllowRemoteVpcDnsResolution = true,
        },
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
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewVpc(ctx, "bar", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		fooVpcPeeringConnection, err := ec2.NewVpcPeeringConnection(ctx, "fooVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
			VpcId:      fooVpc.ID(),
			PeerVpcId:  bar.ID(),
			AutoAccept: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewPeeringConnectionOptions(ctx, "fooPeeringConnectionOptions", &ec2.PeeringConnectionOptionsArgs{
			VpcPeeringConnectionId: fooVpcPeeringConnection.ID(),
			Accepter: &ec2.PeeringConnectionOptionsAccepterArgs{
				AllowRemoteVpcDnsResolution: pulumi.Bool(true),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.PeeringConnectionOptions;
import com.pulumi.aws.ec2.PeeringConnectionOptionsArgs;
import com.pulumi.aws.ec2.inputs.PeeringConnectionOptionsAccepterArgs;
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
        var fooVpc = new Vpc("fooVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var bar = new Vpc("bar", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        var fooVpcPeeringConnection = new VpcPeeringConnection("fooVpcPeeringConnection", VpcPeeringConnectionArgs.builder()        
            .vpcId(fooVpc.id())
            .peerVpcId(bar.id())
            .autoAccept(true)
            .build());

        var fooPeeringConnectionOptions = new PeeringConnectionOptions("fooPeeringConnectionOptions", PeeringConnectionOptionsArgs.builder()        
            .vpcPeeringConnectionId(fooVpcPeeringConnection.id())
            .accepter(PeeringConnectionOptionsAccepterArgs.builder()
                .allowRemoteVpcDnsResolution(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  fooVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  bar:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  fooVpcPeeringConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      vpcId: ${fooVpc.id}
      peerVpcId: ${bar.id}
      autoAccept: true
  fooPeeringConnectionOptions:
    type: aws:ec2:PeeringConnectionOptions
    properties:
      vpcPeeringConnectionId: ${fooVpcPeeringConnection.id}
      accepter:
        allowRemoteVpcDnsResolution: true
```
{{% /example %}}
{{% example %}}
### Cross-Account Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const requester = new aws.Provider("requester", {});
// Requester's credentials.
const accepter = new aws.Provider("accepter", {});
// Accepter's credentials.
const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
    enableDnsSupport: true,
    enableDnsHostnames: true,
}, {
    provider: aws.requester,
});
const peerVpc = new aws.ec2.Vpc("peerVpc", {
    cidrBlock: "10.1.0.0/16",
    enableDnsSupport: true,
    enableDnsHostnames: true,
}, {
    provider: aws.accepter,
});
const peerCallerIdentity = aws.getCallerIdentity({});
// Requester's side of the connection.
const peerVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection", {
    vpcId: main.id,
    peerVpcId: peerVpc.id,
    peerOwnerId: peerCallerIdentity.then(peerCallerIdentity => peerCallerIdentity.accountId),
    autoAccept: false,
    tags: {
        Side: "Requester",
    },
}, {
    provider: aws.requester,
});
// Accepter's side of the connection.
const peerVpcPeeringConnectionAccepter = new aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", {
    vpcPeeringConnectionId: peerVpcPeeringConnection.id,
    autoAccept: true,
    tags: {
        Side: "Accepter",
    },
}, {
    provider: aws.accepter,
});
const requesterPeeringConnectionOptions = new aws.ec2.PeeringConnectionOptions("requesterPeeringConnectionOptions", {
    vpcPeeringConnectionId: peerVpcPeeringConnectionAccepter.id,
    requester: {
        allowRemoteVpcDnsResolution: true,
    },
}, {
    provider: aws.requester,
});
const accepterPeeringConnectionOptions = new aws.ec2.PeeringConnectionOptions("accepterPeeringConnectionOptions", {
    vpcPeeringConnectionId: peerVpcPeeringConnectionAccepter.id,
    accepter: {
        allowRemoteVpcDnsResolution: true,
    },
}, {
    provider: aws.accepter,
});
```
```python
import pulumi
import pulumi_aws as aws

requester = aws.Provider("requester")
# Requester's credentials.
accepter = aws.Provider("accepter")
# Accepter's credentials.
main = aws.ec2.Vpc("main",
    cidr_block="10.0.0.0/16",
    enable_dns_support=True,
    enable_dns_hostnames=True,
    opts=pulumi.ResourceOptions(provider=aws["requester"]))
peer_vpc = aws.ec2.Vpc("peerVpc",
    cidr_block="10.1.0.0/16",
    enable_dns_support=True,
    enable_dns_hostnames=True,
    opts=pulumi.ResourceOptions(provider=aws["accepter"]))
peer_caller_identity = aws.get_caller_identity()
# Requester's side of the connection.
peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
    vpc_id=main.id,
    peer_vpc_id=peer_vpc.id,
    peer_owner_id=peer_caller_identity.account_id,
    auto_accept=False,
    tags={
        "Side": "Requester",
    },
    opts=pulumi.ResourceOptions(provider=aws["requester"]))
# Accepter's side of the connection.
peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
    vpc_peering_connection_id=peer_vpc_peering_connection.id,
    auto_accept=True,
    tags={
        "Side": "Accepter",
    },
    opts=pulumi.ResourceOptions(provider=aws["accepter"]))
requester_peering_connection_options = aws.ec2.PeeringConnectionOptions("requesterPeeringConnectionOptions",
    vpc_peering_connection_id=peer_vpc_peering_connection_accepter.id,
    requester=aws.ec2.PeeringConnectionOptionsRequesterArgs(
        allow_remote_vpc_dns_resolution=True,
    ),
    opts=pulumi.ResourceOptions(provider=aws["requester"]))
accepter_peering_connection_options = aws.ec2.PeeringConnectionOptions("accepterPeeringConnectionOptions",
    vpc_peering_connection_id=peer_vpc_peering_connection_accepter.id,
    accepter=aws.ec2.PeeringConnectionOptionsAccepterArgs(
        allow_remote_vpc_dns_resolution=True,
    ),
    opts=pulumi.ResourceOptions(provider=aws["accepter"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var requester = new Aws.Provider("requester");

    // Requester's credentials.
    var accepter = new Aws.Provider("accepter");

    // Accepter's credentials.
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
        EnableDnsSupport = true,
        EnableDnsHostnames = true,
    }, new CustomResourceOptions
    {
        Provider = aws.Requester,
    });

    var peerVpc = new Aws.Ec2.Vpc("peerVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
        EnableDnsSupport = true,
        EnableDnsHostnames = true,
    }, new CustomResourceOptions
    {
        Provider = aws.Accepter,
    });

    var peerCallerIdentity = Aws.GetCallerIdentity.Invoke();

    // Requester's side of the connection.
    var peerVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("peerVpcPeeringConnection", new()
    {
        VpcId = main.Id,
        PeerVpcId = peerVpc.Id,
        PeerOwnerId = peerCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        AutoAccept = false,
        Tags = 
        {
            { "Side", "Requester" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Requester,
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
        Provider = aws.Accepter,
    });

    var requesterPeeringConnectionOptions = new Aws.Ec2.PeeringConnectionOptions("requesterPeeringConnectionOptions", new()
    {
        VpcPeeringConnectionId = peerVpcPeeringConnectionAccepter.Id,
        Requester = new Aws.Ec2.Inputs.PeeringConnectionOptionsRequesterArgs
        {
            AllowRemoteVpcDnsResolution = true,
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Requester,
    });

    var accepterPeeringConnectionOptions = new Aws.Ec2.PeeringConnectionOptions("accepterPeeringConnectionOptions", new()
    {
        VpcPeeringConnectionId = peerVpcPeeringConnectionAccepter.Id,
        Accepter = new Aws.Ec2.Inputs.PeeringConnectionOptionsAccepterArgs
        {
            AllowRemoteVpcDnsResolution = true,
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Accepter,
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
		_, err := aws.NewProvider(ctx, "requester", nil)
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "accepter", nil)
		if err != nil {
			return err
		}
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.0.0.0/16"),
			EnableDnsSupport:   pulumi.Bool(true),
			EnableDnsHostnames: pulumi.Bool(true),
		}, pulumi.Provider(aws.Requester))
		if err != nil {
			return err
		}
		peerVpc, err := ec2.NewVpc(ctx, "peerVpc", &ec2.VpcArgs{
			CidrBlock:          pulumi.String("10.1.0.0/16"),
			EnableDnsSupport:   pulumi.Bool(true),
			EnableDnsHostnames: pulumi.Bool(true),
		}, pulumi.Provider(aws.Accepter))
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
			AutoAccept:  pulumi.Bool(false),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Requester"),
			},
		}, pulumi.Provider(aws.Requester))
		if err != nil {
			return err
		}
		peerVpcPeeringConnectionAccepter, err := ec2.NewVpcPeeringConnectionAccepter(ctx, "peerVpcPeeringConnectionAccepter", &ec2.VpcPeeringConnectionAccepterArgs{
			VpcPeeringConnectionId: peerVpcPeeringConnection.ID(),
			AutoAccept:             pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Accepter"),
			},
		}, pulumi.Provider(aws.Accepter))
		if err != nil {
			return err
		}
		_, err = ec2.NewPeeringConnectionOptions(ctx, "requesterPeeringConnectionOptions", &ec2.PeeringConnectionOptionsArgs{
			VpcPeeringConnectionId: peerVpcPeeringConnectionAccepter.ID(),
			Requester: &ec2.PeeringConnectionOptionsRequesterArgs{
				AllowRemoteVpcDnsResolution: pulumi.Bool(true),
			},
		}, pulumi.Provider(aws.Requester))
		if err != nil {
			return err
		}
		_, err = ec2.NewPeeringConnectionOptions(ctx, "accepterPeeringConnectionOptions", &ec2.PeeringConnectionOptionsArgs{
			VpcPeeringConnectionId: peerVpcPeeringConnectionAccepter.ID(),
			Accepter: &ec2.PeeringConnectionOptionsAccepterArgs{
				AllowRemoteVpcDnsResolution: pulumi.Bool(true),
			},
		}, pulumi.Provider(aws.Accepter))
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepter;
import com.pulumi.aws.ec2.VpcPeeringConnectionAccepterArgs;
import com.pulumi.aws.ec2.PeeringConnectionOptions;
import com.pulumi.aws.ec2.PeeringConnectionOptionsArgs;
import com.pulumi.aws.ec2.inputs.PeeringConnectionOptionsRequesterArgs;
import com.pulumi.aws.ec2.inputs.PeeringConnectionOptionsAccepterArgs;
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
        var requester = new Provider("requester");

        var accepter = new Provider("accepter");

        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .enableDnsSupport(true)
            .enableDnsHostnames(true)
            .build(), CustomResourceOptions.builder()
                .provider(aws.requester())
                .build());

        var peerVpc = new Vpc("peerVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .enableDnsSupport(true)
            .enableDnsHostnames(true)
            .build(), CustomResourceOptions.builder()
                .provider(aws.accepter())
                .build());

        final var peerCallerIdentity = AwsFunctions.getCallerIdentity();

        var peerVpcPeeringConnection = new VpcPeeringConnection("peerVpcPeeringConnection", VpcPeeringConnectionArgs.builder()        
            .vpcId(main.id())
            .peerVpcId(peerVpc.id())
            .peerOwnerId(peerCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .autoAccept(false)
            .tags(Map.of("Side", "Requester"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.requester())
                .build());

        var peerVpcPeeringConnectionAccepter = new VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter", VpcPeeringConnectionAccepterArgs.builder()        
            .vpcPeeringConnectionId(peerVpcPeeringConnection.id())
            .autoAccept(true)
            .tags(Map.of("Side", "Accepter"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.accepter())
                .build());

        var requesterPeeringConnectionOptions = new PeeringConnectionOptions("requesterPeeringConnectionOptions", PeeringConnectionOptionsArgs.builder()        
            .vpcPeeringConnectionId(peerVpcPeeringConnectionAccepter.id())
            .requester(PeeringConnectionOptionsRequesterArgs.builder()
                .allowRemoteVpcDnsResolution(true)
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(aws.requester())
                .build());

        var accepterPeeringConnectionOptions = new PeeringConnectionOptions("accepterPeeringConnectionOptions", PeeringConnectionOptionsArgs.builder()        
            .vpcPeeringConnectionId(peerVpcPeeringConnectionAccepter.id())
            .accepter(PeeringConnectionOptionsAccepterArgs.builder()
                .allowRemoteVpcDnsResolution(true)
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(aws.accepter())
                .build());

    }
}
```
```yaml
resources:
  requester:
    type: pulumi:providers:aws
  accepter:
    type: pulumi:providers:aws
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
      enableDnsSupport: true
      enableDnsHostnames: true
    options:
      provider: ${aws.requester}
  peerVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
      enableDnsSupport: true
      enableDnsHostnames: true
    options:
      provider: ${aws.accepter}
  # Requester's side of the connection.
  peerVpcPeeringConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      vpcId: ${main.id}
      peerVpcId: ${peerVpc.id}
      peerOwnerId: ${peerCallerIdentity.accountId}
      autoAccept: false
      tags:
        Side: Requester
    options:
      provider: ${aws.requester}
  # Accepter's side of the connection.
  peerVpcPeeringConnectionAccepter:
    type: aws:ec2:VpcPeeringConnectionAccepter
    properties:
      vpcPeeringConnectionId: ${peerVpcPeeringConnection.id}
      autoAccept: true
      tags:
        Side: Accepter
    options:
      provider: ${aws.accepter}
  requesterPeeringConnectionOptions:
    type: aws:ec2:PeeringConnectionOptions
    properties:
      # As options can't be set until the connection has been accepted
      #   # create an explicit dependency on the accepter.
      vpcPeeringConnectionId: ${peerVpcPeeringConnectionAccepter.id}
      requester:
        allowRemoteVpcDnsResolution: true
    options:
      provider: ${aws.requester}
  accepterPeeringConnectionOptions:
    type: aws:ec2:PeeringConnectionOptions
    properties:
      vpcPeeringConnectionId: ${peerVpcPeeringConnectionAccepter.id}
      accepter:
        allowRemoteVpcDnsResolution: true
    options:
      provider: ${aws.accepter}
variables:
  peerCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Peering Connection Options using the VPC peering `id`. For example:

```sh
 $ pulumi import aws:ec2/peeringConnectionOptions:PeeringConnectionOptions foo pcx-111aaa111
```
 