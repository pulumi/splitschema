Manages a VPC Endpoint Route Table Association

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpcEndpointRouteTableAssociation("example", {
    routeTableId: aws_route_table.example.id,
    vpcEndpointId: aws_vpc_endpoint.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.VpcEndpointRouteTableAssociation("example",
    route_table_id=aws_route_table["example"]["id"],
    vpc_endpoint_id=aws_vpc_endpoint["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.VpcEndpointRouteTableAssociation("example", new()
    {
        RouteTableId = aws_route_table.Example.Id,
        VpcEndpointId = aws_vpc_endpoint.Example.Id,
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
		_, err := ec2.NewVpcEndpointRouteTableAssociation(ctx, "example", &ec2.VpcEndpointRouteTableAssociationArgs{
			RouteTableId:  pulumi.Any(aws_route_table.Example.Id),
			VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Example.Id),
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
import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociation;
import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociationArgs;
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
        var example = new VpcEndpointRouteTableAssociation("example", VpcEndpointRouteTableAssociationArgs.builder()        
            .routeTableId(aws_route_table.example().id())
            .vpcEndpointId(aws_vpc_endpoint.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpcEndpointRouteTableAssociation
    properties:
      routeTableId: ${aws_route_table.example.id}
      vpcEndpointId: ${aws_vpc_endpoint.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoint Route Table Associations using `vpc_endpoint_id` together with `route_table_id`. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation example vpce-aaaaaaaa/rtb-bbbbbbbb
```
 