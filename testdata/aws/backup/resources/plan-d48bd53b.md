Provides an AWS Backup plan resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Plan("example", {
    rules: [{
        ruleName: "my_example_backup_rule",
        targetVaultName: aws_backup_vault.test.name,
        schedule: "cron(0 12 * * ? *)",
        lifecycle: {
            deleteAfter: 14,
        },
    }],
    advancedBackupSettings: [{
        backupOptions: {
            WindowsVSS: "enabled",
        },
        resourceType: "EC2",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.Plan("example",
    rules=[aws.backup.PlanRuleArgs(
        rule_name="my_example_backup_rule",
        target_vault_name=aws_backup_vault["test"]["name"],
        schedule="cron(0 12 * * ? *)",
        lifecycle=aws.backup.PlanRuleLifecycleArgs(
            delete_after=14,
        ),
    )],
    advanced_backup_settings=[aws.backup.PlanAdvancedBackupSettingArgs(
        backup_options={
            "WindowsVSS": "enabled",
        },
        resource_type="EC2",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.Plan("example", new()
    {
        Rules = new[]
        {
            new Aws.Backup.Inputs.PlanRuleArgs
            {
                RuleName = "my_example_backup_rule",
                TargetVaultName = aws_backup_vault.Test.Name,
                Schedule = "cron(0 12 * * ? *)",
                Lifecycle = new Aws.Backup.Inputs.PlanRuleLifecycleArgs
                {
                    DeleteAfter = 14,
                },
            },
        },
        AdvancedBackupSettings = new[]
        {
            new Aws.Backup.Inputs.PlanAdvancedBackupSettingArgs
            {
                BackupOptions = 
                {
                    { "WindowsVSS", "enabled" },
                },
                ResourceType = "EC2",
            },
        },
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
		_, err := backup.NewPlan(ctx, "example", &backup.PlanArgs{
			Rules: backup.PlanRuleArray{
				&backup.PlanRuleArgs{
					RuleName:        pulumi.String("my_example_backup_rule"),
					TargetVaultName: pulumi.Any(aws_backup_vault.Test.Name),
					Schedule:        pulumi.String("cron(0 12 * * ? *)"),
					Lifecycle: &backup.PlanRuleLifecycleArgs{
						DeleteAfter: pulumi.Int(14),
					},
				},
			},
			AdvancedBackupSettings: backup.PlanAdvancedBackupSettingArray{
				&backup.PlanAdvancedBackupSettingArgs{
					BackupOptions: pulumi.StringMap{
						"WindowsVSS": pulumi.String("enabled"),
					},
					ResourceType: pulumi.String("EC2"),
				},
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
import com.pulumi.aws.backup.Plan;
import com.pulumi.aws.backup.PlanArgs;
import com.pulumi.aws.backup.inputs.PlanRuleArgs;
import com.pulumi.aws.backup.inputs.PlanRuleLifecycleArgs;
import com.pulumi.aws.backup.inputs.PlanAdvancedBackupSettingArgs;
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
        var example = new Plan("example", PlanArgs.builder()        
            .rules(PlanRuleArgs.builder()
                .ruleName("my_example_backup_rule")
                .targetVaultName(aws_backup_vault.test().name())
                .schedule("cron(0 12 * * ? *)")
                .lifecycle(PlanRuleLifecycleArgs.builder()
                    .deleteAfter(14)
                    .build())
                .build())
            .advancedBackupSettings(PlanAdvancedBackupSettingArgs.builder()
                .backupOptions(Map.of("WindowsVSS", "enabled"))
                .resourceType("EC2")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:Plan
    properties:
      rules:
        - ruleName: my_example_backup_rule
          targetVaultName: ${aws_backup_vault.test.name}
          schedule: cron(0 12 * * ? *)
          lifecycle:
            deleteAfter: 14
      advancedBackupSettings:
        - backupOptions:
            WindowsVSS: enabled
          resourceType: EC2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup Plan using the `id`. For example:

```sh
 $ pulumi import aws:backup/plan:Plan test <id>
```
 