Provides information for multiple EC2 Transit Gateway Route Table Associations, such as their identifiers.

{{% examples %}}
## Example Usage
{{% example %}}
### By Transit Gateway Identifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getRouteTableAssociations({
    transitGatewayRouteTableId: aws_ec2_transit_gateway_route_table.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_route_table_associations(transit_gateway_route_table_id=aws_ec2_transit_gateway_route_table["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetRouteTableAssociations.Invoke(new()
    {
        TransitGatewayRouteTableId = aws_ec2_transit_gateway_route_table.Example.Id,
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
		_, err := ec2transitgateway.GetRouteTableAssociations(ctx, &ec2transitgateway.GetRouteTableAssociationsArgs{
			TransitGatewayRouteTableId: aws_ec2_transit_gateway_route_table.Example.Id,
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
import com.pulumi.aws.ec2transitgateway.inputs.GetRouteTableAssociationsArgs;
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
        final var example = Ec2transitgatewayFunctions.getRouteTableAssociations(GetRouteTableAssociationsArgs.builder()
            .transitGatewayRouteTableId(aws_ec2_transit_gateway_route_table.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getRouteTableAssociations
      Arguments:
        transitGatewayRouteTableId: ${aws_ec2_transit_gateway_route_table.example.id}
```
{{% /example %}}
{{% /examples %}}