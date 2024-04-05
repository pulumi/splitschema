Provides an RDS DB cluster parameter group resource. Documentation of the available parameters for various Aurora engines can be found at:

* [Aurora MySQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Reference.html)
* [Aurora PostgreSQL Parameters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraPostgreSQL.Reference.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.ClusterParameterGroup("default", {
    description: "RDS default cluster parameter group",
    family: "aurora5.6",
    parameters: [
        {
            name: "character_set_server",
            value: "utf8",
        },
        {
            name: "character_set_client",
            value: "utf8",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.ClusterParameterGroup("default",
    description="RDS default cluster parameter group",
    family="aurora5.6",
    parameters=[
        aws.rds.ClusterParameterGroupParameterArgs(
            name="character_set_server",
            value="utf8",
        ),
        aws.rds.ClusterParameterGroupParameterArgs(
            name="character_set_client",
            value="utf8",
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
    var @default = new Aws.Rds.ClusterParameterGroup("default", new()
    {
        Description = "RDS default cluster parameter group",
        Family = "aurora5.6",
        Parameters = new[]
        {
            new Aws.Rds.Inputs.ClusterParameterGroupParameterArgs
            {
                Name = "character_set_server",
                Value = "utf8",
            },
            new Aws.Rds.Inputs.ClusterParameterGroupParameterArgs
            {
                Name = "character_set_client",
                Value = "utf8",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewClusterParameterGroup(ctx, "default", &rds.ClusterParameterGroupArgs{
			Description: pulumi.String("RDS default cluster parameter group"),
			Family:      pulumi.String("aurora5.6"),
			Parameters: rds.ClusterParameterGroupParameterArray{
				&rds.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_server"),
					Value: pulumi.String("utf8"),
				},
				&rds.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("character_set_client"),
					Value: pulumi.String("utf8"),
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
import com.pulumi.aws.rds.ClusterParameterGroup;
import com.pulumi.aws.rds.ClusterParameterGroupArgs;
import com.pulumi.aws.rds.inputs.ClusterParameterGroupParameterArgs;
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
        var default_ = new ClusterParameterGroup("default", ClusterParameterGroupArgs.builder()        
            .description("RDS default cluster parameter group")
            .family("aurora5.6")
            .parameters(            
                ClusterParameterGroupParameterArgs.builder()
                    .name("character_set_server")
                    .value("utf8")
                    .build(),
                ClusterParameterGroupParameterArgs.builder()
                    .name("character_set_client")
                    .value("utf8")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:ClusterParameterGroup
    properties:
      description: RDS default cluster parameter group
      family: aurora5.6
      parameters:
        - name: character_set_server
          value: utf8
        - name: character_set_client
          value: utf8
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RDS Cluster Parameter Groups using the `name`. For example:

```sh
 $ pulumi import aws:rds/clusterParameterGroup:ClusterParameterGroup cluster_pg production-pg-1
```
 