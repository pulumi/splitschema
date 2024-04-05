Manages an EC2 Transit Gateway Policy Table.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.PolicyTable("example", {
    transitGatewayId: aws_ec2_transit_gateway.example.id,
    tags: {
        Name: "Example Policy Table",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.PolicyTable("example",
    transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
    tags={
        "Name": "Example Policy Table",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.PolicyTable("example", new()
    {
        TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
        Tags = 
        {
            { "Name", "Example Policy Table" },
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
		_, err := ec2transitgateway.NewPolicyTable(ctx, "example", &ec2transitgateway.PolicyTableArgs{
			TransitGatewayId: pulumi.Any(aws_ec2_transit_gateway.Example.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Policy Table"),
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
import com.pulumi.aws.ec2transitgateway.PolicyTable;
import com.pulumi.aws.ec2transitgateway.PolicyTableArgs;
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
        var example = new PolicyTable("example", PolicyTableArgs.builder()        
            .transitGatewayId(aws_ec2_transit_gateway.example().id())
            .tags(Map.of("Name", "Example Policy Table"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:PolicyTable
    properties:
      transitGatewayId: ${aws_ec2_transit_gateway.example.id}
      tags:
        Name: Example Policy Table
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_policy_table` using the EC2 Transit Gateway Policy Table identifier. For example:

```sh
 $ pulumi import aws:ec2transitgateway/policyTable:PolicyTable example tgw-rtb-12345678
```
 