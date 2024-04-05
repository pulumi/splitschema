Provides a resource to manage a Kinesis Stream Consumer.

> **Note:** You can register up to 20 consumers per stream. A given consumer can only be registered with one stream at a time.

For more details, see the [Amazon Kinesis Stream Consumer Documentation](https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleStream = new aws.kinesis.Stream("exampleStream", {shardCount: 1});
const exampleStreamConsumer = new aws.kinesis.StreamConsumer("exampleStreamConsumer", {streamArn: exampleStream.arn});
```
```python
import pulumi
import pulumi_aws as aws

example_stream = aws.kinesis.Stream("exampleStream", shard_count=1)
example_stream_consumer = aws.kinesis.StreamConsumer("exampleStreamConsumer", stream_arn=example_stream.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleStream = new Aws.Kinesis.Stream("exampleStream", new()
    {
        ShardCount = 1,
    });

    var exampleStreamConsumer = new Aws.Kinesis.StreamConsumer("exampleStreamConsumer", new()
    {
        StreamArn = exampleStream.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleStream, err := kinesis.NewStream(ctx, "exampleStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewStreamConsumer(ctx, "exampleStreamConsumer", &kinesis.StreamConsumerArgs{
			StreamArn: exampleStream.Arn,
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
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.kinesis.StreamConsumer;
import com.pulumi.aws.kinesis.StreamConsumerArgs;
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
        var exampleStream = new Stream("exampleStream", StreamArgs.builder()        
            .shardCount(1)
            .build());

        var exampleStreamConsumer = new StreamConsumer("exampleStreamConsumer", StreamConsumerArgs.builder()        
            .streamArn(exampleStream.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 1
  exampleStreamConsumer:
    type: aws:kinesis:StreamConsumer
    properties:
      streamArn: ${exampleStream.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Kinesis Stream Consumers using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:kinesis/streamConsumer:StreamConsumer example arn:aws:kinesis:us-west-2:123456789012:stream/example/consumer/example:1616044553
```
 