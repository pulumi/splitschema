Use this data source to get the instance IDs of SSM managed instances.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ssm.getInstances({
    filters: [{
        name: "PlatformTypes",
        values: ["Linux"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.get_instances(filters=[aws.ssm.GetInstancesFilterArgs(
    name="PlatformTypes",
    values=["Linux"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ssm.GetInstances.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ssm.Inputs.GetInstancesFilterInputArgs
            {
                Name = "PlatformTypes",
                Values = new[]
                {
                    "Linux",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.GetInstances(ctx, &ssm.GetInstancesArgs{
			Filters: []ssm.GetInstancesFilter{
				{
					Name: "PlatformTypes",
					Values: []string{
						"Linux",
					},
				},
			},
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
import com.pulumi.aws.ssm.SsmFunctions;
import com.pulumi.aws.ssm.inputs.GetInstancesArgs;
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
        final var example = SsmFunctions.getInstances(GetInstancesArgs.builder()
            .filters(GetInstancesFilterArgs.builder()
                .name("PlatformTypes")
                .values("Linux")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ssm:getInstances
      Arguments:
        filters:
          - name: PlatformTypes
            values:
              - Linux
```
{{% /example %}}
{{% /examples %}}