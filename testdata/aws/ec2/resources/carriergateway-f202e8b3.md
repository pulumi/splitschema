Manages an EC2 Carrier Gateway. See the AWS [documentation](https://docs.aws.amazon.com/vpc/latest/userguide/Carrier_Gateway.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.CarrierGateway("example", {
    vpcId: aws_vpc.example.id,
    tags: {
        Name: "example-carrier-gateway",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.CarrierGateway("example",
    vpc_id=aws_vpc["example"]["id"],
    tags={
        "Name": "example-carrier-gateway",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.CarrierGateway("example", new()
    {
        VpcId = aws_vpc.Example.Id,
        Tags = 
        {
            { "Name", "example-carrier-gateway" },
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
		_, err := ec2.NewCarrierGateway(ctx, "example", &ec2.CarrierGatewayArgs{
			VpcId: pulumi.Any(aws_vpc.Example.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-carrier-gateway"),
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
import com.pulumi.aws.ec2.CarrierGateway;
import com.pulumi.aws.ec2.CarrierGatewayArgs;
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
        var example = new CarrierGateway("example", CarrierGatewayArgs.builder()        
            .vpcId(aws_vpc.example().id())
            .tags(Map.of("Name", "example-carrier-gateway"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:CarrierGateway
    properties:
      vpcId: ${aws_vpc.example.id}
      tags:
        Name: example-carrier-gateway
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_carrier_gateway` using the carrier gateway's ID. For example:

```sh
 $ pulumi import aws:ec2/carrierGateway:CarrierGateway example cgw-12345
```
 