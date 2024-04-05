Provides a resource to manage AWS Secrets Manager secret policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSecret = new aws.secretsmanager.Secret("exampleSecret", {});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "EnableAnotherAWSAccountToReadTheSecret",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["arn:aws:iam::123456789012:root"],
        }],
        actions: ["secretsmanager:GetSecretValue"],
        resources: ["*"],
    }],
});
const exampleSecretPolicy = new aws.secretsmanager.SecretPolicy("exampleSecretPolicy", {
    secretArn: exampleSecret.arn,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_secret = aws.secretsmanager.Secret("exampleSecret")
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="EnableAnotherAWSAccountToReadTheSecret",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["arn:aws:iam::123456789012:root"],
    )],
    actions=["secretsmanager:GetSecretValue"],
    resources=["*"],
)])
example_secret_policy = aws.secretsmanager.SecretPolicy("exampleSecretPolicy",
    secret_arn=example_secret.arn,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSecret = new Aws.SecretsManager.Secret("exampleSecret");

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "EnableAnotherAWSAccountToReadTheSecret",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "arn:aws:iam::123456789012:root",
                        },
                    },
                },
                Actions = new[]
                {
                    "secretsmanager:GetSecretValue",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var exampleSecretPolicy = new Aws.SecretsManager.SecretPolicy("exampleSecretPolicy", new()
    {
        SecretArn = exampleSecret.Arn,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSecret, err := secretsmanager.NewSecret(ctx, "exampleSecret", nil)
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("EnableAnotherAWSAccountToReadTheSecret"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								"arn:aws:iam::123456789012:root",
							},
						},
					},
					Actions: []string{
						"secretsmanager:GetSecretValue",
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
		_, err = secretsmanager.NewSecretPolicy(ctx, "exampleSecretPolicy", &secretsmanager.SecretPolicyArgs{
			SecretArn: exampleSecret.Arn,
			Policy:    *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.secretsmanager.Secret;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.secretsmanager.SecretPolicy;
import com.pulumi.aws.secretsmanager.SecretPolicyArgs;
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
        var exampleSecret = new Secret("exampleSecret");

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("EnableAnotherAWSAccountToReadTheSecret")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("arn:aws:iam::123456789012:root")
                    .build())
                .actions("secretsmanager:GetSecretValue")
                .resources("*")
                .build())
            .build());

        var exampleSecretPolicy = new SecretPolicy("exampleSecretPolicy", SecretPolicyArgs.builder()        
            .secretArn(exampleSecret.arn())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleSecret:
    type: aws:secretsmanager:Secret
  exampleSecretPolicy:
    type: aws:secretsmanager:SecretPolicy
    properties:
      secretArn: ${exampleSecret.arn}
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: EnableAnotherAWSAccountToReadTheSecret
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - arn:aws:iam::123456789012:root
            actions:
              - secretsmanager:GetSecretValue
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_secretsmanager_secret_policy` using the secret Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:secretsmanager/secretPolicy:SecretPolicy example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
```
 