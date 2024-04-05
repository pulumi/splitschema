Adds a streaming configuration for the specified Amazon Chime Voice Connector. The streaming configuration specifies whether media streaming is enabled for sending to Amazon Kinesis.
It also sets the retention period, in hours, for the Amazon Kinesis data.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultVoiceConnector = new aws.chime.VoiceConnector("defaultVoiceConnector", {requireEncryption: true});
const defaultVoiceConnectorStreaming = new aws.chime.VoiceConnectorStreaming("defaultVoiceConnectorStreaming", {
    disabled: false,
    voiceConnectorId: defaultVoiceConnector.id,
    dataRetention: 7,
    streamingNotificationTargets: ["SQS"],
});
```
```python
import pulumi
import pulumi_aws as aws

default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
default_voice_connector_streaming = aws.chime.VoiceConnectorStreaming("defaultVoiceConnectorStreaming",
    disabled=False,
    voice_connector_id=default_voice_connector.id,
    data_retention=7,
    streaming_notification_targets=["SQS"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultVoiceConnector = new Aws.Chime.VoiceConnector("defaultVoiceConnector", new()
    {
        RequireEncryption = true,
    });

    var defaultVoiceConnectorStreaming = new Aws.Chime.VoiceConnectorStreaming("defaultVoiceConnectorStreaming", new()
    {
        Disabled = false,
        VoiceConnectorId = defaultVoiceConnector.Id,
        DataRetention = 7,
        StreamingNotificationTargets = new[]
        {
            "SQS",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chime"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultVoiceConnector, err := chime.NewVoiceConnector(ctx, "defaultVoiceConnector", &chime.VoiceConnectorArgs{
			RequireEncryption: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = chime.NewVoiceConnectorStreaming(ctx, "defaultVoiceConnectorStreaming", &chime.VoiceConnectorStreamingArgs{
			Disabled:         pulumi.Bool(false),
			VoiceConnectorId: defaultVoiceConnector.ID(),
			DataRetention:    pulumi.Int(7),
			StreamingNotificationTargets: pulumi.StringArray{
				pulumi.String("SQS"),
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
import com.pulumi.aws.chime.VoiceConnector;
import com.pulumi.aws.chime.VoiceConnectorArgs;
import com.pulumi.aws.chime.VoiceConnectorStreaming;
import com.pulumi.aws.chime.VoiceConnectorStreamingArgs;
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
        var defaultVoiceConnector = new VoiceConnector("defaultVoiceConnector", VoiceConnectorArgs.builder()        
            .requireEncryption(true)
            .build());

        var defaultVoiceConnectorStreaming = new VoiceConnectorStreaming("defaultVoiceConnectorStreaming", VoiceConnectorStreamingArgs.builder()        
            .disabled(false)
            .voiceConnectorId(defaultVoiceConnector.id())
            .dataRetention(7)
            .streamingNotificationTargets("SQS")
            .build());

    }
}
```
```yaml
resources:
  defaultVoiceConnector:
    type: aws:chime:VoiceConnector
    properties:
      requireEncryption: true
  defaultVoiceConnectorStreaming:
    type: aws:chime:VoiceConnectorStreaming
    properties:
      disabled: false
      voiceConnectorId: ${defaultVoiceConnector.id}
      dataRetention: 7
      streamingNotificationTargets:
        - SQS
```
{{% /example %}}
{{% example %}}
### Example Usage With Media Insights

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultVoiceConnector = new aws.chime.VoiceConnector("defaultVoiceConnector", {requireEncryption: true});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["mediapipelines.chime.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleStream = new aws.kinesis.Stream("exampleStream", {shardCount: 2});
const exampleMediaInsightsPipelineConfiguration = new aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("exampleMediaInsightsPipelineConfiguration", {
    resourceAccessRoleArn: exampleRole.arn,
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
                insightsTarget: exampleStream.arn,
            },
        },
    ],
});
const defaultVoiceConnectorStreaming = new aws.chime.VoiceConnectorStreaming("defaultVoiceConnectorStreaming", {
    disabled: false,
    voiceConnectorId: defaultVoiceConnector.id,
    dataRetention: 7,
    streamingNotificationTargets: ["SQS"],
    mediaInsightsConfiguration: {
        disabled: false,
        configurationArn: exampleMediaInsightsPipelineConfiguration.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["mediapipelines.chime.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_stream = aws.kinesis.Stream("exampleStream", shard_count=2)
example_media_insights_pipeline_configuration = aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration("exampleMediaInsightsPipelineConfiguration",
    resource_access_role_arn=example_role.arn,
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
                insights_target=example_stream.arn,
            ),
        ),
    ])
default_voice_connector_streaming = aws.chime.VoiceConnectorStreaming("defaultVoiceConnectorStreaming",
    disabled=False,
    voice_connector_id=default_voice_connector.id,
    data_retention=7,
    streaming_notification_targets=["SQS"],
    media_insights_configuration=aws.chime.VoiceConnectorStreamingMediaInsightsConfigurationArgs(
        disabled=False,
        configuration_arn=example_media_insights_pipeline_configuration.arn,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultVoiceConnector = new Aws.Chime.VoiceConnector("defaultVoiceConnector", new()
    {
        RequireEncryption = true,
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

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleStream = new Aws.Kinesis.Stream("exampleStream", new()
    {
        ShardCount = 2,
    });

    var exampleMediaInsightsPipelineConfiguration = new Aws.ChimeSDKMediaPipelines.MediaInsightsPipelineConfiguration("exampleMediaInsightsPipelineConfiguration", new()
    {
        ResourceAccessRoleArn = exampleRole.Arn,
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
                    InsightsTarget = exampleStream.Arn,
                },
            },
        },
    });

    var defaultVoiceConnectorStreaming = new Aws.Chime.VoiceConnectorStreaming("defaultVoiceConnectorStreaming", new()
    {
        Disabled = false,
        VoiceConnectorId = defaultVoiceConnector.Id,
        DataRetention = 7,
        StreamingNotificationTargets = new[]
        {
            "SQS",
        },
        MediaInsightsConfiguration = new Aws.Chime.Inputs.VoiceConnectorStreamingMediaInsightsConfigurationArgs
        {
            Disabled = false,
            ConfigurationArn = exampleMediaInsightsPipelineConfiguration.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chime"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chimesdkmediapipelines"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultVoiceConnector, err := chime.NewVoiceConnector(ctx, "defaultVoiceConnector", &chime.VoiceConnectorArgs{
			RequireEncryption: pulumi.Bool(true),
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
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		exampleStream, err := kinesis.NewStream(ctx, "exampleStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(2),
		})
		if err != nil {
			return err
		}
		exampleMediaInsightsPipelineConfiguration, err := chimesdkmediapipelines.NewMediaInsightsPipelineConfiguration(ctx, "exampleMediaInsightsPipelineConfiguration", &chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs{
			ResourceAccessRoleArn: exampleRole.Arn,
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
						InsightsTarget: exampleStream.Arn,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = chime.NewVoiceConnectorStreaming(ctx, "defaultVoiceConnectorStreaming", &chime.VoiceConnectorStreamingArgs{
			Disabled:         pulumi.Bool(false),
			VoiceConnectorId: defaultVoiceConnector.ID(),
			DataRetention:    pulumi.Int(7),
			StreamingNotificationTargets: pulumi.StringArray{
				pulumi.String("SQS"),
			},
			MediaInsightsConfiguration: &chime.VoiceConnectorStreamingMediaInsightsConfigurationArgs{
				Disabled:         pulumi.Bool(false),
				ConfigurationArn: exampleMediaInsightsPipelineConfiguration.Arn,
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
import com.pulumi.aws.chime.VoiceConnector;
import com.pulumi.aws.chime.VoiceConnectorArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfiguration;
import com.pulumi.aws.chimesdkmediapipelines.MediaInsightsPipelineConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationArgs;
import com.pulumi.aws.chimesdkmediapipelines.inputs.MediaInsightsPipelineConfigurationElementKinesisDataStreamSinkConfigurationArgs;
import com.pulumi.aws.chime.VoiceConnectorStreaming;
import com.pulumi.aws.chime.VoiceConnectorStreamingArgs;
import com.pulumi.aws.chime.inputs.VoiceConnectorStreamingMediaInsightsConfigurationArgs;
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
        var defaultVoiceConnector = new VoiceConnector("defaultVoiceConnector", VoiceConnectorArgs.builder()        
            .requireEncryption(true)
            .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("mediapipelines.chime.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleStream = new Stream("exampleStream", StreamArgs.builder()        
            .shardCount(2)
            .build());

        var exampleMediaInsightsPipelineConfiguration = new MediaInsightsPipelineConfiguration("exampleMediaInsightsPipelineConfiguration", MediaInsightsPipelineConfigurationArgs.builder()        
            .resourceAccessRoleArn(exampleRole.arn())
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
                        .insightsTarget(exampleStream.arn())
                        .build())
                    .build())
            .build());

        var defaultVoiceConnectorStreaming = new VoiceConnectorStreaming("defaultVoiceConnectorStreaming", VoiceConnectorStreamingArgs.builder()        
            .disabled(false)
            .voiceConnectorId(defaultVoiceConnector.id())
            .dataRetention(7)
            .streamingNotificationTargets("SQS")
            .mediaInsightsConfiguration(VoiceConnectorStreamingMediaInsightsConfigurationArgs.builder()
                .disabled(false)
                .configurationArn(exampleMediaInsightsPipelineConfiguration.arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  defaultVoiceConnector:
    type: aws:chime:VoiceConnector
    properties:
      requireEncryption: true
  defaultVoiceConnectorStreaming:
    type: aws:chime:VoiceConnectorStreaming
    properties:
      disabled: false
      voiceConnectorId: ${defaultVoiceConnector.id}
      dataRetention: 7
      streamingNotificationTargets:
        - SQS
      mediaInsightsConfiguration:
        disabled: false
        configurationArn: ${exampleMediaInsightsPipelineConfiguration.arn}
  exampleMediaInsightsPipelineConfiguration:
    type: aws:chimesdkmediapipelines:MediaInsightsPipelineConfiguration
    properties:
      resourceAccessRoleArn: ${exampleRole.arn}
      elements:
        - type: AmazonTranscribeCallAnalyticsProcessor
          amazonTranscribeCallAnalyticsProcessorConfiguration:
            languageCode: en-US
        - type: KinesisDataStreamSink
          kinesisDataStreamSinkConfiguration:
            insightsTarget: ${exampleStream.arn}
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 2
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
                  - mediapipelines.chime.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Chime Voice Connector Streaming using the `voice_connector_id`. For example:

```sh
 $ pulumi import aws:chime/voiceConnectorStreaming:VoiceConnectorStreaming default abcdef1ghij2klmno3pqr4
```
 