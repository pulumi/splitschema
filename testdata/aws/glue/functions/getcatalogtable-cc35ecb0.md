This data source can be used to fetch information about an AWS Glue Data Catalog Table.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.glue.getCatalogTable({
    databaseName: "MyCatalogDatabase",
    name: "MyCatalogTable",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.get_catalog_table(database_name="MyCatalogDatabase",
    name="MyCatalogTable")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Glue.GetCatalogTable.Invoke(new()
    {
        DatabaseName = "MyCatalogDatabase",
        Name = "MyCatalogTable",
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
		_, err := glue.LookupCatalogTable(ctx, &glue.LookupCatalogTableArgs{
			DatabaseName: "MyCatalogDatabase",
			Name:         "MyCatalogTable",
		}, nil)
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
import com.pulumi.aws.glue.GlueFunctions;
import com.pulumi.aws.glue.inputs.GetCatalogTableArgs;
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
        final var example = GlueFunctions.getCatalogTable(GetCatalogTableArgs.builder()
            .databaseName("MyCatalogDatabase")
            .name("MyCatalogTable")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:glue:getCatalogTable
      Arguments:
        databaseName: MyCatalogDatabase
        name: MyCatalogTable
```
{{% /example %}}
{{% /examples %}}