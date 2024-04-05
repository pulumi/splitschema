Get information on EC2 Transit Gateway VPC Attachments.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const filtered = aws.ec2transitgateway.getVpcAttachments({
    filters: [{
        name: "state",
        values: ["pendingAcceptance"],
    }],
});
const unit = .map(__index => (aws.ec2transitgateway.getVpcAttachment({
    id: _arg0_.ids[__index],
})));
```
```python
import pulumi
import pulumi_aws as aws

filtered = aws.ec2transitgateway.get_vpc_attachments(filters=[aws.ec2transitgateway.GetVpcAttachmentsFilterArgs(
    name="state",
    values=["pendingAcceptance"],
)])
unit = [aws.ec2transitgateway.get_vpc_attachment(id=filtered.ids[__index]) for __index in range(len(filtered.ids))]
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var filtered = Aws.Ec2TransitGateway.GetVpcAttachments.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetVpcAttachmentsFilterInputArgs
            {
                Name = "state",
                Values = new[]
                {
                    "pendingAcceptance",
                },
            },
        },
    });

    var unit = ;

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
		filtered, err := ec2transitgateway.GetVpcAttachments(ctx, &ec2transitgateway.GetVpcAttachmentsArgs{
			Filters: []ec2transitgateway.GetVpcAttachmentsFilter{
				{
					Name: "state",
					Values: []string{
						"pendingAcceptance",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_ := "TODO: For expression"
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}