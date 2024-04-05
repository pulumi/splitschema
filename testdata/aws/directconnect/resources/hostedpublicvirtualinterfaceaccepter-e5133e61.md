Provides a resource to manage the accepter's side of a Direct Connect hosted public virtual interface.
This resource accepts ownership of a public virtual interface created by another AWS account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const accepter = new aws.Provider("accepter", {});
// Accepter's credentials.
const accepterCallerIdentity = aws.getCallerIdentity({});
// Creator's side of the VIF
const creator = new aws.directconnect.HostedPublicVirtualInterface("creator", {
    connectionId: "dxcon-zzzzzzzz",
    ownerAccountId: accepterCallerIdentity.then(accepterCallerIdentity => accepterCallerIdentity.accountId),
    vlan: 4094,
    addressFamily: "ipv4",
    bgpAsn: 65352,
    customerAddress: "175.45.176.1/30",
    amazonAddress: "175.45.176.2/30",
    routeFilterPrefixes: [
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
});
// Accepter's side of the VIF.
const accepterHostedPublicVirtualInterfaceAccepter = new aws.directconnect.HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter", {
    virtualInterfaceId: creator.id,
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
# Creator's side of the VIF
creator = aws.directconnect.HostedPublicVirtualInterface("creator",
    connection_id="dxcon-zzzzzzzz",
    owner_account_id=accepter_caller_identity.account_id,
    vlan=4094,
    address_family="ipv4",
    bgp_asn=65352,
    customer_address="175.45.176.1/30",
    amazon_address="175.45.176.2/30",
    route_filter_prefixes=[
        "210.52.109.0/24",
        "175.45.176.0/22",
    ])
# Accepter's side of the VIF.
accepter_hosted_public_virtual_interface_accepter = aws.directconnect.HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter",
    virtual_interface_id=creator.id,
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

    // Creator's side of the VIF
    var creator = new Aws.DirectConnect.HostedPublicVirtualInterface("creator", new()
    {
        ConnectionId = "dxcon-zzzzzzzz",
        OwnerAccountId = accepterCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        Vlan = 4094,
        AddressFamily = "ipv4",
        BgpAsn = 65352,
        CustomerAddress = "175.45.176.1/30",
        AmazonAddress = "175.45.176.2/30",
        RouteFilterPrefixes = new[]
        {
            "210.52.109.0/24",
            "175.45.176.0/22",
        },
    });

    // Accepter's side of the VIF.
    var accepterHostedPublicVirtualInterfaceAccepter = new Aws.DirectConnect.HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter", new()
    {
        VirtualInterfaceId = creator.Id,
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
		creator, err := directconnect.NewHostedPublicVirtualInterface(ctx, "creator", &directconnect.HostedPublicVirtualInterfaceArgs{
			ConnectionId:    pulumi.String("dxcon-zzzzzzzz"),
			OwnerAccountId:  *pulumi.String(accepterCallerIdentity.AccountId),
			Vlan:            pulumi.Int(4094),
			AddressFamily:   pulumi.String("ipv4"),
			BgpAsn:          pulumi.Int(65352),
			CustomerAddress: pulumi.String("175.45.176.1/30"),
			AmazonAddress:   pulumi.String("175.45.176.2/30"),
			RouteFilterPrefixes: pulumi.StringArray{
				pulumi.String("210.52.109.0/24"),
				pulumi.String("175.45.176.0/22"),
			},
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewHostedPublicVirtualInterfaceAccepter(ctx, "accepterHostedPublicVirtualInterfaceAccepter", &directconnect.HostedPublicVirtualInterfaceAccepterArgs{
			VirtualInterfaceId: creator.ID(),
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
import com.pulumi.aws.directconnect.HostedPublicVirtualInterface;
import com.pulumi.aws.directconnect.HostedPublicVirtualInterfaceArgs;
import com.pulumi.aws.directconnect.HostedPublicVirtualInterfaceAccepter;
import com.pulumi.aws.directconnect.HostedPublicVirtualInterfaceAccepterArgs;
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

        var creator = new HostedPublicVirtualInterface("creator", HostedPublicVirtualInterfaceArgs.builder()        
            .connectionId("dxcon-zzzzzzzz")
            .ownerAccountId(accepterCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .vlan(4094)
            .addressFamily("ipv4")
            .bgpAsn(65352)
            .customerAddress("175.45.176.1/30")
            .amazonAddress("175.45.176.2/30")
            .routeFilterPrefixes(            
                "210.52.109.0/24",
                "175.45.176.0/22")
            .build());

        var accepterHostedPublicVirtualInterfaceAccepter = new HostedPublicVirtualInterfaceAccepter("accepterHostedPublicVirtualInterfaceAccepter", HostedPublicVirtualInterfaceAccepterArgs.builder()        
            .virtualInterfaceId(creator.id())
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
    type: aws:directconnect:HostedPublicVirtualInterface
    properties:
      connectionId: dxcon-zzzzzzzz
      ownerAccountId: ${accepterCallerIdentity.accountId}
      vlan: 4094
      addressFamily: ipv4
      bgpAsn: 65352
      customerAddress: 175.45.176.1/30
      amazonAddress: 175.45.176.2/30
      routeFilterPrefixes:
        - 210.52.109.0/24
        - 175.45.176.0/22
  # Accepter's side of the VIF.
  accepterHostedPublicVirtualInterfaceAccepter:
    type: aws:directconnect:HostedPublicVirtualInterfaceAccepter
    properties:
      virtualInterfaceId: ${creator.id}
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

Using `pulumi import`, import Direct Connect hosted public virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/hostedPublicVirtualInterfaceAccepter:HostedPublicVirtualInterfaceAccepter test dxvif-33cc44dd
```
 