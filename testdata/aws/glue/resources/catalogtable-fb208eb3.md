Provides a Glue Catalog Table Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Table

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogTable = new aws.glue.CatalogTable("awsGlueCatalogTable", {
    databaseName: "MyCatalogDatabase",
    name: "MyCatalogTable",
});
```
```python
import pulumi
import pulumi_aws as aws

aws_glue_catalog_table = aws.glue.CatalogTable("awsGlueCatalogTable",
    database_name="MyCatalogDatabase",
    name="MyCatalogTable")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsGlueCatalogTable = new Aws.Glue.CatalogTable("awsGlueCatalogTable", new()
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
		_, err := glue.NewCatalogTable(ctx, "awsGlueCatalogTable", &glue.CatalogTableArgs{
			DatabaseName: pulumi.String("MyCatalogDatabase"),
			Name:         pulumi.String("MyCatalogTable"),
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
import com.pulumi.aws.glue.CatalogTable;
import com.pulumi.aws.glue.CatalogTableArgs;
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
        var awsGlueCatalogTable = new CatalogTable("awsGlueCatalogTable", CatalogTableArgs.builder()        
            .databaseName("MyCatalogDatabase")
            .name("MyCatalogTable")
            .build());

    }
}
```
```yaml
resources:
  awsGlueCatalogTable:
    type: aws:glue:CatalogTable
    properties:
      databaseName: MyCatalogDatabase
      name: MyCatalogTable
```
{{% /example %}}
{{% example %}}
### Parquet Table for Athena

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsGlueCatalogTable = new aws.glue.CatalogTable("awsGlueCatalogTable", {
    databaseName: "MyCatalogDatabase",
    name: "MyCatalogTable",
    parameters: {
        EXTERNAL: "TRUE",
        "parquet.compression": "SNAPPY",
    },
    storageDescriptor: {
        columns: [
            {
                name: "my_string",
                type: "string",
            },
            {
                name: "my_double",
                type: "double",
            },
            {
                comment: "",
                name: "my_date",
                type: "date",
            },
            {
                comment: "",
                name: "my_bigint",
                type: "bigint",
            },
            {
                comment: "",
                name: "my_struct",
                type: "struct<my_nested_string:string>",
            },
        ],
        inputFormat: "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
        location: "s3://my-bucket/event-streams/my-stream",
        outputFormat: "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
        serDeInfo: {
            name: "my-stream",
            parameters: {
                "serialization.format": "1",
            },
            serializationLibrary: "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
        },
    },
    tableType: "EXTERNAL_TABLE",
});
```
```python
import pulumi
import pulumi_aws as aws

aws_glue_catalog_table = aws.glue.CatalogTable("awsGlueCatalogTable",
    database_name="MyCatalogDatabase",
    name="MyCatalogTable",
    parameters={
        "EXTERNAL": "TRUE",
        "parquet.compression": "SNAPPY",
    },
    storage_descriptor=aws.glue.CatalogTableStorageDescriptorArgs(
        columns=[
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                name="my_string",
                type="string",
            ),
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                name="my_double",
                type="double",
            ),
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                comment="",
                name="my_date",
                type="date",
            ),
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                comment="",
                name="my_bigint",
                type="bigint",
            ),
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                comment="",
                name="my_struct",
                type="struct<my_nested_string:string>",
            ),
        ],
        input_format="org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
        location="s3://my-bucket/event-streams/my-stream",
        output_format="org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
        ser_de_info=aws.glue.CatalogTableStorageDescriptorSerDeInfoArgs(
            name="my-stream",
            parameters={
                "serialization.format": "1",
            },
            serialization_library="org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
        ),
    ),
    table_type="EXTERNAL_TABLE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsGlueCatalogTable = new Aws.Glue.CatalogTable("awsGlueCatalogTable", new()
    {
        DatabaseName = "MyCatalogDatabase",
        Name = "MyCatalogTable",
        Parameters = 
        {
            { "EXTERNAL", "TRUE" },
            { "parquet.compression", "SNAPPY" },
        },
        StorageDescriptor = new Aws.Glue.Inputs.CatalogTableStorageDescriptorArgs
        {
            Columns = new[]
            {
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Name = "my_string",
                    Type = "string",
                },
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Name = "my_double",
                    Type = "double",
                },
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Comment = "",
                    Name = "my_date",
                    Type = "date",
                },
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Comment = "",
                    Name = "my_bigint",
                    Type = "bigint",
                },
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Comment = "",
                    Name = "my_struct",
                    Type = "struct<my_nested_string:string>",
                },
            },
            InputFormat = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
            Location = "s3://my-bucket/event-streams/my-stream",
            OutputFormat = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
            SerDeInfo = new Aws.Glue.Inputs.CatalogTableStorageDescriptorSerDeInfoArgs
            {
                Name = "my-stream",
                Parameters = 
                {
                    { "serialization.format", "1" },
                },
                SerializationLibrary = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
            },
        },
        TableType = "EXTERNAL_TABLE",
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
		_, err := glue.NewCatalogTable(ctx, "awsGlueCatalogTable", &glue.CatalogTableArgs{
			DatabaseName: pulumi.String("MyCatalogDatabase"),
			Name:         pulumi.String("MyCatalogTable"),
			Parameters: pulumi.StringMap{
				"EXTERNAL":            pulumi.String("TRUE"),
				"parquet.compression": pulumi.String("SNAPPY"),
			},
			StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
				Columns: glue.CatalogTableStorageDescriptorColumnArray{
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name: pulumi.String("my_string"),
						Type: pulumi.String("string"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name: pulumi.String("my_double"),
						Type: pulumi.String("double"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Comment: pulumi.String(""),
						Name:    pulumi.String("my_date"),
						Type:    pulumi.String("date"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Comment: pulumi.String(""),
						Name:    pulumi.String("my_bigint"),
						Type:    pulumi.String("bigint"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Comment: pulumi.String(""),
						Name:    pulumi.String("my_struct"),
						Type:    pulumi.String("struct<my_nested_string:string>"),
					},
				},
				InputFormat:  pulumi.String("org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"),
				Location:     pulumi.String("s3://my-bucket/event-streams/my-stream"),
				OutputFormat: pulumi.String("org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"),
				SerDeInfo: &glue.CatalogTableStorageDescriptorSerDeInfoArgs{
					Name: pulumi.String("my-stream"),
					Parameters: pulumi.StringMap{
						"serialization.format": pulumi.String("1"),
					},
					SerializationLibrary: pulumi.String("org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"),
				},
			},
			TableType: pulumi.String("EXTERNAL_TABLE"),
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
import com.pulumi.aws.glue.CatalogTable;
import com.pulumi.aws.glue.CatalogTableArgs;
import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorArgs;
import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorSerDeInfoArgs;
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
        var awsGlueCatalogTable = new CatalogTable("awsGlueCatalogTable", CatalogTableArgs.builder()        
            .databaseName("MyCatalogDatabase")
            .name("MyCatalogTable")
            .parameters(Map.ofEntries(
                Map.entry("EXTERNAL", "TRUE"),
                Map.entry("parquet.compression", "SNAPPY")
            ))
            .storageDescriptor(CatalogTableStorageDescriptorArgs.builder()
                .columns(                
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .name("my_string")
                        .type("string")
                        .build(),
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .name("my_double")
                        .type("double")
                        .build(),
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .comment("")
                        .name("my_date")
                        .type("date")
                        .build(),
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .comment("")
                        .name("my_bigint")
                        .type("bigint")
                        .build(),
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .comment("")
                        .name("my_struct")
                        .type("struct<my_nested_string:string>")
                        .build())
                .inputFormat("org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat")
                .location("s3://my-bucket/event-streams/my-stream")
                .outputFormat("org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat")
                .serDeInfo(CatalogTableStorageDescriptorSerDeInfoArgs.builder()
                    .name("my-stream")
                    .parameters(Map.of("serialization.format", 1))
                    .serializationLibrary("org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe")
                    .build())
                .build())
            .tableType("EXTERNAL_TABLE")
            .build());

    }
}
```
```yaml
resources:
  awsGlueCatalogTable:
    type: aws:glue:CatalogTable
    properties:
      databaseName: MyCatalogDatabase
      name: MyCatalogTable
      parameters:
        EXTERNAL: TRUE
        parquet.compression: SNAPPY
      storageDescriptor:
        columns:
          - name: my_string
            type: string
          - name: my_double
            type: double
          - comment:
            name: my_date
            type: date
          - comment:
            name: my_bigint
            type: bigint
          - comment:
            name: my_struct
            type: struct<my_nested_string:string>
        inputFormat: org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat
        location: s3://my-bucket/event-streams/my-stream
        outputFormat: org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat
        serDeInfo:
          name: my-stream
          parameters:
            serialization.format: 1
          serializationLibrary: org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
      tableType: EXTERNAL_TABLE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Tables using the catalog ID (usually AWS account ID), database name, and table name. For example:

```sh
 $ pulumi import aws:glue/catalogTable:CatalogTable MyTable 123456789012:MyDatabase:MyTable
```
 