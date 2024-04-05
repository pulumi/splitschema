Provides a resource to create a VPC Internet Gateway Attachment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.1.0.0/16"});
const exampleInternetGateway = new aws.ec2.InternetGateway("exampleInternetGateway", {});
const exampleInternetGatewayAttachment = new aws.ec2.InternetGatewayAttachment("exampleInternetGatewayAttachment", {
    internetGatewayId: exampleInternetGateway.id,
    vpcId: exampleVpc.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.1.0.0/16")
example_internet_gateway = aws.ec2.InternetGateway("exampleInternetGateway")
example_internet_gateway_attachment = aws.ec2.InternetGatewayAttachment("exampleInternetGatewayAttachment",
    internet_gateway_id=example_internet_gateway.id,
    vpc_id=example_vpc.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var exampleInternetGateway = new Aws.Ec2.InternetGateway("exampleInternetGateway");

    var exampleInternetGatewayAttachment = new Aws.Ec2.InternetGatewayAttachment("exampleInternetGatewayAttachment", new()
    {
        InternetGatewayId = exampleInternetGateway.Id,
        VpcId = exampleVpc.Id,
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
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		exampleInternetGateway, err := ec2.NewInternetGateway(ctx, "exampleInternetGateway", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInternetGatewayAttachment(ctx, "exampleInternetGatewayAttachment", &ec2.InternetGatewayAttachmentArgs{
			InternetGatewayId: exampleInternetGateway.ID(),
			VpcId:             exampleVpc.ID(),
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
import com.pulumi.aws.ec2.InternetGateway;
import com.pulumi.aws.ec2.InternetGatewayAttachment;
import com.pulumi.aws.ec2.InternetGatewayAttachmentArgs;
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
        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        var exampleInternetGateway = new InternetGateway("exampleInternetGateway");

        var exampleInternetGatewayAttachment = new InternetGatewayAttachment("exampleInternetGatewayAttachment", InternetGatewayAttachmentArgs.builder()        
            .internetGatewayId(exampleInternetGateway.id())
            .vpcId(exampleVpc.id())
            .build());

    }
}
```
```yaml
resources:
  exampleInternetGatewayAttachment:
    type: aws:ec2:InternetGatewayAttachment
    properties:
      internetGatewayId: ${exampleInternetGateway.id}
      vpcId: ${exampleVpc.id}
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  exampleInternetGateway:
    type: aws:ec2:InternetGateway
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Internet Gateway Attachments using the `id`. For example:

```sh
 $ pulumi import aws:ec2/internetGatewayAttachment:InternetGatewayAttachment example igw-c0a643a9:vpc-123456
```
 