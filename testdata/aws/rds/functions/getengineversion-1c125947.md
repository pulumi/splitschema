Information about an RDS engine version.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.rds.getEngineVersion({
    engine: "mysql",
    preferredVersions: [
        "8.0.27",
        "8.0.26",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.rds.get_engine_version(engine="mysql",
    preferred_versions=[
        "8.0.27",
        "8.0.26",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Rds.GetEngineVersion.Invoke(new()
    {
        Engine = "mysql",
        PreferredVersions = new[]
        {
            "8.0.27",
            "8.0.26",
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
		_, err := rds.GetEngineVersion(ctx, &rds.GetEngineVersionArgs{
			Engine: "mysql",
			PreferredVersions: []string{
				"8.0.27",
				"8.0.26",
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
import com.pulumi.aws.rds.inputs.GetEngineVersionArgs;
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
        final var test = RdsFunctions.getEngineVersion(GetEngineVersionArgs.builder()
            .engine("mysql")
            .preferredVersions(            
                "8.0.27",
                "8.0.26")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:rds:getEngineVersion
      Arguments:
        engine: mysql
        preferredVersions:
          - 8.0.27
          - 8.0.26
```
{{% /example %}}
{{% example %}}
### With `filter`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.rds.getEngineVersion({
    engine: "aurora-postgresql",
    filters: [{
        name: "engine-mode",
        values: ["serverless"],
    }],
    includeAll: true,
    version: "10.14",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.rds.get_engine_version(engine="aurora-postgresql",
    filters=[aws.rds.GetEngineVersionFilterArgs(
        name="engine-mode",
        values=["serverless"],
    )],
    include_all=True,
    version="10.14")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Rds.GetEngineVersion.Invoke(new()
    {
        Engine = "aurora-postgresql",
        Filters = new[]
        {
            new Aws.Rds.Inputs.GetEngineVersionFilterInputArgs
            {
                Name = "engine-mode",
                Values = new[]
                {
                    "serverless",
                },
            },
        },
        IncludeAll = true,
        Version = "10.14",
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
		_, err := rds.GetEngineVersion(ctx, &rds.GetEngineVersionArgs{
			Engine: "aurora-postgresql",
			Filters: []rds.GetEngineVersionFilter{
				{
					Name: "engine-mode",
					Values: []string{
						"serverless",
					},
				},
			},
			IncludeAll: pulumi.BoolRef(true),
			Version:    pulumi.StringRef("10.14"),
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
import com.pulumi.aws.rds.inputs.GetEngineVersionArgs;
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
        final var test = RdsFunctions.getEngineVersion(GetEngineVersionArgs.builder()
            .engine("aurora-postgresql")
            .filters(GetEngineVersionFilterArgs.builder()
                .name("engine-mode")
                .values("serverless")
                .build())
            .includeAll(true)
            .version("10.14")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:rds:getEngineVersion
      Arguments:
        engine: aurora-postgresql
        filters:
          - name: engine-mode
            values:
              - serverless
        includeAll: true
        version: '10.14'
```
{{% /example %}}
{{% /examples %}}