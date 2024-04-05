Manages a Kinesis Analytics v2 Application.
This resource can be used to manage both Kinesis Data Analytics for SQL applications and Kinesis Data Analytics for Apache Flink applications.

> **Note:** Kinesis Data Analytics for SQL applications created using this resource cannot currently be viewed in the AWS Console. To manage Kinesis Data Analytics for SQL applications that can also be viewed in the AWS Console, use the `aws.kinesis.AnalyticsApplication` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Apache Flink Application

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketObjectv2 = new aws.s3.BucketObjectv2("exampleBucketObjectv2", {
    bucket: exampleBucketV2.id,
    key: "example-flink-application",
    source: new pulumi.asset.FileAsset("flink-app.jar"),
});
const exampleApplication = new aws.kinesisanalyticsv2.Application("exampleApplication", {
    runtimeEnvironment: "FLINK-1_8",
    serviceExecutionRole: aws_iam_role.example.arn,
    applicationConfiguration: {
        applicationCodeConfiguration: {
            codeContent: {
                s3ContentLocation: {
                    bucketArn: exampleBucketV2.arn,
                    fileKey: exampleBucketObjectv2.key,
                },
            },
            codeContentType: "ZIPFILE",
        },
        environmentProperties: {
            propertyGroups: [
                {
                    propertyGroupId: "PROPERTY-GROUP-1",
                    propertyMap: {
                        Key1: "Value1",
                    },
                },
                {
                    propertyGroupId: "PROPERTY-GROUP-2",
                    propertyMap: {
                        KeyA: "ValueA",
                        KeyB: "ValueB",
                    },
                },
            ],
        },
        flinkApplicationConfiguration: {
            checkpointConfiguration: {
                configurationType: "DEFAULT",
            },
            monitoringConfiguration: {
                configurationType: "CUSTOM",
                logLevel: "DEBUG",
                metricsLevel: "TASK",
            },
            parallelismConfiguration: {
                autoScalingEnabled: true,
                configurationType: "CUSTOM",
                parallelism: 10,
                parallelismPerKpu: 4,
            },
        },
    },
    tags: {
        Environment: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_objectv2 = aws.s3.BucketObjectv2("exampleBucketObjectv2",
    bucket=example_bucket_v2.id,
    key="example-flink-application",
    source=pulumi.FileAsset("flink-app.jar"))
example_application = aws.kinesisanalyticsv2.Application("exampleApplication",
    runtime_environment="FLINK-1_8",
    service_execution_role=aws_iam_role["example"]["arn"],
    application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationArgs(
        application_code_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs(
            code_content=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs(
                s3_content_location=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs(
                    bucket_arn=example_bucket_v2.arn,
                    file_key=example_bucket_objectv2.key,
                ),
            ),
            code_content_type="ZIPFILE",
        ),
        environment_properties=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesArgs(
            property_groups=[
                aws.kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs(
                    property_group_id="PROPERTY-GROUP-1",
                    property_map={
                        "Key1": "Value1",
                    },
                ),
                aws.kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs(
                    property_group_id="PROPERTY-GROUP-2",
                    property_map={
                        "KeyA": "ValueA",
                        "KeyB": "ValueB",
                    },
                ),
            ],
        ),
        flink_application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs(
            checkpoint_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs(
                configuration_type="DEFAULT",
            ),
            monitoring_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs(
                configuration_type="CUSTOM",
                log_level="DEBUG",
                metrics_level="TASK",
            ),
            parallelism_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs(
                auto_scaling_enabled=True,
                configuration_type="CUSTOM",
                parallelism=10,
                parallelism_per_kpu=4,
            ),
        ),
    ),
    tags={
        "Environment": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketObjectv2 = new Aws.S3.BucketObjectv2("exampleBucketObjectv2", new()
    {
        Bucket = exampleBucketV2.Id,
        Key = "example-flink-application",
        Source = new FileAsset("flink-app.jar"),
    });

    var exampleApplication = new Aws.KinesisAnalyticsV2.Application("exampleApplication", new()
    {
        RuntimeEnvironment = "FLINK-1_8",
        ServiceExecutionRole = aws_iam_role.Example.Arn,
        ApplicationConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationArgs
        {
            ApplicationCodeConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs
            {
                CodeContent = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs
                {
                    S3ContentLocation = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs
                    {
                        BucketArn = exampleBucketV2.Arn,
                        FileKey = exampleBucketObjectv2.Key,
                    },
                },
                CodeContentType = "ZIPFILE",
            },
            EnvironmentProperties = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationEnvironmentPropertiesArgs
            {
                PropertyGroups = new[]
                {
                    new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs
                    {
                        PropertyGroupId = "PROPERTY-GROUP-1",
                        PropertyMap = 
                        {
                            { "Key1", "Value1" },
                        },
                    },
                    new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs
                    {
                        PropertyGroupId = "PROPERTY-GROUP-2",
                        PropertyMap = 
                        {
                            { "KeyA", "ValueA" },
                            { "KeyB", "ValueB" },
                        },
                    },
                },
            },
            FlinkApplicationConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs
            {
                CheckpointConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs
                {
                    ConfigurationType = "DEFAULT",
                },
                MonitoringConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs
                {
                    ConfigurationType = "CUSTOM",
                    LogLevel = "DEBUG",
                    MetricsLevel = "TASK",
                },
                ParallelismConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs
                {
                    AutoScalingEnabled = true,
                    ConfigurationType = "CUSTOM",
                    Parallelism = 10,
                    ParallelismPerKpu = 4,
                },
            },
        },
        Tags = 
        {
            { "Environment", "test" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesisanalyticsv2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleBucketObjectv2, err := s3.NewBucketObjectv2(ctx, "exampleBucketObjectv2", &s3.BucketObjectv2Args{
			Bucket: exampleBucketV2.ID(),
			Key:    pulumi.String("example-flink-application"),
			Source: pulumi.NewFileAsset("flink-app.jar"),
		})
		if err != nil {
			return err
		}
		_, err = kinesisanalyticsv2.NewApplication(ctx, "exampleApplication", &kinesisanalyticsv2.ApplicationArgs{
			RuntimeEnvironment:   pulumi.String("FLINK-1_8"),
			ServiceExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
			ApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationArgs{
				ApplicationCodeConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs{
					CodeContent: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs{
						S3ContentLocation: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs{
							BucketArn: exampleBucketV2.Arn,
							FileKey:   exampleBucketObjectv2.Key,
						},
					},
					CodeContentType: pulumi.String("ZIPFILE"),
				},
				EnvironmentProperties: &kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesArgs{
					PropertyGroups: kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArray{
						&kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs{
							PropertyGroupId: pulumi.String("PROPERTY-GROUP-1"),
							PropertyMap: pulumi.StringMap{
								"Key1": pulumi.String("Value1"),
							},
						},
						&kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs{
							PropertyGroupId: pulumi.String("PROPERTY-GROUP-2"),
							PropertyMap: pulumi.StringMap{
								"KeyA": pulumi.String("ValueA"),
								"KeyB": pulumi.String("ValueB"),
							},
						},
					},
				},
				FlinkApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs{
					CheckpointConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs{
						ConfigurationType: pulumi.String("DEFAULT"),
					},
					MonitoringConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs{
						ConfigurationType: pulumi.String("CUSTOM"),
						LogLevel:          pulumi.String("DEBUG"),
						MetricsLevel:      pulumi.String("TASK"),
					},
					ParallelismConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs{
						AutoScalingEnabled: pulumi.Bool(true),
						ConfigurationType:  pulumi.String("CUSTOM"),
						Parallelism:        pulumi.Int(10),
						ParallelismPerKpu:  pulumi.Int(4),
					},
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.kinesisanalyticsv2.Application;
import com.pulumi.aws.kinesisanalyticsv2.ApplicationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationEnvironmentPropertiesArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs;
import com.pulumi.asset.FileAsset;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketObjectv2 = new BucketObjectv2("exampleBucketObjectv2", BucketObjectv2Args.builder()        
            .bucket(exampleBucketV2.id())
            .key("example-flink-application")
            .source(new FileAsset("flink-app.jar"))
            .build());

        var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()        
            .runtimeEnvironment("FLINK-1_8")
            .serviceExecutionRole(aws_iam_role.example().arn())
            .applicationConfiguration(ApplicationApplicationConfigurationArgs.builder()
                .applicationCodeConfiguration(ApplicationApplicationConfigurationApplicationCodeConfigurationArgs.builder()
                    .codeContent(ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs.builder()
                        .s3ContentLocation(ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs.builder()
                            .bucketArn(exampleBucketV2.arn())
                            .fileKey(exampleBucketObjectv2.key())
                            .build())
                        .build())
                    .codeContentType("ZIPFILE")
                    .build())
                .environmentProperties(ApplicationApplicationConfigurationEnvironmentPropertiesArgs.builder()
                    .propertyGroups(                    
                        ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs.builder()
                            .propertyGroupId("PROPERTY-GROUP-1")
                            .propertyMap(Map.of("Key1", "Value1"))
                            .build(),
                        ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs.builder()
                            .propertyGroupId("PROPERTY-GROUP-2")
                            .propertyMap(Map.ofEntries(
                                Map.entry("KeyA", "ValueA"),
                                Map.entry("KeyB", "ValueB")
                            ))
                            .build())
                    .build())
                .flinkApplicationConfiguration(ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs.builder()
                    .checkpointConfiguration(ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs.builder()
                        .configurationType("DEFAULT")
                        .build())
                    .monitoringConfiguration(ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs.builder()
                        .configurationType("CUSTOM")
                        .logLevel("DEBUG")
                        .metricsLevel("TASK")
                        .build())
                    .parallelismConfiguration(ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs.builder()
                        .autoScalingEnabled(true)
                        .configurationType("CUSTOM")
                        .parallelism(10)
                        .parallelismPerKpu(4)
                        .build())
                    .build())
                .build())
            .tags(Map.of("Environment", "test"))
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleBucketV2.id}
      key: example-flink-application
      source:
        fn::FileAsset: flink-app.jar
  exampleApplication:
    type: aws:kinesisanalyticsv2:Application
    properties:
      runtimeEnvironment: FLINK-1_8
      serviceExecutionRole: ${aws_iam_role.example.arn}
      applicationConfiguration:
        applicationCodeConfiguration:
          codeContent:
            s3ContentLocation:
              bucketArn: ${exampleBucketV2.arn}
              fileKey: ${exampleBucketObjectv2.key}
          codeContentType: ZIPFILE
        environmentProperties:
          propertyGroups:
            - propertyGroupId: PROPERTY-GROUP-1
              propertyMap:
                Key1: Value1
            - propertyGroupId: PROPERTY-GROUP-2
              propertyMap:
                KeyA: ValueA
                KeyB: ValueB
        flinkApplicationConfiguration:
          checkpointConfiguration:
            configurationType: DEFAULT
          monitoringConfiguration:
            configurationType: CUSTOM
            logLevel: DEBUG
            metricsLevel: TASK
          parallelismConfiguration:
            autoScalingEnabled: true
            configurationType: CUSTOM
            parallelism: 10
            parallelismPerKpu: 4
      tags:
        Environment: test
```
{{% /example %}}
{{% example %}}
### SQL Application

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const exampleLogStream = new aws.cloudwatch.LogStream("exampleLogStream", {logGroupName: exampleLogGroup.name});
const exampleApplication = new aws.kinesisanalyticsv2.Application("exampleApplication", {
    runtimeEnvironment: "SQL-1_0",
    serviceExecutionRole: aws_iam_role.example.arn,
    applicationConfiguration: {
        applicationCodeConfiguration: {
            codeContent: {
                textContent: "SELECT 1;\n",
            },
            codeContentType: "PLAINTEXT",
        },
        sqlApplicationConfiguration: {
            input: {
                namePrefix: "PREFIX_1",
                inputParallelism: {
                    count: 3,
                },
                inputSchema: {
                    recordColumns: [
                        {
                            name: "COLUMN_1",
                            sqlType: "VARCHAR(8)",
                            mapping: "MAPPING-1",
                        },
                        {
                            name: "COLUMN_2",
                            sqlType: "DOUBLE",
                        },
                    ],
                    recordEncoding: "UTF-8",
                    recordFormat: {
                        recordFormatType: "CSV",
                        mappingParameters: {
                            csvMappingParameters: {
                                recordColumnDelimiter: ",",
                                recordRowDelimiter: "\n",
                            },
                        },
                    },
                },
                kinesisStreamsInput: {
                    resourceArn: aws_kinesis_stream.example.arn,
                },
            },
            outputs: [
                {
                    name: "OUTPUT_1",
                    destinationSchema: {
                        recordFormatType: "JSON",
                    },
                    lambdaOutput: {
                        resourceArn: aws_lambda_function.example.arn,
                    },
                },
                {
                    name: "OUTPUT_2",
                    destinationSchema: {
                        recordFormatType: "CSV",
                    },
                    kinesisFirehoseOutput: {
                        resourceArn: aws_kinesis_firehose_delivery_stream.example.arn,
                    },
                },
            ],
            referenceDataSource: {
                tableName: "TABLE-1",
                referenceSchema: {
                    recordColumns: [{
                        name: "COLUMN_1",
                        sqlType: "INTEGER",
                    }],
                    recordFormat: {
                        recordFormatType: "JSON",
                        mappingParameters: {
                            jsonMappingParameters: {
                                recordRowPath: "$",
                            },
                        },
                    },
                },
                s3ReferenceDataSource: {
                    bucketArn: aws_s3_bucket.example.arn,
                    fileKey: "KEY-1",
                },
            },
        },
    },
    cloudwatchLoggingOptions: {
        logStreamArn: exampleLogStream.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_log_stream = aws.cloudwatch.LogStream("exampleLogStream", log_group_name=example_log_group.name)
example_application = aws.kinesisanalyticsv2.Application("exampleApplication",
    runtime_environment="SQL-1_0",
    service_execution_role=aws_iam_role["example"]["arn"],
    application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationArgs(
        application_code_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs(
            code_content=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs(
                text_content="SELECT 1;\n",
            ),
            code_content_type="PLAINTEXT",
        ),
        sql_application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationArgs(
            input=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputArgs(
                name_prefix="PREFIX_1",
                input_parallelism=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs(
                    count=3,
                ),
                input_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs(
                    record_columns=[
                        aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs(
                            name="COLUMN_1",
                            sql_type="VARCHAR(8)",
                            mapping="MAPPING-1",
                        ),
                        aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs(
                            name="COLUMN_2",
                            sql_type="DOUBLE",
                        ),
                    ],
                    record_encoding="UTF-8",
                    record_format=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs(
                        record_format_type="CSV",
                        mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs(
                            csv_mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs(
                                record_column_delimiter=",",
                                record_row_delimiter="\n",
                            ),
                        ),
                    ),
                ),
                kinesis_streams_input=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs(
                    resource_arn=aws_kinesis_stream["example"]["arn"],
                ),
            ),
            outputs=[
                aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs(
                    name="OUTPUT_1",
                    destination_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs(
                        record_format_type="JSON",
                    ),
                    lambda_output=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs(
                        resource_arn=aws_lambda_function["example"]["arn"],
                    ),
                ),
                aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs(
                    name="OUTPUT_2",
                    destination_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs(
                        record_format_type="CSV",
                    ),
                    kinesis_firehose_output=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs(
                        resource_arn=aws_kinesis_firehose_delivery_stream["example"]["arn"],
                    ),
                ),
            ],
            reference_data_source=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs(
                table_name="TABLE-1",
                reference_schema=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs(
                    record_columns=[aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArgs(
                        name="COLUMN_1",
                        sql_type="INTEGER",
                    )],
                    record_format=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs(
                        record_format_type="JSON",
                        mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs(
                            json_mapping_parameters=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs(
                                record_row_path="$",
                            ),
                        ),
                    ),
                ),
                s3_reference_data_source=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs(
                    bucket_arn=aws_s3_bucket["example"]["arn"],
                    file_key="KEY-1",
                ),
            ),
        ),
    ),
    cloudwatch_logging_options=aws.kinesisanalyticsv2.ApplicationCloudwatchLoggingOptionsArgs(
        log_stream_arn=example_log_stream.arn,
    ))
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

    var exampleApplication = new Aws.KinesisAnalyticsV2.Application("exampleApplication", new()
    {
        RuntimeEnvironment = "SQL-1_0",
        ServiceExecutionRole = aws_iam_role.Example.Arn,
        ApplicationConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationArgs
        {
            ApplicationCodeConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs
            {
                CodeContent = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs
                {
                    TextContent = @"SELECT 1;
",
                },
                CodeContentType = "PLAINTEXT",
            },
            SqlApplicationConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationArgs
            {
                Input = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputArgs
                {
                    NamePrefix = "PREFIX_1",
                    InputParallelism = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs
                    {
                        Count = 3,
                    },
                    InputSchema = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs
                    {
                        RecordColumns = new[]
                        {
                            new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs
                            {
                                Name = "COLUMN_1",
                                SqlType = "VARCHAR(8)",
                                Mapping = "MAPPING-1",
                            },
                            new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs
                            {
                                Name = "COLUMN_2",
                                SqlType = "DOUBLE",
                            },
                        },
                        RecordEncoding = "UTF-8",
                        RecordFormat = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs
                        {
                            RecordFormatType = "CSV",
                            MappingParameters = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs
                            {
                                CsvMappingParameters = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs
                                {
                                    RecordColumnDelimiter = ",",
                                    RecordRowDelimiter = @"
",
                                },
                            },
                        },
                    },
                    KinesisStreamsInput = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs
                    {
                        ResourceArn = aws_kinesis_stream.Example.Arn,
                    },
                },
                Outputs = new[]
                {
                    new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs
                    {
                        Name = "OUTPUT_1",
                        DestinationSchema = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs
                        {
                            RecordFormatType = "JSON",
                        },
                        LambdaOutput = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs
                        {
                            ResourceArn = aws_lambda_function.Example.Arn,
                        },
                    },
                    new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs
                    {
                        Name = "OUTPUT_2",
                        DestinationSchema = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs
                        {
                            RecordFormatType = "CSV",
                        },
                        KinesisFirehoseOutput = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs
                        {
                            ResourceArn = aws_kinesis_firehose_delivery_stream.Example.Arn,
                        },
                    },
                },
                ReferenceDataSource = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs
                {
                    TableName = "TABLE-1",
                    ReferenceSchema = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs
                    {
                        RecordColumns = new[]
                        {
                            new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArgs
                            {
                                Name = "COLUMN_1",
                                SqlType = "INTEGER",
                            },
                        },
                        RecordFormat = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs
                        {
                            RecordFormatType = "JSON",
                            MappingParameters = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs
                            {
                                JsonMappingParameters = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs
                                {
                                    RecordRowPath = "$",
                                },
                            },
                        },
                    },
                    S3ReferenceDataSource = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs
                    {
                        BucketArn = aws_s3_bucket.Example.Arn,
                        FileKey = "KEY-1",
                    },
                },
            },
        },
        CloudwatchLoggingOptions = new Aws.KinesisAnalyticsV2.Inputs.ApplicationCloudwatchLoggingOptionsArgs
        {
            LogStreamArn = exampleLogStream.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesisanalyticsv2"
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
		_, err = kinesisanalyticsv2.NewApplication(ctx, "exampleApplication", &kinesisanalyticsv2.ApplicationArgs{
			RuntimeEnvironment:   pulumi.String("SQL-1_0"),
			ServiceExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
			ApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationArgs{
				ApplicationCodeConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs{
					CodeContent: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs{
						TextContent: pulumi.String("SELECT 1;\n"),
					},
					CodeContentType: pulumi.String("PLAINTEXT"),
				},
				SqlApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationArgs{
					Input: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputTypeArgs{
						NamePrefix: pulumi.String("PREFIX_1"),
						InputParallelism: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs{
							Count: pulumi.Int(3),
						},
						InputSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs{
							RecordColumns: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArray{
								&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs{
									Name:    pulumi.String("COLUMN_1"),
									SqlType: pulumi.String("VARCHAR(8)"),
									Mapping: pulumi.String("MAPPING-1"),
								},
								&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs{
									Name:    pulumi.String("COLUMN_2"),
									SqlType: pulumi.String("DOUBLE"),
								},
							},
							RecordEncoding: pulumi.String("UTF-8"),
							RecordFormat: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs{
								RecordFormatType: pulumi.String("CSV"),
								MappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs{
									CsvMappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs{
										RecordColumnDelimiter: pulumi.String(","),
										RecordRowDelimiter:    pulumi.String("\n"),
									},
								},
							},
						},
						KinesisStreamsInput: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs{
							ResourceArn: pulumi.Any(aws_kinesis_stream.Example.Arn),
						},
					},
					Outputs: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputTypeArray{
						&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputTypeArgs{
							Name: pulumi.String("OUTPUT_1"),
							DestinationSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs{
								RecordFormatType: pulumi.String("JSON"),
							},
							LambdaOutput: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs{
								ResourceArn: pulumi.Any(aws_lambda_function.Example.Arn),
							},
						},
						&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputTypeArgs{
							Name: pulumi.String("OUTPUT_2"),
							DestinationSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs{
								RecordFormatType: pulumi.String("CSV"),
							},
							KinesisFirehoseOutput: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs{
								ResourceArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
							},
						},
					},
					ReferenceDataSource: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs{
						TableName: pulumi.String("TABLE-1"),
						ReferenceSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs{
							RecordColumns: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArray{
								&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArgs{
									Name:    pulumi.String("COLUMN_1"),
									SqlType: pulumi.String("INTEGER"),
								},
							},
							RecordFormat: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs{
								RecordFormatType: pulumi.String("JSON"),
								MappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs{
									JsonMappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs{
										RecordRowPath: pulumi.String("$"),
									},
								},
							},
						},
						S3ReferenceDataSource: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs{
							BucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
							FileKey:   pulumi.String("KEY-1"),
						},
					},
				},
			},
			CloudwatchLoggingOptions: &kinesisanalyticsv2.ApplicationCloudwatchLoggingOptionsArgs{
				LogStreamArn: exampleLogStream.Arn,
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogStream;
import com.pulumi.aws.cloudwatch.LogStreamArgs;
import com.pulumi.aws.kinesisanalyticsv2.Application;
import com.pulumi.aws.kinesisanalyticsv2.ApplicationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationCloudwatchLoggingOptionsArgs;
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

        var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()        
            .runtimeEnvironment("SQL-1_0")
            .serviceExecutionRole(aws_iam_role.example().arn())
            .applicationConfiguration(ApplicationApplicationConfigurationArgs.builder()
                .applicationCodeConfiguration(ApplicationApplicationConfigurationApplicationCodeConfigurationArgs.builder()
                    .codeContent(ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs.builder()
                        .textContent("""
SELECT 1;
                        """)
                        .build())
                    .codeContentType("PLAINTEXT")
                    .build())
                .sqlApplicationConfiguration(ApplicationApplicationConfigurationSqlApplicationConfigurationArgs.builder()
                    .input(ApplicationApplicationConfigurationSqlApplicationConfigurationInputArgs.builder()
                        .namePrefix("PREFIX_1")
                        .inputParallelism(ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs.builder()
                            .count(3)
                            .build())
                        .inputSchema(ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs.builder()
                            .recordColumns(                            
                                ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs.builder()
                                    .name("COLUMN_1")
                                    .sqlType("VARCHAR(8)")
                                    .mapping("MAPPING-1")
                                    .build(),
                                ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs.builder()
                                    .name("COLUMN_2")
                                    .sqlType("DOUBLE")
                                    .build())
                            .recordEncoding("UTF-8")
                            .recordFormat(ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs.builder()
                                .recordFormatType("CSV")
                                .mappingParameters(ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs.builder()
                                    .csvMappingParameters(ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs.builder()
                                        .recordColumnDelimiter(",")
                                        .recordRowDelimiter("""

                                        """)
                                        .build())
                                    .build())
                                .build())
                            .build())
                        .kinesisStreamsInput(ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs.builder()
                            .resourceArn(aws_kinesis_stream.example().arn())
                            .build())
                        .build())
                    .outputs(                    
                        ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs.builder()
                            .name("OUTPUT_1")
                            .destinationSchema(ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs.builder()
                                .recordFormatType("JSON")
                                .build())
                            .lambdaOutput(ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs.builder()
                                .resourceArn(aws_lambda_function.example().arn())
                                .build())
                            .build(),
                        ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs.builder()
                            .name("OUTPUT_2")
                            .destinationSchema(ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs.builder()
                                .recordFormatType("CSV")
                                .build())
                            .kinesisFirehoseOutput(ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs.builder()
                                .resourceArn(aws_kinesis_firehose_delivery_stream.example().arn())
                                .build())
                            .build())
                    .referenceDataSource(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs.builder()
                        .tableName("TABLE-1")
                        .referenceSchema(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs.builder()
                            .recordColumns(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArgs.builder()
                                .name("COLUMN_1")
                                .sqlType("INTEGER")
                                .build())
                            .recordFormat(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs.builder()
                                .recordFormatType("JSON")
                                .mappingParameters(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs.builder()
                                    .jsonMappingParameters(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs.builder()
                                        .recordRowPath("$")
                                        .build())
                                    .build())
                                .build())
                            .build())
                        .s3ReferenceDataSource(ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs.builder()
                            .bucketArn(aws_s3_bucket.example().arn())
                            .fileKey("KEY-1")
                            .build())
                        .build())
                    .build())
                .build())
            .cloudwatchLoggingOptions(ApplicationCloudwatchLoggingOptionsArgs.builder()
                .logStreamArn(exampleLogStream.arn())
                .build())
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
  exampleApplication:
    type: aws:kinesisanalyticsv2:Application
    properties:
      runtimeEnvironment: SQL-1_0
      serviceExecutionRole: ${aws_iam_role.example.arn}
      applicationConfiguration:
        applicationCodeConfiguration:
          codeContent:
            textContent: |
              SELECT 1;
          codeContentType: PLAINTEXT
        sqlApplicationConfiguration:
          input:
            namePrefix: PREFIX_1
            inputParallelism:
              count: 3
            inputSchema:
              recordColumns:
                - name: COLUMN_1
                  sqlType: VARCHAR(8)
                  mapping: MAPPING-1
                - name: COLUMN_2
                  sqlType: DOUBLE
              recordEncoding: UTF-8
              recordFormat:
                recordFormatType: CSV
                mappingParameters:
                  csvMappingParameters:
                    recordColumnDelimiter: ','
                    recordRowDelimiter: |2+
            kinesisStreamsInput:
              resourceArn: ${aws_kinesis_stream.example.arn}
          outputs:
            - name: OUTPUT_1
              destinationSchema:
                recordFormatType: JSON
              lambdaOutput:
                resourceArn: ${aws_lambda_function.example.arn}
            - name: OUTPUT_2
              destinationSchema:
                recordFormatType: CSV
              kinesisFirehoseOutput:
                resourceArn: ${aws_kinesis_firehose_delivery_stream.example.arn}
          referenceDataSource:
            tableName: TABLE-1
            referenceSchema:
              recordColumns:
                - name: COLUMN_1
                  sqlType: INTEGER
              recordFormat:
                recordFormatType: JSON
                mappingParameters:
                  jsonMappingParameters:
                    recordRowPath: $
            s3ReferenceDataSource:
              bucketArn: ${aws_s3_bucket.example.arn}
              fileKey: KEY-1
      cloudwatchLoggingOptions:
        logStreamArn: ${exampleLogStream.arn}
```
{{% /example %}}
{{% example %}}
### VPC Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketObjectv2 = new aws.s3.BucketObjectv2("exampleBucketObjectv2", {
    bucket: exampleBucketV2.id,
    key: "example-flink-application",
    source: new pulumi.asset.FileAsset("flink-app.jar"),
});
const exampleApplication = new aws.kinesisanalyticsv2.Application("exampleApplication", {
    runtimeEnvironment: "FLINK-1_8",
    serviceExecutionRole: aws_iam_role.example.arn,
    applicationConfiguration: {
        applicationCodeConfiguration: {
            codeContent: {
                s3ContentLocation: {
                    bucketArn: exampleBucketV2.arn,
                    fileKey: exampleBucketObjectv2.key,
                },
            },
            codeContentType: "ZIPFILE",
        },
        vpcConfiguration: {
            securityGroupIds: [
                aws_security_group.example[0].id,
                aws_security_group.example[1].id,
            ],
            subnetIds: [aws_subnet.example.id],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_objectv2 = aws.s3.BucketObjectv2("exampleBucketObjectv2",
    bucket=example_bucket_v2.id,
    key="example-flink-application",
    source=pulumi.FileAsset("flink-app.jar"))
example_application = aws.kinesisanalyticsv2.Application("exampleApplication",
    runtime_environment="FLINK-1_8",
    service_execution_role=aws_iam_role["example"]["arn"],
    application_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationArgs(
        application_code_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs(
            code_content=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs(
                s3_content_location=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs(
                    bucket_arn=example_bucket_v2.arn,
                    file_key=example_bucket_objectv2.key,
                ),
            ),
            code_content_type="ZIPFILE",
        ),
        vpc_configuration=aws.kinesisanalyticsv2.ApplicationApplicationConfigurationVpcConfigurationArgs(
            security_group_ids=[
                aws_security_group["example"][0]["id"],
                aws_security_group["example"][1]["id"],
            ],
            subnet_ids=[aws_subnet["example"]["id"]],
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
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketObjectv2 = new Aws.S3.BucketObjectv2("exampleBucketObjectv2", new()
    {
        Bucket = exampleBucketV2.Id,
        Key = "example-flink-application",
        Source = new FileAsset("flink-app.jar"),
    });

    var exampleApplication = new Aws.KinesisAnalyticsV2.Application("exampleApplication", new()
    {
        RuntimeEnvironment = "FLINK-1_8",
        ServiceExecutionRole = aws_iam_role.Example.Arn,
        ApplicationConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationArgs
        {
            ApplicationCodeConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs
            {
                CodeContent = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs
                {
                    S3ContentLocation = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs
                    {
                        BucketArn = exampleBucketV2.Arn,
                        FileKey = exampleBucketObjectv2.Key,
                    },
                },
                CodeContentType = "ZIPFILE",
            },
            VpcConfiguration = new Aws.KinesisAnalyticsV2.Inputs.ApplicationApplicationConfigurationVpcConfigurationArgs
            {
                SecurityGroupIds = new[]
                {
                    aws_security_group.Example[0].Id,
                    aws_security_group.Example[1].Id,
                },
                SubnetIds = new[]
                {
                    aws_subnet.Example.Id,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesisanalyticsv2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleBucketObjectv2, err := s3.NewBucketObjectv2(ctx, "exampleBucketObjectv2", &s3.BucketObjectv2Args{
			Bucket: exampleBucketV2.ID(),
			Key:    pulumi.String("example-flink-application"),
			Source: pulumi.NewFileAsset("flink-app.jar"),
		})
		if err != nil {
			return err
		}
		_, err = kinesisanalyticsv2.NewApplication(ctx, "exampleApplication", &kinesisanalyticsv2.ApplicationArgs{
			RuntimeEnvironment:   pulumi.String("FLINK-1_8"),
			ServiceExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
			ApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationArgs{
				ApplicationCodeConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs{
					CodeContent: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs{
						S3ContentLocation: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs{
							BucketArn: exampleBucketV2.Arn,
							FileKey:   exampleBucketObjectv2.Key,
						},
					},
					CodeContentType: pulumi.String("ZIPFILE"),
				},
				VpcConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationVpcConfigurationArgs{
					SecurityGroupIds: pulumi.StringArray{
						aws_security_group.Example[0].Id,
						aws_security_group.Example[1].Id,
					},
					SubnetIds: pulumi.StringArray{
						aws_subnet.Example.Id,
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.kinesisanalyticsv2.Application;
import com.pulumi.aws.kinesisanalyticsv2.ApplicationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs;
import com.pulumi.aws.kinesisanalyticsv2.inputs.ApplicationApplicationConfigurationVpcConfigurationArgs;
import com.pulumi.asset.FileAsset;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketObjectv2 = new BucketObjectv2("exampleBucketObjectv2", BucketObjectv2Args.builder()        
            .bucket(exampleBucketV2.id())
            .key("example-flink-application")
            .source(new FileAsset("flink-app.jar"))
            .build());

        var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()        
            .runtimeEnvironment("FLINK-1_8")
            .serviceExecutionRole(aws_iam_role.example().arn())
            .applicationConfiguration(ApplicationApplicationConfigurationArgs.builder()
                .applicationCodeConfiguration(ApplicationApplicationConfigurationApplicationCodeConfigurationArgs.builder()
                    .codeContent(ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs.builder()
                        .s3ContentLocation(ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs.builder()
                            .bucketArn(exampleBucketV2.arn())
                            .fileKey(exampleBucketObjectv2.key())
                            .build())
                        .build())
                    .codeContentType("ZIPFILE")
                    .build())
                .vpcConfiguration(ApplicationApplicationConfigurationVpcConfigurationArgs.builder()
                    .securityGroupIds(                    
                        aws_security_group.example()[0].id(),
                        aws_security_group.example()[1].id())
                    .subnetIds(aws_subnet.example().id())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleBucketV2.id}
      key: example-flink-application
      source:
        fn::FileAsset: flink-app.jar
  exampleApplication:
    type: aws:kinesisanalyticsv2:Application
    properties:
      runtimeEnvironment: FLINK-1_8
      serviceExecutionRole: ${aws_iam_role.example.arn}
      applicationConfiguration:
        applicationCodeConfiguration:
          codeContent:
            s3ContentLocation:
              bucketArn: ${exampleBucketV2.arn}
              fileKey: ${exampleBucketObjectv2.key}
          codeContentType: ZIPFILE
        vpcConfiguration:
          securityGroupIds:
            - ${aws_security_group.example[0].id}
            - ${aws_security_group.example[1].id}
          subnetIds:
            - ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_kinesisanalyticsv2_application` using the application ARN. For example:

```sh
 $ pulumi import aws:kinesisanalyticsv2/application:Application example arn:aws:kinesisanalytics:us-west-2:123456789012:application/example-sql-application
```
 