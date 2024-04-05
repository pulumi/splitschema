Manages an EC2 Transit Gateway Prefix List Reference.

{{% examples %}}
## Example Usage
{{% example %}}
### Attachment Routing

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.PrefixListReference("example", {
    prefixListId: aws_ec2_managed_prefix_list.example.id,
    transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment.example.id,
    transitGatewayRouteTableId: aws_ec2_transit_gateway.example.association_default_route_table_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.PrefixListReference("example",
    prefix_list_id=aws_ec2_managed_prefix_list["example"]["id"],
    transit_gateway_attachment_id=aws_ec2_transit_gateway_vpc_attachment["example"]["id"],
    transit_gateway_route_table_id=aws_ec2_transit_gateway["example"]["association_default_route_table_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.PrefixListReference("example", new()
    {
        PrefixListId = aws_ec2_managed_prefix_list.Example.Id,
        TransitGatewayAttachmentId = aws_ec2_transit_gateway_vpc_attachment.Example.Id,
        TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
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
		_, err := ec2transitgateway.NewPrefixListReference(ctx, "example", &ec2transitgateway.PrefixListReferenceArgs{
			PrefixListId:               pulumi.Any(aws_ec2_managed_prefix_list.Example.Id),
			TransitGatewayAttachmentId: pulumi.Any(aws_ec2_transit_gateway_vpc_attachment.Example.Id),
			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
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
import com.pulumi.aws.ec2transitgateway.PrefixListReference;
import com.pulumi.aws.ec2transitgateway.PrefixListReferenceArgs;
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
        var example = new PrefixListReference("example", PrefixListReferenceArgs.builder()        
            .prefixListId(aws_ec2_managed_prefix_list.example().id())
            .transitGatewayAttachmentId(aws_ec2_transit_gateway_vpc_attachment.example().id())
            .transitGatewayRouteTableId(aws_ec2_transit_gateway.example().association_default_route_table_id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:PrefixListReference
    properties:
      prefixListId: ${aws_ec2_managed_prefix_list.example.id}
      transitGatewayAttachmentId: ${aws_ec2_transit_gateway_vpc_attachment.example.id}
      transitGatewayRouteTableId: ${aws_ec2_transit_gateway.example.association_default_route_table_id}
```
{{% /example %}}
{{% example %}}
### Blackhole Routing

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2transitgateway.PrefixListReference("example", {
    blackhole: true,
    prefixListId: aws_ec2_managed_prefix_list.example.id,
    transitGatewayRouteTableId: aws_ec2_transit_gateway.example.association_default_route_table_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.PrefixListReference("example",
    blackhole=True,
    prefix_list_id=aws_ec2_managed_prefix_list["example"]["id"],
    transit_gateway_route_table_id=aws_ec2_transit_gateway["example"]["association_default_route_table_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2TransitGateway.PrefixListReference("example", new()
    {
        Blackhole = true,
        PrefixListId = aws_ec2_managed_prefix_list.Example.Id,
        TransitGatewayRouteTableId = aws_ec2_transit_gateway.Example.Association_default_route_table_id,
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
		_, err := ec2transitgateway.NewPrefixListReference(ctx, "example", &ec2transitgateway.PrefixListReferenceArgs{
			Blackhole:                  pulumi.Bool(true),
			PrefixListId:               pulumi.Any(aws_ec2_managed_prefix_list.Example.Id),
			TransitGatewayRouteTableId: pulumi.Any(aws_ec2_transit_gateway.Example.Association_default_route_table_id),
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
import com.pulumi.aws.ec2transitgateway.PrefixListReference;
import com.pulumi.aws.ec2transitgateway.PrefixListReferenceArgs;
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
        var example = new PrefixListReference("example", PrefixListReferenceArgs.builder()        
            .blackhole(true)
            .prefixListId(aws_ec2_managed_prefix_list.example().id())
            .transitGatewayRouteTableId(aws_ec2_transit_gateway.example().association_default_route_table_id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2transitgateway:PrefixListReference
    properties:
      blackhole: true
      prefixListId: ${aws_ec2_managed_prefix_list.example.id}
      transitGatewayRouteTableId: ${aws_ec2_transit_gateway.example.association_default_route_table_id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ec2_transit_gateway_prefix_list_reference` using the EC2 Transit Gateway Route Table identifier and EC2 Prefix List identifier, separated by an underscore (`_`). For example:

```sh
 $ pulumi import aws:ec2transitgateway/prefixListReference:PrefixListReference example tgw-rtb-12345678_pl-12345678
```
 