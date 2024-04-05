Provides a resource to create a VPC VPN Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpnGw = new aws.ec2.VpnGateway("vpnGw", {
    vpcId: aws_vpc.main.id,
    tags: {
        Name: "main",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

vpn_gw = aws.ec2.VpnGateway("vpnGw",
    vpc_id=aws_vpc["main"]["id"],
    tags={
        "Name": "main",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var vpnGw = new Aws.Ec2.VpnGateway("vpnGw", new()
    {
        VpcId = aws_vpc.Main.Id,
        Tags = 
        {
            { "Name", "main" },
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
		_, err := ec2.NewVpnGateway(ctx, "vpnGw", &ec2.VpnGatewayArgs{
			VpcId: pulumi.Any(aws_vpc.Main.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
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
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
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
        var vpnGw = new VpnGateway("vpnGw", VpnGatewayArgs.builder()        
            .vpcId(aws_vpc.main().id())
            .tags(Map.of("Name", "main"))
            .build());

    }
}
```
```yaml
resources:
  vpnGw:
    type: aws:ec2:VpnGateway
    properties:
      vpcId: ${aws_vpc.main.id}
      tags:
        Name: main
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPN Gateways using the VPN gateway `id`. For example:

```sh
 $ pulumi import aws:ec2/vpnGateway:VpnGateway testvpngateway vgw-9a4cacf3
```
 