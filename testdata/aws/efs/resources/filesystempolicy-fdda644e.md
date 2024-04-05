Provides an Elastic File System (EFS) File System Policy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fs = new aws.efs.FileSystem("fs", {});
const policyPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        sid: "ExampleStatement01",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["*"],
        }],
        actions: [
            "elasticfilesystem:ClientMount",
            "elasticfilesystem:ClientWrite",
        ],
        resources: [fs.arn],
        conditions: [{
            test: "Bool",
            variable: "aws:SecureTransport",
            values: ["true"],
        }],
    }],
});
const policyFileSystemPolicy = new aws.efs.FileSystemPolicy("policyFileSystemPolicy", {
    fileSystemId: fs.id,
    policy: policyPolicyDocument.apply(policyPolicyDocument => policyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

fs = aws.efs.FileSystem("fs")
policy_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="ExampleStatement01",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["*"],
    )],
    actions=[
        "elasticfilesystem:ClientMount",
        "elasticfilesystem:ClientWrite",
    ],
    resources=[fs.arn],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="Bool",
        variable="aws:SecureTransport",
        values=["true"],
    )],
)])
policy_file_system_policy = aws.efs.FileSystemPolicy("policyFileSystemPolicy",
    file_system_id=fs.id,
    policy=policy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fs = new Aws.Efs.FileSystem("fs");

    var policyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "ExampleStatement01",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "elasticfilesystem:ClientMount",
                    "elasticfilesystem:ClientWrite",
                },
                Resources = new[]
                {
                    fs.Arn,
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

    var policyFileSystemPolicy = new Aws.Efs.FileSystemPolicy("policyFileSystemPolicy", new()
    {
        FileSystemId = fs.Id,
        Policy = policyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fs, err := efs.NewFileSystem(ctx, "fs", nil)
		if err != nil {
			return err
		}
		policyPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Sid:    pulumi.String("ExampleStatement01"),
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("AWS"),
							Identifiers: pulumi.StringArray{
								pulumi.String("*"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("elasticfilesystem:ClientMount"),
						pulumi.String("elasticfilesystem:ClientWrite"),
					},
					Resources: pulumi.StringArray{
						fs.Arn,
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
		_, err = efs.NewFileSystemPolicy(ctx, "policyFileSystemPolicy", &efs.FileSystemPolicyArgs{
			FileSystemId: fs.ID(),
			Policy: policyPolicyDocument.ApplyT(func(policyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &policyPolicyDocument.Json, nil
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
import com.pulumi.aws.efs.FileSystem;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.efs.FileSystemPolicy;
import com.pulumi.aws.efs.FileSystemPolicyArgs;
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
        var fs = new FileSystem("fs");

        final var policyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("ExampleStatement01")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("*")
                    .build())
                .actions(                
                    "elasticfilesystem:ClientMount",
                    "elasticfilesystem:ClientWrite")
                .resources(fs.arn())
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("Bool")
                    .variable("aws:SecureTransport")
                    .values("true")
                    .build())
                .build())
            .build());

        var policyFileSystemPolicy = new FileSystemPolicy("policyFileSystemPolicy", FileSystemPolicyArgs.builder()        
            .fileSystemId(fs.id())
            .policy(policyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(policyPolicyDocument -> policyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  fs:
    type: aws:efs:FileSystem
  policyFileSystemPolicy:
    type: aws:efs:FileSystemPolicy
    properties:
      fileSystemId: ${fs.id}
      policy: ${policyPolicyDocument.json}
variables:
  policyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: ExampleStatement01
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '*'
            actions:
              - elasticfilesystem:ClientMount
              - elasticfilesystem:ClientWrite
            resources:
              - ${fs.arn}
            conditions:
              - test: Bool
                variable: aws:SecureTransport
                values:
                  - 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the EFS file system policies using the `id`. For example:

```sh
 $ pulumi import aws:efs/fileSystemPolicy:FileSystemPolicy foo fs-6fa144c6
```
 