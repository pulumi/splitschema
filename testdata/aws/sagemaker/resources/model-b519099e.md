Provides a SageMaker model resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["sagemaker.amazonaws.com"],
        }],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const test = aws.sagemaker.getPrebuiltEcrImage({
    repositoryName: "kmeans",
});
const exampleModel = new aws.sagemaker.Model("exampleModel", {
    executionRoleArn: exampleRole.arn,
    primaryContainer: {
        image: test.then(test => test.registryPath),
    },
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["sagemaker.amazonaws.com"],
    )],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
test = aws.sagemaker.get_prebuilt_ecr_image(repository_name="kmeans")
example_model = aws.sagemaker.Model("exampleModel",
    execution_role_arn=example_role.arn,
    primary_container=aws.sagemaker.ModelPrimaryContainerArgs(
        image=test.registry_path,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "sagemaker.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var test = Aws.Sagemaker.GetPrebuiltEcrImage.Invoke(new()
    {
        RepositoryName = "kmeans",
    });

    var exampleModel = new Aws.Sagemaker.Model("exampleModel", new()
    {
        ExecutionRoleArn = exampleRole.Arn,
        PrimaryContainer = new Aws.Sagemaker.Inputs.ModelPrimaryContainerArgs
        {
            Image = test.Apply(getPrebuiltEcrImageResult => getPrebuiltEcrImageResult.RegistryPath),
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"sagemaker.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		test, err := sagemaker.GetPrebuiltEcrImage(ctx, &sagemaker.GetPrebuiltEcrImageArgs{
			RepositoryName: "kmeans",
		}, nil)
		if err != nil {
			return err
		}
		_, err = sagemaker.NewModel(ctx, "exampleModel", &sagemaker.ModelArgs{
			ExecutionRoleArn: exampleRole.Arn,
			PrimaryContainer: &sagemaker.ModelPrimaryContainerArgs{
				Image: *pulumi.String(test.RegistryPath),
			},
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.sagemaker.SagemakerFunctions;
import com.pulumi.aws.sagemaker.inputs.GetPrebuiltEcrImageArgs;
import com.pulumi.aws.sagemaker.Model;
import com.pulumi.aws.sagemaker.ModelArgs;
import com.pulumi.aws.sagemaker.inputs.ModelPrimaryContainerArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("sagemaker.amazonaws.com")
                    .build())
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var test = SagemakerFunctions.getPrebuiltEcrImage(GetPrebuiltEcrImageArgs.builder()
            .repositoryName("kmeans")
            .build());

        var exampleModel = new Model("exampleModel", ModelArgs.builder()        
            .executionRoleArn(exampleRole.arn())
            .primaryContainer(ModelPrimaryContainerArgs.builder()
                .image(test.applyValue(getPrebuiltEcrImageResult -> getPrebuiltEcrImageResult.registryPath()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleModel:
    type: aws:sagemaker:Model
    properties:
      executionRoleArn: ${exampleRole.arn}
      primaryContainer:
        image: ${test.registryPath}
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - sagemaker.amazonaws.com
  test:
    fn::invoke:
      Function: aws:sagemaker:getPrebuiltEcrImage
      Arguments:
        repositoryName: kmeans
```
{{% /example %}}
{{% /examples %}}
## Inference Execution Config

* `mode` - (Required) How containers in a multi-container are run. The following values are valid `Serial` and `Direct`.


## Import

Using `pulumi import`, import models using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/model:Model test_model model-foo
```
 