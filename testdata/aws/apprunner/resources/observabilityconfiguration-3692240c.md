Manages an App Runner Observability Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apprunner.ObservabilityConfiguration("example", {
    observabilityConfigurationName: "example",
    tags: {
        Name: "example-apprunner-observability-configuration",
    },
    traceConfiguration: {
        vendor: "AWSXRAY",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apprunner.ObservabilityConfiguration("example",
    observability_configuration_name="example",
    tags={
        "Name": "example-apprunner-observability-configuration",
    },
    trace_configuration=aws.apprunner.ObservabilityConfigurationTraceConfigurationArgs(
        vendor="AWSXRAY",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppRunner.ObservabilityConfiguration("example", new()
    {
        ObservabilityConfigurationName = "example",
        Tags = 
        {
            { "Name", "example-apprunner-observability-configuration" },
        },
        TraceConfiguration = new Aws.AppRunner.Inputs.ObservabilityConfigurationTraceConfigurationArgs
        {
            Vendor = "AWSXRAY",
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
		_, err := apprunner.NewObservabilityConfiguration(ctx, "example", &apprunner.ObservabilityConfigurationArgs{
			ObservabilityConfigurationName: pulumi.String("example"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-apprunner-observability-configuration"),
			},
			TraceConfiguration: &apprunner.ObservabilityConfigurationTraceConfigurationArgs{
				Vendor: pulumi.String("AWSXRAY"),
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
import com.pulumi.aws.apprunner.ObservabilityConfiguration;
import com.pulumi.aws.apprunner.ObservabilityConfigurationArgs;
import com.pulumi.aws.apprunner.inputs.ObservabilityConfigurationTraceConfigurationArgs;
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
        var example = new ObservabilityConfiguration("example", ObservabilityConfigurationArgs.builder()        
            .observabilityConfigurationName("example")
            .tags(Map.of("Name", "example-apprunner-observability-configuration"))
            .traceConfiguration(ObservabilityConfigurationTraceConfigurationArgs.builder()
                .vendor("AWSXRAY")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apprunner:ObservabilityConfiguration
    properties:
      observabilityConfigurationName: example
      tags:
        Name: example-apprunner-observability-configuration
      traceConfiguration:
        vendor: AWSXRAY
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner Observability Configuration using the `arn`. For example:

```sh
 $ pulumi import aws:apprunner/observabilityConfiguration:ObservabilityConfiguration example arn:aws:apprunner:us-east-1:1234567890:observabilityconfiguration/example/1/d75bc7ea55b71e724fe5c23452fe22a1
```
 