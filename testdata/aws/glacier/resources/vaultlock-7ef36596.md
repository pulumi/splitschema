Manages a Glacier Vault Lock. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-lock.html) for a full explanation of the Glacier Vault Lock functionality.

> **NOTE:** This resource allows you to test Glacier Vault Lock policies by setting the `complete_lock` argument to `false`. When testing policies in this manner, the Glacier Vault Lock automatically expires after 24 hours and this provider will show this resource as needing recreation after that time. To permanently apply the policy, set the `complete_lock` argument to `true`. When changing `complete_lock` to `true`, it is expected the resource will show as recreating.

!> **WARNING:** Once a Glacier Vault Lock is completed, it is immutable. The deletion of the Glacier Vault Lock is not be possible and attempting to remove it from this provider will return an error. Set the `ignore_deletion_error` argument to `true` and apply this configuration before attempting to delete this resource via this provider or remove this resource from this provider's management.

{{% examples %}}
## Example Usage
{{% example %}}
### Testing Glacier Vault Lock Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVault = new aws.glacier.Vault("exampleVault", {});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        actions: ["glacier:DeleteArchive"],
        effect: "Deny",
        resources: [exampleVault.arn],
        conditions: [{
            test: "NumericLessThanEquals",
            variable: "glacier:ArchiveAgeinDays",
            values: ["365"],
        }],
    }],
});
const exampleVaultLock = new aws.glacier.VaultLock("exampleVaultLock", {
    completeLock: false,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
    vaultName: exampleVault.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example_vault = aws.glacier.Vault("exampleVault")
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["glacier:DeleteArchive"],
    effect="Deny",
    resources=[example_vault.arn],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="NumericLessThanEquals",
        variable="glacier:ArchiveAgeinDays",
        values=["365"],
    )],
)])
example_vault_lock = aws.glacier.VaultLock("exampleVaultLock",
    complete_lock=False,
    policy=example_policy_document.json,
    vault_name=example_vault.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVault = new Aws.Glacier.Vault("exampleVault");

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "glacier:DeleteArchive",
                },
                Effect = "Deny",
                Resources = new[]
                {
                    exampleVault.Arn,
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "NumericLessThanEquals",
                        Variable = "glacier:ArchiveAgeinDays",
                        Values = new[]
                        {
                            "365",
                        },
                    },
                },
            },
        },
    });

    var exampleVaultLock = new Aws.Glacier.VaultLock("exampleVaultLock", new()
    {
        CompleteLock = false,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        VaultName = exampleVault.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glacier"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVault, err := glacier.NewVault(ctx, "exampleVault", nil)
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("glacier:DeleteArchive"),
					},
					Effect: pulumi.String("Deny"),
					Resources: pulumi.StringArray{
						exampleVault.Arn,
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("NumericLessThanEquals"),
							Variable: pulumi.String("glacier:ArchiveAgeinDays"),
							Values: pulumi.StringArray{
								pulumi.String("365"),
							},
						},
					},
				},
			},
		}, nil)
		_, err = glacier.NewVaultLock(ctx, "exampleVaultLock", &glacier.VaultLockArgs{
			CompleteLock: pulumi.Bool(false),
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &examplePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
			VaultName: exampleVault.Name,
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
import com.pulumi.aws.glacier.Vault;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.glacier.VaultLock;
import com.pulumi.aws.glacier.VaultLockArgs;
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
                .actions("glacier:DeleteArchive")
                .effect("Deny")
                .resources(exampleVault.arn())
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("NumericLessThanEquals")
                    .variable("glacier:ArchiveAgeinDays")
                    .values("365")
                    .build())
                .build())
            .build());

        var exampleVaultLock = new VaultLock("exampleVaultLock", VaultLockArgs.builder()        
            .completeLock(false)
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .vaultName(exampleVault.name())
            .build());

    }
}
```
```yaml
resources:
  exampleVault:
    type: aws:glacier:Vault
  exampleVaultLock:
    type: aws:glacier:VaultLock
    properties:
      completeLock: false
      policy: ${examplePolicyDocument.json}
      vaultName: ${exampleVault.name}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - glacier:DeleteArchive
            effect: Deny
            resources:
              - ${exampleVault.arn}
            conditions:
              - test: NumericLessThanEquals
                variable: glacier:ArchiveAgeinDays
                values:
                  - '365'
```
{{% /example %}}
{{% example %}}
### Permanently Applying Glacier Vault Lock Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glacier.VaultLock("example", {
    completeLock: true,
    policy: data.aws_iam_policy_document.example.json,
    vaultName: aws_glacier_vault.example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glacier.VaultLock("example",
    complete_lock=True,
    policy=data["aws_iam_policy_document"]["example"]["json"],
    vault_name=aws_glacier_vault["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glacier.VaultLock("example", new()
    {
        CompleteLock = true,
        Policy = data.Aws_iam_policy_document.Example.Json,
        VaultName = aws_glacier_vault.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glacier"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glacier.NewVaultLock(ctx, "example", &glacier.VaultLockArgs{
			CompleteLock: pulumi.Bool(true),
			Policy:       pulumi.Any(data.Aws_iam_policy_document.Example.Json),
			VaultName:    pulumi.Any(aws_glacier_vault.Example.Name),
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
import com.pulumi.aws.glacier.VaultLock;
import com.pulumi.aws.glacier.VaultLockArgs;
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
        var example = new VaultLock("example", VaultLockArgs.builder()        
            .completeLock(true)
            .policy(data.aws_iam_policy_document().example().json())
            .vaultName(aws_glacier_vault.example().name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glacier:VaultLock
    properties:
      completeLock: true
      policy: ${data.aws_iam_policy_document.example.json}
      vaultName: ${aws_glacier_vault.example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glacier Vault Locks using the Glacier Vault name. For example:

```sh
 $ pulumi import aws:glacier/vaultLock:VaultLock example example-vault
```
 