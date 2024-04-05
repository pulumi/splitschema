Data source for listing RDS Database Instances.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.rds.getInstances({
    filters: [{
        name: "db-instance-id",
        values: ["my-database-id"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.get_instances(filters=[aws.rds.GetInstancesFilterArgs(
    name="db-instance-id",
    values=["my-database-id"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Rds.GetInstances.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Rds.Inputs.GetInstancesFilterInputArgs
            {
                Name = "db-instance-id",
                Values = new[]
                {
                    "my-database-id",
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
		_, err := rds.GetInstances(ctx, &rds.GetInstancesArgs{
			Filters: []rds.GetInstancesFilter{
				{
					Name: "db-instance-id",
					Values: []string{
						"my-database-id",
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
import com.pulumi.aws.rds.inputs.GetInstancesArgs;
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
        final var example = RdsFunctions.getInstances(GetInstancesArgs.builder()
            .filters(GetInstancesFilterArgs.builder()
                .name("db-instance-id")
                .values("my-database-id")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:rds:getInstances
      Arguments:
        filters:
          - name: db-instance-id
            values:
              - my-database-id
```
{{% /example %}}
{{% example %}}
### Using tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.rds.getInstances({
    tags: {
        Env: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.get_instances(tags={
    "Env": "test",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Rds.GetInstances.Invoke(new()
    {
        Tags = 
        {
            { "Env", "test" },
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
		_, err := rds.GetInstances(ctx, &rds.GetInstancesArgs{
			Tags: map[string]interface{}{
				"Env": "test",
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
import com.pulumi.aws.rds.inputs.GetInstancesArgs;
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
        final var example = RdsFunctions.getInstances(GetInstancesArgs.builder()
            .tags(Map.of("Env", "test"))
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:rds:getInstances
      Arguments:
        tags:
          Env: test
```
{{% /example %}}
{{% /examples %}}