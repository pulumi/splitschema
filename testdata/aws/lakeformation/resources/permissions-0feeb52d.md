Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, tables, LF-tags, and LF-tag policies. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).

!> **WARNING:** Lake Formation permissions are not in effect by default within AWS. Using this resource will not secure your data and will result in errors if you do not change the security settings for existing resources and the default security settings for new resources. See Default Behavior and `IAMAllowedPrincipals` for additional details.

> **NOTE:** In general, the `principal` should _NOT_ be a Lake Formation administrator or the entity (e.g., IAM role) that is running the deployment. Administrators have implicit permissions. These should be managed by granting or not granting administrator rights using `aws.lakeformation.DataLakeSettings`, _not_ with this resource.

## Default Behavior and `IAMAllowedPrincipals`

**_Lake Formation permissions are not in effect by default within AWS._** `IAMAllowedPrincipals` (i.e., `IAM_ALLOWED_PRINCIPALS`) conflicts with individual Lake Formation permissions (i.e., non-`IAMAllowedPrincipals` permissions), will cause unexpected behavior, and may result in errors.

When using Lake Formation, choose ONE of the following options as they are mutually exclusive:

1. Use this resource (`aws.lakeformation.Permissions`), change the default security settings using `aws.lakeformation.DataLakeSettings`, and remove existing `IAMAllowedPrincipals` permissions
2. Use `IAMAllowedPrincipals` without `aws.lakeformation.Permissions`

This example shows removing the `IAMAllowedPrincipals` default security settings and making the caller a Lake Formation admin. Since `create_database_default_permissions` and `create_table_default_permissions` are not set in the `aws.lakeformation.DataLakeSettings` resource, they are cleared.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentSessionContext = currentCallerIdentity.then(currentCallerIdentity => aws.iam.getSessionContext({
    arn: currentCallerIdentity.arn,
}));
const test = new aws.lakeformation.DataLakeSettings("test", {admins: [currentSessionContext.then(currentSessionContext => currentSessionContext.issuerArn)]});
```
```python
import pulumi
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_session_context = aws.iam.get_session_context(arn=current_caller_identity.arn)
test = aws.lakeformation.DataLakeSettings("test", admins=[current_session_context.issuer_arn])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentSessionContext = Aws.Iam.GetSessionContext.Invoke(new()
    {
        Arn = currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
    });

    var test = new Aws.LakeFormation.DataLakeSettings("test", new()
    {
        Admins = new[]
        {
            currentSessionContext.Apply(getSessionContextResult => getSessionContextResult.IssuerArn),
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentSessionContext, err := iam.GetSessionContext(ctx, &iam.GetSessionContextArgs{
			Arn: currentCallerIdentity.Arn,
		}, nil)
		if err != nil {
			return err
		}
		_, err = lakeformation.NewDataLakeSettings(ctx, "test", &lakeformation.DataLakeSettingsArgs{
			Admins: pulumi.StringArray{
				*pulumi.String(currentSessionContext.IssuerArn),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetSessionContextArgs;
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
        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentSessionContext = IamFunctions.getSessionContext(GetSessionContextArgs.builder()
            .arn(currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn()))
            .build());

        var test = new DataLakeSettings("test", DataLakeSettingsArgs.builder()        
            .admins(currentSessionContext.applyValue(getSessionContextResult -> getSessionContextResult.issuerArn()))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lakeformation:DataLakeSettings
    properties:
      admins:
        - ${currentSessionContext.issuerArn}
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentSessionContext:
    fn::invoke:
      Function: aws:iam:getSessionContext
      Arguments:
        arn: ${currentCallerIdentity.arn}
```

To remove existing `IAMAllowedPrincipals` permissions, use the [AWS Lake Formation Console](https://console.aws.amazon.com/lakeformation/) or [AWS CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/lakeformation/batch-revoke-permissions.html).

`IAMAllowedPrincipals` is a hook to maintain backwards compatibility with AWS Glue. `IAMAllowedPrincipals` is a pseudo-entity group that acts like a Lake Formation principal. The group includes any IAM users and roles that are allowed access to your Data Catalog resources by your IAM policies.

This is Lake Formation's default behavior:

* Lake Formation grants `Super` permission to `IAMAllowedPrincipals` on all existing AWS Glue Data Catalog resources.
* Lake Formation enables "Use only IAM access control" for new Data Catalog resources.

For more details, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html).

### Problem Using `IAMAllowedPrincipals`

AWS does not support combining `IAMAllowedPrincipals` permissions and non-`IAMAllowedPrincipals` permissions. Doing so results in unexpected permissions and behaviors. For example, this configuration grants a user `SELECT` on a column in a table.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCatalogDatabase = new aws.glue.CatalogDatabase("exampleCatalogDatabase", {name: "sadabate"});
const exampleCatalogTable = new aws.glue.CatalogTable("exampleCatalogTable", {
    name: "abelt",
    databaseName: aws_glue_catalog_database.test.name,
    storageDescriptor: {
        columns: [{
            name: "event",
            type: "string",
        }],
    },
});
const examplePermissions = new aws.lakeformation.Permissions("examplePermissions", {
    permissions: ["SELECT"],
    principal: "arn:aws:iam:us-east-1:123456789012:user/SanHolo",
    tableWithColumns: {
        databaseName: exampleCatalogTable.databaseName,
        name: exampleCatalogTable.name,
        columnNames: ["event"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_catalog_database = aws.glue.CatalogDatabase("exampleCatalogDatabase", name="sadabate")
example_catalog_table = aws.glue.CatalogTable("exampleCatalogTable",
    name="abelt",
    database_name=aws_glue_catalog_database["test"]["name"],
    storage_descriptor=aws.glue.CatalogTableStorageDescriptorArgs(
        columns=[aws.glue.CatalogTableStorageDescriptorColumnArgs(
            name="event",
            type="string",
        )],
    ))
example_permissions = aws.lakeformation.Permissions("examplePermissions",
    permissions=["SELECT"],
    principal="arn:aws:iam:us-east-1:123456789012:user/SanHolo",
    table_with_columns=aws.lakeformation.PermissionsTableWithColumnsArgs(
        database_name=example_catalog_table.database_name,
        name=example_catalog_table.name,
        column_names=["event"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCatalogDatabase = new Aws.Glue.CatalogDatabase("exampleCatalogDatabase", new()
    {
        Name = "sadabate",
    });

    var exampleCatalogTable = new Aws.Glue.CatalogTable("exampleCatalogTable", new()
    {
        Name = "abelt",
        DatabaseName = aws_glue_catalog_database.Test.Name,
        StorageDescriptor = new Aws.Glue.Inputs.CatalogTableStorageDescriptorArgs
        {
            Columns = new[]
            {
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Name = "event",
                    Type = "string",
                },
            },
        },
    });

    var examplePermissions = new Aws.LakeFormation.Permissions("examplePermissions", new()
    {
        PermissionDetails = new[]
        {
            "SELECT",
        },
        Principal = "arn:aws:iam:us-east-1:123456789012:user/SanHolo",
        TableWithColumns = new Aws.LakeFormation.Inputs.PermissionsTableWithColumnsArgs
        {
            DatabaseName = exampleCatalogTable.DatabaseName,
            Name = exampleCatalogTable.Name,
            ColumnNames = new[]
            {
                "event",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCatalogDatabase(ctx, "exampleCatalogDatabase", &glue.CatalogDatabaseArgs{
			Name: pulumi.String("sadabate"),
		})
		if err != nil {
			return err
		}
		exampleCatalogTable, err := glue.NewCatalogTable(ctx, "exampleCatalogTable", &glue.CatalogTableArgs{
			Name:         pulumi.String("abelt"),
			DatabaseName: pulumi.Any(aws_glue_catalog_database.Test.Name),
			StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
				Columns: glue.CatalogTableStorageDescriptorColumnArray{
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name: pulumi.String("event"),
						Type: pulumi.String("string"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = lakeformation.NewPermissions(ctx, "examplePermissions", &lakeformation.PermissionsArgs{
			Permissions: pulumi.StringArray{
				pulumi.String("SELECT"),
			},
			Principal: pulumi.String("arn:aws:iam:us-east-1:123456789012:user/SanHolo"),
			TableWithColumns: &lakeformation.PermissionsTableWithColumnsArgs{
				DatabaseName: exampleCatalogTable.DatabaseName,
				Name:         exampleCatalogTable.Name,
				ColumnNames: pulumi.StringArray{
					pulumi.String("event"),
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
import com.pulumi.aws.glue.CatalogDatabase;
import com.pulumi.aws.glue.CatalogDatabaseArgs;
import com.pulumi.aws.glue.CatalogTable;
import com.pulumi.aws.glue.CatalogTableArgs;
import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorArgs;
import com.pulumi.aws.lakeformation.Permissions;
import com.pulumi.aws.lakeformation.PermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.PermissionsTableWithColumnsArgs;
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
        var exampleCatalogDatabase = new CatalogDatabase("exampleCatalogDatabase", CatalogDatabaseArgs.builder()        
            .name("sadabate")
            .build());

        var exampleCatalogTable = new CatalogTable("exampleCatalogTable", CatalogTableArgs.builder()        
            .name("abelt")
            .databaseName(aws_glue_catalog_database.test().name())
            .storageDescriptor(CatalogTableStorageDescriptorArgs.builder()
                .columns(CatalogTableStorageDescriptorColumnArgs.builder()
                    .name("event")
                    .type("string")
                    .build())
                .build())
            .build());

        var examplePermissions = new Permissions("examplePermissions", PermissionsArgs.builder()        
            .permissions("SELECT")
            .principal("arn:aws:iam:us-east-1:123456789012:user/SanHolo")
            .tableWithColumns(PermissionsTableWithColumnsArgs.builder()
                .databaseName(exampleCatalogTable.databaseName())
                .name(exampleCatalogTable.name())
                .columnNames("event")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleCatalogDatabase:
    type: aws:glue:CatalogDatabase
    properties:
      name: sadabate
  exampleCatalogTable:
    type: aws:glue:CatalogTable
    properties:
      name: abelt
      databaseName: ${aws_glue_catalog_database.test.name}
      storageDescriptor:
        columns:
          - name: event
            type: string
  examplePermissions:
    type: aws:lakeformation:Permissions
    properties:
      permissions:
        - SELECT
      principal: arn:aws:iam:us-east-1:123456789012:user/SanHolo
      tableWithColumns:
        databaseName: ${exampleCatalogTable.databaseName}
        name: ${exampleCatalogTable.name}
        columnNames:
          - event
```

The resulting permissions depend on whether the table had `IAMAllowedPrincipals` (IAP) permissions or not.

| Result With IAP | Result Without IAP |
| ---- | ---- |
| `SELECT` column wildcard (i.e., all columns) | `SELECT` on `"event"` (as expected) |

## Using Lake Formation Permissions

Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. These implicit permissions cannot be revoked _per se_. If this resource reads implicit permissions, it will attempt to revoke them, which causes an error when the resource is destroyed.

There are two ways to avoid these errors. First, and the way we recommend, is to avoid using this resource with principals that have implicit permissions. A second, error-prone option, is to grant explicit permissions (and `permissions_with_grant_option`) to "overwrite" a principal's implicit permissions, which you can then revoke with this resource. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).

If the `principal` is also a data lake administrator, AWS grants implicit permissions that can cause errors using this resource. For example, AWS implicitly grants a `principal`/administrator `permissions` and `permissions_with_grant_option` of `ALL`, `ALTER`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT` on a table. If you use this resource to explicitly grant the `principal`/administrator `permissions` but _not_ `permissions_with_grant_option` of `ALL`, `ALTER`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT` on the table, this resource will read the implicit `permissions_with_grant_option` and attempt to revoke them when the resource is destroyed. Doing so will cause an `InvalidInputException: No permissions revoked` error because you cannot revoke implicit permissions _per se_. To workaround this problem, explicitly grant the `principal`/administrator `permissions` _and_ `permissions_with_grant_option`, which can then be revoked. Similarly, granting a `principal`/administrator permissions on a table with columns and providing `column_names`, will result in a `InvalidInputException: Permissions modification is invalid` error because you are narrowing the implicit permissions. Instead, set `wildcard` to `true` and remove the `column_names`.

{{% examples %}}
## Example Usage
{{% example %}}
### Grant Permissions For A Lake Formation S3 Resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lakeformation.Permissions("example", {
    principal: aws_iam_role.workflow_role.arn,
    permissions: ["ALL"],
    dataLocation: {
        arn: aws_lakeformation_resource.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lakeformation.Permissions("example",
    principal=aws_iam_role["workflow_role"]["arn"],
    permissions=["ALL"],
    data_location=aws.lakeformation.PermissionsDataLocationArgs(
        arn=aws_lakeformation_resource["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LakeFormation.Permissions("example", new()
    {
        Principal = aws_iam_role.Workflow_role.Arn,
        PermissionDetails = new[]
        {
            "ALL",
        },
        DataLocation = new Aws.LakeFormation.Inputs.PermissionsDataLocationArgs
        {
            Arn = aws_lakeformation_resource.Example.Arn,
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
		_, err := lakeformation.NewPermissions(ctx, "example", &lakeformation.PermissionsArgs{
			Principal: pulumi.Any(aws_iam_role.Workflow_role.Arn),
			Permissions: pulumi.StringArray{
				pulumi.String("ALL"),
			},
			DataLocation: &lakeformation.PermissionsDataLocationArgs{
				Arn: pulumi.Any(aws_lakeformation_resource.Example.Arn),
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
import com.pulumi.aws.lakeformation.Permissions;
import com.pulumi.aws.lakeformation.PermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.PermissionsDataLocationArgs;
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
        var example = new Permissions("example", PermissionsArgs.builder()        
            .principal(aws_iam_role.workflow_role().arn())
            .permissions("ALL")
            .dataLocation(PermissionsDataLocationArgs.builder()
                .arn(aws_lakeformation_resource.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lakeformation:Permissions
    properties:
      principal: ${aws_iam_role.workflow_role.arn}
      permissions:
        - ALL
      dataLocation:
        arn: ${aws_lakeformation_resource.example.arn}
```
{{% /example %}}
{{% example %}}
### Grant Permissions For A Glue Catalog Database

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lakeformation.Permissions("example", {
    principal: aws_iam_role.workflow_role.arn,
    permissions: [
        "CREATE_TABLE",
        "ALTER",
        "DROP",
    ],
    database: {
        name: aws_glue_catalog_database.example.name,
        catalogId: "110376042874",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lakeformation.Permissions("example",
    principal=aws_iam_role["workflow_role"]["arn"],
    permissions=[
        "CREATE_TABLE",
        "ALTER",
        "DROP",
    ],
    database=aws.lakeformation.PermissionsDatabaseArgs(
        name=aws_glue_catalog_database["example"]["name"],
        catalog_id="110376042874",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.LakeFormation.Permissions("example", new()
    {
        Principal = aws_iam_role.Workflow_role.Arn,
        PermissionDetails = new[]
        {
            "CREATE_TABLE",
            "ALTER",
            "DROP",
        },
        Database = new Aws.LakeFormation.Inputs.PermissionsDatabaseArgs
        {
            Name = aws_glue_catalog_database.Example.Name,
            CatalogId = "110376042874",
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
		_, err := lakeformation.NewPermissions(ctx, "example", &lakeformation.PermissionsArgs{
			Principal: pulumi.Any(aws_iam_role.Workflow_role.Arn),
			Permissions: pulumi.StringArray{
				pulumi.String("CREATE_TABLE"),
				pulumi.String("ALTER"),
				pulumi.String("DROP"),
			},
			Database: &lakeformation.PermissionsDatabaseArgs{
				Name:      pulumi.Any(aws_glue_catalog_database.Example.Name),
				CatalogId: pulumi.String("110376042874"),
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
import com.pulumi.aws.lakeformation.Permissions;
import com.pulumi.aws.lakeformation.PermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.PermissionsDatabaseArgs;
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
        var example = new Permissions("example", PermissionsArgs.builder()        
            .principal(aws_iam_role.workflow_role().arn())
            .permissions(            
                "CREATE_TABLE",
                "ALTER",
                "DROP")
            .database(PermissionsDatabaseArgs.builder()
                .name(aws_glue_catalog_database.example().name())
                .catalogId("110376042874")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lakeformation:Permissions
    properties:
      principal: ${aws_iam_role.workflow_role.arn}
      permissions:
        - CREATE_TABLE
        - ALTER
        - DROP
      database:
        name: ${aws_glue_catalog_database.example.name}
        catalogId: '110376042874'
```
{{% /example %}}
{{% example %}}
### Grant Permissions Using Tag-Based Access Control

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lakeformation.Permissions("test", {
    principal: aws_iam_role.sales_role.arn,
    permissions: [
        "CREATE_TABLE",
        "ALTER",
        "DROP",
    ],
    lfTagPolicy: {
        resourceType: "DATABASE",
        expressions: [
            {
                key: "Team",
                values: ["Sales"],
            },
            {
                key: "Environment",
                values: [
                    "Dev",
                    "Production",
                ],
            },
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lakeformation.Permissions("test",
    principal=aws_iam_role["sales_role"]["arn"],
    permissions=[
        "CREATE_TABLE",
        "ALTER",
        "DROP",
    ],
    lf_tag_policy=aws.lakeformation.PermissionsLfTagPolicyArgs(
        resource_type="DATABASE",
        expressions=[
            aws.lakeformation.PermissionsLfTagPolicyExpressionArgs(
                key="Team",
                values=["Sales"],
            ),
            aws.lakeformation.PermissionsLfTagPolicyExpressionArgs(
                key="Environment",
                values=[
                    "Dev",
                    "Production",
                ],
            ),
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LakeFormation.Permissions("test", new()
    {
        Principal = aws_iam_role.Sales_role.Arn,
        PermissionDetails = new[]
        {
            "CREATE_TABLE",
            "ALTER",
            "DROP",
        },
        LfTagPolicy = new Aws.LakeFormation.Inputs.PermissionsLfTagPolicyArgs
        {
            ResourceType = "DATABASE",
            Expressions = new[]
            {
                new Aws.LakeFormation.Inputs.PermissionsLfTagPolicyExpressionArgs
                {
                    Key = "Team",
                    Values = new[]
                    {
                        "Sales",
                    },
                },
                new Aws.LakeFormation.Inputs.PermissionsLfTagPolicyExpressionArgs
                {
                    Key = "Environment",
                    Values = new[]
                    {
                        "Dev",
                        "Production",
                    },
                },
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
		_, err := lakeformation.NewPermissions(ctx, "test", &lakeformation.PermissionsArgs{
			Principal: pulumi.Any(aws_iam_role.Sales_role.Arn),
			Permissions: pulumi.StringArray{
				pulumi.String("CREATE_TABLE"),
				pulumi.String("ALTER"),
				pulumi.String("DROP"),
			},
			LfTagPolicy: &lakeformation.PermissionsLfTagPolicyArgs{
				ResourceType: pulumi.String("DATABASE"),
				Expressions: lakeformation.PermissionsLfTagPolicyExpressionArray{
					&lakeformation.PermissionsLfTagPolicyExpressionArgs{
						Key: pulumi.String("Team"),
						Values: pulumi.StringArray{
							pulumi.String("Sales"),
						},
					},
					&lakeformation.PermissionsLfTagPolicyExpressionArgs{
						Key: pulumi.String("Environment"),
						Values: pulumi.StringArray{
							pulumi.String("Dev"),
							pulumi.String("Production"),
						},
					},
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
import com.pulumi.aws.lakeformation.Permissions;
import com.pulumi.aws.lakeformation.PermissionsArgs;
import com.pulumi.aws.lakeformation.inputs.PermissionsLfTagPolicyArgs;
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
        var test = new Permissions("test", PermissionsArgs.builder()        
            .principal(aws_iam_role.sales_role().arn())
            .permissions(            
                "CREATE_TABLE",
                "ALTER",
                "DROP")
            .lfTagPolicy(PermissionsLfTagPolicyArgs.builder()
                .resourceType("DATABASE")
                .expressions(                
                    PermissionsLfTagPolicyExpressionArgs.builder()
                        .key("Team")
                        .values("Sales")
                        .build(),
                    PermissionsLfTagPolicyExpressionArgs.builder()
                        .key("Environment")
                        .values(                        
                            "Dev",
                            "Production")
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lakeformation:Permissions
    properties:
      principal: ${aws_iam_role.sales_role.arn}
      permissions:
        - CREATE_TABLE
        - ALTER
        - DROP
      lfTagPolicy:
        resourceType: DATABASE
        expressions:
          - key: Team
            values:
              - Sales
          - key: Environment
            values:
              - Dev
              - Production
```
{{% /example %}}
{{% /examples %}}