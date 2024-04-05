Provides a CodeDeploy deployment config for an application

{{% examples %}}
## Example Usage
{{% example %}}
### Server Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooDeploymentConfig = new aws.codedeploy.DeploymentConfig("fooDeploymentConfig", {
    deploymentConfigName: "test-deployment-config",
    minimumHealthyHosts: {
        type: "HOST_COUNT",
        value: 2,
    },
});
const fooDeploymentGroup = new aws.codedeploy.DeploymentGroup("fooDeploymentGroup", {
    appName: aws_codedeploy_app.foo_app.name,
    deploymentGroupName: "bar",
    serviceRoleArn: aws_iam_role.foo_role.arn,
    deploymentConfigName: fooDeploymentConfig.id,
    ec2TagFilters: [{
        key: "filterkey",
        type: "KEY_AND_VALUE",
        value: "filtervalue",
    }],
    triggerConfigurations: [{
        triggerEvents: ["DeploymentFailure"],
        triggerName: "foo-trigger",
        triggerTargetArn: "foo-topic-arn",
    }],
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_FAILURE"],
    },
    alarmConfiguration: {
        alarms: ["my-alarm-name"],
        enabled: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo_deployment_config = aws.codedeploy.DeploymentConfig("fooDeploymentConfig",
    deployment_config_name="test-deployment-config",
    minimum_healthy_hosts=aws.codedeploy.DeploymentConfigMinimumHealthyHostsArgs(
        type="HOST_COUNT",
        value=2,
    ))
foo_deployment_group = aws.codedeploy.DeploymentGroup("fooDeploymentGroup",
    app_name=aws_codedeploy_app["foo_app"]["name"],
    deployment_group_name="bar",
    service_role_arn=aws_iam_role["foo_role"]["arn"],
    deployment_config_name=foo_deployment_config.id,
    ec2_tag_filters=[aws.codedeploy.DeploymentGroupEc2TagFilterArgs(
        key="filterkey",
        type="KEY_AND_VALUE",
        value="filtervalue",
    )],
    trigger_configurations=[aws.codedeploy.DeploymentGroupTriggerConfigurationArgs(
        trigger_events=["DeploymentFailure"],
        trigger_name="foo-trigger",
        trigger_target_arn="foo-topic-arn",
    )],
    auto_rollback_configuration=aws.codedeploy.DeploymentGroupAutoRollbackConfigurationArgs(
        enabled=True,
        events=["DEPLOYMENT_FAILURE"],
    ),
    alarm_configuration=aws.codedeploy.DeploymentGroupAlarmConfigurationArgs(
        alarms=["my-alarm-name"],
        enabled=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooDeploymentConfig = new Aws.CodeDeploy.DeploymentConfig("fooDeploymentConfig", new()
    {
        DeploymentConfigName = "test-deployment-config",
        MinimumHealthyHosts = new Aws.CodeDeploy.Inputs.DeploymentConfigMinimumHealthyHostsArgs
        {
            Type = "HOST_COUNT",
            Value = 2,
        },
    });

    var fooDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("fooDeploymentGroup", new()
    {
        AppName = aws_codedeploy_app.Foo_app.Name,
        DeploymentGroupName = "bar",
        ServiceRoleArn = aws_iam_role.Foo_role.Arn,
        DeploymentConfigName = fooDeploymentConfig.Id,
        Ec2TagFilters = new[]
        {
            new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagFilterArgs
            {
                Key = "filterkey",
                Type = "KEY_AND_VALUE",
                Value = "filtervalue",
            },
        },
        TriggerConfigurations = new[]
        {
            new Aws.CodeDeploy.Inputs.DeploymentGroupTriggerConfigurationArgs
            {
                TriggerEvents = new[]
                {
                    "DeploymentFailure",
                },
                TriggerName = "foo-trigger",
                TriggerTargetArn = "foo-topic-arn",
            },
        },
        AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
        {
            Enabled = true,
            Events = new[]
            {
                "DEPLOYMENT_FAILURE",
            },
        },
        AlarmConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAlarmConfigurationArgs
        {
            Alarms = new[]
            {
                "my-alarm-name",
            },
            Enabled = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooDeploymentConfig, err := codedeploy.NewDeploymentConfig(ctx, "fooDeploymentConfig", &codedeploy.DeploymentConfigArgs{
			DeploymentConfigName: pulumi.String("test-deployment-config"),
			MinimumHealthyHosts: &codedeploy.DeploymentConfigMinimumHealthyHostsArgs{
				Type:  pulumi.String("HOST_COUNT"),
				Value: pulumi.Int(2),
			},
		})
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "fooDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AppName:              pulumi.Any(aws_codedeploy_app.Foo_app.Name),
			DeploymentGroupName:  pulumi.String("bar"),
			ServiceRoleArn:       pulumi.Any(aws_iam_role.Foo_role.Arn),
			DeploymentConfigName: fooDeploymentConfig.ID(),
			Ec2TagFilters: codedeploy.DeploymentGroupEc2TagFilterArray{
				&codedeploy.DeploymentGroupEc2TagFilterArgs{
					Key:   pulumi.String("filterkey"),
					Type:  pulumi.String("KEY_AND_VALUE"),
					Value: pulumi.String("filtervalue"),
				},
			},
			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
					TriggerEvents: pulumi.StringArray{
						pulumi.String("DeploymentFailure"),
					},
					TriggerName:      pulumi.String("foo-trigger"),
					TriggerTargetArn: pulumi.String("foo-topic-arn"),
				},
			},
			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
				Enabled: pulumi.Bool(true),
				Events: pulumi.StringArray{
					pulumi.String("DEPLOYMENT_FAILURE"),
				},
			},
			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
				Alarms: pulumi.StringArray{
					pulumi.String("my-alarm-name"),
				},
				Enabled: pulumi.Bool(true),
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
import com.pulumi.aws.codedeploy.DeploymentConfig;
import com.pulumi.aws.codedeploy.DeploymentConfigArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentConfigMinimumHealthyHostsArgs;
import com.pulumi.aws.codedeploy.DeploymentGroup;
import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupEc2TagFilterArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupTriggerConfigurationArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAutoRollbackConfigurationArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAlarmConfigurationArgs;
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
        var fooDeploymentConfig = new DeploymentConfig("fooDeploymentConfig", DeploymentConfigArgs.builder()        
            .deploymentConfigName("test-deployment-config")
            .minimumHealthyHosts(DeploymentConfigMinimumHealthyHostsArgs.builder()
                .type("HOST_COUNT")
                .value(2)
                .build())
            .build());

        var fooDeploymentGroup = new DeploymentGroup("fooDeploymentGroup", DeploymentGroupArgs.builder()        
            .appName(aws_codedeploy_app.foo_app().name())
            .deploymentGroupName("bar")
            .serviceRoleArn(aws_iam_role.foo_role().arn())
            .deploymentConfigName(fooDeploymentConfig.id())
            .ec2TagFilters(DeploymentGroupEc2TagFilterArgs.builder()
                .key("filterkey")
                .type("KEY_AND_VALUE")
                .value("filtervalue")
                .build())
            .triggerConfigurations(DeploymentGroupTriggerConfigurationArgs.builder()
                .triggerEvents("DeploymentFailure")
                .triggerName("foo-trigger")
                .triggerTargetArn("foo-topic-arn")
                .build())
            .autoRollbackConfiguration(DeploymentGroupAutoRollbackConfigurationArgs.builder()
                .enabled(true)
                .events("DEPLOYMENT_FAILURE")
                .build())
            .alarmConfiguration(DeploymentGroupAlarmConfigurationArgs.builder()
                .alarms("my-alarm-name")
                .enabled(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  fooDeploymentConfig:
    type: aws:codedeploy:DeploymentConfig
    properties:
      deploymentConfigName: test-deployment-config
      minimumHealthyHosts:
        type: HOST_COUNT
        value: 2
  fooDeploymentGroup:
    type: aws:codedeploy:DeploymentGroup
    properties:
      appName: ${aws_codedeploy_app.foo_app.name}
      deploymentGroupName: bar
      serviceRoleArn: ${aws_iam_role.foo_role.arn}
      deploymentConfigName: ${fooDeploymentConfig.id}
      ec2TagFilters:
        - key: filterkey
          type: KEY_AND_VALUE
          value: filtervalue
      triggerConfigurations:
        - triggerEvents:
            - DeploymentFailure
          triggerName: foo-trigger
          triggerTargetArn: foo-topic-arn
      autoRollbackConfiguration:
        enabled: true
        events:
          - DEPLOYMENT_FAILURE
      alarmConfiguration:
        alarms:
          - my-alarm-name
        enabled: true
```
{{% /example %}}
{{% example %}}
### Lambda Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooDeploymentConfig = new aws.codedeploy.DeploymentConfig("fooDeploymentConfig", {
    deploymentConfigName: "test-deployment-config",
    computePlatform: "Lambda",
    trafficRoutingConfig: {
        type: "TimeBasedLinear",
        timeBasedLinear: {
            interval: 10,
            percentage: 10,
        },
    },
});
const fooDeploymentGroup = new aws.codedeploy.DeploymentGroup("fooDeploymentGroup", {
    appName: aws_codedeploy_app.foo_app.name,
    deploymentGroupName: "bar",
    serviceRoleArn: aws_iam_role.foo_role.arn,
    deploymentConfigName: fooDeploymentConfig.id,
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_STOP_ON_ALARM"],
    },
    alarmConfiguration: {
        alarms: ["my-alarm-name"],
        enabled: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo_deployment_config = aws.codedeploy.DeploymentConfig("fooDeploymentConfig",
    deployment_config_name="test-deployment-config",
    compute_platform="Lambda",
    traffic_routing_config=aws.codedeploy.DeploymentConfigTrafficRoutingConfigArgs(
        type="TimeBasedLinear",
        time_based_linear=aws.codedeploy.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs(
            interval=10,
            percentage=10,
        ),
    ))
foo_deployment_group = aws.codedeploy.DeploymentGroup("fooDeploymentGroup",
    app_name=aws_codedeploy_app["foo_app"]["name"],
    deployment_group_name="bar",
    service_role_arn=aws_iam_role["foo_role"]["arn"],
    deployment_config_name=foo_deployment_config.id,
    auto_rollback_configuration=aws.codedeploy.DeploymentGroupAutoRollbackConfigurationArgs(
        enabled=True,
        events=["DEPLOYMENT_STOP_ON_ALARM"],
    ),
    alarm_configuration=aws.codedeploy.DeploymentGroupAlarmConfigurationArgs(
        alarms=["my-alarm-name"],
        enabled=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooDeploymentConfig = new Aws.CodeDeploy.DeploymentConfig("fooDeploymentConfig", new()
    {
        DeploymentConfigName = "test-deployment-config",
        ComputePlatform = "Lambda",
        TrafficRoutingConfig = new Aws.CodeDeploy.Inputs.DeploymentConfigTrafficRoutingConfigArgs
        {
            Type = "TimeBasedLinear",
            TimeBasedLinear = new Aws.CodeDeploy.Inputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs
            {
                Interval = 10,
                Percentage = 10,
            },
        },
    });

    var fooDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("fooDeploymentGroup", new()
    {
        AppName = aws_codedeploy_app.Foo_app.Name,
        DeploymentGroupName = "bar",
        ServiceRoleArn = aws_iam_role.Foo_role.Arn,
        DeploymentConfigName = fooDeploymentConfig.Id,
        AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
        {
            Enabled = true,
            Events = new[]
            {
                "DEPLOYMENT_STOP_ON_ALARM",
            },
        },
        AlarmConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAlarmConfigurationArgs
        {
            Alarms = new[]
            {
                "my-alarm-name",
            },
            Enabled = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codedeploy"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fooDeploymentConfig, err := codedeploy.NewDeploymentConfig(ctx, "fooDeploymentConfig", &codedeploy.DeploymentConfigArgs{
			DeploymentConfigName: pulumi.String("test-deployment-config"),
			ComputePlatform:      pulumi.String("Lambda"),
			TrafficRoutingConfig: &codedeploy.DeploymentConfigTrafficRoutingConfigArgs{
				Type: pulumi.String("TimeBasedLinear"),
				TimeBasedLinear: &codedeploy.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs{
					Interval:   pulumi.Int(10),
					Percentage: pulumi.Int(10),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "fooDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AppName:              pulumi.Any(aws_codedeploy_app.Foo_app.Name),
			DeploymentGroupName:  pulumi.String("bar"),
			ServiceRoleArn:       pulumi.Any(aws_iam_role.Foo_role.Arn),
			DeploymentConfigName: fooDeploymentConfig.ID(),
			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
				Enabled: pulumi.Bool(true),
				Events: pulumi.StringArray{
					pulumi.String("DEPLOYMENT_STOP_ON_ALARM"),
				},
			},
			AlarmConfiguration: &codedeploy.DeploymentGroupAlarmConfigurationArgs{
				Alarms: pulumi.StringArray{
					pulumi.String("my-alarm-name"),
				},
				Enabled: pulumi.Bool(true),
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
import com.pulumi.aws.codedeploy.DeploymentConfig;
import com.pulumi.aws.codedeploy.DeploymentConfigArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentConfigTrafficRoutingConfigArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs;
import com.pulumi.aws.codedeploy.DeploymentGroup;
import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAutoRollbackConfigurationArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAlarmConfigurationArgs;
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
        var fooDeploymentConfig = new DeploymentConfig("fooDeploymentConfig", DeploymentConfigArgs.builder()        
            .deploymentConfigName("test-deployment-config")
            .computePlatform("Lambda")
            .trafficRoutingConfig(DeploymentConfigTrafficRoutingConfigArgs.builder()
                .type("TimeBasedLinear")
                .timeBasedLinear(DeploymentConfigTrafficRoutingConfigTimeBasedLinearArgs.builder()
                    .interval(10)
                    .percentage(10)
                    .build())
                .build())
            .build());

        var fooDeploymentGroup = new DeploymentGroup("fooDeploymentGroup", DeploymentGroupArgs.builder()        
            .appName(aws_codedeploy_app.foo_app().name())
            .deploymentGroupName("bar")
            .serviceRoleArn(aws_iam_role.foo_role().arn())
            .deploymentConfigName(fooDeploymentConfig.id())
            .autoRollbackConfiguration(DeploymentGroupAutoRollbackConfigurationArgs.builder()
                .enabled(true)
                .events("DEPLOYMENT_STOP_ON_ALARM")
                .build())
            .alarmConfiguration(DeploymentGroupAlarmConfigurationArgs.builder()
                .alarms("my-alarm-name")
                .enabled(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  fooDeploymentConfig:
    type: aws:codedeploy:DeploymentConfig
    properties:
      deploymentConfigName: test-deployment-config
      computePlatform: Lambda
      trafficRoutingConfig:
        type: TimeBasedLinear
        timeBasedLinear:
          interval: 10
          percentage: 10
  fooDeploymentGroup:
    type: aws:codedeploy:DeploymentGroup
    properties:
      appName: ${aws_codedeploy_app.foo_app.name}
      deploymentGroupName: bar
      serviceRoleArn: ${aws_iam_role.foo_role.arn}
      deploymentConfigName: ${fooDeploymentConfig.id}
      autoRollbackConfiguration:
        enabled: true
        events:
          - DEPLOYMENT_STOP_ON_ALARM
      alarmConfiguration:
        alarms:
          - my-alarm-name
        enabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeDeploy Deployment Configurations using the `deployment_config_name`. For example:

```sh
 $ pulumi import aws:codedeploy/deploymentConfig:DeploymentConfig example my-deployment-config
```
 