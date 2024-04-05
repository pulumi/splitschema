Manages the default App Runner auto scaling configuration.
When creating or updating this resource the existing default auto scaling configuration will be set to non-default automatically.
When creating or updating this resource the configuration is automatically assigned as the default to the new services you create in the future. The new default designation doesn't affect the associations that were previously set for existing services.
Each account can have only one default auto scaling configuration per Region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAutoScalingConfigurationVersion = new aws.apprunner.AutoScalingConfigurationVersion("exampleAutoScalingConfigurationVersion", {
    autoScalingConfigurationName: "example",
    maxConcurrency: 50,
    maxSize: 10,
    minSize: 2,
});
const exampleDefaultAutoScalingConfigurationVersion = new aws.apprunner.DefaultAutoScalingConfigurationVersion("exampleDefaultAutoScalingConfigurationVersion", {autoScalingConfigurationArn: exampleAutoScalingConfigurationVersion.arn});
```
```python
import pulumi
import pulumi_aws as aws

example_auto_scaling_configuration_version = aws.apprunner.AutoScalingConfigurationVersion("exampleAutoScalingConfigurationVersion",
    auto_scaling_configuration_name="example",
    max_concurrency=50,
    max_size=10,
    min_size=2)
example_default_auto_scaling_configuration_version = aws.apprunner.DefaultAutoScalingConfigurationVersion("exampleDefaultAutoScalingConfigurationVersion", auto_scaling_configuration_arn=example_auto_scaling_configuration_version.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAutoScalingConfigurationVersion = new Aws.AppRunner.AutoScalingConfigurationVersion("exampleAutoScalingConfigurationVersion", new()
    {
        AutoScalingConfigurationName = "example",
        MaxConcurrency = 50,
        MaxSize = 10,
        MinSize = 2,
    });

    var exampleDefaultAutoScalingConfigurationVersion = new Aws.AppRunner.DefaultAutoScalingConfigurationVersion("exampleDefaultAutoScalingConfigurationVersion", new()
    {
        AutoScalingConfigurationArn = exampleAutoScalingConfigurationVersion.Arn,
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
		exampleAutoScalingConfigurationVersion, err := apprunner.NewAutoScalingConfigurationVersion(ctx, "exampleAutoScalingConfigurationVersion", &apprunner.AutoScalingConfigurationVersionArgs{
			AutoScalingConfigurationName: pulumi.String("example"),
			MaxConcurrency:               pulumi.Int(50),
			MaxSize:                      pulumi.Int(10),
			MinSize:                      pulumi.Int(2),
		})
		if err != nil {
			return err
		}
		_, err = apprunner.NewDefaultAutoScalingConfigurationVersion(ctx, "exampleDefaultAutoScalingConfigurationVersion", &apprunner.DefaultAutoScalingConfigurationVersionArgs{
			AutoScalingConfigurationArn: exampleAutoScalingConfigurationVersion.Arn,
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
import com.pulumi.aws.apprunner.DefaultAutoScalingConfigurationVersion;
import com.pulumi.aws.apprunner.DefaultAutoScalingConfigurationVersionArgs;
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
        var exampleAutoScalingConfigurationVersion = new AutoScalingConfigurationVersion("exampleAutoScalingConfigurationVersion", AutoScalingConfigurationVersionArgs.builder()        
            .autoScalingConfigurationName("example")
            .maxConcurrency(50)
            .maxSize(10)
            .minSize(2)
            .build());

        var exampleDefaultAutoScalingConfigurationVersion = new DefaultAutoScalingConfigurationVersion("exampleDefaultAutoScalingConfigurationVersion", DefaultAutoScalingConfigurationVersionArgs.builder()        
            .autoScalingConfigurationArn(exampleAutoScalingConfigurationVersion.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleAutoScalingConfigurationVersion:
    type: aws:apprunner:AutoScalingConfigurationVersion
    properties:
      autoScalingConfigurationName: example
      maxConcurrency: 50
      maxSize: 10
      minSize: 2
  exampleDefaultAutoScalingConfigurationVersion:
    type: aws:apprunner:DefaultAutoScalingConfigurationVersion
    properties:
      autoScalingConfigurationArn: ${exampleAutoScalingConfigurationVersion.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner default auto scaling configurations using the current Region. For example:

```sh
 $ pulumi import aws:apprunner/defaultAutoScalingConfigurationVersion:DefaultAutoScalingConfigurationVersion example us-west-2
```
 