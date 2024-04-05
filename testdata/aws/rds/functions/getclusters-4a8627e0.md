Data source for managing an AWS RDS (Relational Database) Clusters.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.rds.getClusters({
    filters: [{
        name: "engine",
        values: ["aurora-postgresql"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.get_clusters(filters=[aws.rds.GetClustersFilterArgs(
    name="engine",
    values=["aurora-postgresql"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Rds.GetClusters.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Rds.Inputs.GetClustersFilterInputArgs
            {
                Name = "engine",
                Values = new[]
                {
                    "aurora-postgresql",
                },
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
		_, err := rds.GetClusters(ctx, &rds.GetClustersArgs{
			Filters: []rds.GetClustersFilter{
				{
					Name: "engine",
					Values: []string{
						"aurora-postgresql",
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetClustersArgs;
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
        final var example = RdsFunctions.getClusters(GetClustersArgs.builder()
            .filters(GetClustersFilterArgs.builder()
                .name("engine")
                .values("aurora-postgresql")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:rds:getClusters
      Arguments:
        filters:
          - name: engine
            values:
              - aurora-postgresql
```
{{% /example %}}
{{% /examples %}}