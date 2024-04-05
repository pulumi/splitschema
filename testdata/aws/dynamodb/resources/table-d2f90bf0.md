Provides a DynamoDB table resource.

> **Note:** It is recommended to use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) for `read_capacity` and/or `write_capacity` if there's `autoscaling policy` attached to the table.

> **Note:** When using aws.dynamodb.TableReplica with this resource, use `lifecycle` `ignore_changes` for `replica`, _e.g._, `lifecycle { ignore_changes = [replica] }`.

## DynamoDB Table attributes

Only define attributes on the table object that are going to be used as:

* Table hash key or range key
* LSI or GSI hash key or range key

The DynamoDB API expects attribute structure (name and type) to be passed along when creating or updating GSI/LSIs or creating the initial table. In these cases it expects the Hash / Range keys to be provided. Because these get re-used in numerous places (i.e the table's range key could be a part of one or more GSIs), they are stored on the table object to prevent duplication and increase consistency. If you add attributes here that are not used in these scenarios it can cause an infinite loop in planning.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

The following dynamodb table description models the table and GSI shown in the [AWS SDK example documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const basic_dynamodb_table = new aws.dynamodb.Table("basic-dynamodb-table", {
    attributes: [
        {
            name: "UserId",
            type: "S",
        },
        {
            name: "GameTitle",
            type: "S",
        },
        {
            name: "TopScore",
            type: "N",
        },
    ],
    billingMode: "PROVISIONED",
    globalSecondaryIndexes: [{
        hashKey: "GameTitle",
        name: "GameTitleIndex",
        nonKeyAttributes: ["UserId"],
        projectionType: "INCLUDE",
        rangeKey: "TopScore",
        readCapacity: 10,
        writeCapacity: 10,
    }],
    hashKey: "UserId",
    rangeKey: "GameTitle",
    readCapacity: 20,
    tags: {
        Environment: "production",
        Name: "dynamodb-table-1",
    },
    ttl: {
        attributeName: "TimeToExist",
        enabled: false,
    },
    writeCapacity: 20,
});
```
```python
import pulumi
import pulumi_aws as aws

basic_dynamodb_table = aws.dynamodb.Table("basic-dynamodb-table",
    attributes=[
        aws.dynamodb.TableAttributeArgs(
            name="UserId",
            type="S",
        ),
        aws.dynamodb.TableAttributeArgs(
            name="GameTitle",
            type="S",
        ),
        aws.dynamodb.TableAttributeArgs(
            name="TopScore",
            type="N",
        ),
    ],
    billing_mode="PROVISIONED",
    global_secondary_indexes=[aws.dynamodb.TableGlobalSecondaryIndexArgs(
        hash_key="GameTitle",
        name="GameTitleIndex",
        non_key_attributes=["UserId"],
        projection_type="INCLUDE",
        range_key="TopScore",
        read_capacity=10,
        write_capacity=10,
    )],
    hash_key="UserId",
    range_key="GameTitle",
    read_capacity=20,
    tags={
        "Environment": "production",
        "Name": "dynamodb-table-1",
    },
    ttl=aws.dynamodb.TableTtlArgs(
        attribute_name="TimeToExist",
        enabled=False,
    ),
    write_capacity=20)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var basic_dynamodb_table = new Aws.DynamoDB.Table("basic-dynamodb-table", new()
    {
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "UserId",
                Type = "S",
            },
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "GameTitle",
                Type = "S",
            },
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "TopScore",
                Type = "N",
            },
        },
        BillingMode = "PROVISIONED",
        GlobalSecondaryIndexes = new[]
        {
            new Aws.DynamoDB.Inputs.TableGlobalSecondaryIndexArgs
            {
                HashKey = "GameTitle",
                Name = "GameTitleIndex",
                NonKeyAttributes = new[]
                {
                    "UserId",
                },
                ProjectionType = "INCLUDE",
                RangeKey = "TopScore",
                ReadCapacity = 10,
                WriteCapacity = 10,
            },
        },
        HashKey = "UserId",
        RangeKey = "GameTitle",
        ReadCapacity = 20,
        Tags = 
        {
            { "Environment", "production" },
            { "Name", "dynamodb-table-1" },
        },
        Ttl = new Aws.DynamoDB.Inputs.TableTtlArgs
        {
            AttributeName = "TimeToExist",
            Enabled = false,
        },
        WriteCapacity = 20,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewTable(ctx, "basic-dynamodb-table", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("UserId"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("GameTitle"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("TopScore"),
					Type: pulumi.String("N"),
				},
			},
			BillingMode: pulumi.String("PROVISIONED"),
			GlobalSecondaryIndexes: dynamodb.TableGlobalSecondaryIndexArray{
				&dynamodb.TableGlobalSecondaryIndexArgs{
					HashKey: pulumi.String("GameTitle"),
					Name:    pulumi.String("GameTitleIndex"),
					NonKeyAttributes: pulumi.StringArray{
						pulumi.String("UserId"),
					},
					ProjectionType: pulumi.String("INCLUDE"),
					RangeKey:       pulumi.String("TopScore"),
					ReadCapacity:   pulumi.Int(10),
					WriteCapacity:  pulumi.Int(10),
				},
			},
			HashKey:      pulumi.String("UserId"),
			RangeKey:     pulumi.String("GameTitle"),
			ReadCapacity: pulumi.Int(20),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
				"Name":        pulumi.String("dynamodb-table-1"),
			},
			Ttl: &dynamodb.TableTtlArgs{
				AttributeName: pulumi.String("TimeToExist"),
				Enabled:       pulumi.Bool(false),
			},
			WriteCapacity: pulumi.Int(20),
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
import com.pulumi.aws.dynamodb.Table;
import com.pulumi.aws.dynamodb.TableArgs;
import com.pulumi.aws.dynamodb.inputs.TableAttributeArgs;
import com.pulumi.aws.dynamodb.inputs.TableGlobalSecondaryIndexArgs;
import com.pulumi.aws.dynamodb.inputs.TableTtlArgs;
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
        var basic_dynamodb_table = new Table("basic-dynamodb-table", TableArgs.builder()        
            .attributes(            
                TableAttributeArgs.builder()
                    .name("UserId")
                    .type("S")
                    .build(),
                TableAttributeArgs.builder()
                    .name("GameTitle")
                    .type("S")
                    .build(),
                TableAttributeArgs.builder()
                    .name("TopScore")
                    .type("N")
                    .build())
            .billingMode("PROVISIONED")
            .globalSecondaryIndexes(TableGlobalSecondaryIndexArgs.builder()
                .hashKey("GameTitle")
                .name("GameTitleIndex")
                .nonKeyAttributes("UserId")
                .projectionType("INCLUDE")
                .rangeKey("TopScore")
                .readCapacity(10)
                .writeCapacity(10)
                .build())
            .hashKey("UserId")
            .rangeKey("GameTitle")
            .readCapacity(20)
            .tags(Map.ofEntries(
                Map.entry("Environment", "production"),
                Map.entry("Name", "dynamodb-table-1")
            ))
            .ttl(TableTtlArgs.builder()
                .attributeName("TimeToExist")
                .enabled(false)
                .build())
            .writeCapacity(20)
            .build());

    }
}
```
```yaml
resources:
  basic-dynamodb-table:
    type: aws:dynamodb:Table
    properties:
      attributes:
        - name: UserId
          type: S
        - name: GameTitle
          type: S
        - name: TopScore
          type: N
      billingMode: PROVISIONED
      globalSecondaryIndexes:
        - hashKey: GameTitle
          name: GameTitleIndex
          nonKeyAttributes:
            - UserId
          projectionType: INCLUDE
          rangeKey: TopScore
          readCapacity: 10
          writeCapacity: 10
      hashKey: UserId
      rangeKey: GameTitle
      readCapacity: 20
      tags:
        Environment: production
        Name: dynamodb-table-1
      ttl:
        attributeName: TimeToExist
        enabled: false
      writeCapacity: 20
```
{{% /example %}}
{{% example %}}
### Global Tables

This resource implements support for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) via `replica` configuration blocks. For working with [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html), see the `aws.dynamodb.GlobalTable` resource.

> **Note:** aws.dynamodb.TableReplica is an alternate way of configuring Global Tables. Do not use `replica` configuration blocks of `aws.dynamodb.Table` together with aws_dynamodb_table_replica.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dynamodb.Table("example", {
    attributes: [{
        name: "TestTableHashKey",
        type: "S",
    }],
    billingMode: "PAY_PER_REQUEST",
    hashKey: "TestTableHashKey",
    replicas: [
        {
            regionName: "us-east-2",
        },
        {
            regionName: "us-west-2",
        },
    ],
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dynamodb.Table("example",
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="TestTableHashKey",
        type="S",
    )],
    billing_mode="PAY_PER_REQUEST",
    hash_key="TestTableHashKey",
    replicas=[
        aws.dynamodb.TableReplicaArgs(
            region_name="us-east-2",
        ),
        aws.dynamodb.TableReplicaArgs(
            region_name="us-west-2",
        ),
    ],
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DynamoDB.Table("example", new()
    {
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "TestTableHashKey",
                Type = "S",
            },
        },
        BillingMode = "PAY_PER_REQUEST",
        HashKey = "TestTableHashKey",
        Replicas = new[]
        {
            new Aws.DynamoDB.Inputs.TableReplicaArgs
            {
                RegionName = "us-east-2",
            },
            new Aws.DynamoDB.Inputs.TableReplicaArgs
            {
                RegionName = "us-west-2",
            },
        },
        StreamEnabled = true,
        StreamViewType = "NEW_AND_OLD_IMAGES",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewTable(ctx, "example", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("TestTableHashKey"),
					Type: pulumi.String("S"),
				},
			},
			BillingMode: pulumi.String("PAY_PER_REQUEST"),
			HashKey:     pulumi.String("TestTableHashKey"),
			Replicas: dynamodb.TableReplicaTypeArray{
				&dynamodb.TableReplicaTypeArgs{
					RegionName: pulumi.String("us-east-2"),
				},
				&dynamodb.TableReplicaTypeArgs{
					RegionName: pulumi.String("us-west-2"),
				},
			},
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
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
import com.pulumi.aws.dynamodb.Table;
import com.pulumi.aws.dynamodb.TableArgs;
import com.pulumi.aws.dynamodb.inputs.TableAttributeArgs;
import com.pulumi.aws.dynamodb.inputs.TableReplicaArgs;
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
            .attributes(TableAttributeArgs.builder()
                .name("TestTableHashKey")
                .type("S")
                .build())
            .billingMode("PAY_PER_REQUEST")
            .hashKey("TestTableHashKey")
            .replicas(            
                TableReplicaArgs.builder()
                    .regionName("us-east-2")
                    .build(),
                TableReplicaArgs.builder()
                    .regionName("us-west-2")
                    .build())
            .streamEnabled(true)
            .streamViewType("NEW_AND_OLD_IMAGES")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dynamodb:Table
    properties:
      attributes:
        - name: TestTableHashKey
          type: S
      billingMode: PAY_PER_REQUEST
      hashKey: TestTableHashKey
      replicas:
        - regionName: us-east-2
        - regionName: us-west-2
      streamEnabled: true
      streamViewType: NEW_AND_OLD_IMAGES
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DynamoDB tables using the `name`. For example:

```sh
 $ pulumi import aws:dynamodb/table:Table basic-dynamodb-table GameScores
```
 