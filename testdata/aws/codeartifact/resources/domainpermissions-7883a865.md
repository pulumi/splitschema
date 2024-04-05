Provides a CodeArtifact Domains Permissions Policy Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {description: "domain key"});
const exampleDomain = new aws.codeartifact.Domain("exampleDomain", {
    domain: "example",
    encryptionKey: exampleKey.arn,
});
const testPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: ["codeartifact:CreateRepository"],
        resources: [exampleDomain.arn],
    }],
});
const testDomainPermissions = new aws.codeartifact.DomainPermissions("testDomainPermissions", {
    domain: exampleDomain.domain,
    policyDocument: testPolicyDocument.apply(testPolicyDocument => testPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey", description="domain key")
example_domain = aws.codeartifact.Domain("exampleDomain",
    domain="example",
    encryption_key=example_key.arn)
test_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=["codeartifact:CreateRepository"],
    resources=[example_domain.arn],
)])
test_domain_permissions = aws.codeartifact.DomainPermissions("testDomainPermissions",
    domain=example_domain.domain,
    policy_document=test_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "domain key",
    });

    var exampleDomain = new Aws.CodeArtifact.Domain("exampleDomain", new()
    {
        DomainName = "example",
        EncryptionKey = exampleKey.Arn,
    });

    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "*",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "codeartifact:CreateRepository",
                },
                Resources = new[]
                {
                    exampleDomain.Arn,
                },
            },
        },
    });

    var testDomainPermissions = new Aws.CodeArtifact.DomainPermissions("testDomainPermissions", new()
    {
        Domain = exampleDomain.DomainName,
        PolicyDocument = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeartifact"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description: pulumi.String("domain key"),
		})
		if err != nil {
			return err
		}
		exampleDomain, err := codeartifact.NewDomain(ctx, "exampleDomain", &codeartifact.DomainArgs{
			Domain:        pulumi.String("example"),
			EncryptionKey: exampleKey.Arn,
		})
		if err != nil {
			return err
		}
		testPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("*"),
							Identifiers: pulumi.StringArray{
								pulumi.String("*"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("codeartifact:CreateRepository"),
					},
					Resources: pulumi.StringArray{
						exampleDomain.Arn,
					},
				},
			},
		}, nil)
		_, err = codeartifact.NewDomainPermissions(ctx, "testDomainPermissions", &codeartifact.DomainPermissionsArgs{
			Domain: exampleDomain.Domain,
			PolicyDocument: testPolicyDocument.ApplyT(func(testPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &testPolicyDocument.Json, nil
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.codeartifact.Domain;
import com.pulumi.aws.codeartifact.DomainArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.codeartifact.DomainPermissions;
import com.pulumi.aws.codeartifact.DomainPermissionsArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("domain key")
            .build());

        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .domain("example")
            .encryptionKey(exampleKey.arn())
            .build());

        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions("codeartifact:CreateRepository")
                .resources(exampleDomain.arn())
                .build())
            .build());

        var testDomainPermissions = new DomainPermissions("testDomainPermissions", DomainPermissionsArgs.builder()        
            .domain(exampleDomain.domain())
            .policyDocument(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(testPolicyDocument -> testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: domain key
  exampleDomain:
    type: aws:codeartifact:Domain
    properties:
      domain: example
      encryptionKey: ${exampleKey.arn}
  testDomainPermissions:
    type: aws:codeartifact:DomainPermissions
    properties:
      domain: ${exampleDomain.domain}
      policyDocument: ${testPolicyDocument.json}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: '*'
                identifiers:
                  - '*'
            actions:
              - codeartifact:CreateRepository
            resources:
              - ${exampleDomain.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeArtifact Domain Permissions Policies using the CodeArtifact Domain ARN. For example:

```sh
 $ pulumi import aws:codeartifact/domainPermissions:DomainPermissions example arn:aws:codeartifact:us-west-2:012345678912:domain/tf-acc-test-1928056699409417367
```
 