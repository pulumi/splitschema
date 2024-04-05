Provides an EventBridge Target resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}
### Kinesis Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const console = new aws.cloudwatch.EventRule("console", {
    description: "Capture all EC2 scaling events",
    eventPattern: JSON.stringify({
        source: ["aws.autoscaling"],
        "detail-type": [
            "EC2 Instance Launch Successful",
            "EC2 Instance Terminate Successful",
            "EC2 Instance Launch Unsuccessful",
            "EC2 Instance Terminate Unsuccessful",
        ],
    }),
});
const testStream = new aws.kinesis.Stream("testStream", {shardCount: 1});
const yada = new aws.cloudwatch.EventTarget("yada", {
    rule: console.name,
    arn: testStream.arn,
    runCommandTargets: [
        {
            key: "tag:Name",
            values: ["FooBar"],
        },
        {
            key: "InstanceIds",
            values: ["i-162058cd308bffec2"],
        },
    ],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

console = aws.cloudwatch.EventRule("console",
    description="Capture all EC2 scaling events",
    event_pattern=json.dumps({
        "source": ["aws.autoscaling"],
        "detail-type": [
            "EC2 Instance Launch Successful",
            "EC2 Instance Terminate Successful",
            "EC2 Instance Launch Unsuccessful",
            "EC2 Instance Terminate Unsuccessful",
        ],
    }))
test_stream = aws.kinesis.Stream("testStream", shard_count=1)
yada = aws.cloudwatch.EventTarget("yada",
    rule=console.name,
    arn=test_stream.arn,
    run_command_targets=[
        aws.cloudwatch.EventTargetRunCommandTargetArgs(
            key="tag:Name",
            values=["FooBar"],
        ),
        aws.cloudwatch.EventTargetRunCommandTargetArgs(
            key="InstanceIds",
            values=["i-162058cd308bffec2"],
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var console = new Aws.CloudWatch.EventRule("console", new()
    {
        Description = "Capture all EC2 scaling events",
        EventPattern = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["source"] = new[]
            {
                "aws.autoscaling",
            },
            ["detail-type"] = new[]
            {
                "EC2 Instance Launch Successful",
                "EC2 Instance Terminate Successful",
                "EC2 Instance Launch Unsuccessful",
                "EC2 Instance Terminate Unsuccessful",
            },
        }),
    });

    var testStream = new Aws.Kinesis.Stream("testStream", new()
    {
        ShardCount = 1,
    });

    var yada = new Aws.CloudWatch.EventTarget("yada", new()
    {
        Rule = console.Name,
        Arn = testStream.Arn,
        RunCommandTargets = new[]
        {
            new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
            {
                Key = "tag:Name",
                Values = new[]
                {
                    "FooBar",
                },
            },
            new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    "i-162058cd308bffec2",
                },
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"source": []string{
				"aws.autoscaling",
			},
			"detail-type": []string{
				"EC2 Instance Launch Successful",
				"EC2 Instance Terminate Successful",
				"EC2 Instance Launch Unsuccessful",
				"EC2 Instance Terminate Unsuccessful",
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		console, err := cloudwatch.NewEventRule(ctx, "console", &cloudwatch.EventRuleArgs{
			Description:  pulumi.String("Capture all EC2 scaling events"),
			EventPattern: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "yada", &cloudwatch.EventTargetArgs{
			Rule: console.Name,
			Arn:  testStream.Arn,
			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
				&cloudwatch.EventTargetRunCommandTargetArgs{
					Key: pulumi.String("tag:Name"),
					Values: pulumi.StringArray{
						pulumi.String("FooBar"),
					},
				},
				&cloudwatch.EventTargetRunCommandTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						pulumi.String("i-162058cd308bffec2"),
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
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventRuleArgs;
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.cloudwatch.inputs.EventTargetRunCommandTargetArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var console = new EventRule("console", EventRuleArgs.builder()        
            .description("Capture all EC2 scaling events")
            .eventPattern(serializeJson(
                jsonObject(
                    jsonProperty("source", jsonArray("aws.autoscaling")),
                    jsonProperty("detail-type", jsonArray(
                        "EC2 Instance Launch Successful", 
                        "EC2 Instance Terminate Successful", 
                        "EC2 Instance Launch Unsuccessful", 
                        "EC2 Instance Terminate Unsuccessful"
                    ))
                )))
            .build());

        var testStream = new Stream("testStream", StreamArgs.builder()        
            .shardCount(1)
            .build());

        var yada = new EventTarget("yada", EventTargetArgs.builder()        
            .rule(console.name())
            .arn(testStream.arn())
            .runCommandTargets(            
                EventTargetRunCommandTargetArgs.builder()
                    .key("tag:Name")
                    .values("FooBar")
                    .build(),
                EventTargetRunCommandTargetArgs.builder()
                    .key("InstanceIds")
                    .values("i-162058cd308bffec2")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  yada:
    type: aws:cloudwatch:EventTarget
    properties:
      rule: ${console.name}
      arn: ${testStream.arn}
      runCommandTargets:
        - key: tag:Name
          values:
            - FooBar
        - key: InstanceIds
          values:
            - i-162058cd308bffec2
  console:
    type: aws:cloudwatch:EventRule
    properties:
      description: Capture all EC2 scaling events
      eventPattern:
        fn::toJSON:
          source:
            - aws.autoscaling
          detail-type:
            - EC2 Instance Launch Successful
            - EC2 Instance Terminate Successful
            - EC2 Instance Launch Unsuccessful
            - EC2 Instance Terminate Unsuccessful
  testStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 1
```
{{% /example %}}
{{% example %}}
### SSM Document Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ssmLifecycleTrust = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["events.amazonaws.com"],
        }],
    }],
});
const stopInstance = new aws.ssm.Document("stopInstance", {
    documentType: "Command",
    content: JSON.stringify({
        schemaVersion: "1.2",
        description: "Stop an instance",
        parameters: {},
        runtimeConfig: {
            "aws:runShellScript": {
                properties: [{
                    id: "0.aws:runShellScript",
                    runCommand: ["halt"],
                }],
            },
        },
    }),
});
const ssmLifecyclePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: ["ssm:SendCommand"],
            resources: ["arn:aws:ec2:eu-west-1:1234567890:instance/*"],
            conditions: [{
                test: "StringEquals",
                variable: "ec2:ResourceTag/Terminate",
                values: ["*"],
            }],
        },
        {
            effect: "Allow",
            actions: ["ssm:SendCommand"],
            resources: [stopInstance.arn],
        },
    ],
});
const ssmLifecycleRole = new aws.iam.Role("ssmLifecycleRole", {assumeRolePolicy: ssmLifecycleTrust.then(ssmLifecycleTrust => ssmLifecycleTrust.json)});
const ssmLifecyclePolicy = new aws.iam.Policy("ssmLifecyclePolicy", {policy: ssmLifecyclePolicyDocument.apply(ssmLifecyclePolicyDocument => ssmLifecyclePolicyDocument.json)});
const ssmLifecycleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("ssmLifecycleRolePolicyAttachment", {
    policyArn: ssmLifecyclePolicy.arn,
    role: ssmLifecycleRole.name,
});
const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
    description: "Stop instances nightly",
    scheduleExpression: "cron(0 0 * * ? *)",
});
const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
    arn: stopInstance.arn,
    rule: stopInstancesEventRule.name,
    roleArn: ssmLifecycleRole.arn,
    runCommandTargets: [{
        key: "tag:Terminate",
        values: ["midnight"],
    }],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

ssm_lifecycle_trust = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["events.amazonaws.com"],
    )],
)])
stop_instance = aws.ssm.Document("stopInstance",
    document_type="Command",
    content=json.dumps({
        "schemaVersion": "1.2",
        "description": "Stop an instance",
        "parameters": {},
        "runtimeConfig": {
            "aws:runShellScript": {
                "properties": [{
                    "id": "0.aws:runShellScript",
                    "runCommand": ["halt"],
                }],
            },
        },
    }))
ssm_lifecycle_policy_document = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["ssm:SendCommand"],
        resources=["arn:aws:ec2:eu-west-1:1234567890:instance/*"],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="ec2:ResourceTag/Terminate",
            values=["*"],
        )],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["ssm:SendCommand"],
        resources=[stop_instance.arn],
    ),
])
ssm_lifecycle_role = aws.iam.Role("ssmLifecycleRole", assume_role_policy=ssm_lifecycle_trust.json)
ssm_lifecycle_policy = aws.iam.Policy("ssmLifecyclePolicy", policy=ssm_lifecycle_policy_document.json)
ssm_lifecycle_role_policy_attachment = aws.iam.RolePolicyAttachment("ssmLifecycleRolePolicyAttachment",
    policy_arn=ssm_lifecycle_policy.arn,
    role=ssm_lifecycle_role.name)
stop_instances_event_rule = aws.cloudwatch.EventRule("stopInstancesEventRule",
    description="Stop instances nightly",
    schedule_expression="cron(0 0 * * ? *)")
stop_instances_event_target = aws.cloudwatch.EventTarget("stopInstancesEventTarget",
    arn=stop_instance.arn,
    rule=stop_instances_event_rule.name,
    role_arn=ssm_lifecycle_role.arn,
    run_command_targets=[aws.cloudwatch.EventTargetRunCommandTargetArgs(
        key="tag:Terminate",
        values=["midnight"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ssmLifecycleTrust = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "events.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var stopInstance = new Aws.Ssm.Document("stopInstance", new()
    {
        DocumentType = "Command",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["schemaVersion"] = "1.2",
            ["description"] = "Stop an instance",
            ["parameters"] = new Dictionary<string, object?>
            {
            },
            ["runtimeConfig"] = new Dictionary<string, object?>
            {
                ["aws:runShellScript"] = new Dictionary<string, object?>
                {
                    ["properties"] = new[]
                    {
                        new Dictionary<string, object?>
                        {
                            ["id"] = "0.aws:runShellScript",
                            ["runCommand"] = new[]
                            {
                                "halt",
                            },
                        },
                    },
                },
            },
        }),
    });

    var ssmLifecyclePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "ssm:SendCommand",
                },
                Resources = new[]
                {
                    "arn:aws:ec2:eu-west-1:1234567890:instance/*",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "ec2:ResourceTag/Terminate",
                        Values = new[]
                        {
                            "*",
                        },
                    },
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "ssm:SendCommand",
                },
                Resources = new[]
                {
                    stopInstance.Arn,
                },
            },
        },
    });

    var ssmLifecycleRole = new Aws.Iam.Role("ssmLifecycleRole", new()
    {
        AssumeRolePolicy = ssmLifecycleTrust.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var ssmLifecyclePolicy = new Aws.Iam.Policy("ssmLifecyclePolicy", new()
    {
        PolicyDocument = ssmLifecyclePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var ssmLifecycleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("ssmLifecycleRolePolicyAttachment", new()
    {
        PolicyArn = ssmLifecyclePolicy.Arn,
        Role = ssmLifecycleRole.Name,
    });

    var stopInstancesEventRule = new Aws.CloudWatch.EventRule("stopInstancesEventRule", new()
    {
        Description = "Stop instances nightly",
        ScheduleExpression = "cron(0 0 * * ? *)",
    });

    var stopInstancesEventTarget = new Aws.CloudWatch.EventTarget("stopInstancesEventTarget", new()
    {
        Arn = stopInstance.Arn,
        Rule = stopInstancesEventRule.Name,
        RoleArn = ssmLifecycleRole.Arn,
        RunCommandTargets = new[]
        {
            new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
            {
                Key = "tag:Terminate",
                Values = new[]
                {
                    "midnight",
                },
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ssmLifecycleTrust, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"events.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"schemaVersion": "1.2",
			"description":   "Stop an instance",
			"parameters":    nil,
			"runtimeConfig": map[string]interface{}{
				"aws:runShellScript": map[string]interface{}{
					"properties": []map[string]interface{}{
						map[string]interface{}{
							"id": "0.aws:runShellScript",
							"runCommand": []string{
								"halt",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		stopInstance, err := ssm.NewDocument(ctx, "stopInstance", &ssm.DocumentArgs{
			DocumentType: pulumi.String("Command"),
			Content:      pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		ssmLifecyclePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("ssm:SendCommand"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("arn:aws:ec2:eu-west-1:1234567890:instance/*"),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("StringEquals"),
							Variable: pulumi.String("ec2:ResourceTag/Terminate"),
							Values: pulumi.StringArray{
								pulumi.String("*"),
							},
						},
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("ssm:SendCommand"),
					},
					Resources: pulumi.StringArray{
						stopInstance.Arn,
					},
				},
			},
		}, nil)
		ssmLifecycleRole, err := iam.NewRole(ctx, "ssmLifecycleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(ssmLifecycleTrust.Json),
		})
		if err != nil {
			return err
		}
		ssmLifecyclePolicy, err := iam.NewPolicy(ctx, "ssmLifecyclePolicy", &iam.PolicyArgs{
			Policy: ssmLifecyclePolicyDocument.ApplyT(func(ssmLifecyclePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &ssmLifecyclePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "ssmLifecycleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: ssmLifecyclePolicy.Arn,
			Role:      ssmLifecycleRole.Name,
		})
		if err != nil {
			return err
		}
		stopInstancesEventRule, err := cloudwatch.NewEventRule(ctx, "stopInstancesEventRule", &cloudwatch.EventRuleArgs{
			Description:        pulumi.String("Stop instances nightly"),
			ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "stopInstancesEventTarget", &cloudwatch.EventTargetArgs{
			Arn:     stopInstance.Arn,
			Rule:    stopInstancesEventRule.Name,
			RoleArn: ssmLifecycleRole.Arn,
			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
				&cloudwatch.EventTargetRunCommandTargetArgs{
					Key: pulumi.String("tag:Terminate"),
					Values: pulumi.StringArray{
						pulumi.String("midnight"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.ssm.Document;
import com.pulumi.aws.ssm.DocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventRuleArgs;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.cloudwatch.inputs.EventTargetRunCommandTargetArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        final var ssmLifecycleTrust = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("events.amazonaws.com")
                    .build())
                .build())
            .build());

        var stopInstance = new Document("stopInstance", DocumentArgs.builder()        
            .documentType("Command")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("schemaVersion", "1.2"),
                    jsonProperty("description", "Stop an instance"),
                    jsonProperty("parameters", jsonObject(

                    )),
                    jsonProperty("runtimeConfig", jsonObject(
                        jsonProperty("aws:runShellScript", jsonObject(
                            jsonProperty("properties", jsonArray(jsonObject(
                                jsonProperty("id", "0.aws:runShellScript"),
                                jsonProperty("runCommand", jsonArray("halt"))
                            )))
                        ))
                    ))
                )))
            .build());

        final var ssmLifecyclePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("ssm:SendCommand")
                    .resources("arn:aws:ec2:eu-west-1:1234567890:instance/*")
                    .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                        .test("StringEquals")
                        .variable("ec2:ResourceTag/Terminate")
                        .values("*")
                        .build())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("ssm:SendCommand")
                    .resources(stopInstance.arn())
                    .build())
            .build());

        var ssmLifecycleRole = new Role("ssmLifecycleRole", RoleArgs.builder()        
            .assumeRolePolicy(ssmLifecycleTrust.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var ssmLifecyclePolicy = new Policy("ssmLifecyclePolicy", PolicyArgs.builder()        
            .policy(ssmLifecyclePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(ssmLifecyclePolicyDocument -> ssmLifecyclePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var ssmLifecycleRolePolicyAttachment = new RolePolicyAttachment("ssmLifecycleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .policyArn(ssmLifecyclePolicy.arn())
            .role(ssmLifecycleRole.name())
            .build());

        var stopInstancesEventRule = new EventRule("stopInstancesEventRule", EventRuleArgs.builder()        
            .description("Stop instances nightly")
            .scheduleExpression("cron(0 0 * * ? *)")
            .build());

        var stopInstancesEventTarget = new EventTarget("stopInstancesEventTarget", EventTargetArgs.builder()        
            .arn(stopInstance.arn())
            .rule(stopInstancesEventRule.name())
            .roleArn(ssmLifecycleRole.arn())
            .runCommandTargets(EventTargetRunCommandTargetArgs.builder()
                .key("tag:Terminate")
                .values("midnight")
                .build())
            .build());

    }
}
```
```yaml
resources:
  ssmLifecycleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${ssmLifecycleTrust.json}
  ssmLifecyclePolicy:
    type: aws:iam:Policy
    properties:
      policy: ${ssmLifecyclePolicyDocument.json}
  ssmLifecycleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: ${ssmLifecyclePolicy.arn}
      role: ${ssmLifecycleRole.name}
  stopInstance:
    type: aws:ssm:Document
    properties:
      documentType: Command
      content:
        fn::toJSON:
          schemaVersion: '1.2'
          description: Stop an instance
          parameters: {}
          runtimeConfig:
            aws:runShellScript:
              properties:
                - id: 0.aws:runShellScript
                  runCommand:
                    - halt
  stopInstancesEventRule:
    type: aws:cloudwatch:EventRule
    properties:
      description: Stop instances nightly
      scheduleExpression: cron(0 0 * * ? *)
  stopInstancesEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      arn: ${stopInstance.arn}
      rule: ${stopInstancesEventRule.name}
      roleArn: ${ssmLifecycleRole.arn}
      runCommandTargets:
        - key: tag:Terminate
          values:
            - midnight
variables:
  ssmLifecycleTrust:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - events.amazonaws.com
  ssmLifecyclePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - ssm:SendCommand
            resources:
              - arn:aws:ec2:eu-west-1:1234567890:instance/*
            conditions:
              - test: StringEquals
                variable: ec2:ResourceTag/Terminate
                values:
                  - '*'
          - effect: Allow
            actions:
              - ssm:SendCommand
            resources:
              - ${stopInstance.arn}
```
{{% /example %}}
{{% example %}}
### RunCommand Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
    description: "Stop instances nightly",
    scheduleExpression: "cron(0 0 * * ? *)",
});
const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
    arn: `arn:aws:ssm:${_var.aws_region}::document/AWS-RunShellScript`,
    input: "{\"commands\":[\"halt\"]}",
    rule: stopInstancesEventRule.name,
    roleArn: aws_iam_role.ssm_lifecycle.arn,
    runCommandTargets: [{
        key: "tag:Terminate",
        values: ["midnight"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

stop_instances_event_rule = aws.cloudwatch.EventRule("stopInstancesEventRule",
    description="Stop instances nightly",
    schedule_expression="cron(0 0 * * ? *)")
stop_instances_event_target = aws.cloudwatch.EventTarget("stopInstancesEventTarget",
    arn=f"arn:aws:ssm:{var['aws_region']}::document/AWS-RunShellScript",
    input="{\"commands\":[\"halt\"]}",
    rule=stop_instances_event_rule.name,
    role_arn=aws_iam_role["ssm_lifecycle"]["arn"],
    run_command_targets=[aws.cloudwatch.EventTargetRunCommandTargetArgs(
        key="tag:Terminate",
        values=["midnight"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var stopInstancesEventRule = new Aws.CloudWatch.EventRule("stopInstancesEventRule", new()
    {
        Description = "Stop instances nightly",
        ScheduleExpression = "cron(0 0 * * ? *)",
    });

    var stopInstancesEventTarget = new Aws.CloudWatch.EventTarget("stopInstancesEventTarget", new()
    {
        Arn = $"arn:aws:ssm:{@var.Aws_region}::document/AWS-RunShellScript",
        Input = "{\"commands\":[\"halt\"]}",
        Rule = stopInstancesEventRule.Name,
        RoleArn = aws_iam_role.Ssm_lifecycle.Arn,
        RunCommandTargets = new[]
        {
            new Aws.CloudWatch.Inputs.EventTargetRunCommandTargetArgs
            {
                Key = "tag:Terminate",
                Values = new[]
                {
                    "midnight",
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stopInstancesEventRule, err := cloudwatch.NewEventRule(ctx, "stopInstancesEventRule", &cloudwatch.EventRuleArgs{
			Description:        pulumi.String("Stop instances nightly"),
			ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "stopInstancesEventTarget", &cloudwatch.EventTargetArgs{
			Arn:     pulumi.String(fmt.Sprintf("arn:aws:ssm:%v::document/AWS-RunShellScript", _var.Aws_region)),
			Input:   pulumi.String("{\"commands\":[\"halt\"]}"),
			Rule:    stopInstancesEventRule.Name,
			RoleArn: pulumi.Any(aws_iam_role.Ssm_lifecycle.Arn),
			RunCommandTargets: cloudwatch.EventTargetRunCommandTargetArray{
				&cloudwatch.EventTargetRunCommandTargetArgs{
					Key: pulumi.String("tag:Terminate"),
					Values: pulumi.StringArray{
						pulumi.String("midnight"),
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
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventRuleArgs;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.cloudwatch.inputs.EventTargetRunCommandTargetArgs;
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
        var stopInstancesEventRule = new EventRule("stopInstancesEventRule", EventRuleArgs.builder()        
            .description("Stop instances nightly")
            .scheduleExpression("cron(0 0 * * ? *)")
            .build());

        var stopInstancesEventTarget = new EventTarget("stopInstancesEventTarget", EventTargetArgs.builder()        
            .arn(String.format("arn:aws:ssm:%s::document/AWS-RunShellScript", var_.aws_region()))
            .input("{\"commands\":[\"halt\"]}")
            .rule(stopInstancesEventRule.name())
            .roleArn(aws_iam_role.ssm_lifecycle().arn())
            .runCommandTargets(EventTargetRunCommandTargetArgs.builder()
                .key("tag:Terminate")
                .values("midnight")
                .build())
            .build());

    }
}
```
```yaml
resources:
  stopInstancesEventRule:
    type: aws:cloudwatch:EventRule
    properties:
      description: Stop instances nightly
      scheduleExpression: cron(0 0 * * ? *)
  stopInstancesEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      arn: arn:aws:ssm:${var.aws_region}::document/AWS-RunShellScript
      input: '{"commands":["halt"]}'
      rule: ${stopInstancesEventRule.name}
      roleArn: ${aws_iam_role.ssm_lifecycle.arn}
      runCommandTargets:
        - key: tag:Terminate
          values:
            - midnight
```
{{% /example %}}
{{% example %}}
### API Gateway target

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
// ...
const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {restApi: aws_api_gateway_rest_api.example.id});
// ...
const exampleStage = new aws.apigateway.Stage("exampleStage", {
    restApi: aws_api_gateway_rest_api.example.id,
    deployment: exampleDeployment.id,
});
// ...
const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
    arn: pulumi.interpolate`${exampleStage.executionArn}/GET`,
    rule: exampleEventRule.id,
    httpTarget: {
        queryStringParameters: {
            Body: "$.detail.body",
        },
        headerParameters: {
            Env: "Test",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_event_rule = aws.cloudwatch.EventRule("exampleEventRule")
# ...
example_deployment = aws.apigateway.Deployment("exampleDeployment", rest_api=aws_api_gateway_rest_api["example"]["id"])
# ...
example_stage = aws.apigateway.Stage("exampleStage",
    rest_api=aws_api_gateway_rest_api["example"]["id"],
    deployment=example_deployment.id)
# ...
example_event_target = aws.cloudwatch.EventTarget("exampleEventTarget",
    arn=example_stage.execution_arn.apply(lambda execution_arn: f"{execution_arn}/GET"),
    rule=example_event_rule.id,
    http_target=aws.cloudwatch.EventTargetHttpTargetArgs(
        query_string_parameters={
            "Body": "$.detail.body",
        },
        header_parameters={
            "Env": "Test",
        },
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule");

    // ...
    var exampleDeployment = new Aws.ApiGateway.Deployment("exampleDeployment", new()
    {
        RestApi = aws_api_gateway_rest_api.Example.Id,
    });

    // ...
    var exampleStage = new Aws.ApiGateway.Stage("exampleStage", new()
    {
        RestApi = aws_api_gateway_rest_api.Example.Id,
        Deployment = exampleDeployment.Id,
    });

    // ...
    var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new()
    {
        Arn = exampleStage.ExecutionArn.Apply(executionArn => $"{executionArn}/GET"),
        Rule = exampleEventRule.Id,
        HttpTarget = new Aws.CloudWatch.Inputs.EventTargetHttpTargetArgs
        {
            QueryStringParameters = 
            {
                { "Body", "$.detail.body" },
            },
            HeaderParameters = 
            {
                { "Env", "Test" },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleEventRule, err := cloudwatch.NewEventRule(ctx, "exampleEventRule", nil)
		if err != nil {
			return err
		}
		exampleDeployment, err := apigateway.NewDeployment(ctx, "exampleDeployment", &apigateway.DeploymentArgs{
			RestApi: pulumi.Any(aws_api_gateway_rest_api.Example.Id),
		})
		if err != nil {
			return err
		}
		exampleStage, err := apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
			RestApi:    pulumi.Any(aws_api_gateway_rest_api.Example.Id),
			Deployment: exampleDeployment.ID(),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "exampleEventTarget", &cloudwatch.EventTargetArgs{
			Arn: exampleStage.ExecutionArn.ApplyT(func(executionArn string) (string, error) {
				return fmt.Sprintf("%v/GET", executionArn), nil
			}).(pulumi.StringOutput),
			Rule: exampleEventRule.ID(),
			HttpTarget: &cloudwatch.EventTargetHttpTargetArgs{
				QueryStringParameters: pulumi.StringMap{
					"Body": pulumi.String("$.detail.body"),
				},
				HeaderParameters: pulumi.StringMap{
					"Env": pulumi.String("Test"),
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
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.apigateway.Deployment;
import com.pulumi.aws.apigateway.DeploymentArgs;
import com.pulumi.aws.apigateway.Stage;
import com.pulumi.aws.apigateway.StageArgs;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.cloudwatch.inputs.EventTargetHttpTargetArgs;
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
        var exampleEventRule = new EventRule("exampleEventRule");

        var exampleDeployment = new Deployment("exampleDeployment", DeploymentArgs.builder()        
            .restApi(aws_api_gateway_rest_api.example().id())
            .build());

        var exampleStage = new Stage("exampleStage", StageArgs.builder()        
            .restApi(aws_api_gateway_rest_api.example().id())
            .deployment(exampleDeployment.id())
            .build());

        var exampleEventTarget = new EventTarget("exampleEventTarget", EventTargetArgs.builder()        
            .arn(exampleStage.executionArn().applyValue(executionArn -> String.format("%s/GET", executionArn)))
            .rule(exampleEventRule.id())
            .httpTarget(EventTargetHttpTargetArgs.builder()
                .queryStringParameters(Map.of("Body", "$.detail.body"))
                .headerParameters(Map.of("Env", "Test"))
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      arn: ${exampleStage.executionArn}/GET
      rule: ${exampleEventRule.id}
      httpTarget:
        queryStringParameters:
          Body: $.detail.body
        headerParameters:
          Env: Test
  exampleEventRule:
    type: aws:cloudwatch:EventRule
  exampleDeployment:
    type: aws:apigateway:Deployment
    properties:
      restApi: ${aws_api_gateway_rest_api.example.id}
  exampleStage:
    type: aws:apigateway:Stage
    properties:
      restApi: ${aws_api_gateway_rest_api.example.id}
      deployment: ${exampleDeployment.id}
```
{{% /example %}}
{{% example %}}
### Cross-Account Event Bus target

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["events.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const eventBusInvokeRemoteEventBusRole = new aws.iam.Role("eventBusInvokeRemoteEventBusRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const eventBusInvokeRemoteEventBusPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["events:PutEvents"],
        resources: ["arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus"],
    }],
});
const eventBusInvokeRemoteEventBusPolicy = new aws.iam.Policy("eventBusInvokeRemoteEventBusPolicy", {policy: eventBusInvokeRemoteEventBusPolicyDocument.then(eventBusInvokeRemoteEventBusPolicyDocument => eventBusInvokeRemoteEventBusPolicyDocument.json)});
const eventBusInvokeRemoteEventBusRolePolicyAttachment = new aws.iam.RolePolicyAttachment("eventBusInvokeRemoteEventBusRolePolicyAttachment", {
    role: eventBusInvokeRemoteEventBusRole.name,
    policyArn: eventBusInvokeRemoteEventBusPolicy.arn,
});
const stopInstancesEventRule = new aws.cloudwatch.EventRule("stopInstancesEventRule", {
    description: "Stop instances nightly",
    scheduleExpression: "cron(0 0 * * ? *)",
});
const stopInstancesEventTarget = new aws.cloudwatch.EventTarget("stopInstancesEventTarget", {
    arn: "arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus",
    rule: stopInstancesEventRule.name,
    roleArn: eventBusInvokeRemoteEventBusRole.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["events.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
event_bus_invoke_remote_event_bus_role = aws.iam.Role("eventBusInvokeRemoteEventBusRole", assume_role_policy=assume_role.json)
event_bus_invoke_remote_event_bus_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["events:PutEvents"],
    resources=["arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus"],
)])
event_bus_invoke_remote_event_bus_policy = aws.iam.Policy("eventBusInvokeRemoteEventBusPolicy", policy=event_bus_invoke_remote_event_bus_policy_document.json)
event_bus_invoke_remote_event_bus_role_policy_attachment = aws.iam.RolePolicyAttachment("eventBusInvokeRemoteEventBusRolePolicyAttachment",
    role=event_bus_invoke_remote_event_bus_role.name,
    policy_arn=event_bus_invoke_remote_event_bus_policy.arn)
stop_instances_event_rule = aws.cloudwatch.EventRule("stopInstancesEventRule",
    description="Stop instances nightly",
    schedule_expression="cron(0 0 * * ? *)")
stop_instances_event_target = aws.cloudwatch.EventTarget("stopInstancesEventTarget",
    arn="arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus",
    rule=stop_instances_event_rule.name,
    role_arn=event_bus_invoke_remote_event_bus_role.arn)
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
                            "events.amazonaws.com",
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

    var eventBusInvokeRemoteEventBusRole = new Aws.Iam.Role("eventBusInvokeRemoteEventBusRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var eventBusInvokeRemoteEventBusPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "events:PutEvents",
                },
                Resources = new[]
                {
                    "arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus",
                },
            },
        },
    });

    var eventBusInvokeRemoteEventBusPolicy = new Aws.Iam.Policy("eventBusInvokeRemoteEventBusPolicy", new()
    {
        PolicyDocument = eventBusInvokeRemoteEventBusPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var eventBusInvokeRemoteEventBusRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("eventBusInvokeRemoteEventBusRolePolicyAttachment", new()
    {
        Role = eventBusInvokeRemoteEventBusRole.Name,
        PolicyArn = eventBusInvokeRemoteEventBusPolicy.Arn,
    });

    var stopInstancesEventRule = new Aws.CloudWatch.EventRule("stopInstancesEventRule", new()
    {
        Description = "Stop instances nightly",
        ScheduleExpression = "cron(0 0 * * ? *)",
    });

    var stopInstancesEventTarget = new Aws.CloudWatch.EventTarget("stopInstancesEventTarget", new()
    {
        Arn = "arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus",
        Rule = stopInstancesEventRule.Name,
        RoleArn = eventBusInvokeRemoteEventBusRole.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
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
								"events.amazonaws.com",
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
		eventBusInvokeRemoteEventBusRole, err := iam.NewRole(ctx, "eventBusInvokeRemoteEventBusRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		eventBusInvokeRemoteEventBusPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"events:PutEvents",
					},
					Resources: []string{
						"arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		eventBusInvokeRemoteEventBusPolicy, err := iam.NewPolicy(ctx, "eventBusInvokeRemoteEventBusPolicy", &iam.PolicyArgs{
			Policy: *pulumi.String(eventBusInvokeRemoteEventBusPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "eventBusInvokeRemoteEventBusRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      eventBusInvokeRemoteEventBusRole.Name,
			PolicyArn: eventBusInvokeRemoteEventBusPolicy.Arn,
		})
		if err != nil {
			return err
		}
		stopInstancesEventRule, err := cloudwatch.NewEventRule(ctx, "stopInstancesEventRule", &cloudwatch.EventRuleArgs{
			Description:        pulumi.String("Stop instances nightly"),
			ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "stopInstancesEventTarget", &cloudwatch.EventTargetArgs{
			Arn:     pulumi.String("arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus"),
			Rule:    stopInstancesEventRule.Name,
			RoleArn: eventBusInvokeRemoteEventBusRole.Arn,
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
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventRuleArgs;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
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
                    .identifiers("events.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var eventBusInvokeRemoteEventBusRole = new Role("eventBusInvokeRemoteEventBusRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var eventBusInvokeRemoteEventBusPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("events:PutEvents")
                .resources("arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus")
                .build())
            .build());

        var eventBusInvokeRemoteEventBusPolicy = new Policy("eventBusInvokeRemoteEventBusPolicy", PolicyArgs.builder()        
            .policy(eventBusInvokeRemoteEventBusPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var eventBusInvokeRemoteEventBusRolePolicyAttachment = new RolePolicyAttachment("eventBusInvokeRemoteEventBusRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(eventBusInvokeRemoteEventBusRole.name())
            .policyArn(eventBusInvokeRemoteEventBusPolicy.arn())
            .build());

        var stopInstancesEventRule = new EventRule("stopInstancesEventRule", EventRuleArgs.builder()        
            .description("Stop instances nightly")
            .scheduleExpression("cron(0 0 * * ? *)")
            .build());

        var stopInstancesEventTarget = new EventTarget("stopInstancesEventTarget", EventTargetArgs.builder()        
            .arn("arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus")
            .rule(stopInstancesEventRule.name())
            .roleArn(eventBusInvokeRemoteEventBusRole.arn())
            .build());

    }
}
```
```yaml
resources:
  eventBusInvokeRemoteEventBusRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  eventBusInvokeRemoteEventBusPolicy:
    type: aws:iam:Policy
    properties:
      policy: ${eventBusInvokeRemoteEventBusPolicyDocument.json}
  eventBusInvokeRemoteEventBusRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${eventBusInvokeRemoteEventBusRole.name}
      policyArn: ${eventBusInvokeRemoteEventBusPolicy.arn}
  stopInstancesEventRule:
    type: aws:cloudwatch:EventRule
    properties:
      description: Stop instances nightly
      scheduleExpression: cron(0 0 * * ? *)
  stopInstancesEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      arn: arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus
      rule: ${stopInstancesEventRule.name}
      roleArn: ${eventBusInvokeRemoteEventBusRole.arn}
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
                  - events.amazonaws.com
            actions:
              - sts:AssumeRole
  eventBusInvokeRemoteEventBusPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - events:PutEvents
            resources:
              - arn:aws:events:eu-west-1:1234567890:event-bus/My-Event-Bus
```
{{% /example %}}
{{% example %}}
### Input Transformer Usage - JSON Object

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
// ...
const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
    arn: aws_lambda_function.example.arn,
    rule: exampleEventRule.id,
    inputTransformer: {
        inputPaths: {
            instance: "$.detail.instance",
            status: "$.detail.status",
        },
        inputTemplate: `{
  "instance_id": <instance>,
  "instance_status": <status>
}
`,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_event_rule = aws.cloudwatch.EventRule("exampleEventRule")
# ...
example_event_target = aws.cloudwatch.EventTarget("exampleEventTarget",
    arn=aws_lambda_function["example"]["arn"],
    rule=example_event_rule.id,
    input_transformer=aws.cloudwatch.EventTargetInputTransformerArgs(
        input_paths={
            "instance": "$.detail.instance",
            "status": "$.detail.status",
        },
        input_template="""{
  "instance_id": <instance>,
  "instance_status": <status>
}
""",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule");

    // ...
    var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new()
    {
        Arn = aws_lambda_function.Example.Arn,
        Rule = exampleEventRule.Id,
        InputTransformer = new Aws.CloudWatch.Inputs.EventTargetInputTransformerArgs
        {
            InputPaths = 
            {
                { "instance", "$.detail.instance" },
                { "status", "$.detail.status" },
            },
            InputTemplate = @"{
  ""instance_id"": <instance>,
  ""instance_status"": <status>
}
",
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
		exampleEventRule, err := cloudwatch.NewEventRule(ctx, "exampleEventRule", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "exampleEventTarget", &cloudwatch.EventTargetArgs{
			Arn:  pulumi.Any(aws_lambda_function.Example.Arn),
			Rule: exampleEventRule.ID(),
			InputTransformer: &cloudwatch.EventTargetInputTransformerArgs{
				InputPaths: pulumi.StringMap{
					"instance": pulumi.String("$.detail.instance"),
					"status":   pulumi.String("$.detail.status"),
				},
				InputTemplate: pulumi.String("{\n  \"instance_id\": <instance>,\n  \"instance_status\": <status>\n}\n"),
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
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.cloudwatch.inputs.EventTargetInputTransformerArgs;
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
        var exampleEventRule = new EventRule("exampleEventRule");

        var exampleEventTarget = new EventTarget("exampleEventTarget", EventTargetArgs.builder()        
            .arn(aws_lambda_function.example().arn())
            .rule(exampleEventRule.id())
            .inputTransformer(EventTargetInputTransformerArgs.builder()
                .inputPaths(Map.ofEntries(
                    Map.entry("instance", "$.detail.instance"),
                    Map.entry("status", "$.detail.status")
                ))
                .inputTemplate("""
{
  "instance_id": <instance>,
  "instance_status": <status>
}
                """)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      arn: ${aws_lambda_function.example.arn}
      rule: ${exampleEventRule.id}
      inputTransformer:
        inputPaths:
          instance: $.detail.instance
          status: $.detail.status
        inputTemplate: |
          {
            "instance_id": <instance>,
            "instance_status": <status>
          }
  exampleEventRule:
    type: aws:cloudwatch:EventRule
```
{{% /example %}}
{{% example %}}
### Input Transformer Usage - Simple String

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {});
// ...
const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
    arn: aws_lambda_function.example.arn,
    rule: exampleEventRule.id,
    inputTransformer: {
        inputPaths: {
            instance: "$.detail.instance",
            status: "$.detail.status",
        },
        inputTemplate: "\"<instance> is in state <status>\"",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_event_rule = aws.cloudwatch.EventRule("exampleEventRule")
# ...
example_event_target = aws.cloudwatch.EventTarget("exampleEventTarget",
    arn=aws_lambda_function["example"]["arn"],
    rule=example_event_rule.id,
    input_transformer=aws.cloudwatch.EventTargetInputTransformerArgs(
        input_paths={
            "instance": "$.detail.instance",
            "status": "$.detail.status",
        },
        input_template="\"<instance> is in state <status>\"",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule");

    // ...
    var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new()
    {
        Arn = aws_lambda_function.Example.Arn,
        Rule = exampleEventRule.Id,
        InputTransformer = new Aws.CloudWatch.Inputs.EventTargetInputTransformerArgs
        {
            InputPaths = 
            {
                { "instance", "$.detail.instance" },
                { "status", "$.detail.status" },
            },
            InputTemplate = "\"<instance> is in state <status>\"",
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
		exampleEventRule, err := cloudwatch.NewEventRule(ctx, "exampleEventRule", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "exampleEventTarget", &cloudwatch.EventTargetArgs{
			Arn:  pulumi.Any(aws_lambda_function.Example.Arn),
			Rule: exampleEventRule.ID(),
			InputTransformer: &cloudwatch.EventTargetInputTransformerArgs{
				InputPaths: pulumi.StringMap{
					"instance": pulumi.String("$.detail.instance"),
					"status":   pulumi.String("$.detail.status"),
				},
				InputTemplate: pulumi.String("\"<instance> is in state <status>\""),
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
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import com.pulumi.aws.cloudwatch.inputs.EventTargetInputTransformerArgs;
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
        var exampleEventRule = new EventRule("exampleEventRule");

        var exampleEventTarget = new EventTarget("exampleEventTarget", EventTargetArgs.builder()        
            .arn(aws_lambda_function.example().arn())
            .rule(exampleEventRule.id())
            .inputTransformer(EventTargetInputTransformerArgs.builder()
                .inputPaths(Map.ofEntries(
                    Map.entry("instance", "$.detail.instance"),
                    Map.entry("status", "$.detail.status")
                ))
                .inputTemplate("\"<instance> is in state <status>\"")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      arn: ${aws_lambda_function.example.arn}
      rule: ${exampleEventRule.id}
      inputTransformer:
        inputPaths:
          instance: $.detail.instance
          status: $.detail.status
        inputTemplate: '"<instance> is in state <status>"'
  exampleEventRule:
    type: aws:cloudwatch:EventRule
```
{{% /example %}}
{{% example %}}
### Cloudwatch Log Group Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 1});
const exampleEventRule = new aws.cloudwatch.EventRule("exampleEventRule", {
    description: "GuardDuty Findings",
    eventPattern: JSON.stringify({
        source: ["aws.guardduty"],
    }),
    tags: {
        Environment: "example",
    },
});
const exampleLogPolicy = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: ["logs:CreateLogStream"],
            resources: [pulumi.interpolate`${exampleLogGroup.arn}:*`],
            principals: [{
                type: "Service",
                identifiers: [
                    "events.amazonaws.com",
                    "delivery.logs.amazonaws.com",
                ],
            }],
        },
        {
            effect: "Allow",
            actions: ["logs:PutLogEvents"],
            resources: [pulumi.interpolate`${exampleLogGroup.arn}:*:*`],
            principals: [{
                type: "Service",
                identifiers: [
                    "events.amazonaws.com",
                    "delivery.logs.amazonaws.com",
                ],
            }],
            conditions: [{
                test: "ArnEquals",
                values: [exampleEventRule.arn],
                variable: "aws:SourceArn",
            }],
        },
    ],
});
const exampleLogResourcePolicy = new aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy", {
    policyDocument: exampleLogPolicy.apply(exampleLogPolicy => exampleLogPolicy.json),
    policyName: "guardduty-log-publishing-policy",
});
const exampleEventTarget = new aws.cloudwatch.EventTarget("exampleEventTarget", {
    rule: exampleEventRule.name,
    arn: exampleLogGroup.arn,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=1)
example_event_rule = aws.cloudwatch.EventRule("exampleEventRule",
    description="GuardDuty Findings",
    event_pattern=json.dumps({
        "source": ["aws.guardduty"],
    }),
    tags={
        "Environment": "example",
    })
example_log_policy = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["logs:CreateLogStream"],
        resources=[example_log_group.arn.apply(lambda arn: f"{arn}:*")],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=[
                "events.amazonaws.com",
                "delivery.logs.amazonaws.com",
            ],
        )],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["logs:PutLogEvents"],
        resources=[example_log_group.arn.apply(lambda arn: f"{arn}:*:*")],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=[
                "events.amazonaws.com",
                "delivery.logs.amazonaws.com",
            ],
        )],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="ArnEquals",
            values=[example_event_rule.arn],
            variable="aws:SourceArn",
        )],
    ),
])
example_log_resource_policy = aws.cloudwatch.LogResourcePolicy("exampleLogResourcePolicy",
    policy_document=example_log_policy.json,
    policy_name="guardduty-log-publishing-policy")
example_event_target = aws.cloudwatch.EventTarget("exampleEventTarget",
    rule=example_event_rule.name,
    arn=example_log_group.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new()
    {
        RetentionInDays = 1,
    });

    var exampleEventRule = new Aws.CloudWatch.EventRule("exampleEventRule", new()
    {
        Description = "GuardDuty Findings",
        EventPattern = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["source"] = new[]
            {
                "aws.guardduty",
            },
        }),
        Tags = 
        {
            { "Environment", "example" },
        },
    });

    var exampleLogPolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "logs:CreateLogStream",
                },
                Resources = new[]
                {
                    $"{exampleLogGroup.Arn}:*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "events.amazonaws.com",
                            "delivery.logs.amazonaws.com",
                        },
                    },
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "logs:PutLogEvents",
                },
                Resources = new[]
                {
                    $"{exampleLogGroup.Arn}:*:*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "events.amazonaws.com",
                            "delivery.logs.amazonaws.com",
                        },
                    },
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "ArnEquals",
                        Values = new[]
                        {
                            exampleEventRule.Arn,
                        },
                        Variable = "aws:SourceArn",
                    },
                },
            },
        },
    });

    var exampleLogResourcePolicy = new Aws.CloudWatch.LogResourcePolicy("exampleLogResourcePolicy", new()
    {
        PolicyDocument = exampleLogPolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        PolicyName = "guardduty-log-publishing-policy",
    });

    var exampleEventTarget = new Aws.CloudWatch.EventTarget("exampleEventTarget", new()
    {
        Rule = exampleEventRule.Name,
        Arn = exampleLogGroup.Arn,
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"source": []string{
				"aws.guardduty",
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleEventRule, err := cloudwatch.NewEventRule(ctx, "exampleEventRule", &cloudwatch.EventRuleArgs{
			Description:  pulumi.String("GuardDuty Findings"),
			EventPattern: pulumi.String(json0),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("example"),
			},
		})
		if err != nil {
			return err
		}
		exampleLogPolicy := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("logs:CreateLogStream"),
					},
					Resources: pulumi.StringArray{
						exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v:*", arn), nil
						}).(pulumi.StringOutput),
					},
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("events.amazonaws.com"),
								pulumi.String("delivery.logs.amazonaws.com"),
							},
						},
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("logs:PutLogEvents"),
					},
					Resources: pulumi.StringArray{
						exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v:*:*", arn), nil
						}).(pulumi.StringOutput),
					},
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("events.amazonaws.com"),
								pulumi.String("delivery.logs.amazonaws.com"),
							},
						},
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test: pulumi.String("ArnEquals"),
							Values: pulumi.StringArray{
								exampleEventRule.Arn,
							},
							Variable: pulumi.String("aws:SourceArn"),
						},
					},
				},
			},
		}, nil)
		_, err = cloudwatch.NewLogResourcePolicy(ctx, "exampleLogResourcePolicy", &cloudwatch.LogResourcePolicyArgs{
			PolicyDocument: exampleLogPolicy.ApplyT(func(exampleLogPolicy iam.GetPolicyDocumentResult) (*string, error) {
				return &exampleLogPolicy.Json, nil
			}).(pulumi.StringPtrOutput),
			PolicyName: pulumi.String("guardduty-log-publishing-policy"),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewEventTarget(ctx, "exampleEventTarget", &cloudwatch.EventTargetArgs{
			Rule: exampleEventRule.Name,
			Arn:  exampleLogGroup.Arn,
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
import com.pulumi.aws.cloudwatch.LogGroupArgs;
import com.pulumi.aws.cloudwatch.EventRule;
import com.pulumi.aws.cloudwatch.EventRuleArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudwatch.LogResourcePolicy;
import com.pulumi.aws.cloudwatch.LogResourcePolicyArgs;
import com.pulumi.aws.cloudwatch.EventTarget;
import com.pulumi.aws.cloudwatch.EventTargetArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup", LogGroupArgs.builder()        
            .retentionInDays(1)
            .build());

        var exampleEventRule = new EventRule("exampleEventRule", EventRuleArgs.builder()        
            .description("GuardDuty Findings")
            .eventPattern(serializeJson(
                jsonObject(
                    jsonProperty("source", jsonArray("aws.guardduty"))
                )))
            .tags(Map.of("Environment", "example"))
            .build());

        final var exampleLogPolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("logs:CreateLogStream")
                    .resources(exampleLogGroup.arn().applyValue(arn -> String.format("%s:*", arn)))
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers(                        
                            "events.amazonaws.com",
                            "delivery.logs.amazonaws.com")
                        .build())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("logs:PutLogEvents")
                    .resources(exampleLogGroup.arn().applyValue(arn -> String.format("%s:*:*", arn)))
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers(                        
                            "events.amazonaws.com",
                            "delivery.logs.amazonaws.com")
                        .build())
                    .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                        .test("ArnEquals")
                        .values(exampleEventRule.arn())
                        .variable("aws:SourceArn")
                        .build())
                    .build())
            .build());

        var exampleLogResourcePolicy = new LogResourcePolicy("exampleLogResourcePolicy", LogResourcePolicyArgs.builder()        
            .policyDocument(exampleLogPolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(exampleLogPolicy -> exampleLogPolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .policyName("guardduty-log-publishing-policy")
            .build());

        var exampleEventTarget = new EventTarget("exampleEventTarget", EventTargetArgs.builder()        
            .rule(exampleEventRule.name())
            .arn(exampleLogGroup.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
    properties:
      retentionInDays: 1
  exampleLogResourcePolicy:
    type: aws:cloudwatch:LogResourcePolicy
    properties:
      policyDocument: ${exampleLogPolicy.json}
      policyName: guardduty-log-publishing-policy
  exampleEventRule:
    type: aws:cloudwatch:EventRule
    properties:
      description: GuardDuty Findings
      eventPattern:
        fn::toJSON:
          source:
            - aws.guardduty
      tags:
        Environment: example
  exampleEventTarget:
    type: aws:cloudwatch:EventTarget
    properties:
      rule: ${exampleEventRule.name}
      arn: ${exampleLogGroup.arn}
variables:
  exampleLogPolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - logs:CreateLogStream
            resources:
              - ${exampleLogGroup.arn}:*
            principals:
              - type: Service
                identifiers:
                  - events.amazonaws.com
                  - delivery.logs.amazonaws.com
          - effect: Allow
            actions:
              - logs:PutLogEvents
            resources:
              - ${exampleLogGroup.arn}:*:*
            principals:
              - type: Service
                identifiers:
                  - events.amazonaws.com
                  - delivery.logs.amazonaws.com
            conditions:
              - test: ArnEquals
                values:
                  - ${exampleEventRule.arn}
                variable: aws:SourceArn
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge Targets using `event_bus_name/rule-name/target-id` (if you omit `event_bus_name`, the `default` event bus will be used). For example:

```sh
 $ pulumi import aws:cloudwatch/eventTarget:EventTarget test-event-target rule-name/target-id
```
 