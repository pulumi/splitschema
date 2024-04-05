Provides a Virtual Private Gateway attachment resource, allowing for an existing
hardware VPN gateway to be attached and/or detached from a VPC.

> **Note:** The `aws.ec2.VpnGateway`
resource can also automatically attach the Virtual Private Gateway it creates
to an existing VPC by setting the `vpc_id` attribute accordingly.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const network = new aws.ec2.Vpc("network", {cidrBlock: "10.0.0.0/16"});
const vpn = new aws.ec2.VpnGateway("vpn", {tags: {
    Name: "example-vpn-gateway",
}});
const vpnAttachment = new aws.ec2.VpnGatewayAttachment("vpnAttachment", {
    vpcId: network.id,
    vpnGatewayId: vpn.id,
});
```
```python
import pulumi
import pulumi_aws as aws

network = aws.ec2.Vpc("network", cidr_block="10.0.0.0/16")
vpn = aws.ec2.VpnGateway("vpn", tags={
    "Name": "example-vpn-gateway",
})
vpn_attachment = aws.ec2.VpnGatewayAttachment("vpnAttachment",
    vpc_id=network.id,
    vpn_gateway_id=vpn.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var network = new Aws.Ec2.Vpc("network", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var vpn = new Aws.Ec2.VpnGateway("vpn", new()
    {
        Tags = 
        {
            { "Name", "example-vpn-gateway" },
        },
    });

    var vpnAttachment = new Aws.Ec2.VpnGatewayAttachment("vpnAttachment", new()
    {
        VpcId = network.Id,
        VpnGatewayId = vpn.Id,
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
		network, err := ec2.NewVpc(ctx, "network", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		vpn, err := ec2.NewVpnGateway(ctx, "vpn", &ec2.VpnGatewayArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-vpn-gateway"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpnGatewayAttachment(ctx, "vpnAttachment", &ec2.VpnGatewayAttachmentArgs{
			VpcId:        network.ID(),
			VpnGatewayId: vpn.ID(),
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
import com.pulumi.aws.ec2.VpnGateway;
import com.pulumi.aws.ec2.VpnGatewayArgs;
import com.pulumi.aws.ec2.VpnGatewayAttachment;
import com.pulumi.aws.ec2.VpnGatewayAttachmentArgs;
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
        var network = new Vpc("network", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var vpn = new VpnGateway("vpn", VpnGatewayArgs.builder()        
            .tags(Map.of("Name", "example-vpn-gateway"))
            .build());

        var vpnAttachment = new VpnGatewayAttachment("vpnAttachment", VpnGatewayAttachmentArgs.builder()        
            .vpcId(network.id())
            .vpnGatewayId(vpn.id())
            .build());

    }
}
```
```yaml
resources:
  network:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  vpn:
    type: aws:ec2:VpnGateway
    properties:
      tags:
        Name: example-vpn-gateway
  vpnAttachment:
    type: aws:ec2:VpnGatewayAttachment
    properties:
      vpcId: ${network.id}
      vpnGatewayId: ${vpn.id}
```

See [Virtual Private Cloud](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Introduction.html)
and [Virtual Private Gateway](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_VPN.html) user
guides for more information.
{{% /example %}}
{{% /examples %}}

## Import

You cannot import this resource. 