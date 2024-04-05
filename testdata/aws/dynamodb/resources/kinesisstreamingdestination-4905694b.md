Enables a [Kinesis streaming destination](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/kds.html) for data replication of a DynamoDB table.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTable = new aws.dynamodb.Table("exampleTable", {
    hashKey: "id",
    attributes: [{
        name: "id",
        type: "S",
    }],
});
const exampleStream = new aws.kinesis.Stream("exampleStream", {shardCount: 1});
const exampleKinesisStreamingDestination = new aws.dynamodb.KinesisStreamingDestination("exampleKinesisStreamingDestination", {
    streamArn: exampleStream.arn,
    tableName: exampleTable.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example_table = aws.dynamodb.Table("exampleTable",
    hash_key="id",
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="id",
        type="S",
    )])
example_stream = aws.kinesis.Stream("exampleStream", shard_count=1)
example_kinesis_streaming_destination = aws.dynamodb.KinesisStreamingDestination("exampleKinesisStreamingDestination",
    stream_arn=example_stream.arn,
    table_name=example_table.name)
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
        HashKey = "id",
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "id",
                Type = "S",
            },
        },
    });

    var exampleStream = new Aws.Kinesis.Stream("exampleStream", new()
    {
        ShardCount = 1,
    });

    var exampleKinesisStreamingDestination = new Aws.DynamoDB.KinesisStreamingDestination("exampleKinesisStreamingDestination", new()
    {
        StreamArn = exampleStream.Arn,
        TableName = exampleTable.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
			HashKey: pulumi.String("id"),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("id"),
					Type: pulumi.String("S"),
				},
			},
		})
		if err != nil {
			return err
		}
		exampleStream, err := kinesis.NewStream(ctx, "exampleStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = dynamodb.NewKinesisStreamingDestination(ctx, "exampleKinesisStreamingDestination", &dynamodb.KinesisStreamingDestinationArgs{
			StreamArn: exampleStream.Arn,
			TableName: exampleTable.Name,
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
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.dynamodb.KinesisStreamingDestination;
import com.pulumi.aws.dynamodb.KinesisStreamingDestinationArgs;
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
            .hashKey("id")
            .attributes(TableAttributeArgs.builder()
                .name("id")
                .type("S")
                .build())
            .build());

        var exampleStream = new Stream("exampleStream", StreamArgs.builder()        
            .shardCount(1)
            .build());

        var exampleKinesisStreamingDestination = new KinesisStreamingDestination("exampleKinesisStreamingDestination", KinesisStreamingDestinationArgs.builder()        
            .streamArn(exampleStream.arn())
            .tableName(exampleTable.name())
            .build());

    }
}
```
```yaml
resources:
  exampleTable:
    type: aws:dynamodb:Table
    properties:
      hashKey: id
      attributes:
        - name: id
          type: S
  exampleStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 1
  exampleKinesisStreamingDestination:
    type: aws:dynamodb:KinesisStreamingDestination
    properties:
      streamArn: ${exampleStream.arn}
      tableName: ${exampleTable.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DynamoDB Kinesis Streaming Destinations using the `table_name` and `stream_arn` separated by `,`. For example:

```sh
 $ pulumi import aws:dynamodb/kinesisStreamingDestination:KinesisStreamingDestination example example,arn:aws:kinesis:us-east-1:111122223333:exampleStreamName
```
 