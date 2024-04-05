Provides a MemoryDB Parameter Group.

More information about parameter groups can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/parametergroups.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.memorydb.ParameterGroup("example", {
    family: "memorydb_redis6",
    parameters: [{
        name: "activedefrag",
        value: "yes",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.ParameterGroup("example",
    family="memorydb_redis6",
    parameters=[aws.memorydb.ParameterGroupParameterArgs(
        name="activedefrag",
        value="yes",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.MemoryDb.ParameterGroup("example", new()
    {
        Family = "memorydb_redis6",
        Parameters = new[]
        {
            new Aws.MemoryDb.Inputs.ParameterGroupParameterArgs
            {
                Name = "activedefrag",
                Value = "yes",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/memorydb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := memorydb.NewParameterGroup(ctx, "example", &memorydb.ParameterGroupArgs{
			Family: pulumi.String("memorydb_redis6"),
			Parameters: memorydb.ParameterGroupParameterArray{
				&memorydb.ParameterGroupParameterArgs{
					Name:  pulumi.String("activedefrag"),
					Value: pulumi.String("yes"),
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
import com.pulumi.aws.memorydb.ParameterGroup;
import com.pulumi.aws.memorydb.ParameterGroupArgs;
import com.pulumi.aws.memorydb.inputs.ParameterGroupParameterArgs;
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
        var example = new ParameterGroup("example", ParameterGroupArgs.builder()        
            .family("memorydb_redis6")
            .parameters(ParameterGroupParameterArgs.builder()
                .name("activedefrag")
                .value("yes")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:memorydb:ParameterGroup
    properties:
      family: memorydb_redis6
      parameters:
        - name: activedefrag
          value: yes
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a parameter group using the `name`. For example:

```sh
 $ pulumi import aws:memorydb/parameterGroup:ParameterGroup example my-parameter-group
```
 