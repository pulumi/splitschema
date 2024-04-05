Manages a S3 Bucket Notification Configuration. For additional information, see the [Configuring S3 Event Notifications section in the Amazon S3 Developer Guide](https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html).

> **NOTE:** S3 Buckets only support a single notification configuration. Declaring multiple `aws.s3.BucketNotification` resources to the same S3 Bucket will cause a perpetual difference in configuration. See the example "Trigger multiple Lambda functions" for an option.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Add notification configuration to SNS Topic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.BucketV2("bucket", {});
const topicPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["s3.amazonaws.com"],
        }],
        actions: ["SNS:Publish"],
        resources: ["arn:aws:sns:*:*:s3-event-notification-topic"],
        conditions: [{
            test: "ArnLike",
            variable: "aws:SourceArn",
            values: [bucket.arn],
        }],
    }],
});
const topicTopic = new aws.sns.Topic("topicTopic", {policy: topicPolicyDocument.apply(topicPolicyDocument => topicPolicyDocument.json)});
const bucketNotification = new aws.s3.BucketNotification("bucketNotification", {
    bucket: bucket.id,
    topics: [{
        topicArn: topicTopic.arn,
        events: ["s3:ObjectCreated:*"],
        filterSuffix: ".log",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.BucketV2("bucket")
topic_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["s3.amazonaws.com"],
    )],
    actions=["SNS:Publish"],
    resources=["arn:aws:sns:*:*:s3-event-notification-topic"],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="ArnLike",
        variable="aws:SourceArn",
        values=[bucket.arn],
    )],
)])
topic_topic = aws.sns.Topic("topicTopic", policy=topic_policy_document.json)
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    topics=[aws.s3.BucketNotificationTopicArgs(
        topic_arn=topic_topic.arn,
        events=["s3:ObjectCreated:*"],
        filter_suffix=".log",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucket = new Aws.S3.BucketV2("bucket");

    var topicPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "s3.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "SNS:Publish",
                },
                Resources = new[]
                {
                    "arn:aws:sns:*:*:s3-event-notification-topic",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "ArnLike",
                        Variable = "aws:SourceArn",
                        Values = new[]
                        {
                            bucket.Arn,
                        },
                    },
                },
            },
        },
    });

    var topicTopic = new Aws.Sns.Topic("topicTopic", new()
    {
        Policy = topicPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new()
    {
        Bucket = bucket.Id,
        Topics = new[]
        {
            new Aws.S3.Inputs.BucketNotificationTopicArgs
            {
                TopicArn = topicTopic.Arn,
                Events = new[]
                {
                    "s3:ObjectCreated:*",
                },
                FilterSuffix = ".log",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		topicPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("s3.amazonaws.com"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("SNS:Publish"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("arn:aws:sns:*:*:s3-event-notification-topic"),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("ArnLike"),
							Variable: pulumi.String("aws:SourceArn"),
							Values: pulumi.StringArray{
								bucket.Arn,
							},
						},
					},
				},
			},
		}, nil)
		topicTopic, err := sns.NewTopic(ctx, "topicTopic", &sns.TopicArgs{
			Policy: topicPolicyDocument.ApplyT(func(topicPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &topicPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket: bucket.ID(),
			Topics: s3.BucketNotificationTopicArray{
				&s3.BucketNotificationTopicArgs{
					TopicArn: topicTopic.Arn,
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterSuffix: pulumi.String(".log"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.sns.TopicArgs;
import com.pulumi.aws.s3.BucketNotification;
import com.pulumi.aws.s3.BucketNotificationArgs;
import com.pulumi.aws.s3.inputs.BucketNotificationTopicArgs;
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
        var bucket = new BucketV2("bucket");

        final var topicPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("s3.amazonaws.com")
                    .build())
                .actions("SNS:Publish")
                .resources("arn:aws:sns:*:*:s3-event-notification-topic")
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("ArnLike")
                    .variable("aws:SourceArn")
                    .values(bucket.arn())
                    .build())
                .build())
            .build());

        var topicTopic = new Topic("topicTopic", TopicArgs.builder()        
            .policy(topicPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(topicPolicyDocument -> topicPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var bucketNotification = new BucketNotification("bucketNotification", BucketNotificationArgs.builder()        
            .bucket(bucket.id())
            .topics(BucketNotificationTopicArgs.builder()
                .topicArn(topicTopic.arn())
                .events("s3:ObjectCreated:*")
                .filterSuffix(".log")
                .build())
            .build());

    }
}
```
```yaml
resources:
  topicTopic:
    type: aws:sns:Topic
    properties:
      policy: ${topicPolicyDocument.json}
  bucket:
    type: aws:s3:BucketV2
  bucketNotification:
    type: aws:s3:BucketNotification
    properties:
      bucket: ${bucket.id}
      topics:
        - topicArn: ${topicTopic.arn}
          events:
            - s3:ObjectCreated:*
          filterSuffix: .log
variables:
  topicPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - s3.amazonaws.com
            actions:
              - SNS:Publish
            resources:
              - arn:aws:sns:*:*:s3-event-notification-topic
            conditions:
              - test: ArnLike
                variable: aws:SourceArn
                values:
                  - ${bucket.arn}
```
{{% /example %}}
{{% example %}}
### Add notification configuration to SQS Queue

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.BucketV2("bucket", {});
const queuePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: ["sqs:SendMessage"],
        resources: ["arn:aws:sqs:*:*:s3-event-notification-queue"],
        conditions: [{
            test: "ArnEquals",
            variable: "aws:SourceArn",
            values: [bucket.arn],
        }],
    }],
});
const queueQueue = new aws.sqs.Queue("queueQueue", {policy: queuePolicyDocument.apply(queuePolicyDocument => queuePolicyDocument.json)});
const bucketNotification = new aws.s3.BucketNotification("bucketNotification", {
    bucket: bucket.id,
    queues: [{
        queueArn: queueQueue.arn,
        events: ["s3:ObjectCreated:*"],
        filterSuffix: ".log",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.BucketV2("bucket")
queue_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=["sqs:SendMessage"],
    resources=["arn:aws:sqs:*:*:s3-event-notification-queue"],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="ArnEquals",
        variable="aws:SourceArn",
        values=[bucket.arn],
    )],
)])
queue_queue = aws.sqs.Queue("queueQueue", policy=queue_policy_document.json)
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    queues=[aws.s3.BucketNotificationQueueArgs(
        queue_arn=queue_queue.arn,
        events=["s3:ObjectCreated:*"],
        filter_suffix=".log",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucket = new Aws.S3.BucketV2("bucket");

    var queuePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "*",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "sqs:SendMessage",
                },
                Resources = new[]
                {
                    "arn:aws:sqs:*:*:s3-event-notification-queue",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "ArnEquals",
                        Variable = "aws:SourceArn",
                        Values = new[]
                        {
                            bucket.Arn,
                        },
                    },
                },
            },
        },
    });

    var queueQueue = new Aws.Sqs.Queue("queueQueue", new()
    {
        Policy = queuePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new()
    {
        Bucket = bucket.Id,
        Queues = new[]
        {
            new Aws.S3.Inputs.BucketNotificationQueueArgs
            {
                QueueArn = queueQueue.Arn,
                Events = new[]
                {
                    "s3:ObjectCreated:*",
                },
                FilterSuffix = ".log",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		queuePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("*"),
							Identifiers: pulumi.StringArray{
								pulumi.String("*"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("sqs:SendMessage"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("arn:aws:sqs:*:*:s3-event-notification-queue"),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("ArnEquals"),
							Variable: pulumi.String("aws:SourceArn"),
							Values: pulumi.StringArray{
								bucket.Arn,
							},
						},
					},
				},
			},
		}, nil)
		queueQueue, err := sqs.NewQueue(ctx, "queueQueue", &sqs.QueueArgs{
			Policy: queuePolicyDocument.ApplyT(func(queuePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &queuePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket: bucket.ID(),
			Queues: s3.BucketNotificationQueueArray{
				&s3.BucketNotificationQueueArgs{
					QueueArn: queueQueue.Arn,
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterSuffix: pulumi.String(".log"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.sqs.QueueArgs;
import com.pulumi.aws.s3.BucketNotification;
import com.pulumi.aws.s3.BucketNotificationArgs;
import com.pulumi.aws.s3.inputs.BucketNotificationQueueArgs;
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
        var bucket = new BucketV2("bucket");

        final var queuePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions("sqs:SendMessage")
                .resources("arn:aws:sqs:*:*:s3-event-notification-queue")
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("ArnEquals")
                    .variable("aws:SourceArn")
                    .values(bucket.arn())
                    .build())
                .build())
            .build());

        var queueQueue = new Queue("queueQueue", QueueArgs.builder()        
            .policy(queuePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(queuePolicyDocument -> queuePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var bucketNotification = new BucketNotification("bucketNotification", BucketNotificationArgs.builder()        
            .bucket(bucket.id())
            .queues(BucketNotificationQueueArgs.builder()
                .queueArn(queueQueue.arn())
                .events("s3:ObjectCreated:*")
                .filterSuffix(".log")
                .build())
            .build());

    }
}
```
```yaml
resources:
  queueQueue:
    type: aws:sqs:Queue
    properties:
      policy: ${queuePolicyDocument.json}
  bucket:
    type: aws:s3:BucketV2
  bucketNotification:
    type: aws:s3:BucketNotification
    properties:
      bucket: ${bucket.id}
      queues:
        - queueArn: ${queueQueue.arn}
          events:
            - s3:ObjectCreated:*
          filterSuffix: .log
variables:
  queuePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: '*'
                identifiers:
                  - '*'
            actions:
              - sqs:SendMessage
            resources:
              - arn:aws:sqs:*:*:s3-event-notification-queue
            conditions:
              - test: ArnEquals
                variable: aws:SourceArn
                values:
                  - ${bucket.arn}
```
{{% /example %}}
{{% example %}}
### Add notification configuration to Lambda Function

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["lambda.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const iamForLambda = new aws.iam.Role("iamForLambda", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const func = new aws.lambda.Function("func", {
    code: new pulumi.asset.FileArchive("your-function.zip"),
    role: iamForLambda.arn,
    handler: "exports.example",
    runtime: "go1.x",
});
const bucket = new aws.s3.BucketV2("bucket", {});
const allowBucket = new aws.lambda.Permission("allowBucket", {
    action: "lambda:InvokeFunction",
    "function": func.arn,
    principal: "s3.amazonaws.com",
    sourceArn: bucket.arn,
});
const bucketNotification = new aws.s3.BucketNotification("bucketNotification", {
    bucket: bucket.id,
    lambdaFunctions: [{
        lambdaFunctionArn: func.arn,
        events: ["s3:ObjectCreated:*"],
        filterPrefix: "AWSLogs/",
        filterSuffix: ".log",
    }],
}, {
    dependsOn: [allowBucket],
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["lambda.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
iam_for_lambda = aws.iam.Role("iamForLambda", assume_role_policy=assume_role.json)
func = aws.lambda_.Function("func",
    code=pulumi.FileArchive("your-function.zip"),
    role=iam_for_lambda.arn,
    handler="exports.example",
    runtime="go1.x")
bucket = aws.s3.BucketV2("bucket")
allow_bucket = aws.lambda_.Permission("allowBucket",
    action="lambda:InvokeFunction",
    function=func.arn,
    principal="s3.amazonaws.com",
    source_arn=bucket.arn)
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    lambda_functions=[aws.s3.BucketNotificationLambdaFunctionArgs(
        lambda_function_arn=func.arn,
        events=["s3:ObjectCreated:*"],
        filter_prefix="AWSLogs/",
        filter_suffix=".log",
    )],
    opts=pulumi.ResourceOptions(depends_on=[allow_bucket]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
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
                            "lambda.amazonaws.com",
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

    var iamForLambda = new Aws.Iam.Role("iamForLambda", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var func = new Aws.Lambda.Function("func", new()
    {
        Code = new FileArchive("your-function.zip"),
        Role = iamForLambda.Arn,
        Handler = "exports.example",
        Runtime = "go1.x",
    });

    var bucket = new Aws.S3.BucketV2("bucket");

    var allowBucket = new Aws.Lambda.Permission("allowBucket", new()
    {
        Action = "lambda:InvokeFunction",
        Function = func.Arn,
        Principal = "s3.amazonaws.com",
        SourceArn = bucket.Arn,
    });

    var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new()
    {
        Bucket = bucket.Id,
        LambdaFunctions = new[]
        {
            new Aws.S3.Inputs.BucketNotificationLambdaFunctionArgs
            {
                LambdaFunctionArn = func.Arn,
                Events = new[]
                {
                    "s3:ObjectCreated:*",
                },
                FilterPrefix = "AWSLogs/",
                FilterSuffix = ".log",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            allowBucket,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"lambda.amazonaws.com",
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
		iamForLambda, err := iam.NewRole(ctx, "iamForLambda", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "func", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("your-function.zip"),
			Role:    iamForLambda.Arn,
			Handler: pulumi.String("exports.example"),
			Runtime: pulumi.String("go1.x"),
		})
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		allowBucket, err := lambda.NewPermission(ctx, "allowBucket", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  _func.Arn,
			Principal: pulumi.String("s3.amazonaws.com"),
			SourceArn: bucket.Arn,
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket: bucket.ID(),
			LambdaFunctions: s3.BucketNotificationLambdaFunctionArray{
				&s3.BucketNotificationLambdaFunctionArgs{
					LambdaFunctionArn: _func.Arn,
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterPrefix: pulumi.String("AWSLogs/"),
					FilterSuffix: pulumi.String(".log"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			allowBucket,
		}))
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.s3.BucketNotification;
import com.pulumi.aws.s3.BucketNotificationArgs;
import com.pulumi.aws.s3.inputs.BucketNotificationLambdaFunctionArgs;
import com.pulumi.resources.CustomResourceOptions;
import com.pulumi.asset.FileArchive;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("lambda.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var iamForLambda = new Role("iamForLambda", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var func = new Function("func", FunctionArgs.builder()        
            .code(new FileArchive("your-function.zip"))
            .role(iamForLambda.arn())
            .handler("exports.example")
            .runtime("go1.x")
            .build());

        var bucket = new BucketV2("bucket");

        var allowBucket = new Permission("allowBucket", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(func.arn())
            .principal("s3.amazonaws.com")
            .sourceArn(bucket.arn())
            .build());

        var bucketNotification = new BucketNotification("bucketNotification", BucketNotificationArgs.builder()        
            .bucket(bucket.id())
            .lambdaFunctions(BucketNotificationLambdaFunctionArgs.builder()
                .lambdaFunctionArn(func.arn())
                .events("s3:ObjectCreated:*")
                .filterPrefix("AWSLogs/")
                .filterSuffix(".log")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(allowBucket)
                .build());

    }
}
```
```yaml
resources:
  iamForLambda:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  allowBucket:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${func.arn}
      principal: s3.amazonaws.com
      sourceArn: ${bucket.arn}
  func:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: your-function.zip
      role: ${iamForLambda.arn}
      handler: exports.example
      runtime: go1.x
  bucket:
    type: aws:s3:BucketV2
  bucketNotification:
    type: aws:s3:BucketNotification
    properties:
      bucket: ${bucket.id}
      lambdaFunctions:
        - lambdaFunctionArn: ${func.arn}
          events:
            - s3:ObjectCreated:*
          filterPrefix: AWSLogs/
          filterSuffix: .log
    options:
      dependson:
        - ${allowBucket}
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
                  - lambda.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Trigger multiple Lambda functions

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.s3.BucketNotification;
import com.pulumi.aws.s3.BucketNotificationArgs;
import com.pulumi.aws.s3.inputs.BucketNotificationLambdaFunctionArgs;
import com.pulumi.resources.CustomResourceOptions;
import com.pulumi.asset.FileArchive;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .effect("Allow")
            .principals(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .actions("sts:AssumeRole")
            .build());

        var iamForLambda = new Role("iamForLambda", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var func1 = new Function("func1", FunctionArgs.builder()        
            .code(new FileArchive("your-function1.zip"))
            .role(iamForLambda.arn())
            .handler("exports.example")
            .runtime("go1.x")
            .build());

        var bucket = new BucketV2("bucket");

        var allowBucket1 = new Permission("allowBucket1", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(func1.arn())
            .principal("s3.amazonaws.com")
            .sourceArn(bucket.arn())
            .build());

        var func2 = new Function("func2", FunctionArgs.builder()        
            .code(new FileArchive("your-function2.zip"))
            .role(iamForLambda.arn())
            .handler("exports.example")
            .build());

        var allowBucket2 = new Permission("allowBucket2", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(func2.arn())
            .principal("s3.amazonaws.com")
            .sourceArn(bucket.arn())
            .build());

        var bucketNotification = new BucketNotification("bucketNotification", BucketNotificationArgs.builder()        
            .bucket(bucket.id())
            .lambdaFunctions(            
                BucketNotificationLambdaFunctionArgs.builder()
                    .lambdaFunctionArn(func1.arn())
                    .events("s3:ObjectCreated:*")
                    .filterPrefix("AWSLogs/")
                    .filterSuffix(".log")
                    .build(),
                BucketNotificationLambdaFunctionArgs.builder()
                    .lambdaFunctionArn(func2.arn())
                    .events("s3:ObjectCreated:*")
                    .filterPrefix("OtherLogs/")
                    .filterSuffix(".log")
                    .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    allowBucket1,
                    allowBucket2)
                .build());

    }
}
```
```yaml
resources:
  iamForLambda:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  allowBucket1:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${func1.arn}
      principal: s3.amazonaws.com
      sourceArn: ${bucket.arn}
  func1:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: your-function1.zip
      role: ${iamForLambda.arn}
      handler: exports.example
      runtime: go1.x
  allowBucket2:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${func2.arn}
      principal: s3.amazonaws.com
      sourceArn: ${bucket.arn}
  func2:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: your-function2.zip
      role: ${iamForLambda.arn}
      handler: exports.example
  bucket:
    type: aws:s3:BucketV2
  bucketNotification:
    type: aws:s3:BucketNotification
    properties:
      bucket: ${bucket.id}
      lambdaFunctions:
        - lambdaFunctionArn: ${func1.arn}
          events:
            - s3:ObjectCreated:*
          filterPrefix: AWSLogs/
          filterSuffix: .log
        - lambdaFunctionArn: ${func2.arn}
          events:
            - s3:ObjectCreated:*
          filterPrefix: OtherLogs/
          filterSuffix: .log
    options:
      dependson:
        - ${allowBucket1}
        - ${allowBucket2}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        effect: Allow
        principals:
          - type: Service
            identifiers:
              - lambda.amazonaws.com
        actions:
          - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Add multiple notification configurations to SQS Queue

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.BucketV2("bucket", {});
const queuePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: ["sqs:SendMessage"],
        resources: ["arn:aws:sqs:*:*:s3-event-notification-queue"],
        conditions: [{
            test: "ArnEquals",
            variable: "aws:SourceArn",
            values: [bucket.arn],
        }],
    }],
});
const queueQueue = new aws.sqs.Queue("queueQueue", {policy: queuePolicyDocument.apply(queuePolicyDocument => queuePolicyDocument.json)});
const bucketNotification = new aws.s3.BucketNotification("bucketNotification", {
    bucket: bucket.id,
    queues: [
        {
            id: "image-upload-event",
            queueArn: queueQueue.arn,
            events: ["s3:ObjectCreated:*"],
            filterPrefix: "images/",
        },
        {
            id: "video-upload-event",
            queueArn: queueQueue.arn,
            events: ["s3:ObjectCreated:*"],
            filterPrefix: "videos/",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.BucketV2("bucket")
queue_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=["sqs:SendMessage"],
    resources=["arn:aws:sqs:*:*:s3-event-notification-queue"],
    conditions=[aws.iam.GetPolicyDocumentStatementConditionArgs(
        test="ArnEquals",
        variable="aws:SourceArn",
        values=[bucket.arn],
    )],
)])
queue_queue = aws.sqs.Queue("queueQueue", policy=queue_policy_document.json)
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    queues=[
        aws.s3.BucketNotificationQueueArgs(
            id="image-upload-event",
            queue_arn=queue_queue.arn,
            events=["s3:ObjectCreated:*"],
            filter_prefix="images/",
        ),
        aws.s3.BucketNotificationQueueArgs(
            id="video-upload-event",
            queue_arn=queue_queue.arn,
            events=["s3:ObjectCreated:*"],
            filter_prefix="videos/",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucket = new Aws.S3.BucketV2("bucket");

    var queuePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "*",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "sqs:SendMessage",
                },
                Resources = new[]
                {
                    "arn:aws:sqs:*:*:s3-event-notification-queue",
                },
                Conditions = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementConditionInputArgs
                    {
                        Test = "ArnEquals",
                        Variable = "aws:SourceArn",
                        Values = new[]
                        {
                            bucket.Arn,
                        },
                    },
                },
            },
        },
    });

    var queueQueue = new Aws.Sqs.Queue("queueQueue", new()
    {
        Policy = queuePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new()
    {
        Bucket = bucket.Id,
        Queues = new[]
        {
            new Aws.S3.Inputs.BucketNotificationQueueArgs
            {
                Id = "image-upload-event",
                QueueArn = queueQueue.Arn,
                Events = new[]
                {
                    "s3:ObjectCreated:*",
                },
                FilterPrefix = "images/",
            },
            new Aws.S3.Inputs.BucketNotificationQueueArgs
            {
                Id = "video-upload-event",
                QueueArn = queueQueue.Arn,
                Events = new[]
                {
                    "s3:ObjectCreated:*",
                },
                FilterPrefix = "videos/",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		queuePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("*"),
							Identifiers: pulumi.StringArray{
								pulumi.String("*"),
							},
						},
					},
					Actions: pulumi.StringArray{
						pulumi.String("sqs:SendMessage"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("arn:aws:sqs:*:*:s3-event-notification-queue"),
					},
					Conditions: iam.GetPolicyDocumentStatementConditionArray{
						&iam.GetPolicyDocumentStatementConditionArgs{
							Test:     pulumi.String("ArnEquals"),
							Variable: pulumi.String("aws:SourceArn"),
							Values: pulumi.StringArray{
								bucket.Arn,
							},
						},
					},
				},
			},
		}, nil)
		queueQueue, err := sqs.NewQueue(ctx, "queueQueue", &sqs.QueueArgs{
			Policy: queuePolicyDocument.ApplyT(func(queuePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &queuePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket: bucket.ID(),
			Queues: s3.BucketNotificationQueueArray{
				&s3.BucketNotificationQueueArgs{
					Id:       pulumi.String("image-upload-event"),
					QueueArn: queueQueue.Arn,
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterPrefix: pulumi.String("images/"),
				},
				&s3.BucketNotificationQueueArgs{
					Id:       pulumi.String("video-upload-event"),
					QueueArn: queueQueue.Arn,
					Events: pulumi.StringArray{
						pulumi.String("s3:ObjectCreated:*"),
					},
					FilterPrefix: pulumi.String("videos/"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.sqs.QueueArgs;
import com.pulumi.aws.s3.BucketNotification;
import com.pulumi.aws.s3.BucketNotificationArgs;
import com.pulumi.aws.s3.inputs.BucketNotificationQueueArgs;
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
        var bucket = new BucketV2("bucket");

        final var queuePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions("sqs:SendMessage")
                .resources("arn:aws:sqs:*:*:s3-event-notification-queue")
                .conditions(GetPolicyDocumentStatementConditionArgs.builder()
                    .test("ArnEquals")
                    .variable("aws:SourceArn")
                    .values(bucket.arn())
                    .build())
                .build())
            .build());

        var queueQueue = new Queue("queueQueue", QueueArgs.builder()        
            .policy(queuePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(queuePolicyDocument -> queuePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var bucketNotification = new BucketNotification("bucketNotification", BucketNotificationArgs.builder()        
            .bucket(bucket.id())
            .queues(            
                BucketNotificationQueueArgs.builder()
                    .id("image-upload-event")
                    .queueArn(queueQueue.arn())
                    .events("s3:ObjectCreated:*")
                    .filterPrefix("images/")
                    .build(),
                BucketNotificationQueueArgs.builder()
                    .id("video-upload-event")
                    .queueArn(queueQueue.arn())
                    .events("s3:ObjectCreated:*")
                    .filterPrefix("videos/")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  queueQueue:
    type: aws:sqs:Queue
    properties:
      policy: ${queuePolicyDocument.json}
  bucket:
    type: aws:s3:BucketV2
  bucketNotification:
    type: aws:s3:BucketNotification
    properties:
      bucket: ${bucket.id}
      queues:
        - id: image-upload-event
          queueArn: ${queueQueue.arn}
          events:
            - s3:ObjectCreated:*
          filterPrefix: images/
        - id: video-upload-event
          queueArn: ${queueQueue.arn}
          events:
            - s3:ObjectCreated:*
          filterPrefix: videos/
variables:
  queuePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: '*'
                identifiers:
                  - '*'
            actions:
              - sqs:SendMessage
            resources:
              - arn:aws:sqs:*:*:s3-event-notification-queue
            conditions:
              - test: ArnEquals
                variable: aws:SourceArn
                values:
                  - ${bucket.arn}
```

For JSON syntax, use an array instead of defining the `queue` key twice.

```typescript
import * as pulumi from "@pulumi/pulumi";
```
```python
import pulumi
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() => 
{
});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
```yaml
{}
```
{{% /example %}}
{{% example %}}
### Emit events to EventBridge

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.BucketV2("bucket", {});
const bucketNotification = new aws.s3.BucketNotification("bucketNotification", {
    bucket: bucket.id,
    eventbridge: true,
});
```
```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.BucketV2("bucket")
bucket_notification = aws.s3.BucketNotification("bucketNotification",
    bucket=bucket.id,
    eventbridge=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucket = new Aws.S3.BucketV2("bucket");

    var bucketNotification = new Aws.S3.BucketNotification("bucketNotification", new()
    {
        Bucket = bucket.Id,
        Eventbridge = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketNotification(ctx, "bucketNotification", &s3.BucketNotificationArgs{
			Bucket:      bucket.ID(),
			Eventbridge: pulumi.Bool(true),
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
import com.pulumi.aws.s3.BucketNotification;
import com.pulumi.aws.s3.BucketNotificationArgs;
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
        var bucket = new BucketV2("bucket");

        var bucketNotification = new BucketNotification("bucketNotification", BucketNotificationArgs.builder()        
            .bucket(bucket.id())
            .eventbridge(true)
            .build());

    }
}
```
```yaml
resources:
  bucket:
    type: aws:s3:BucketV2
  bucketNotification:
    type: aws:s3:BucketNotification
    properties:
      bucket: ${bucket.id}
      eventbridge: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket notification using the `bucket`. For example:

```sh
 $ pulumi import aws:s3/bucketNotification:BucketNotification bucket_notification bucket-name
```
 