Get information on an EC2 Transit Gateway VPC Attachment.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getVpcAttachment({
    filters: [{
        name: "vpc-id",
        values: ["vpc-12345678"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_vpc_attachment(filters=[aws.ec2transitgateway.GetVpcAttachmentFilterArgs(
    name="vpc-id",
    values=["vpc-12345678"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetVpcAttachment.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetVpcAttachmentFilterInputArgs
            {
                Name = "vpc-id",
                Values = new[]
                {
                    "vpc-12345678",
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
		_, err := ec2transitgateway.LookupVpcAttachment(ctx, &ec2transitgateway.LookupVpcAttachmentArgs{
			Filters: []ec2transitgateway.GetVpcAttachmentFilter{
				{
					Name: "vpc-id",
					Values: []string{
						"vpc-12345678",
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
import com.pulumi.aws.ec2transitgateway.inputs.GetVpcAttachmentArgs;
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
        final var example = Ec2transitgatewayFunctions.getVpcAttachment(GetVpcAttachmentArgs.builder()
            .filters(GetVpcAttachmentFilterArgs.builder()
                .name("vpc-id")
                .values("vpc-12345678")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getVpcAttachment
      Arguments:
        filters:
          - name: vpc-id
            values:
              - vpc-12345678
```
{{% /example %}}
{{% example %}}
### By Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getVpcAttachment({
    id: "tgw-attach-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_vpc_attachment(id="tgw-attach-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetVpcAttachment.Invoke(new()
    {
        Id = "tgw-attach-12345678",
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
		_, err := ec2transitgateway.LookupVpcAttachment(ctx, &ec2transitgateway.LookupVpcAttachmentArgs{
			Id: pulumi.StringRef("tgw-attach-12345678"),
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
import com.pulumi.aws.ec2transitgateway.inputs.GetVpcAttachmentArgs;
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
        final var example = Ec2transitgatewayFunctions.getVpcAttachment(GetVpcAttachmentArgs.builder()
            .id("tgw-attach-12345678")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getVpcAttachment
      Arguments:
        id: tgw-attach-12345678
```
{{% /example %}}
{{% /examples %}}