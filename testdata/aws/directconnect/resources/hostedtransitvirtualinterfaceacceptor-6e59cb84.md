Provides a resource to manage the accepter's side of a Direct Connect hosted transit virtual interface.
This resource accepts ownership of a transit virtual interface created by another AWS account.

> **NOTE:** AWS allows a Direct Connect hosted transit virtual interface to be deleted from either the allocator's or accepter's side. However, this provider only allows the Direct Connect hosted transit virtual interface to be deleted from the allocator's side by removing the corresponding `aws.directconnect.HostedTransitVirtualInterface` resource from your configuration. Removing a `aws.directconnect.HostedTransitVirtualInterfaceAcceptor` resource from your configuration will remove it from your statefile and management, **but will not delete the Direct Connect virtual interface.**

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const accepter = new aws.Provider("accepter", {});
// Accepter's credentials.
const accepterCallerIdentity = aws.getCallerIdentity({});
// Accepter's side of the VIF.
const example = new aws.directconnect.Gateway("example", {amazonSideAsn: "64512"}, {
    provider: aws.accepter,
});
// Creator's side of the VIF
const creator = new aws.directconnect.HostedTransitVirtualInterface("creator", {
    connectionId: "dxcon-zzzzzzzz",
    ownerAccountId: accepterCallerIdentity.then(accepterCallerIdentity => accepterCallerIdentity.accountId),
    vlan: 4094,
    addressFamily: "ipv4",
    bgpAsn: 65352,
}, {
    dependsOn: [example],
});
const accepterHostedTransitVirtualInterfaceAcceptor = new aws.directconnect.HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor", {
    virtualInterfaceId: creator.id,
    dxGatewayId: example.id,
    tags: {
        Side: "Accepter",
    },
}, {
    provider: aws.accepter,
});
```
```python
import pulumi
import pulumi_aws as aws

accepter = aws.Provider("accepter")
# Accepter's credentials.
accepter_caller_identity = aws.get_caller_identity()
# Accepter's side of the VIF.
example = aws.directconnect.Gateway("example", amazon_side_asn="64512",
opts=pulumi.ResourceOptions(provider=aws["accepter"]))
# Creator's side of the VIF
creator = aws.directconnect.HostedTransitVirtualInterface("creator",
    connection_id="dxcon-zzzzzzzz",
    owner_account_id=accepter_caller_identity.account_id,
    vlan=4094,
    address_family="ipv4",
    bgp_asn=65352,
    opts=pulumi.ResourceOptions(depends_on=[example]))
accepter_hosted_transit_virtual_interface_acceptor = aws.directconnect.HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor",
    virtual_interface_id=creator.id,
    dx_gateway_id=example.id,
    tags={
        "Side": "Accepter",
    },
    opts=pulumi.ResourceOptions(provider=aws["accepter"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var accepter = new Aws.Provider("accepter");

    // Accepter's credentials.
    var accepterCallerIdentity = Aws.GetCallerIdentity.Invoke();

    // Accepter's side of the VIF.
    var example = new Aws.DirectConnect.Gateway("example", new()
    {
        AmazonSideAsn = "64512",
    }, new CustomResourceOptions
    {
        Provider = aws.Accepter,
    });

    // Creator's side of the VIF
    var creator = new Aws.DirectConnect.HostedTransitVirtualInterface("creator", new()
    {
        ConnectionId = "dxcon-zzzzzzzz",
        OwnerAccountId = accepterCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        Vlan = 4094,
        AddressFamily = "ipv4",
        BgpAsn = 65352,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            example,
        },
    });

    var accepterHostedTransitVirtualInterfaceAcceptor = new Aws.DirectConnect.HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor", new()
    {
        VirtualInterfaceId = creator.Id,
        DxGatewayId = example.Id,
        Tags = 
        {
            { "Side", "Accepter" },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "accepter", nil)
		if err != nil {
			return err
		}
		accepterCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		example, err := directconnect.NewGateway(ctx, "example", &directconnect.GatewayArgs{
			AmazonSideAsn: pulumi.String("64512"),
		}, pulumi.Provider(aws.Accepter))
		if err != nil {
			return err
		}
		creator, err := directconnect.NewHostedTransitVirtualInterface(ctx, "creator", &directconnect.HostedTransitVirtualInterfaceArgs{
			ConnectionId:   pulumi.String("dxcon-zzzzzzzz"),
			OwnerAccountId: *pulumi.String(accepterCallerIdentity.AccountId),
			Vlan:           pulumi.Int(4094),
			AddressFamily:  pulumi.String("ipv4"),
			BgpAsn:         pulumi.Int(65352),
		}, pulumi.DependsOn([]pulumi.Resource{
			example,
		}))
		if err != nil {
			return err
		}
		_, err = directconnect.NewHostedTransitVirtualInterfaceAcceptor(ctx, "accepterHostedTransitVirtualInterfaceAcceptor", &directconnect.HostedTransitVirtualInterfaceAcceptorArgs{
			VirtualInterfaceId: creator.ID(),
			DxGatewayId:        example.ID(),
			Tags: pulumi.StringMap{
				"Side": pulumi.String("Accepter"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.directconnect.Gateway;
import com.pulumi.aws.directconnect.GatewayArgs;
import com.pulumi.aws.directconnect.HostedTransitVirtualInterface;
import com.pulumi.aws.directconnect.HostedTransitVirtualInterfaceArgs;
import com.pulumi.aws.directconnect.HostedTransitVirtualInterfaceAcceptor;
import com.pulumi.aws.directconnect.HostedTransitVirtualInterfaceAcceptorArgs;
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
        var accepter = new Provider("accepter");

        final var accepterCallerIdentity = AwsFunctions.getCallerIdentity();

        var example = new Gateway("example", GatewayArgs.builder()        
            .amazonSideAsn(64512)
            .build(), CustomResourceOptions.builder()
                .provider(aws.accepter())
                .build());

        var creator = new HostedTransitVirtualInterface("creator", HostedTransitVirtualInterfaceArgs.builder()        
            .connectionId("dxcon-zzzzzzzz")
            .ownerAccountId(accepterCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .vlan(4094)
            .addressFamily("ipv4")
            .bgpAsn(65352)
            .build(), CustomResourceOptions.builder()
                .dependsOn(example)
                .build());

        var accepterHostedTransitVirtualInterfaceAcceptor = new HostedTransitVirtualInterfaceAcceptor("accepterHostedTransitVirtualInterfaceAcceptor", HostedTransitVirtualInterfaceAcceptorArgs.builder()        
            .virtualInterfaceId(creator.id())
            .dxGatewayId(example.id())
            .tags(Map.of("Side", "Accepter"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.accepter())
                .build());

    }
}
```
```yaml
resources:
  accepter:
    type: pulumi:providers:aws
  # Creator's side of the VIF
  creator:
    type: aws:directconnect:HostedTransitVirtualInterface
    properties:
      connectionId: dxcon-zzzzzzzz
      ownerAccountId: ${accepterCallerIdentity.accountId}
      vlan: 4094
      addressFamily: ipv4
      bgpAsn: 65352
    options:
      dependson:
        - ${example}
  # Accepter's side of the VIF.
  example:
    type: aws:directconnect:Gateway
    properties:
      amazonSideAsn: 64512
    options:
      provider: ${aws.accepter}
  accepterHostedTransitVirtualInterfaceAcceptor:
    type: aws:directconnect:HostedTransitVirtualInterfaceAcceptor
    properties:
      virtualInterfaceId: ${creator.id}
      dxGatewayId: ${example.id}
      tags:
        Side: Accepter
    options:
      provider: ${aws.accepter}
variables:
  accepterCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect hosted transit virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor test dxvif-33cc44dd
```
 