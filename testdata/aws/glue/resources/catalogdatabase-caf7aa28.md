Provides a Glue Catalog Database Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogDatabase = new aws.glue.CatalogDatabase("awsGlueCatalogDatabase", {name: "MyCatalogDatabase"});
```
```python
import pulumi
import pulumi_aws as aws

aws_glue_catalog_database = aws.glue.CatalogDatabase("awsGlueCatalogDatabase", name="MyCatalogDatabase")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsGlueCatalogDatabase = new Aws.Glue.CatalogDatabase("awsGlueCatalogDatabase", new()
    {
        Name = "MyCatalogDatabase",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCatalogDatabase(ctx, "awsGlueCatalogDatabase", &glue.CatalogDatabaseArgs{
			Name: pulumi.String("MyCatalogDatabase"),
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
        var awsGlueCatalogDatabase = new CatalogDatabase("awsGlueCatalogDatabase", CatalogDatabaseArgs.builder()        
            .name("MyCatalogDatabase")
            .build());

    }
}
```
```yaml
resources:
  awsGlueCatalogDatabase:
    type: aws:glue:CatalogDatabase
    properties:
      name: MyCatalogDatabase
```
{{% /example %}}
{{% example %}}
### Create Table Default Permissions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogDatabase = new aws.glue.CatalogDatabase("awsGlueCatalogDatabase", {
    createTableDefaultPermissions: [{
        permissions: ["SELECT"],
        principal: {
            dataLakePrincipalIdentifier: "IAM_ALLOWED_PRINCIPALS",
        },
    }],
    name: "MyCatalogDatabase",
});
```
```python
import pulumi
import pulumi_aws as aws

aws_glue_catalog_database = aws.glue.CatalogDatabase("awsGlueCatalogDatabase",
    create_table_default_permissions=[aws.glue.CatalogDatabaseCreateTableDefaultPermissionArgs(
        permissions=["SELECT"],
        principal=aws.glue.CatalogDatabaseCreateTableDefaultPermissionPrincipalArgs(
            data_lake_principal_identifier="IAM_ALLOWED_PRINCIPALS",
        ),
    )],
    name="MyCatalogDatabase")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsGlueCatalogDatabase = new Aws.Glue.CatalogDatabase("awsGlueCatalogDatabase", new()
    {
        CreateTableDefaultPermissions = new[]
        {
            new Aws.Glue.Inputs.CatalogDatabaseCreateTableDefaultPermissionArgs
            {
                Permissions = new[]
                {
                    "SELECT",
                },
                Principal = new Aws.Glue.Inputs.CatalogDatabaseCreateTableDefaultPermissionPrincipalArgs
                {
                    DataLakePrincipalIdentifier = "IAM_ALLOWED_PRINCIPALS",
                },
            },
        },
        Name = "MyCatalogDatabase",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewCatalogDatabase(ctx, "awsGlueCatalogDatabase", &glue.CatalogDatabaseArgs{
			CreateTableDefaultPermissions: glue.CatalogDatabaseCreateTableDefaultPermissionArray{
				&glue.CatalogDatabaseCreateTableDefaultPermissionArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("SELECT"),
					},
					Principal: &glue.CatalogDatabaseCreateTableDefaultPermissionPrincipalArgs{
						DataLakePrincipalIdentifier: pulumi.String("IAM_ALLOWED_PRINCIPALS"),
					},
				},
			},
			Name: pulumi.String("MyCatalogDatabase"),
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
import com.pulumi.aws.glue.inputs.CatalogDatabaseCreateTableDefaultPermissionArgs;
import com.pulumi.aws.glue.inputs.CatalogDatabaseCreateTableDefaultPermissionPrincipalArgs;
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
        var awsGlueCatalogDatabase = new CatalogDatabase("awsGlueCatalogDatabase", CatalogDatabaseArgs.builder()        
            .createTableDefaultPermissions(CatalogDatabaseCreateTableDefaultPermissionArgs.builder()
                .permissions("SELECT")
                .principal(CatalogDatabaseCreateTableDefaultPermissionPrincipalArgs.builder()
                    .dataLakePrincipalIdentifier("IAM_ALLOWED_PRINCIPALS")
                    .build())
                .build())
            .name("MyCatalogDatabase")
            .build());

    }
}
```
```yaml
resources:
  awsGlueCatalogDatabase:
    type: aws:glue:CatalogDatabase
    properties:
      createTableDefaultPermissions:
        - permissions:
            - SELECT
          principal:
            dataLakePrincipalIdentifier: IAM_ALLOWED_PRINCIPALS
      name: MyCatalogDatabase
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Catalog Databases using the `catalog_id:name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in. For example:

```sh
 $ pulumi import aws:glue/catalogDatabase:CatalogDatabase database 123456789012:my_database
```
 