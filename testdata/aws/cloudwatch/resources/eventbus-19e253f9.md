Provides an EventBridge event bus resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const messenger = new aws.cloudwatch.EventBus("messenger", {});
```
```python
import pulumi
import pulumi_aws as aws

messenger = aws.cloudwatch.EventBus("messenger")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var messenger = new Aws.CloudWatch.EventBus("messenger");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventBus(ctx, "messenger", nil)
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

    }
}
```
```yaml
resources:
  messenger:
    type: aws:cloudwatch:EventBus
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplepartnerEventSource = aws.cloudwatch.getEventSource({
    namePrefix: "aws.partner/examplepartner.com",
});
const examplepartnerEventBus = new aws.cloudwatch.EventBus("examplepartnerEventBus", {eventSourceName: examplepartnerEventSource.then(examplepartnerEventSource => examplepartnerEventSource.name)});
```
```python
import pulumi
import pulumi_aws as aws

examplepartner_event_source = aws.cloudwatch.get_event_source(name_prefix="aws.partner/examplepartner.com")
examplepartner_event_bus = aws.cloudwatch.EventBus("examplepartnerEventBus", event_source_name=examplepartner_event_source.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplepartnerEventSource = Aws.CloudWatch.GetEventSource.Invoke(new()
    {
        NamePrefix = "aws.partner/examplepartner.com",
    });

    var examplepartnerEventBus = new Aws.CloudWatch.EventBus("examplepartnerEventBus", new()
    {
        EventSourceName = examplepartnerEventSource.Apply(getEventSourceResult => getEventSourceResult.Name),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplepartnerEventSource, err := cloudwatch.GetEventSource(ctx, &cloudwatch.GetEventSourceArgs{
			NamePrefix: pulumi.StringRef("aws.partner/examplepartner.com"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventBus(ctx, "examplepartnerEventBus", &cloudwatch.EventBusArgs{
			EventSourceName: *pulumi.String(examplepartnerEventSource.Name),
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
import com.pulumi.aws.cloudwatch.CloudwatchFunctions;
import com.pulumi.aws.cloudwatch.inputs.GetEventSourceArgs;
import com.pulumi.aws.cloudwatch.EventBus;
import com.pulumi.aws.cloudwatch.EventBusArgs;
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
        final var examplepartnerEventSource = CloudwatchFunctions.getEventSource(GetEventSourceArgs.builder()
            .namePrefix("aws.partner/examplepartner.com")
            .build());

        var examplepartnerEventBus = new EventBus("examplepartnerEventBus", EventBusArgs.builder()        
            .eventSourceName(examplepartnerEventSource.applyValue(getEventSourceResult -> getEventSourceResult.name()))
            .build());

    }
}
```
```yaml
resources:
  examplepartnerEventBus:
    type: aws:cloudwatch:EventBus
    properties:
      eventSourceName: ${examplepartnerEventSource.name}
variables:
  examplepartnerEventSource:
    fn::invoke:
      Function: aws:cloudwatch:getEventSource
      Arguments:
        namePrefix: aws.partner/examplepartner.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge event buses using the `name` (which can also be a partner event source name). For example:

```sh
 $ pulumi import aws:cloudwatch/eventBus:EventBus messenger chat-messages
```
 