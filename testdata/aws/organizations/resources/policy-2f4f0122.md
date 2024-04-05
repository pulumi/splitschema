Provides a resource to manage an [AWS Organizations policy](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["*"],
        resources: ["*"],
    }],
});
const examplePolicy = new aws.organizations.Policy("examplePolicy", {content: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json)});
```
```python
import pulumi
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["*"],
    resources=["*"],
)])
example_policy = aws.organizations.Policy("examplePolicy", content=example_policy_document.json)
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
                Effect = "Allow",
                Actions = new[]
                {
                    "*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var examplePolicy = new Aws.Organizations.Policy("examplePolicy", new()
    {
        Content = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"*",
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
		_, err = organizations.NewPolicy(ctx, "examplePolicy", &organizations.PolicyArgs{
			Content: *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.organizations.Policy;
import com.pulumi.aws.organizations.PolicyArgs;
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
                .effect("Allow")
                .actions("*")
                .resources("*")
                .build())
            .build());

        var examplePolicy = new Policy("examplePolicy", PolicyArgs.builder()        
            .content(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  examplePolicy:
    type: aws:organizations:Policy
    properties:
      content: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - '*'
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_organizations_policy` using the policy ID. For example:

```sh
 $ pulumi import aws:organizations/policy:Policy example p-12345678
```
 