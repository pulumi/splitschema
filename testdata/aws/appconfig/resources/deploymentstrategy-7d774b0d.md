Provides an AppConfig Deployment Strategy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.DeploymentStrategy("example", {
    deploymentDurationInMinutes: 3,
    description: "Example Deployment Strategy",
    finalBakeTimeInMinutes: 4,
    growthFactor: 10,
    growthType: "LINEAR",
    replicateTo: "NONE",
    tags: {
        Type: "AppConfig Deployment Strategy",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appconfig.DeploymentStrategy("example",
    deployment_duration_in_minutes=3,
    description="Example Deployment Strategy",
    final_bake_time_in_minutes=4,
    growth_factor=10,
    growth_type="LINEAR",
    replicate_to="NONE",
    tags={
        "Type": "AppConfig Deployment Strategy",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppConfig.DeploymentStrategy("example", new()
    {
        DeploymentDurationInMinutes = 3,
        Description = "Example Deployment Strategy",
        FinalBakeTimeInMinutes = 4,
        GrowthFactor = 10,
        GrowthType = "LINEAR",
        ReplicateTo = "NONE",
        Tags = 
        {
            { "Type", "AppConfig Deployment Strategy" },
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
		_, err := appconfig.NewDeploymentStrategy(ctx, "example", &appconfig.DeploymentStrategyArgs{
			DeploymentDurationInMinutes: pulumi.Int(3),
			Description:                 pulumi.String("Example Deployment Strategy"),
			FinalBakeTimeInMinutes:      pulumi.Int(4),
			GrowthFactor:                pulumi.Float64(10),
			GrowthType:                  pulumi.String("LINEAR"),
			ReplicateTo:                 pulumi.String("NONE"),
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Deployment Strategy"),
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
import com.pulumi.aws.appconfig.DeploymentStrategy;
import com.pulumi.aws.appconfig.DeploymentStrategyArgs;
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
        var example = new DeploymentStrategy("example", DeploymentStrategyArgs.builder()        
            .deploymentDurationInMinutes(3)
            .description("Example Deployment Strategy")
            .finalBakeTimeInMinutes(4)
            .growthFactor(10)
            .growthType("LINEAR")
            .replicateTo("NONE")
            .tags(Map.of("Type", "AppConfig Deployment Strategy"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:DeploymentStrategy
    properties:
      deploymentDurationInMinutes: 3
      description: Example Deployment Strategy
      finalBakeTimeInMinutes: 4
      growthFactor: 10
      growthType: LINEAR
      replicateTo: NONE
      tags:
        Type: AppConfig Deployment Strategy
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Deployment Strategies using their deployment strategy ID. For example:

```sh
 $ pulumi import aws:appconfig/deploymentStrategy:DeploymentStrategy example 11xxxxx
```
 