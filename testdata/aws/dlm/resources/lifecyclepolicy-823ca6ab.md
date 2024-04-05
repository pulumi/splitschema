Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.dlm.LifecyclePolicy;
import com.pulumi.aws.dlm.LifecyclePolicyArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsArgs;
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
                    .identifiers("dlm.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var dlmLifecycleRole = new Role("dlmLifecycleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var dlmLifecyclePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "ec2:CreateSnapshot",
                        "ec2:CreateSnapshots",
                        "ec2:DeleteSnapshot",
                        "ec2:DescribeInstances",
                        "ec2:DescribeVolumes",
                        "ec2:DescribeSnapshots")
                    .resources("*")
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("ec2:CreateTags")
                    .resources("arn:aws:ec2:*::snapshot/*")
                    .build())
            .build());

        var dlmLifecycleRolePolicy = new RolePolicy("dlmLifecycleRolePolicy", RolePolicyArgs.builder()        
            .role(dlmLifecycleRole.id())
            .policy(dlmLifecyclePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var example = new LifecyclePolicy("example", LifecyclePolicyArgs.builder()        
            .description("example DLM lifecycle policy")
            .executionRoleArn(dlmLifecycleRole.arn())
            .state("ENABLED")
            .policyDetails(LifecyclePolicyPolicyDetailsArgs.builder()
                .resourceTypes("VOLUME")
                .schedules(LifecyclePolicyPolicyDetailsScheduleArgs.builder()
                    .name("2 weeks of daily snapshots")
                    .createRule(LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs.builder()
                        .interval(24)
                        .intervalUnit("HOURS")
                        .times("23:45")
                        .build())
                    .retainRule(LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs.builder()
                        .count(14)
                        .build())
                    .tagsToAdd(Map.of("SnapshotCreator", "DLM"))
                    .copyTags(false)
                    .build())
                .targetTags(Map.of("Snapshot", "true"))
                .build())
            .build());

    }
}
```
```yaml
resources:
  dlmLifecycleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  dlmLifecycleRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${dlmLifecycleRole.id}
      policy: ${dlmLifecyclePolicyDocument.json}
  example:
    type: aws:dlm:LifecyclePolicy
    properties:
      description: example DLM lifecycle policy
      executionRoleArn: ${dlmLifecycleRole.arn}
      state: ENABLED
      policyDetails:
        resourceTypes:
          - VOLUME
        schedules:
          - name: 2 weeks of daily snapshots
            createRule:
              interval: 24
              intervalUnit: HOURS
              times:
                - 23:45
            retainRule:
              count: 14
            tagsToAdd:
              SnapshotCreator: DLM
            copyTags: false
        targetTags:
          Snapshot: 'true'
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
                  - dlm.amazonaws.com
            actions:
              - sts:AssumeRole
  dlmLifecyclePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - ec2:CreateSnapshot
              - ec2:CreateSnapshots
              - ec2:DeleteSnapshot
              - ec2:DescribeInstances
              - ec2:DescribeVolumes
              - ec2:DescribeSnapshots
            resources:
              - '*'
          - effect: Allow
            actions:
              - ec2:CreateTags
            resources:
              - arn:aws:ec2:*::snapshot/*
```
{{% /example %}}
{{% example %}}
### Example Cross-Region Snapshot Copy Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.dlm.LifecyclePolicy;
import com.pulumi.aws.dlm.LifecyclePolicyArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var key = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("Enable IAM User Permissions")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(String.format("arn:aws:iam::%s:root", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                    .build())
                .actions("kms:*")
                .resources("*")
                .build())
            .build());

        var dlmCrossRegionCopyCmk = new Key("dlmCrossRegionCopyCmk", KeyArgs.builder()        
            .description("Example Alternate Region KMS Key")
            .policy(key.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build(), CustomResourceOptions.builder()
                .provider(aws.alternate())
                .build());

        var example = new LifecyclePolicy("example", LifecyclePolicyArgs.builder()        
            .description("example DLM lifecycle policy")
            .executionRoleArn(aws_iam_role.dlm_lifecycle_role().arn())
            .state("ENABLED")
            .policyDetails(LifecyclePolicyPolicyDetailsArgs.builder()
                .resourceTypes("VOLUME")
                .schedules(LifecyclePolicyPolicyDetailsScheduleArgs.builder()
                    .name("2 weeks of daily snapshots")
                    .createRule(LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs.builder()
                        .interval(24)
                        .intervalUnit("HOURS")
                        .times("23:45")
                        .build())
                    .retainRule(LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs.builder()
                        .count(14)
                        .build())
                    .tagsToAdd(Map.of("SnapshotCreator", "DLM"))
                    .copyTags(false)
                    .crossRegionCopyRules(LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleArgs.builder()
                        .target("us-west-2")
                        .encrypted(true)
                        .cmkArn(dlmCrossRegionCopyCmk.arn())
                        .copyTags(true)
                        .retainRule(LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleArgs.builder()
                            .interval(30)
                            .intervalUnit("DAYS")
                            .build())
                        .build())
                    .build())
                .targetTags(Map.of("Snapshot", "true"))
                .build())
            .build());

    }
}
```
```yaml
resources:
  dlmCrossRegionCopyCmk:
    type: aws:kms:Key
    properties:
      description: Example Alternate Region KMS Key
      policy: ${key.json}
    options:
      provider: ${aws.alternate}
  example:
    type: aws:dlm:LifecyclePolicy
    properties:
      description: example DLM lifecycle policy
      executionRoleArn: ${aws_iam_role.dlm_lifecycle_role.arn}
      state: ENABLED
      policyDetails:
        resourceTypes:
          - VOLUME
        schedules:
          - name: 2 weeks of daily snapshots
            createRule:
              interval: 24
              intervalUnit: HOURS
              times:
                - 23:45
            retainRule:
              count: 14
            tagsToAdd:
              SnapshotCreator: DLM
            copyTags: false
            crossRegionCopyRules:
              - target: us-west-2
                encrypted: true
                cmkArn: ${dlmCrossRegionCopyCmk.arn}
                copyTags: true
                retainRule:
                  interval: 30
                  intervalUnit: DAYS
        targetTags:
          Snapshot: 'true'
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  key:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: Enable IAM User Permissions
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - arn:aws:iam::${current.accountId}:root
            actions:
              - kms:*
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### Example Event Based Policy Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const exampleLifecyclePolicy = new aws.dlm.LifecyclePolicy("exampleLifecyclePolicy", {
    description: "tf-acc-basic",
    executionRoleArn: aws_iam_role.example.arn,
    policyDetails: {
        policyType: "EVENT_BASED_POLICY",
        action: {
            name: "tf-acc-basic",
            crossRegionCopies: [{
                encryptionConfiguration: {},
                retainRule: {
                    interval: 15,
                    intervalUnit: "MONTHS",
                },
                target: "us-east-1",
            }],
        },
        eventSource: {
            type: "MANAGED_CWE",
            parameters: {
                descriptionRegex: "^.*Created for policy: policy-1234567890abcdef0.*$",
                eventType: "shareSnapshot",
                snapshotOwners: [current.then(current => current.accountId)],
            },
        },
    },
});
const examplePolicy = aws.iam.getPolicy({
    name: "AWSDataLifecycleManagerServiceRole",
});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment", {
    role: aws_iam_role.example.id,
    policyArn: examplePolicy.then(examplePolicy => examplePolicy.arn),
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_lifecycle_policy = aws.dlm.LifecyclePolicy("exampleLifecyclePolicy",
    description="tf-acc-basic",
    execution_role_arn=aws_iam_role["example"]["arn"],
    policy_details=aws.dlm.LifecyclePolicyPolicyDetailsArgs(
        policy_type="EVENT_BASED_POLICY",
        action=aws.dlm.LifecyclePolicyPolicyDetailsActionArgs(
            name="tf-acc-basic",
            cross_region_copies=[aws.dlm.LifecyclePolicyPolicyDetailsActionCrossRegionCopyArgs(
                encryption_configuration=aws.dlm.LifecyclePolicyPolicyDetailsActionCrossRegionCopyEncryptionConfigurationArgs(),
                retain_rule=aws.dlm.LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs(
                    interval=15,
                    interval_unit="MONTHS",
                ),
                target="us-east-1",
            )],
        ),
        event_source=aws.dlm.LifecyclePolicyPolicyDetailsEventSourceArgs(
            type="MANAGED_CWE",
            parameters=aws.dlm.LifecyclePolicyPolicyDetailsEventSourceParametersArgs(
                description_regex="^.*Created for policy: policy-1234567890abcdef0.*$",
                event_type="shareSnapshot",
                snapshot_owners=[current.account_id],
            ),
        ),
    ))
example_policy = aws.iam.get_policy(name="AWSDataLifecycleManagerServiceRole")
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    role=aws_iam_role["example"]["id"],
    policy_arn=example_policy.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var exampleLifecyclePolicy = new Aws.Dlm.LifecyclePolicy("exampleLifecyclePolicy", new()
    {
        Description = "tf-acc-basic",
        ExecutionRoleArn = aws_iam_role.Example.Arn,
        PolicyDetails = new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsArgs
        {
            PolicyType = "EVENT_BASED_POLICY",
            Action = new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsActionArgs
            {
                Name = "tf-acc-basic",
                CrossRegionCopies = new[]
                {
                    new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsActionCrossRegionCopyArgs
                    {
                        EncryptionConfiguration = null,
                        RetainRule = new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs
                        {
                            Interval = 15,
                            IntervalUnit = "MONTHS",
                        },
                        Target = "us-east-1",
                    },
                },
            },
            EventSource = new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsEventSourceArgs
            {
                Type = "MANAGED_CWE",
                Parameters = new Aws.Dlm.Inputs.LifecyclePolicyPolicyDetailsEventSourceParametersArgs
                {
                    DescriptionRegex = "^.*Created for policy: policy-1234567890abcdef0.*$",
                    EventType = "shareSnapshot",
                    SnapshotOwners = new[]
                    {
                        current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                    },
                },
            },
        },
    });

    var examplePolicy = Aws.Iam.GetPolicy.Invoke(new()
    {
        Name = "AWSDataLifecycleManagerServiceRole",
    });

    var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    {
        Role = aws_iam_role.Example.Id,
        PolicyArn = examplePolicy.Apply(getPolicyResult => getPolicyResult.Arn),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dlm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = dlm.NewLifecyclePolicy(ctx, "exampleLifecyclePolicy", &dlm.LifecyclePolicyArgs{
			Description:      pulumi.String("tf-acc-basic"),
			ExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			PolicyDetails: &dlm.LifecyclePolicyPolicyDetailsArgs{
				PolicyType: pulumi.String("EVENT_BASED_POLICY"),
				Action: &dlm.LifecyclePolicyPolicyDetailsActionArgs{
					Name: pulumi.String("tf-acc-basic"),
					CrossRegionCopies: dlm.LifecyclePolicyPolicyDetailsActionCrossRegionCopyArray{
						&dlm.LifecyclePolicyPolicyDetailsActionCrossRegionCopyArgs{
							EncryptionConfiguration: nil,
							RetainRule: &dlm.LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs{
								Interval:     pulumi.Int(15),
								IntervalUnit: pulumi.String("MONTHS"),
							},
							Target: pulumi.String("us-east-1"),
						},
					},
				},
				EventSource: &dlm.LifecyclePolicyPolicyDetailsEventSourceArgs{
					Type: pulumi.String("MANAGED_CWE"),
					Parameters: &dlm.LifecyclePolicyPolicyDetailsEventSourceParametersArgs{
						DescriptionRegex: pulumi.String("^.*Created for policy: policy-1234567890abcdef0.*$"),
						EventType:        pulumi.String("shareSnapshot"),
						SnapshotOwners: pulumi.StringArray{
							*pulumi.String(current.AccountId),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		examplePolicy, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
			Name: pulumi.StringRef("AWSDataLifecycleManagerServiceRole"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      pulumi.Any(aws_iam_role.Example.Id),
			PolicyArn: *pulumi.String(examplePolicy.Arn),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.dlm.LifecyclePolicy;
import com.pulumi.aws.dlm.LifecyclePolicyArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsActionArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsEventSourceArgs;
import com.pulumi.aws.dlm.inputs.LifecyclePolicyPolicyDetailsEventSourceParametersArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var exampleLifecyclePolicy = new LifecyclePolicy("exampleLifecyclePolicy", LifecyclePolicyArgs.builder()        
            .description("tf-acc-basic")
            .executionRoleArn(aws_iam_role.example().arn())
            .policyDetails(LifecyclePolicyPolicyDetailsArgs.builder()
                .policyType("EVENT_BASED_POLICY")
                .action(LifecyclePolicyPolicyDetailsActionArgs.builder()
                    .name("tf-acc-basic")
                    .crossRegionCopies(LifecyclePolicyPolicyDetailsActionCrossRegionCopyArgs.builder()
                        .encryptionConfiguration()
                        .retainRule(LifecyclePolicyPolicyDetailsActionCrossRegionCopyRetainRuleArgs.builder()
                            .interval(15)
                            .intervalUnit("MONTHS")
                            .build())
                        .target("us-east-1")
                        .build())
                    .build())
                .eventSource(LifecyclePolicyPolicyDetailsEventSourceArgs.builder()
                    .type("MANAGED_CWE")
                    .parameters(LifecyclePolicyPolicyDetailsEventSourceParametersArgs.builder()
                        .descriptionRegex("^.*Created for policy: policy-1234567890abcdef0.*$")
                        .eventType("shareSnapshot")
                        .snapshotOwners(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                        .build())
                    .build())
                .build())
            .build());

        final var examplePolicy = IamFunctions.getPolicy(GetPolicyArgs.builder()
            .name("AWSDataLifecycleManagerServiceRole")
            .build());

        var exampleRolePolicyAttachment = new RolePolicyAttachment("exampleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(aws_iam_role.example().id())
            .policyArn(examplePolicy.applyValue(getPolicyResult -> getPolicyResult.arn()))
            .build());

    }
}
```
```yaml
resources:
  exampleLifecyclePolicy:
    type: aws:dlm:LifecyclePolicy
    properties:
      description: tf-acc-basic
      executionRoleArn: ${aws_iam_role.example.arn}
      policyDetails:
        policyType: EVENT_BASED_POLICY
        action:
          name: tf-acc-basic
          crossRegionCopies:
            - encryptionConfiguration: {}
              retainRule:
                interval: 15
                intervalUnit: MONTHS
              target: us-east-1
        eventSource:
          type: MANAGED_CWE
          parameters:
            descriptionRegex: '^.*Created for policy: policy-1234567890abcdef0.*$'
            eventType: shareSnapshot
            snapshotOwners:
              - ${current.accountId}
  exampleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${aws_iam_role.example.id}
      policyArn: ${examplePolicy.arn}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  examplePolicy:
    fn::invoke:
      Function: aws:iam:getPolicy
      Arguments:
        name: AWSDataLifecycleManagerServiceRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DLM lifecycle policies using their policy ID. For example:

```sh
 $ pulumi import aws:dlm/lifecyclePolicy:LifecyclePolicy example policy-abcdef12345678901
```
 