Provides a list of Elastic IPs in a region.

{{% examples %}}
## Example Usage
{{% example %}}

The following shows outputting all Elastic IPs with the a specific tag value.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getEips({
    tags: {
        Env: "dev",
    },
});
export const allocationIds = example.then(example => example.allocationIds);
export const publicIps = example.then(example => example.publicIps);
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_eips(tags={
    "Env": "dev",
})
pulumi.export("allocationIds", example.allocation_ids)
pulumi.export("publicIps", example.public_ips)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetEips.Invoke(new()
    {
        Tags = 
        {
            { "Env", "dev" },
        },
    });

    return new Dictionary<string, object?>
    {
        ["allocationIds"] = example.Apply(getEipsResult => getEipsResult.AllocationIds),
        ["publicIps"] = example.Apply(getEipsResult => getEipsResult.PublicIps),
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
		example, err := ec2.GetEips(ctx, &ec2.GetEipsArgs{
			Tags: map[string]interface{}{
				"Env": "dev",
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("allocationIds", example.AllocationIds)
		ctx.Export("publicIps", example.PublicIps)
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
import com.pulumi.aws.ec2.inputs.GetEipsArgs;
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
        final var example = Ec2Functions.getEips(GetEipsArgs.builder()
            .tags(Map.of("Env", "dev"))
            .build());

        ctx.export("allocationIds", example.applyValue(getEipsResult -> getEipsResult.allocationIds()));
        ctx.export("publicIps", example.applyValue(getEipsResult -> getEipsResult.publicIps()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getEips
      Arguments:
        tags:
          Env: dev
outputs:
  # VPC EIPs.
  allocationIds: ${example.allocationIds}
  # EC2-Classic EIPs.
  publicIps: ${example.publicIps}
```
{{% /example %}}
{{% /examples %}}