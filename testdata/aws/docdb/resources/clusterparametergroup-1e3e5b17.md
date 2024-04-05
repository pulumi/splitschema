Manages a DocumentDB Cluster Parameter Group

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.docdb.ClusterParameterGroup("example", {
    description: "docdb cluster parameter group",
    family: "docdb3.6",
    parameters: [{
        name: "tls",
        value: "enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.docdb.ClusterParameterGroup("example",
    description="docdb cluster parameter group",
    family="docdb3.6",
    parameters=[aws.docdb.ClusterParameterGroupParameterArgs(
        name="tls",
        value="enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DocDB.ClusterParameterGroup("example", new()
    {
        Description = "docdb cluster parameter group",
        Family = "docdb3.6",
        Parameters = new[]
        {
            new Aws.DocDB.Inputs.ClusterParameterGroupParameterArgs
            {
                Name = "tls",
                Value = "enabled",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.NewClusterParameterGroup(ctx, "example", &docdb.ClusterParameterGroupArgs{
			Description: pulumi.String("docdb cluster parameter group"),
			Family:      pulumi.String("docdb3.6"),
			Parameters: docdb.ClusterParameterGroupParameterArray{
				&docdb.ClusterParameterGroupParameterArgs{
					Name:  pulumi.String("tls"),
					Value: pulumi.String("enabled"),
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
import com.pulumi.aws.docdb.ClusterParameterGroup;
import com.pulumi.aws.docdb.ClusterParameterGroupArgs;
import com.pulumi.aws.docdb.inputs.ClusterParameterGroupParameterArgs;
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
            .description("docdb cluster parameter group")
            .family("docdb3.6")
            .parameters(ClusterParameterGroupParameterArgs.builder()
                .name("tls")
                .value("enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:docdb:ClusterParameterGroup
    properties:
      description: docdb cluster parameter group
      family: docdb3.6
      parameters:
        - name: tls
          value: enabled
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DocumentDB Cluster Parameter Groups using the `name`. For example:

```sh
 $ pulumi import aws:docdb/clusterParameterGroup:ClusterParameterGroup cluster_pg production-pg-1
```
 