Get information on an EC2 Transit Gateway VPN Attachment.

> EC2 Transit Gateway VPN Attachments are implicitly created by VPN Connections referencing an EC2 Transit Gateway so there is no managed resource. For ease, the `aws.ec2.VpnConnection` resource includes a `transit_gateway_attachment_id` attribute which can replace some usage of this data source. For tagging the attachment, see the `aws.ec2.Tag` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### By Transit Gateway and VPN Connection Identifiers

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getVpnAttachment({
    transitGatewayId: aws_ec2_transit_gateway.example.id,
    vpnConnectionId: aws_vpn_connection.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_vpn_attachment(transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
    vpn_connection_id=aws_vpn_connection["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetVpnAttachment.Invoke(new()
    {
        TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
        VpnConnectionId = aws_vpn_connection.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.GetVpnAttachment(ctx, &ec2transitgateway.GetVpnAttachmentArgs{
			TransitGatewayId: pulumi.StringRef(aws_ec2_transit_gateway.Example.Id),
			VpnConnectionId:  pulumi.StringRef(aws_vpn_connection.Example.Id),
		}, nil)
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
import com.pulumi.aws.ec2transitgateway.Ec2transitgatewayFunctions;
import com.pulumi.aws.ec2transitgateway.inputs.GetVpnAttachmentArgs;
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
        final var example = Ec2transitgatewayFunctions.getVpnAttachment(GetVpnAttachmentArgs.builder()
            .transitGatewayId(aws_ec2_transit_gateway.example().id())
            .vpnConnectionId(aws_vpn_connection.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getVpnAttachment
      Arguments:
        transitGatewayId: ${aws_ec2_transit_gateway.example.id}
        vpnConnectionId: ${aws_vpn_connection.example.id}
```
{{% /example %}}
{{% example %}}
### Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ec2transitgateway.getVpnAttachment({
    filters: [{
        name: "resource-id",
        values: ["some-resource"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2transitgateway.get_vpn_attachment(filters=[aws.ec2transitgateway.GetVpnAttachmentFilterArgs(
    name="resource-id",
    values=["some-resource"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ec2TransitGateway.GetVpnAttachment.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetVpnAttachmentFilterInputArgs
            {
                Name = "resource-id",
                Values = new[]
                {
                    "some-resource",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.GetVpnAttachment(ctx, &ec2transitgateway.GetVpnAttachmentArgs{
			Filters: []ec2transitgateway.GetVpnAttachmentFilter{
				{
					Name: "resource-id",
					Values: []string{
						"some-resource",
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.ec2transitgateway.Ec2transitgatewayFunctions;
import com.pulumi.aws.ec2transitgateway.inputs.GetVpnAttachmentArgs;
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
        final var test = Ec2transitgatewayFunctions.getVpnAttachment(GetVpnAttachmentArgs.builder()
            .filters(GetVpnAttachmentFilterArgs.builder()
                .name("resource-id")
                .values("some-resource")
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ec2transitgateway:getVpnAttachment
      Arguments:
        filters:
          - name: resource-id
            values:
              - some-resource
```
{{% /example %}}
{{% /examples %}}