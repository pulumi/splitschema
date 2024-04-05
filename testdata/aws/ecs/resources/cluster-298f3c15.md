Provides an ECS cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ecs.Cluster("foo", {settings: [{
    name: "containerInsights",
    value: "enabled",
}]});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ecs.Cluster("foo", settings=[aws.ecs.ClusterSettingArgs(
    name="containerInsights",
    value="enabled",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ecs.Cluster("foo", new()
    {
        Settings = new[]
        {
            new Aws.Ecs.Inputs.ClusterSettingArgs
            {
                Name = "containerInsights",
                Value = "enabled",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewCluster(ctx, "foo", &ecs.ClusterArgs{
			Settings: ecs.ClusterSettingArray{
				&ecs.ClusterSettingArgs{
					Name:  pulumi.String("containerInsights"),
					Value: pulumi.String("enabled"),
				},
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
import com.pulumi.aws.ecs.Cluster;
import com.pulumi.aws.ecs.ClusterArgs;
import com.pulumi.aws.ecs.inputs.ClusterSettingArgs;
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
        var foo = new Cluster("foo", ClusterArgs.builder()        
            .settings(ClusterSettingArgs.builder()
                .name("containerInsights")
                .value("enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ecs:Cluster
    properties:
      settings:
        - name: containerInsights
          value: enabled
```
{{% /example %}}
{{% example %}}
### Example with Log Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "example",
    deletionWindowInDays: 7,
});
const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const test = new aws.ecs.Cluster("test", {configuration: {
    executeCommandConfiguration: {
        kmsKeyId: exampleKey.arn,
        logging: "OVERRIDE",
        logConfiguration: {
            cloudWatchEncryptionEnabled: true,
            cloudWatchLogGroupName: exampleLogGroup.name,
        },
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="example",
    deletion_window_in_days=7)
example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
test = aws.ecs.Cluster("test", configuration=aws.ecs.ClusterConfigurationArgs(
    execute_command_configuration=aws.ecs.ClusterConfigurationExecuteCommandConfigurationArgs(
        kms_key_id=example_key.arn,
        logging="OVERRIDE",
        log_configuration=aws.ecs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs(
            cloud_watch_encryption_enabled=True,
            cloud_watch_log_group_name=example_log_group.name,
        ),
    ),
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "example",
        DeletionWindowInDays = 7,
    });

    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var test = new Aws.Ecs.Cluster("test", new()
    {
        Configuration = new Aws.Ecs.Inputs.ClusterConfigurationArgs
        {
            ExecuteCommandConfiguration = new Aws.Ecs.Inputs.ClusterConfigurationExecuteCommandConfigurationArgs
            {
                KmsKeyId = exampleKey.Arn,
                Logging = "OVERRIDE",
                LogConfiguration = new Aws.Ecs.Inputs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs
                {
                    CloudWatchEncryptionEnabled = true,
                    CloudWatchLogGroupName = exampleLogGroup.Name,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("example"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		_, err = ecs.NewCluster(ctx, "test", &ecs.ClusterArgs{
			Configuration: &ecs.ClusterConfigurationArgs{
				ExecuteCommandConfiguration: &ecs.ClusterConfigurationExecuteCommandConfigurationArgs{
					KmsKeyId: exampleKey.Arn,
					Logging:  pulumi.String("OVERRIDE"),
					LogConfiguration: &ecs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs{
						CloudWatchEncryptionEnabled: pulumi.Bool(true),
						CloudWatchLogGroupName:      exampleLogGroup.Name,
					},
				},
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.ecs.Cluster;
import com.pulumi.aws.ecs.ClusterArgs;
import com.pulumi.aws.ecs.inputs.ClusterConfigurationArgs;
import com.pulumi.aws.ecs.inputs.ClusterConfigurationExecuteCommandConfigurationArgs;
import com.pulumi.aws.ecs.inputs.ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("example")
            .deletionWindowInDays(7)
            .build());

        var exampleLogGroup = new LogGroup("exampleLogGroup");

        var test = new Cluster("test", ClusterArgs.builder()        
            .configuration(ClusterConfigurationArgs.builder()
                .executeCommandConfiguration(ClusterConfigurationExecuteCommandConfigurationArgs.builder()
                    .kmsKeyId(exampleKey.arn())
                    .logging("OVERRIDE")
                    .logConfiguration(ClusterConfigurationExecuteCommandConfigurationLogConfigurationArgs.builder()
                        .cloudWatchEncryptionEnabled(true)
                        .cloudWatchLogGroupName(exampleLogGroup.name())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: example
      deletionWindowInDays: 7
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  test:
    type: aws:ecs:Cluster
    properties:
      configuration:
        executeCommandConfiguration:
          kmsKeyId: ${exampleKey.arn}
          logging: OVERRIDE
          logConfiguration:
            cloudWatchEncryptionEnabled: true
            cloudWatchLogGroupName: ${exampleLogGroup.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS clusters using the `name`. For example:

```sh
 $ pulumi import aws:ecs/cluster:Cluster stateless stateless-app
```
 