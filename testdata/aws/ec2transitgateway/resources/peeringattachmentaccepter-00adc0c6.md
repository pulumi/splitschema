Manages the accepter's side of an EC2 Transit Gateway Peering Attachment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.PeeringAttachmentAccepter("example", {
    transitGatewayAttachmentId: aws_ec2_transit_gateway_peering_attachment.example.id,
    tags: {
        Name: "Example cross-account attachment",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.PeeringAttachmentAccepter("example",
    transit_gateway_attachment_id=aws_ec2_transit_gateway_peering_attachment["example"]["id"],
    tags={
        "Name": "Example cross-account attachment",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.PeeringAttachmentAccepter("example", new()
    {
        TransitGatewayAttachmentId = aws_ec2_transit_gateway_peering_attachment.Example.Id,
        Tags = 
        {
            { "Name", "Example cross-account attachment" },
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
		_, err := ec2transitgateway.NewPeeringAttachmentAccepter(ctx, "example", &ec2transitgateway.PeeringAttachmentAccepterArgs{
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_peering_attachment.Example.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example cross-account attachment"),
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
import com.pulumi.aws.ec2transitgateway.PeeringAttachmentAccepter;
import com.pulumi.aws.ec2transitgateway.PeeringAttachmentAccepterArgs;
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
        var example = new PeeringAttachmentAccepter("example", PeeringAttachmentAccepterArgs.builder()        
            .transitGatewayAttachmentId(aws_ec2_transit_gateway_peering_attachment.example().id())
            .tags(Map.of("Name", "Example cross-account attachment"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:PeeringAttachmentAccepter
    properties:
      transitGatewayAttachmentId: ${aws_ec2_transit_gateway_peering_attachment.example.id}
      tags:
        Name: Example cross-account attachment
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_peering_attachment_accepter` using the EC2 Transit Gateway Attachment identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/peeringAttachmentAccepter:PeeringAttachmentAccepter example tgw-attach-12345678
```
 