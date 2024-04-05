Provides an AppConfig Environment resource for an `aws.appconfig.Application` resource. One or more environments can be defined for an application.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApplication = new aws.appconfig.Application("exampleApplication", {
    description: "Example AppConfig Application",
    tags: {
        Type: "AppConfig Application",
    },
});
const exampleEnvironment = new aws.appconfig.Environment("exampleEnvironment", {
    description: "Example AppConfig Environment",
    applicationId: exampleApplication.id,
    monitors: [{
        alarmArn: aws_cloudwatch_metric_alarm.example.arn,
        alarmRoleArn: aws_iam_role.example.arn,
    }],
    tags: {
        Type: "AppConfig Environment",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_application = aws.appconfig.Application("exampleApplication",
    description="Example AppConfig Application",
    tags={
        "Type": "AppConfig Application",
    })
example_environment = aws.appconfig.Environment("exampleEnvironment",
    description="Example AppConfig Environment",
    application_id=example_application.id,
    monitors=[aws.appconfig.EnvironmentMonitorArgs(
        alarm_arn=aws_cloudwatch_metric_alarm["example"]["arn"],
        alarm_role_arn=aws_iam_role["example"]["arn"],
    )],
    tags={
        "Type": "AppConfig Environment",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleApplication = new Aws.AppConfig.Application("exampleApplication", new()
    {
        Description = "Example AppConfig Application",
        Tags = 
        {
            { "Type", "AppConfig Application" },
        },
    });

    var exampleEnvironment = new Aws.AppConfig.Environment("exampleEnvironment", new()
    {
        Description = "Example AppConfig Environment",
        ApplicationId = exampleApplication.Id,
        Monitors = new[]
        {
            new Aws.AppConfig.Inputs.EnvironmentMonitorArgs
            {
                AlarmArn = aws_cloudwatch_metric_alarm.Example.Arn,
                AlarmRoleArn = aws_iam_role.Example.Arn,
            },
        },
        Tags = 
        {
            { "Type", "AppConfig Environment" },
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
		exampleApplication, err := appconfig.NewApplication(ctx, "exampleApplication", &appconfig.ApplicationArgs{
			Description: pulumi.String("Example AppConfig Application"),
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Application"),
			},
		})
		if err != nil {
			return err
		}
		_, err = appconfig.NewEnvironment(ctx, "exampleEnvironment", &appconfig.EnvironmentArgs{
			Description:   pulumi.String("Example AppConfig Environment"),
			ApplicationId: exampleApplication.ID(),
			Monitors: appconfig.EnvironmentMonitorArray{
				&appconfig.EnvironmentMonitorArgs{
					AlarmArn:     pulumi.Any(aws_cloudwatch_metric_alarm.Example.Arn),
					AlarmRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
				},
			},
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Environment"),
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
import com.pulumi.aws.appconfig.Application;
import com.pulumi.aws.appconfig.ApplicationArgs;
import com.pulumi.aws.appconfig.Environment;
import com.pulumi.aws.appconfig.EnvironmentArgs;
import com.pulumi.aws.appconfig.inputs.EnvironmentMonitorArgs;
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
        var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()        
            .description("Example AppConfig Application")
            .tags(Map.of("Type", "AppConfig Application"))
            .build());

        var exampleEnvironment = new Environment("exampleEnvironment", EnvironmentArgs.builder()        
            .description("Example AppConfig Environment")
            .applicationId(exampleApplication.id())
            .monitors(EnvironmentMonitorArgs.builder()
                .alarmArn(aws_cloudwatch_metric_alarm.example().arn())
                .alarmRoleArn(aws_iam_role.example().arn())
                .build())
            .tags(Map.of("Type", "AppConfig Environment"))
            .build());

    }
}
```
```yaml
resources:
  exampleEnvironment:
    type: aws:appconfig:Environment
    properties:
      description: Example AppConfig Environment
      applicationId: ${exampleApplication.id}
      monitors:
        - alarmArn: ${aws_cloudwatch_metric_alarm.example.arn}
          alarmRoleArn: ${aws_iam_role.example.arn}
      tags:
        Type: AppConfig Environment
  exampleApplication:
    type: aws:appconfig:Application
    properties:
      description: Example AppConfig Application
      tags:
        Type: AppConfig Application
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Environments using the environment ID and application ID separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:appconfig/environment:Environment example 71abcde:11xxxxx
```
 