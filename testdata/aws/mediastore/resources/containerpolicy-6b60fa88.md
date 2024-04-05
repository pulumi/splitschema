Provides a MediaStore Container Policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const exampleContainer = new aws.mediastore.Container("exampleContainer", {});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        sid: "MediaStoreFullAccess",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: [currentCallerIdentity.then(currentCallerIdentity => `arn:aws:iam::${currentCallerIdentity.accountId}:root`)],
        }],
        actions: ["mediastore:*"],
        resources: [pulumi.all([currentRegion, currentCallerIdentity, exampleContainer.name]).apply(([currentRegion, currentCallerIdentity, name]) => `arn:aws:mediastore:${currentRegion.name}:${currentCallerIdentity.accountId}:container/${name}/*`)],
        conditions: [{
            test: "Bool",
            variable: "aws:SecureTransport",
            values: ["true"],
        }],
    }],
});
const exampleContainerPolicy = new aws.mediastore.ContainerPolicy("exampleContainerPolicy", {
    containerName: exampleContainer.name,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
example_container = aws.mediastore.Container("exampleContainer")
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="MediaStoreFullAccess",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[f"arn:aws:iam::{current_caller_identity.account_id}:root"],
    )],
    actions=["mediastore:*"],
    resources=[example_container.name.apply(lambda name: f"arn:aws:mediastore:{current_region.name}:{current_caller_identity.account_id}:container/{name}/*")],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="Bool",
        variable="aws:SecureTransport",
        values=["true"],
    )],
)])
example_container_policy = aws.mediastore.ContainerPolicy("exampleContainerPolicy",
    container_name=example_container.name,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var exampleContainer = new Aws.MediaStore.Container("exampleContainer");

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "MediaStoreFullAccess",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            $"arn:aws:iam::{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                    },
                },
                Actions = new[]
                {
                    "mediastore:*",
                },
                Resources = new[]
                {
                    $"arn:aws:mediastore:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:container/{exampleContainer.Name}/*",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "Bool",
                        Variable = "aws:SecureTransport",
                        Values = new[]
                        {
                            "true",
                        },
                    },
                },
            },
        },
    });

    var exampleContainerPolicy = new Aws.MediaStore.ContainerPolicy("exampleContainerPolicy", new()
    {
        ContainerName = exampleContainer.Name,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mediastore"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleContainer, err := mediastore.NewContainer(ctx, "exampleContainer", nil)
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Sid:    pulumi.String("MediaStoreFullAccess"),
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("AWS"),
							Identifiers: pulumi.StringArray{
								pulumi.String(fmt.Sprintf("arn:aws:iam::%v:root", currentCallerIdentity.AccountId)),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("mediastore:*"),
					},
					Resources: pulumi.StringArray{
						exampleContainer.Name.ApplyT(func(name string) (string, error) {
							return fmt.Sprintf("arn:aws:mediastore:%v:%v:container/%v/*", currentRegion.Name, currentCallerIdentity.AccountId, name), nil
						}).(pulumi.StringOutput),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("Bool"),
							Variable: pulumi.String("aws:SecureTransport"),
							Values: pulumi.StringArray{
								pulumi.String("true"),
							},
						},
					},
				},
			},
		}, nil)
		_, err = mediastore.NewContainerPolicy(ctx, "exampleContainerPolicy", &mediastore.ContainerPolicyArgs{
			ContainerName: exampleContainer.Name,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.mediastore.Container;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.mediastore.ContainerPolicy;
import com.pulumi.aws.mediastore.ContainerPolicyArgs;
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
        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        var exampleContainer = new Container("exampleContainer");

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("MediaStoreFullAccess")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(String.format("arn:aws:iam::%s:root", currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                    .build())
                .actions("mediastore:*")
                .resources(exampleContainer.name().applyValue(name -> String.format("arn:aws:mediastore:%s:%s:container/%s/*", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),name)))
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("Bool")
                    .variable("aws:SecureTransport")
                    .values("true")
                    .build())
                .build())
            .build());

        var exampleContainerPolicy = new ContainerPolicy("exampleContainerPolicy", ContainerPolicyArgs.builder()        
            .containerName(exampleContainer.name())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  exampleContainer:
    type: aws:mediastore:Container
  exampleContainerPolicy:
    type: aws:mediastore:ContainerPolicy
    properties:
      containerName: ${exampleContainer.name}
      policy: ${examplePolicyDocument.json}
variables:
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: MediaStoreFullAccess
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - arn:aws:iam::${currentCallerIdentity.accountId}:root
            actions:
              - mediastore:*
            resources:
              - arn:aws:mediastore:${currentRegion.name}:${currentCallerIdentity.accountId}:container/${exampleContainer.name}/*
            conditions:
              - test: Bool
                variable: aws:SecureTransport
                values:
                  - 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MediaStore Container Policy using the MediaStore Container Name. For example:

```sh
 $ pulumi import aws:mediastore/containerPolicy:ContainerPolicy example example
```
 