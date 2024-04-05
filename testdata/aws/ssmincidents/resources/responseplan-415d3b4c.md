Provides a resource to manage response plans in AWS Systems Manager Incident Manager.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssmincidents.ResponsePlan("example", {
    incidentTemplate: {
        title: "title",
        impact: 3,
    },
    tags: {
        key: "value",
    },
}, {
    dependsOn: [aws_ssmincidents_replication_set.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssmincidents.ResponsePlan("example",
    incident_template=aws.ssmincidents.ResponsePlanIncidentTemplateArgs(
        title="title",
        impact=3,
    ),
    tags={
        "key": "value",
    },
    opts=pulumi.ResourceOptions(depends_on=[aws_ssmincidents_replication_set["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsmIncidents.ResponsePlan("example", new()
    {
        IncidentTemplate = new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateArgs
        {
            Title = "title",
            Impact = 3,
        },
        Tags = 
        {
            { "key", "value" },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_ssmincidents_replication_set.Example,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmincidents.NewResponsePlan(ctx, "example", &ssmincidents.ResponsePlanArgs{
			IncidentTemplate: &ssmincidents.ResponsePlanIncidentTemplateArgs{
				Title:  pulumi.String("title"),
				Impact: pulumi.Int(3),
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_ssmincidents_replication_set.Example,
		}))
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
import com.pulumi.aws.ssmincidents.ResponsePlan;
import com.pulumi.aws.ssmincidents.ResponsePlanArgs;
import com.pulumi.aws.ssmincidents.inputs.ResponsePlanIncidentTemplateArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var example = new ResponsePlan("example", ResponsePlanArgs.builder()        
            .incidentTemplate(ResponsePlanIncidentTemplateArgs.builder()
                .title("title")
                .impact("3")
                .build())
            .tags(Map.of("key", "value"))
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_ssmincidents_replication_set.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssmincidents:ResponsePlan
    properties:
      incidentTemplate:
        title: title
        impact: '3'
      tags:
        key: value
    options:
      dependson:
        - ${aws_ssmincidents_replication_set.example}
```
{{% /example %}}
{{% example %}}
### Usage With All Fields

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssmincidents.ResponsePlan("example", {
    incidentTemplate: {
        title: "title",
        impact: 3,
        dedupeString: "dedupe",
        incidentTags: {
            key: "value",
        },
        notificationTargets: [
            {
                snsTopicArn: aws_sns_topic.example1.arn,
            },
            {
                snsTopicArn: aws_sns_topic.example2.arn,
            },
        ],
        summary: "summary",
    },
    displayName: "display name",
    chatChannels: [aws_sns_topic.topic.arn],
    engagements: ["arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1"],
    action: {
        ssmAutomations: [{
            documentName: aws_ssm_document.document1.name,
            roleArn: aws_iam_role.role1.arn,
            documentVersion: "version1",
            targetAccount: "RESPONSE_PLAN_OWNER_ACCOUNT",
            parameters: [
                {
                    name: "key",
                    values: [
                        "value1",
                        "value2",
                    ],
                },
                {
                    name: "foo",
                    values: ["bar"],
                },
            ],
            dynamicParameters: {
                someKey: "INVOLVED_RESOURCES",
                anotherKey: "INCIDENT_RECORD_ARN",
            },
        }],
    },
    integration: {
        pagerduties: [{
            name: "pagerdutyIntergration",
            serviceId: "example",
            secretId: "example",
        }],
    },
    tags: {
        key: "value",
    },
}, {
    dependsOn: [aws_ssmincidents_replication_set.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssmincidents.ResponsePlan("example",
    incident_template=aws.ssmincidents.ResponsePlanIncidentTemplateArgs(
        title="title",
        impact=3,
        dedupe_string="dedupe",
        incident_tags={
            "key": "value",
        },
        notification_targets=[
            aws.ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs(
                sns_topic_arn=aws_sns_topic["example1"]["arn"],
            ),
            aws.ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs(
                sns_topic_arn=aws_sns_topic["example2"]["arn"],
            ),
        ],
        summary="summary",
    ),
    display_name="display name",
    chat_channels=[aws_sns_topic["topic"]["arn"]],
    engagements=["arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1"],
    action=aws.ssmincidents.ResponsePlanActionArgs(
        ssm_automations=[aws.ssmincidents.ResponsePlanActionSsmAutomationArgs(
            document_name=aws_ssm_document["document1"]["name"],
            role_arn=aws_iam_role["role1"]["arn"],
            document_version="version1",
            target_account="RESPONSE_PLAN_OWNER_ACCOUNT",
            parameters=[
                aws.ssmincidents.ResponsePlanActionSsmAutomationParameterArgs(
                    name="key",
                    values=[
                        "value1",
                        "value2",
                    ],
                ),
                aws.ssmincidents.ResponsePlanActionSsmAutomationParameterArgs(
                    name="foo",
                    values=["bar"],
                ),
            ],
            dynamic_parameters={
                "someKey": "INVOLVED_RESOURCES",
                "anotherKey": "INCIDENT_RECORD_ARN",
            },
        )],
    ),
    integration=aws.ssmincidents.ResponsePlanIntegrationArgs(
        pagerduties=[aws.ssmincidents.ResponsePlanIntegrationPagerdutyArgs(
            name="pagerdutyIntergration",
            service_id="example",
            secret_id="example",
        )],
    ),
    tags={
        "key": "value",
    },
    opts=pulumi.ResourceOptions(depends_on=[aws_ssmincidents_replication_set["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsmIncidents.ResponsePlan("example", new()
    {
        IncidentTemplate = new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateArgs
        {
            Title = "title",
            Impact = 3,
            DedupeString = "dedupe",
            IncidentTags = 
            {
                { "key", "value" },
            },
            NotificationTargets = new[]
            {
                new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateNotificationTargetArgs
                {
                    SnsTopicArn = aws_sns_topic.Example1.Arn,
                },
                new Aws.SsmIncidents.Inputs.ResponsePlanIncidentTemplateNotificationTargetArgs
                {
                    SnsTopicArn = aws_sns_topic.Example2.Arn,
                },
            },
            Summary = "summary",
        },
        DisplayName = "display name",
        ChatChannels = new[]
        {
            aws_sns_topic.Topic.Arn,
        },
        Engagements = new[]
        {
            "arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1",
        },
        Action = new Aws.SsmIncidents.Inputs.ResponsePlanActionArgs
        {
            SsmAutomations = new[]
            {
                new Aws.SsmIncidents.Inputs.ResponsePlanActionSsmAutomationArgs
                {
                    DocumentName = aws_ssm_document.Document1.Name,
                    RoleArn = aws_iam_role.Role1.Arn,
                    DocumentVersion = "version1",
                    TargetAccount = "RESPONSE_PLAN_OWNER_ACCOUNT",
                    Parameters = new[]
                    {
                        new Aws.SsmIncidents.Inputs.ResponsePlanActionSsmAutomationParameterArgs
                        {
                            Name = "key",
                            Values = new[]
                            {
                                "value1",
                                "value2",
                            },
                        },
                        new Aws.SsmIncidents.Inputs.ResponsePlanActionSsmAutomationParameterArgs
                        {
                            Name = "foo",
                            Values = new[]
                            {
                                "bar",
                            },
                        },
                    },
                    DynamicParameters = 
                    {
                        { "someKey", "INVOLVED_RESOURCES" },
                        { "anotherKey", "INCIDENT_RECORD_ARN" },
                    },
                },
            },
        },
        Integration = new Aws.SsmIncidents.Inputs.ResponsePlanIntegrationArgs
        {
            Pagerduties = new[]
            {
                new Aws.SsmIncidents.Inputs.ResponsePlanIntegrationPagerdutyArgs
                {
                    Name = "pagerdutyIntergration",
                    ServiceId = "example",
                    SecretId = "example",
                },
            },
        },
        Tags = 
        {
            { "key", "value" },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_ssmincidents_replication_set.Example,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmincidents.NewResponsePlan(ctx, "example", &ssmincidents.ResponsePlanArgs{
			IncidentTemplate: &ssmincidents.ResponsePlanIncidentTemplateArgs{
				Title:        pulumi.String("title"),
				Impact:       pulumi.Int(3),
				DedupeString: pulumi.String("dedupe"),
				IncidentTags: pulumi.StringMap{
					"key": pulumi.String("value"),
				},
				NotificationTargets: ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArray{
					&ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs{
						SnsTopicArn: pulumi.Any(aws_sns_topic.Example1.Arn),
					},
					&ssmincidents.ResponsePlanIncidentTemplateNotificationTargetArgs{
						SnsTopicArn: pulumi.Any(aws_sns_topic.Example2.Arn),
					},
				},
				Summary: pulumi.String("summary"),
			},
			DisplayName: pulumi.String("display name"),
			ChatChannels: pulumi.StringArray{
				aws_sns_topic.Topic.Arn,
			},
			Engagements: pulumi.StringArray{
				pulumi.String("arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1"),
			},
			Action: &ssmincidents.ResponsePlanActionArgs{
				SsmAutomations: ssmincidents.ResponsePlanActionSsmAutomationArray{
					&ssmincidents.ResponsePlanActionSsmAutomationArgs{
						DocumentName:    pulumi.Any(aws_ssm_document.Document1.Name),
						RoleArn:         pulumi.Any(aws_iam_role.Role1.Arn),
						DocumentVersion: pulumi.String("version1"),
						TargetAccount:   pulumi.String("RESPONSE_PLAN_OWNER_ACCOUNT"),
						Parameters: ssmincidents.ResponsePlanActionSsmAutomationParameterArray{
							&ssmincidents.ResponsePlanActionSsmAutomationParameterArgs{
								Name: pulumi.String("key"),
								Values: pulumi.StringArray{
									pulumi.String("value1"),
									pulumi.String("value2"),
								},
							},
							&ssmincidents.ResponsePlanActionSsmAutomationParameterArgs{
								Name: pulumi.String("foo"),
								Values: pulumi.StringArray{
									pulumi.String("bar"),
								},
							},
						},
						DynamicParameters: pulumi.StringMap{
							"someKey":    pulumi.String("INVOLVED_RESOURCES"),
							"anotherKey": pulumi.String("INCIDENT_RECORD_ARN"),
						},
					},
				},
			},
			Integration: &ssmincidents.ResponsePlanIntegrationArgs{
				Pagerduties: ssmincidents.ResponsePlanIntegrationPagerdutyArray{
					&ssmincidents.ResponsePlanIntegrationPagerdutyArgs{
						Name:      pulumi.String("pagerdutyIntergration"),
						ServiceId: pulumi.String("example"),
						SecretId:  pulumi.String("example"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_ssmincidents_replication_set.Example,
		}))
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
import com.pulumi.aws.ssmincidents.ResponsePlan;
import com.pulumi.aws.ssmincidents.ResponsePlanArgs;
import com.pulumi.aws.ssmincidents.inputs.ResponsePlanIncidentTemplateArgs;
import com.pulumi.aws.ssmincidents.inputs.ResponsePlanActionArgs;
import com.pulumi.aws.ssmincidents.inputs.ResponsePlanIntegrationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var example = new ResponsePlan("example", ResponsePlanArgs.builder()        
            .incidentTemplate(ResponsePlanIncidentTemplateArgs.builder()
                .title("title")
                .impact("3")
                .dedupeString("dedupe")
                .incidentTags(Map.of("key", "value"))
                .notificationTargets(                
                    ResponsePlanIncidentTemplateNotificationTargetArgs.builder()
                        .snsTopicArn(aws_sns_topic.example1().arn())
                        .build(),
                    ResponsePlanIncidentTemplateNotificationTargetArgs.builder()
                        .snsTopicArn(aws_sns_topic.example2().arn())
                        .build())
                .summary("summary")
                .build())
            .displayName("display name")
            .chatChannels(aws_sns_topic.topic().arn())
            .engagements("arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1")
            .action(ResponsePlanActionArgs.builder()
                .ssmAutomations(ResponsePlanActionSsmAutomationArgs.builder()
                    .documentName(aws_ssm_document.document1().name())
                    .roleArn(aws_iam_role.role1().arn())
                    .documentVersion("version1")
                    .targetAccount("RESPONSE_PLAN_OWNER_ACCOUNT")
                    .parameters(                    
                        ResponsePlanActionSsmAutomationParameterArgs.builder()
                            .name("key")
                            .values(                            
                                "value1",
                                "value2")
                            .build(),
                        ResponsePlanActionSsmAutomationParameterArgs.builder()
                            .name("foo")
                            .values("bar")
                            .build())
                    .dynamicParameters(Map.ofEntries(
                        Map.entry("someKey", "INVOLVED_RESOURCES"),
                        Map.entry("anotherKey", "INCIDENT_RECORD_ARN")
                    ))
                    .build())
                .build())
            .integration(ResponsePlanIntegrationArgs.builder()
                .pagerduties(ResponsePlanIntegrationPagerdutyArgs.builder()
                    .name("pagerdutyIntergration")
                    .serviceId("example")
                    .secretId("example")
                    .build())
                .build())
            .tags(Map.of("key", "value"))
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_ssmincidents_replication_set.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssmincidents:ResponsePlan
    properties:
      incidentTemplate:
        title: title
        impact: '3'
        dedupeString: dedupe
        incidentTags:
          key: value
        notificationTargets:
          - snsTopicArn: ${aws_sns_topic.example1.arn}
          - snsTopicArn: ${aws_sns_topic.example2.arn}
        summary: summary
      displayName: display name
      chatChannels:
        - ${aws_sns_topic.topic.arn}
      engagements:
        - arn:aws:ssm-contacts:us-east-2:111122223333:contact/test1
      action:
        ssmAutomations:
          - documentName: ${aws_ssm_document.document1.name}
            roleArn: ${aws_iam_role.role1.arn}
            documentVersion: version1
            targetAccount: RESPONSE_PLAN_OWNER_ACCOUNT
            parameters:
              - name: key
                values:
                  - value1
                  - value2
              - name: foo
                values:
                  - bar
            dynamicParameters:
              someKey: INVOLVED_RESOURCES
              anotherKey: INCIDENT_RECORD_ARN
      integration:
        pagerduties:
          - name: pagerdutyIntergration
            serviceId: example
            secretId: example
      tags:
        key: value
    options:
      dependson:
        - ${aws_ssmincidents_replication_set.example}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an Incident Manager response plan using the response plan ARN. You can find the response plan ARN in the AWS Management Console. For example:

```sh
 $ pulumi import aws:ssmincidents/responsePlan:ResponsePlan responsePlanName ARNValue
```
 