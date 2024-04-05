Provides a Glue Data Quality Ruleset Resource. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/glue-data-quality.html) for a full explanation of the Glue Data Quality Ruleset functionality

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.DataQualityRuleset("example", {ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.DataQualityRuleset("example", ruleset="Rules = [Completeness \"colA\" between 0.4 and 0.8]")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.DataQualityRuleset("example", new()
    {
        Ruleset = "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewDataQualityRuleset(ctx, "example", &glue.DataQualityRulesetArgs{
			Ruleset: pulumi.String("Rules = [Completeness \"colA\" between 0.4 and 0.8]"),
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
import com.pulumi.aws.glue.DataQualityRuleset;
import com.pulumi.aws.glue.DataQualityRulesetArgs;
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
        var example = new DataQualityRuleset("example", DataQualityRulesetArgs.builder()        
            .ruleset("Rules = [Completeness \"colA\" between 0.4 and 0.8]")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:DataQualityRuleset
    properties:
      ruleset: Rules = [Completeness "colA" between 0.4 and 0.8]
```
{{% /example %}}
{{% example %}}
### With description

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.DataQualityRuleset("example", {
    description: "example",
    ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.DataQualityRuleset("example",
    description="example",
    ruleset="Rules = [Completeness \"colA\" between 0.4 and 0.8]")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.DataQualityRuleset("example", new()
    {
        Description = "example",
        Ruleset = "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewDataQualityRuleset(ctx, "example", &glue.DataQualityRulesetArgs{
			Description: pulumi.String("example"),
			Ruleset:     pulumi.String("Rules = [Completeness \"colA\" between 0.4 and 0.8]"),
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
import com.pulumi.aws.glue.DataQualityRuleset;
import com.pulumi.aws.glue.DataQualityRulesetArgs;
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
        var example = new DataQualityRuleset("example", DataQualityRulesetArgs.builder()        
            .description("example")
            .ruleset("Rules = [Completeness \"colA\" between 0.4 and 0.8]")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:DataQualityRuleset
    properties:
      description: example
      ruleset: Rules = [Completeness "colA" between 0.4 and 0.8]
```
{{% /example %}}
{{% example %}}
### With tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.DataQualityRuleset("example", {
    ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
    tags: {
        hello: "world",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.DataQualityRuleset("example",
    ruleset="Rules = [Completeness \"colA\" between 0.4 and 0.8]",
    tags={
        "hello": "world",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.DataQualityRuleset("example", new()
    {
        Ruleset = "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
        Tags = 
        {
            { "hello", "world" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewDataQualityRuleset(ctx, "example", &glue.DataQualityRulesetArgs{
			Ruleset: pulumi.String("Rules = [Completeness \"colA\" between 0.4 and 0.8]"),
			Tags: pulumi.StringMap{
				"hello": pulumi.String("world"),
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
import com.pulumi.aws.glue.DataQualityRuleset;
import com.pulumi.aws.glue.DataQualityRulesetArgs;
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
        var example = new DataQualityRuleset("example", DataQualityRulesetArgs.builder()        
            .ruleset("Rules = [Completeness \"colA\" between 0.4 and 0.8]")
            .tags(Map.of("hello", "world"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:DataQualityRuleset
    properties:
      ruleset: Rules = [Completeness "colA" between 0.4 and 0.8]
      tags:
        hello: world
```
{{% /example %}}
{{% example %}}
### With target_table

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.DataQualityRuleset("example", {
    ruleset: "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
    targetTable: {
        databaseName: aws_glue_catalog_database.example.name,
        tableName: aws_glue_catalog_table.example.name,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.DataQualityRuleset("example",
    ruleset="Rules = [Completeness \"colA\" between 0.4 and 0.8]",
    target_table=aws.glue.DataQualityRulesetTargetTableArgs(
        database_name=aws_glue_catalog_database["example"]["name"],
        table_name=aws_glue_catalog_table["example"]["name"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.DataQualityRuleset("example", new()
    {
        Ruleset = "Rules = [Completeness \"colA\" between 0.4 and 0.8]",
        TargetTable = new Aws.Glue.Inputs.DataQualityRulesetTargetTableArgs
        {
            DatabaseName = aws_glue_catalog_database.Example.Name,
            TableName = aws_glue_catalog_table.Example.Name,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewDataQualityRuleset(ctx, "example", &glue.DataQualityRulesetArgs{
			Ruleset: pulumi.String("Rules = [Completeness \"colA\" between 0.4 and 0.8]"),
			TargetTable: &glue.DataQualityRulesetTargetTableArgs{
				DatabaseName: pulumi.Any(aws_glue_catalog_database.Example.Name),
				TableName:    pulumi.Any(aws_glue_catalog_table.Example.Name),
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
import com.pulumi.aws.glue.DataQualityRuleset;
import com.pulumi.aws.glue.DataQualityRulesetArgs;
import com.pulumi.aws.glue.inputs.DataQualityRulesetTargetTableArgs;
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
        var example = new DataQualityRuleset("example", DataQualityRulesetArgs.builder()        
            .ruleset("Rules = [Completeness \"colA\" between 0.4 and 0.8]")
            .targetTable(DataQualityRulesetTargetTableArgs.builder()
                .databaseName(aws_glue_catalog_database.example().name())
                .tableName(aws_glue_catalog_table.example().name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:DataQualityRuleset
    properties:
      ruleset: Rules = [Completeness "colA" between 0.4 and 0.8]
      targetTable:
        databaseName: ${aws_glue_catalog_database.example.name}
        tableName: ${aws_glue_catalog_table.example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Data Quality Ruleset using the `name`. For example:

```sh
 $ pulumi import aws:glue/dataQualityRuleset:DataQualityRuleset example exampleName
```
 