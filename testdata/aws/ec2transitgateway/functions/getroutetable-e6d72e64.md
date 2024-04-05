Get information on an EC2 Transit Gateway Route Table.

{{% examples %}}
## Example Usage
{{% example %}}
### By Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getRouteTable({
    filters: [
        {
            name: "default-association-route-table",
            values: ["true"],
        },
        {
            name: "transit-gateway-id",
            values: ["tgw-12345678"],
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_route_table(filters=[
    aws.ec2transitgateway.GetRouteTableFilterArgs(
        name="default-association-route-table",
        values=["true"],
    ),
    aws.ec2transitgateway.GetRouteTableFilterArgs(
        name="transit-gateway-id",
        values=["tgw-12345678"],
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetRouteTable.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetRouteTableFilterInputArgs
            {
                Name = "default-association-route-table",
                Values = new[]
                {
                    "true",
                },
            },
            new Aws.Ec2TransitGateway.Inputs.GetRouteTableFilterInputArgs
            {
                Name = "transit-gateway-id",
                Values = new[]
                {
                    "tgw-12345678",
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
		_, err := ec2transitgateway.LookupRouteTable(ctx, &ec2transitgateway.LookupRouteTableArgs{
			Filters: []ec2transitgateway.GetRouteTableFilter{
				{
					Name: "default-association-route-table",
					Values: []string{
						"true",
					},
				},
				{
					Name: "transit-gateway-id",
					Values: []string{
						"tgw-12345678",
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
import com.pulumi.aws.ec2transitgateway.inputs.GetRouteTableArgs;
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
        final var example = Ec2transitgatewayFunctions.getRouteTable(GetRouteTableArgs.builder()
            .filters(            
                GetRouteTableFilterArgs.builder()
                    .name("default-association-route-table")
                    .values("true")
                    .build(),
                GetRouteTableFilterArgs.builder()
                    .name("transit-gateway-id")
                    .values("tgw-12345678")
                    .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getRouteTable
      Arguments:
        filters:
          - name: default-association-route-table
            values:
              - 'true'
          - name: transit-gateway-id
            values:
              - tgw-12345678
```
{{% /example %}}
{{% example %}}
### By Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getRouteTable({
    id: "tgw-rtb-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_route_table(id="tgw-rtb-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetRouteTable.Invoke(new()
    {
        Id = "tgw-rtb-12345678",
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
		_, err := ec2transitgateway.LookupRouteTable(ctx, &ec2transitgateway.LookupRouteTableArgs{
			Id: pulumi.StringRef("tgw-rtb-12345678"),
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
import com.pulumi.aws.ec2transitgateway.inputs.GetRouteTableArgs;
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
        final var example = Ec2transitgatewayFunctions.getRouteTable(GetRouteTableArgs.builder()
            .id("tgw-rtb-12345678")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getRouteTable
      Arguments:
        id: tgw-rtb-12345678
```
{{% /example %}}
{{% /examples %}}