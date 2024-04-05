Provides a CodeArtifact Repostory Permissions Policy Resource.

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
const exampleRepository = new aws.codeartifact.Repository("exampleRepository", {
    repository: "example",
    domain: exampleDomain.domain,
});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: ["codeartifact:ReadFromRepository"],
        resources: [exampleRepository.arn],
    }],
});
const exampleRepositoryPermissionsPolicy = new aws.codeartifact.RepositoryPermissionsPolicy("exampleRepositoryPermissionsPolicy", {
    repository: exampleRepository.repository,
    domain: exampleDomain.domain,
    policyDocument: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey", description="domain key")
example_domain = aws.codeartifact.Domain("exampleDomain",
    domain="example",
    encryption_key=example_key.arn)
example_repository = aws.codeartifact.Repository("exampleRepository",
    repository="example",
    domain=example_domain.domain)
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=["codeartifact:ReadFromRepository"],
    resources=[example_repository.arn],
)])
example_repository_permissions_policy = aws.codeartifact.RepositoryPermissionsPolicy("exampleRepositoryPermissionsPolicy",
    repository=example_repository.repository,
    domain=example_domain.domain,
    policy_document=example_policy_document.json)
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

    var exampleRepository = new Aws.CodeArtifact.Repository("exampleRepository", new()
    {
        RepositoryName = "example",
        Domain = exampleDomain.DomainName,
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                    "codeartifact:ReadFromRepository",
                },
                Resources = new[]
                {
                    exampleRepository.Arn,
                },
            },
        },
    });

    var exampleRepositoryPermissionsPolicy = new Aws.CodeArtifact.RepositoryPermissionsPolicy("exampleRepositoryPermissionsPolicy", new()
    {
        Repository = exampleRepository.RepositoryName,
        Domain = exampleDomain.DomainName,
        PolicyDocument = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
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
		exampleRepository, err := codeartifact.NewRepository(ctx, "exampleRepository", &codeartifact.RepositoryArgs{
			Repository: pulumi.String("example"),
			Domain:     exampleDomain.Domain,
		})
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
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
						pulumi.String("codeartifact:ReadFromRepository"),
					},
					Resources: pulumi.StringArray{
						exampleRepository.Arn,
					},
				},
			},
		}, nil)
		_, err = codeartifact.NewRepositoryPermissionsPolicy(ctx, "exampleRepositoryPermissionsPolicy", &codeartifact.RepositoryPermissionsPolicyArgs{
			Repository: exampleRepository.Repository,
			Domain:     exampleDomain.Domain,
			PolicyDocument: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.codeartifact.Domain;
import com.pulumi.aws.codeartifact.DomainArgs;
import com.pulumi.aws.codeartifact.Repository;
import com.pulumi.aws.codeartifact.RepositoryArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.codeartifact.RepositoryPermissionsPolicy;
import com.pulumi.aws.codeartifact.RepositoryPermissionsPolicyArgs;
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

        var exampleRepository = new Repository("exampleRepository", RepositoryArgs.builder()        
            .repository("example")
            .domain(exampleDomain.domain())
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions("codeartifact:ReadFromRepository")
                .resources(exampleRepository.arn())
                .build())
            .build());

        var exampleRepositoryPermissionsPolicy = new RepositoryPermissionsPolicy("exampleRepositoryPermissionsPolicy", RepositoryPermissionsPolicyArgs.builder()        
            .repository(exampleRepository.repository())
            .domain(exampleDomain.domain())
            .policyDocument(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
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
  exampleRepository:
    type: aws:codeartifact:Repository
    properties:
      repository: example
      domain: ${exampleDomain.domain}
  exampleRepositoryPermissionsPolicy:
    type: aws:codeartifact:RepositoryPermissionsPolicy
    properties:
      repository: ${exampleRepository.repository}
      domain: ${exampleDomain.domain}
      policyDocument: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
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
              - codeartifact:ReadFromRepository
            resources:
              - ${exampleRepository.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeArtifact Repository Permissions Policies using the CodeArtifact Repository ARN. For example:

```sh
 $ pulumi import aws:codeartifact/repositoryPermissionsPolicy:RepositoryPermissionsPolicy example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
```
 