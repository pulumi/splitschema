Manages Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.

> **NOTE:** Lake Formation introduces fine-grained access control for data in your data lake. Part of the changes include the `IAMAllowedPrincipals` principal in order to make Lake Formation backwards compatible with existing IAM and Glue permissions. For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) and [Upgrading AWS Glue Data Permissions to the AWS Lake Formation Model](https://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Data Lake Admins

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lakeformation.DataLakeSettings("example", {admins: [
    aws_iam_user.test.arn,
    aws_iam_role.test.arn,
]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lakeformation.DataLakeSettings("example", admins=[
    aws_iam_user["test"]["arn"],
    aws_iam_role["test"]["arn"],
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LakeFormation.DataLakeSettings("example", new()
    {
        Admins = new[]
        {
            aws_iam_user.Test.Arn,
            aws_iam_role.Test.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.NewDataLakeSettings(ctx, "example", &lakeformation.DataLakeSettingsArgs{
			Admins: pulumi.StringArray{
				aws_iam_user.Test.Arn,
				aws_iam_role.Test.Arn,
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
import com.pulumi.aws.lakeformation.DataLakeSettings;
import com.pulumi.aws.lakeformation.DataLakeSettingsArgs;
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
        var example = new DataLakeSettings("example", DataLakeSettingsArgs.builder()        
            .admins(            
                aws_iam_user.test().arn(),
                aws_iam_role.test().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lakeformation:DataLakeSettings
    properties:
      admins:
        - ${aws_iam_user.test.arn}
        - ${aws_iam_role.test.arn}
```
{{% /example %}}
{{% example %}}
### Create Default Permissions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lakeformation.DataLakeSettings("example", {
    admins: [
        aws_iam_user.test.arn,
        aws_iam_role.test.arn,
    ],
    createDatabaseDefaultPermissions: [{
        permissions: [
            "SELECT",
            "ALTER",
            "DROP",
        ],
        principal: aws_iam_user.test.arn,
    }],
    createTableDefaultPermissions: [{
        permissions: ["ALL"],
        principal: aws_iam_role.test.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lakeformation.DataLakeSettings("example",
    admins=[
        aws_iam_user["test"]["arn"],
        aws_iam_role["test"]["arn"],
    ],
    create_database_default_permissions=[aws.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs(
        permissions=[
            "SELECT",
            "ALTER",
            "DROP",
        ],
        principal=aws_iam_user["test"]["arn"],
    )],
    create_table_default_permissions=[aws.lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs(
        permissions=["ALL"],
        principal=aws_iam_role["test"]["arn"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LakeFormation.DataLakeSettings("example", new()
    {
        Admins = new[]
        {
            aws_iam_user.Test.Arn,
            aws_iam_role.Test.Arn,
        },
        CreateDatabaseDefaultPermissions = new[]
        {
            new Aws.LakeFormation.Inputs.DataLakeSettingsCreateDatabaseDefaultPermissionArgs
            {
                Permissions = new[]
                {
                    "SELECT",
                    "ALTER",
                    "DROP",
                },
                Principal = aws_iam_user.Test.Arn,
            },
        },
        CreateTableDefaultPermissions = new[]
        {
            new Aws.LakeFormation.Inputs.DataLakeSettingsCreateTableDefaultPermissionArgs
            {
                Permissions = new[]
                {
                    "ALL",
                },
                Principal = aws_iam_role.Test.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.NewDataLakeSettings(ctx, "example", &lakeformation.DataLakeSettingsArgs{
			Admins: pulumi.StringArray{
				aws_iam_user.Test.Arn,
				aws_iam_role.Test.Arn,
			},
			CreateDatabaseDefaultPermissions: lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArray{
				&lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("SELECT"),
						pulumi.String("ALTER"),
						pulumi.String("DROP"),
					},
					Principal: pulumi.Any(aws_iam_user.Test.Arn),
				},
			},
			CreateTableDefaultPermissions: lakeformation.DataLakeSettingsCreateTableDefaultPermissionArray{
				&lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("ALL"),
					},
					Principal: pulumi.Any(aws_iam_role.Test.Arn),
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
import com.pulumi.aws.lakeformation.DataLakeSettings;
import com.pulumi.aws.lakeformation.DataLakeSettingsArgs;
import com.pulumi.aws.lakeformation.inputs.DataLakeSettingsCreateDatabaseDefaultPermissionArgs;
import com.pulumi.aws.lakeformation.inputs.DataLakeSettingsCreateTableDefaultPermissionArgs;
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
        var example = new DataLakeSettings("example", DataLakeSettingsArgs.builder()        
            .admins(            
                aws_iam_user.test().arn(),
                aws_iam_role.test().arn())
            .createDatabaseDefaultPermissions(DataLakeSettingsCreateDatabaseDefaultPermissionArgs.builder()
                .permissions(                
                    "SELECT",
                    "ALTER",
                    "DROP")
                .principal(aws_iam_user.test().arn())
                .build())
            .createTableDefaultPermissions(DataLakeSettingsCreateTableDefaultPermissionArgs.builder()
                .permissions("ALL")
                .principal(aws_iam_role.test().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lakeformation:DataLakeSettings
    properties:
      admins:
        - ${aws_iam_user.test.arn}
        - ${aws_iam_role.test.arn}
      createDatabaseDefaultPermissions:
        - permissions:
            - SELECT
            - ALTER
            - DROP
          principal: ${aws_iam_user.test.arn}
      createTableDefaultPermissions:
        - permissions:
            - ALL
          principal: ${aws_iam_role.test.arn}
```
{{% /example %}}
{{% example %}}
### Enable EMR access to LakeFormation resources

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lakeformation.DataLakeSettings("example", {
    admins: [
        aws_iam_user.test.arn,
        aws_iam_role.test.arn,
    ],
    createDatabaseDefaultPermissions: [{
        permissions: [
            "SELECT",
            "ALTER",
            "DROP",
        ],
        principal: aws_iam_user.test.arn,
    }],
    createTableDefaultPermissions: [{
        permissions: ["ALL"],
        principal: aws_iam_role.test.arn,
    }],
    allowExternalDataFiltering: true,
    externalDataFilteringAllowLists: [
        data.aws_caller_identity.current.account_id,
        data.aws_caller_identity.third_party.account_id,
    ],
    authorizedSessionTagValueLists: ["Amazon EMR"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lakeformation.DataLakeSettings("example",
    admins=[
        aws_iam_user["test"]["arn"],
        aws_iam_role["test"]["arn"],
    ],
    create_database_default_permissions=[aws.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs(
        permissions=[
            "SELECT",
            "ALTER",
            "DROP",
        ],
        principal=aws_iam_user["test"]["arn"],
    )],
    create_table_default_permissions=[aws.lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs(
        permissions=["ALL"],
        principal=aws_iam_role["test"]["arn"],
    )],
    allow_external_data_filtering=True,
    external_data_filtering_allow_lists=[
        data["aws_caller_identity"]["current"]["account_id"],
        data["aws_caller_identity"]["third_party"]["account_id"],
    ],
    authorized_session_tag_value_lists=["Amazon EMR"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LakeFormation.DataLakeSettings("example", new()
    {
        Admins = new[]
        {
            aws_iam_user.Test.Arn,
            aws_iam_role.Test.Arn,
        },
        CreateDatabaseDefaultPermissions = new[]
        {
            new Aws.LakeFormation.Inputs.DataLakeSettingsCreateDatabaseDefaultPermissionArgs
            {
                Permissions = new[]
                {
                    "SELECT",
                    "ALTER",
                    "DROP",
                },
                Principal = aws_iam_user.Test.Arn,
            },
        },
        CreateTableDefaultPermissions = new[]
        {
            new Aws.LakeFormation.Inputs.DataLakeSettingsCreateTableDefaultPermissionArgs
            {
                Permissions = new[]
                {
                    "ALL",
                },
                Principal = aws_iam_role.Test.Arn,
            },
        },
        AllowExternalDataFiltering = true,
        ExternalDataFilteringAllowLists = new[]
        {
            data.Aws_caller_identity.Current.Account_id,
            data.Aws_caller_identity.Third_party.Account_id,
        },
        AuthorizedSessionTagValueLists = new[]
        {
            "Amazon EMR",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.NewDataLakeSettings(ctx, "example", &lakeformation.DataLakeSettingsArgs{
			Admins: pulumi.StringArray{
				aws_iam_user.Test.Arn,
				aws_iam_role.Test.Arn,
			},
			CreateDatabaseDefaultPermissions: lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArray{
				&lakeformation.DataLakeSettingsCreateDatabaseDefaultPermissionArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("SELECT"),
						pulumi.String("ALTER"),
						pulumi.String("DROP"),
					},
					Principal: pulumi.Any(aws_iam_user.Test.Arn),
				},
			},
			CreateTableDefaultPermissions: lakeformation.DataLakeSettingsCreateTableDefaultPermissionArray{
				&lakeformation.DataLakeSettingsCreateTableDefaultPermissionArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("ALL"),
					},
					Principal: pulumi.Any(aws_iam_role.Test.Arn),
				},
			},
			AllowExternalDataFiltering: pulumi.Bool(true),
			ExternalDataFilteringAllowLists: pulumi.StringArray{
				data.Aws_caller_identity.Current.Account_id,
				data.Aws_caller_identity.Third_party.Account_id,
			},
			AuthorizedSessionTagValueLists: pulumi.StringArray{
				pulumi.String("Amazon EMR"),
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
import com.pulumi.aws.lakeformation.DataLakeSettings;
import com.pulumi.aws.lakeformation.DataLakeSettingsArgs;
import com.pulumi.aws.lakeformation.inputs.DataLakeSettingsCreateDatabaseDefaultPermissionArgs;
import com.pulumi.aws.lakeformation.inputs.DataLakeSettingsCreateTableDefaultPermissionArgs;
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
        var example = new DataLakeSettings("example", DataLakeSettingsArgs.builder()        
            .admins(            
                aws_iam_user.test().arn(),
                aws_iam_role.test().arn())
            .createDatabaseDefaultPermissions(DataLakeSettingsCreateDatabaseDefaultPermissionArgs.builder()
                .permissions(                
                    "SELECT",
                    "ALTER",
                    "DROP")
                .principal(aws_iam_user.test().arn())
                .build())
            .createTableDefaultPermissions(DataLakeSettingsCreateTableDefaultPermissionArgs.builder()
                .permissions("ALL")
                .principal(aws_iam_role.test().arn())
                .build())
            .allowExternalDataFiltering(true)
            .externalDataFilteringAllowLists(            
                data.aws_caller_identity().current().account_id(),
                data.aws_caller_identity().third_party().account_id())
            .authorizedSessionTagValueLists("Amazon EMR")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lakeformation:DataLakeSettings
    properties:
      admins:
        - ${aws_iam_user.test.arn}
        - ${aws_iam_role.test.arn}
      createDatabaseDefaultPermissions:
        - permissions:
            - SELECT
            - ALTER
            - DROP
          principal: ${aws_iam_user.test.arn}
      createTableDefaultPermissions:
        - permissions:
            - ALL
          principal: ${aws_iam_role.test.arn}
      allowExternalDataFiltering: true
      externalDataFilteringAllowLists:
        - ${data.aws_caller_identity.current.account_id}
        - ${data.aws_caller_identity.third_party.account_id}
      authorizedSessionTagValueLists:
        - Amazon EMR
```
{{% /example %}}
{{% /examples %}}