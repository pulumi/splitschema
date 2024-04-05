Provides details about a specific VPC NAT Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const default = aws.ec2.getNatGateway({
    subnetId: aws_subnet["public"].id,
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.ec2.get_nat_gateway(subnet_id=aws_subnet["public"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = Aws.Ec2.GetNatGateway.Invoke(new()
    {
        SubnetId = aws_subnet.Public.Id,
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
		_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
			SubnetId: pulumi.StringRef(aws_subnet.Public.Id),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetNatGatewayArgs;
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
        final var default = Ec2Functions.getNatGateway(GetNatGatewayArgs.builder()
            .subnetId(aws_subnet.public().id())
            .build());

    }
}
```
```yaml
variables:
  default:
    fn::invoke:
      Function: aws:ec2:getNatGateway
      Arguments:
        subnetId: ${aws_subnet.public.id}
```
{{% /example %}}
{{% example %}}
### With tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const default = aws.ec2.getNatGateway({
    subnetId: aws_subnet["public"].id,
    tags: {
        Name: "gw NAT",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.ec2.get_nat_gateway(subnet_id=aws_subnet["public"]["id"],
    tags={
        "Name": "gw NAT",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = Aws.Ec2.GetNatGateway.Invoke(new()
    {
        SubnetId = aws_subnet.Public.Id,
        Tags = 
        {
            { "Name", "gw NAT" },
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
		_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
			SubnetId: pulumi.StringRef(aws_subnet.Public.Id),
			Tags: map[string]interface{}{
				"Name": "gw NAT",
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetNatGatewayArgs;
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
        final var default = Ec2Functions.getNatGateway(GetNatGatewayArgs.builder()
            .subnetId(aws_subnet.public().id())
            .tags(Map.of("Name", "gw NAT"))
            .build());

    }
}
```
```yaml
variables:
  default:
    fn::invoke:
      Function: aws:ec2:getNatGateway
      Arguments:
        subnetId: ${aws_subnet.public.id}
        tags:
          Name: gw NAT
```
{{% /example %}}
{{% /examples %}}