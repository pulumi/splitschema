Provides a CloudWatch Log Metric Filter resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dada = new aws.cloudwatch.LogGroup("dada", {});
const yada = new aws.cloudwatch.LogMetricFilter("yada", {
    pattern: "",
    logGroupName: dada.name,
    metricTransformation: {
        name: "EventCount",
        namespace: "YourNamespace",
        value: "1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

dada = aws.cloudwatch.LogGroup("dada")
yada = aws.cloudwatch.LogMetricFilter("yada",
    pattern="",
    log_group_name=dada.name,
    metric_transformation=aws.cloudwatch.LogMetricFilterMetricTransformationArgs(
        name="EventCount",
        namespace="YourNamespace",
        value="1",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dada = new Aws.CloudWatch.LogGroup("dada");

    var yada = new Aws.CloudWatch.LogMetricFilter("yada", new()
    {
        Pattern = "",
        LogGroupName = dada.Name,
        MetricTransformation = new Aws.CloudWatch.Inputs.LogMetricFilterMetricTransformationArgs
        {
            Name = "EventCount",
            Namespace = "YourNamespace",
            Value = "1",
        },
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
		dada, err := cloudwatch.NewLogGroup(ctx, "dada", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogMetricFilter(ctx, "yada", &cloudwatch.LogMetricFilterArgs{
			Pattern:      pulumi.String(""),
			LogGroupName: dada.Name,
			MetricTransformation: &cloudwatch.LogMetricFilterMetricTransformationArgs{
				Name:      pulumi.String("EventCount"),
				Namespace: pulumi.String("YourNamespace"),
				Value:     pulumi.String("1"),
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogMetricFilter;
import com.pulumi.aws.cloudwatch.LogMetricFilterArgs;
import com.pulumi.aws.cloudwatch.inputs.LogMetricFilterMetricTransformationArgs;
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
        var dada = new LogGroup("dada");

        var yada = new LogMetricFilter("yada", LogMetricFilterArgs.builder()        
            .pattern("")
            .logGroupName(dada.name())
            .metricTransformation(LogMetricFilterMetricTransformationArgs.builder()
                .name("EventCount")
                .namespace("YourNamespace")
                .value("1")
                .build())
            .build());

    }
}
```
```yaml
resources:
  yada:
    type: aws:cloudwatch:LogMetricFilter
    properties:
      pattern:
      logGroupName: ${dada.name}
      metricTransformation:
        name: EventCount
        namespace: YourNamespace
        value: '1'
  dada:
    type: aws:cloudwatch:LogGroup
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Log Metric Filter using the `log_group_name:name`. For example:

```sh
 $ pulumi import aws:cloudwatch/logMetricFilter:LogMetricFilter test /aws/lambda/function:test
```
 