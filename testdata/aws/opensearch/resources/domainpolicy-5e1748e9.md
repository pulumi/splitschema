Allows setting policy to an OpenSearch domain while referencing domain attributes (e.g., ARN).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.opensearch.Domain("example", {engineVersion: "OpenSearch_1.1"});
const mainPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: ["es:*"],
        resources: [pulumi.interpolate`${example.arn}/*`],
        conditions: [{
            test: "IpAddress",
            variable: "aws:SourceIp",
            values: ["127.0.0.1/32"],
        }],
    }],
});
const mainDomainPolicy = new aws.opensearch.DomainPolicy("mainDomainPolicy", {
    domainName: example.domainName,
    accessPolicies: mainPolicyDocument.apply(mainPolicyDocument => mainPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opensearch.Domain("example", engine_version="OpenSearch_1.1")
main_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=["es:*"],
    resources=[example.arn.apply(lambda arn: f"{arn}/*")],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="IpAddress",
        variable="aws:SourceIp",
        values=["127.0.0.1/32"],
    )],
)])
main_domain_policy = aws.opensearch.DomainPolicy("mainDomainPolicy",
    domain_name=example.domain_name,
    access_policies=main_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.OpenSearch.Domain("example", new()
    {
        EngineVersion = "OpenSearch_1.1",
    });

    var mainPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                    "es:*",
                },
                Resources = new[]
                {
                    $"{example.Arn}/*",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "IpAddress",
                        Variable = "aws:SourceIp",
                        Values = new[]
                        {
                            "127.0.0.1/32",
                        },
                    },
                },
            },
        },
    });

    var mainDomainPolicy = new Aws.OpenSearch.DomainPolicy("mainDomainPolicy", new()
    {
        DomainName = example.DomainName,
        AccessPolicies = mainPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := opensearch.NewDomain(ctx, "example", &opensearch.DomainArgs{
			EngineVersion: pulumi.String("OpenSearch_1.1"),
		})
		if err != nil {
			return err
		}
		mainPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
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
						pulumi.String("es:*"),
					},
					Resources: pulumi.StringArray{
						example.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("IpAddress"),
							Variable: pulumi.String("aws:SourceIp"),
							Values: pulumi.StringArray{
								pulumi.String("127.0.0.1/32"),
							},
						},
					},
				},
			},
		}, nil)
		_, err = opensearch.NewDomainPolicy(ctx, "mainDomainPolicy", &opensearch.DomainPolicyArgs{
			DomainName: example.DomainName,
			AccessPolicies: mainPolicyDocument.ApplyT(func(mainPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &mainPolicyDocument.Json, nil
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
import com.pulumi.aws.opensearch.Domain;
import com.pulumi.aws.opensearch.DomainArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.opensearch.DomainPolicy;
import com.pulumi.aws.opensearch.DomainPolicyArgs;
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
        var example = new Domain("example", DomainArgs.builder()        
            .engineVersion("OpenSearch_1.1")
            .build());

        final var mainPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions("es:*")
                .resources(example.arn().applyValue(arn -> String.format("%s/*", arn)))
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("IpAddress")
                    .variable("aws:SourceIp")
                    .values("127.0.0.1/32")
                    .build())
                .build())
            .build());

        var mainDomainPolicy = new DomainPolicy("mainDomainPolicy", DomainPolicyArgs.builder()        
            .domainName(example.domainName())
            .accessPolicies(mainPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(mainPolicyDocument -> mainPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:Domain
    properties:
      engineVersion: OpenSearch_1.1
  mainDomainPolicy:
    type: aws:opensearch:DomainPolicy
    properties:
      domainName: ${example.domainName}
      accessPolicies: ${mainPolicyDocument.json}
variables:
  mainPolicyDocument:
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
              - es:*
            resources:
              - ${example.arn}/*
            conditions:
              - test: IpAddress
                variable: aws:SourceIp
                values:
                  - 127.0.0.1/32
```
{{% /example %}}
{{% /examples %}}