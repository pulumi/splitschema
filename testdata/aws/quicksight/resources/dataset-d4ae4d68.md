Resource for managing a QuickSight Data Set.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.DataSet("example", {
    dataSetId: "example-id",
    importMode: "SPICE",
    physicalTableMaps: [{
        physicalTableMapId: "example-id",
        s3Source: {
            dataSourceArn: aws_quicksight_data_source.example.arn,
            inputColumns: [{
                name: "Column1",
                type: "STRING",
            }],
            uploadSettings: {
                format: "JSON",
            },
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.DataSet("example",
    data_set_id="example-id",
    import_mode="SPICE",
    physical_table_maps=[aws.quicksight.DataSetPhysicalTableMapArgs(
        physical_table_map_id="example-id",
        s3_source=aws.quicksight.DataSetPhysicalTableMapS3SourceArgs(
            data_source_arn=aws_quicksight_data_source["example"]["arn"],
            input_columns=[aws.quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs(
                name="Column1",
                type="STRING",
            )],
            upload_settings=aws.quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs(
                format="JSON",
            ),
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.DataSet("example", new()
    {
        DataSetId = "example-id",
        ImportMode = "SPICE",
        PhysicalTableMaps = new[]
        {
            new Aws.Quicksight.Inputs.DataSetPhysicalTableMapArgs
            {
                PhysicalTableMapId = "example-id",
                S3Source = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceArgs
                {
                    DataSourceArn = aws_quicksight_data_source.Example.Arn,
                    InputColumns = new[]
                    {
                        new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceInputColumnArgs
                        {
                            Name = "Column1",
                            Type = "STRING",
                        },
                    },
                    UploadSettings = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs
                    {
                        Format = "JSON",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewDataSet(ctx, "example", &quicksight.DataSetArgs{
			DataSetId:  pulumi.String("example-id"),
			ImportMode: pulumi.String("SPICE"),
			PhysicalTableMaps: quicksight.DataSetPhysicalTableMapArray{
				&quicksight.DataSetPhysicalTableMapArgs{
					PhysicalTableMapId: pulumi.String("example-id"),
					S3Source: &quicksight.DataSetPhysicalTableMapS3SourceArgs{
						DataSourceArn: pulumi.Any(aws_quicksight_data_source.Example.Arn),
						InputColumns: quicksight.DataSetPhysicalTableMapS3SourceInputColumnArray{
							&quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs{
								Name: pulumi.String("Column1"),
								Type: pulumi.String("STRING"),
							},
						},
						UploadSettings: &quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs{
							Format: pulumi.String("JSON"),
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
import com.pulumi.aws.quicksight.DataSet;
import com.pulumi.aws.quicksight.DataSetArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
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
        var example = new DataSet("example", DataSetArgs.builder()        
            .dataSetId("example-id")
            .importMode("SPICE")
            .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
                .physicalTableMapId("example-id")
                .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
                    .dataSourceArn(aws_quicksight_data_source.example().arn())
                    .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
                        .name("Column1")
                        .type("STRING")
                        .build())
                    .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
                        .format("JSON")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:DataSet
    properties:
      dataSetId: example-id
      importMode: SPICE
      physicalTableMaps:
        - physicalTableMapId: example-id
          s3Source:
            dataSourceArn: ${aws_quicksight_data_source.example.arn}
            inputColumns:
              - name: Column1
                type: STRING
            uploadSettings:
              format: JSON
```
{{% /example %}}
{{% example %}}
### With Column Level Permission Rules

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.DataSet("example", {
    dataSetId: "example-id",
    importMode: "SPICE",
    physicalTableMaps: [{
        physicalTableMapId: "example-id",
        s3Source: {
            dataSourceArn: aws_quicksight_data_source.example.arn,
            inputColumns: [{
                name: "Column1",
                type: "STRING",
            }],
            uploadSettings: {
                format: "JSON",
            },
        },
    }],
    columnLevelPermissionRules: [{
        columnNames: ["Column1"],
        principals: [aws_quicksight_user.example.arn],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.DataSet("example",
    data_set_id="example-id",
    import_mode="SPICE",
    physical_table_maps=[aws.quicksight.DataSetPhysicalTableMapArgs(
        physical_table_map_id="example-id",
        s3_source=aws.quicksight.DataSetPhysicalTableMapS3SourceArgs(
            data_source_arn=aws_quicksight_data_source["example"]["arn"],
            input_columns=[aws.quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs(
                name="Column1",
                type="STRING",
            )],
            upload_settings=aws.quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs(
                format="JSON",
            ),
        ),
    )],
    column_level_permission_rules=[aws.quicksight.DataSetColumnLevelPermissionRuleArgs(
        column_names=["Column1"],
        principals=[aws_quicksight_user["example"]["arn"]],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.DataSet("example", new()
    {
        DataSetId = "example-id",
        ImportMode = "SPICE",
        PhysicalTableMaps = new[]
        {
            new Aws.Quicksight.Inputs.DataSetPhysicalTableMapArgs
            {
                PhysicalTableMapId = "example-id",
                S3Source = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceArgs
                {
                    DataSourceArn = aws_quicksight_data_source.Example.Arn,
                    InputColumns = new[]
                    {
                        new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceInputColumnArgs
                        {
                            Name = "Column1",
                            Type = "STRING",
                        },
                    },
                    UploadSettings = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs
                    {
                        Format = "JSON",
                    },
                },
            },
        },
        ColumnLevelPermissionRules = new[]
        {
            new Aws.Quicksight.Inputs.DataSetColumnLevelPermissionRuleArgs
            {
                ColumnNames = new[]
                {
                    "Column1",
                },
                Principals = new[]
                {
                    aws_quicksight_user.Example.Arn,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewDataSet(ctx, "example", &quicksight.DataSetArgs{
			DataSetId:  pulumi.String("example-id"),
			ImportMode: pulumi.String("SPICE"),
			PhysicalTableMaps: quicksight.DataSetPhysicalTableMapArray{
				&quicksight.DataSetPhysicalTableMapArgs{
					PhysicalTableMapId: pulumi.String("example-id"),
					S3Source: &quicksight.DataSetPhysicalTableMapS3SourceArgs{
						DataSourceArn: pulumi.Any(aws_quicksight_data_source.Example.Arn),
						InputColumns: quicksight.DataSetPhysicalTableMapS3SourceInputColumnArray{
							&quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs{
								Name: pulumi.String("Column1"),
								Type: pulumi.String("STRING"),
							},
						},
						UploadSettings: &quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs{
							Format: pulumi.String("JSON"),
						},
					},
				},
			},
			ColumnLevelPermissionRules: quicksight.DataSetColumnLevelPermissionRuleArray{
				&quicksight.DataSetColumnLevelPermissionRuleArgs{
					ColumnNames: pulumi.StringArray{
						pulumi.String("Column1"),
					},
					Principals: pulumi.StringArray{
						aws_quicksight_user.Example.Arn,
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
import com.pulumi.aws.quicksight.DataSet;
import com.pulumi.aws.quicksight.DataSetArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
import com.pulumi.aws.quicksight.inputs.DataSetColumnLevelPermissionRuleArgs;
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
        var example = new DataSet("example", DataSetArgs.builder()        
            .dataSetId("example-id")
            .importMode("SPICE")
            .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
                .physicalTableMapId("example-id")
                .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
                    .dataSourceArn(aws_quicksight_data_source.example().arn())
                    .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
                        .name("Column1")
                        .type("STRING")
                        .build())
                    .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
                        .format("JSON")
                        .build())
                    .build())
                .build())
            .columnLevelPermissionRules(DataSetColumnLevelPermissionRuleArgs.builder()
                .columnNames("Column1")
                .principals(aws_quicksight_user.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:DataSet
    properties:
      dataSetId: example-id
      importMode: SPICE
      physicalTableMaps:
        - physicalTableMapId: example-id
          s3Source:
            dataSourceArn: ${aws_quicksight_data_source.example.arn}
            inputColumns:
              - name: Column1
                type: STRING
            uploadSettings:
              format: JSON
      columnLevelPermissionRules:
        - columnNames:
            - Column1
          principals:
            - ${aws_quicksight_user.example.arn}
```
{{% /example %}}
{{% example %}}
### With Field Folders

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.DataSet("example", {
    dataSetId: "example-id",
    importMode: "SPICE",
    physicalTableMaps: [{
        physicalTableMapId: "example-id",
        s3Source: {
            dataSourceArn: aws_quicksight_data_source.example.arn,
            inputColumns: [{
                name: "Column1",
                type: "STRING",
            }],
            uploadSettings: {
                format: "JSON",
            },
        },
    }],
    fieldFolders: [{
        fieldFoldersId: "example-id",
        columns: ["Column1"],
        description: "example description",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.DataSet("example",
    data_set_id="example-id",
    import_mode="SPICE",
    physical_table_maps=[aws.quicksight.DataSetPhysicalTableMapArgs(
        physical_table_map_id="example-id",
        s3_source=aws.quicksight.DataSetPhysicalTableMapS3SourceArgs(
            data_source_arn=aws_quicksight_data_source["example"]["arn"],
            input_columns=[aws.quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs(
                name="Column1",
                type="STRING",
            )],
            upload_settings=aws.quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs(
                format="JSON",
            ),
        ),
    )],
    field_folders=[aws.quicksight.DataSetFieldFolderArgs(
        field_folders_id="example-id",
        columns=["Column1"],
        description="example description",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.DataSet("example", new()
    {
        DataSetId = "example-id",
        ImportMode = "SPICE",
        PhysicalTableMaps = new[]
        {
            new Aws.Quicksight.Inputs.DataSetPhysicalTableMapArgs
            {
                PhysicalTableMapId = "example-id",
                S3Source = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceArgs
                {
                    DataSourceArn = aws_quicksight_data_source.Example.Arn,
                    InputColumns = new[]
                    {
                        new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceInputColumnArgs
                        {
                            Name = "Column1",
                            Type = "STRING",
                        },
                    },
                    UploadSettings = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs
                    {
                        Format = "JSON",
                    },
                },
            },
        },
        FieldFolders = new[]
        {
            new Aws.Quicksight.Inputs.DataSetFieldFolderArgs
            {
                FieldFoldersId = "example-id",
                Columns = new[]
                {
                    "Column1",
                },
                Description = "example description",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewDataSet(ctx, "example", &quicksight.DataSetArgs{
			DataSetId:  pulumi.String("example-id"),
			ImportMode: pulumi.String("SPICE"),
			PhysicalTableMaps: quicksight.DataSetPhysicalTableMapArray{
				&quicksight.DataSetPhysicalTableMapArgs{
					PhysicalTableMapId: pulumi.String("example-id"),
					S3Source: &quicksight.DataSetPhysicalTableMapS3SourceArgs{
						DataSourceArn: pulumi.Any(aws_quicksight_data_source.Example.Arn),
						InputColumns: quicksight.DataSetPhysicalTableMapS3SourceInputColumnArray{
							&quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs{
								Name: pulumi.String("Column1"),
								Type: pulumi.String("STRING"),
							},
						},
						UploadSettings: &quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs{
							Format: pulumi.String("JSON"),
						},
					},
				},
			},
			FieldFolders: quicksight.DataSetFieldFolderArray{
				&quicksight.DataSetFieldFolderArgs{
					FieldFoldersId: pulumi.String("example-id"),
					Columns: pulumi.StringArray{
						pulumi.String("Column1"),
					},
					Description: pulumi.String("example description"),
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
import com.pulumi.aws.quicksight.DataSet;
import com.pulumi.aws.quicksight.DataSetArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
import com.pulumi.aws.quicksight.inputs.DataSetFieldFolderArgs;
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
        var example = new DataSet("example", DataSetArgs.builder()        
            .dataSetId("example-id")
            .importMode("SPICE")
            .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
                .physicalTableMapId("example-id")
                .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
                    .dataSourceArn(aws_quicksight_data_source.example().arn())
                    .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
                        .name("Column1")
                        .type("STRING")
                        .build())
                    .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
                        .format("JSON")
                        .build())
                    .build())
                .build())
            .fieldFolders(DataSetFieldFolderArgs.builder()
                .fieldFoldersId("example-id")
                .columns("Column1")
                .description("example description")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:DataSet
    properties:
      dataSetId: example-id
      importMode: SPICE
      physicalTableMaps:
        - physicalTableMapId: example-id
          s3Source:
            dataSourceArn: ${aws_quicksight_data_source.example.arn}
            inputColumns:
              - name: Column1
                type: STRING
            uploadSettings:
              format: JSON
      fieldFolders:
        - fieldFoldersId: example-id
          columns:
            - Column1
          description: example description
```
{{% /example %}}
{{% example %}}
### With Permissions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.DataSet("example", {
    dataSetId: "example-id",
    importMode: "SPICE",
    physicalTableMaps: [{
        physicalTableMapId: "example-id",
        s3Source: {
            dataSourceArn: aws_quicksight_data_source.example.arn,
            inputColumns: [{
                name: "Column1",
                type: "STRING",
            }],
            uploadSettings: {
                format: "JSON",
            },
        },
    }],
    permissions: [{
        actions: [
            "quicksight:DescribeDataSet",
            "quicksight:DescribeDataSetPermissions",
            "quicksight:PassDataSet",
            "quicksight:DescribeIngestion",
            "quicksight:ListIngestions",
        ],
        principal: aws_quicksight_user.example.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.DataSet("example",
    data_set_id="example-id",
    import_mode="SPICE",
    physical_table_maps=[aws.quicksight.DataSetPhysicalTableMapArgs(
        physical_table_map_id="example-id",
        s3_source=aws.quicksight.DataSetPhysicalTableMapS3SourceArgs(
            data_source_arn=aws_quicksight_data_source["example"]["arn"],
            input_columns=[aws.quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs(
                name="Column1",
                type="STRING",
            )],
            upload_settings=aws.quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs(
                format="JSON",
            ),
        ),
    )],
    permissions=[aws.quicksight.DataSetPermissionArgs(
        actions=[
            "quicksight:DescribeDataSet",
            "quicksight:DescribeDataSetPermissions",
            "quicksight:PassDataSet",
            "quicksight:DescribeIngestion",
            "quicksight:ListIngestions",
        ],
        principal=aws_quicksight_user["example"]["arn"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.DataSet("example", new()
    {
        DataSetId = "example-id",
        ImportMode = "SPICE",
        PhysicalTableMaps = new[]
        {
            new Aws.Quicksight.Inputs.DataSetPhysicalTableMapArgs
            {
                PhysicalTableMapId = "example-id",
                S3Source = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceArgs
                {
                    DataSourceArn = aws_quicksight_data_source.Example.Arn,
                    InputColumns = new[]
                    {
                        new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceInputColumnArgs
                        {
                            Name = "Column1",
                            Type = "STRING",
                        },
                    },
                    UploadSettings = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs
                    {
                        Format = "JSON",
                    },
                },
            },
        },
        Permissions = new[]
        {
            new Aws.Quicksight.Inputs.DataSetPermissionArgs
            {
                Actions = new[]
                {
                    "quicksight:DescribeDataSet",
                    "quicksight:DescribeDataSetPermissions",
                    "quicksight:PassDataSet",
                    "quicksight:DescribeIngestion",
                    "quicksight:ListIngestions",
                },
                Principal = aws_quicksight_user.Example.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewDataSet(ctx, "example", &quicksight.DataSetArgs{
			DataSetId:  pulumi.String("example-id"),
			ImportMode: pulumi.String("SPICE"),
			PhysicalTableMaps: quicksight.DataSetPhysicalTableMapArray{
				&quicksight.DataSetPhysicalTableMapArgs{
					PhysicalTableMapId: pulumi.String("example-id"),
					S3Source: &quicksight.DataSetPhysicalTableMapS3SourceArgs{
						DataSourceArn: pulumi.Any(aws_quicksight_data_source.Example.Arn),
						InputColumns: quicksight.DataSetPhysicalTableMapS3SourceInputColumnArray{
							&quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs{
								Name: pulumi.String("Column1"),
								Type: pulumi.String("STRING"),
							},
						},
						UploadSettings: &quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs{
							Format: pulumi.String("JSON"),
						},
					},
				},
			},
			Permissions: quicksight.DataSetPermissionArray{
				&quicksight.DataSetPermissionArgs{
					Actions: pulumi.StringArray{
						pulumi.String("quicksight:DescribeDataSet"),
						pulumi.String("quicksight:DescribeDataSetPermissions"),
						pulumi.String("quicksight:PassDataSet"),
						pulumi.String("quicksight:DescribeIngestion"),
						pulumi.String("quicksight:ListIngestions"),
					},
					Principal: pulumi.Any(aws_quicksight_user.Example.Arn),
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
import com.pulumi.aws.quicksight.DataSet;
import com.pulumi.aws.quicksight.DataSetArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPermissionArgs;
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
        var example = new DataSet("example", DataSetArgs.builder()        
            .dataSetId("example-id")
            .importMode("SPICE")
            .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
                .physicalTableMapId("example-id")
                .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
                    .dataSourceArn(aws_quicksight_data_source.example().arn())
                    .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
                        .name("Column1")
                        .type("STRING")
                        .build())
                    .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
                        .format("JSON")
                        .build())
                    .build())
                .build())
            .permissions(DataSetPermissionArgs.builder()
                .actions(                
                    "quicksight:DescribeDataSet",
                    "quicksight:DescribeDataSetPermissions",
                    "quicksight:PassDataSet",
                    "quicksight:DescribeIngestion",
                    "quicksight:ListIngestions")
                .principal(aws_quicksight_user.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:DataSet
    properties:
      dataSetId: example-id
      importMode: SPICE
      physicalTableMaps:
        - physicalTableMapId: example-id
          s3Source:
            dataSourceArn: ${aws_quicksight_data_source.example.arn}
            inputColumns:
              - name: Column1
                type: STRING
            uploadSettings:
              format: JSON
      permissions:
        - actions:
            - quicksight:DescribeDataSet
            - quicksight:DescribeDataSetPermissions
            - quicksight:PassDataSet
            - quicksight:DescribeIngestion
            - quicksight:ListIngestions
          principal: ${aws_quicksight_user.example.arn}
```
{{% /example %}}
{{% example %}}
### With Row Level Permission Tag Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.DataSet("example", {
    dataSetId: "example-id",
    importMode: "SPICE",
    physicalTableMaps: [{
        physicalTableMapId: "example-id",
        s3Source: {
            dataSourceArn: aws_quicksight_data_source.example.arn,
            inputColumns: [{
                name: "Column1",
                type: "STRING",
            }],
            uploadSettings: {
                format: "JSON",
            },
        },
    }],
    rowLevelPermissionTagConfiguration: {
        status: "ENABLED",
        tagRules: [{
            columnName: "Column1",
            tagKey: "tagkey",
            matchAllValue: "*",
            tagMultiValueDelimiter: ",",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.DataSet("example",
    data_set_id="example-id",
    import_mode="SPICE",
    physical_table_maps=[aws.quicksight.DataSetPhysicalTableMapArgs(
        physical_table_map_id="example-id",
        s3_source=aws.quicksight.DataSetPhysicalTableMapS3SourceArgs(
            data_source_arn=aws_quicksight_data_source["example"]["arn"],
            input_columns=[aws.quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs(
                name="Column1",
                type="STRING",
            )],
            upload_settings=aws.quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs(
                format="JSON",
            ),
        ),
    )],
    row_level_permission_tag_configuration=aws.quicksight.DataSetRowLevelPermissionTagConfigurationArgs(
        status="ENABLED",
        tag_rules=[aws.quicksight.DataSetRowLevelPermissionTagConfigurationTagRuleArgs(
            column_name="Column1",
            tag_key="tagkey",
            match_all_value="*",
            tag_multi_value_delimiter=",",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.DataSet("example", new()
    {
        DataSetId = "example-id",
        ImportMode = "SPICE",
        PhysicalTableMaps = new[]
        {
            new Aws.Quicksight.Inputs.DataSetPhysicalTableMapArgs
            {
                PhysicalTableMapId = "example-id",
                S3Source = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceArgs
                {
                    DataSourceArn = aws_quicksight_data_source.Example.Arn,
                    InputColumns = new[]
                    {
                        new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceInputColumnArgs
                        {
                            Name = "Column1",
                            Type = "STRING",
                        },
                    },
                    UploadSettings = new Aws.Quicksight.Inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs
                    {
                        Format = "JSON",
                    },
                },
            },
        },
        RowLevelPermissionTagConfiguration = new Aws.Quicksight.Inputs.DataSetRowLevelPermissionTagConfigurationArgs
        {
            Status = "ENABLED",
            TagRules = new[]
            {
                new Aws.Quicksight.Inputs.DataSetRowLevelPermissionTagConfigurationTagRuleArgs
                {
                    ColumnName = "Column1",
                    TagKey = "tagkey",
                    MatchAllValue = "*",
                    TagMultiValueDelimiter = ",",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewDataSet(ctx, "example", &quicksight.DataSetArgs{
			DataSetId:  pulumi.String("example-id"),
			ImportMode: pulumi.String("SPICE"),
			PhysicalTableMaps: quicksight.DataSetPhysicalTableMapArray{
				&quicksight.DataSetPhysicalTableMapArgs{
					PhysicalTableMapId: pulumi.String("example-id"),
					S3Source: &quicksight.DataSetPhysicalTableMapS3SourceArgs{
						DataSourceArn: pulumi.Any(aws_quicksight_data_source.Example.Arn),
						InputColumns: quicksight.DataSetPhysicalTableMapS3SourceInputColumnArray{
							&quicksight.DataSetPhysicalTableMapS3SourceInputColumnArgs{
								Name: pulumi.String("Column1"),
								Type: pulumi.String("STRING"),
							},
						},
						UploadSettings: &quicksight.DataSetPhysicalTableMapS3SourceUploadSettingsArgs{
							Format: pulumi.String("JSON"),
						},
					},
				},
			},
			RowLevelPermissionTagConfiguration: &quicksight.DataSetRowLevelPermissionTagConfigurationArgs{
				Status: pulumi.String("ENABLED"),
				TagRules: quicksight.DataSetRowLevelPermissionTagConfigurationTagRuleArray{
					&quicksight.DataSetRowLevelPermissionTagConfigurationTagRuleArgs{
						ColumnName:             pulumi.String("Column1"),
						TagKey:                 pulumi.String("tagkey"),
						MatchAllValue:          pulumi.String("*"),
						TagMultiValueDelimiter: pulumi.String(","),
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
import com.pulumi.aws.quicksight.DataSet;
import com.pulumi.aws.quicksight.DataSetArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSetPhysicalTableMapS3SourceUploadSettingsArgs;
import com.pulumi.aws.quicksight.inputs.DataSetRowLevelPermissionTagConfigurationArgs;
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
        var example = new DataSet("example", DataSetArgs.builder()        
            .dataSetId("example-id")
            .importMode("SPICE")
            .physicalTableMaps(DataSetPhysicalTableMapArgs.builder()
                .physicalTableMapId("example-id")
                .s3Source(DataSetPhysicalTableMapS3SourceArgs.builder()
                    .dataSourceArn(aws_quicksight_data_source.example().arn())
                    .inputColumns(DataSetPhysicalTableMapS3SourceInputColumnArgs.builder()
                        .name("Column1")
                        .type("STRING")
                        .build())
                    .uploadSettings(DataSetPhysicalTableMapS3SourceUploadSettingsArgs.builder()
                        .format("JSON")
                        .build())
                    .build())
                .build())
            .rowLevelPermissionTagConfiguration(DataSetRowLevelPermissionTagConfigurationArgs.builder()
                .status("ENABLED")
                .tagRules(DataSetRowLevelPermissionTagConfigurationTagRuleArgs.builder()
                    .columnName("Column1")
                    .tagKey("tagkey")
                    .matchAllValue("*")
                    .tagMultiValueDelimiter(",")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:DataSet
    properties:
      dataSetId: example-id
      importMode: SPICE
      physicalTableMaps:
        - physicalTableMapId: example-id
          s3Source:
            dataSourceArn: ${aws_quicksight_data_source.example.arn}
            inputColumns:
              - name: Column1
                type: STRING
            uploadSettings:
              format: JSON
      rowLevelPermissionTagConfiguration:
        status: ENABLED
        tagRules:
          - columnName: Column1
            tagKey: tagkey
            matchAllValue: '*'
            tagMultiValueDelimiter: ','
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a QuickSight Data Set using the AWS account ID and data set ID separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:quicksight/dataSet:DataSet example 123456789012,example-id
```
 