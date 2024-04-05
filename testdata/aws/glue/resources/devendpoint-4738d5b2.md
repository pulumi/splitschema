Provides a Glue Development Endpoint resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["glue.amazonaws.com"],
        }],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json)});
const exampleDevEndpoint = new aws.glue.DevEndpoint("exampleDevEndpoint", {roleArn: exampleRole.arn});
const example_AWSGlueServiceRole = new aws.iam.RolePolicyAttachment("example-AWSGlueServiceRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole",
    role: exampleRole.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["glue.amazonaws.com"],
    )],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=example_policy_document.json)
example_dev_endpoint = aws.glue.DevEndpoint("exampleDevEndpoint", role_arn=example_role.arn)
example__aws_glue_service_role = aws.iam.RolePolicyAttachment("example-AWSGlueServiceRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole",
    role=example_role.name)
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
                            "glue.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleDevEndpoint = new Aws.Glue.DevEndpoint("exampleDevEndpoint", new()
    {
        RoleArn = exampleRole.Arn,
    });

    var example_AWSGlueServiceRole = new Aws.Iam.RolePolicyAttachment("example-AWSGlueServiceRole", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole",
        Role = exampleRole.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"glue.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = glue.NewDevEndpoint(ctx, "exampleDevEndpoint", &glue.DevEndpointArgs{
			RoleArn: exampleRole.Arn,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AWSGlueServiceRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"),
			Role:      exampleRole.Name,
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
import com.pulumi.aws.glue.DevEndpoint;
import com.pulumi.aws.glue.DevEndpointArgs;
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
        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("glue.amazonaws.com")
                    .build())
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleDevEndpoint = new DevEndpoint("exampleDevEndpoint", DevEndpointArgs.builder()        
            .roleArn(exampleRole.arn())
            .build());

        var example_AWSGlueServiceRole = new RolePolicyAttachment("example-AWSGlueServiceRole", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole")
            .role(exampleRole.name())
            .build());

    }
}
```
```yaml
resources:
  exampleDevEndpoint:
    type: aws:glue:DevEndpoint
    properties:
      roleArn: ${exampleRole.arn}
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${examplePolicyDocument.json}
  example-AWSGlueServiceRole:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole
      role: ${exampleRole.name}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - glue.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a Glue Development Endpoint using the `name`. For example:

```sh
 $ pulumi import aws:glue/devEndpoint:DevEndpoint example foo
```
 