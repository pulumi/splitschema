Provides a Direct Connect hosted transit virtual interface resource.
This resource represents the allocator's side of the hosted virtual interface.
A hosted virtual interface is a virtual interface that is owned by another AWS account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.HostedTransitVirtualInterface("example", {
    connectionId: aws_dx_connection.example.id,
    vlan: 4094,
    addressFamily: "ipv4",
    bgpAsn: 65352,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.HostedTransitVirtualInterface("example",
    connection_id=aws_dx_connection["example"]["id"],
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
    var example = new Aws.DirectConnect.HostedTransitVirtualInterface("example", new()
    {
        ConnectionId = aws_dx_connection.Example.Id,
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
		_, err := directconnect.NewHostedTransitVirtualInterface(ctx, "example", &directconnect.HostedTransitVirtualInterfaceArgs{
			ConnectionId:  pulumi.Any(aws_dx_connection.Example.Id),
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
import com.pulumi.aws.directconnect.HostedTransitVirtualInterface;
import com.pulumi.aws.directconnect.HostedTransitVirtualInterfaceArgs;
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
        var example = new HostedTransitVirtualInterface("example", HostedTransitVirtualInterfaceArgs.builder()        
            .connectionId(aws_dx_connection.example().id())
            .vlan(4094)
            .addressFamily("ipv4")
            .bgpAsn(65352)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directconnect:HostedTransitVirtualInterface
    properties:
      connectionId: ${aws_dx_connection.example.id}
      vlan: 4094
      addressFamily: ipv4
      bgpAsn: 65352
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect hosted transit virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/hostedTransitVirtualInterface:HostedTransitVirtualInterface test dxvif-33cc44dd
```
 