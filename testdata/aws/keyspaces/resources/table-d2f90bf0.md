Provides a Keyspaces Table.

More information about Keyspaces tables can be found in the [Keyspaces Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/working-with-tables.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.keyspaces.Table("example", {
    keyspaceName: aws_keyspaces_keyspace.example.name,
    tableName: "my_table",
    schemaDefinition: {
        columns: [{
            name: "Message",
            type: "ASCII",
        }],
        partitionKeys: [{
            name: "Message",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.keyspaces.Table("example",
    keyspace_name=aws_keyspaces_keyspace["example"]["name"],
    table_name="my_table",
    schema_definition=aws.keyspaces.TableSchemaDefinitionArgs(
        columns=[aws.keyspaces.TableSchemaDefinitionColumnArgs(
            name="Message",
            type="ASCII",
        )],
        partition_keys=[aws.keyspaces.TableSchemaDefinitionPartitionKeyArgs(
            name="Message",
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
    var example = new Aws.Keyspaces.Table("example", new()
    {
        KeyspaceName = aws_keyspaces_keyspace.Example.Name,
        TableName = "my_table",
        SchemaDefinition = new Aws.Keyspaces.Inputs.TableSchemaDefinitionArgs
        {
            Columns = new[]
            {
                new Aws.Keyspaces.Inputs.TableSchemaDefinitionColumnArgs
                {
                    Name = "Message",
                    Type = "ASCII",
                },
            },
            PartitionKeys = new[]
            {
                new Aws.Keyspaces.Inputs.TableSchemaDefinitionPartitionKeyArgs
                {
                    Name = "Message",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/keyspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyspaces.NewTable(ctx, "example", &keyspaces.TableArgs{
			KeyspaceName: pulumi.Any(aws_keyspaces_keyspace.Example.Name),
			TableName:    pulumi.String("my_table"),
			SchemaDefinition: &keyspaces.TableSchemaDefinitionArgs{
				Columns: keyspaces.TableSchemaDefinitionColumnArray{
					&keyspaces.TableSchemaDefinitionColumnArgs{
						Name: pulumi.String("Message"),
						Type: pulumi.String("ASCII"),
					},
				},
				PartitionKeys: keyspaces.TableSchemaDefinitionPartitionKeyArray{
					&keyspaces.TableSchemaDefinitionPartitionKeyArgs{
						Name: pulumi.String("Message"),
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
import com.pulumi.aws.keyspaces.Table;
import com.pulumi.aws.keyspaces.TableArgs;
import com.pulumi.aws.keyspaces.inputs.TableSchemaDefinitionArgs;
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
            .keyspaceName(aws_keyspaces_keyspace.example().name())
            .tableName("my_table")
            .schemaDefinition(TableSchemaDefinitionArgs.builder()
                .columns(TableSchemaDefinitionColumnArgs.builder()
                    .name("Message")
                    .type("ASCII")
                    .build())
                .partitionKeys(TableSchemaDefinitionPartitionKeyArgs.builder()
                    .name("Message")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:keyspaces:Table
    properties:
      keyspaceName: ${aws_keyspaces_keyspace.example.name}
      tableName: my_table
      schemaDefinition:
        columns:
          - name: Message
            type: ASCII
        partitionKeys:
          - name: Message
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a table using the `keyspace_name` and `table_name` separated by `/`. For example:

```sh
 $ pulumi import aws:keyspaces/table:Table example my_keyspace/my_table
```
 