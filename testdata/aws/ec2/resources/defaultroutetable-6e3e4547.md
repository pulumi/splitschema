Provides a resource to manage a default route table of a VPC. This resource can manage the default route table of the default or a non-default VPC.

> **NOTE:** This is an advanced resource with special caveats. Please read this document in its entirety before using this resource. The `aws.ec2.DefaultRouteTable` resource behaves differently from normal resources. This provider does not _create_ this resource but instead attempts to "adopt" it into management. **Do not** use both `aws.ec2.DefaultRouteTable` to manage a default route table **and** `aws.ec2.MainRouteTableAssociation` with the same VPC due to possible route conflicts. See aws.ec2.MainRouteTableAssociation documentation for more details.

Every VPC has a default route table that can be managed but not destroyed. When the provider first adopts a default route table, it **immediately removes all defined routes**. It then proceeds to create any routes specified in the configuration. This step is required so that only the routes specified in the configuration exist in the default route table.

For more information, see the Amazon VPC User Guide on [Route Tables](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html). For information about managing normal route tables in this provider, see `aws.ec2.RouteTable`.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.DefaultRouteTable("example", {
    defaultRouteTableId: aws_vpc.example.default_route_table_id,
    routes: [
        {
            cidrBlock: "10.0.1.0/24",
            gatewayId: aws_internet_gateway.example.id,
        },
        {
            ipv6CidrBlock: "::/0",
            egressOnlyGatewayId: aws_egress_only_internet_gateway.example.id,
        },
    ],
    tags: {
        Name: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.DefaultRouteTable("example",
    default_route_table_id=aws_vpc["example"]["default_route_table_id"],
    routes=[
        aws.ec2.DefaultRouteTableRouteArgs(
            cidr_block="10.0.1.0/24",
            gateway_id=aws_internet_gateway["example"]["id"],
        ),
        aws.ec2.DefaultRouteTableRouteArgs(
            ipv6_cidr_block="::/0",
            egress_only_gateway_id=aws_egress_only_internet_gateway["example"]["id"],
        ),
    ],
    tags={
        "Name": "example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.DefaultRouteTable("example", new()
    {
        DefaultRouteTableId = aws_vpc.Example.Default_route_table_id,
        Routes = new[]
        {
            new Aws.Ec2.Inputs.DefaultRouteTableRouteArgs
            {
                CidrBlock = "10.0.1.0/24",
                GatewayId = aws_internet_gateway.Example.Id,
            },
            new Aws.Ec2.Inputs.DefaultRouteTableRouteArgs
            {
                Ipv6CidrBlock = "::/0",
                EgressOnlyGatewayId = aws_egress_only_internet_gateway.Example.Id,
            },
        },
        Tags = 
        {
            { "Name", "example" },
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
		_, err := ec2.NewDefaultRouteTable(ctx, "example", &ec2.DefaultRouteTableArgs{
			DefaultRouteTableId: pulumi.Any(aws_vpc.Example.Default_route_table_id),
			Routes: ec2.DefaultRouteTableRouteArray{
				&ec2.DefaultRouteTableRouteArgs{
					CidrBlock: pulumi.String("10.0.1.0/24"),
					GatewayId: pulumi.Any(aws_internet_gateway.Example.Id),
				},
				&ec2.DefaultRouteTableRouteArgs{
					Ipv6CidrBlock:       pulumi.String("::/0"),
					EgressOnlyGatewayId: pulumi.Any(aws_egress_only_internet_gateway.Example.Id),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
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
import com.pulumi.aws.ec2.DefaultRouteTable;
import com.pulumi.aws.ec2.DefaultRouteTableArgs;
import com.pulumi.aws.ec2.inputs.DefaultRouteTableRouteArgs;
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
        var example = new DefaultRouteTable("example", DefaultRouteTableArgs.builder()        
            .defaultRouteTableId(aws_vpc.example().default_route_table_id())
            .routes(            
                DefaultRouteTableRouteArgs.builder()
                    .cidrBlock("10.0.1.0/24")
                    .gatewayId(aws_internet_gateway.example().id())
                    .build(),
                DefaultRouteTableRouteArgs.builder()
                    .ipv6CidrBlock("::/0")
                    .egressOnlyGatewayId(aws_egress_only_internet_gateway.example().id())
                    .build())
            .tags(Map.of("Name", "example"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:DefaultRouteTable
    properties:
      defaultRouteTableId: ${aws_vpc.example.default_route_table_id}
      routes:
        - cidrBlock: 10.0.1.0/24
          gatewayId: ${aws_internet_gateway.example.id}
        - ipv6CidrBlock: ::/0
          egressOnlyGatewayId: ${aws_egress_only_internet_gateway.example.id}
      tags:
        Name: example
```

To subsequently remove all managed routes:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.DefaultRouteTable("example", {
    defaultRouteTableId: aws_vpc.example.default_route_table_id,
    routes: [],
    tags: {
        Name: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.DefaultRouteTable("example",
    default_route_table_id=aws_vpc["example"]["default_route_table_id"],
    routes=[],
    tags={
        "Name": "example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.DefaultRouteTable("example", new()
    {
        DefaultRouteTableId = aws_vpc.Example.Default_route_table_id,
        Routes = new[] {},
        Tags = 
        {
            { "Name", "example" },
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
		_, err := ec2.NewDefaultRouteTable(ctx, "example", &ec2.DefaultRouteTableArgs{
			DefaultRouteTableId: pulumi.Any(aws_vpc.Example.Default_route_table_id),
			Routes:              ec2.DefaultRouteTableRouteArray{},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
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
import com.pulumi.aws.ec2.DefaultRouteTable;
import com.pulumi.aws.ec2.DefaultRouteTableArgs;
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
        var example = new DefaultRouteTable("example", DefaultRouteTableArgs.builder()        
            .defaultRouteTableId(aws_vpc.example().default_route_table_id())
            .routes()
            .tags(Map.of("Name", "example"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:DefaultRouteTable
    properties:
      defaultRouteTableId: ${aws_vpc.example.default_route_table_id}
      routes: []
      tags:
        Name: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Default VPC route tables using the `vpc_id`. For example:

```sh
 $ pulumi import aws:ec2/defaultRouteTable:DefaultRouteTable example vpc-33cc44dd
```
 