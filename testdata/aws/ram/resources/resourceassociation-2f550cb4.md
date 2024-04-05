Manages a Resource Access Manager (RAM) Resource Association.

> *NOTE:* Certain AWS resources (e.g., EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ram.ResourceAssociation("example", {
    resourceArn: aws_subnet.example.arn,
    resourceShareArn: aws_ram_resource_share.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ram.ResourceAssociation("example",
    resource_arn=aws_subnet["example"]["arn"],
    resource_share_arn=aws_ram_resource_share["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ram.ResourceAssociation("example", new()
    {
        ResourceArn = aws_subnet.Example.Arn,
        ResourceShareArn = aws_ram_resource_share.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.NewResourceAssociation(ctx, "example", &ram.ResourceAssociationArgs{
			ResourceArn:      pulumi.Any(aws_subnet.Example.Arn),
			ResourceShareArn: pulumi.Any(aws_ram_resource_share.Example.Arn),
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
import com.pulumi.aws.ram.ResourceAssociation;
import com.pulumi.aws.ram.ResourceAssociationArgs;
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
        var example = new ResourceAssociation("example", ResourceAssociationArgs.builder()        
            .resourceArn(aws_subnet.example().arn())
            .resourceShareArn(aws_ram_resource_share.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ram:ResourceAssociation
    properties:
      resourceArn: ${aws_subnet.example.arn}
      resourceShareArn: ${aws_ram_resource_share.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RAM Resource Associations using their Resource Share ARN and Resource ARN separated by a comma. For example:

```sh
 $ pulumi import aws:ram/resourceAssociation:ResourceAssociation example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,arn:aws:ec2:eu-west-1:123456789012:subnet/subnet-12345678
```
 