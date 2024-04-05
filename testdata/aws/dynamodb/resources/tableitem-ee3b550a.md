Provides a DynamoDB table item resource

> **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
  You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTable = new aws.dynamodb.Table("exampleTable", {
    readCapacity: 10,
    writeCapacity: 10,
    hashKey: "exampleHashKey",
    attributes: [{
        name: "exampleHashKey",
        type: "S",
    }],
});
const exampleTableItem = new aws.dynamodb.TableItem("exampleTableItem", {
    tableName: exampleTable.name,
    hashKey: exampleTable.hashKey,
    item: `{
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

example_table = aws.dynamodb.Table("exampleTable",
    read_capacity=10,
    write_capacity=10,
    hash_key="exampleHashKey",
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="exampleHashKey",
        type="S",
    )])
example_table_item = aws.dynamodb.TableItem("exampleTableItem",
    table_name=example_table.name,
    hash_key=example_table.hash_key,
    item="""{
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleTable = new Aws.DynamoDB.Table("exampleTable", new()
    {
        ReadCapacity = 10,
        WriteCapacity = 10,
        HashKey = "exampleHashKey",
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "exampleHashKey",
                Type = "S",
            },
        },
    });

    var exampleTableItem = new Aws.DynamoDB.TableItem("exampleTableItem", new()
    {
        TableName = exampleTable.Name,
        HashKey = exampleTable.HashKey,
        Item = @"{
  ""exampleHashKey"": {""S"": ""something""},
  ""one"": {""N"": ""11111""},
  ""two"": {""N"": ""22222""},
  ""three"": {""N"": ""33333""},
  ""four"": {""N"": ""44444""}
}
",
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
		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
			ReadCapacity:  pulumi.Int(10),
			WriteCapacity: pulumi.Int(10),
			HashKey:       pulumi.String("exampleHashKey"),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("exampleHashKey"),
					Type: pulumi.String("S"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTableItem(ctx, "exampleTableItem", &dynamodb.TableItemArgs{
			TableName: exampleTable.Name,
			HashKey:   exampleTable.HashKey,
			Item: pulumi.String(`{
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}
`),
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
import com.pulumi.aws.dynamodb.TableItem;
import com.pulumi.aws.dynamodb.TableItemArgs;
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
        var exampleTable = new Table("exampleTable", TableArgs.builder()        
            .readCapacity(10)
            .writeCapacity(10)
            .hashKey("exampleHashKey")
            .attributes(TableAttributeArgs.builder()
                .name("exampleHashKey")
                .type("S")
                .build())
            .build());

        var exampleTableItem = new TableItem("exampleTableItem", TableItemArgs.builder()        
            .tableName(exampleTable.name())
            .hashKey(exampleTable.hashKey())
            .item("""
{
  "exampleHashKey": {"S": "something"},
  "one": {"N": "11111"},
  "two": {"N": "22222"},
  "three": {"N": "33333"},
  "four": {"N": "44444"}
}
            """)
            .build());

    }
}
```
```yaml
resources:
  exampleTableItem:
    type: aws:dynamodb:TableItem
    properties:
      tableName: ${exampleTable.name}
      hashKey: ${exampleTable.hashKey}
      item: |
        {
          "exampleHashKey": {"S": "something"},
          "one": {"N": "11111"},
          "two": {"N": "22222"},
          "three": {"N": "33333"},
          "four": {"N": "44444"}
        }
  exampleTable:
    type: aws:dynamodb:Table
    properties:
      readCapacity: 10
      writeCapacity: 10
      hashKey: exampleHashKey
      attributes:
        - name: exampleHashKey
          type: S
```
{{% /example %}}
{{% /examples %}}

## Import

You cannot import DynamoDB table items. 