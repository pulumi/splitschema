Resource for managing an AWS Chime SDK Media Pipelines Media Insights Pipeline Configuration.
Consult the [Call analytics developer guide](https://docs.aws.amazon.com/chime-sdk/latest/dg/call-analytics.html) for more detailed information about usage.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kinesis.Stream("example", {shardCount: 2});
const mediaPipelinesAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["mediapipelines.chime.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const callAnalyticsRole = new aws.iam.Role("callAnalyticsRole", {assumeRolePolicy: mediaPipelinesAssumeRole.then(mediaPipelinesAssumeRole => mediaPipelinesAssumeRole.json)});
const myConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration", {
    resourceAccessRoleArn: callAnalyticsRole.arn,
    elements: [
        {
            type: "AmazonTranscribeCallAnalyticsProcessor",
            amazonTranscribeCallAnalyticsProcessorConfiguration: {
                languageCode: "en-US",
            },
        },
        {
            type: "KinesisDataStreamSink",
            kinesisDataStreamSinkConfiguration: {
                insightsTarget: example.arn,
            },
        },
    ],
    tags: {
        Key1: "Value1",
        Key2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kinesis.Stream("example", shard_count=2)
media_pipelines_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["mediapipelines.chime.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
call_analytics_role = aws.iam.Role("callAnalyticsRole", assume_role_policy=media_pipelines_assume_role.json)
my_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration",
    resource_access_role_arn=call_analytics_role.arn,
    elements=[
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="AmazonTranscribeCallAnalyticsProcessor",
            amazon_transcribe_call_analytics_processor_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs(
                language_code="en-US",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="KinesisDataStreamSink",
            kinesis_data_stream_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs(
                insights_target=example.arn,
            ),
        ),
    ],
    tags={
        "Key1": "Value1",
        "Key2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kinesis.Stream("example", new()
    {
        ShardCount = 2,
    });

    var mediaPipelinesAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "mediapipelines.chime.amazonaws.com",
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

    var callAnalyticsRole = new Aws.Iam.Role("callAnalyticsRole", new()
    {
        AssumeRolePolicy = mediaPipelinesAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var myConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("myConfiguration", new()
    {
        ResourceAccessRoleArn = callAnalyticsRole.Arn,
        Elements = new[]
        {
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "AmazonTranscribeCallAnalyticsProcessor",
                AmazonTranscribeCallAnalyticsProcessorConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs
                {
                    LanguageCode = "en-US",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "KinesisDataStreamSink",
                KinesisDataStreamSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs
                {
                    InsightsTarget = example.Arn,
                },
            },
        },
        Tags = 
        {
            { "Key1", "Value1" },
            { "Key2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := kinesis.NewStream(ctx, "example", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(2),
		})
		if err != nil {
			return err
		}
		mediaPipelinesAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"mediapipelines.chime.amazonaws.com",
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
		callAnalyticsRole, err := iam.NewRole(ctx, "callAnalyticsRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(mediaPipelinesAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "myConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: callAnalyticsRole.Arn,
			Elements: chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArray{
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("AmazonTranscribeCallAnalyticsProcessor"),
					AmazonTranscribeCallAnalyticsProcessorConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs{
						LanguageCode: pulumi.String("en-US"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("KinesisDataStreamSink"),
					KinesisDataStreamSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs{
						InsightsTarget: example.Arn,
					},
				},
			},
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("Value1"),
				"Key2": pulumi.String("Value2"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
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
            .shardCount(2)
            .build());

        final var mediaPipelinesAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("mediapipelines.chime.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var callAnalyticsRole = new Role("callAnalyticsRole", RoleArgs.builder()        
            .assumeRolePolicy(mediaPipelinesAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var myConfiguration = new MediaInsightsPipelineConfiguration("myConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(callAnalyticsRole.arn())
            .elements(            
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("AmazonTranscribeCallAnalyticsProcessor")
                    .amazonTranscribeCallAnalyticsProcessorConfiguration(MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs.builder()
                        .languageCode("en-US")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("KinesisDataStreamSink")
                    .kinesisDataStreamSinkConfiguration(MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs.builder()
                        .insightsTarget(example.arn())
                        .build())
                    .build())
            .tags(Map.ofEntries(
                Map.entry("Key1", "Value1"),
                Map.entry("Key2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  myConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${callAnalyticsRole.arn}
      elements:
        - type: AmazonTranscribeCallAnalyticsProcessor
          amazonTranscribeCallAnalyticsProcessorConfiguration:
            languageCode: en-US
        - type: KinesisDataStreamSink
          kinesisDataStreamSinkConfiguration:
            insightsTarget: ${example.arn}
      tags:
        Key1: Value1
        Key2: Value2
  example:
    type: aws:kinesis:Stream
    properties:
      shardCount: 2
  callAnalyticsRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${mediaPipelinesAssumeRole.json}
variables:
  mediaPipelinesAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - mediapipelines.chime.amazonaws.com
            actions:
              - sts:AssumeRole
```

- The required policies on `call_analytics_role` will vary based on the selected processors. See [Call analytics resource access role](https://docs.aws.amazon.com/chime-sdk/latest/dg/ca-resource-access-role.html) for directions on choosing appropriate policies.
{{% /example %}}
{{% example %}}
### Transcribe Call Analytics processor usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const transcribeAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["transcribe.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const postCallRole = new aws.iam.Role("postCallRole", {assumeRolePolicy: transcribeAssumeRole.then(transcribeAssumeRole => transcribeAssumeRole.json)});
const myConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration", {
    resourceAccessRoleArn: aws_iam_role.example.arn,
    elements: [
        {
            type: "AmazonTranscribeCallAnalyticsProcessor",
            amazonTranscribeCallAnalyticsProcessorConfiguration: {
                callAnalyticsStreamCategories: [
                    "category_1",
                    "category_2",
                ],
                contentRedactionType: "PII",
                enablePartialResultsStabilization: true,
                filterPartialResults: true,
                languageCode: "en-US",
                languageModelName: "MyLanguageModel",
                partialResultsStability: "high",
                piiEntityTypes: "ADDRESS,BANK_ACCOUNT_NUMBER",
                postCallAnalyticsSettings: {
                    contentRedactionOutput: "redacted",
                    dataAccessRoleArn: postCallRole.arn,
                    outputEncryptionKmsKeyId: "MyKmsKeyId",
                    outputLocation: "s3://MyBucket",
                },
                vocabularyFilterMethod: "mask",
                vocabularyFilterName: "MyVocabularyFilter",
                vocabularyName: "MyVocabulary",
            },
        },
        {
            type: "KinesisDataStreamSink",
            kinesisDataStreamSinkConfiguration: {
                insightsTarget: aws_kinesis_stream.example.arn,
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

transcribe_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["transcribe.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
post_call_role = aws.iam.Role("postCallRole", assume_role_policy=transcribe_assume_role.json)
my_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration",
    resource_access_role_arn=aws_iam_role["example"]["arn"],
    elements=[
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="AmazonTranscribeCallAnalyticsProcessor",
            amazon_transcribe_call_analytics_processor_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs(
                call_analytics_stream_categories=[
                    "category_1",
                    "category_2",
                ],
                content_redaction_type="PII",
                enable_partial_results_stabilization=True,
                filter_partial_results=True,
                language_code="en-US",
                language_model_name="MyLanguageModel",
                partial_results_stability="high",
                pii_entity_types="ADDRESS,BANK_ACCOUNT_NUMBER",
                post_call_analytics_settings=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsArgs(
                    content_redaction_output="redacted",
                    data_access_role_arn=post_call_role.arn,
                    output_encryption_kms_key_id="MyKmsKeyId",
                    output_location="s3://MyBucket",
                ),
                vocabulary_filter_method="mask",
                vocabulary_filter_name="MyVocabularyFilter",
                vocabulary_name="MyVocabulary",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="KinesisDataStreamSink",
            kinesis_data_stream_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs(
                insights_target=aws_kinesis_stream["example"]["arn"],
            ),
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
    var transcribeAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "transcribe.amazonaws.com",
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

    var postCallRole = new Aws.Iam.Role("postCallRole", new()
    {
        AssumeRolePolicy = transcribeAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var myConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("myConfiguration", new()
    {
        ResourceAccessRoleArn = aws_iam_role.Example.Arn,
        Elements = new[]
        {
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "AmazonTranscribeCallAnalyticsProcessor",
                AmazonTranscribeCallAnalyticsProcessorConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs
                {
                    CallAnalyticsStreamCategories = new[]
                    {
                        "category_1",
                        "category_2",
                    },
                    ContentRedactionType = "PII",
                    EnablePartialResultsStabilization = true,
                    FilterPartialResults = true,
                    LanguageCode = "en-US",
                    LanguageModelName = "MyLanguageModel",
                    PartialResultsStability = "high",
                    PiiEntityTypes = "ADDRESS,BANK_ACCOUNT_NUMBER",
                    PostCallAnalyticsSettings = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsArgs
                    {
                        ContentRedactionOutput = "redacted",
                        DataAccessRoleArn = postCallRole.Arn,
                        OutputEncryptionKmsKeyId = "MyKmsKeyId",
                        OutputLocation = "s3://MyBucket",
                    },
                    VocabularyFilterMethod = "mask",
                    VocabularyFilterName = "MyVocabularyFilter",
                    VocabularyName = "MyVocabulary",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "KinesisDataStreamSink",
                KinesisDataStreamSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs
                {
                    InsightsTarget = aws_kinesis_stream.Example.Arn,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		transcribeAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"transcribe.amazonaws.com",
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
		postCallRole, err := iam.NewRole(ctx, "postCallRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(transcribeAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "myConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Elements: chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArray{
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("AmazonTranscribeCallAnalyticsProcessor"),
					AmazonTranscribeCallAnalyticsProcessorConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs{
						CallAnalyticsStreamCategories: pulumi.StringArray{
							pulumi.String("category_1"),
							pulumi.String("category_2"),
						},
						ContentRedactionType:              pulumi.String("PII"),
						EnablePartialResultsStabilization: pulumi.Bool(true),
						FilterPartialResults:              pulumi.Bool(true),
						LanguageCode:                      pulumi.String("en-US"),
						LanguageModelName:                 pulumi.String("MyLanguageModel"),
						PartialResultsStability:           pulumi.String("high"),
						PiiEntityTypes:                    pulumi.String("ADDRESS,BANK_ACCOUNT_NUMBER"),
						PostCallAnalyticsSettings: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsArgs{
							ContentRedactionOutput:   pulumi.String("redacted"),
							DataAccessRoleArn:        postCallRole.Arn,
							OutputEncryptionKmsKeyId: pulumi.String("MyKmsKeyId"),
							OutputLocation:           pulumi.String("s3://MyBucket"),
						},
						VocabularyFilterMethod: pulumi.String("mask"),
						VocabularyFilterName:   pulumi.String("MyVocabularyFilter"),
						VocabularyName:         pulumi.String("MyVocabulary"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("KinesisDataStreamSink"),
					KinesisDataStreamSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs{
						InsightsTarget: pulumi.Any(aws_kinesis_stream.Example.Arn),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
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
        final var transcribeAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("transcribe.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var postCallRole = new Role("postCallRole", RoleArgs.builder()        
            .assumeRolePolicy(transcribeAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var myConfiguration = new MediaInsightsPipelineConfiguration("myConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(aws_iam_role.example().arn())
            .elements(            
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("AmazonTranscribeCallAnalyticsProcessor")
                    .amazonTranscribeCallAnalyticsProcessorConfiguration(MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs.builder()
                        .callAnalyticsStreamCategories(                        
                            "category_1",
                            "category_2")
                        .contentRedactionType("PII")
                        .enablePartialResultsStabilization(true)
                        .filterPartialResults(true)
                        .languageCode("en-US")
                        .languageModelName("MyLanguageModel")
                        .partialResultsStability("high")
                        .piiEntityTypes("ADDRESS,BANK_ACCOUNT_NUMBER")
                        .postCallAnalyticsSettings(MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettingsArgs.builder()
                            .contentRedactionOutput("redacted")
                            .dataAccessRoleArn(postCallRole.arn())
                            .outputEncryptionKmsKeyId("MyKmsKeyId")
                            .outputLocation("s3://MyBucket")
                            .build())
                        .vocabularyFilterMethod("mask")
                        .vocabularyFilterName("MyVocabularyFilter")
                        .vocabularyName("MyVocabulary")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("KinesisDataStreamSink")
                    .kinesisDataStreamSinkConfiguration(MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs.builder()
                        .insightsTarget(aws_kinesis_stream.example().arn())
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  myConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${aws_iam_role.example.arn}
      elements:
        - type: AmazonTranscribeCallAnalyticsProcessor
          amazonTranscribeCallAnalyticsProcessorConfiguration:
            callAnalyticsStreamCategories:
              - category_1
              - category_2
            contentRedactionType: PII
            enablePartialResultsStabilization: true
            filterPartialResults: true
            languageCode: en-US
            languageModelName: MyLanguageModel
            partialResultsStability: high
            piiEntityTypes: ADDRESS,BANK_ACCOUNT_NUMBER
            postCallAnalyticsSettings:
              contentRedactionOutput: redacted
              dataAccessRoleArn: ${postCallRole.arn}
              outputEncryptionKmsKeyId: MyKmsKeyId
              outputLocation: s3://MyBucket
            vocabularyFilterMethod: mask
            vocabularyFilterName: MyVocabularyFilter
            vocabularyName: MyVocabulary
        - type: KinesisDataStreamSink
          kinesisDataStreamSinkConfiguration:
            insightsTarget: ${aws_kinesis_stream.example.arn}
  postCallRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${transcribeAssumeRole.json}
variables:
  transcribeAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - transcribe.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Real time alerts usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration", {
    resourceAccessRoleArn: aws_iam_role.call_analytics_role.arn,
    elements: [
        {
            type: "AmazonTranscribeCallAnalyticsProcessor",
            amazonTranscribeCallAnalyticsProcessorConfiguration: {
                languageCode: "en-US",
            },
        },
        {
            type: "KinesisDataStreamSink",
            kinesisDataStreamSinkConfiguration: {
                insightsTarget: aws_kinesis_stream.example.arn,
            },
        },
    ],
    realTimeAlertConfiguration: {
        disabled: false,
        rules: [
            {
                type: "IssueDetection",
                issueDetectionConfiguration: {
                    ruleName: "MyIssueDetectionRule",
                },
            },
            {
                type: "KeywordMatch",
                keywordMatchConfiguration: {
                    keywords: [
                        "keyword1",
                        "keyword2",
                    ],
                    negate: false,
                    ruleName: "MyKeywordMatchRule",
                },
            },
            {
                type: "Sentiment",
                sentimentConfiguration: {
                    ruleName: "MySentimentRule",
                    sentimentType: "NEGATIVE",
                    timePeriod: 60,
                },
            },
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

my_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration",
    resource_access_role_arn=aws_iam_role["call_analytics_role"]["arn"],
    elements=[
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="AmazonTranscribeCallAnalyticsProcessor",
            amazon_transcribe_call_analytics_processor_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs(
                language_code="en-US",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="KinesisDataStreamSink",
            kinesis_data_stream_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs(
                insights_target=aws_kinesis_stream["example"]["arn"],
            ),
        ),
    ],
    real_time_alert_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationArgs(
        disabled=False,
        rules=[
            aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs(
                type="IssueDetection",
                issue_detection_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleIssueDetectionConfigurationArgs(
                    rule_name="MyIssueDetectionRule",
                ),
            ),
            aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs(
                type="KeywordMatch",
                keyword_match_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleKeywordMatchConfigurationArgs(
                    keywords=[
                        "keyword1",
                        "keyword2",
                    ],
                    negate=False,
                    rule_name="MyKeywordMatchRule",
                ),
            ),
            aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs(
                type="Sentiment",
                sentiment_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleSentimentConfigurationArgs(
                    rule_name="MySentimentRule",
                    sentiment_type="NEGATIVE",
                    time_period=60,
                ),
            ),
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("myConfiguration", new()
    {
        ResourceAccessRoleArn = aws_iam_role.Call_analytics_role.Arn,
        Elements = new[]
        {
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "AmazonTranscribeCallAnalyticsProcessor",
                AmazonTranscribeCallAnalyticsProcessorConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs
                {
                    LanguageCode = "en-US",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "KinesisDataStreamSink",
                KinesisDataStreamSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs
                {
                    InsightsTarget = aws_kinesis_stream.Example.Arn,
                },
            },
        },
        RealTimeAlertConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationArgs
        {
            Disabled = false,
            Rules = new[]
            {
                new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs
                {
                    Type = "IssueDetection",
                    IssueDetectionConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleIssueDetectionConfigurationArgs
                    {
                        RuleName = "MyIssueDetectionRule",
                    },
                },
                new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs
                {
                    Type = "KeywordMatch",
                    KeywordMatchConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleKeywordMatchConfigurationArgs
                    {
                        Keywords = new[]
                        {
                            "keyword1",
                            "keyword2",
                        },
                        Negate = false,
                        RuleName = "MyKeywordMatchRule",
                    },
                },
                new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs
                {
                    Type = "Sentiment",
                    SentimentConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleSentimentConfigurationArgs
                    {
                        RuleName = "MySentimentRule",
                        SentimentType = "NEGATIVE",
                        TimePeriod = 60,
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "myConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: pulumi.Any(aws_iam_role.Call_analytics_role.Arn),
			Elements: chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArray{
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("AmazonTranscribeCallAnalyticsProcessor"),
					AmazonTranscribeCallAnalyticsProcessorConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs{
						LanguageCode: pulumi.String("en-US"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("KinesisDataStreamSink"),
					KinesisDataStreamSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs{
						InsightsTarget: pulumi.Any(aws_kinesis_stream.Example.Arn),
					},
				},
			},
			RealTimeAlertConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationArgs{
				Disabled: pulumi.Bool(false),
				Rules: chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArray{
					&chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs{
						Type: pulumi.String("IssueDetection"),
						IssueDetectionConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleIssueDetectionConfigurationArgs{
							RuleName: pulumi.String("MyIssueDetectionRule"),
						},
					},
					&chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs{
						Type: pulumi.String("KeywordMatch"),
						KeywordMatchConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleKeywordMatchConfigurationArgs{
							Keywords: pulumi.StringArray{
								pulumi.String("keyword1"),
								pulumi.String("keyword2"),
							},
							Negate:   pulumi.Bool(false),
							RuleName: pulumi.String("MyKeywordMatchRule"),
						},
					},
					&chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs{
						Type: pulumi.String("Sentiment"),
						SentimentConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleSentimentConfigurationArgs{
							RuleName:      pulumi.String("MySentimentRule"),
							SentimentType: pulumi.String("NEGATIVE"),
							TimePeriod:    pulumi.Int(60),
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
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationRealTimeAlertConfigurationArgs;
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
        var myConfiguration = new MediaInsightsPipelineConfiguration("myConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(aws_iam_role.call_analytics_role().arn())
            .elements(            
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("AmazonTranscribeCallAnalyticsProcessor")
                    .amazonTranscribeCallAnalyticsProcessorConfiguration(MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs.builder()
                        .languageCode("en-US")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("KinesisDataStreamSink")
                    .kinesisDataStreamSinkConfiguration(MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs.builder()
                        .insightsTarget(aws_kinesis_stream.example().arn())
                        .build())
                    .build())
            .realTimeAlertConfiguration(MediaInsightsPipelineConfigurationRealTimeAlertConfigurationArgs.builder()
                .disabled(false)
                .rules(                
                    MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs.builder()
                        .type("IssueDetection")
                        .issueDetectionConfiguration(MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleIssueDetectionConfigurationArgs.builder()
                            .ruleName("MyIssueDetectionRule")
                            .build())
                        .build(),
                    MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs.builder()
                        .type("KeywordMatch")
                        .keywordMatchConfiguration(MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleKeywordMatchConfigurationArgs.builder()
                            .keywords(                            
                                "keyword1",
                                "keyword2")
                            .negate(false)
                            .ruleName("MyKeywordMatchRule")
                            .build())
                        .build(),
                    MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleArgs.builder()
                        .type("Sentiment")
                        .sentimentConfiguration(MediaInsightsPipelineConfigurationRealTimeAlertConfigurationRuleSentimentConfigurationArgs.builder()
                            .ruleName("MySentimentRule")
                            .sentimentType("NEGATIVE")
                            .timePeriod(60)
                            .build())
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  myConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${aws_iam_role.call_analytics_role.arn}
      elements:
        - type: AmazonTranscribeCallAnalyticsProcessor
          amazonTranscribeCallAnalyticsProcessorConfiguration:
            languageCode: en-US
        - type: KinesisDataStreamSink
          kinesisDataStreamSinkConfiguration:
            insightsTarget: ${aws_kinesis_stream.example.arn}
      realTimeAlertConfiguration:
        disabled: false
        rules:
          - type: IssueDetection
            issueDetectionConfiguration:
              ruleName: MyIssueDetectionRule
          - type: KeywordMatch
            keywordMatchConfiguration:
              keywords:
                - keyword1
                - keyword2
              negate: false
              ruleName: MyKeywordMatchRule
          - type: Sentiment
            sentimentConfiguration:
              ruleName: MySentimentRule
              sentimentType: NEGATIVE
              timePeriod: 60
```
{{% /example %}}
{{% example %}}
### Transcribe processor usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration", {
    resourceAccessRoleArn: aws_iam_role.example.arn,
    elements: [
        {
            type: "AmazonTranscribeProcessor",
            amazonTranscribeProcessorConfiguration: {
                contentIdentificationType: "PII",
                enablePartialResultsStabilization: true,
                filterPartialResults: true,
                languageCode: "en-US",
                languageModelName: "MyLanguageModel",
                partialResultsStability: "high",
                piiEntityTypes: "ADDRESS,BANK_ACCOUNT_NUMBER",
                showSpeakerLabel: true,
                vocabularyFilterMethod: "mask",
                vocabularyFilterName: "MyVocabularyFilter",
                vocabularyName: "MyVocabulary",
            },
        },
        {
            type: "KinesisDataStreamSink",
            kinesisDataStreamSinkConfiguration: {
                insightsTarget: aws_kinesis_stream.example.arn,
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

my_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration",
    resource_access_role_arn=aws_iam_role["example"]["arn"],
    elements=[
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="AmazonTranscribeProcessor",
            amazon_transcribe_processor_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeProcessorConfigurationArgs(
                content_identification_type="PII",
                enable_partial_results_stabilization=True,
                filter_partial_results=True,
                language_code="en-US",
                language_model_name="MyLanguageModel",
                partial_results_stability="high",
                pii_entity_types="ADDRESS,BANK_ACCOUNT_NUMBER",
                show_speaker_label=True,
                vocabulary_filter_method="mask",
                vocabulary_filter_name="MyVocabularyFilter",
                vocabulary_name="MyVocabulary",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="KinesisDataStreamSink",
            kinesis_data_stream_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs(
                insights_target=aws_kinesis_stream["example"]["arn"],
            ),
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
    var myConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("myConfiguration", new()
    {
        ResourceAccessRoleArn = aws_iam_role.Example.Arn,
        Elements = new[]
        {
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "AmazonTranscribeProcessor",
                AmazonTranscribeProcessorConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeProcessorConfigurationArgs
                {
                    ContentIdentificationType = "PII",
                    EnablePartialResultsStabilization = true,
                    FilterPartialResults = true,
                    LanguageCode = "en-US",
                    LanguageModelName = "MyLanguageModel",
                    PartialResultsStability = "high",
                    PiiEntityTypes = "ADDRESS,BANK_ACCOUNT_NUMBER",
                    ShowSpeakerLabel = true,
                    VocabularyFilterMethod = "mask",
                    VocabularyFilterName = "MyVocabularyFilter",
                    VocabularyName = "MyVocabulary",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "KinesisDataStreamSink",
                KinesisDataStreamSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs
                {
                    InsightsTarget = aws_kinesis_stream.Example.Arn,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "myConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Elements: chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArray{
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("AmazonTranscribeProcessor"),
					AmazonTranscribeProcessorConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementAmazonTranscribeProcessorConfigurationArgs{
						ContentIdentificationType:         pulumi.String("PII"),
						EnablePartialResultsStabilization: pulumi.Bool(true),
						FilterPartialResults:              pulumi.Bool(true),
						LanguageCode:                      pulumi.String("en-US"),
						LanguageModelName:                 pulumi.String("MyLanguageModel"),
						PartialResultsStability:           pulumi.String("high"),
						PiiEntityTypes:                    pulumi.String("ADDRESS,BANK_ACCOUNT_NUMBER"),
						ShowSpeakerLabel:                  pulumi.Bool(true),
						VocabularyFilterMethod:            pulumi.String("mask"),
						VocabularyFilterName:              pulumi.String("MyVocabularyFilter"),
						VocabularyName:                    pulumi.String("MyVocabulary"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("KinesisDataStreamSink"),
					KinesisDataStreamSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs{
						InsightsTarget: pulumi.Any(aws_kinesis_stream.Example.Arn),
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
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeProcessorConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
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
        var myConfiguration = new MediaInsightsPipelineConfiguration("myConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(aws_iam_role.example().arn())
            .elements(            
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("AmazonTranscribeProcessor")
                    .amazonTranscribeProcessorConfiguration(MediaInsightsPipelineConfigurationElementAmazonTranscribeProcessorConfigurationArgs.builder()
                        .contentIdentificationType("PII")
                        .enablePartialResultsStabilization(true)
                        .filterPartialResults(true)
                        .languageCode("en-US")
                        .languageModelName("MyLanguageModel")
                        .partialResultsStability("high")
                        .piiEntityTypes("ADDRESS,BANK_ACCOUNT_NUMBER")
                        .showSpeakerLabel(true)
                        .vocabularyFilterMethod("mask")
                        .vocabularyFilterName("MyVocabularyFilter")
                        .vocabularyName("MyVocabulary")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("KinesisDataStreamSink")
                    .kinesisDataStreamSinkConfiguration(MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs.builder()
                        .insightsTarget(aws_kinesis_stream.example().arn())
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  myConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${aws_iam_role.example.arn}
      elements:
        - type: AmazonTranscribeProcessor
          amazonTranscribeProcessorConfiguration:
            contentIdentificationType: PII
            enablePartialResultsStabilization: true
            filterPartialResults: true
            languageCode: en-US
            languageModelName: MyLanguageModel
            partialResultsStability: high
            piiEntityTypes: ADDRESS,BANK_ACCOUNT_NUMBER
            showSpeakerLabel: true
            vocabularyFilterMethod: mask
            vocabularyFilterName: MyVocabularyFilter
            vocabularyName: MyVocabulary
        - type: KinesisDataStreamSink
          kinesisDataStreamSinkConfiguration:
            insightsTarget: ${aws_kinesis_stream.example.arn}
```
{{% /example %}}
{{% example %}}
### Voice analytics processor usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration", {
    resourceAccessRoleArn: aws_iam_role.example.arn,
    elements: [
        {
            type: "VoiceAnalyticsProcessor",
            voiceAnalyticsProcessorConfiguration: {
                speakerSearchStatus: "Enabled",
                voiceToneAnalysisStatus: "Enabled",
            },
        },
        {
            type: "LambdaFunctionSink",
            lambdaFunctionSinkConfiguration: {
                insightsTarget: "arn:aws:lambda:us-west-2:1111111111:function:MyFunction",
            },
        },
        {
            type: "SnsTopicSink",
            snsTopicSinkConfiguration: {
                insightsTarget: "arn:aws:sns:us-west-2:1111111111:topic/MyTopic",
            },
        },
        {
            type: "SqsQueueSink",
            sqsQueueSinkConfiguration: {
                insightsTarget: "arn:aws:sqs:us-west-2:1111111111:queue/MyQueue",
            },
        },
        {
            type: "KinesisDataStreamSink",
            kinesisDataStreamSinkConfiguration: {
                insightsTarget: aws_kinesis_stream.test.arn,
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

my_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration",
    resource_access_role_arn=aws_iam_role["example"]["arn"],
    elements=[
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="VoiceAnalyticsProcessor",
            voice_analytics_processor_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementVoiceAnalyticsProcessorConfigurationArgs(
                speaker_search_status="Enabled",
                voice_tone_analysis_status="Enabled",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="LambdaFunctionSink",
            lambda_function_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementLambdaFunctionSinkConfigurationArgs(
                insights_target="arn:aws:lambda:us-west-2:1111111111:function:MyFunction",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="SnsTopicSink",
            sns_topic_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementSnsTopicSinkConfigurationArgs(
                insights_target="arn:aws:sns:us-west-2:1111111111:topic/MyTopic",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="SqsQueueSink",
            sqs_queue_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementSqsQueueSinkConfigurationArgs(
                insights_target="arn:aws:sqs:us-west-2:1111111111:queue/MyQueue",
            ),
        ),
        aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
            type="KinesisDataStreamSink",
            kinesis_data_stream_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs(
                insights_target=aws_kinesis_stream["test"]["arn"],
            ),
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
    var myConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("myConfiguration", new()
    {
        ResourceAccessRoleArn = aws_iam_role.Example.Arn,
        Elements = new[]
        {
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "VoiceAnalyticsProcessor",
                VoiceAnalyticsProcessorConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementVoiceAnalyticsProcessorConfigurationArgs
                {
                    SpeakerSearchStatus = "Enabled",
                    VoiceToneAnalysisStatus = "Enabled",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "LambdaFunctionSink",
                LambdaFunctionSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementLambdaFunctionSinkConfigurationArgs
                {
                    InsightsTarget = "arn:aws:lambda:us-west-2:1111111111:function:MyFunction",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "SnsTopicSink",
                SnsTopicSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementSnsTopicSinkConfigurationArgs
                {
                    InsightsTarget = "arn:aws:sns:us-west-2:1111111111:topic/MyTopic",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "SqsQueueSink",
                SqsQueueSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementSqsQueueSinkConfigurationArgs
                {
                    InsightsTarget = "arn:aws:sqs:us-west-2:1111111111:queue/MyQueue",
                },
            },
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "KinesisDataStreamSink",
                KinesisDataStreamSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs
                {
                    InsightsTarget = aws_kinesis_stream.Test.Arn,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "myConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Elements: chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArray{
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("VoiceAnalyticsProcessor"),
					VoiceAnalyticsProcessorConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementVoiceAnalyticsProcessorConfigurationArgs{
						SpeakerSearchStatus:     pulumi.String("Enabled"),
						VoiceToneAnalysisStatus: pulumi.String("Enabled"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("LambdaFunctionSink"),
					LambdaFunctionSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementLambdaFunctionSinkConfigurationArgs{
						InsightsTarget: pulumi.String("arn:aws:lambda:us-west-2:1111111111:function:MyFunction"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("SnsTopicSink"),
					SnsTopicSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementSnsTopicSinkConfigurationArgs{
						InsightsTarget: pulumi.String("arn:aws:sns:us-west-2:1111111111:topic/MyTopic"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("SqsQueueSink"),
					SqsQueueSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementSqsQueueSinkConfigurationArgs{
						InsightsTarget: pulumi.String("arn:aws:sqs:us-west-2:1111111111:queue/MyQueue"),
					},
				},
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("KinesisDataStreamSink"),
					KinesisDataStreamSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs{
						InsightsTarget: pulumi.Any(aws_kinesis_stream.Test.Arn),
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
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementVoiceAnalyticsProcessorConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementLambdaFunctionSinkConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementSnsTopicSinkConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementSqsQueueSinkConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
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
        var myConfiguration = new MediaInsightsPipelineConfiguration("myConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(aws_iam_role.example().arn())
            .elements(            
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("VoiceAnalyticsProcessor")
                    .voiceAnalyticsProcessorConfiguration(MediaInsightsPipelineConfigurationElementVoiceAnalyticsProcessorConfigurationArgs.builder()
                        .speakerSearchStatus("Enabled")
                        .voiceToneAnalysisStatus("Enabled")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("LambdaFunctionSink")
                    .lambdaFunctionSinkConfiguration(MediaInsightsPipelineConfigurationElementLambdaFunctionSinkConfigurationArgs.builder()
                        .insightsTarget("arn:aws:lambda:us-west-2:1111111111:function:MyFunction")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("SnsTopicSink")
                    .snsTopicSinkConfiguration(MediaInsightsPipelineConfigurationElementSnsTopicSinkConfigurationArgs.builder()
                        .insightsTarget("arn:aws:sns:us-west-2:1111111111:topic/MyTopic")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("SqsQueueSink")
                    .sqsQueueSinkConfiguration(MediaInsightsPipelineConfigurationElementSqsQueueSinkConfigurationArgs.builder()
                        .insightsTarget("arn:aws:sqs:us-west-2:1111111111:queue/MyQueue")
                        .build())
                    .build(),
                MediaInsightsPipelineConfigurationElementArgs.builder()
                    .type("KinesisDataStreamSink")
                    .kinesisDataStreamSinkConfiguration(MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs.builder()
                        .insightsTarget(aws_kinesis_stream.test().arn())
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  myConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${aws_iam_role.example.arn}
      elements:
        - type: VoiceAnalyticsProcessor
          voiceAnalyticsProcessorConfiguration:
            speakerSearchStatus: Enabled
            voiceToneAnalysisStatus: Enabled
        - type: LambdaFunctionSink
          lambdaFunctionSinkConfiguration:
            insightsTarget: arn:aws:lambda:us-west-2:1111111111:function:MyFunction
        - type: SnsTopicSink
          snsTopicSinkConfiguration:
            insightsTarget: arn:aws:sns:us-west-2:1111111111:topic/MyTopic
        - type: SqsQueueSink
          sqsQueueSinkConfiguration:
            insightsTarget: arn:aws:sqs:us-west-2:1111111111:queue/MyQueue
        - type: KinesisDataStreamSink
          kinesisDataStreamSinkConfiguration:
            insightsTarget: ${aws_kinesis_stream.test.arn}
```
{{% /example %}}
{{% example %}}
### S3 Recording sink usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration", {
    resourceAccessRoleArn: aws_iam_role.example.arn,
    elements: [{
        type: "S3RecordingSink",
        s3RecordingSinkConfiguration: {
            destination: "arn:aws:s3:::MyBucket",
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

my_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("myConfiguration",
    resource_access_role_arn=aws_iam_role["example"]["arn"],
    elements=[aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs(
        type="S3RecordingSink",
        s3_recording_sink_configuration=aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementS3RecordingSinkConfigurationArgs(
            destination="arn:aws:s3:::MyBucket",
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("myConfiguration", new()
    {
        ResourceAccessRoleArn = aws_iam_role.Example.Arn,
        Elements = new[]
        {
            new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementArgs
            {
                Type = "S3RecordingSink",
                S3RecordingSinkConfiguration = new Aws.ChimeSDKMediaPipelines.Inputs.MediaInsightsPipelineConfigurationElementS3RecordingSinkConfigurationArgs
                {
                    Destination = "arn:aws:s3:::MyBucket",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "myConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Elements: chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArray{
				&chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementArgs{
					Type: pulumi.String("S3RecordingSink"),
					S3RecordingSinkConfiguration: &chimesdkmediapipelines.MediaInsightsPipelineConfigurationElementS3RecordingSinkConfigurationArgs{
						Destination: pulumi.String("arn:aws:s3:::MyBucket"),
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
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementS3RecordingSinkConfigurationArgs;
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
        var myConfiguration = new MediaInsightsPipelineConfiguration("myConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(aws_iam_role.example().arn())
            .elements(MediaInsightsPipelineConfigurationElementArgs.builder()
                .type("S3RecordingSink")
                .s3RecordingSinkConfiguration(MediaInsightsPipelineConfigurationElementS3RecordingSinkConfigurationArgs.builder()
                    .destination("arn:aws:s3:::MyBucket")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  myConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${aws_iam_role.example.arn}
      elements:
        - type: S3RecordingSink
          s3RecordingSinkConfiguration:
            destination: arn:aws:s3:::MyBucket
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Chime SDK Media Pipelines Media Insights Pipeline Configuration using the `id`. For example:

```sh
 $ pulumi import aws:chimesdkmediapipelines/mediaInsightsPipelineConfiguration:MediaInsightsPipelineConfiguration example abcdef123456
```
 