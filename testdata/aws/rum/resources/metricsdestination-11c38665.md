Provides a CloudWatch RUM Metrics Destination resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rum.MetricsDestination("example", {
    appMonitorName: aws_rum_app_monitor.example.name,
    destination: "CloudWatch",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rum.MetricsDestination("example",
    app_monitor_name=aws_rum_app_monitor["example"]["name"],
    destination="CloudWatch")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rum.MetricsDestination("example", new()
    {
        AppMonitorName = aws_rum_app_monitor.Example.Name,
        Destination = "CloudWatch",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rum"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rum.NewMetricsDestination(ctx, "example", &rum.MetricsDestinationArgs{
			AppMonitorName: pulumi.Any(aws_rum_app_monitor.Example.Name),
			Destination:    pulumi.String("CloudWatch"),
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
import com.pulumi.aws.rum.MetricsDestination;
import com.pulumi.aws.rum.MetricsDestinationArgs;
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
        var example = new MetricsDestination("example", MetricsDestinationArgs.builder()        
            .appMonitorName(aws_rum_app_monitor.example().name())
            .destination("CloudWatch")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rum:MetricsDestination
    properties:
      appMonitorName: ${aws_rum_app_monitor.example.name}
      destination: CloudWatch
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudwatch RUM Metrics Destination using the `id`. For example:

```sh
 $ pulumi import aws:rum/metricsDestination:MetricsDestination example example
```
 