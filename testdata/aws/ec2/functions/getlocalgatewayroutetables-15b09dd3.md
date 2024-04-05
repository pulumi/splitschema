Provides information for multiple EC2 Local Gateway Route Tables, such as their identifiers.

{{% examples %}}
## Example Usage
{{% example %}}

The following shows outputting all Local Gateway Route Table Ids.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooLocalGatewayRouteTables = aws.ec2.getLocalGatewayRouteTables({});
export const foo = fooLocalGatewayRouteTables.then(fooLocalGatewayRouteTables => fooLocalGatewayRouteTables.ids);
```
```python
import pulumi
import pulumi_aws as aws

foo_local_gateway_route_tables = aws.ec2.get_local_gateway_route_tables()
pulumi.export("foo", foo_local_gateway_route_tables.ids)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooLocalGatewayRouteTables = Aws.Ec2.GetLocalGatewayRouteTables.Invoke();

    return new Dictionary<string, object?>
    {
        ["foo"] = fooLocalGatewayRouteTables.Apply(getLocalGatewayRouteTablesResult => getLocalGatewayRouteTablesResult.Ids),
    };
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
		fooLocalGatewayRouteTables, err := ec2.GetLocalGatewayRouteTables(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("foo", fooLocalGatewayRouteTables.Ids)
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
import com.pulumi.aws.ec2.inputs.GetLocalGatewayRouteTablesArgs;
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
        final var fooLocalGatewayRouteTables = Ec2Functions.getLocalGatewayRouteTables();

        ctx.export("foo", fooLocalGatewayRouteTables.applyValue(getLocalGatewayRouteTablesResult -> getLocalGatewayRouteTablesResult.ids()));
    }
}
```
```yaml
variables:
  fooLocalGatewayRouteTables:
    fn::invoke:
      Function: aws:ec2:getLocalGatewayRouteTables
      Arguments: {}
outputs:
  foo: ${fooLocalGatewayRouteTables.ids}
```
{{% /example %}}
{{% /examples %}}