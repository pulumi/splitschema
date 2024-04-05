Provides an AWS Quantum Ledger Database (QLDB) Stream resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.qldb.Stream("example", {
    inclusiveStartTime: "2021-01-01T00:00:00Z",
    kinesisConfiguration: {
        aggregationEnabled: false,
        streamArn: "arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream",
    },
    ledgerName: "existing-ledger-name",
    roleArn: "sample-role-arn",
    streamName: "sample-ledger-stream",
    tags: {
        example: "tag",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.qldb.Stream("example",
    inclusive_start_time="2021-01-01T00:00:00Z",
    kinesis_configuration=aws.qldb.StreamKinesisConfigurationArgs(
        aggregation_enabled=False,
        stream_arn="arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream",
    ),
    ledger_name="existing-ledger-name",
    role_arn="sample-role-arn",
    stream_name="sample-ledger-stream",
    tags={
        "example": "tag",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Qldb.Stream("example", new()
    {
        InclusiveStartTime = "2021-01-01T00:00:00Z",
        KinesisConfiguration = new Aws.Qldb.Inputs.StreamKinesisConfigurationArgs
        {
            AggregationEnabled = false,
            StreamArn = "arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream",
        },
        LedgerName = "existing-ledger-name",
        RoleArn = "sample-role-arn",
        StreamName = "sample-ledger-stream",
        Tags = 
        {
            { "example", "tag" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/qldb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := qldb.NewStream(ctx, "example", &qldb.StreamArgs{
			InclusiveStartTime: pulumi.String("2021-01-01T00:00:00Z"),
			KinesisConfiguration: &qldb.StreamKinesisConfigurationArgs{
				AggregationEnabled: pulumi.Bool(false),
				StreamArn:          pulumi.String("arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream"),
			},
			LedgerName: pulumi.String("existing-ledger-name"),
			RoleArn:    pulumi.String("sample-role-arn"),
			StreamName: pulumi.String("sample-ledger-stream"),
			Tags: pulumi.StringMap{
				"example": pulumi.String("tag"),
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
import com.pulumi.aws.qldb.Stream;
import com.pulumi.aws.qldb.StreamArgs;
import com.pulumi.aws.qldb.inputs.StreamKinesisConfigurationArgs;
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
        var example = new Stream("example", StreamArgs.builder()        
            .inclusiveStartTime("2021-01-01T00:00:00Z")
            .kinesisConfiguration(StreamKinesisConfigurationArgs.builder()
                .aggregationEnabled(false)
                .streamArn("arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream")
                .build())
            .ledgerName("existing-ledger-name")
            .roleArn("sample-role-arn")
            .streamName("sample-ledger-stream")
            .tags(Map.of("example", "tag"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:qldb:Stream
    properties:
      inclusiveStartTime: 2021-01-01T00:00:00Z
      kinesisConfiguration:
        aggregationEnabled: false
        streamArn: arn:aws:kinesis:us-east-1:xxxxxxxxxxxx:stream/example-kinesis-stream
      ledgerName: existing-ledger-name
      roleArn: sample-role-arn
      streamName: sample-ledger-stream
      tags:
        example: tag
```
{{% /example %}}
{{% /examples %}}