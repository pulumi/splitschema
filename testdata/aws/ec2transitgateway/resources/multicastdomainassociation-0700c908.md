Associates the specified subnet and transit gateway attachment with the specified transit gateway multicast domain.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("exampleTransitGateway", {multicastSupport: "enable"});
const exampleVpcAttachment = new aws.ec2transitgateway.VpcAttachment("exampleVpcAttachment", {
    subnetIds: [aws_subnet.example.id],
    transitGatewayId: exampleTransitGateway.id,
    vpcId: aws_vpc.example.id,
});
const exampleMulticastDomain = new aws.ec2transitgateway.MulticastDomain("exampleMulticastDomain", {transitGatewayId: exampleTransitGateway.id});
const exampleMulticastDomainAssociation = new aws.ec2transitgateway.MulticastDomainAssociation("exampleMulticastDomainAssociation", {
    subnetId: aws_subnet.example.id,
    transitGatewayAttachmentId: exampleVpcAttachment.id,
    transitGatewayMulticastDomainId: exampleMulticastDomain.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway", multicast_support="enable")
example_vpc_attachment = aws.ec2transitgateway.VpcAttachment("exampleVpcAttachment",
    subnet_ids=[aws_subnet["example"]["id"]],
    transit_gateway_id=example_transit_gateway.id,
    vpc_id=aws_vpc["example"]["id"])
example_multicast_domain = aws.ec2transitgateway.MulticastDomain("exampleMulticastDomain", transit_gateway_id=example_transit_gateway.id)
example_multicast_domain_association = aws.ec2transitgateway.MulticastDomainAssociation("exampleMulticastDomainAssociation",
    subnet_id=aws_subnet["example"]["id"],
    transit_gateway_attachment_id=example_vpc_attachment.id,
    transit_gateway_multicast_domain_id=example_multicast_domain.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleTransitGateway = new Aws.Ec2TransitGateway.TransitGateway("exampleTransitGateway", new()
    {
        MulticastSupport = "enable",
    });

    var exampleVpcAttachment = new Aws.Ec2TransitGateway.VpcAttachment("exampleVpcAttachment", new()
    {
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        TransitGatewayId = exampleTransitGateway.Id,
        VpcId = aws_vpc.Example.Id,
    });

    var exampleMulticastDomain = new Aws.Ec2TransitGateway.MulticastDomain("exampleMulticastDomain", new()
    {
        TransitGatewayId = exampleTransitGateway.Id,
    });

    var exampleMulticastDomainAssociation = new Aws.Ec2TransitGateway.MulticastDomainAssociation("exampleMulticastDomainAssociation", new()
    {
        SubnetId = aws_subnet.Example.Id,
        TransitGatewayAttachmentId = exampleVpcAttachment.Id,
        TransitGatewayMulticastDomainId = exampleMulticastDomain.Id,
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
		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", &ec2transitgateway.TransitGatewayArgs{
			MulticastSupport: pulumi.String("enable"),
		})
		if err != nil {
			return err
		}
		exampleVpcAttachment, err := ec2transitgateway.NewVpcAttachment(ctx, "exampleVpcAttachment", &ec2transitgateway.VpcAttachmentArgs{
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			TransitGatewayId: exampleTransitGateway.ID(),
			VpcId:            pulumi.Any(aws_vpc.Example.Id),
		})
		if err != nil {
			return err
		}
		exampleMulticastDomain, err := ec2transitgateway.NewMulticastDomain(ctx, "exampleMulticastDomain", &ec2transitgateway.MulticastDomainArgs{
			TransitGatewayId: exampleTransitGateway.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2transitgateway.NewMulticastDomainAssociation(ctx, "exampleMulticastDomainAssociation", &ec2transitgateway.MulticastDomainAssociationArgs{
			SubnetId:                        pulumi.Any(aws_subnet.Example.Id),
			TransitGatewayAttachmentId:      exampleVpcAttachment.ID(),
			TransitGatewayMulticastDomainId: exampleMulticastDomain.ID(),
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
import com.pulumi.aws.ec2transitgateway.TransitGateway;
import com.pulumi.aws.ec2transitgateway.TransitGatewayArgs;
import com.pulumi.aws.ec2transitgateway.VpcAttachment;
import com.pulumi.aws.ec2transitgateway.VpcAttachmentArgs;
import com.pulumi.aws.ec2transitgateway.MulticastDomain;
import com.pulumi.aws.ec2transitgateway.MulticastDomainArgs;
import com.pulumi.aws.ec2transitgateway.MulticastDomainAssociation;
import com.pulumi.aws.ec2transitgateway.MulticastDomainAssociationArgs;
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
        var exampleTransitGateway = new TransitGateway("exampleTransitGateway", TransitGatewayArgs.builder()        
            .multicastSupport("enable")
            .build());

        var exampleVpcAttachment = new VpcAttachment("exampleVpcAttachment", VpcAttachmentArgs.builder()        
            .subnetIds(aws_subnet.example().id())
            .transitGatewayId(exampleTransitGateway.id())
            .vpcId(aws_vpc.example().id())
            .build());

        var exampleMulticastDomain = new MulticastDomain("exampleMulticastDomain", MulticastDomainArgs.builder()        
            .transitGatewayId(exampleTransitGateway.id())
            .build());

        var exampleMulticastDomainAssociation = new MulticastDomainAssociation("exampleMulticastDomainAssociation", MulticastDomainAssociationArgs.builder()        
            .subnetId(aws_subnet.example().id())
            .transitGatewayAttachmentId(exampleVpcAttachment.id())
            .transitGatewayMulticastDomainId(exampleMulticastDomain.id())
            .build());

    }
}
```
```yaml
resources:
  exampleTransitGateway:
    type: aws:ec2transitgateway:TransitGateway
    properties:
      multicastSupport: enable
  exampleVpcAttachment:
    type: aws:ec2transitgateway:VpcAttachment
    properties:
      subnetIds:
        - ${aws_subnet.example.id}
      transitGatewayId: ${exampleTransitGateway.id}
      vpcId: ${aws_vpc.example.id}
  exampleMulticastDomain:
    type: aws:ec2transitgateway:MulticastDomain
    properties:
      transitGatewayId: ${exampleTransitGateway.id}
  exampleMulticastDomainAssociation:
    type: aws:ec2transitgateway:MulticastDomainAssociation
    properties:
      subnetId: ${aws_subnet.example.id}
      transitGatewayAttachmentId: ${exampleVpcAttachment.id}
      transitGatewayMulticastDomainId: ${exampleMulticastDomain.id}
```
{{% /example %}}
{{% /examples %}}