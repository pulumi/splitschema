Provides a Kinesis Video Stream resource. Amazon Kinesis Video Streams makes it easy to securely stream video from connected devices to AWS for analytics, machine learning (ML), playback, and other processing.

For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.kinesis.VideoStream("default", {
    dataRetentionInHours: 1,
    deviceName: "kinesis-video-device-name",
    mediaType: "video/h264",
    tags: {
        Name: "kinesis-video-stream",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.kinesis.VideoStream("default",
    data_retention_in_hours=1,
    device_name="kinesis-video-device-name",
    media_type="video/h264",
    tags={
        "Name": "kinesis-video-stream",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Kinesis.VideoStream("default", new()
    {
        DataRetentionInHours = 1,
        DeviceName = "kinesis-video-device-name",
        MediaType = "video/h264",
        Tags = 
        {
            { "Name", "kinesis-video-stream" },
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
		_, err := kinesis.NewVideoStream(ctx, "default", &kinesis.VideoStreamArgs{
			DataRetentionInHours: pulumi.Int(1),
			DeviceName:           pulumi.String("kinesis-video-device-name"),
			MediaType:            pulumi.String("video/h264"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("kinesis-video-stream"),
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
import com.pulumi.aws.kinesis.VideoStream;
import com.pulumi.aws.kinesis.VideoStreamArgs;
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
        var default_ = new VideoStream("default", VideoStreamArgs.builder()        
            .dataRetentionInHours(1)
            .deviceName("kinesis-video-device-name")
            .mediaType("video/h264")
            .tags(Map.of("Name", "kinesis-video-stream"))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:kinesis:VideoStream
    properties:
      dataRetentionInHours: 1
      deviceName: kinesis-video-device-name
      mediaType: video/h264
      tags:
        Name: kinesis-video-stream
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Kinesis Streams using the `arn`. For example:

```sh
 $ pulumi import aws:kinesis/videoStream:VideoStream test_stream arn:aws:kinesisvideo:us-west-2:123456789012:stream/pulumi-kinesis-test/1554978910975
```
 