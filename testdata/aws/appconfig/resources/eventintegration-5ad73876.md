Provides an Amazon AppIntegrations Event Integration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.EventIntegration("example", {
    description: "Example Description",
    eventFilter: {
        source: "aws.partner/examplepartner.com",
    },
    eventbridgeBus: "default",
    tags: {
        Name: "Example Event Integration",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appconfig.EventIntegration("example",
    description="Example Description",
    event_filter=aws.appconfig.EventIntegrationEventFilterArgs(
        source="aws.partner/examplepartner.com",
    ),
    eventbridge_bus="default",
    tags={
        "Name": "Example Event Integration",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppConfig.EventIntegration("example", new()
    {
        Description = "Example Description",
        EventFilter = new Aws.AppConfig.Inputs.EventIntegrationEventFilterArgs
        {
            Source = "aws.partner/examplepartner.com",
        },
        EventbridgeBus = "default",
        Tags = 
        {
            { "Name", "Example Event Integration" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfig.NewEventIntegration(ctx, "example", &appconfig.EventIntegrationArgs{
			Description: pulumi.String("Example Description"),
			EventFilter: &appconfig.EventIntegrationEventFilterArgs{
				Source: pulumi.String("aws.partner/examplepartner.com"),
			},
			EventbridgeBus: pulumi.String("default"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Event Integration"),
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
import com.pulumi.aws.appconfig.EventIntegration;
import com.pulumi.aws.appconfig.EventIntegrationArgs;
import com.pulumi.aws.appconfig.inputs.EventIntegrationEventFilterArgs;
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
        var example = new EventIntegration("example", EventIntegrationArgs.builder()        
            .description("Example Description")
            .eventFilter(EventIntegrationEventFilterArgs.builder()
                .source("aws.partner/examplepartner.com")
                .build())
            .eventbridgeBus("default")
            .tags(Map.of("Name", "Example Event Integration"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:EventIntegration
    properties:
      description: Example Description
      eventFilter:
        source: aws.partner/examplepartner.com
      eventbridgeBus: default
      tags:
        Name: Example Event Integration
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon AppIntegrations Event Integrations using the `name`. For example:

```sh
 $ pulumi import aws:appconfig/eventIntegration:EventIntegration example example-name
```
 