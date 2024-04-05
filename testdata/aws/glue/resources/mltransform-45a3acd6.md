Provides a Glue ML Transform resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCatalogDatabase = new aws.glue.CatalogDatabase("testCatalogDatabase", {name: "example"});
const testCatalogTable = new aws.glue.CatalogTable("testCatalogTable", {
    name: "example",
    databaseName: testCatalogDatabase.name,
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
const testMLTransform = new aws.glue.MLTransform("testMLTransform", {
    roleArn: aws_iam_role.test.arn,
    inputRecordTables: [{
        databaseName: testCatalogTable.databaseName,
        tableName: testCatalogTable.name,
    }],
    parameters: {
        transformType: "FIND_MATCHES",
        findMatchesParameters: {
            primaryKeyColumnName: "my_column_1",
        },
    },
}, {
    dependsOn: [aws_iam_role_policy_attachment.test],
});
```
```python
import pulumi
import pulumi_aws as aws

test_catalog_database = aws.glue.CatalogDatabase("testCatalogDatabase", name="example")
test_catalog_table = aws.glue.CatalogTable("testCatalogTable",
    name="example",
    database_name=test_catalog_database.name,
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
test_ml_transform = aws.glue.MLTransform("testMLTransform",
    role_arn=aws_iam_role["test"]["arn"],
    input_record_tables=[aws.glue.MLTransformInputRecordTableArgs(
        database_name=test_catalog_table.database_name,
        table_name=test_catalog_table.name,
    )],
    parameters=aws.glue.MLTransformParametersArgs(
        transform_type="FIND_MATCHES",
        find_matches_parameters=aws.glue.MLTransformParametersFindMatchesParametersArgs(
            primary_key_column_name="my_column_1",
        ),
    ),
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy_attachment["test"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCatalogDatabase = new Aws.Glue.CatalogDatabase("testCatalogDatabase", new()
    {
        Name = "example",
    });

    var testCatalogTable = new Aws.Glue.CatalogTable("testCatalogTable", new()
    {
        Name = "example",
        DatabaseName = testCatalogDatabase.Name,
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

    var testMLTransform = new Aws.Glue.MLTransform("testMLTransform", new()
    {
        RoleArn = aws_iam_role.Test.Arn,
        InputRecordTables = new[]
        {
            new Aws.Glue.Inputs.MLTransformInputRecordTableArgs
            {
                DatabaseName = testCatalogTable.DatabaseName,
                TableName = testCatalogTable.Name,
            },
        },
        Parameters = new Aws.Glue.Inputs.MLTransformParametersArgs
        {
            TransformType = "FIND_MATCHES",
            FindMatchesParameters = new Aws.Glue.Inputs.MLTransformParametersFindMatchesParametersArgs
            {
                PrimaryKeyColumnName = "my_column_1",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy_attachment.Test,
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
		testCatalogDatabase, err := glue.NewCatalogDatabase(ctx, "testCatalogDatabase", &glue.CatalogDatabaseArgs{
			Name: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		testCatalogTable, err := glue.NewCatalogTable(ctx, "testCatalogTable", &glue.CatalogTableArgs{
			Name:             pulumi.String("example"),
			DatabaseName:     testCatalogDatabase.Name,
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
		_, err = glue.NewMLTransform(ctx, "testMLTransform", &glue.MLTransformArgs{
			RoleArn: pulumi.Any(aws_iam_role.Test.Arn),
			InputRecordTables: glue.MLTransformInputRecordTableArray{
				&glue.MLTransformInputRecordTableArgs{
					DatabaseName: testCatalogTable.DatabaseName,
					TableName:    testCatalogTable.Name,
				},
			},
			Parameters: &glue.MLTransformParametersArgs{
				TransformType: pulumi.String("FIND_MATCHES"),
				FindMatchesParameters: &glue.MLTransformParametersFindMatchesParametersArgs{
					PrimaryKeyColumnName: pulumi.String("my_column_1"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy_attachment.Test,
		}))
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
import com.pulumi.aws.glue.MLTransform;
import com.pulumi.aws.glue.MLTransformArgs;
import com.pulumi.aws.glue.inputs.MLTransformInputRecordTableArgs;
import com.pulumi.aws.glue.inputs.MLTransformParametersArgs;
import com.pulumi.aws.glue.inputs.MLTransformParametersFindMatchesParametersArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var testCatalogDatabase = new CatalogDatabase("testCatalogDatabase", CatalogDatabaseArgs.builder()        
            .name("example")
            .build());

        var testCatalogTable = new CatalogTable("testCatalogTable", CatalogTableArgs.builder()        
            .name("example")
            .databaseName(testCatalogDatabase.name())
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

        var testMLTransform = new MLTransform("testMLTransform", MLTransformArgs.builder()        
            .roleArn(aws_iam_role.test().arn())
            .inputRecordTables(MLTransformInputRecordTableArgs.builder()
                .databaseName(testCatalogTable.databaseName())
                .tableName(testCatalogTable.name())
                .build())
            .parameters(MLTransformParametersArgs.builder()
                .transformType("FIND_MATCHES")
                .findMatchesParameters(MLTransformParametersFindMatchesParametersArgs.builder()
                    .primaryKeyColumnName("my_column_1")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy_attachment.test())
                .build());

    }
}
```
```yaml
resources:
  testMLTransform:
    type: aws:glue:MLTransform
    properties:
      roleArn: ${aws_iam_role.test.arn}
      inputRecordTables:
        - databaseName: ${testCatalogTable.databaseName}
          tableName: ${testCatalogTable.name}
      parameters:
        transformType: FIND_MATCHES
        findMatchesParameters:
          primaryKeyColumnName: my_column_1
    options:
      dependson:
        - ${aws_iam_role_policy_attachment.test}
  testCatalogDatabase:
    type: aws:glue:CatalogDatabase
    properties:
      name: example
  testCatalogTable:
    type: aws:glue:CatalogTable
    properties:
      name: example
      databaseName: ${testCatalogDatabase.name}
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
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue ML Transforms using `id`. For example:

```sh
 $ pulumi import aws:glue/mLTransform:MLTransform example tfm-c2cafbe83b1c575f49eaca9939220e2fcd58e2d5
```
 