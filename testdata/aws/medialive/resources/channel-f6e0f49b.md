Resource for managing an AWS MediaLive Channel.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.medialive.Channel("example", {
    channelClass: "STANDARD",
    roleArn: aws_iam_role.example.arn,
    inputSpecification: {
        codec: "AVC",
        inputResolution: "HD",
        maximumBitrate: "MAX_20_MBPS",
    },
    inputAttachments: [{
        inputAttachmentName: "example-input",
        inputId: aws_medialive_input.example.id,
    }],
    destinations: [{
        id: "destination",
        settings: [
            {
                url: `s3://${aws_s3_bucket.main.id}/test1`,
            },
            {
                url: `s3://${aws_s3_bucket.main2.id}/test2`,
            },
        ],
    }],
    encoderSettings: {
        timecodeConfig: {
            source: "EMBEDDED",
        },
        audioDescriptions: [{
            audioSelectorName: "example audio selector",
            name: "audio-selector",
        }],
        videoDescriptions: [{
            name: "example-video",
        }],
        outputGroups: [{
            outputGroupSettings: {
                archiveGroupSettings: [{
                    destination: {
                        destinationRefId: "destination",
                    },
                }],
            },
            outputs: [{
                outputName: "example-name",
                videoDescriptionName: "example-video",
                audioDescriptionNames: ["audio-selector"],
                outputSettings: {
                    archiveOutputSettings: {
                        nameModifier: "_1",
                        extension: "m2ts",
                        containerSettings: {
                            m2tsSettings: {
                                audioBufferModel: "ATSC",
                                bufferModel: "MULTIPLEX",
                                rateMode: "CBR",
                            },
                        },
                    },
                },
            }],
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.medialive.Channel("example",
    channel_class="STANDARD",
    role_arn=aws_iam_role["example"]["arn"],
    input_specification=aws.medialive.ChannelInputSpecificationArgs(
        codec="AVC",
        input_resolution="HD",
        maximum_bitrate="MAX_20_MBPS",
    ),
    input_attachments=[aws.medialive.ChannelInputAttachmentArgs(
        input_attachment_name="example-input",
        input_id=aws_medialive_input["example"]["id"],
    )],
    destinations=[aws.medialive.ChannelDestinationArgs(
        id="destination",
        settings=[
            aws.medialive.ChannelDestinationSettingArgs(
                url=f"s3://{aws_s3_bucket['main']['id']}/test1",
            ),
            aws.medialive.ChannelDestinationSettingArgs(
                url=f"s3://{aws_s3_bucket['main2']['id']}/test2",
            ),
        ],
    )],
    encoder_settings=aws.medialive.ChannelEncoderSettingsArgs(
        timecode_config=aws.medialive.ChannelEncoderSettingsTimecodeConfigArgs(
            source="EMBEDDED",
        ),
        audio_descriptions=[aws.medialive.ChannelEncoderSettingsAudioDescriptionArgs(
            audio_selector_name="example audio selector",
            name="audio-selector",
        )],
        video_descriptions=[aws.medialive.ChannelEncoderSettingsVideoDescriptionArgs(
            name="example-video",
        )],
        output_groups=[aws.medialive.ChannelEncoderSettingsOutputGroupArgs(
            output_group_settings=aws.medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArgs(
                archive_group_settings=[aws.medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArgs(
                    destination=aws.medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestinationArgs(
                        destination_ref_id="destination",
                    ),
                )],
            ),
            outputs=[aws.medialive.ChannelEncoderSettingsOutputGroupOutputArgs(
                output_name="example-name",
                video_description_name="example-video",
                audio_description_names=["audio-selector"],
                output_settings=aws.medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs(
                    archive_output_settings=aws.medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsArgs(
                        name_modifier="_1",
                        extension="m2ts",
                        container_settings=aws.medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsArgs(
                            m2ts_settings=aws.medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsArgs(
                                audio_buffer_model="ATSC",
                                buffer_model="MULTIPLEX",
                                rate_mode="CBR",
                            ),
                        ),
                    ),
                ),
            )],
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
    var example = new Aws.MediaLive.Channel("example", new()
    {
        ChannelClass = "STANDARD",
        RoleArn = aws_iam_role.Example.Arn,
        InputSpecification = new Aws.MediaLive.Inputs.ChannelInputSpecificationArgs
        {
            Codec = "AVC",
            InputResolution = "HD",
            MaximumBitrate = "MAX_20_MBPS",
        },
        InputAttachments = new[]
        {
            new Aws.MediaLive.Inputs.ChannelInputAttachmentArgs
            {
                InputAttachmentName = "example-input",
                InputId = aws_medialive_input.Example.Id,
            },
        },
        Destinations = new[]
        {
            new Aws.MediaLive.Inputs.ChannelDestinationArgs
            {
                Id = "destination",
                Settings = new[]
                {
                    new Aws.MediaLive.Inputs.ChannelDestinationSettingArgs
                    {
                        Url = $"s3://{aws_s3_bucket.Main.Id}/test1",
                    },
                    new Aws.MediaLive.Inputs.ChannelDestinationSettingArgs
                    {
                        Url = $"s3://{aws_s3_bucket.Main2.Id}/test2",
                    },
                },
            },
        },
        EncoderSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsArgs
        {
            TimecodeConfig = new Aws.MediaLive.Inputs.ChannelEncoderSettingsTimecodeConfigArgs
            {
                Source = "EMBEDDED",
            },
            AudioDescriptions = new[]
            {
                new Aws.MediaLive.Inputs.ChannelEncoderSettingsAudioDescriptionArgs
                {
                    AudioSelectorName = "example audio selector",
                    Name = "audio-selector",
                },
            },
            VideoDescriptions = new[]
            {
                new Aws.MediaLive.Inputs.ChannelEncoderSettingsVideoDescriptionArgs
                {
                    Name = "example-video",
                },
            },
            OutputGroups = new[]
            {
                new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupArgs
                {
                    OutputGroupSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArgs
                    {
                        ArchiveGroupSettings = new[]
                        {
                            new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArgs
                            {
                                Destination = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestinationArgs
                                {
                                    DestinationRefId = "destination",
                                },
                            },
                        },
                    },
                    Outputs = new[]
                    {
                        new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputArgs
                        {
                            OutputName = "example-name",
                            VideoDescriptionName = "example-video",
                            AudioDescriptionNames = new[]
                            {
                                "audio-selector",
                            },
                            OutputSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs
                            {
                                ArchiveOutputSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsArgs
                                {
                                    NameModifier = "_1",
                                    Extension = "m2ts",
                                    ContainerSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsArgs
                                    {
                                        M2tsSettings = new Aws.MediaLive.Inputs.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsArgs
                                        {
                                            AudioBufferModel = "ATSC",
                                            BufferModel = "MULTIPLEX",
                                            RateMode = "CBR",
                                        },
                                    },
                                },
                            },
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
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/medialive"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := medialive.NewChannel(ctx, "example", &medialive.ChannelArgs{
			ChannelClass: pulumi.String("STANDARD"),
			RoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
			InputSpecification: &medialive.ChannelInputSpecificationArgs{
				Codec:           pulumi.String("AVC"),
				InputResolution: pulumi.String("HD"),
				MaximumBitrate:  pulumi.String("MAX_20_MBPS"),
			},
			InputAttachments: medialive.ChannelInputAttachmentArray{
				&medialive.ChannelInputAttachmentArgs{
					InputAttachmentName: pulumi.String("example-input"),
					InputId:             pulumi.Any(aws_medialive_input.Example.Id),
				},
			},
			Destinations: medialive.ChannelDestinationArray{
				&medialive.ChannelDestinationArgs{
					Id: pulumi.String("destination"),
					Settings: medialive.ChannelDestinationSettingArray{
						&medialive.ChannelDestinationSettingArgs{
							Url: pulumi.String(fmt.Sprintf("s3://%v/test1", aws_s3_bucket.Main.Id)),
						},
						&medialive.ChannelDestinationSettingArgs{
							Url: pulumi.String(fmt.Sprintf("s3://%v/test2", aws_s3_bucket.Main2.Id)),
						},
					},
				},
			},
			EncoderSettings: &medialive.ChannelEncoderSettingsArgs{
				TimecodeConfig: &medialive.ChannelEncoderSettingsTimecodeConfigArgs{
					Source: pulumi.String("EMBEDDED"),
				},
				AudioDescriptions: medialive.ChannelEncoderSettingsAudioDescriptionArray{
					&medialive.ChannelEncoderSettingsAudioDescriptionArgs{
						AudioSelectorName: pulumi.String("example audio selector"),
						Name:              pulumi.String("audio-selector"),
					},
				},
				VideoDescriptions: medialive.ChannelEncoderSettingsVideoDescriptionArray{
					&medialive.ChannelEncoderSettingsVideoDescriptionArgs{
						Name: pulumi.String("example-video"),
					},
				},
				OutputGroups: medialive.ChannelEncoderSettingsOutputGroupArray{
					&medialive.ChannelEncoderSettingsOutputGroupArgs{
						OutputGroupSettings: &medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArgs{
							ArchiveGroupSettings: medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArray{
								&medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArgs{
									Destination: &medialive.ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestinationArgs{
										DestinationRefId: pulumi.String("destination"),
									},
								},
							},
						},
						Outputs: medialive.ChannelEncoderSettingsOutputGroupOutputTypeArray{
							&medialive.ChannelEncoderSettingsOutputGroupOutputTypeArgs{
								OutputName:           pulumi.String("example-name"),
								VideoDescriptionName: pulumi.String("example-video"),
								AudioDescriptionNames: pulumi.StringArray{
									pulumi.String("audio-selector"),
								},
								OutputSettings: &medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs{
									ArchiveOutputSettings: &medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsArgs{
										NameModifier: pulumi.String("_1"),
										Extension:    pulumi.String("m2ts"),
										ContainerSettings: &medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsArgs{
											M2tsSettings: &medialive.ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsArgs{
												AudioBufferModel: pulumi.String("ATSC"),
												BufferModel:      pulumi.String("MULTIPLEX"),
												RateMode:         pulumi.String("CBR"),
											},
										},
									},
								},
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
import com.pulumi.aws.medialive.Channel;
import com.pulumi.aws.medialive.ChannelArgs;
import com.pulumi.aws.medialive.inputs.ChannelInputSpecificationArgs;
import com.pulumi.aws.medialive.inputs.ChannelInputAttachmentArgs;
import com.pulumi.aws.medialive.inputs.ChannelDestinationArgs;
import com.pulumi.aws.medialive.inputs.ChannelEncoderSettingsArgs;
import com.pulumi.aws.medialive.inputs.ChannelEncoderSettingsTimecodeConfigArgs;
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
        var example = new Channel("example", ChannelArgs.builder()        
            .channelClass("STANDARD")
            .roleArn(aws_iam_role.example().arn())
            .inputSpecification(ChannelInputSpecificationArgs.builder()
                .codec("AVC")
                .inputResolution("HD")
                .maximumBitrate("MAX_20_MBPS")
                .build())
            .inputAttachments(ChannelInputAttachmentArgs.builder()
                .inputAttachmentName("example-input")
                .inputId(aws_medialive_input.example().id())
                .build())
            .destinations(ChannelDestinationArgs.builder()
                .id("destination")
                .settings(                
                    ChannelDestinationSettingArgs.builder()
                        .url(String.format("s3://%s/test1", aws_s3_bucket.main().id()))
                        .build(),
                    ChannelDestinationSettingArgs.builder()
                        .url(String.format("s3://%s/test2", aws_s3_bucket.main2().id()))
                        .build())
                .build())
            .encoderSettings(ChannelEncoderSettingsArgs.builder()
                .timecodeConfig(ChannelEncoderSettingsTimecodeConfigArgs.builder()
                    .source("EMBEDDED")
                    .build())
                .audioDescriptions(ChannelEncoderSettingsAudioDescriptionArgs.builder()
                    .audioSelectorName("example audio selector")
                    .name("audio-selector")
                    .build())
                .videoDescriptions(ChannelEncoderSettingsVideoDescriptionArgs.builder()
                    .name("example-video")
                    .build())
                .outputGroups(ChannelEncoderSettingsOutputGroupArgs.builder()
                    .outputGroupSettings(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArgs.builder()
                        .archiveGroupSettings(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArgs.builder()
                            .destination(ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestinationArgs.builder()
                                .destinationRefId("destination")
                                .build())
                            .build())
                        .build())
                    .outputs(ChannelEncoderSettingsOutputGroupOutputArgs.builder()
                        .outputName("example-name")
                        .videoDescriptionName("example-video")
                        .audioDescriptionNames("audio-selector")
                        .outputSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettingsArgs.builder()
                            .archiveOutputSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsArgs.builder()
                                .nameModifier("_1")
                                .extension("m2ts")
                                .containerSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsArgs.builder()
                                    .m2tsSettings(ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettingsM2tsSettingsArgs.builder()
                                        .audioBufferModel("ATSC")
                                        .bufferModel("MULTIPLEX")
                                        .rateMode("CBR")
                                        .build())
                                    .build())
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
  example:
    type: aws:medialive:Channel
    properties:
      channelClass: STANDARD
      roleArn: ${aws_iam_role.example.arn}
      inputSpecification:
        codec: AVC
        inputResolution: HD
        maximumBitrate: MAX_20_MBPS
      inputAttachments:
        - inputAttachmentName: example-input
          inputId: ${aws_medialive_input.example.id}
      destinations:
        - id: destination
          settings:
            - url: s3://${aws_s3_bucket.main.id}/test1
            - url: s3://${aws_s3_bucket.main2.id}/test2
      encoderSettings:
        timecodeConfig:
          source: EMBEDDED
        audioDescriptions:
          - audioSelectorName: example audio selector
            name: audio-selector
        videoDescriptions:
          - name: example-video
        outputGroups:
          - outputGroupSettings:
              archiveGroupSettings:
                - destination:
                    destinationRefId: destination
            outputs:
              - outputName: example-name
                videoDescriptionName: example-video
                audioDescriptionNames:
                  - audio-selector
                outputSettings:
                  archiveOutputSettings:
                    nameModifier: _1
                    extension: m2ts
                    containerSettings:
                      m2tsSettings:
                        audioBufferModel: ATSC
                        bufferModel: MULTIPLEX
                        rateMode: CBR
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MediaLive Channel using the `channel_id`. For example:

```sh
 $ pulumi import aws:medialive/channel:Channel example 1234567
```
 