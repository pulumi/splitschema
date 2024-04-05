Provides information for multiple EC2 Local Gateways, such as their identifiers.

{{% examples %}}
## Example Usage
{{% example %}}

The following example retrieves Local Gateways with a resource tag of `service` set to `production`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooLocalGateways = aws.ec2.getLocalGateways({
    tags: {
        service: "production",
    },
});
export const foo = fooLocalGateways.then(fooLocalGateways => fooLocalGateways.ids);
```
```python
import pulumi
import pulumi_aws as aws

foo_local_gateways = aws.ec2.get_local_gateways(tags={
    "service": "production",
})
pulumi.export("foo", foo_local_gateways.ids)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooLocalGateways = Aws.Ec2.GetLocalGateways.Invoke(new()
    {
        Tags = 
        {
            { "service", "production" },
        },
    });

    return new Dictionary<string, object?>
    {
        ["foo"] = fooLocalGateways.Apply(getLocalGatewaysResult => getLocalGatewaysResult.Ids),
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
		fooLocalGateways, err := ec2.GetLocalGateways(ctx, &ec2.GetLocalGatewaysArgs{
			Tags: map[string]interface{}{
				"service": "production",
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("foo", fooLocalGateways.Ids)
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
import com.pulumi.aws.ec2.inputs.GetLocalGatewaysArgs;
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
        final var fooLocalGateways = Ec2Functions.getLocalGateways(GetLocalGatewaysArgs.builder()
            .tags(Map.of("service", "production"))
            .build());

        ctx.export("foo", fooLocalGateways.applyValue(getLocalGatewaysResult -> getLocalGatewaysResult.ids()));
    }
}
```
```yaml
variables:
  fooLocalGateways:
    fn::invoke:
      Function: aws:ec2:getLocalGateways
      Arguments:
        tags:
          service: production
outputs:
  foo: ${fooLocalGateways.ids}
```
{{% /example %}}
{{% /examples %}}