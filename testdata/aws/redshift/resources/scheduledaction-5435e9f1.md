{{% examples %}}
## Example Usage
{{% example %}}
### Pause Cluster Action

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["scheduler.redshift.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: [
            "redshift:PauseCluster",
            "redshift:ResumeCluster",
            "redshift:ResizeCluster",
        ],
        resources: ["*"],
    }],
});
const examplePolicy = new aws.iam.Policy("examplePolicy", {policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json)});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment", {
    policyArn: examplePolicy.arn,
    role: exampleRole.name,
});
const exampleScheduledAction = new aws.redshift.ScheduledAction("exampleScheduledAction", {
    schedule: "cron(00 23 * * ? *)",
    iamRole: exampleRole.arn,
    targetAction: {
        pauseCluster: {
            clusterIdentifier: "tf-redshift001",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["scheduler.redshift.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "redshift:PauseCluster",
        "redshift:ResumeCluster",
        "redshift:ResizeCluster",
    ],
    resources=["*"],
)])
example_policy = aws.iam.Policy("examplePolicy", policy=example_policy_document.json)
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    policy_arn=example_policy.arn,
    role=example_role.name)
example_scheduled_action = aws.redshift.ScheduledAction("exampleScheduledAction",
    schedule="cron(00 23 * * ? *)",
    iam_role=example_role.arn,
    target_action=aws.redshift.ScheduledActionTargetActionArgs(
        pause_cluster=aws.redshift.ScheduledActionTargetActionPauseClusterArgs(
            cluster_identifier="tf-redshift001",
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
                            "scheduler.redshift.amazonaws.com",
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

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "redshift:PauseCluster",
                    "redshift:ResumeCluster",
                    "redshift:ResizeCluster",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var examplePolicy = new Aws.Iam.Policy("examplePolicy", new()
    {
        PolicyDocument = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    {
        PolicyArn = examplePolicy.Arn,
        Role = exampleRole.Name,
    });

    var exampleScheduledAction = new Aws.RedShift.ScheduledAction("exampleScheduledAction", new()
    {
        Schedule = "cron(00 23 * * ? *)",
        IamRole = exampleRole.Arn,
        TargetAction = new Aws.RedShift.Inputs.ScheduledActionTargetActionArgs
        {
            PauseCluster = new Aws.RedShift.Inputs.ScheduledActionTargetActionPauseClusterArgs
            {
                ClusterIdentifier = "tf-redshift001",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
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
								"scheduler.redshift.amazonaws.com",
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
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"redshift:PauseCluster",
						"redshift:ResumeCluster",
						"redshift:ResizeCluster",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		examplePolicy, err := iam.NewPolicy(ctx, "examplePolicy", &iam.PolicyArgs{
			Policy: *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: examplePolicy.Arn,
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewScheduledAction(ctx, "exampleScheduledAction", &redshift.ScheduledActionArgs{
			Schedule: pulumi.String("cron(00 23 * * ? *)"),
			IamRole:  exampleRole.Arn,
			TargetAction: &redshift.ScheduledActionTargetActionArgs{
				PauseCluster: &redshift.ScheduledActionTargetActionPauseClusterArgs{
					ClusterIdentifier: pulumi.String("tf-redshift001"),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.redshift.ScheduledAction;
import com.pulumi.aws.redshift.ScheduledActionArgs;
import com.pulumi.aws.redshift.inputs.ScheduledActionTargetActionArgs;
import com.pulumi.aws.redshift.inputs.ScheduledActionTargetActionPauseClusterArgs;
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
                    .identifiers("scheduler.redshift.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "redshift:PauseCluster",
                    "redshift:ResumeCluster",
                    "redshift:ResizeCluster")
                .resources("*")
                .build())
            .build());

        var examplePolicy = new Policy("examplePolicy", PolicyArgs.builder()        
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleRolePolicyAttachment = new RolePolicyAttachment("exampleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .policyArn(examplePolicy.arn())
            .role(exampleRole.name())
            .build());

        var exampleScheduledAction = new ScheduledAction("exampleScheduledAction", ScheduledActionArgs.builder()        
            .schedule("cron(00 23 * * ? *)")
            .iamRole(exampleRole.arn())
            .targetAction(ScheduledActionTargetActionArgs.builder()
                .pauseCluster(ScheduledActionTargetActionPauseClusterArgs.builder()
                    .clusterIdentifier("tf-redshift001")
                    .build())
                .build())
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
  examplePolicy:
    type: aws:iam:Policy
    properties:
      policy: ${examplePolicyDocument.json}
  exampleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: ${examplePolicy.arn}
      role: ${exampleRole.name}
  exampleScheduledAction:
    type: aws:redshift:ScheduledAction
    properties:
      schedule: cron(00 23 * * ? *)
      iamRole: ${exampleRole.arn}
      targetAction:
        pauseCluster:
          clusterIdentifier: tf-redshift001
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
                  - scheduler.redshift.amazonaws.com
            actions:
              - sts:AssumeRole
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - redshift:PauseCluster
              - redshift:ResumeCluster
              - redshift:ResizeCluster
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### Resize Cluster Action

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.ScheduledAction("example", {
    schedule: "cron(00 23 * * ? *)",
    iamRole: aws_iam_role.example.arn,
    targetAction: {
        resizeCluster: {
            clusterIdentifier: "tf-redshift001",
            clusterType: "multi-node",
            nodeType: "dc1.large",
            numberOfNodes: 2,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.ScheduledAction("example",
    schedule="cron(00 23 * * ? *)",
    iam_role=aws_iam_role["example"]["arn"],
    target_action=aws.redshift.ScheduledActionTargetActionArgs(
        resize_cluster=aws.redshift.ScheduledActionTargetActionResizeClusterArgs(
            cluster_identifier="tf-redshift001",
            cluster_type="multi-node",
            node_type="dc1.large",
            number_of_nodes=2,
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
    var example = new Aws.RedShift.ScheduledAction("example", new()
    {
        Schedule = "cron(00 23 * * ? *)",
        IamRole = aws_iam_role.Example.Arn,
        TargetAction = new Aws.RedShift.Inputs.ScheduledActionTargetActionArgs
        {
            ResizeCluster = new Aws.RedShift.Inputs.ScheduledActionTargetActionResizeClusterArgs
            {
                ClusterIdentifier = "tf-redshift001",
                ClusterType = "multi-node",
                NodeType = "dc1.large",
                NumberOfNodes = 2,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewScheduledAction(ctx, "example", &redshift.ScheduledActionArgs{
			Schedule: pulumi.String("cron(00 23 * * ? *)"),
			IamRole:  pulumi.Any(aws_iam_role.Example.Arn),
			TargetAction: &redshift.ScheduledActionTargetActionArgs{
				ResizeCluster: &redshift.ScheduledActionTargetActionResizeClusterArgs{
					ClusterIdentifier: pulumi.String("tf-redshift001"),
					ClusterType:       pulumi.String("multi-node"),
					NodeType:          pulumi.String("dc1.large"),
					NumberOfNodes:     pulumi.Int(2),
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
import com.pulumi.aws.redshift.ScheduledAction;
import com.pulumi.aws.redshift.ScheduledActionArgs;
import com.pulumi.aws.redshift.inputs.ScheduledActionTargetActionArgs;
import com.pulumi.aws.redshift.inputs.ScheduledActionTargetActionResizeClusterArgs;
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
        var example = new ScheduledAction("example", ScheduledActionArgs.builder()        
            .schedule("cron(00 23 * * ? *)")
            .iamRole(aws_iam_role.example().arn())
            .targetAction(ScheduledActionTargetActionArgs.builder()
                .resizeCluster(ScheduledActionTargetActionResizeClusterArgs.builder()
                    .clusterIdentifier("tf-redshift001")
                    .clusterType("multi-node")
                    .nodeType("dc1.large")
                    .numberOfNodes(2)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:ScheduledAction
    properties:
      schedule: cron(00 23 * * ? *)
      iamRole: ${aws_iam_role.example.arn}
      targetAction:
        resizeCluster:
          clusterIdentifier: tf-redshift001
          clusterType: multi-node
          nodeType: dc1.large
          numberOfNodes: 2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Scheduled Action using the `name`. For example:

```sh
 $ pulumi import aws:redshift/scheduledAction:ScheduledAction example tf-redshift-scheduled-action
```
 