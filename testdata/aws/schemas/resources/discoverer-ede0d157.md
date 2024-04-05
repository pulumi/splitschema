Provides an EventBridge Schema Discoverer resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const messenger = new aws.cloudwatch.EventBus("messenger", {});
const test = new aws.schemas.Discoverer("test", {
    sourceArn: messenger.arn,
    description: "Auto discover event schemas",
});
```
```python
import pulumi
import pulumi_aws as aws

messenger = aws.cloudwatch.EventBus("messenger")
test = aws.schemas.Discoverer("test",
    source_arn=messenger.arn,
    description="Auto discover event schemas")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var messenger = new Aws.CloudWatch.EventBus("messenger");

    var test = new Aws.Schemas.Discoverer("test", new()
    {
        SourceArn = messenger.Arn,
        Description = "Auto discover event schemas",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/schemas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		messenger, err := cloudwatch.NewEventBus(ctx, "messenger", nil)
		if err != nil {
			return err
		}
		_, err = schemas.NewDiscoverer(ctx, "test", &schemas.DiscovererArgs{
			SourceArn:   messenger.Arn,
			Description: pulumi.String("Auto discover event schemas"),
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
import com.pulumi.aws.cloudwatch.EventBus;
import com.pulumi.aws.schemas.Discoverer;
import com.pulumi.aws.schemas.DiscovererArgs;
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
        var messenger = new EventBus("messenger");

        var test = new Discoverer("test", DiscovererArgs.builder()        
            .sourceArn(messenger.arn())
            .description("Auto discover event schemas")
            .build());

    }
}
```
```yaml
resources:
  messenger:
    type: aws:cloudwatch:EventBus
  test:
    type: aws:schemas:Discoverer
    properties:
      sourceArn: ${messenger.arn}
      description: Auto discover event schemas
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge discoverers using the `id`. For example:

```sh
 $ pulumi import aws:schemas/discoverer:Discoverer test 123
```
 