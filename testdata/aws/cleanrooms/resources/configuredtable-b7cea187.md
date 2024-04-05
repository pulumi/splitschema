Provides a AWS Clean Rooms configured table. Configured tables are used to represent references to existing tables in the AWS Glue Data Catalog.

{{% examples %}}
## Example Usage
{{% example %}}
### Configured table with tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testConfiguredTable = new aws.cleanrooms.ConfiguredTable("testConfiguredTable", {
    allowedColumns: [
        "column1",
        "column2",
        "column3",
    ],
    analysisMethod: "DIRECT_QUERY",
    description: "I made this table with Pulumi!",
    tableReference: {
        databaseName: "example_database",
        tableName: "example_table",
    },
    tags: {
        Project: "Pulumi",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_configured_table = aws.cleanrooms.ConfiguredTable("testConfiguredTable",
    allowed_columns=[
        "column1",
        "column2",
        "column3",
    ],
    analysis_method="DIRECT_QUERY",
    description="I made this table with Pulumi!",
    table_reference=aws.cleanrooms.ConfiguredTableTableReferenceArgs(
        database_name="example_database",
        table_name="example_table",
    ),
    tags={
        "Project": "Pulumi",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testConfiguredTable = new Aws.CleanRooms.ConfiguredTable("testConfiguredTable", new()
    {
        AllowedColumns = new[]
        {
            "column1",
            "column2",
            "column3",
        },
        AnalysisMethod = "DIRECT_QUERY",
        Description = "I made this table with Pulumi!",
        TableReference = new Aws.CleanRooms.Inputs.ConfiguredTableTableReferenceArgs
        {
            DatabaseName = "example_database",
            TableName = "example_table",
        },
        Tags = 
        {
            { "Project", "Pulumi" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cleanrooms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cleanrooms.NewConfiguredTable(ctx, "testConfiguredTable", &cleanrooms.ConfiguredTableArgs{
			AllowedColumns: pulumi.StringArray{
				pulumi.String("column1"),
				pulumi.String("column2"),
				pulumi.String("column3"),
			},
			AnalysisMethod: pulumi.String("DIRECT_QUERY"),
			Description:    pulumi.String("I made this table with Pulumi!"),
			TableReference: &cleanrooms.ConfiguredTableTableReferenceArgs{
				DatabaseName: pulumi.String("example_database"),
				TableName:    pulumi.String("example_table"),
			},
			Tags: pulumi.StringMap{
				"Project": pulumi.String("Pulumi"),
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
import com.pulumi.aws.cleanrooms.ConfiguredTable;
import com.pulumi.aws.cleanrooms.ConfiguredTableArgs;
import com.pulumi.aws.cleanrooms.inputs.ConfiguredTableTableReferenceArgs;
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
        var testConfiguredTable = new ConfiguredTable("testConfiguredTable", ConfiguredTableArgs.builder()        
            .allowedColumns(            
                "column1",
                "column2",
                "column3")
            .analysisMethod("DIRECT_QUERY")
            .description("I made this table with Pulumi!")
            .tableReference(ConfiguredTableTableReferenceArgs.builder()
                .databaseName("example_database")
                .tableName("example_table")
                .build())
            .tags(Map.of("Project", "Pulumi"))
            .build());

    }
}
```
```yaml
resources:
  testConfiguredTable:
    type: aws:cleanrooms:ConfiguredTable
    properties:
      allowedColumns:
        - column1
        - column2
        - column3
      analysisMethod: DIRECT_QUERY
      description: I made this table with Pulumi!
      tableReference:
        databaseName: example_database
        tableName: example_table
      tags:
        Project: Pulumi
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_cleanrooms_configured_table` using the `id`. For example:

```sh
 $ pulumi import aws:cleanrooms/configuredTable:ConfiguredTable table 1234abcd-12ab-34cd-56ef-1234567890ab
```
 