Manages an App Runner AutoScaling Configuration Version.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apprunner.AutoScalingConfigurationVersion("example", {
    autoScalingConfigurationName: "example",
    maxConcurrency: 50,
    maxSize: 10,
    minSize: 2,
    tags: {
        Name: "example-apprunner-autoscaling",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apprunner.AutoScalingConfigurationVersion("example",
    auto_scaling_configuration_name="example",
    max_concurrency=50,
    max_size=10,
    min_size=2,
    tags={
        "Name": "example-apprunner-autoscaling",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppRunner.AutoScalingConfigurationVersion("example", new()
    {
        AutoScalingConfigurationName = "example",
        MaxConcurrency = 50,
        MaxSize = 10,
        MinSize = 2,
        Tags = 
        {
            { "Name", "example-apprunner-autoscaling" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apprunner"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apprunner.NewAutoScalingConfigurationVersion(ctx, "example", &apprunner.AutoScalingConfigurationVersionArgs{
			AutoScalingConfigurationName: pulumi.String("example"),
			MaxConcurrency:               pulumi.Int(50),
			MaxSize:                      pulumi.Int(10),
			MinSize:                      pulumi.Int(2),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-apprunner-autoscaling"),
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
import com.pulumi.aws.apprunner.AutoScalingConfigurationVersion;
import com.pulumi.aws.apprunner.AutoScalingConfigurationVersionArgs;
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
        var example = new AutoScalingConfigurationVersion("example", AutoScalingConfigurationVersionArgs.builder()        
            .autoScalingConfigurationName("example")
            .maxConcurrency(50)
            .maxSize(10)
            .minSize(2)
            .tags(Map.of("Name", "example-apprunner-autoscaling"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apprunner:AutoScalingConfigurationVersion
    properties:
      autoScalingConfigurationName: example
      maxConcurrency: 50
      maxSize: 10
      minSize: 2
      tags:
        Name: example-apprunner-autoscaling
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner AutoScaling Configuration Versions using the `arn`. For example:

```sh
 $ pulumi import aws:apprunner/autoScalingConfigurationVersion:AutoScalingConfigurationVersion example "arn:aws:apprunner:us-east-1:1234567890:autoscalingconfiguration/example/1/69bdfe0115224b0db49398b7beb68e0f
```
 