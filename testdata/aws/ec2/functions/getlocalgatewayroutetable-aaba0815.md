Provides details about an EC2 Local Gateway Route Table.

This data source can prove useful when a module accepts a local gateway route table id as
an input variable and needs to, for example, find the associated Outpost or Local Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

The following example returns a specific local gateway route table ID

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const awsEc2LocalGatewayRouteTable = config.requireObject("awsEc2LocalGatewayRouteTable");
const selected = aws.ec2.getLocalGatewayRouteTable({
    localGatewayRouteTableId: awsEc2LocalGatewayRouteTable,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
aws_ec2_local_gateway_route_table = config.require_object("awsEc2LocalGatewayRouteTable")
selected = aws.ec2.get_local_gateway_route_table(local_gateway_route_table_id=aws_ec2_local_gateway_route_table)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var awsEc2LocalGatewayRouteTable = config.RequireObject<dynamic>("awsEc2LocalGatewayRouteTable");
    var selected = Aws.Ec2.GetLocalGatewayRouteTable.Invoke(new()
    {
        LocalGatewayRouteTableId = awsEc2LocalGatewayRouteTable,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		awsEc2LocalGatewayRouteTable := cfg.RequireObject("awsEc2LocalGatewayRouteTable")
		_, err := ec2.GetLocalGatewayRouteTable(ctx, &ec2.GetLocalGatewayRouteTableArgs{
			LocalGatewayRouteTableId: pulumi.StringRef(awsEc2LocalGatewayRouteTable),
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
import com.pulumi.aws.ec2.inputs.GetLocalGatewayRouteTableArgs;
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
        final var config = ctx.config();
        final var awsEc2LocalGatewayRouteTable = config.get("awsEc2LocalGatewayRouteTable");
        final var selected = Ec2Functions.getLocalGatewayRouteTable(GetLocalGatewayRouteTableArgs.builder()
            .localGatewayRouteTableId(awsEc2LocalGatewayRouteTable)
            .build());

    }
}
```
```yaml
configuration:
  awsEc2LocalGatewayRouteTable:
    type: dynamic
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getLocalGatewayRouteTable
      Arguments:
        localGatewayRouteTableId: ${awsEc2LocalGatewayRouteTable}
```
{{% /example %}}
{{% /examples %}}