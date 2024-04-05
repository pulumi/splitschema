Provides an CloudSearch domain service access policy resource.

The provider waits for the domain service access policy to become `Active` when applying a configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomain = new aws.cloudsearch.Domain("exampleDomain", {});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "search_only",
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: [
            "cloudsearch:search",
            "cloudsearch:document",
        ],
        conditions: [{
            test: "IpAddress",
            variable: "aws:SourceIp",
            values: ["192.0.2.0/32"],
        }],
    }],
});
const exampleDomainServiceAccessPolicy = new aws.cloudsearch.DomainServiceAccessPolicy("exampleDomainServiceAccessPolicy", {
    domainName: exampleDomain.id,
    accessPolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_domain = aws.cloudsearch.Domain("exampleDomain")
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="search_only",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=[
        "cloudsearch:search",
        "cloudsearch:document",
    ],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="IpAddress",
        variable="aws:SourceIp",
        values=["192.0.2.0/32"],
    )],
)])
example_domain_service_access_policy = aws.cloudsearch.DomainServiceAccessPolicy("exampleDomainServiceAccessPolicy",
    domain_name=example_domain.id,
    access_policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomain = new Aws.CloudSearch.Domain("exampleDomain");

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "search_only",
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
                    "cloudsearch:search",
                    "cloudsearch:document",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "IpAddress",
                        Variable = "aws:SourceIp",
                        Values = new[]
                        {
                            "192.0.2.0/32",
                        },
                    },
                },
            },
        },
    });

    var exampleDomainServiceAccessPolicy = new Aws.CloudSearch.DomainServiceAccessPolicy("exampleDomainServiceAccessPolicy", new()
    {
        DomainName = exampleDomain.Id,
        AccessPolicy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudsearch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomain, err := cloudsearch.NewDomain(ctx, "exampleDomain", nil)
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("search_only"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "*",
							Identifiers: []string{
								"*",
							},
						},
					},
					Actions: []string{
						"cloudsearch:search",
						"cloudsearch:document",
					},
					Conditions: []iam.GetPolicyDocumentStatementCondition{
						{
							Test:     "IpAddress",
							Variable: "aws:SourceIp",
							Values: []string{
								"192.0.2.0/32",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudsearch.NewDomainServiceAccessPolicy(ctx, "exampleDomainServiceAccessPolicy", &cloudsearch.DomainServiceAccessPolicyArgs{
			DomainName:   exampleDomain.ID(),
			AccessPolicy: *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.cloudsearch.Domain;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.cloudsearch.DomainServiceAccessPolicy;
import com.pulumi.aws.cloudsearch.DomainServiceAccessPolicyArgs;
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
        var exampleDomain = new Domain("exampleDomain");

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("search_only")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions(                
                    "cloudsearch:search",
                    "cloudsearch:document")
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("IpAddress")
                    .variable("aws:SourceIp")
                    .values("192.0.2.0/32")
                    .build())
                .build())
            .build());

        var exampleDomainServiceAccessPolicy = new DomainServiceAccessPolicy("exampleDomainServiceAccessPolicy", DomainServiceAccessPolicyArgs.builder()        
            .domainName(exampleDomain.id())
            .accessPolicy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleDomain:
    type: aws:cloudsearch:Domain
  exampleDomainServiceAccessPolicy:
    type: aws:cloudsearch:DomainServiceAccessPolicy
    properties:
      domainName: ${exampleDomain.id}
      accessPolicy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: search_only
            effect: Allow
            principals:
              - type: '*'
                identifiers:
                  - '*'
            actions:
              - cloudsearch:search
              - cloudsearch:document
            conditions:
              - test: IpAddress
                variable: aws:SourceIp
                values:
                  - 192.0.2.0/32
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudSearch domain service access policies using the domain name. For example:

```sh
 $ pulumi import aws:cloudsearch/domainServiceAccessPolicy:DomainServiceAccessPolicy example example-domain
```
 