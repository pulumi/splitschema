Provides a budget action resource. Budget actions are cost savings controls that run either automatically on your behalf or by using a workflow approval process.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["ec2:Describe*"],
        resources: ["*"],
    }],
});
const examplePolicy = new aws.iam.Policy("examplePolicy", {
    description: "My example policy",
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
const current = aws.getPartition({});
const assumeRole = current.then(current => aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: [`budgets.${current.dnsSuffix}`],
        }],
        actions: ["sts:AssumeRole"],
    }],
}));
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleBudget = new aws.budgets.Budget("exampleBudget", {
    budgetType: "USAGE",
    limitAmount: "10.0",
    limitUnit: "dollars",
    timePeriodStart: "2006-01-02_15:04",
    timeUnit: "MONTHLY",
});
const exampleBudgetAction = new aws.budgets.BudgetAction("exampleBudgetAction", {
    budgetName: exampleBudget.name,
    actionType: "APPLY_IAM_POLICY",
    approvalModel: "AUTOMATIC",
    notificationType: "ACTUAL",
    executionRoleArn: exampleRole.arn,
    actionThreshold: {
        actionThresholdType: "ABSOLUTE_VALUE",
        actionThresholdValue: 100,
    },
    definition: {
        iamActionDefinition: {
            policyArn: examplePolicy.arn,
            roles: [exampleRole.name],
        },
    },
    subscribers: [{
        address: "example@example.example",
        subscriptionType: "EMAIL",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["ec2:Describe*"],
    resources=["*"],
)])
example_policy = aws.iam.Policy("examplePolicy",
    description="My example policy",
    policy=example_policy_document.json)
current = aws.get_partition()
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=[f"budgets.{current.dns_suffix}"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_budget = aws.budgets.Budget("exampleBudget",
    budget_type="USAGE",
    limit_amount="10.0",
    limit_unit="dollars",
    time_period_start="2006-01-02_15:04",
    time_unit="MONTHLY")
example_budget_action = aws.budgets.BudgetAction("exampleBudgetAction",
    budget_name=example_budget.name,
    action_type="APPLY_IAM_POLICY",
    approval_model="AUTOMATIC",
    notification_type="ACTUAL",
    execution_role_arn=example_role.arn,
    action_threshold=aws.budgets.BudgetActionActionThresholdArgs(
        action_threshold_type="ABSOLUTE_VALUE",
        action_threshold_value=100,
    ),
    definition=aws.budgets.BudgetActionDefinitionArgs(
        iam_action_definition=aws.budgets.BudgetActionDefinitionIamActionDefinitionArgs(
            policy_arn=example_policy.arn,
            roles=[example_role.name],
        ),
    ),
    subscribers=[aws.budgets.BudgetActionSubscriberArgs(
        address="example@example.example",
        subscription_type="EMAIL",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "ec2:Describe*",
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
        Description = "My example policy",
        PolicyDocument = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var current = Aws.GetPartition.Invoke();

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
                            $"budgets.{current.Apply(getPartitionResult => getPartitionResult.DnsSuffix)}",
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

    var exampleBudget = new Aws.Budgets.Budget("exampleBudget", new()
    {
        BudgetType = "USAGE",
        LimitAmount = "10.0",
        LimitUnit = "dollars",
        TimePeriodStart = "2006-01-02_15:04",
        TimeUnit = "MONTHLY",
    });

    var exampleBudgetAction = new Aws.Budgets.BudgetAction("exampleBudgetAction", new()
    {
        BudgetName = exampleBudget.Name,
        ActionType = "APPLY_IAM_POLICY",
        ApprovalModel = "AUTOMATIC",
        NotificationType = "ACTUAL",
        ExecutionRoleArn = exampleRole.Arn,
        ActionThreshold = new Aws.Budgets.Inputs.BudgetActionActionThresholdArgs
        {
            ActionThresholdType = "ABSOLUTE_VALUE",
            ActionThresholdValue = 100,
        },
        Definition = new Aws.Budgets.Inputs.BudgetActionDefinitionArgs
        {
            IamActionDefinition = new Aws.Budgets.Inputs.BudgetActionDefinitionIamActionDefinitionArgs
            {
                PolicyArn = examplePolicy.Arn,
                Roles = new[]
                {
                    exampleRole.Name,
                },
            },
        },
        Subscribers = new[]
        {
            new Aws.Budgets.Inputs.BudgetActionSubscriberArgs
            {
                Address = "example@example.example",
                SubscriptionType = "EMAIL",
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/budgets"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"ec2:Describe*",
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
			Description: pulumi.String("My example policy"),
			Policy:      *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		current, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								fmt.Sprintf("budgets.%v", current.DnsSuffix),
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
		exampleBudget, err := budgets.NewBudget(ctx, "exampleBudget", &budgets.BudgetArgs{
			BudgetType:      pulumi.String("USAGE"),
			LimitAmount:     pulumi.String("10.0"),
			LimitUnit:       pulumi.String("dollars"),
			TimePeriodStart: pulumi.String("2006-01-02_15:04"),
			TimeUnit:        pulumi.String("MONTHLY"),
		})
		if err != nil {
			return err
		}
		_, err = budgets.NewBudgetAction(ctx, "exampleBudgetAction", &budgets.BudgetActionArgs{
			BudgetName:       exampleBudget.Name,
			ActionType:       pulumi.String("APPLY_IAM_POLICY"),
			ApprovalModel:    pulumi.String("AUTOMATIC"),
			NotificationType: pulumi.String("ACTUAL"),
			ExecutionRoleArn: exampleRole.Arn,
			ActionThreshold: &budgets.BudgetActionActionThresholdArgs{
				ActionThresholdType:  pulumi.String("ABSOLUTE_VALUE"),
				ActionThresholdValue: pulumi.Float64(100),
			},
			Definition: &budgets.BudgetActionDefinitionArgs{
				IamActionDefinition: &budgets.BudgetActionDefinitionIamActionDefinitionArgs{
					PolicyArn: examplePolicy.Arn,
					Roles: pulumi.StringArray{
						exampleRole.Name,
					},
				},
			},
			Subscribers: budgets.BudgetActionSubscriberArray{
				&budgets.BudgetActionSubscriberArgs{
					Address:          pulumi.String("example@example.example"),
					SubscriptionType: pulumi.String("EMAIL"),
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
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.budgets.Budget;
import com.pulumi.aws.budgets.BudgetArgs;
import com.pulumi.aws.budgets.BudgetAction;
import com.pulumi.aws.budgets.BudgetActionArgs;
import com.pulumi.aws.budgets.inputs.BudgetActionActionThresholdArgs;
import com.pulumi.aws.budgets.inputs.BudgetActionDefinitionArgs;
import com.pulumi.aws.budgets.inputs.BudgetActionDefinitionIamActionDefinitionArgs;
import com.pulumi.aws.budgets.inputs.BudgetActionSubscriberArgs;
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
        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("ec2:Describe*")
                .resources("*")
                .build())
            .build());

        var examplePolicy = new Policy("examplePolicy", PolicyArgs.builder()        
            .description("My example policy")
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var current = AwsFunctions.getPartition();

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers(String.format("budgets.%s", current.applyValue(getPartitionResult -> getPartitionResult.dnsSuffix())))
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleBudget = new Budget("exampleBudget", BudgetArgs.builder()        
            .budgetType("USAGE")
            .limitAmount("10.0")
            .limitUnit("dollars")
            .timePeriodStart("2006-01-02_15:04")
            .timeUnit("MONTHLY")
            .build());

        var exampleBudgetAction = new BudgetAction("exampleBudgetAction", BudgetActionArgs.builder()        
            .budgetName(exampleBudget.name())
            .actionType("APPLY_IAM_POLICY")
            .approvalModel("AUTOMATIC")
            .notificationType("ACTUAL")
            .executionRoleArn(exampleRole.arn())
            .actionThreshold(BudgetActionActionThresholdArgs.builder()
                .actionThresholdType("ABSOLUTE_VALUE")
                .actionThresholdValue(100)
                .build())
            .definition(BudgetActionDefinitionArgs.builder()
                .iamActionDefinition(BudgetActionDefinitionIamActionDefinitionArgs.builder()
                    .policyArn(examplePolicy.arn())
                    .roles(exampleRole.name())
                    .build())
                .build())
            .subscribers(BudgetActionSubscriberArgs.builder()
                .address("example@example.example")
                .subscriptionType("EMAIL")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBudgetAction:
    type: aws:budgets:BudgetAction
    properties:
      budgetName: ${exampleBudget.name}
      actionType: APPLY_IAM_POLICY
      approvalModel: AUTOMATIC
      notificationType: ACTUAL
      executionRoleArn: ${exampleRole.arn}
      actionThreshold:
        actionThresholdType: ABSOLUTE_VALUE
        actionThresholdValue: 100
      definition:
        iamActionDefinition:
          policyArn: ${examplePolicy.arn}
          roles:
            - ${exampleRole.name}
      subscribers:
        - address: example@example.example
          subscriptionType: EMAIL
  examplePolicy:
    type: aws:iam:Policy
    properties:
      description: My example policy
      policy: ${examplePolicyDocument.json}
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleBudget:
    type: aws:budgets:Budget
    properties:
      budgetType: USAGE
      limitAmount: '10.0'
      limitUnit: dollars
      timePeriodStart: 2006-01-02_15:04
      timeUnit: MONTHLY
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - ec2:Describe*
            resources:
              - '*'
  current:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - budgets.${current.dnsSuffix}
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import budget actions using `AccountID:ActionID:BudgetName`. For example:

```sh
 $ pulumi import aws:budgets/budgetAction:BudgetAction myBudget 123456789012:some-id:myBudget
```
 