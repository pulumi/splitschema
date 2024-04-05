Provides a Direct Connect hosted public virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
A hosted virtual interface is a virtual interface that is owned by another AWS account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.directconnect.HostedPublicVirtualInterface("foo", {
    addressFamily: "ipv4",
    amazonAddress: "175.45.176.2/30",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    customerAddress: "175.45.176.1/30",
    routeFilterPrefixes: [
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    vlan: 4094,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.directconnect.HostedPublicVirtualInterface("foo",
    address_family="ipv4",
    amazon_address="175.45.176.2/30",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    customer_address="175.45.176.1/30",
    route_filter_prefixes=[
        "210.52.109.0/24",
        "175.45.176.0/22",
    ],
    vlan=4094)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.DirectConnect.HostedPublicVirtualInterface("foo", new()
    {
        AddressFamily = "ipv4",
        AmazonAddress = "175.45.176.2/30",
        BgpAsn = 65352,
        ConnectionId = "dxcon-zzzzzzzz",
        CustomerAddress = "175.45.176.1/30",
        RouteFilterPrefixes = new[]
        {
            "210.52.109.0/24",
            "175.45.176.0/22",
        },
        Vlan = 4094,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewHostedPublicVirtualInterface(ctx, "foo", &directconnect.HostedPublicVirtualInterfaceArgs{
			AddressFamily:   pulumi.String("ipv4"),
			AmazonAddress:   pulumi.String("175.45.176.2/30"),
			BgpAsn:          pulumi.Int(65352),
			ConnectionId:    pulumi.String("dxcon-zzzzzzzz"),
			CustomerAddress: pulumi.String("175.45.176.1/30"),
			RouteFilterPrefixes: pulumi.StringArray{
				pulumi.String("210.52.109.0/24"),
				pulumi.String("175.45.176.0/22"),
			},
			Vlan: pulumi.Int(4094),
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
import com.pulumi.aws.directconnect.HostedPublicVirtualInterface;
import com.pulumi.aws.directconnect.HostedPublicVirtualInterfaceArgs;
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
        var foo = new HostedPublicVirtualInterface("foo", HostedPublicVirtualInterfaceArgs.builder()        
            .addressFamily("ipv4")
            .amazonAddress("175.45.176.2/30")
            .bgpAsn(65352)
            .connectionId("dxcon-zzzzzzzz")
            .customerAddress("175.45.176.1/30")
            .routeFilterPrefixes(            
                "210.52.109.0/24",
                "175.45.176.0/22")
            .vlan(4094)
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:directconnect:HostedPublicVirtualInterface
    properties:
      addressFamily: ipv4
      amazonAddress: 175.45.176.2/30
      bgpAsn: 65352
      connectionId: dxcon-zzzzzzzz
      customerAddress: 175.45.176.1/30
      routeFilterPrefixes:
        - 210.52.109.0/24
        - 175.45.176.0/22
      vlan: 4094
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect hosted public virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface test dxvif-33cc44dd
```
 