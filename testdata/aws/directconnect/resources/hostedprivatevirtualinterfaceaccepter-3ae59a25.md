Provides a resource to manage the accepter's side of a Direct Connect hosted private virtual interface.
This resource accepts ownership of a private virtual interface created by another AWS account.

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
const vpnGw = new aws.ec2.VpnGateway("vpnGw", {}, {
    provider: aws.accepter,
});
// Creator's side of the VIF
const creator = new aws.directconnect.HostedPrivateVirtualInterface("creator", {
    connectionId: "dxcon-zzzzzzzz",
    ownerAccountId: accepterCallerIdentity.then(accepterCallerIdentity => accepterCallerIdentity.accountId),
    vlan: 4094,
    addressFamily: "ipv4",
    bgpAsn: 65352,
}, {
    dependsOn: [vpnGw],
});
const accepterHostedPrivateVirtualInterfaceAccepter = new aws.directconnect.HostedPrivateVirtualInterfaceAccepter("accepterHostedPrivateVirtualInterfaceAccepter", {
    virtualInterfaceId: creator.id,
    vpnGatewayId: vpnGw.id,
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
vpn_gw = aws.ec2.VpnGateway("vpnGw", opts=pulumi.ResourceOptions(provider=aws["accepter"]))
# Creator's side of the VIF
creator = aws.directconnect.HostedPrivateVirtualInterface("creator",
    connection_id="dxcon-zzzzzzzz",
    owner_account_id=accepter_caller_identity.account_id,
    vlan=4094,
    address_family="ipv4",
    bgp_asn=65352,
    opts=pulumi.ResourceOptions(depends_on=[vpn_gw]))
accepter_hosted_private_virtual_interface_accepter = aws.directconnect.HostedPrivateVirtualInterfaceAccepter("accepterHostedPrivateVirtualInterfaceAccepter",
    virtual_interface_id=creator.id,
    vpn_gateway_id=vpn_gw.id,
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
    var vpnGw = new Aws.Ec2.VpnGateway("vpnGw", new()
    {
    }, new CustomResourceOptions
    {
        Provider = aws.Accepter,
    });

    // Creator's side of the VIF
    var creator = new Aws.DirectConnect.HostedPrivateVirtualInterface("creator", new()
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
            vpnGw,
        },
    });

    var accepterHostedPrivateVirtualInterfaceAccepter = new Aws.DirectConnect.HostedPrivateVirtualInterfaceAccepter("accepterHostedPrivateVirtualInterfaceAccepter", new()
    {
        VirtualInterfaceId = creator.Id,
        VpnGatewayId = vpnGw.Id,
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
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
		vpnGw, err := ec2.NewVpnGateway(ctx, "vpnGw", nil, pulumi.Provider(aws.Accepter))
		if err != nil {
			return err
		}
		creator, err := directconnect.NewHostedPrivateVirtualInterface(ctx, "creator", &directconnect.HostedPrivateVirtualInterfaceArgs{
			ConnectionId:   pulumi.String("dxcon-zzzzzzzz"),
			OwnerAccountId: *pulumi.String(accepterCallerIdentity.AccountId),
			Vlan:           pulumi.Int(4094),
			AddressFamily:  pulumi.String("ipv4"),
			BgpAsn:         pulumi.Int(65352),
		}, pulumi.DependsOn([]pulumi.Resource{
			vpnGw,
		}))
		if err != nil {
			return err
		}
		_, err = directconnect.NewHostedPrivateVirtualInterfaceAccepter(ctx, "accepterHostedPrivateVirtualInterfaceAccepter", &directconnect.HostedPrivateVirtualInterfaceAccepterArgs{
			VirtualInterfaceId: creator.ID(),
			VpnGatewayId:       vpnGw.ID(),
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
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
import com.pulumi.aws.directconnect.HostedPrivateVirtualInterface;
import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceArgs;
import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceAccepter;
import com.pulumi.aws.directconnect.HostedPrivateVirtualInterfaceAccepterArgs;
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

        var vpnGw = new VpnGateway("vpnGw", VpnGatewayArgs.Empty, CustomResourceOptions.builder()
            .provider(aws.accepter())
            .build());

        var creator = new HostedPrivateVirtualInterface("creator", HostedPrivateVirtualInterfaceArgs.builder()        
            .connectionId("dxcon-zzzzzzzz")
            .ownerAccountId(accepterCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .vlan(4094)
            .addressFamily("ipv4")
            .bgpAsn(65352)
            .build(), CustomResourceOptions.builder()
                .dependsOn(vpnGw)
                .build());

        var accepterHostedPrivateVirtualInterfaceAccepter = new HostedPrivateVirtualInterfaceAccepter("accepterHostedPrivateVirtualInterfaceAccepter", HostedPrivateVirtualInterfaceAccepterArgs.builder()        
            .virtualInterfaceId(creator.id())
            .vpnGatewayId(vpnGw.id())
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
    type: aws:directconnect:HostedPrivateVirtualInterface
    properties:
      connectionId: dxcon-zzzzzzzz
      ownerAccountId: ${accepterCallerIdentity.accountId}
      vlan: 4094
      addressFamily: ipv4
      bgpAsn: 65352
    options:
      dependson:
        - ${vpnGw}
  # Accepter's side of the VIF.
  vpnGw:
    type: aws:ec2:VpnGateway
    options:
      provider: ${aws.accepter}
  accepterHostedPrivateVirtualInterfaceAccepter:
    type: aws:directconnect:HostedPrivateVirtualInterfaceAccepter
    properties:
      virtualInterfaceId: ${creator.id}
      vpnGatewayId: ${vpnGw.id}
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

Using `pulumi import`, import Direct Connect hosted private virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter test dxvif-33cc44dd
```
 