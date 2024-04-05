Provides a Direct Connect BGP peer resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const peer = new aws.directconnect.BgpPeer("peer", {
    virtualInterfaceId: aws_dx_private_virtual_interface.foo.id,
    addressFamily: "ipv6",
    bgpAsn: 65351,
});
```
```python
import pulumi
import pulumi_aws as aws

peer = aws.directconnect.BgpPeer("peer",
    virtual_interface_id=aws_dx_private_virtual_interface["foo"]["id"],
    address_family="ipv6",
    bgp_asn=65351)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var peer = new Aws.DirectConnect.BgpPeer("peer", new()
    {
        VirtualInterfaceId = aws_dx_private_virtual_interface.Foo.Id,
        AddressFamily = "ipv6",
        BgpAsn = 65351,
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
		_, err := directconnect.NewBgpPeer(ctx, "peer", &directconnect.BgpPeerArgs{
			VirtualInterfaceId: pulumi.Any(aws_dx_private_virtual_interface.Foo.Id),
			AddressFamily:      pulumi.String("ipv6"),
			BgpAsn:             pulumi.Int(65351),
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
import com.pulumi.aws.directconnect.BgpPeer;
import com.pulumi.aws.directconnect.BgpPeerArgs;
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
        var peer = new BgpPeer("peer", BgpPeerArgs.builder()        
            .virtualInterfaceId(aws_dx_private_virtual_interface.foo().id())
            .addressFamily("ipv6")
            .bgpAsn(65351)
            .build());

    }
}
```
```yaml
resources:
  peer:
    type: aws:directconnect:BgpPeer
    properties:
      virtualInterfaceId: ${aws_dx_private_virtual_interface.foo.id}
      addressFamily: ipv6
      bgpAsn: 65351
```
{{% /example %}}
{{% /examples %}}