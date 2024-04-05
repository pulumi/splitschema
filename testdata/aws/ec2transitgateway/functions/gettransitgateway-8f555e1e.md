Get information on an EC2 Transit Gateway.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getTransitGateway({
    filters: [{
        name: "options.amazon-side-asn",
        values: ["64512"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_transit_gateway(filters=[aws.ec2transitgateway.GetTransitGatewayFilterArgs(
    name="options.amazon-side-asn",
    values=["64512"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetTransitGateway.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetTransitGatewayFilterInputArgs
            {
                Name = "options.amazon-side-asn",
                Values = new[]
                {
                    "64512",
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
		_, err := ec2transitgateway.LookupTransitGateway(ctx, &ec2transitgateway.LookupTransitGatewayArgs{
			Filters: []ec2transitgateway.GetTransitGatewayFilter{
				{
					Name: "options.amazon-side-asn",
					Values: []string{
						"64512",
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
import com.pulumi.aws.ec2transitgateway.inputs.GetTransitGatewayArgs;
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
        final var example = Ec2transitgatewayFunctions.getTransitGateway(GetTransitGatewayArgs.builder()
            .filters(GetTransitGatewayFilterArgs.builder()
                .name("options.amazon-side-asn")
                .values("64512")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getTransitGateway
      Arguments:
        filters:
          - name: options.amazon-side-asn
            values:
              - '64512'
```
{{% /example %}}
{{% example %}}
### By Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getTransitGateway({
    id: "tgw-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_transit_gateway(id="tgw-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetTransitGateway.Invoke(new()
    {
        Id = "tgw-12345678",
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
		_, err := ec2transitgateway.LookupTransitGateway(ctx, &ec2transitgateway.LookupTransitGatewayArgs{
			Id: pulumi.StringRef("tgw-12345678"),
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
import com.pulumi.aws.ec2transitgateway.inputs.GetTransitGatewayArgs;
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
        final var example = Ec2transitgatewayFunctions.getTransitGateway(GetTransitGatewayArgs.builder()
            .id("tgw-12345678")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getTransitGateway
      Arguments:
        id: tgw-12345678
```
{{% /example %}}
{{% /examples %}}