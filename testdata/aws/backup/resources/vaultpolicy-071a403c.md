Provides an AWS Backup vault policy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVault = new aws.backup.Vault("exampleVault", {});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: ["*"],
        }],
        actions: [
            "backup:DescribeBackupVault",
            "backup:DeleteBackupVault",
            "backup:PutBackupVaultAccessPolicy",
            "backup:DeleteBackupVaultAccessPolicy",
            "backup:GetBackupVaultAccessPolicy",
            "backup:StartBackupJob",
            "backup:GetBackupVaultNotifications",
            "backup:PutBackupVaultNotifications",
        ],
        resources: [exampleVault.arn],
    }],
});
const exampleVaultPolicy = new aws.backup.VaultPolicy("exampleVaultPolicy", {
    backupVaultName: exampleVault.name,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_vault = aws.backup.Vault("exampleVault")
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=["*"],
    )],
    actions=[
        "backup:DescribeBackupVault",
        "backup:DeleteBackupVault",
        "backup:PutBackupVaultAccessPolicy",
        "backup:DeleteBackupVaultAccessPolicy",
        "backup:GetBackupVaultAccessPolicy",
        "backup:StartBackupJob",
        "backup:GetBackupVaultNotifications",
        "backup:PutBackupVaultNotifications",
    ],
    resources=[example_vault.arn],
)])
example_vault_policy = aws.backup.VaultPolicy("exampleVaultPolicy",
    backup_vault_name=example_vault.name,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVault = new Aws.Backup.Vault("exampleVault");

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
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "backup:DescribeBackupVault",
                    "backup:DeleteBackupVault",
                    "backup:PutBackupVaultAccessPolicy",
                    "backup:DeleteBackupVaultAccessPolicy",
                    "backup:GetBackupVaultAccessPolicy",
                    "backup:StartBackupJob",
                    "backup:GetBackupVaultNotifications",
                    "backup:PutBackupVaultNotifications",
                },
                Resources = new[]
                {
                    exampleVault.Arn,
                },
            },
        },
    });

    var exampleVaultPolicy = new Aws.Backup.VaultPolicy("exampleVaultPolicy", new()
    {
        BackupVaultName = exampleVault.Name,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVault, err := backup.NewVault(ctx, "exampleVault", nil)
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
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
						pulumi.String("backup:DescribeBackupVault"),
						pulumi.String("backup:DeleteBackupVault"),
						pulumi.String("backup:PutBackupVaultAccessPolicy"),
						pulumi.String("backup:DeleteBackupVaultAccessPolicy"),
						pulumi.String("backup:GetBackupVaultAccessPolicy"),
						pulumi.String("backup:StartBackupJob"),
						pulumi.String("backup:GetBackupVaultNotifications"),
						pulumi.String("backup:PutBackupVaultNotifications"),
					},
					Resources: pulumi.StringArray{
						exampleVault.Arn,
					},
				},
			},
		}, nil)
		_, err = backup.NewVaultPolicy(ctx, "exampleVaultPolicy", &backup.VaultPolicyArgs{
			BackupVaultName: exampleVault.Name,
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
import com.pulumi.aws.backup.Vault;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.backup.VaultPolicy;
import com.pulumi.aws.backup.VaultPolicyArgs;
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
        var exampleVault = new Vault("exampleVault");

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers("*")
                    .build())
                .actions(                
                    "backup:DescribeBackupVault",
                    "backup:DeleteBackupVault",
                    "backup:PutBackupVaultAccessPolicy",
                    "backup:DeleteBackupVaultAccessPolicy",
                    "backup:GetBackupVaultAccessPolicy",
                    "backup:StartBackupJob",
                    "backup:GetBackupVaultNotifications",
                    "backup:PutBackupVaultNotifications")
                .resources(exampleVault.arn())
                .build())
            .build());

        var exampleVaultPolicy = new VaultPolicy("exampleVaultPolicy", VaultPolicyArgs.builder()        
            .backupVaultName(exampleVault.name())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  exampleVault:
    type: aws:backup:Vault
  exampleVaultPolicy:
    type: aws:backup:VaultPolicy
    properties:
      backupVaultName: ${exampleVault.name}
      policy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - '*'
            actions:
              - backup:DescribeBackupVault
              - backup:DeleteBackupVault
              - backup:PutBackupVaultAccessPolicy
              - backup:DeleteBackupVaultAccessPolicy
              - backup:GetBackupVaultAccessPolicy
              - backup:StartBackupJob
              - backup:GetBackupVaultNotifications
              - backup:PutBackupVaultNotifications
            resources:
              - ${exampleVault.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup vault policy using the `name`. For example:

```sh
 $ pulumi import aws:backup/vaultPolicy:VaultPolicy test TestVault
```
 