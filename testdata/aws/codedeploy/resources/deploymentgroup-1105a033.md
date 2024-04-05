Provides a CodeDeploy Deployment Group for a CodeDeploy Application

> **NOTE on blue/green deployments:** When using `green_fleet_provisioning_option` with the `COPY_AUTO_SCALING_GROUP` action, CodeDeploy will create a new ASG with a different name. This ASG is _not_ managed by this provider and will conflict with existing configuration and state. You may want to use a different approach to managing deployments that involve multiple ASG, such as `DISCOVER_EXISTING` with separate blue and green ASG.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["codedeploy.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const aWSCodeDeployRole = new aws.iam.RolePolicyAttachment("aWSCodeDeployRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
    role: exampleRole.name,
});
const exampleApplication = new aws.codedeploy.Application("exampleApplication", {});
const exampleTopic = new aws.sns.Topic("exampleTopic", {});
const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("exampleDeploymentGroup", {
    appName: exampleApplication.name,
    deploymentGroupName: "example-group",
    serviceRoleArn: exampleRole.arn,
    ec2TagSets: [{
        ec2TagFilters: [
            {
                key: "filterkey1",
                type: "KEY_AND_VALUE",
                value: "filtervalue",
            },
            {
                key: "filterkey2",
                type: "KEY_AND_VALUE",
                value: "filtervalue",
            },
        ],
    }],
    triggerConfigurations: [{
        triggerEvents: ["DeploymentFailure"],
        triggerName: "example-trigger",
        triggerTargetArn: exampleTopic.arn,
    }],
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_FAILURE"],
    },
    alarmConfiguration: {
        alarms: ["my-alarm-name"],
        enabled: true,
    },
    outdatedInstancesStrategy: "UPDATE",
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["codedeploy.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
a_ws_code_deploy_role = aws.iam.RolePolicyAttachment("aWSCodeDeployRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
    role=example_role.name)
example_application = aws.codedeploy.Application("exampleApplication")
example_topic = aws.sns.Topic("exampleTopic")
example_deployment_group = aws.codedeploy.DeploymentGroup("exampleDeploymentGroup",
    app_name=example_application.name,
    deployment_group_name="example-group",
    service_role_arn=example_role.arn,
    ec2_tag_sets=[aws.codedeploy.DeploymentGroupEc2TagSetArgs(
        ec2_tag_filters=[
            aws.codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArgs(
                key="filterkey1",
                type="KEY_AND_VALUE",
                value="filtervalue",
            ),
            aws.codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArgs(
                key="filterkey2",
                type="KEY_AND_VALUE",
                value="filtervalue",
            ),
        ],
    )],
    trigger_configurations=[aws.codedeploy.DeploymentGroupTriggerConfigurationArgs(
        trigger_events=["DeploymentFailure"],
        trigger_name="example-trigger",
        trigger_target_arn=example_topic.arn,
    )],
    auto_rollback_configuration=aws.codedeploy.DeploymentGroupAutoRollbackConfigurationArgs(
        enabled=True,
        events=["DEPLOYMENT_FAILURE"],
    ),
    alarm_configuration=aws.codedeploy.DeploymentGroupAlarmConfigurationArgs(
        alarms=["my-alarm-name"],
        enabled=True,
    ),
    outdated_instances_strategy="UPDATE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "codedeploy.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var aWSCodeDeployRole = new Aws.Iam.RolePolicyAttachment("aWSCodeDeployRole", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole",
        Role = exampleRole.Name,
    });

    var exampleApplication = new Aws.CodeDeploy.Application("exampleApplication");

    var exampleTopic = new Aws.Sns.Topic("exampleTopic");

    var exampleDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("exampleDeploymentGroup", new()
    {
        AppName = exampleApplication.Name,
        DeploymentGroupName = "example-group",
        ServiceRoleArn = exampleRole.Arn,
        Ec2TagSets = new[]
        {
            new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagSetArgs
            {
                Ec2TagFilters = new[]
                {
                    new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagSetEc2TagFilterArgs
                    {
                        Key = "filterkey1",
                        Type = "KEY_AND_VALUE",
                        Value = "filtervalue",
                    },
                    new Aws.CodeDeploy.Inputs.DeploymentGroupEc2TagSetEc2TagFilterArgs
                    {
                        Key = "filterkey2",
                        Type = "KEY_AND_VALUE",
                        Value = "filtervalue",
                    },
                },
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
                TriggerName = "example-trigger",
                TriggerTargetArn = exampleTopic.Arn,
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
        OutdatedInstancesStrategy = "UPDATE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codedeploy"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"codedeploy.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "aWSCodeDeployRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole"),
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
		if err != nil {
			return err
		}
		exampleTopic, err := sns.NewTopic(ctx, "exampleTopic", nil)
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AppName:             exampleApplication.Name,
			DeploymentGroupName: pulumi.String("example-group"),
			ServiceRoleArn:      exampleRole.Arn,
			Ec2TagSets: codedeploy.DeploymentGroupEc2TagSetArray{
				&codedeploy.DeploymentGroupEc2TagSetArgs{
					Ec2TagFilters: codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArray{
						&codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArgs{
							Key:   pulumi.String("filterkey1"),
							Type:  pulumi.String("KEY_AND_VALUE"),
							Value: pulumi.String("filtervalue"),
						},
						&codedeploy.DeploymentGroupEc2TagSetEc2TagFilterArgs{
							Key:   pulumi.String("filterkey2"),
							Type:  pulumi.String("KEY_AND_VALUE"),
							Value: pulumi.String("filtervalue"),
						},
					},
				},
			},
			TriggerConfigurations: codedeploy.DeploymentGroupTriggerConfigurationArray{
				&codedeploy.DeploymentGroupTriggerConfigurationArgs{
					TriggerEvents: pulumi.StringArray{
						pulumi.String("DeploymentFailure"),
					},
					TriggerName:      pulumi.String("example-trigger"),
					TriggerTargetArn: exampleTopic.Arn,
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
			OutdatedInstancesStrategy: pulumi.String("UPDATE"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.codedeploy.Application;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.codedeploy.DeploymentGroup;
import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupEc2TagSetArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("codedeploy.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var aWSCodeDeployRole = new RolePolicyAttachment("aWSCodeDeployRole", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole")
            .role(exampleRole.name())
            .build());

        var exampleApplication = new Application("exampleApplication");

        var exampleTopic = new Topic("exampleTopic");

        var exampleDeploymentGroup = new DeploymentGroup("exampleDeploymentGroup", DeploymentGroupArgs.builder()        
            .appName(exampleApplication.name())
            .deploymentGroupName("example-group")
            .serviceRoleArn(exampleRole.arn())
            .ec2TagSets(DeploymentGroupEc2TagSetArgs.builder()
                .ec2TagFilters(                
                    DeploymentGroupEc2TagSetEc2TagFilterArgs.builder()
                        .key("filterkey1")
                        .type("KEY_AND_VALUE")
                        .value("filtervalue")
                        .build(),
                    DeploymentGroupEc2TagSetEc2TagFilterArgs.builder()
                        .key("filterkey2")
                        .type("KEY_AND_VALUE")
                        .value("filtervalue")
                        .build())
                .build())
            .triggerConfigurations(DeploymentGroupTriggerConfigurationArgs.builder()
                .triggerEvents("DeploymentFailure")
                .triggerName("example-trigger")
                .triggerTargetArn(exampleTopic.arn())
                .build())
            .autoRollbackConfiguration(DeploymentGroupAutoRollbackConfigurationArgs.builder()
                .enabled(true)
                .events("DEPLOYMENT_FAILURE")
                .build())
            .alarmConfiguration(DeploymentGroupAlarmConfigurationArgs.builder()
                .alarms("my-alarm-name")
                .enabled(true)
                .build())
            .outdatedInstancesStrategy("UPDATE")
            .build());

    }
}
```
```yaml
resources:
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  aWSCodeDeployRole:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/service-role/AWSCodeDeployRole
      role: ${exampleRole.name}
  exampleApplication:
    type: aws:codedeploy:Application
  exampleTopic:
    type: aws:sns:Topic
  exampleDeploymentGroup:
    type: aws:codedeploy:DeploymentGroup
    properties:
      appName: ${exampleApplication.name}
      deploymentGroupName: example-group
      serviceRoleArn: ${exampleRole.arn}
      ec2TagSets:
        - ec2TagFilters:
            - key: filterkey1
              type: KEY_AND_VALUE
              value: filtervalue
            - key: filterkey2
              type: KEY_AND_VALUE
              value: filtervalue
      triggerConfigurations:
        - triggerEvents:
            - DeploymentFailure
          triggerName: example-trigger
          triggerTargetArn: ${exampleTopic.arn}
      autoRollbackConfiguration:
        enabled: true
        events:
          - DEPLOYMENT_FAILURE
      alarmConfiguration:
        alarms:
          - my-alarm-name
        enabled: true
      outdatedInstancesStrategy: UPDATE
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - codedeploy.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Blue Green Deployments with ECS

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApplication = new aws.codedeploy.Application("exampleApplication", {computePlatform: "ECS"});
const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("exampleDeploymentGroup", {
    appName: exampleApplication.name,
    deploymentConfigName: "CodeDeployDefault.ECSAllAtOnce",
    deploymentGroupName: "example",
    serviceRoleArn: aws_iam_role.example.arn,
    autoRollbackConfiguration: {
        enabled: true,
        events: ["DEPLOYMENT_FAILURE"],
    },
    blueGreenDeploymentConfig: {
        deploymentReadyOption: {
            actionOnTimeout: "CONTINUE_DEPLOYMENT",
        },
        terminateBlueInstancesOnDeploymentSuccess: {
            action: "TERMINATE",
            terminationWaitTimeInMinutes: 5,
        },
    },
    deploymentStyle: {
        deploymentOption: "WITH_TRAFFIC_CONTROL",
        deploymentType: "BLUE_GREEN",
    },
    ecsService: {
        clusterName: aws_ecs_cluster.example.name,
        serviceName: aws_ecs_service.example.name,
    },
    loadBalancerInfo: {
        targetGroupPairInfo: {
            prodTrafficRoute: {
                listenerArns: [aws_lb_listener.example.arn],
            },
            targetGroups: [
                {
                    name: aws_lb_target_group.blue.name,
                },
                {
                    name: aws_lb_target_group.green.name,
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_application = aws.codedeploy.Application("exampleApplication", compute_platform="ECS")
example_deployment_group = aws.codedeploy.DeploymentGroup("exampleDeploymentGroup",
    app_name=example_application.name,
    deployment_config_name="CodeDeployDefault.ECSAllAtOnce",
    deployment_group_name="example",
    service_role_arn=aws_iam_role["example"]["arn"],
    auto_rollback_configuration=aws.codedeploy.DeploymentGroupAutoRollbackConfigurationArgs(
        enabled=True,
        events=["DEPLOYMENT_FAILURE"],
    ),
    blue_green_deployment_config=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs(
        deployment_ready_option=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs(
            action_on_timeout="CONTINUE_DEPLOYMENT",
        ),
        terminate_blue_instances_on_deployment_success=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs(
            action="TERMINATE",
            termination_wait_time_in_minutes=5,
        ),
    ),
    deployment_style=aws.codedeploy.DeploymentGroupDeploymentStyleArgs(
        deployment_option="WITH_TRAFFIC_CONTROL",
        deployment_type="BLUE_GREEN",
    ),
    ecs_service=aws.codedeploy.DeploymentGroupEcsServiceArgs(
        cluster_name=aws_ecs_cluster["example"]["name"],
        service_name=aws_ecs_service["example"]["name"],
    ),
    load_balancer_info=aws.codedeploy.DeploymentGroupLoadBalancerInfoArgs(
        target_group_pair_info=aws.codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs(
            prod_traffic_route=aws.codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs(
                listener_arns=[aws_lb_listener["example"]["arn"]],
            ),
            target_groups=[
                aws.codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs(
                    name=aws_lb_target_group["blue"]["name"],
                ),
                aws.codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs(
                    name=aws_lb_target_group["green"]["name"],
                ),
            ],
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
    var exampleApplication = new Aws.CodeDeploy.Application("exampleApplication", new()
    {
        ComputePlatform = "ECS",
    });

    var exampleDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("exampleDeploymentGroup", new()
    {
        AppName = exampleApplication.Name,
        DeploymentConfigName = "CodeDeployDefault.ECSAllAtOnce",
        DeploymentGroupName = "example",
        ServiceRoleArn = aws_iam_role.Example.Arn,
        AutoRollbackConfiguration = new Aws.CodeDeploy.Inputs.DeploymentGroupAutoRollbackConfigurationArgs
        {
            Enabled = true,
            Events = new[]
            {
                "DEPLOYMENT_FAILURE",
            },
        },
        BlueGreenDeploymentConfig = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigArgs
        {
            DeploymentReadyOption = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs
            {
                ActionOnTimeout = "CONTINUE_DEPLOYMENT",
            },
            TerminateBlueInstancesOnDeploymentSuccess = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs
            {
                Action = "TERMINATE",
                TerminationWaitTimeInMinutes = 5,
            },
        },
        DeploymentStyle = new Aws.CodeDeploy.Inputs.DeploymentGroupDeploymentStyleArgs
        {
            DeploymentOption = "WITH_TRAFFIC_CONTROL",
            DeploymentType = "BLUE_GREEN",
        },
        EcsService = new Aws.CodeDeploy.Inputs.DeploymentGroupEcsServiceArgs
        {
            ClusterName = aws_ecs_cluster.Example.Name,
            ServiceName = aws_ecs_service.Example.Name,
        },
        LoadBalancerInfo = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoArgs
        {
            TargetGroupPairInfo = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs
            {
                ProdTrafficRoute = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs
                {
                    ListenerArns = new[]
                    {
                        aws_lb_listener.Example.Arn,
                    },
                },
                TargetGroups = new[]
                {
                    new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs
                    {
                        Name = aws_lb_target_group.Blue.Name,
                    },
                    new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs
                    {
                        Name = aws_lb_target_group.Green.Name,
                    },
                },
            },
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
		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", &codedeploy.ApplicationArgs{
			ComputePlatform: pulumi.String("ECS"),
		})
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AppName:              exampleApplication.Name,
			DeploymentConfigName: pulumi.String("CodeDeployDefault.ECSAllAtOnce"),
			DeploymentGroupName:  pulumi.String("example"),
			ServiceRoleArn:       pulumi.Any(aws_iam_role.Example.Arn),
			AutoRollbackConfiguration: &codedeploy.DeploymentGroupAutoRollbackConfigurationArgs{
				Enabled: pulumi.Bool(true),
				Events: pulumi.StringArray{
					pulumi.String("DEPLOYMENT_FAILURE"),
				},
			},
			BlueGreenDeploymentConfig: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs{
				DeploymentReadyOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs{
					ActionOnTimeout: pulumi.String("CONTINUE_DEPLOYMENT"),
				},
				TerminateBlueInstancesOnDeploymentSuccess: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs{
					Action:                       pulumi.String("TERMINATE"),
					TerminationWaitTimeInMinutes: pulumi.Int(5),
				},
			},
			DeploymentStyle: &codedeploy.DeploymentGroupDeploymentStyleArgs{
				DeploymentOption: pulumi.String("WITH_TRAFFIC_CONTROL"),
				DeploymentType:   pulumi.String("BLUE_GREEN"),
			},
			EcsService: &codedeploy.DeploymentGroupEcsServiceArgs{
				ClusterName: pulumi.Any(aws_ecs_cluster.Example.Name),
				ServiceName: pulumi.Any(aws_ecs_service.Example.Name),
			},
			LoadBalancerInfo: &codedeploy.DeploymentGroupLoadBalancerInfoArgs{
				TargetGroupPairInfo: &codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs{
					ProdTrafficRoute: &codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs{
						ListenerArns: pulumi.StringArray{
							aws_lb_listener.Example.Arn,
						},
					},
					TargetGroups: codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArray{
						&codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs{
							Name: pulumi.Any(aws_lb_target_group.Blue.Name),
						},
						&codedeploy.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs{
							Name: pulumi.Any(aws_lb_target_group.Green.Name),
						},
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
import com.pulumi.aws.codedeploy.Application;
import com.pulumi.aws.codedeploy.ApplicationArgs;
import com.pulumi.aws.codedeploy.DeploymentGroup;
import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupAutoRollbackConfigurationArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupDeploymentStyleArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupEcsServiceArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupLoadBalancerInfoArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs;
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
            .computePlatform("ECS")
            .build());

        var exampleDeploymentGroup = new DeploymentGroup("exampleDeploymentGroup", DeploymentGroupArgs.builder()        
            .appName(exampleApplication.name())
            .deploymentConfigName("CodeDeployDefault.ECSAllAtOnce")
            .deploymentGroupName("example")
            .serviceRoleArn(aws_iam_role.example().arn())
            .autoRollbackConfiguration(DeploymentGroupAutoRollbackConfigurationArgs.builder()
                .enabled(true)
                .events("DEPLOYMENT_FAILURE")
                .build())
            .blueGreenDeploymentConfig(DeploymentGroupBlueGreenDeploymentConfigArgs.builder()
                .deploymentReadyOption(DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs.builder()
                    .actionOnTimeout("CONTINUE_DEPLOYMENT")
                    .build())
                .terminateBlueInstancesOnDeploymentSuccess(DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs.builder()
                    .action("TERMINATE")
                    .terminationWaitTimeInMinutes(5)
                    .build())
                .build())
            .deploymentStyle(DeploymentGroupDeploymentStyleArgs.builder()
                .deploymentOption("WITH_TRAFFIC_CONTROL")
                .deploymentType("BLUE_GREEN")
                .build())
            .ecsService(DeploymentGroupEcsServiceArgs.builder()
                .clusterName(aws_ecs_cluster.example().name())
                .serviceName(aws_ecs_service.example().name())
                .build())
            .loadBalancerInfo(DeploymentGroupLoadBalancerInfoArgs.builder()
                .targetGroupPairInfo(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoArgs.builder()
                    .prodTrafficRoute(DeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRouteArgs.builder()
                        .listenerArns(aws_lb_listener.example().arn())
                        .build())
                    .targetGroups(                    
                        DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs.builder()
                            .name(aws_lb_target_group.blue().name())
                            .build(),
                        DeploymentGroupLoadBalancerInfoTargetGroupPairInfoTargetGroupArgs.builder()
                            .name(aws_lb_target_group.green().name())
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleApplication:
    type: aws:codedeploy:Application
    properties:
      computePlatform: ECS
  exampleDeploymentGroup:
    type: aws:codedeploy:DeploymentGroup
    properties:
      appName: ${exampleApplication.name}
      deploymentConfigName: CodeDeployDefault.ECSAllAtOnce
      deploymentGroupName: example
      serviceRoleArn: ${aws_iam_role.example.arn}
      autoRollbackConfiguration:
        enabled: true
        events:
          - DEPLOYMENT_FAILURE
      blueGreenDeploymentConfig:
        deploymentReadyOption:
          actionOnTimeout: CONTINUE_DEPLOYMENT
        terminateBlueInstancesOnDeploymentSuccess:
          action: TERMINATE
          terminationWaitTimeInMinutes: 5
      deploymentStyle:
        deploymentOption: WITH_TRAFFIC_CONTROL
        deploymentType: BLUE_GREEN
      ecsService:
        clusterName: ${aws_ecs_cluster.example.name}
        serviceName: ${aws_ecs_service.example.name}
      loadBalancerInfo:
        targetGroupPairInfo:
          prodTrafficRoute:
            listenerArns:
              - ${aws_lb_listener.example.arn}
          targetGroups:
            - name: ${aws_lb_target_group.blue.name}
            - name: ${aws_lb_target_group.green.name}
```
{{% /example %}}
{{% example %}}
### Blue Green Deployments with Servers and Classic ELB

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApplication = new aws.codedeploy.Application("exampleApplication", {});
const exampleDeploymentGroup = new aws.codedeploy.DeploymentGroup("exampleDeploymentGroup", {
    appName: exampleApplication.name,
    deploymentGroupName: "example-group",
    serviceRoleArn: aws_iam_role.example.arn,
    deploymentStyle: {
        deploymentOption: "WITH_TRAFFIC_CONTROL",
        deploymentType: "BLUE_GREEN",
    },
    loadBalancerInfo: {
        elbInfos: [{
            name: aws_elb.example.name,
        }],
    },
    blueGreenDeploymentConfig: {
        deploymentReadyOption: {
            actionOnTimeout: "STOP_DEPLOYMENT",
            waitTimeInMinutes: 60,
        },
        greenFleetProvisioningOption: {
            action: "DISCOVER_EXISTING",
        },
        terminateBlueInstancesOnDeploymentSuccess: {
            action: "KEEP_ALIVE",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_application = aws.codedeploy.Application("exampleApplication")
example_deployment_group = aws.codedeploy.DeploymentGroup("exampleDeploymentGroup",
    app_name=example_application.name,
    deployment_group_name="example-group",
    service_role_arn=aws_iam_role["example"]["arn"],
    deployment_style=aws.codedeploy.DeploymentGroupDeploymentStyleArgs(
        deployment_option="WITH_TRAFFIC_CONTROL",
        deployment_type="BLUE_GREEN",
    ),
    load_balancer_info=aws.codedeploy.DeploymentGroupLoadBalancerInfoArgs(
        elb_infos=[aws.codedeploy.DeploymentGroupLoadBalancerInfoElbInfoArgs(
            name=aws_elb["example"]["name"],
        )],
    ),
    blue_green_deployment_config=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs(
        deployment_ready_option=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs(
            action_on_timeout="STOP_DEPLOYMENT",
            wait_time_in_minutes=60,
        ),
        green_fleet_provisioning_option=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs(
            action="DISCOVER_EXISTING",
        ),
        terminate_blue_instances_on_deployment_success=aws.codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs(
            action="KEEP_ALIVE",
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
    var exampleApplication = new Aws.CodeDeploy.Application("exampleApplication");

    var exampleDeploymentGroup = new Aws.CodeDeploy.DeploymentGroup("exampleDeploymentGroup", new()
    {
        AppName = exampleApplication.Name,
        DeploymentGroupName = "example-group",
        ServiceRoleArn = aws_iam_role.Example.Arn,
        DeploymentStyle = new Aws.CodeDeploy.Inputs.DeploymentGroupDeploymentStyleArgs
        {
            DeploymentOption = "WITH_TRAFFIC_CONTROL",
            DeploymentType = "BLUE_GREEN",
        },
        LoadBalancerInfo = new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoArgs
        {
            ElbInfos = new[]
            {
                new Aws.CodeDeploy.Inputs.DeploymentGroupLoadBalancerInfoElbInfoArgs
                {
                    Name = aws_elb.Example.Name,
                },
            },
        },
        BlueGreenDeploymentConfig = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigArgs
        {
            DeploymentReadyOption = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs
            {
                ActionOnTimeout = "STOP_DEPLOYMENT",
                WaitTimeInMinutes = 60,
            },
            GreenFleetProvisioningOption = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs
            {
                Action = "DISCOVER_EXISTING",
            },
            TerminateBlueInstancesOnDeploymentSuccess = new Aws.CodeDeploy.Inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs
            {
                Action = "KEEP_ALIVE",
            },
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
		exampleApplication, err := codedeploy.NewApplication(ctx, "exampleApplication", nil)
		if err != nil {
			return err
		}
		_, err = codedeploy.NewDeploymentGroup(ctx, "exampleDeploymentGroup", &codedeploy.DeploymentGroupArgs{
			AppName:             exampleApplication.Name,
			DeploymentGroupName: pulumi.String("example-group"),
			ServiceRoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
			DeploymentStyle: &codedeploy.DeploymentGroupDeploymentStyleArgs{
				DeploymentOption: pulumi.String("WITH_TRAFFIC_CONTROL"),
				DeploymentType:   pulumi.String("BLUE_GREEN"),
			},
			LoadBalancerInfo: &codedeploy.DeploymentGroupLoadBalancerInfoArgs{
				ElbInfos: codedeploy.DeploymentGroupLoadBalancerInfoElbInfoArray{
					&codedeploy.DeploymentGroupLoadBalancerInfoElbInfoArgs{
						Name: pulumi.Any(aws_elb.Example.Name),
					},
				},
			},
			BlueGreenDeploymentConfig: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigArgs{
				DeploymentReadyOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs{
					ActionOnTimeout:   pulumi.String("STOP_DEPLOYMENT"),
					WaitTimeInMinutes: pulumi.Int(60),
				},
				GreenFleetProvisioningOption: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs{
					Action: pulumi.String("DISCOVER_EXISTING"),
				},
				TerminateBlueInstancesOnDeploymentSuccess: &codedeploy.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs{
					Action: pulumi.String("KEEP_ALIVE"),
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
import com.pulumi.aws.codedeploy.Application;
import com.pulumi.aws.codedeploy.DeploymentGroup;
import com.pulumi.aws.codedeploy.DeploymentGroupArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupDeploymentStyleArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupLoadBalancerInfoArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs;
import com.pulumi.aws.codedeploy.inputs.DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs;
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
        var exampleApplication = new Application("exampleApplication");

        var exampleDeploymentGroup = new DeploymentGroup("exampleDeploymentGroup", DeploymentGroupArgs.builder()        
            .appName(exampleApplication.name())
            .deploymentGroupName("example-group")
            .serviceRoleArn(aws_iam_role.example().arn())
            .deploymentStyle(DeploymentGroupDeploymentStyleArgs.builder()
                .deploymentOption("WITH_TRAFFIC_CONTROL")
                .deploymentType("BLUE_GREEN")
                .build())
            .loadBalancerInfo(DeploymentGroupLoadBalancerInfoArgs.builder()
                .elbInfos(DeploymentGroupLoadBalancerInfoElbInfoArgs.builder()
                    .name(aws_elb.example().name())
                    .build())
                .build())
            .blueGreenDeploymentConfig(DeploymentGroupBlueGreenDeploymentConfigArgs.builder()
                .deploymentReadyOption(DeploymentGroupBlueGreenDeploymentConfigDeploymentReadyOptionArgs.builder()
                    .actionOnTimeout("STOP_DEPLOYMENT")
                    .waitTimeInMinutes(60)
                    .build())
                .greenFleetProvisioningOption(DeploymentGroupBlueGreenDeploymentConfigGreenFleetProvisioningOptionArgs.builder()
                    .action("DISCOVER_EXISTING")
                    .build())
                .terminateBlueInstancesOnDeploymentSuccess(DeploymentGroupBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccessArgs.builder()
                    .action("KEEP_ALIVE")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleApplication:
    type: aws:codedeploy:Application
  exampleDeploymentGroup:
    type: aws:codedeploy:DeploymentGroup
    properties:
      appName: ${exampleApplication.name}
      deploymentGroupName: example-group
      serviceRoleArn: ${aws_iam_role.example.arn}
      deploymentStyle:
        deploymentOption: WITH_TRAFFIC_CONTROL
        deploymentType: BLUE_GREEN
      loadBalancerInfo:
        elbInfos:
          - name: ${aws_elb.example.name}
      blueGreenDeploymentConfig:
        deploymentReadyOption:
          actionOnTimeout: STOP_DEPLOYMENT
          waitTimeInMinutes: 60
        greenFleetProvisioningOption:
          action: DISCOVER_EXISTING
        terminateBlueInstancesOnDeploymentSuccess:
          action: KEEP_ALIVE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeDeploy Deployment Groups using `app_name`, a colon, and `deployment_group_name`. For example:

```sh
 $ pulumi import aws:codedeploy/deploymentGroup:DeploymentGroup example my-application:my-deployment-group
```
 