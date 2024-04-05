Provides an Elastic Container Registry Repository Policy.

Note that currently only one policy may be applied to a repository.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ecr.Repository("foo", {});
const foopolicyPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "new policy",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["123456789012"],
        }],
        actions: [
            "ecr:GetDownloadUrlForLayer",
            "ecr:BatchGetImage",
            "ecr:BatchCheckLayerAvailability",
            "ecr:PutImage",
            "ecr:InitiateLayerUpload",
            "ecr:UploadLayerPart",
            "ecr:CompleteLayerUpload",
            "ecr:DescribeRepositories",
            "ecr:GetRepositoryPolicy",
            "ecr:ListImages",
            "ecr:DeleteRepository",
            "ecr:BatchDeleteImage",
            "ecr:SetRepositoryPolicy",
            "ecr:DeleteRepositoryPolicy",
        ],
    }],
});
const foopolicyRepositoryPolicy = new aws.ecr.RepositoryPolicy("foopolicyRepositoryPolicy", {
    repository: foo.name,
    policy: foopolicyPolicyDocument.then(foopolicyPolicyDocument => foopolicyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo")
foopolicy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="new policy",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["123456789012"],
    )],
    actions=[
        "ecr:GetDownloadUrlForLayer",
        "ecr:BatchGetImage",
        "ecr:BatchCheckLayerAvailability",
        "ecr:PutImage",
        "ecr:InitiateLayerUpload",
        "ecr:UploadLayerPart",
        "ecr:CompleteLayerUpload",
        "ecr:DescribeRepositories",
        "ecr:GetRepositoryPolicy",
        "ecr:ListImages",
        "ecr:DeleteRepository",
        "ecr:BatchDeleteImage",
        "ecr:SetRepositoryPolicy",
        "ecr:DeleteRepositoryPolicy",
    ],
)])
foopolicy_repository_policy = aws.ecr.RepositoryPolicy("foopolicyRepositoryPolicy",
    repository=foo.name,
    policy=foopolicy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ecr.Repository("foo");

    var foopolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "new policy",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "123456789012",
                        },
                    },
                },
                Actions = new[]
                {
                    "ecr:GetDownloadUrlForLayer",
                    "ecr:BatchGetImage",
                    "ecr:BatchCheckLayerAvailability",
                    "ecr:PutImage",
                    "ecr:InitiateLayerUpload",
                    "ecr:UploadLayerPart",
                    "ecr:CompleteLayerUpload",
                    "ecr:DescribeRepositories",
                    "ecr:GetRepositoryPolicy",
                    "ecr:ListImages",
                    "ecr:DeleteRepository",
                    "ecr:BatchDeleteImage",
                    "ecr:SetRepositoryPolicy",
                    "ecr:DeleteRepositoryPolicy",
                },
            },
        },
    });

    var foopolicyRepositoryPolicy = new Aws.Ecr.RepositoryPolicy("foopolicyRepositoryPolicy", new()
    {
        Repository = foo.Name,
        Policy = foopolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ecr.NewRepository(ctx, "foo", nil)
		if err != nil {
			return err
		}
		foopolicyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("new policy"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								"123456789012",
							},
						},
					},
					Actions: []string{
						"ecr:GetDownloadUrlForLayer",
						"ecr:BatchGetImage",
						"ecr:BatchCheckLayerAvailability",
						"ecr:PutImage",
						"ecr:InitiateLayerUpload",
						"ecr:UploadLayerPart",
						"ecr:CompleteLayerUpload",
						"ecr:DescribeRepositories",
						"ecr:GetRepositoryPolicy",
						"ecr:ListImages",
						"ecr:DeleteRepository",
						"ecr:BatchDeleteImage",
						"ecr:SetRepositoryPolicy",
						"ecr:DeleteRepositoryPolicy",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewRepositoryPolicy(ctx, "foopolicyRepositoryPolicy", &ecr.RepositoryPolicyArgs{
			Repository: foo.Name,
			Policy:     *pulumi.String(foopolicyPolicyDocument.Json),
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
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.ecr.RepositoryPolicy;
import com.pulumi.aws.ecr.RepositoryPolicyArgs;
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
        var foo = new Repository("foo");

        final var foopolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("new policy")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("123456789012")
                    .build())
                .actions(                
                    "ecr:GetDownloadUrlForLayer",
                    "ecr:BatchGetImage",
                    "ecr:BatchCheckLayerAvailability",
                    "ecr:PutImage",
                    "ecr:InitiateLayerUpload",
                    "ecr:UploadLayerPart",
                    "ecr:CompleteLayerUpload",
                    "ecr:DescribeRepositories",
                    "ecr:GetRepositoryPolicy",
                    "ecr:ListImages",
                    "ecr:DeleteRepository",
                    "ecr:BatchDeleteImage",
                    "ecr:SetRepositoryPolicy",
                    "ecr:DeleteRepositoryPolicy")
                .build())
            .build());

        var foopolicyRepositoryPolicy = new RepositoryPolicy("foopolicyRepositoryPolicy", RepositoryPolicyArgs.builder()        
            .repository(foo.name())
            .policy(foopolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ecr:Repository
  foopolicyRepositoryPolicy:
    type: aws:ecr:RepositoryPolicy
    properties:
      repository: ${foo.name}
      policy: ${foopolicyPolicyDocument.json}
variables:
  foopolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: new policy
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '123456789012'
            actions:
              - ecr:GetDownloadUrlForLayer
              - ecr:BatchGetImage
              - ecr:BatchCheckLayerAvailability
              - ecr:PutImage
              - ecr:InitiateLayerUpload
              - ecr:UploadLayerPart
              - ecr:CompleteLayerUpload
              - ecr:DescribeRepositories
              - ecr:GetRepositoryPolicy
              - ecr:ListImages
              - ecr:DeleteRepository
              - ecr:BatchDeleteImage
              - ecr:SetRepositoryPolicy
              - ecr:DeleteRepositoryPolicy
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECR Repository Policy using the repository name. For example:

```sh
 $ pulumi import aws:ecr/repositoryPolicy:RepositoryPolicy example example
```
 