Get information on an EC2 Transit Gateway Peering Attachment.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getPeeringAttachment({
    filters: [{
        name: "transit-gateway-attachment-id",
        values: ["tgw-attach-12345678"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_peering_attachment(filters=[aws.ec2transitgateway.GetPeeringAttachmentFilterArgs(
    name="transit-gateway-attachment-id",
    values=["tgw-attach-12345678"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetPeeringAttachment.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetPeeringAttachmentFilterInputArgs
            {
                Name = "transit-gateway-attachment-id",
                Values = new[]
                {
                    "tgw-attach-12345678",
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
		_, err := ec2transitgateway.LookupPeeringAttachment(ctx, &ec2transitgateway.LookupPeeringAttachmentArgs{
			Filters: []ec2transitgateway.GetPeeringAttachmentFilter{
				{
					Name: "transit-gateway-attachment-id",
					Values: []string{
						"tgw-attach-12345678",
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
import com.pulumi.aws.ec2transitgateway.inputs.GetPeeringAttachmentArgs;
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
        final var example = Ec2transitgatewayFunctions.getPeeringAttachment(GetPeeringAttachmentArgs.builder()
            .filters(GetPeeringAttachmentFilterArgs.builder()
                .name("transit-gateway-attachment-id")
                .values("tgw-attach-12345678")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getPeeringAttachment
      Arguments:
        filters:
          - name: transit-gateway-attachment-id
            values:
              - tgw-attach-12345678
```
{{% /example %}}
{{% example %}}
### By Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const attachment = aws.ec2transitgateway.getPeeringAttachment({
    id: "tgw-attach-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

attachment = aws.ec2transitgateway.get_peering_attachment(id="tgw-attach-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var attachment = Aws.Ec2TransitGateway.GetPeeringAttachment.Invoke(new()
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
		_, err := ec2transitgateway.LookupPeeringAttachment(ctx, &ec2transitgateway.LookupPeeringAttachmentArgs{
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
import com.pulumi.aws.ec2transitgateway.inputs.GetPeeringAttachmentArgs;
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
        final var attachment = Ec2transitgatewayFunctions.getPeeringAttachment(GetPeeringAttachmentArgs.builder()
            .id("tgw-attach-12345678")
            .build());

    }
}
```
```yaml
variables:
  attachment:
    fn::invoke:
      Function: aws:ec2transitgateway:getPeeringAttachment
      Arguments:
        id: tgw-attach-12345678
```
{{% /example %}}
{{% /examples %}}