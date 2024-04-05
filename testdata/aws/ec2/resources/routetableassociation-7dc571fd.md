Provides a resource to create an association between a route table and a subnet or a route table and an
internet gateway or virtual private gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const routeTableAssociation = new aws.ec2.RouteTableAssociation("routeTableAssociation", {
    subnetId: aws_subnet.foo.id,
    routeTableId: aws_route_table.bar.id,
});
```
```python
import pulumi
import pulumi_aws as aws

route_table_association = aws.ec2.RouteTableAssociation("routeTableAssociation",
    subnet_id=aws_subnet["foo"]["id"],
    route_table_id=aws_route_table["bar"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var routeTableAssociation = new Aws.Ec2.RouteTableAssociation("routeTableAssociation", new()
    {
        SubnetId = aws_subnet.Foo.Id,
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
		_, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation", &ec2.RouteTableAssociationArgs{
			SubnetId:     pulumi.Any(aws_subnet.Foo.Id),
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
import com.pulumi.aws.ec2.RouteTableAssociation;
import com.pulumi.aws.ec2.RouteTableAssociationArgs;
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
        var routeTableAssociation = new RouteTableAssociation("routeTableAssociation", RouteTableAssociationArgs.builder()        
            .subnetId(aws_subnet.foo().id())
            .routeTableId(aws_route_table.bar().id())
            .build());

    }
}
```
```yaml
resources:
  routeTableAssociation:
    type: aws:ec2:RouteTableAssociation
    properties:
      subnetId: ${aws_subnet.foo.id}
      routeTableId: ${aws_route_table.bar.id}
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const routeTableAssociation = new aws.ec2.RouteTableAssociation("routeTableAssociation", {
    gatewayId: aws_internet_gateway.foo.id,
    routeTableId: aws_route_table.bar.id,
});
```
```python
import pulumi
import pulumi_aws as aws

route_table_association = aws.ec2.RouteTableAssociation("routeTableAssociation",
    gateway_id=aws_internet_gateway["foo"]["id"],
    route_table_id=aws_route_table["bar"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var routeTableAssociation = new Aws.Ec2.RouteTableAssociation("routeTableAssociation", new()
    {
        GatewayId = aws_internet_gateway.Foo.Id,
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
		_, err := ec2.NewRouteTableAssociation(ctx, "routeTableAssociation", &ec2.RouteTableAssociationArgs{
			GatewayId:    pulumi.Any(aws_internet_gateway.Foo.Id),
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
import com.pulumi.aws.ec2.RouteTableAssociation;
import com.pulumi.aws.ec2.RouteTableAssociationArgs;
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
        var routeTableAssociation = new RouteTableAssociation("routeTableAssociation", RouteTableAssociationArgs.builder()        
            .gatewayId(aws_internet_gateway.foo().id())
            .routeTableId(aws_route_table.bar().id())
            .build());

    }
}
```
```yaml
resources:
  routeTableAssociation:
    type: aws:ec2:RouteTableAssociation
    properties:
      gatewayId: ${aws_internet_gateway.foo.id}
      routeTableId: ${aws_route_table.bar.id}
```
{{% /example %}}
{{% /examples %}}

## Import

With EC2 Internet Gateways:

__Using `pulumi import` to import__ EC2 Route Table Associations using the associated resource ID and Route Table ID separated by a forward slash (`/`). For example:

With EC2 Subnets:

```sh
 $ pulumi import aws:ec2/routeTableAssociation:RouteTableAssociation assoc subnet-6777656e646f6c796e/rtb-656c65616e6f72
```
 With EC2 Internet Gateways:

```sh
 $ pulumi import aws:ec2/routeTableAssociation:RouteTableAssociation assoc igw-01b3a60780f8d034a/rtb-656c65616e6f72
```
 