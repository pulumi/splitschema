Provides a resource for managing the main routing table of a VPC.

> **NOTE:** **Do not** use both `aws.ec2.DefaultRouteTable` to manage a default route table **and** `aws.ec2.MainRouteTableAssociation` with the same VPC due to possible route conflicts. See aws.ec2.DefaultRouteTable documentation for more details.
For more information, see the Amazon VPC User Guide on [Route Tables][aws-route-tables]. For information about managing normal route tables in Pulumi, see [`aws.ec2.RouteTable`][tf-route-tables].

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainRouteTableAssociation = new aws.ec2.MainRouteTableAssociation("mainRouteTableAssociation", {
    vpcId: aws_vpc.foo.id,
    routeTableId: aws_route_table.bar.id,
});
```
```python
import pulumi
import pulumi_aws as aws

main_route_table_association = aws.ec2.MainRouteTableAssociation("mainRouteTableAssociation",
    vpc_id=aws_vpc["foo"]["id"],
    route_table_id=aws_route_table["bar"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mainRouteTableAssociation = new Aws.Ec2.MainRouteTableAssociation("mainRouteTableAssociation", new()
    {
        VpcId = aws_vpc.Foo.Id,
        RouteTableId = aws_route_table.Bar.Id,
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
		_, err := ec2.NewMainRouteTableAssociation(ctx, "mainRouteTableAssociation", &ec2.MainRouteTableAssociationArgs{
			VpcId:        pulumi.Any(aws_vpc.Foo.Id),
			RouteTableId: pulumi.Any(aws_route_table.Bar.Id),
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
import com.pulumi.aws.ec2.MainRouteTableAssociation;
import com.pulumi.aws.ec2.MainRouteTableAssociationArgs;
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
        var mainRouteTableAssociation = new MainRouteTableAssociation("mainRouteTableAssociation", MainRouteTableAssociationArgs.builder()        
            .vpcId(aws_vpc.foo().id())
            .routeTableId(aws_route_table.bar().id())
            .build());

    }
}
```
```yaml
resources:
  mainRouteTableAssociation:
    type: aws:ec2:MainRouteTableAssociation
    properties:
      vpcId: ${aws_vpc.foo.id}
      routeTableId: ${aws_route_table.bar.id}
```
{{% /example %}}
{{% /examples %}}
## Notes

On VPC creation, the AWS API always creates an initial Main Route Table. This
resource records the ID of that Route Table under `original_route_table_id`.
The "Delete" action for a `main_route_table_association` consists of resetting
this original table as the Main Route Table for the VPC. You'll see this
additional Route Table in the AWS console; it must remain intact in order for
the `main_route_table_association` delete to work properly.
