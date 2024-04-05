Resource for managing an AWS MediaLive MultiplexProgram.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
});
const exampleMultiplex = new aws.medialive.Multiplex("exampleMultiplex", {
    availabilityZones: [
        available.then(available => available.names?.[0]),
        available.then(available => available.names?.[1]),
    ],
    multiplexSettings: {
        transportStreamBitrate: 1000000,
        transportStreamId: 1,
        transportStreamReservedBitrate: 1,
        maximumVideoBufferDelayMilliseconds: 1000,
    },
    startMultiplex: true,
    tags: {
        tag1: "value1",
    },
});
const exampleMultiplexProgram = new aws.medialive.MultiplexProgram("exampleMultiplexProgram", {
    programName: "example_program",
    multiplexId: exampleMultiplex.id,
    multiplexProgramSettings: {
        programNumber: 1,
        preferredChannelPipeline: "CURRENTLY_ACTIVE",
        videoSettings: {
            constantBitrate: 100000,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available")
example_multiplex = aws.medialive.Multiplex("exampleMultiplex",
    availability_zones=[
        available.names[0],
        available.names[1],
    ],
    multiplex_settings=aws.medialive.MultiplexMultiplexSettingsArgs(
        transport_stream_bitrate=1000000,
        transport_stream_id=1,
        transport_stream_reserved_bitrate=1,
        maximum_video_buffer_delay_milliseconds=1000,
    ),
    start_multiplex=True,
    tags={
        "tag1": "value1",
    })
example_multiplex_program = aws.medialive.MultiplexProgram("exampleMultiplexProgram",
    program_name="example_program",
    multiplex_id=example_multiplex.id,
    multiplex_program_settings=aws.medialive.MultiplexProgramMultiplexProgramSettingsArgs(
        program_number=1,
        preferred_channel_pipeline="CURRENTLY_ACTIVE",
        video_settings=aws.medialive.MultiplexProgramMultiplexProgramSettingsVideoSettingsArgs(
            constant_bitrate=100000,
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
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
    });

    var exampleMultiplex = new Aws.MediaLive.Multiplex("exampleMultiplex", new()
    {
        AvailabilityZones = new[]
        {
            available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
            available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[1]),
        },
        MultiplexSettings = new Aws.MediaLive.Inputs.MultiplexMultiplexSettingsArgs
        {
            TransportStreamBitrate = 1000000,
            TransportStreamId = 1,
            TransportStreamReservedBitrate = 1,
            MaximumVideoBufferDelayMilliseconds = 1000,
        },
        StartMultiplex = true,
        Tags = 
        {
            { "tag1", "value1" },
        },
    });

    var exampleMultiplexProgram = new Aws.MediaLive.MultiplexProgram("exampleMultiplexProgram", new()
    {
        ProgramName = "example_program",
        MultiplexId = exampleMultiplex.Id,
        MultiplexProgramSettings = new Aws.MediaLive.Inputs.MultiplexProgramMultiplexProgramSettingsArgs
        {
            ProgramNumber = 1,
            PreferredChannelPipeline = "CURRENTLY_ACTIVE",
            VideoSettings = new Aws.MediaLive.Inputs.MultiplexProgramMultiplexProgramSettingsVideoSettingsArgs
            {
                ConstantBitrate = 100000,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/medialive"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
		}, nil)
		if err != nil {
			return err
		}
		exampleMultiplex, err := medialive.NewMultiplex(ctx, "exampleMultiplex", &medialive.MultiplexArgs{
			AvailabilityZones: pulumi.StringArray{
				*pulumi.String(available.Names[0]),
				*pulumi.String(available.Names[1]),
			},
			MultiplexSettings: &medialive.MultiplexMultiplexSettingsArgs{
				TransportStreamBitrate:              pulumi.Int(1000000),
				TransportStreamId:                   pulumi.Int(1),
				TransportStreamReservedBitrate:      pulumi.Int(1),
				MaximumVideoBufferDelayMilliseconds: pulumi.Int(1000),
			},
			StartMultiplex: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		_, err = medialive.NewMultiplexProgram(ctx, "exampleMultiplexProgram", &medialive.MultiplexProgramArgs{
			ProgramName: pulumi.String("example_program"),
			MultiplexId: exampleMultiplex.ID(),
			MultiplexProgramSettings: &medialive.MultiplexProgramMultiplexProgramSettingsArgs{
				ProgramNumber:            pulumi.Int(1),
				PreferredChannelPipeline: pulumi.String("CURRENTLY_ACTIVE"),
				VideoSettings: &medialive.MultiplexProgramMultiplexProgramSettingsVideoSettingsArgs{
					ConstantBitrate: pulumi.Int(100000),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.medialive.Multiplex;
import com.pulumi.aws.medialive.MultiplexArgs;
import com.pulumi.aws.medialive.inputs.MultiplexMultiplexSettingsArgs;
import com.pulumi.aws.medialive.MultiplexProgram;
import com.pulumi.aws.medialive.MultiplexProgramArgs;
import com.pulumi.aws.medialive.inputs.MultiplexProgramMultiplexProgramSettingsArgs;
import com.pulumi.aws.medialive.inputs.MultiplexProgramMultiplexProgramSettingsVideoSettingsArgs;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .build());

        var exampleMultiplex = new Multiplex("exampleMultiplex", MultiplexArgs.builder()        
            .availabilityZones(            
                available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]),
                available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[1]))
            .multiplexSettings(MultiplexMultiplexSettingsArgs.builder()
                .transportStreamBitrate(1000000)
                .transportStreamId(1)
                .transportStreamReservedBitrate(1)
                .maximumVideoBufferDelayMilliseconds(1000)
                .build())
            .startMultiplex(true)
            .tags(Map.of("tag1", "value1"))
            .build());

        var exampleMultiplexProgram = new MultiplexProgram("exampleMultiplexProgram", MultiplexProgramArgs.builder()        
            .programName("example_program")
            .multiplexId(exampleMultiplex.id())
            .multiplexProgramSettings(MultiplexProgramMultiplexProgramSettingsArgs.builder()
                .programNumber(1)
                .preferredChannelPipeline("CURRENTLY_ACTIVE")
                .videoSettings(MultiplexProgramMultiplexProgramSettingsVideoSettingsArgs.builder()
                    .constantBitrate(100000)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleMultiplex:
    type: aws:medialive:Multiplex
    properties:
      availabilityZones:
        - ${available.names[0]}
        - ${available.names[1]}
      multiplexSettings:
        transportStreamBitrate: 1e+06
        transportStreamId: 1
        transportStreamReservedBitrate: 1
        maximumVideoBufferDelayMilliseconds: 1000
      startMultiplex: true
      tags:
        tag1: value1
  exampleMultiplexProgram:
    type: aws:medialive:MultiplexProgram
    properties:
      programName: example_program
      multiplexId: ${exampleMultiplex.id}
      multiplexProgramSettings:
        programNumber: 1
        preferredChannelPipeline: CURRENTLY_ACTIVE
        videoSettings:
          constantBitrate: 100000
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MediaLive MultiplexProgram using the `id`, or a combination of "`program_name`/`multiplex_id`". For example:

```sh
 $ pulumi import aws:medialive/multiplexProgram:MultiplexProgram example example_program/1234567
```
 