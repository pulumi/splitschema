Provides an ElastiCache parameter group resource.

> **NOTE:** Attempting to remove the `reserved-memory` parameter when `family` is set to `redis2.6` or `redis2.8` may show a perpetual difference in this provider due to an ElastiCache API limitation. Leave that parameter configured with any value to workaround the issue.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.elasticache.ParameterGroup("default", {
    family: "redis2.8",
    parameters: [
        {
            name: "activerehashing",
            value: "yes",
        },
        {
            name: "min-slaves-to-write",
            value: "2",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.elasticache.ParameterGroup("default",
    family="redis2.8",
    parameters=[
        aws.elasticache.ParameterGroupParameterArgs(
            name="activerehashing",
            value="yes",
        ),
        aws.elasticache.ParameterGroupParameterArgs(
            name="min-slaves-to-write",
            value="2",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.ElastiCache.ParameterGroup("default", new()
    {
        Family = "redis2.8",
        Parameters = new[]
        {
            new Aws.ElastiCache.Inputs.ParameterGroupParameterArgs
            {
                Name = "activerehashing",
                Value = "yes",
            },
            new Aws.ElastiCache.Inputs.ParameterGroupParameterArgs
            {
                Name = "min-slaves-to-write",
                Value = "2",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewParameterGroup(ctx, "default", &elasticache.ParameterGroupArgs{
			Family: pulumi.String("redis2.8"),
			Parameters: elasticache.ParameterGroupParameterArray{
				&elasticache.ParameterGroupParameterArgs{
					Name:  pulumi.String("activerehashing"),
					Value: pulumi.String("yes"),
				},
				&elasticache.ParameterGroupParameterArgs{
					Name:  pulumi.String("min-slaves-to-write"),
					Value: pulumi.String("2"),
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
import com.pulumi.aws.elasticache.ParameterGroup;
import com.pulumi.aws.elasticache.ParameterGroupArgs;
import com.pulumi.aws.elasticache.inputs.ParameterGroupParameterArgs;
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
        var default_ = new ParameterGroup("default", ParameterGroupArgs.builder()        
            .family("redis2.8")
            .parameters(            
                ParameterGroupParameterArgs.builder()
                    .name("activerehashing")
                    .value("yes")
                    .build(),
                ParameterGroupParameterArgs.builder()
                    .name("min-slaves-to-write")
                    .value("2")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:elasticache:ParameterGroup
    properties:
      family: redis2.8
      parameters:
        - name: activerehashing
          value: yes
        - name: min-slaves-to-write
          value: '2'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache Parameter Groups using the `name`. For example:

```sh
 $ pulumi import aws:elasticache/parameterGroup:ParameterGroup default redis-params
```
 