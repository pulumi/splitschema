Manages a SES Identity Policy. More information about SES Sending Authorization Policies can be found in the [SES Developer Guide](https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomainIdentity = new aws.ses.DomainIdentity("exampleDomainIdentity", {domain: "example.com"});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        actions: [
            "SES:SendEmail",
            "SES:SendRawEmail",
        ],
        resources: [exampleDomainIdentity.arn],
        principals: [{
            identifiers: ["*"],
            type: "AWS",
        }],
    }],
});
const exampleIdentityPolicy = new aws.ses.IdentityPolicy("exampleIdentityPolicy", {
    identity: exampleDomainIdentity.arn,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_domain_identity = aws.ses.DomainIdentity("exampleDomainIdentity", domain="example.com")
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=[
        "SES:SendEmail",
        "SES:SendRawEmail",
    ],
    resources=[example_domain_identity.arn],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=["*"],
        type="AWS",
    )],
)])
example_identity_policy = aws.ses.IdentityPolicy("exampleIdentityPolicy",
    identity=example_domain_identity.arn,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomainIdentity = new Aws.Ses.DomainIdentity("exampleDomainIdentity", new()
    {
        Domain = "example.com",
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "SES:SendEmail",
                    "SES:SendRawEmail",
                },
                Resources = new[]
                {
                    exampleDomainIdentity.Arn,
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Identifiers = new[]
                        {
                            "*",
                        },
                        Type = "AWS",
                    },
                },
            },
        },
    });

    var exampleIdentityPolicy = new Aws.Ses.IdentityPolicy("exampleIdentityPolicy", new()
    {
        Identity = exampleDomainIdentity.Arn,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomainIdentity, err := ses.NewDomainIdentity(ctx, "exampleDomainIdentity", &ses.DomainIdentityArgs{
			Domain: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("SES:SendEmail"),
						pulumi.String("SES:SendRawEmail"),
					},
					Resources: pulumi.StringArray{
						exampleDomainIdentity.Arn,
					},
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Identifiers: pulumi.StringArray{
								pulumi.String("*"),
							},
							Type: pulumi.String("AWS"),
						},
					},
				},
			},
		}, nil)
		_, err = ses.NewIdentityPolicy(ctx, "exampleIdentityPolicy", &ses.IdentityPolicyArgs{
			Identity: exampleDomainIdentity.Arn,
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &examplePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.ses.DomainIdentity;
import com.pulumi.aws.ses.DomainIdentityArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.ses.IdentityPolicy;
import com.pulumi.aws.ses.IdentityPolicyArgs;
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
        var exampleDomainIdentity = new DomainIdentity("exampleDomainIdentity", DomainIdentityArgs.builder()        
            .domain("example.com")
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions(                
                    "SES:SendEmail",
                    "SES:SendRawEmail")
                .resources(exampleDomainIdentity.arn())
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers("*")
                    .type("AWS")
                    .build())
                .build())
            .build());

        var exampleIdentityPolicy = new IdentityPolicy("exampleIdentityPolicy", IdentityPolicyArgs.builder()        
            .identity(exampleDomainIdentity.arn())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  exampleDomainIdentity:
    type: aws:ses:DomainIdentity
    properties:
      domain: example.com
  exampleIdentityPolicy:
    type: aws:ses:IdentityPolicy
    properties:
      identity: ${exampleDomainIdentity.arn}
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - SES:SendEmail
              - SES:SendRawEmail
            resources:
              - ${exampleDomainIdentity.arn}
            principals:
              - identifiers:
                  - '*'
                type: AWS
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES Identity Policies using the identity and policy name, separated by a pipe character (`|`). For example:

```sh
 $ pulumi import aws:ses/identityPolicy:IdentityPolicy example 'example.com|example'
```
 