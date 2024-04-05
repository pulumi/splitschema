Provides an AWS Backup vault lock configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.backup.VaultLockConfiguration("test", {
    backupVaultName: "example_backup_vault",
    changeableForDays: 3,
    maxRetentionDays: 1200,
    minRetentionDays: 7,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.backup.VaultLockConfiguration("test",
    backup_vault_name="example_backup_vault",
    changeable_for_days=3,
    max_retention_days=1200,
    min_retention_days=7)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Backup.VaultLockConfiguration("test", new()
    {
        BackupVaultName = "example_backup_vault",
        ChangeableForDays = 3,
        MaxRetentionDays = 1200,
        MinRetentionDays = 7,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewVaultLockConfiguration(ctx, "test", &backup.VaultLockConfigurationArgs{
			BackupVaultName:   pulumi.String("example_backup_vault"),
			ChangeableForDays: pulumi.Int(3),
			MaxRetentionDays:  pulumi.Int(1200),
			MinRetentionDays:  pulumi.Int(7),
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
import com.pulumi.aws.backup.VaultLockConfiguration;
import com.pulumi.aws.backup.VaultLockConfigurationArgs;
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
        var test = new VaultLockConfiguration("test", VaultLockConfigurationArgs.builder()        
            .backupVaultName("example_backup_vault")
            .changeableForDays(3)
            .maxRetentionDays(1200)
            .minRetentionDays(7)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:backup:VaultLockConfiguration
    properties:
      backupVaultName: example_backup_vault
      changeableForDays: 3
      maxRetentionDays: 1200
      minRetentionDays: 7
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup vault lock configuration using the `name`. For example:

```sh
 $ pulumi import aws:backup/vaultLockConfiguration:VaultLockConfiguration test TestVault
```
 