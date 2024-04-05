Manages the accepter's side of an EC2 Transit Gateway VPC Attachment.

When a cross-account (requester's AWS account differs from the accepter's AWS account) EC2 Transit Gateway VPC Attachment
is created, an EC2 Transit Gateway VPC Attachment resource is automatically created in the accepter's account.
The requester can use the `aws.ec2transitgateway.VpcAttachment` resource to manage its side of the connection
and the accepter can use the `aws.ec2transitgateway.VpcAttachmentAccepter` resource to "adopt" its side of the
connection into management.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.VpcAttachmentAccepter("example", {
    transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment.example.id,
    tags: {
        Name: "Example cross-account attachment",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.VpcAttachmentAccepter("example",
    transit_gateway_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"],
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
    var example = new Aws.Ec2TransitGateway.VpcAttachmentAccepter("example", new()
    {
        TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
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
		_, err := ec2transitgateway.NewVpcAttachmentAccepter(ctx, "example", &ec2transitgateway.VpcAttachmentAccepterArgs{
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
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
import com.pulumi.aws.ec2transitgateway.VpcAttachmentAccepter;
import com.pulumi.aws.ec2transitgateway.VpcAttachmentAccepterArgs;
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
        var example = new VpcAttachmentAccepter("example", VpcAttachmentAccepterArgs.builder()        
            .transitGatewayAttachmentId(aws_ec2_transit_gateway_vpc_attachment.example().id())
            .tags(Map.of("Name", "Example cross-account attachment"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:VpcAttachmentAccepter
    properties:
      transitGatewayAttachmentId: ${aws_ec2_transit_gateway_vpc_attachment.example.id}
      tags:
        Name: Example cross-account attachment
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_vpc_attachment_accepter` using the EC2 Transit Gateway Attachment identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter example tgw-attach-12345678
```
 