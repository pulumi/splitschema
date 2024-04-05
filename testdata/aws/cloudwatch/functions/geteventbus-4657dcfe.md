This data source can be used to fetch information about a specific
EventBridge event bus. Use this data source to compute the ARN of
an event bus, given the name of the bus.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.cloudwatch.getEventBus({
    name: "example-bus-name",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudwatch.get_event_bus(name="example-bus-name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CloudWatch.GetEventBus.Invoke(new()
    {
        Name = "example-bus-name",
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
		_, err := cloudwatch.LookupEventBus(ctx, &cloudwatch.LookupEventBusArgs{
			Name: "example-bus-name",
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
import com.pulumi.aws.cloudwatch.CloudwatchFunctions;
import com.pulumi.aws.cloudwatch.inputs.GetEventBusArgs;
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
        final var example = CloudwatchFunctions.getEventBus(GetEventBusArgs.builder()
            .name("example-bus-name")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:cloudwatch:getEventBus
      Arguments:
        name: example-bus-name
```
{{% /example %}}
{{% /examples %}}