Provides a Direct Connect private virtual interface resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.directconnect.PrivateVirtualInterface("foo", {
    addressFamily: "ipv4",
    bgpAsn: 65352,
    connectionId: "dxcon-zzzzzzzz",
    vlan: 4094,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.directconnect.PrivateVirtualInterface("foo",
    address_family="ipv4",
    bgp_asn=65352,
    connection_id="dxcon-zzzzzzzz",
    vlan=4094)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.DirectConnect.PrivateVirtualInterface("foo", new()
    {
        AddressFamily = "ipv4",
        BgpAsn = 65352,
        ConnectionId = "dxcon-zzzzzzzz",
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
		_, err := directconnect.NewPrivateVirtualInterface(ctx, "foo", &directconnect.PrivateVirtualInterfaceArgs{
			AddressFamily: pulumi.String("ipv4"),
			BgpAsn:        pulumi.Int(65352),
			ConnectionId:  pulumi.String("dxcon-zzzzzzzz"),
			Vlan:          pulumi.Int(4094),
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
import com.pulumi.aws.directconnect.PrivateVirtualInterface;
import com.pulumi.aws.directconnect.PrivateVirtualInterfaceArgs;
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
        var foo = new PrivateVirtualInterface("foo", PrivateVirtualInterfaceArgs.builder()        
            .addressFamily("ipv4")
            .bgpAsn(65352)
            .connectionId("dxcon-zzzzzzzz")
            .vlan(4094)
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:directconnect:PrivateVirtualInterface
    properties:
      addressFamily: ipv4
      bgpAsn: 65352
      connectionId: dxcon-zzzzzzzz
      vlan: 4094
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect private virtual interfaces using the VIF `id`. For example:

```sh
 $ pulumi import aws:directconnect/privateVirtualInterface:PrivateVirtualInterface test dxvif-33cc44dd
```
 