Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
allows processing and analyzing streaming data using standard SQL.

For more details, see the [Amazon Kinesis Analytics Documentation](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/what-is.html).

> **Note:** To manage Amazon Kinesis Data Analytics for Apache Flink applications, use the `aws.kinesisanalyticsv2.Application` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Kinesis Stream Input

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStream = new aws.kinesis.Stream("testStream", {shardCount: 1});
const testApplication = new aws.kinesis.AnalyticsApplication("testApplication", {inputs: {
    namePrefix: "test_prefix",
    kinesisStream: {
        resourceArn: testStream.arn,
        roleArn: aws_iam_role.test.arn,
    },
    parallelism: {
        count: 1,
    },
    schema: {
        recordColumns: [{
            mapping: "$.test",
            name: "test",
            sqlType: "VARCHAR(8)",
        }],
        recordEncoding: "UTF-8",
        recordFormat: {
            mappingParameters: {
                json: {
                    recordRowPath: "$",
                },
            },
        },
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

test_stream = aws.kinesis.Stream("testStream", shard_count=1)
test_application = aws.kinesis.AnalyticsApplication("testApplication", inputs=aws.kinesis.AnalyticsApplicationInputsArgs(
    name_prefix="test_prefix",
    kinesis_stream=aws.kinesis.AnalyticsApplicationInputsKinesisStreamArgs(
        resource_arn=test_stream.arn,
        role_arn=aws_iam_role["test"]["arn"],
    ),
    parallelism=aws.kinesis.AnalyticsApplicationInputsParallelismArgs(
        count=1,
    ),
    schema=aws.kinesis.AnalyticsApplicationInputsSchemaArgs(
        record_columns=[aws.kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs(
            mapping="$.test",
            name="test",
            sql_type="VARCHAR(8)",
        )],
        record_encoding="UTF-8",
        record_format=aws.kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs(
            mapping_parameters=aws.kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs(
                json=aws.kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs(
                    record_row_path="$",
                ),
            ),
        ),
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
    var testStream = new Aws.Kinesis.Stream("testStream", new()
    {
        ShardCount = 1,
    });

    var testApplication = new Aws.Kinesis.AnalyticsApplication("testApplication", new()
    {
        Inputs = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsArgs
        {
            NamePrefix = "test_prefix",
            KinesisStream = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsKinesisStreamArgs
            {
                ResourceArn = testStream.Arn,
                RoleArn = aws_iam_role.Test.Arn,
            },
            Parallelism = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsParallelismArgs
            {
                Count = 1,
            },
            Schema = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaArgs
            {
                RecordColumns = new[]
                {
                    new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordColumnArgs
                    {
                        Mapping = "$.test",
                        Name = "test",
                        SqlType = "VARCHAR(8)",
                    },
                },
                RecordEncoding = "UTF-8",
                RecordFormat = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatArgs
                {
                    MappingParameters = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs
                    {
                        Json = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs
                        {
                            RecordRowPath = "$",
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewAnalyticsApplication(ctx, "testApplication", &kinesis.AnalyticsApplicationArgs{
			Inputs: &kinesis.AnalyticsApplicationInputsArgs{
				NamePrefix: pulumi.String("test_prefix"),
				KinesisStream: &kinesis.AnalyticsApplicationInputsKinesisStreamArgs{
					ResourceArn: testStream.Arn,
					RoleArn:     pulumi.Any(aws_iam_role.Test.Arn),
				},
				Parallelism: &kinesis.AnalyticsApplicationInputsParallelismArgs{
					Count: pulumi.Int(1),
				},
				Schema: &kinesis.AnalyticsApplicationInputsSchemaArgs{
					RecordColumns: kinesis.AnalyticsApplicationInputsSchemaRecordColumnArray{
						&kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs{
							Mapping: pulumi.String("$.test"),
							Name:    pulumi.String("test"),
							SqlType: pulumi.String("VARCHAR(8)"),
						},
					},
					RecordEncoding: pulumi.String("UTF-8"),
					RecordFormat: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs{
						MappingParameters: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs{
							Json: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs{
								RecordRowPath: pulumi.String("$"),
							},
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
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.kinesis.AnalyticsApplication;
import com.pulumi.aws.kinesis.AnalyticsApplicationArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsKinesisStreamArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsParallelismArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaRecordFormatArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs;
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
        var testStream = new Stream("testStream", StreamArgs.builder()        
            .shardCount(1)
            .build());

        var testApplication = new AnalyticsApplication("testApplication", AnalyticsApplicationArgs.builder()        
            .inputs(AnalyticsApplicationInputsArgs.builder()
                .namePrefix("test_prefix")
                .kinesisStream(AnalyticsApplicationInputsKinesisStreamArgs.builder()
                    .resourceArn(testStream.arn())
                    .roleArn(aws_iam_role.test().arn())
                    .build())
                .parallelism(AnalyticsApplicationInputsParallelismArgs.builder()
                    .count(1)
                    .build())
                .schema(AnalyticsApplicationInputsSchemaArgs.builder()
                    .recordColumns(AnalyticsApplicationInputsSchemaRecordColumnArgs.builder()
                        .mapping("$.test")
                        .name("test")
                        .sqlType("VARCHAR(8)")
                        .build())
                    .recordEncoding("UTF-8")
                    .recordFormat(AnalyticsApplicationInputsSchemaRecordFormatArgs.builder()
                        .mappingParameters(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs.builder()
                            .json(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs.builder()
                                .recordRowPath("$")
                                .build())
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 1
  testApplication:
    type: aws:kinesis:AnalyticsApplication
    properties:
      inputs:
        namePrefix: test_prefix
        kinesisStream:
          resourceArn: ${testStream.arn}
          roleArn: ${aws_iam_role.test.arn}
        parallelism:
          count: 1
        schema:
          recordColumns:
            - mapping: $.test
              name: test
              sqlType: VARCHAR(8)
          recordEncoding: UTF-8
          recordFormat:
            mappingParameters:
              json:
                recordRowPath: $
```
{{% /example %}}
{{% example %}}
### Starting An Application

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const exampleLogStream = new aws.cloudwatch.LogStream("exampleLogStream", {logGroupName: exampleLogGroup.name});
const exampleStream = new aws.kinesis.Stream("exampleStream", {shardCount: 1});
const exampleFirehoseDeliveryStream = new aws.kinesis.FirehoseDeliveryStream("exampleFirehoseDeliveryStream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        bucketArn: aws_s3_bucket.example.arn,
        roleArn: aws_iam_role.example.arn,
    },
});
const test = new aws.kinesis.AnalyticsApplication("test", {
    cloudwatchLoggingOptions: {
        logStreamArn: exampleLogStream.arn,
        roleArn: aws_iam_role.example.arn,
    },
    inputs: {
        namePrefix: "example_prefix",
        schema: {
            recordColumns: [{
                name: "COLUMN_1",
                sqlType: "INTEGER",
            }],
            recordFormat: {
                mappingParameters: {
                    csv: {
                        recordColumnDelimiter: ",",
                        recordRowDelimiter: "|",
                    },
                },
            },
        },
        kinesisStream: {
            resourceArn: exampleStream.arn,
            roleArn: aws_iam_role.example.arn,
        },
        startingPositionConfigurations: [{
            startingPosition: "NOW",
        }],
    },
    outputs: [{
        name: "OUTPUT_1",
        schema: {
            recordFormatType: "CSV",
        },
        kinesisFirehose: {
            resourceArn: exampleFirehoseDeliveryStream.arn,
            roleArn: aws_iam_role.example.arn,
        },
    }],
    startApplication: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_log_stream = aws.cloudwatch.LogStream("exampleLogStream", log_group_name=example_log_group.name)
example_stream = aws.kinesis.Stream("exampleStream", shard_count=1)
example_firehose_delivery_stream = aws.kinesis.FirehoseDeliveryStream("exampleFirehoseDeliveryStream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        bucket_arn=aws_s3_bucket["example"]["arn"],
        role_arn=aws_iam_role["example"]["arn"],
    ))
test = aws.kinesis.AnalyticsApplication("test",
    cloudwatch_logging_options=aws.kinesis.AnalyticsApplicationCloudwatchLoggingOptionsArgs(
        log_stream_arn=example_log_stream.arn,
        role_arn=aws_iam_role["example"]["arn"],
    ),
    inputs=aws.kinesis.AnalyticsApplicationInputsArgs(
        name_prefix="example_prefix",
        schema=aws.kinesis.AnalyticsApplicationInputsSchemaArgs(
            record_columns=[aws.kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs(
                name="COLUMN_1",
                sql_type="INTEGER",
            )],
            record_format=aws.kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs(
                mapping_parameters=aws.kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs(
                    csv=aws.kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs(
                        record_column_delimiter=",",
                        record_row_delimiter="|",
                    ),
                ),
            ),
        ),
        kinesis_stream=aws.kinesis.AnalyticsApplicationInputsKinesisStreamArgs(
            resource_arn=example_stream.arn,
            role_arn=aws_iam_role["example"]["arn"],
        ),
        starting_position_configurations=[aws.kinesis.AnalyticsApplicationInputsStartingPositionConfigurationArgs(
            starting_position="NOW",
        )],
    ),
    outputs=[aws.kinesis.AnalyticsApplicationOutputArgs(
        name="OUTPUT_1",
        schema=aws.kinesis.AnalyticsApplicationOutputSchemaArgs(
            record_format_type="CSV",
        ),
        kinesis_firehose=aws.kinesis.AnalyticsApplicationOutputKinesisFirehoseArgs(
            resource_arn=example_firehose_delivery_stream.arn,
            role_arn=aws_iam_role["example"]["arn"],
        ),
    )],
    start_application=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var exampleLogStream = new Aws.CloudWatch.LogStream("exampleLogStream", new()
    {
        LogGroupName = exampleLogGroup.Name,
    });

    var exampleStream = new Aws.Kinesis.Stream("exampleStream", new()
    {
        ShardCount = 1,
    });

    var exampleFirehoseDeliveryStream = new Aws.Kinesis.FirehoseDeliveryStream("exampleFirehoseDeliveryStream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            BucketArn = aws_s3_bucket.Example.Arn,
            RoleArn = aws_iam_role.Example.Arn,
        },
    });

    var test = new Aws.Kinesis.AnalyticsApplication("test", new()
    {
        CloudwatchLoggingOptions = new Aws.Kinesis.Inputs.AnalyticsApplicationCloudwatchLoggingOptionsArgs
        {
            LogStreamArn = exampleLogStream.Arn,
            RoleArn = aws_iam_role.Example.Arn,
        },
        Inputs = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsArgs
        {
            NamePrefix = "example_prefix",
            Schema = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaArgs
            {
                RecordColumns = new[]
                {
                    new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordColumnArgs
                    {
                        Name = "COLUMN_1",
                        SqlType = "INTEGER",
                    },
                },
                RecordFormat = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatArgs
                {
                    MappingParameters = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs
                    {
                        Csv = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs
                        {
                            RecordColumnDelimiter = ",",
                            RecordRowDelimiter = "|",
                        },
                    },
                },
            },
            KinesisStream = new Aws.Kinesis.Inputs.AnalyticsApplicationInputsKinesisStreamArgs
            {
                ResourceArn = exampleStream.Arn,
                RoleArn = aws_iam_role.Example.Arn,
            },
            StartingPositionConfigurations = new[]
            {
                new Aws.Kinesis.Inputs.AnalyticsApplicationInputsStartingPositionConfigurationArgs
                {
                    StartingPosition = "NOW",
                },
            },
        },
        Outputs = new[]
        {
            new Aws.Kinesis.Inputs.AnalyticsApplicationOutputArgs
            {
                Name = "OUTPUT_1",
                Schema = new Aws.Kinesis.Inputs.AnalyticsApplicationOutputSchemaArgs
                {
                    RecordFormatType = "CSV",
                },
                KinesisFirehose = new Aws.Kinesis.Inputs.AnalyticsApplicationOutputKinesisFirehoseArgs
                {
                    ResourceArn = exampleFirehoseDeliveryStream.Arn,
                    RoleArn = aws_iam_role.Example.Arn,
                },
            },
        },
        StartApplication = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		exampleLogStream, err := cloudwatch.NewLogStream(ctx, "exampleLogStream", &cloudwatch.LogStreamArgs{
			LogGroupName: exampleLogGroup.Name,
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
		exampleFirehoseDeliveryStream, err := kinesis.NewFirehoseDeliveryStream(ctx, "exampleFirehoseDeliveryStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				BucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
				RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
			},
		})
		if err != nil {
			return err
		}
		_, err = kinesis.NewAnalyticsApplication(ctx, "test", &kinesis.AnalyticsApplicationArgs{
			CloudwatchLoggingOptions: &kinesis.AnalyticsApplicationCloudwatchLoggingOptionsArgs{
				LogStreamArn: exampleLogStream.Arn,
				RoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
			},
			Inputs: &kinesis.AnalyticsApplicationInputsArgs{
				NamePrefix: pulumi.String("example_prefix"),
				Schema: &kinesis.AnalyticsApplicationInputsSchemaArgs{
					RecordColumns: kinesis.AnalyticsApplicationInputsSchemaRecordColumnArray{
						&kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs{
							Name:    pulumi.String("COLUMN_1"),
							SqlType: pulumi.String("INTEGER"),
						},
					},
					RecordFormat: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs{
						MappingParameters: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs{
							Csv: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs{
								RecordColumnDelimiter: pulumi.String(","),
								RecordRowDelimiter:    pulumi.String("|"),
							},
						},
					},
				},
				KinesisStream: &kinesis.AnalyticsApplicationInputsKinesisStreamArgs{
					ResourceArn: exampleStream.Arn,
					RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
				},
				StartingPositionConfigurations: kinesis.AnalyticsApplicationInputsStartingPositionConfigurationArray{
					&kinesis.AnalyticsApplicationInputsStartingPositionConfigurationArgs{
						StartingPosition: pulumi.String("NOW"),
					},
				},
			},
			Outputs: kinesis.AnalyticsApplicationOutputTypeArray{
				&kinesis.AnalyticsApplicationOutputTypeArgs{
					Name: pulumi.String("OUTPUT_1"),
					Schema: &kinesis.AnalyticsApplicationOutputSchemaArgs{
						RecordFormatType: pulumi.String("CSV"),
					},
					KinesisFirehose: &kinesis.AnalyticsApplicationOutputKinesisFirehoseArgs{
						ResourceArn: exampleFirehoseDeliveryStream.Arn,
						RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
					},
				},
			},
			StartApplication: pulumi.Bool(true),
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogStream;
import com.pulumi.aws.cloudwatch.LogStreamArgs;
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.kinesis.AnalyticsApplication;
import com.pulumi.aws.kinesis.AnalyticsApplicationArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationCloudwatchLoggingOptionsArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaRecordFormatArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationInputsKinesisStreamArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationOutputArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationOutputSchemaArgs;
import com.pulumi.aws.kinesis.inputs.AnalyticsApplicationOutputKinesisFirehoseArgs;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup");

        var exampleLogStream = new LogStream("exampleLogStream", LogStreamArgs.builder()        
            .logGroupName(exampleLogGroup.name())
            .build());

        var exampleStream = new Stream("exampleStream", StreamArgs.builder()        
            .shardCount(1)
            .build());

        var exampleFirehoseDeliveryStream = new FirehoseDeliveryStream("exampleFirehoseDeliveryStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .bucketArn(aws_s3_bucket.example().arn())
                .roleArn(aws_iam_role.example().arn())
                .build())
            .build());

        var test = new AnalyticsApplication("test", AnalyticsApplicationArgs.builder()        
            .cloudwatchLoggingOptions(AnalyticsApplicationCloudwatchLoggingOptionsArgs.builder()
                .logStreamArn(exampleLogStream.arn())
                .roleArn(aws_iam_role.example().arn())
                .build())
            .inputs(AnalyticsApplicationInputsArgs.builder()
                .namePrefix("example_prefix")
                .schema(AnalyticsApplicationInputsSchemaArgs.builder()
                    .recordColumns(AnalyticsApplicationInputsSchemaRecordColumnArgs.builder()
                        .name("COLUMN_1")
                        .sqlType("INTEGER")
                        .build())
                    .recordFormat(AnalyticsApplicationInputsSchemaRecordFormatArgs.builder()
                        .mappingParameters(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs.builder()
                            .csv(AnalyticsApplicationInputsSchemaRecordFormatMappingParametersCsvArgs.builder()
                                .recordColumnDelimiter(",")
                                .recordRowDelimiter("|")
                                .build())
                            .build())
                        .build())
                    .build())
                .kinesisStream(AnalyticsApplicationInputsKinesisStreamArgs.builder()
                    .resourceArn(exampleStream.arn())
                    .roleArn(aws_iam_role.example().arn())
                    .build())
                .startingPositionConfigurations(AnalyticsApplicationInputsStartingPositionConfigurationArgs.builder()
                    .startingPosition("NOW")
                    .build())
                .build())
            .outputs(AnalyticsApplicationOutputArgs.builder()
                .name("OUTPUT_1")
                .schema(AnalyticsApplicationOutputSchemaArgs.builder()
                    .recordFormatType("CSV")
                    .build())
                .kinesisFirehose(AnalyticsApplicationOutputKinesisFirehoseArgs.builder()
                    .resourceArn(exampleFirehoseDeliveryStream.arn())
                    .roleArn(aws_iam_role.example().arn())
                    .build())
                .build())
            .startApplication(true)
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleLogStream:
    type: aws:cloudwatch:LogStream
    properties:
      logGroupName: ${exampleLogGroup.name}
  exampleStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 1
  exampleFirehoseDeliveryStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        bucketArn: ${aws_s3_bucket.example.arn}
        roleArn: ${aws_iam_role.example.arn}
  test:
    type: aws:kinesis:AnalyticsApplication
    properties:
      cloudwatchLoggingOptions:
        logStreamArn: ${exampleLogStream.arn}
        roleArn: ${aws_iam_role.example.arn}
      inputs:
        namePrefix: example_prefix
        schema:
          recordColumns:
            - name: COLUMN_1
              sqlType: INTEGER
          recordFormat:
            mappingParameters:
              csv:
                recordColumnDelimiter: ','
                recordRowDelimiter: '|'
        kinesisStream:
          resourceArn: ${exampleStream.arn}
          roleArn: ${aws_iam_role.example.arn}
        startingPositionConfigurations:
          - startingPosition: NOW
      outputs:
        - name: OUTPUT_1
          schema:
            recordFormatType: CSV
          kinesisFirehose:
            resourceArn: ${exampleFirehoseDeliveryStream.arn}
            roleArn: ${aws_iam_role.example.arn}
      startApplication: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Kinesis Analytics Application using ARN. For example:

```sh
 $ pulumi import aws:kinesis/analyticsApplication:AnalyticsApplication example arn:aws:kinesisanalytics:us-west-2:1234567890:application/example
```
 