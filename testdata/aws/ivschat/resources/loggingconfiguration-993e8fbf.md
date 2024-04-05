Resource for managing an AWS IVS (Interactive Video) Chat Logging Configuration.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage - Logging to CloudWatch

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const exampleLoggingConfiguration = new aws.ivschat.LoggingConfiguration("exampleLoggingConfiguration", {destinationConfiguration: {
    cloudwatchLogs: {
        logGroupName: exampleLogGroup.name,
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_logging_configuration = aws.ivschat.LoggingConfiguration("exampleLoggingConfiguration", destination_configuration=aws.ivschat.LoggingConfigurationDestinationConfigurationArgs(
    cloudwatch_logs=aws.ivschat.LoggingConfigurationDestinationConfigurationCloudwatchLogsArgs(
        log_group_name=example_log_group.name,
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
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var exampleLoggingConfiguration = new Aws.IvsChat.LoggingConfiguration("exampleLoggingConfiguration", new()
    {
        DestinationConfiguration = new Aws.IvsChat.Inputs.LoggingConfigurationDestinationConfigurationArgs
        {
            CloudwatchLogs = new Aws.IvsChat.Inputs.LoggingConfigurationDestinationConfigurationCloudwatchLogsArgs
            {
                LogGroupName = exampleLogGroup.Name,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivschat"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		_, err = ivschat.NewLoggingConfiguration(ctx, "exampleLoggingConfiguration", &ivschat.LoggingConfigurationArgs{
			DestinationConfiguration: &ivschat.LoggingConfigurationDestinationConfigurationArgs{
				CloudwatchLogs: &ivschat.LoggingConfigurationDestinationConfigurationCloudwatchLogsArgs{
					LogGroupName: exampleLogGroup.Name,
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.ivschat.LoggingConfiguration;
import com.pulumi.aws.ivschat.LoggingConfigurationArgs;
import com.pulumi.aws.ivschat.inputs.LoggingConfigurationDestinationConfigurationArgs;
import com.pulumi.aws.ivschat.inputs.LoggingConfigurationDestinationConfigurationCloudwatchLogsArgs;
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

        var exampleLoggingConfiguration = new LoggingConfiguration("exampleLoggingConfiguration", LoggingConfigurationArgs.builder()        
            .destinationConfiguration(LoggingConfigurationDestinationConfigurationArgs.builder()
                .cloudwatchLogs(LoggingConfigurationDestinationConfigurationCloudwatchLogsArgs.builder()
                    .logGroupName(exampleLogGroup.name())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleLoggingConfiguration:
    type: aws:ivschat:LoggingConfiguration
    properties:
      destinationConfiguration:
        cloudwatchLogs:
          logGroupName: ${exampleLogGroup.name}
```
{{% /example %}}
{{% example %}}
### Basic Usage - Logging to Kinesis Firehose with Extended S3

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {bucketPrefix: "tf-ivschat-logging-bucket"});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["firehose.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleFirehoseDeliveryStream = new aws.kinesis.FirehoseDeliveryStream("exampleFirehoseDeliveryStream", {
    destination: "extended_s3",
    extendedS3Configuration: {
        roleArn: exampleRole.arn,
        bucketArn: exampleBucketV2.arn,
    },
    tags: {
        LogDeliveryEnabled: "true",
    },
});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    acl: "private",
});
const exampleLoggingConfiguration = new aws.ivschat.LoggingConfiguration("exampleLoggingConfiguration", {destinationConfiguration: {
    firehose: {
        deliveryStreamName: exampleFirehoseDeliveryStream.name,
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", bucket_prefix="tf-ivschat-logging-bucket")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["firehose.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_firehose_delivery_stream = aws.kinesis.FirehoseDeliveryStream("exampleFirehoseDeliveryStream",
    destination="extended_s3",
    extended_s3_configuration=aws.kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs(
        role_arn=example_role.arn,
        bucket_arn=example_bucket_v2.arn,
    ),
    tags={
        "LogDeliveryEnabled": "true",
    })
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    acl="private")
example_logging_configuration = aws.ivschat.LoggingConfiguration("exampleLoggingConfiguration", destination_configuration=aws.ivschat.LoggingConfigurationDestinationConfigurationArgs(
    firehose=aws.ivschat.LoggingConfigurationDestinationConfigurationFirehoseArgs(
        delivery_stream_name=example_firehose_delivery_stream.name,
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
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        BucketPrefix = "tf-ivschat-logging-bucket",
    });

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "firehose.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleFirehoseDeliveryStream = new Aws.Kinesis.FirehoseDeliveryStream("exampleFirehoseDeliveryStream", new()
    {
        Destination = "extended_s3",
        ExtendedS3Configuration = new Aws.Kinesis.Inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs
        {
            RoleArn = exampleRole.Arn,
            BucketArn = exampleBucketV2.Arn,
        },
        Tags = 
        {
            { "LogDeliveryEnabled", "true" },
        },
    });

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Acl = "private",
    });

    var exampleLoggingConfiguration = new Aws.IvsChat.LoggingConfiguration("exampleLoggingConfiguration", new()
    {
        DestinationConfiguration = new Aws.IvsChat.Inputs.LoggingConfigurationDestinationConfigurationArgs
        {
            Firehose = new Aws.IvsChat.Inputs.LoggingConfigurationDestinationConfigurationFirehoseArgs
            {
                DeliveryStreamName = exampleFirehoseDeliveryStream.Name,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivschat"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			BucketPrefix: pulumi.String("tf-ivschat-logging-bucket"),
		})
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"firehose.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		exampleFirehoseDeliveryStream, err := kinesis.NewFirehoseDeliveryStream(ctx, "exampleFirehoseDeliveryStream", &kinesis.FirehoseDeliveryStreamArgs{
			Destination: pulumi.String("extended_s3"),
			ExtendedS3Configuration: &kinesis.FirehoseDeliveryStreamExtendedS3ConfigurationArgs{
				RoleArn:   exampleRole.Arn,
				BucketArn: exampleBucketV2.Arn,
			},
			Tags: pulumi.StringMap{
				"LogDeliveryEnabled": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: exampleBucketV2.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = ivschat.NewLoggingConfiguration(ctx, "exampleLoggingConfiguration", &ivschat.LoggingConfigurationArgs{
			DestinationConfiguration: &ivschat.LoggingConfigurationDestinationConfigurationArgs{
				Firehose: &ivschat.LoggingConfigurationDestinationConfigurationFirehoseArgs{
					DeliveryStreamName: exampleFirehoseDeliveryStream.Name,
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
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.kinesis.FirehoseDeliveryStream;
import com.pulumi.aws.kinesis.FirehoseDeliveryStreamArgs;
import com.pulumi.aws.kinesis.inputs.FirehoseDeliveryStreamExtendedS3ConfigurationArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.ivschat.LoggingConfiguration;
import com.pulumi.aws.ivschat.LoggingConfigurationArgs;
import com.pulumi.aws.ivschat.inputs.LoggingConfigurationDestinationConfigurationArgs;
import com.pulumi.aws.ivschat.inputs.LoggingConfigurationDestinationConfigurationFirehoseArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .bucketPrefix("tf-ivschat-logging-bucket")
            .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("firehose.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleFirehoseDeliveryStream = new FirehoseDeliveryStream("exampleFirehoseDeliveryStream", FirehoseDeliveryStreamArgs.builder()        
            .destination("extended_s3")
            .extendedS3Configuration(FirehoseDeliveryStreamExtendedS3ConfigurationArgs.builder()
                .roleArn(exampleRole.arn())
                .bucketArn(exampleBucketV2.arn())
                .build())
            .tags(Map.of("LogDeliveryEnabled", "true"))
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("private")
            .build());

        var exampleLoggingConfiguration = new LoggingConfiguration("exampleLoggingConfiguration", LoggingConfigurationArgs.builder()        
            .destinationConfiguration(LoggingConfigurationDestinationConfigurationArgs.builder()
                .firehose(LoggingConfigurationDestinationConfigurationFirehoseArgs.builder()
                    .deliveryStreamName(exampleFirehoseDeliveryStream.name())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleFirehoseDeliveryStream:
    type: aws:kinesis:FirehoseDeliveryStream
    properties:
      destination: extended_s3
      extendedS3Configuration:
        roleArn: ${exampleRole.arn}
        bucketArn: ${exampleBucketV2.arn}
      tags:
        LogDeliveryEnabled: 'true'
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      bucketPrefix: tf-ivschat-logging-bucket
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: private
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleLoggingConfiguration:
    type: aws:ivschat:LoggingConfiguration
    properties:
      destinationConfiguration:
        firehose:
          deliveryStreamName: ${exampleFirehoseDeliveryStream.name}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - firehose.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Basic Usage - Logging to S3

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.ivschat.LoggingConfiguration;
import com.pulumi.aws.ivschat.LoggingConfigurationArgs;
import com.pulumi.aws.ivschat.inputs.LoggingConfigurationDestinationConfigurationArgs;
import com.pulumi.aws.ivschat.inputs.LoggingConfigurationDestinationConfigurationS3Args;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .bucketName("tf-ivschat-logging")
            .forceDestroy(true)
            .build());

        var exampleLoggingConfiguration = new LoggingConfiguration("exampleLoggingConfiguration", LoggingConfigurationArgs.builder()        
            .destinationConfiguration(LoggingConfigurationDestinationConfigurationArgs.builder()
                .s3(LoggingConfigurationDestinationConfigurationS3Args.builder()
                    .bucketName(exampleBucketV2.id())
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
    properties:
      bucketName: tf-ivschat-logging
      forceDestroy: true
  exampleLoggingConfiguration:
    type: aws:ivschat:LoggingConfiguration
    properties:
      destinationConfiguration:
        s3:
          bucketName: ${exampleBucketV2.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IVS (Interactive Video) Chat Logging Configuration using the ARN. For example:

```sh
 $ pulumi import aws:ivschat/loggingConfiguration:LoggingConfiguration example arn:aws:ivschat:us-west-2:326937407773:logging-configuration/MMUQc8wcqZmC
```
 