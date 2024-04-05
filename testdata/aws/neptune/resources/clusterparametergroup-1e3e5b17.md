Manages a Neptune Cluster Parameter Group

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.neptune.ClusterParameterGroup("example", {
    description: "neptune cluster parameter group",
    family: "neptune1",
    parameters: [{
        name: "neptune_enable_audit_log",
        value: "1",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.neptune.ClusterParameterGroup("example",
    description="neptune cluster parameter group",
    family="neptune1",
    parameters=[aws.neptune.ClusterParameterGroupParameterArgs(
        name="neptune_enable_audit_log",
        value="1",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Neptune.ClusterParameterGroup("example", new()
    {
        Description = "neptune cluster parameter group",
        Family = "neptune1",
        Parameters = new[]
        {
            new Aws.Neptune.Inputs.ClusterParameterGroupParameterArgs
            {
                Name = "neptune_enable_audit_log",
                Value = "1",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/neptune"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := neptune.NewClusterParameterGroup(ctx, "example", &neptune.ClusterParameterGroupArgs{
			Description: pulumi.String("neptune cluster parameter group"),
			Family:      pulumi.String("neptune1"),
			Parameters: neptune.ClusterParameterGroupParameterArray{
				&neptune.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("neptune_enable_audit_log"),
					Value: pulumi.String("1"),
				},
			},
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
import com.pulumi.aws.neptune.ClusterParameterGroup;
import com.pulumi.aws.neptune.ClusterParameterGroupArgs;
import com.pulumi.aws.neptune.inputs.ClusterParameterGroupParameterArgs;
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
        var example = new ClusterParameterGroup("example", ClusterParameterGroupArgs.builder()        
            .description("neptune cluster parameter group")
            .family("neptune1")
            .parameters(ClusterParameterGroupParameterArgs.builder()
                .name("neptune_enable_audit_log")
                .value(1)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:neptune:ClusterParameterGroup
    properties:
      description: neptune cluster parameter group
      family: neptune1
      parameters:
        - name: neptune_enable_audit_log
          value: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Neptune Cluster Parameter Groups using the `name`. For example:

```sh
 $ pulumi import aws:neptune/clusterParameterGroup:ClusterParameterGroup cluster_pg production-pg-1
```
 