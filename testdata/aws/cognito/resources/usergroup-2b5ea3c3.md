Provides a Cognito User Group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mainUserPool = new aws.cognito.UserPool("mainUserPool", {});
const groupRolePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Federated",
            identifiers: ["cognito-identity.amazonaws.com"],
        }],
        actions: ["sts:AssumeRoleWithWebIdentity"],
        conditions: [
            {
                test: "StringEquals",
                variable: "cognito-identity.amazonaws.com:aud",
                values: ["us-east-1:12345678-dead-beef-cafe-123456790ab"],
            },
            {
                test: "ForAnyValue:StringLike",
                variable: "cognito-identity.amazonaws.com:amr",
                values: ["authenticated"],
            },
        ],
    }],
});
const groupRoleRole = new aws.iam.Role("groupRoleRole", {assumeRolePolicy: groupRolePolicyDocument.then(groupRolePolicyDocument => groupRolePolicyDocument.json)});
const mainUserGroup = new aws.cognito.UserGroup("mainUserGroup", {
    userPoolId: mainUserPool.id,
    description: "Managed by Pulumi",
    precedence: 42,
    roleArn: groupRoleRole.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

main_user_pool = aws.cognito.UserPool("mainUserPool")
group_role_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Federated",
        identifiers=["cognito-identity.amazonaws.com"],
    )],
    actions=["sts:AssumeRoleWithWebIdentity"],
    conditions=[
        aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="cognito-identity.amazonaws.com:aud",
            values=["us-east-1:12345678-dead-beef-cafe-123456790ab"],
        ),
        aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="ForAnyValue:StringLike",
            variable="cognito-identity.amazonaws.com:amr",
            values=["authenticated"],
        ),
    ],
)])
group_role_role = aws.iam.Role("groupRoleRole", assume_role_policy=group_role_policy_document.json)
main_user_group = aws.cognito.UserGroup("mainUserGroup",
    user_pool_id=main_user_pool.id,
    description="Managed by Pulumi",
    precedence=42,
    role_arn=group_role_role.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mainUserPool = new Aws.Cognito.UserPool("mainUserPool");

    var groupRolePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "Federated",
                        Identifiers = new[]
                        {
                            "cognito-identity.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRoleWithWebIdentity",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "cognito-identity.amazonaws.com:aud",
                        Values = new[]
                        {
                            "us-east-1:12345678-dead-beef-cafe-123456790ab",
                        },
                    },
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "ForAnyValue:StringLike",
                        Variable = "cognito-identity.amazonaws.com:amr",
                        Values = new[]
                        {
                            "authenticated",
                        },
                    },
                },
            },
        },
    });

    var groupRoleRole = new Aws.Iam.Role("groupRoleRole", new()
    {
        AssumeRolePolicy = groupRolePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var mainUserGroup = new Aws.Cognito.UserGroup("mainUserGroup", new()
    {
        UserPoolId = mainUserPool.Id,
        Description = "Managed by Pulumi",
        Precedence = 42,
        RoleArn = groupRoleRole.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mainUserPool, err := cognito.NewUserPool(ctx, "mainUserPool", nil)
		if err != nil {
			return err
		}
		groupRolePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Federated",
							Identifiers: []string{
								"cognito-identity.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRoleWithWebIdentity",
					},
					Conditions: []iam.GetPolicyDocumentStatementCondition{
						{
							Test:     "StringEquals",
							Variable: "cognito-identity.amazonaws.com:aud",
							Values: []string{
								"us-east-1:12345678-dead-beef-cafe-123456790ab",
							},
						},
						{
							Test:     "ForAnyValue:StringLike",
							Variable: "cognito-identity.amazonaws.com:amr",
							Values: []string{
								"authenticated",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		groupRoleRole, err := iam.NewRole(ctx, "groupRoleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(groupRolePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserGroup(ctx, "mainUserGroup", &cognito.UserGroupArgs{
			UserPoolId:  mainUserPool.ID(),
			Description: pulumi.String("Managed by Pulumi"),
			Precedence:  pulumi.Int(42),
			RoleArn:     groupRoleRole.Arn,
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.cognito.UserGroup;
import com.pulumi.aws.cognito.UserGroupArgs;
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
        var mainUserPool = new UserPool("mainUserPool");

        final var groupRolePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Federated")
                    .identifiers("cognito-identity.amazonaws.com")
                    .build())
                .actions("sts:AssumeRoleWithWebIdentity")
                .conditions(                
                    GetPolicyDocumentStatementConditionArgs.builder()
                        .test("StringEquals")
                        .variable("cognito-identity.amazonaws.com:aud")
                        .values("us-east-1:12345678-dead-beef-cafe-123456790ab")
                        .build(),
                    GetPolicyDocumentStatementConditionArgs.builder()
                        .test("ForAnyValue:StringLike")
                        .variable("cognito-identity.amazonaws.com:amr")
                        .values("authenticated")
                        .build())
                .build())
            .build());

        var groupRoleRole = new Role("groupRoleRole", RoleArgs.builder()        
            .assumeRolePolicy(groupRolePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var mainUserGroup = new UserGroup("mainUserGroup", UserGroupArgs.builder()        
            .userPoolId(mainUserPool.id())
            .description("Managed by Pulumi")
            .precedence(42)
            .roleArn(groupRoleRole.arn())
            .build());

    }
}
```
```yaml
resources:
  mainUserPool:
    type: aws:cognito:UserPool
  groupRoleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${groupRolePolicyDocument.json}
  mainUserGroup:
    type: aws:cognito:UserGroup
    properties:
      userPoolId: ${mainUserPool.id}
      description: Managed by Pulumi
      precedence: 42
      roleArn: ${groupRoleRole.arn}
variables:
  groupRolePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Federated
                identifiers:
                  - cognito-identity.amazonaws.com
            actions:
              - sts:AssumeRoleWithWebIdentity
            conditions:
              - test: StringEquals
                variable: cognito-identity.amazonaws.com:aud
                values:
                  - us-east-1:12345678-dead-beef-cafe-123456790ab
              - test: ForAnyValue:StringLike
                variable: cognito-identity.amazonaws.com:amr
                values:
                  - authenticated
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito User Groups using the `user_pool_id`/`name` attributes concatenated. For example:

```sh
 $ pulumi import aws:cognito/userGroup:UserGroup group us-east-1_vG78M4goG/user-group
```
 