Resource for managing an AWS EventBridge Schemas Registry Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "example",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["109876543210"],
        }],
        actions: ["schemas:*"],
        resources: [
            "arn:aws:schemas:us-east-1:012345678901:registry/example",
            "arn:aws:schemas:us-east-1:012345678901:schema/example*",
        ],
    }],
});
const exampleRegistryPolicy = new aws.schemas.RegistryPolicy("exampleRegistryPolicy", {
    registryName: "example",
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="example",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["109876543210"],
    )],
    actions=["schemas:*"],
    resources=[
        "arn:aws:schemas:us-east-1:012345678901:registry/example",
        "arn:aws:schemas:us-east-1:012345678901:schema/example*",
    ],
)])
example_registry_policy = aws.schemas.RegistryPolicy("exampleRegistryPolicy",
    registry_name="example",
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
                Sid = "example",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "109876543210",
                        },
                    },
                },
                Actions = new[]
                {
                    "schemas:*",
                },
                Resources = new[]
                {
                    "arn:aws:schemas:us-east-1:012345678901:registry/example",
                    "arn:aws:schemas:us-east-1:012345678901:schema/example*",
                },
            },
        },
    });

    var exampleRegistryPolicy = new Aws.Schemas.RegistryPolicy("exampleRegistryPolicy", new()
    {
        RegistryName = "example",
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/schemas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("example"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								"109876543210",
							},
						},
					},
					Actions: []string{
						"schemas:*",
					},
					Resources: []string{
						"arn:aws:schemas:us-east-1:012345678901:registry/example",
						"arn:aws:schemas:us-east-1:012345678901:schema/example*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = schemas.NewRegistryPolicy(ctx, "exampleRegistryPolicy", &schemas.RegistryPolicyArgs{
			RegistryName: pulumi.String("example"),
			Policy:       *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.schemas.RegistryPolicy;
import com.pulumi.aws.schemas.RegistryPolicyArgs;
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
                .sid("example")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("109876543210")
                    .build())
                .actions("schemas:*")
                .resources(                
                    "arn:aws:schemas:us-east-1:012345678901:registry/example",
                    "arn:aws:schemas:us-east-1:012345678901:schema/example*")
                .build())
            .build());

        var exampleRegistryPolicy = new RegistryPolicy("exampleRegistryPolicy", RegistryPolicyArgs.builder()        
            .registryName("example")
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleRegistryPolicy:
    type: aws:schemas:RegistryPolicy
    properties:
      registryName: example
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: example
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '109876543210'
            actions:
              - schemas:*
            resources:
              - arn:aws:schemas:us-east-1:012345678901:registry/example
              - arn:aws:schemas:us-east-1:012345678901:schema/example*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge Schema Registry Policy using the `registry_name`. For example:

```sh
 $ pulumi import aws:schemas/registryPolicy:RegistryPolicy example example
```
 