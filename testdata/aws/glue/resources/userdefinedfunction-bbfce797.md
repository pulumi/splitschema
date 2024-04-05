Provides a Glue User Defined Function Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCatalogDatabase = new aws.glue.CatalogDatabase("exampleCatalogDatabase", {name: "my_database"});
const exampleUserDefinedFunction = new aws.glue.UserDefinedFunction("exampleUserDefinedFunction", {
    catalogId: exampleCatalogDatabase.catalogId,
    databaseName: exampleCatalogDatabase.name,
    className: "class",
    ownerName: "owner",
    ownerType: "GROUP",
    resourceUris: [{
        resourceType: "ARCHIVE",
        uri: "uri",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_catalog_database = aws.glue.CatalogDatabase("exampleCatalogDatabase", name="my_database")
example_user_defined_function = aws.glue.UserDefinedFunction("exampleUserDefinedFunction",
    catalog_id=example_catalog_database.catalog_id,
    database_name=example_catalog_database.name,
    class_name="class",
    owner_name="owner",
    owner_type="GROUP",
    resource_uris=[aws.glue.UserDefinedFunctionResourceUriArgs(
        resource_type="ARCHIVE",
        uri="uri",
    )])
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
        Name = "my_database",
    });

    var exampleUserDefinedFunction = new Aws.Glue.UserDefinedFunction("exampleUserDefinedFunction", new()
    {
        CatalogId = exampleCatalogDatabase.CatalogId,
        DatabaseName = exampleCatalogDatabase.Name,
        ClassName = "class",
        OwnerName = "owner",
        OwnerType = "GROUP",
        ResourceUris = new[]
        {
            new Aws.Glue.Inputs.UserDefinedFunctionResourceUriArgs
            {
                ResourceType = "ARCHIVE",
                Uri = "uri",
            },
        },
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
		exampleCatalogDatabase, err := glue.NewCatalogDatabase(ctx, "exampleCatalogDatabase", &glue.CatalogDatabaseArgs{
			Name: pulumi.String("my_database"),
		})
		if err != nil {
			return err
		}
		_, err = glue.NewUserDefinedFunction(ctx, "exampleUserDefinedFunction", &glue.UserDefinedFunctionArgs{
			CatalogId:    exampleCatalogDatabase.CatalogId,
			DatabaseName: exampleCatalogDatabase.Name,
			ClassName:    pulumi.String("class"),
			OwnerName:    pulumi.String("owner"),
			OwnerType:    pulumi.String("GROUP"),
			ResourceUris: glue.UserDefinedFunctionResourceUriArray{
				&glue.UserDefinedFunctionResourceUriArgs{
					ResourceType: pulumi.String("ARCHIVE"),
					Uri:          pulumi.String("uri"),
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
import com.pulumi.aws.glue.UserDefinedFunction;
import com.pulumi.aws.glue.UserDefinedFunctionArgs;
import com.pulumi.aws.glue.inputs.UserDefinedFunctionResourceUriArgs;
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
            .name("my_database")
            .build());

        var exampleUserDefinedFunction = new UserDefinedFunction("exampleUserDefinedFunction", UserDefinedFunctionArgs.builder()        
            .catalogId(exampleCatalogDatabase.catalogId())
            .databaseName(exampleCatalogDatabase.name())
            .className("class")
            .ownerName("owner")
            .ownerType("GROUP")
            .resourceUris(UserDefinedFunctionResourceUriArgs.builder()
                .resourceType("ARCHIVE")
                .uri("uri")
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
      name: my_database
  exampleUserDefinedFunction:
    type: aws:glue:UserDefinedFunction
    properties:
      catalogId: ${exampleCatalogDatabase.catalogId}
      databaseName: ${exampleCatalogDatabase.name}
      className: class
      ownerName: owner
      ownerType: GROUP
      resourceUris:
        - resourceType: ARCHIVE
          uri: uri
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue User Defined Functions using the `catalog_id:database_name:function_name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in. For example:

```sh
 $ pulumi import aws:glue/userDefinedFunction:UserDefinedFunction func 123456789012:my_database:my_func
```
 