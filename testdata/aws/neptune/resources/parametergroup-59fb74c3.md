Manages a Neptune Parameter Group

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.neptune.ParameterGroup("example", {
    family: "neptune1",
    parameters: [{
        name: "neptune_query_timeout",
        value: "25",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.neptune.ParameterGroup("example",
    family="neptune1",
    parameters=[aws.neptune.ParameterGroupParameterArgs(
        name="neptune_query_timeout",
        value="25",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Neptune.ParameterGroup("example", new()
    {
        Family = "neptune1",
        Parameters = new[]
        {
            new Aws.Neptune.Inputs.ParameterGroupParameterArgs
            {
                Name = "neptune_query_timeout",
                Value = "25",
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
		_, err := neptune.NewParameterGroup(ctx, "example", &neptune.ParameterGroupArgs{
			Family: pulumi.String("neptune1"),
			Parameters: neptune.ParameterGroupParameterArray{
				&neptune.ParameterGroupParameterArgs{
					Name:  pulumi.String("neptune_query_timeout"),
					Value: pulumi.String("25"),
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
import com.pulumi.aws.neptune.ParameterGroup;
import com.pulumi.aws.neptune.ParameterGroupArgs;
import com.pulumi.aws.neptune.inputs.ParameterGroupParameterArgs;
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
            .family("neptune1")
            .parameters(ParameterGroupParameterArgs.builder()
                .name("neptune_query_timeout")
                .value("25")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:neptune:ParameterGroup
    properties:
      family: neptune1
      parameters:
        - name: neptune_query_timeout
          value: '25'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Neptune Parameter Groups using the `name`. For example:

```sh
 $ pulumi import aws:neptune/parameterGroup:ParameterGroup some_pg some-pg
```
 