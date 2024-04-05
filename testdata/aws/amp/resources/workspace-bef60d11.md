Manages an Amazon Managed Service for Prometheus (AMP) Workspace.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.amp.Workspace("example", {
    alias: "example",
    tags: {
        Environment: "production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.amp.Workspace("example",
    alias="example",
    tags={
        "Environment": "production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Amp.Workspace("example", new()
    {
        Alias = "example",
        Tags = 
        {
            { "Environment", "production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := amp.NewWorkspace(ctx, "example", &amp.WorkspaceArgs{
			Alias: pulumi.String("example"),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
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
import com.pulumi.aws.amp.Workspace;
import com.pulumi.aws.amp.WorkspaceArgs;
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
        var example = new Workspace("example", WorkspaceArgs.builder()        
            .alias("example")
            .tags(Map.of("Environment", "production"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:amp:Workspace
    properties:
      alias: example
      tags:
        Environment: production
```
{{% /example %}}
{{% example %}}
### CloudWatch Logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const exampleWorkspace = new aws.amp.Workspace("exampleWorkspace", {loggingConfiguration: {
    logGroupArn: pulumi.interpolate`${exampleLogGroup.arn}:*`,
}});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_workspace = aws.amp.Workspace("exampleWorkspace", logging_configuration=aws.amp.WorkspaceLoggingConfigurationArgs(
    log_group_arn=example_log_group.arn.apply(lambda arn: f"{arn}:*"),
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var exampleWorkspace = new Aws.Amp.Workspace("exampleWorkspace", new()
    {
        LoggingConfiguration = new Aws.Amp.Inputs.WorkspaceLoggingConfigurationArgs
        {
            LogGroupArn = exampleLogGroup.Arn.Apply(arn => $"{arn}:*"),
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amp"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		_, err = amp.NewWorkspace(ctx, "exampleWorkspace", &amp.WorkspaceArgs{
			LoggingConfiguration: &amp.WorkspaceLoggingConfigurationArgs{
				LogGroupArn: exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
					return fmt.Sprintf("%v:*", arn), nil
				}).(pulumi.StringOutput),
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
import com.pulumi.aws.amp.Workspace;
import com.pulumi.aws.amp.WorkspaceArgs;
import com.pulumi.aws.amp.inputs.WorkspaceLoggingConfigurationArgs;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup");

        var exampleWorkspace = new Workspace("exampleWorkspace", WorkspaceArgs.builder()        
            .loggingConfiguration(WorkspaceLoggingConfigurationArgs.builder()
                .logGroupArn(exampleLogGroup.arn().applyValue(arn -> String.format("%s:*", arn)))
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleWorkspace:
    type: aws:amp:Workspace
    properties:
      loggingConfiguration:
        logGroupArn: ${exampleLogGroup.arn}:*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AMP Workspaces using the identifier. For example:

```sh
 $ pulumi import aws:amp/workspace:Workspace demo ws-C6DCB907-F2D7-4D96-957B-66691F865D8B
```
 