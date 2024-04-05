Provides an RDS DB option group resource. Documentation of the available options for various RDS engines can be found at:

* [MariaDB Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MariaDB.Options.html)
* [Microsoft SQL Server Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.Options.html)
* [MySQL Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.MySQL.Options.html)
* [Oracle Options](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.Options.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.OptionGroup("example", {
    optionGroupDescription: "Option Group",
    engineName: "sqlserver-ee",
    majorEngineVersion: "11.00",
    options: [
        {
            optionName: "Timezone",
            optionSettings: [{
                name: "TIME_ZONE",
                value: "UTC",
            }],
        },
        {
            optionName: "SQLSERVER_BACKUP_RESTORE",
            optionSettings: [{
                name: "IAM_ROLE_ARN",
                value: aws_iam_role.example.arn,
            }],
        },
        {
            optionName: "TDE",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.OptionGroup("example",
    option_group_description="Option Group",
    engine_name="sqlserver-ee",
    major_engine_version="11.00",
    options=[
        aws.rds.OptionGroupOptionArgs(
            option_name="Timezone",
            option_settings=[aws.rds.OptionGroupOptionOptionSettingArgs(
                name="TIME_ZONE",
                value="UTC",
            )],
        ),
        aws.rds.OptionGroupOptionArgs(
            option_name="SQLSERVER_BACKUP_RESTORE",
            option_settings=[aws.rds.OptionGroupOptionOptionSettingArgs(
                name="IAM_ROLE_ARN",
                value=aws_iam_role["example"]["arn"],
            )],
        ),
        aws.rds.OptionGroupOptionArgs(
            option_name="TDE",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.OptionGroup("example", new()
    {
        OptionGroupDescription = "Option Group",
        EngineName = "sqlserver-ee",
        MajorEngineVersion = "11.00",
        Options = new[]
        {
            new Aws.Rds.Inputs.OptionGroupOptionArgs
            {
                OptionName = "Timezone",
                OptionSettings = new[]
                {
                    new Aws.Rds.Inputs.OptionGroupOptionOptionSettingArgs
                    {
                        Name = "TIME_ZONE",
                        Value = "UTC",
                    },
                },
            },
            new Aws.Rds.Inputs.OptionGroupOptionArgs
            {
                OptionName = "SQLSERVER_BACKUP_RESTORE",
                OptionSettings = new[]
                {
                    new Aws.Rds.Inputs.OptionGroupOptionOptionSettingArgs
                    {
                        Name = "IAM_ROLE_ARN",
                        Value = aws_iam_role.Example.Arn,
                    },
                },
            },
            new Aws.Rds.Inputs.OptionGroupOptionArgs
            {
                OptionName = "TDE",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewOptionGroup(ctx, "example", &rds.OptionGroupArgs{
			OptionGroupDescription: pulumi.String("Option Group"),
			EngineName:             pulumi.String("sqlserver-ee"),
			MajorEngineVersion:     pulumi.String("11.00"),
			Options: rds.OptionGroupOptionArray{
				&rds.OptionGroupOptionArgs{
					OptionName: pulumi.String("Timezone"),
					OptionSettings: rds.OptionGroupOptionOptionSettingArray{
						&rds.OptionGroupOptionOptionSettingArgs{
							Name:  pulumi.String("TIME_ZONE"),
							Value: pulumi.String("UTC"),
						},
					},
				},
				&rds.OptionGroupOptionArgs{
					OptionName: pulumi.String("SQLSERVER_BACKUP_RESTORE"),
					OptionSettings: rds.OptionGroupOptionOptionSettingArray{
						&rds.OptionGroupOptionOptionSettingArgs{
							Name:  pulumi.String("IAM_ROLE_ARN"),
							Value: pulumi.Any(aws_iam_role.Example.Arn),
						},
					},
				},
				&rds.OptionGroupOptionArgs{
					OptionName: pulumi.String("TDE"),
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
import com.pulumi.aws.rds.OptionGroup;
import com.pulumi.aws.rds.OptionGroupArgs;
import com.pulumi.aws.rds.inputs.OptionGroupOptionArgs;
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
        var example = new OptionGroup("example", OptionGroupArgs.builder()        
            .optionGroupDescription("Option Group")
            .engineName("sqlserver-ee")
            .majorEngineVersion("11.00")
            .options(            
                OptionGroupOptionArgs.builder()
                    .optionName("Timezone")
                    .optionSettings(OptionGroupOptionOptionSettingArgs.builder()
                        .name("TIME_ZONE")
                        .value("UTC")
                        .build())
                    .build(),
                OptionGroupOptionArgs.builder()
                    .optionName("SQLSERVER_BACKUP_RESTORE")
                    .optionSettings(OptionGroupOptionOptionSettingArgs.builder()
                        .name("IAM_ROLE_ARN")
                        .value(aws_iam_role.example().arn())
                        .build())
                    .build(),
                OptionGroupOptionArgs.builder()
                    .optionName("TDE")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:OptionGroup
    properties:
      optionGroupDescription: Option Group
      engineName: sqlserver-ee
      majorEngineVersion: '11.00'
      options:
        - optionName: Timezone
          optionSettings:
            - name: TIME_ZONE
              value: UTC
        - optionName: SQLSERVER_BACKUP_RESTORE
          optionSettings:
            - name: IAM_ROLE_ARN
              value: ${aws_iam_role.example.arn}
        - optionName: TDE
```

> **Note:** Any modifications to the `aws.rds.OptionGroup` are set to happen immediately as we default to applying immediately.

> **WARNING:** You can perform a destroy on a `aws.rds.OptionGroup`, as long as it is not associated with any Amazon RDS resource. An option group can be associated with a DB instance, a manual DB snapshot, or an automated DB snapshot.

If you try to delete an option group that is associated with an Amazon RDS resource, an error similar to the following is returned:

> An error occurred (InvalidOptionGroupStateFault) when calling the DeleteOptionGroup operation: The option group 'optionGroupName' cannot be deleted because it is in use.

More information about this can be found [here](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithOptionGroups.html#USER_WorkingWithOptionGroups.Delete).
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DB option groups using the `name`. For example:

```sh
 $ pulumi import aws:rds/optionGroup:OptionGroup example mysql-option-group
```
 