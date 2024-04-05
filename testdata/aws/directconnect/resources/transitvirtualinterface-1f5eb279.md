Provides a Direct Connect transit virtual interface resource.
A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one or more transit gateways.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGateway = new aws.directconnect.Gateway("exampleGateway", {amazonSideAsn: "64512"});
const exampleTransitVirtualInterface = new aws.directconnect.TransitVirtualInterface("exampleTransitVirtualInterface", {
    connectionId: aws_dx_connection.example.id,
    dxGatewayId: exampleGateway.id,
    vlan: 4094,
    addressFamily: "ipv4",
    bgpAsn: 65352,
});
```
```python
import pulumi
import pulumi_aws as aws

example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
example_transit_virtual_interface = aws.directconnect.TransitVirtualInterface("exampleTransitVirtualInterface",
    connection_id=aws_dx_connection["example"]["id"],
    dx_gateway_id=example_gateway.id,
    vlan=4094,
    address_family="ipv4",
    bgp_asn=65352)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGateway = new Aws.DirectConnect.Gateway("exampleGateway", new()
    {
        AmazonSideAsn = "64512",
    });

    var exampleTransitVirtualInterface = new Aws.DirectConnect.TransitVirtualInterface("exampleTransitVirtualInterface", new()
    {
        ConnectionId = aws_dx_connection.Example.Id,
        DxGatewayId = exampleGateway.Id,
        Vlan = 4094,
        AddressFamily = "ipv4",
        BgpAsn = 65352,
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
		exampleGateway, err := directconnect.NewGateway(ctx, "exampleGateway", &directconnect.GatewayArgs{
			AmazonSideAsn: pulumi.String("64512"),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewTransitVirtualInterface(ctx, "exampleTransitVirtualInterface", &directconnect.TransitVirtualInterfaceArgs{
			ConnectionId:  pulumi.Any(aws_dx_connection.Example.Id),
			DxGatewayId:   exampleGateway.ID(),
			Vlan:          pulumi.Int(4094),
			AddressFamily: pulumi.String("ipv4"),
			BgpAsn:        pulumi.Int(65352),
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
import com.pulumi.aws.directconnect.Gateway;
import com.pulumi.aws.directconnect.GatewayArgs;
import com.pulumi.aws.directconnect.TransitVirtualInterface;
import com.pulumi.aws.directconnect.TransitVirtualInterfaceArgs;
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
        var exampleGateway = new Gateway("exampleGateway", GatewayArgs.builder()        
            .amazonSideAsn(64512)
            .build());

        var exampleTransitVirtualInterface = new TransitVirtualInterface("exampleTransitVirtualInterface", TransitVirtualInterfaceArgs.builder()        
            .connectionId(aws_dx_connection.example().id())
            .dxGatewayId(exampleGateway.id())
            .vlan(4094)
            .addressFamily("ipv4")
            .bgpAsn(65352)
            .build());

    }
}
```
```yaml
resources:
  exampleGateway:
    type: aws:directconnect:Gateway
    properties:
      amazonSideAsn: 64512
  exampleTransitVirtualInterface:
    type: aws:directconnect:TransitVirtualInterface
    properties:
      connectionId: ${aws_dx_connection.example.id}
      dxGatewayId: ${exampleGateway.id}
      vlan: 4094
      addressFamily: ipv4
      bgpAsn: 65352
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect transit virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/transitVirtualInterface:TransitVirtualInterface test dxvif-33cc44dd
```
 