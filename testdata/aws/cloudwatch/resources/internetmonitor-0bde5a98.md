Provides a Internet Monitor Monitor resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudwatch.InternetMonitor("example", {monitorName: "exmple"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudwatch.InternetMonitor("example", monitor_name="exmple")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudWatch.InternetMonitor("example", new()
    {
        MonitorName = "exmple",
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
		_, err := cloudwatch.NewInternetMonitor(ctx, "example", &cloudwatch.InternetMonitorArgs{
			MonitorName: pulumi.String("exmple"),
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
import com.pulumi.aws.cloudwatch.InternetMonitor;
import com.pulumi.aws.cloudwatch.InternetMonitorArgs;
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
        var example = new InternetMonitor("example", InternetMonitorArgs.builder()        
            .monitorName("exmple")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudwatch:InternetMonitor
    properties:
      monitorName: exmple
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Internet Monitor Monitors using the `monitor_name`. For example:

```sh
 $ pulumi import aws:cloudwatch/internetMonitor:InternetMonitor some some-monitor
```
 