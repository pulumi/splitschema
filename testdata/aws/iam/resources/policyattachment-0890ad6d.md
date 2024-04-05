Attaches a Managed IAM Policy to user(s), role(s), and/or group(s)

!> **WARNING:** The aws.iam.PolicyAttachment resource creates **exclusive** attachments of IAM policies. Across the entire AWS account, all of the users/roles/groups to which a single policy is attached must be declared by a single aws.iam.PolicyAttachment resource. This means that even any users/roles/groups that have the attached policy via any other mechanism (including other resources managed by this provider) will have that attached policy revoked by this resource. Consider `aws.iam.RolePolicyAttachment`, `aws.iam.UserPolicyAttachment`, or `aws.iam.GroupPolicyAttachment` instead. These resources do not enforce exclusive attachment of an IAM policy.

> **NOTE:** The usage of this resource conflicts with the `aws.iam.GroupPolicyAttachment`, `aws.iam.RolePolicyAttachment`, and `aws.iam.UserPolicyAttachment` resources and will permanently show a difference if both are defined.

> **NOTE:** For a given role, this resource is incompatible with using the `aws.iam.Role` resource `managed_policy_arns` argument. When using that argument and this resource, both will attempt to manage the role's managed policy attachments and the provider will show a permanent difference.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const user = new aws.iam.User("user", {});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["ec2.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const role = new aws.iam.Role("role", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const group = new aws.iam.Group("group", {});
const policyPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["ec2:Describe*"],
        resources: ["*"],
    }],
});
const policyPolicy = new aws.iam.Policy("policyPolicy", {
    description: "A test policy",
    policy: policyPolicyDocument.then(policyPolicyDocument => policyPolicyDocument.json),
});
const test_attach = new aws.iam.PolicyAttachment("test-attach", {
    users: [user.name],
    roles: [role.name],
    groups: [group.name],
    policyArn: policyPolicy.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

user = aws.iam.User("user")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["ec2.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
role = aws.iam.Role("role", assume_role_policy=assume_role.json)
group = aws.iam.Group("group")
policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["ec2:Describe*"],
    resources=["*"],
)])
policy_policy = aws.iam.Policy("policyPolicy",
    description="A test policy",
    policy=policy_policy_document.json)
test_attach = aws.iam.PolicyAttachment("test-attach",
    users=[user.name],
    roles=[role.name],
    groups=[group.name],
    policy_arn=policy_policy.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var user = new Aws.Iam.User("user");

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
                            "ec2.amazonaws.com",
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

    var role = new Aws.Iam.Role("role", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var @group = new Aws.Iam.Group("group");

    var policyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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

    var policyPolicy = new Aws.Iam.Policy("policyPolicy", new()
    {
        Description = "A test policy",
        PolicyDocument = policyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var test_attach = new Aws.Iam.PolicyAttachment("test-attach", new()
    {
        Users = new[]
        {
            user.Name,
        },
        Roles = new[]
        {
            role.Name,
        },
        Groups = new[]
        {
            @group.Name,
        },
        PolicyArn = policyPolicy.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		user, err := iam.NewUser(ctx, "user", nil)
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
								"ec2.amazonaws.com",
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
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		group, err := iam.NewGroup(ctx, "group", nil)
		if err != nil {
			return err
		}
		policyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
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
		policyPolicy, err := iam.NewPolicy(ctx, "policyPolicy", &iam.PolicyArgs{
			Description: pulumi.String("A test policy"),
			Policy:      *pulumi.String(policyPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewPolicyAttachment(ctx, "test-attach", &iam.PolicyAttachmentArgs{
			Users: pulumi.Array{
				user.Name,
			},
			Roles: pulumi.Array{
				role.Name,
			},
			Groups: pulumi.Array{
				group.Name,
			},
			PolicyArn: policyPolicy.Arn,
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
import com.pulumi.aws.iam.User;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.Group;
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.PolicyAttachment;
import com.pulumi.aws.iam.PolicyAttachmentArgs;
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
        var user = new User("user");

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("ec2.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var role = new Role("role", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var group = new Group("group");

        final var policyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("ec2:Describe*")
                .resources("*")
                .build())
            .build());

        var policyPolicy = new Policy("policyPolicy", PolicyArgs.builder()        
            .description("A test policy")
            .policy(policyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var test_attach = new PolicyAttachment("test-attach", PolicyAttachmentArgs.builder()        
            .users(user.name())
            .roles(role.name())
            .groups(group.name())
            .policyArn(policyPolicy.arn())
            .build());

    }
}
```
```yaml
resources:
  user:
    type: aws:iam:User
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  group:
    type: aws:iam:Group
  policyPolicy:
    type: aws:iam:Policy
    properties:
      description: A test policy
      policy: ${policyPolicyDocument.json}
  test-attach:
    type: aws:iam:PolicyAttachment
    properties:
      users:
        - ${user.name}
      roles:
        - ${role.name}
      groups:
        - ${group.name}
      policyArn: ${policyPolicy.arn}
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
                  - ec2.amazonaws.com
            actions:
              - sts:AssumeRole
  policyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - ec2:Describe*
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}