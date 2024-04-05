Attaches a resource based policy to a private CA.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [
        {
            sid: "1",
            effect: "Allow",
            principals: [{
                type: "AWS",
                identifiers: [data.aws_caller_identity.current.account_id],
            }],
            actions: [
                "acm-pca:DescribeCertificateAuthority",
                "acm-pca:GetCertificate",
                "acm-pca:GetCertificateAuthorityCertificate",
                "acm-pca:ListPermissions",
                "acm-pca:ListTags",
            ],
            resources: [aws_acmpca_certificate_authority.example.arn],
        },
        {
            sid: "2",
            effect: Allow,
            principals: [{
                type: "AWS",
                identifiers: [data.aws_caller_identity.current.account_id],
            }],
            actions: ["acm-pca:IssueCertificate"],
            resources: [aws_acmpca_certificate_authority.example.arn],
            conditions: [{
                test: "StringEquals",
                variable: "acm-pca:TemplateArn",
                values: ["arn:aws:acm-pca:::template/EndEntityCertificate/V1"],
            }],
        },
    ],
});
const examplePolicy = new aws.acmpca.Policy("examplePolicy", {
    resourceArn: aws_acmpca_certificate_authority.example.arn,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="1",
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[data["aws_caller_identity"]["current"]["account_id"]],
        )],
        actions=[
            "acm-pca:DescribeCertificateAuthority",
            "acm-pca:GetCertificate",
            "acm-pca:GetCertificateAuthorityCertificate",
            "acm-pca:ListPermissions",
            "acm-pca:ListTags",
        ],
        resources=[aws_acmpca_certificate_authority["example"]["arn"]],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="2",
        effect=allow,
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[data["aws_caller_identity"]["current"]["account_id"]],
        )],
        actions=["acm-pca:IssueCertificate"],
        resources=[aws_acmpca_certificate_authority["example"]["arn"]],
        conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
            test="StringEquals",
            variable="acm-pca:TemplateArn",
            values=["arn:aws:acm-pca:::template/EndEntityCertificate/V1"],
        )],
    ),
])
example_policy = aws.acmpca.Policy("examplePolicy",
    resource_arn=aws_acmpca_certificate_authority["example"]["arn"],
    policy=example_policy_document.json)
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
                Sid = "1",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            data.Aws_caller_identity.Current.Account_id,
                        },
                    },
                },
                Actions = new[]
                {
                    "acm-pca:DescribeCertificateAuthority",
                    "acm-pca:GetCertificate",
                    "acm-pca:GetCertificateAuthorityCertificate",
                    "acm-pca:ListPermissions",
                    "acm-pca:ListTags",
                },
                Resources = new[]
                {
                    aws_acmpca_certificate_authority.Example.Arn,
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "2",
                Effect = Allow,
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            data.Aws_caller_identity.Current.Account_id,
                        },
                    },
                },
                Actions = new[]
                {
                    "acm-pca:IssueCertificate",
                },
                Resources = new[]
                {
                    aws_acmpca_certificate_authority.Example.Arn,
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "StringEquals",
                        Variable = "acm-pca:TemplateArn",
                        Values = new[]
                        {
                            "arn:aws:acm-pca:::template/EndEntityCertificate/V1",
                        },
                    },
                },
            },
        },
    });

    var examplePolicy = new Aws.Acmpca.Policy("examplePolicy", new()
    {
        ResourceArn = aws_acmpca_certificate_authority.Example.Arn,
        PolicyDetails = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acmpca"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Sid: pulumi.StringRef("1"),
Effect: pulumi.StringRef("Allow"),
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
data.Aws_caller_identity.Current.Account_id,
},
},
},
Actions: []string{
"acm-pca:DescribeCertificateAuthority",
"acm-pca:GetCertificate",
"acm-pca:GetCertificateAuthorityCertificate",
"acm-pca:ListPermissions",
"acm-pca:ListTags",
},
Resources: interface{}{
aws_acmpca_certificate_authority.Example.Arn,
},
},
{
Sid: pulumi.StringRef("2"),
Effect: pulumi.StringRef(Allow),
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
data.Aws_caller_identity.Current.Account_id,
},
},
},
Actions: []string{
"acm-pca:IssueCertificate",
},
Resources: interface{}{
aws_acmpca_certificate_authority.Example.Arn,
},
Conditions: []iam.GetPolicyDocumentStatementCondition{
{
Test: "StringEquals",
Variable: "acm-pca:TemplateArn",
Values: []string{
"arn:aws:acm-pca:::template/EndEntityCertificate/V1",
},
},
},
},
},
}, nil);
if err != nil {
return err
}
_, err = acmpca.NewPolicy(ctx, "examplePolicy", &acmpca.PolicyArgs{
ResourceArn: pulumi.Any(aws_acmpca_certificate_authority.Example.Arn),
Policy: *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.acmpca.Policy;
import com.pulumi.aws.acmpca.PolicyArgs;
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
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("1")
                    .effect("Allow")
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(data.aws_caller_identity().current().account_id())
                        .build())
                    .actions(                    
                        "acm-pca:DescribeCertificateAuthority",
                        "acm-pca:GetCertificate",
                        "acm-pca:GetCertificateAuthorityCertificate",
                        "acm-pca:ListPermissions",
                        "acm-pca:ListTags")
                    .resources(aws_acmpca_certificate_authority.example().arn())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("2")
                    .effect(Allow)
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(data.aws_caller_identity().current().account_id())
                        .build())
                    .actions("acm-pca:IssueCertificate")
                    .resources(aws_acmpca_certificate_authority.example().arn())
                    .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                        .test("StringEquals")
                        .variable("acm-pca:TemplateArn")
                        .values("arn:aws:acm-pca:::template/EndEntityCertificate/V1")
                        .build())
                    .build())
            .build());

        var examplePolicy = new Policy("examplePolicy", PolicyArgs.builder()        
            .resourceArn(aws_acmpca_certificate_authority.example().arn())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  examplePolicy:
    type: aws:acmpca:Policy
    properties:
      resourceArn: ${aws_acmpca_certificate_authority.example.arn}
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: '1'
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${data.aws_caller_identity.current.account_id}
            actions:
              - acm-pca:DescribeCertificateAuthority
              - acm-pca:GetCertificate
              - acm-pca:GetCertificateAuthorityCertificate
              - acm-pca:ListPermissions
              - acm-pca:ListTags
            resources:
              - ${aws_acmpca_certificate_authority.example.arn}
          - sid: '2'
            effect: ${Allow}
            principals:
              - type: AWS
                identifiers:
                  - ${data.aws_caller_identity.current.account_id}
            actions:
              - acm-pca:IssueCertificate
            resources:
              - ${aws_acmpca_certificate_authority.example.arn}
            conditions:
              - test: StringEquals
                variable: acm-pca:TemplateArn
                values:
                  - arn:aws:acm-pca:::template/EndEntityCertificate/V1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_acmpca_policy` using the `resource_arn` value. For example:

```sh
 $ pulumi import aws:acmpca/policy:Policy example arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
```
 