Use this data source to get the window IDs of SSM maintenance windows.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ssm.getMaintenanceWindows({
    filters: [{
        name: "Enabled",
        values: ["true"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.get_maintenance_windows(filters=[aws.ssm.GetMaintenanceWindowsFilterArgs(
    name="Enabled",
    values=["true"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ssm.GetMaintenanceWindows.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ssm.Inputs.GetMaintenanceWindowsFilterInputArgs
            {
                Name = "Enabled",
                Values = new[]
                {
                    "true",
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
		_, err := ssm.GetMaintenanceWindows(ctx, &ssm.GetMaintenanceWindowsArgs{
			Filters: []ssm.GetMaintenanceWindowsFilter{
				{
					Name: "Enabled",
					Values: []string{
						"true",
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
import com.pulumi.aws.ssm.inputs.GetMaintenanceWindowsArgs;
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
        final var example = SsmFunctions.getMaintenanceWindows(GetMaintenanceWindowsArgs.builder()
            .filters(GetMaintenanceWindowsFilterArgs.builder()
                .name("Enabled")
                .values("true")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ssm:getMaintenanceWindows
      Arguments:
        filters:
          - name: Enabled
            values:
              - 'true'
```
{{% /example %}}
{{% /examples %}}