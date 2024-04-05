Provides a Timestream table resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.timestreamwrite.Table("example", {
    databaseName: aws_timestreamwrite_database.example.database_name,
    tableName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.timestreamwrite.Table("example",
    database_name=aws_timestreamwrite_database["example"]["database_name"],
    table_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.TimestreamWrite.Table("example", new()
    {
        DatabaseName = aws_timestreamwrite_database.Example.Database_name,
        TableName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timestreamwrite.NewTable(ctx, "example", &timestreamwrite.TableArgs{
			DatabaseName: pulumi.Any(aws_timestreamwrite_database.Example.Database_name),
			TableName:    pulumi.String("example"),
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
import com.pulumi.aws.timestreamwrite.Table;
import com.pulumi.aws.timestreamwrite.TableArgs;
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
        var example = new Table("example", TableArgs.builder()        
            .databaseName(aws_timestreamwrite_database.example().database_name())
            .tableName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:timestreamwrite:Table
    properties:
      databaseName: ${aws_timestreamwrite_database.example.database_name}
      tableName: example
```
{{% /example %}}
{{% example %}}
### Full usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.timestreamwrite.Table("example", {
    databaseName: aws_timestreamwrite_database.example.database_name,
    tableName: "example",
    retentionProperties: {
        magneticStoreRetentionPeriodInDays: 30,
        memoryStoreRetentionPeriodInHours: 8,
    },
    tags: {
        Name: "example-timestream-table",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.timestreamwrite.Table("example",
    database_name=aws_timestreamwrite_database["example"]["database_name"],
    table_name="example",
    retention_properties=aws.timestreamwrite.TableRetentionPropertiesArgs(
        magnetic_store_retention_period_in_days=30,
        memory_store_retention_period_in_hours=8,
    ),
    tags={
        "Name": "example-timestream-table",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.TimestreamWrite.Table("example", new()
    {
        DatabaseName = aws_timestreamwrite_database.Example.Database_name,
        TableName = "example",
        RetentionProperties = new Aws.TimestreamWrite.Inputs.TableRetentionPropertiesArgs
        {
            MagneticStoreRetentionPeriodInDays = 30,
            MemoryStoreRetentionPeriodInHours = 8,
        },
        Tags = 
        {
            { "Name", "example-timestream-table" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timestreamwrite.NewTable(ctx, "example", &timestreamwrite.TableArgs{
			DatabaseName: pulumi.Any(aws_timestreamwrite_database.Example.Database_name),
			TableName:    pulumi.String("example"),
			RetentionProperties: &timestreamwrite.TableRetentionPropertiesArgs{
				MagneticStoreRetentionPeriodInDays: pulumi.Int(30),
				MemoryStoreRetentionPeriodInHours:  pulumi.Int(8),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-timestream-table"),
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
import com.pulumi.aws.timestreamwrite.Table;
import com.pulumi.aws.timestreamwrite.TableArgs;
import com.pulumi.aws.timestreamwrite.inputs.TableRetentionPropertiesArgs;
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
        var example = new Table("example", TableArgs.builder()        
            .databaseName(aws_timestreamwrite_database.example().database_name())
            .tableName("example")
            .retentionProperties(TableRetentionPropertiesArgs.builder()
                .magneticStoreRetentionPeriodInDays(30)
                .memoryStoreRetentionPeriodInHours(8)
                .build())
            .tags(Map.of("Name", "example-timestream-table"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:timestreamwrite:Table
    properties:
      databaseName: ${aws_timestreamwrite_database.example.database_name}
      tableName: example
      retentionProperties:
        magneticStoreRetentionPeriodInDays: 30
        memoryStoreRetentionPeriodInHours: 8
      tags:
        Name: example-timestream-table
```
{{% /example %}}
{{% example %}}
### Customer-defined Partition Key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.timestreamwrite.Table("example", {
    databaseName: aws_timestreamwrite_database.example.database_name,
    tableName: "example",
    schema: {
        compositePartitionKey: {
            enforcementInRecord: "REQUIRED",
            name: "attr1",
            type: "DIMENSION",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.timestreamwrite.Table("example",
    database_name=aws_timestreamwrite_database["example"]["database_name"],
    table_name="example",
    schema=aws.timestreamwrite.TableSchemaArgs(
        composite_partition_key=aws.timestreamwrite.TableSchemaCompositePartitionKeyArgs(
            enforcement_in_record="REQUIRED",
            name="attr1",
            type="DIMENSION",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.TimestreamWrite.Table("example", new()
    {
        DatabaseName = aws_timestreamwrite_database.Example.Database_name,
        TableName = "example",
        Schema = new Aws.TimestreamWrite.Inputs.TableSchemaArgs
        {
            CompositePartitionKey = new Aws.TimestreamWrite.Inputs.TableSchemaCompositePartitionKeyArgs
            {
                EnforcementInRecord = "REQUIRED",
                Name = "attr1",
                Type = "DIMENSION",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timestreamwrite.NewTable(ctx, "example", &timestreamwrite.TableArgs{
			DatabaseName: pulumi.Any(aws_timestreamwrite_database.Example.Database_name),
			TableName:    pulumi.String("example"),
			Schema: &timestreamwrite.TableSchemaArgs{
				CompositePartitionKey: &timestreamwrite.TableSchemaCompositePartitionKeyArgs{
					EnforcementInRecord: pulumi.String("REQUIRED"),
					Name:                pulumi.String("attr1"),
					Type:                pulumi.String("DIMENSION"),
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
import com.pulumi.aws.timestreamwrite.Table;
import com.pulumi.aws.timestreamwrite.TableArgs;
import com.pulumi.aws.timestreamwrite.inputs.TableSchemaArgs;
import com.pulumi.aws.timestreamwrite.inputs.TableSchemaCompositePartitionKeyArgs;
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
        var example = new Table("example", TableArgs.builder()        
            .databaseName(aws_timestreamwrite_database.example().database_name())
            .tableName("example")
            .schema(TableSchemaArgs.builder()
                .compositePartitionKey(TableSchemaCompositePartitionKeyArgs.builder()
                    .enforcementInRecord("REQUIRED")
                    .name("attr1")
                    .type("DIMENSION")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:timestreamwrite:Table
    properties:
      databaseName: ${aws_timestreamwrite_database.example.database_name}
      tableName: example
      schema:
        compositePartitionKey:
          enforcementInRecord: REQUIRED
          name: attr1
          type: DIMENSION
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Timestream tables using the `table_name` and `database_name` separate by a colon (`:`). For example:

```sh
 $ pulumi import aws:timestreamwrite/table:Table example ExampleTable:ExampleDatabase
```
 