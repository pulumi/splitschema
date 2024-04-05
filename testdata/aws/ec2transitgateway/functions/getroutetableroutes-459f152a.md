Provides informations for routes of a specific transit gateway, such as state, type, cidr

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ec2transitgateway.getRouteTableRoutes({
    filters: [{
        name: "type",
        values: ["propagated"],
    }],
    transitGatewayRouteTableId: aws_ec2_transit_gateway_route_table.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2transitgateway.get_route_table_routes(filters=[aws.ec2transitgateway.GetRouteTableRoutesFilterArgs(
        name="type",
        values=["propagated"],
    )],
    transit_gateway_route_table_id=aws_ec2_transit_gateway_route_table["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ec2TransitGateway.GetRouteTableRoutes.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2TransitGateway.Inputs.GetRouteTableRoutesFilterInputArgs
            {
                Name = "type",
                Values = new[]
                {
                    "propagated",
                },
            },
        },
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
		_, err := ec2transitgateway.GetRouteTableRoutes(ctx, &ec2transitgateway.GetRouteTableRoutesArgs{
			Filters: []ec2transitgateway.GetRouteTableRoutesFilter{
				{
					Name: "type",
					Values: []string{
						"propagated",
					},
				},
			},
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
import com.pulumi.aws.ec2transitgateway.inputs.GetRouteTableRoutesArgs;
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
        final var test = Ec2transitgatewayFunctions.getRouteTableRoutes(GetRouteTableRoutesArgs.builder()
            .filters(GetRouteTableRoutesFilterArgs.builder()
                .name("type")
                .values("propagated")
                .build())
            .transitGatewayRouteTableId(aws_ec2_transit_gateway_route_table.example().id())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ec2transitgateway:getRouteTableRoutes
      Arguments:
        filters:
          - name: type
            values:
              - propagated
        transitGatewayRouteTableId: ${aws_ec2_transit_gateway_route_table.example.id}
```
{{% /example %}}
{{% /examples %}}