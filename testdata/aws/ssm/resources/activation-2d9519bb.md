Registers an on-premises server or virtual machine with Amazon EC2 so that it can be managed using Run Command.

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
            identifiers: ["ssm.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const testRole = new aws.iam.Role("testRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const testAttach = new aws.iam.RolePolicyAttachment("testAttach", {
    role: testRole.name,
    policyArn: "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
});
const foo = new aws.ssm.Activation("foo", {
    description: "Test",
    iamRole: testRole.id,
    registrationLimit: 5,
}, {
    dependsOn: [testAttach],
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["ssm.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
test_role = aws.iam.Role("testRole", assume_role_policy=assume_role.json)
test_attach = aws.iam.RolePolicyAttachment("testAttach",
    role=test_role.name,
    policy_arn="arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore")
foo = aws.ssm.Activation("foo",
    description="Test",
    iam_role=test_role.id,
    registration_limit=5,
    opts=pulumi.ResourceOptions(depends_on=[test_attach]))
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
                            "ssm.amazonaws.com",
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

    var testRole = new Aws.Iam.Role("testRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var testAttach = new Aws.Iam.RolePolicyAttachment("testAttach", new()
    {
        Role = testRole.Name,
        PolicyArn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore",
    });

    var foo = new Aws.Ssm.Activation("foo", new()
    {
        Description = "Test",
        IamRole = testRole.Id,
        RegistrationLimit = 5,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            testAttach,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
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
								"ssm.amazonaws.com",
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
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		testAttach, err := iam.NewRolePolicyAttachment(ctx, "testAttach", &iam.RolePolicyAttachmentArgs{
			Role:      testRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"),
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewActivation(ctx, "foo", &ssm.ActivationArgs{
			Description:       pulumi.String("Test"),
			IamRole:           testRole.ID(),
			RegistrationLimit: pulumi.Int(5),
		}, pulumi.DependsOn([]pulumi.Resource{
			testAttach,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.ssm.Activation;
import com.pulumi.aws.ssm.ActivationArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("ssm.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var testRole = new Role("testRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var testAttach = new RolePolicyAttachment("testAttach", RolePolicyAttachmentArgs.builder()        
            .role(testRole.name())
            .policyArn("arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore")
            .build());

        var foo = new Activation("foo", ActivationArgs.builder()        
            .description("Test")
            .iamRole(testRole.id())
            .registrationLimit("5")
            .build(), CustomResourceOptions.builder()
                .dependsOn(testAttach)
                .build());

    }
}
```
```yaml
resources:
  testRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  testAttach:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${testRole.name}
      policyArn: arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore
  foo:
    type: aws:ssm:Activation
    properties:
      description: Test
      iamRole: ${testRole.id}
      registrationLimit: '5'
    options:
      dependson:
        - ${testAttach}
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
                  - ssm.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS SSM Activation using the `id`. For example:

```sh
 $ pulumi import aws:ssm/activation:Activation example e488f2f6-e686-4afb-8a04-ef6dfEXAMPLE
```
 -> __Note:__ The `activation_code` attribute cannot be imported.

