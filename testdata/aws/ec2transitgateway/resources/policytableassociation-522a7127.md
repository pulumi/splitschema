Manages an EC2 Transit Gateway Policy Table association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.PolicyTableAssociation("example", {
    transitGatewayAttachmentId: aws_networkmanager_transit_gateway_peering.example.transit_gateway_peering_attachment_id,
    transitGatewayPolicyTableId: aws_ec2_transit_gateway_policy_table.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.PolicyTableAssociation("example",
    transit_gateway_attachment_id=aws_networkmanager_transit_gateway_peering["example"]["transit_gateway_peering_attachment_id"],
    transit_gateway_policy_table_id=aws_ec2_transit_gateway_policy_table["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.PolicyTableAssociation("example", new()
    {
        TransitGatewayAttachmentId = aws_networkmanager_transit_gateway_peering.Example.Transit_gateway_peering_attachment_id,
        TransitGatewayPolicyTableId = aws_ec2_transit_gateway_policy_table.Example.Id,
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
		_, err := ec2transitgateway.NewPolicyTableAssociation(ctx, "example", &ec2transitgateway.PolicyTableAssociationArgs{
			TransitGatewayAttachmentId:  pulumi.Any(aws_networkmanager_transit_gateway_peering.Example.Transit_gateway_peering_attachment_id),
			TransitGatewayPolicyTableId: pulumi.Any(aws_ec2_transit_gateway_policy_table.Example.Id),
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
import com.pulumi.aws.ec2transitgateway.PolicyTableAssociation;
import com.pulumi.aws.ec2transitgateway.PolicyTableAssociationArgs;
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
        var example = new PolicyTableAssociation("example", PolicyTableAssociationArgs.builder()        
            .transitGatewayAttachmentId(aws_networkmanager_transit_gateway_peering.example().transit_gateway_peering_attachment_id())
            .transitGatewayPolicyTableId(aws_ec2_transit_gateway_policy_table.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:PolicyTableAssociation
    properties:
      transitGatewayAttachmentId: ${aws_networkmanager_transit_gateway_peering.example.transit_gateway_peering_attachment_id}
      transitGatewayPolicyTableId: ${aws_ec2_transit_gateway_policy_table.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_policy_table_association` using the EC2 Transit Gateway Policy Table identifier, an underscore, and the EC2 Transit Gateway Attachment identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/policyTableAssociation:PolicyTableAssociation example tgw-rtb-12345678_tgw-attach-87654321
```
 