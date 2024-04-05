{{% examples %}}
## Example Usage
{{% example %}}

List the event categories of all the RDS resources.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEventCategories = aws.rds.getEventCategories({});
export const example = exampleEventCategories.then(exampleEventCategories => exampleEventCategories.eventCategories);
```
```python
import pulumi
import pulumi_aws as aws

example_event_categories = aws.rds.get_event_categories()
pulumi.export("example", example_event_categories.event_categories)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEventCategories = Aws.Rds.GetEventCategories.Invoke();

    return new Dictionary<string, object?>
    {
        ["example"] = exampleEventCategories.Apply(getEventCategoriesResult => getEventCategoriesResult.EventCategories),
    };
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
		exampleEventCategories, err := rds.GetEventCategories(ctx, nil, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleEventCategories.EventCategories)
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
import com.pulumi.aws.rds.inputs.GetEventCategoriesArgs;
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
        final var exampleEventCategories = RdsFunctions.getEventCategories();

        ctx.export("example", exampleEventCategories.applyValue(getEventCategoriesResult -> getEventCategoriesResult.eventCategories()));
    }
}
```
```yaml
variables:
  exampleEventCategories:
    fn::invoke:
      Function: aws:rds:getEventCategories
      Arguments: {}
outputs:
  example: ${exampleEventCategories.eventCategories}
```

List the event categories specific to the RDS resource `db-snapshot`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEventCategories = aws.rds.getEventCategories({
    sourceType: "db-snapshot",
});
export const example = exampleEventCategories.then(exampleEventCategories => exampleEventCategories.eventCategories);
```
```python
import pulumi
import pulumi_aws as aws

example_event_categories = aws.rds.get_event_categories(source_type="db-snapshot")
pulumi.export("example", example_event_categories.event_categories)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEventCategories = Aws.Rds.GetEventCategories.Invoke(new()
    {
        SourceType = "db-snapshot",
    });

    return new Dictionary<string, object?>
    {
        ["example"] = exampleEventCategories.Apply(getEventCategoriesResult => getEventCategoriesResult.EventCategories),
    };
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
		exampleEventCategories, err := rds.GetEventCategories(ctx, &rds.GetEventCategoriesArgs{
			SourceType: pulumi.StringRef("db-snapshot"),
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("example", exampleEventCategories.EventCategories)
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
import com.pulumi.aws.rds.inputs.GetEventCategoriesArgs;
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
        final var exampleEventCategories = RdsFunctions.getEventCategories(GetEventCategoriesArgs.builder()
            .sourceType("db-snapshot")
            .build());

        ctx.export("example", exampleEventCategories.applyValue(getEventCategoriesResult -> getEventCategoriesResult.eventCategories()));
    }
}
```
```yaml
variables:
  exampleEventCategories:
    fn::invoke:
      Function: aws:rds:getEventCategories
      Arguments:
        sourceType: db-snapshot
outputs:
  example: ${exampleEventCategories.eventCategories}
```
{{% /example %}}
{{% /examples %}}