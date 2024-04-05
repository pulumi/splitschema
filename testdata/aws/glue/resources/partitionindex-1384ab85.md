{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCatalogDatabase = new aws.glue.CatalogDatabase("exampleCatalogDatabase", {name: "example"});
const exampleCatalogTable = new aws.glue.CatalogTable("exampleCatalogTable", {
    name: "example",
    databaseName: exampleCatalogDatabase.name,
    owner: "my_owner",
    retention: 1,
    tableType: "VIRTUAL_VIEW",
    viewExpandedText: "view_expanded_text_1",
    viewOriginalText: "view_original_text_1",
    storageDescriptor: {
        bucketColumns: ["bucket_column_1"],
        compressed: false,
        inputFormat: "SequenceFileInputFormat",
        location: "my_location",
        numberOfBuckets: 1,
        outputFormat: "SequenceFileInputFormat",
        storedAsSubDirectories: false,
        parameters: {
            param1: "param1_val",
        },
        columns: [
            {
                name: "my_column_1",
                type: "int",
                comment: "my_column1_comment",
            },
            {
                name: "my_column_2",
                type: "string",
                comment: "my_column2_comment",
            },
        ],
        serDeInfo: {
            name: "ser_de_name",
            parameters: {
                param1: "param_val_1",
            },
            serializationLibrary: "org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
        },
        sortColumns: [{
            column: "my_column_1",
            sortOrder: 1,
        }],
        skewedInfo: {
            skewedColumnNames: ["my_column_1"],
            skewedColumnValueLocationMaps: {
                my_column_1: "my_column_1_val_loc_map",
            },
            skewedColumnValues: ["skewed_val_1"],
        },
    },
    partitionKeys: [
        {
            name: "my_column_1",
            type: "int",
            comment: "my_column_1_comment",
        },
        {
            name: "my_column_2",
            type: "string",
            comment: "my_column_2_comment",
        },
    ],
    parameters: {
        param1: "param1_val",
    },
});
const examplePartitionIndex = new aws.glue.PartitionIndex("examplePartitionIndex", {
    databaseName: exampleCatalogDatabase.name,
    tableName: exampleCatalogTable.name,
    partitionIndex: {
        indexName: "example",
        keys: [
            "my_column_1",
            "my_column_2",
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_catalog_database = aws.glue.CatalogDatabase("exampleCatalogDatabase", name="example")
example_catalog_table = aws.glue.CatalogTable("exampleCatalogTable",
    name="example",
    database_name=example_catalog_database.name,
    owner="my_owner",
    retention=1,
    table_type="VIRTUAL_VIEW",
    view_expanded_text="view_expanded_text_1",
    view_original_text="view_original_text_1",
    storage_descriptor=aws.glue.CatalogTableStorageDescriptorArgs(
        bucket_columns=["bucket_column_1"],
        compressed=False,
        input_format="SequenceFileInputFormat",
        location="my_location",
        number_of_buckets=1,
        output_format="SequenceFileInputFormat",
        stored_as_sub_directories=False,
        parameters={
            "param1": "param1_val",
        },
        columns=[
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                name="my_column_1",
                type="int",
                comment="my_column1_comment",
            ),
            aws.glue.CatalogTableStorageDescriptorColumnArgs(
                name="my_column_2",
                type="string",
                comment="my_column2_comment",
            ),
        ],
        ser_de_info=aws.glue.CatalogTableStorageDescriptorSerDeInfoArgs(
            name="ser_de_name",
            parameters={
                "param1": "param_val_1",
            },
            serialization_library="org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
        ),
        sort_columns=[aws.glue.CatalogTableStorageDescriptorSortColumnArgs(
            column="my_column_1",
            sort_order=1,
        )],
        skewed_info=aws.glue.CatalogTableStorageDescriptorSkewedInfoArgs(
            skewed_column_names=["my_column_1"],
            skewed_column_value_location_maps={
                "my_column_1": "my_column_1_val_loc_map",
            },
            skewed_column_values=["skewed_val_1"],
        ),
    ),
    partition_keys=[
        aws.glue.CatalogTablePartitionKeyArgs(
            name="my_column_1",
            type="int",
            comment="my_column_1_comment",
        ),
        aws.glue.CatalogTablePartitionKeyArgs(
            name="my_column_2",
            type="string",
            comment="my_column_2_comment",
        ),
    ],
    parameters={
        "param1": "param1_val",
    })
example_partition_index = aws.glue.PartitionIndex("examplePartitionIndex",
    database_name=example_catalog_database.name,
    table_name=example_catalog_table.name,
    partition_index=aws.glue.PartitionIndexPartitionIndexArgs(
        index_name="example",
        keys=[
            "my_column_1",
            "my_column_2",
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
    var exampleCatalogDatabase = new Aws.Glue.CatalogDatabase("exampleCatalogDatabase", new()
    {
        Name = "example",
    });

    var exampleCatalogTable = new Aws.Glue.CatalogTable("exampleCatalogTable", new()
    {
        Name = "example",
        DatabaseName = exampleCatalogDatabase.Name,
        Owner = "my_owner",
        Retention = 1,
        TableType = "VIRTUAL_VIEW",
        ViewExpandedText = "view_expanded_text_1",
        ViewOriginalText = "view_original_text_1",
        StorageDescriptor = new Aws.Glue.Inputs.CatalogTableStorageDescriptorArgs
        {
            BucketColumns = new[]
            {
                "bucket_column_1",
            },
            Compressed = false,
            InputFormat = "SequenceFileInputFormat",
            Location = "my_location",
            NumberOfBuckets = 1,
            OutputFormat = "SequenceFileInputFormat",
            StoredAsSubDirectories = false,
            Parameters = 
            {
                { "param1", "param1_val" },
            },
            Columns = new[]
            {
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Name = "my_column_1",
                    Type = "int",
                    Comment = "my_column1_comment",
                },
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorColumnArgs
                {
                    Name = "my_column_2",
                    Type = "string",
                    Comment = "my_column2_comment",
                },
            },
            SerDeInfo = new Aws.Glue.Inputs.CatalogTableStorageDescriptorSerDeInfoArgs
            {
                Name = "ser_de_name",
                Parameters = 
                {
                    { "param1", "param_val_1" },
                },
                SerializationLibrary = "org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
            },
            SortColumns = new[]
            {
                new Aws.Glue.Inputs.CatalogTableStorageDescriptorSortColumnArgs
                {
                    Column = "my_column_1",
                    SortOrder = 1,
                },
            },
            SkewedInfo = new Aws.Glue.Inputs.CatalogTableStorageDescriptorSkewedInfoArgs
            {
                SkewedColumnNames = new[]
                {
                    "my_column_1",
                },
                SkewedColumnValueLocationMaps = 
                {
                    { "my_column_1", "my_column_1_val_loc_map" },
                },
                SkewedColumnValues = new[]
                {
                    "skewed_val_1",
                },
            },
        },
        PartitionKeys = new[]
        {
            new Aws.Glue.Inputs.CatalogTablePartitionKeyArgs
            {
                Name = "my_column_1",
                Type = "int",
                Comment = "my_column_1_comment",
            },
            new Aws.Glue.Inputs.CatalogTablePartitionKeyArgs
            {
                Name = "my_column_2",
                Type = "string",
                Comment = "my_column_2_comment",
            },
        },
        Parameters = 
        {
            { "param1", "param1_val" },
        },
    });

    var examplePartitionIndex = new Aws.Glue.PartitionIndex("examplePartitionIndex", new()
    {
        DatabaseName = exampleCatalogDatabase.Name,
        TableName = exampleCatalogTable.Name,
        PartitionIndexConfig = new Aws.Glue.Inputs.PartitionIndexPartitionIndexArgs
        {
            IndexName = "example",
            Keys = new[]
            {
                "my_column_1",
                "my_column_2",
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
			Name: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleCatalogTable, err := glue.NewCatalogTable(ctx, "exampleCatalogTable", &glue.CatalogTableArgs{
			Name:             pulumi.String("example"),
			DatabaseName:     exampleCatalogDatabase.Name,
			Owner:            pulumi.String("my_owner"),
			Retention:        pulumi.Int(1),
			TableType:        pulumi.String("VIRTUAL_VIEW"),
			ViewExpandedText: pulumi.String("view_expanded_text_1"),
			ViewOriginalText: pulumi.String("view_original_text_1"),
			StorageDescriptor: &glue.CatalogTableStorageDescriptorArgs{
				BucketColumns: pulumi.StringArray{
					pulumi.String("bucket_column_1"),
				},
				Compressed:             pulumi.Bool(false),
				InputFormat:            pulumi.String("SequenceFileInputFormat"),
				Location:               pulumi.String("my_location"),
				NumberOfBuckets:        pulumi.Int(1),
				OutputFormat:           pulumi.String("SequenceFileInputFormat"),
				StoredAsSubDirectories: pulumi.Bool(false),
				Parameters: pulumi.StringMap{
					"param1": pulumi.String("param1_val"),
				},
				Columns: glue.CatalogTableStorageDescriptorColumnArray{
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name:    pulumi.String("my_column_1"),
						Type:    pulumi.String("int"),
						Comment: pulumi.String("my_column1_comment"),
					},
					&glue.CatalogTableStorageDescriptorColumnArgs{
						Name:    pulumi.String("my_column_2"),
						Type:    pulumi.String("string"),
						Comment: pulumi.String("my_column2_comment"),
					},
				},
				SerDeInfo: &glue.CatalogTableStorageDescriptorSerDeInfoArgs{
					Name: pulumi.String("ser_de_name"),
					Parameters: pulumi.StringMap{
						"param1": pulumi.String("param_val_1"),
					},
					SerializationLibrary: pulumi.String("org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe"),
				},
				SortColumns: glue.CatalogTableStorageDescriptorSortColumnArray{
					&glue.CatalogTableStorageDescriptorSortColumnArgs{
						Column:    pulumi.String("my_column_1"),
						SortOrder: pulumi.Int(1),
					},
				},
				SkewedInfo: &glue.CatalogTableStorageDescriptorSkewedInfoArgs{
					SkewedColumnNames: pulumi.StringArray{
						pulumi.String("my_column_1"),
					},
					SkewedColumnValueLocationMaps: pulumi.StringMap{
						"my_column_1": pulumi.String("my_column_1_val_loc_map"),
					},
					SkewedColumnValues: pulumi.StringArray{
						pulumi.String("skewed_val_1"),
					},
				},
			},
			PartitionKeys: glue.CatalogTablePartitionKeyArray{
				&glue.CatalogTablePartitionKeyArgs{
					Name:    pulumi.String("my_column_1"),
					Type:    pulumi.String("int"),
					Comment: pulumi.String("my_column_1_comment"),
				},
				&glue.CatalogTablePartitionKeyArgs{
					Name:    pulumi.String("my_column_2"),
					Type:    pulumi.String("string"),
					Comment: pulumi.String("my_column_2_comment"),
				},
			},
			Parameters: pulumi.StringMap{
				"param1": pulumi.String("param1_val"),
			},
		})
		if err != nil {
			return err
		}
		_, err = glue.NewPartitionIndex(ctx, "examplePartitionIndex", &glue.PartitionIndexArgs{
			DatabaseName: exampleCatalogDatabase.Name,
			TableName:    exampleCatalogTable.Name,
			PartitionIndex: &glue.PartitionIndexPartitionIndexArgs{
				IndexName: pulumi.String("example"),
				Keys: pulumi.StringArray{
					pulumi.String("my_column_1"),
					pulumi.String("my_column_2"),
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
import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorSerDeInfoArgs;
import com.pulumi.aws.glue.inputs.CatalogTableStorageDescriptorSkewedInfoArgs;
import com.pulumi.aws.glue.inputs.CatalogTablePartitionKeyArgs;
import com.pulumi.aws.glue.PartitionIndex;
import com.pulumi.aws.glue.PartitionIndexArgs;
import com.pulumi.aws.glue.inputs.PartitionIndexPartitionIndexArgs;
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
            .name("example")
            .build());

        var exampleCatalogTable = new CatalogTable("exampleCatalogTable", CatalogTableArgs.builder()        
            .name("example")
            .databaseName(exampleCatalogDatabase.name())
            .owner("my_owner")
            .retention(1)
            .tableType("VIRTUAL_VIEW")
            .viewExpandedText("view_expanded_text_1")
            .viewOriginalText("view_original_text_1")
            .storageDescriptor(CatalogTableStorageDescriptorArgs.builder()
                .bucketColumns("bucket_column_1")
                .compressed(false)
                .inputFormat("SequenceFileInputFormat")
                .location("my_location")
                .numberOfBuckets(1)
                .outputFormat("SequenceFileInputFormat")
                .storedAsSubDirectories(false)
                .parameters(Map.of("param1", "param1_val"))
                .columns(                
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .name("my_column_1")
                        .type("int")
                        .comment("my_column1_comment")
                        .build(),
                    CatalogTableStorageDescriptorColumnArgs.builder()
                        .name("my_column_2")
                        .type("string")
                        .comment("my_column2_comment")
                        .build())
                .serDeInfo(CatalogTableStorageDescriptorSerDeInfoArgs.builder()
                    .name("ser_de_name")
                    .parameters(Map.of("param1", "param_val_1"))
                    .serializationLibrary("org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe")
                    .build())
                .sortColumns(CatalogTableStorageDescriptorSortColumnArgs.builder()
                    .column("my_column_1")
                    .sortOrder(1)
                    .build())
                .skewedInfo(CatalogTableStorageDescriptorSkewedInfoArgs.builder()
                    .skewedColumnNames("my_column_1")
                    .skewedColumnValueLocationMaps(Map.of("my_column_1", "my_column_1_val_loc_map"))
                    .skewedColumnValues("skewed_val_1")
                    .build())
                .build())
            .partitionKeys(            
                CatalogTablePartitionKeyArgs.builder()
                    .name("my_column_1")
                    .type("int")
                    .comment("my_column_1_comment")
                    .build(),
                CatalogTablePartitionKeyArgs.builder()
                    .name("my_column_2")
                    .type("string")
                    .comment("my_column_2_comment")
                    .build())
            .parameters(Map.of("param1", "param1_val"))
            .build());

        var examplePartitionIndex = new PartitionIndex("examplePartitionIndex", PartitionIndexArgs.builder()        
            .databaseName(exampleCatalogDatabase.name())
            .tableName(exampleCatalogTable.name())
            .partitionIndex(PartitionIndexPartitionIndexArgs.builder()
                .indexName("example")
                .keys(                
                    "my_column_1",
                    "my_column_2")
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
      name: example
  exampleCatalogTable:
    type: aws:glue:CatalogTable
    properties:
      name: example
      databaseName: ${exampleCatalogDatabase.name}
      owner: my_owner
      retention: 1
      tableType: VIRTUAL_VIEW
      viewExpandedText: view_expanded_text_1
      viewOriginalText: view_original_text_1
      storageDescriptor:
        bucketColumns:
          - bucket_column_1
        compressed: false
        inputFormat: SequenceFileInputFormat
        location: my_location
        numberOfBuckets: 1
        outputFormat: SequenceFileInputFormat
        storedAsSubDirectories: false
        parameters:
          param1: param1_val
        columns:
          - name: my_column_1
            type: int
            comment: my_column1_comment
          - name: my_column_2
            type: string
            comment: my_column2_comment
        serDeInfo:
          name: ser_de_name
          parameters:
            param1: param_val_1
          serializationLibrary: org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe
        sortColumns:
          - column: my_column_1
            sortOrder: 1
        skewedInfo:
          skewedColumnNames:
            - my_column_1
          skewedColumnValueLocationMaps:
            my_column_1: my_column_1_val_loc_map
          skewedColumnValues:
            - skewed_val_1
      partitionKeys:
        - name: my_column_1
          type: int
          comment: my_column_1_comment
        - name: my_column_2
          type: string
          comment: my_column_2_comment
      parameters:
        param1: param1_val
  examplePartitionIndex:
    type: aws:glue:PartitionIndex
    properties:
      databaseName: ${exampleCatalogDatabase.name}
      tableName: ${exampleCatalogTable.name}
      partitionIndex:
        indexName: example
        keys:
          - my_column_1
          - my_column_2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Partition Indexes using the catalog ID (usually AWS account ID), database name, table name, and index name. For example:

```sh
 $ pulumi import aws:glue/partitionIndex:PartitionIndex example 123456789012:MyDatabase:MyTable:index-name
```
 