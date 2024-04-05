Provides an Amazon Connect Routing Profile resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.RoutingProfile("example", {
    defaultOutboundQueueId: "12345678-1234-1234-1234-123456789012",
    description: "example description",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    mediaConcurrencies: [{
        channel: "VOICE",
        concurrency: 1,
    }],
    queueConfigs: [{
        channel: "VOICE",
        delay: 2,
        priority: 1,
        queueId: "12345678-1234-1234-1234-123456789012",
    }],
    tags: {
        Name: "Example Routing Profile",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.RoutingProfile("example",
    default_outbound_queue_id="12345678-1234-1234-1234-123456789012",
    description="example description",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    media_concurrencies=[aws.connect.RoutingProfileMediaConcurrencyArgs(
        channel="VOICE",
        concurrency=1,
    )],
    queue_configs=[aws.connect.RoutingProfileQueueConfigArgs(
        channel="VOICE",
        delay=2,
        priority=1,
        queue_id="12345678-1234-1234-1234-123456789012",
    )],
    tags={
        "Name": "Example Routing Profile",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.RoutingProfile("example", new()
    {
        DefaultOutboundQueueId = "12345678-1234-1234-1234-123456789012",
        Description = "example description",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        MediaConcurrencies = new[]
        {
            new Aws.Connect.Inputs.RoutingProfileMediaConcurrencyArgs
            {
                Channel = "VOICE",
                Concurrency = 1,
            },
        },
        QueueConfigs = new[]
        {
            new Aws.Connect.Inputs.RoutingProfileQueueConfigArgs
            {
                Channel = "VOICE",
                Delay = 2,
                Priority = 1,
                QueueId = "12345678-1234-1234-1234-123456789012",
            },
        },
        Tags = 
        {
            { "Name", "Example Routing Profile" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewRoutingProfile(ctx, "example", &connect.RoutingProfileArgs{
			DefaultOutboundQueueId: pulumi.String("12345678-1234-1234-1234-123456789012"),
			Description:            pulumi.String("example description"),
			InstanceId:             pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			MediaConcurrencies: connect.RoutingProfileMediaConcurrencyArray{
				&connect.RoutingProfileMediaConcurrencyArgs{
					Channel:     pulumi.String("VOICE"),
					Concurrency: pulumi.Int(1),
				},
			},
			QueueConfigs: connect.RoutingProfileQueueConfigArray{
				&connect.RoutingProfileQueueConfigArgs{
					Channel:  pulumi.String("VOICE"),
					Delay:    pulumi.Int(2),
					Priority: pulumi.Int(1),
					QueueId:  pulumi.String("12345678-1234-1234-1234-123456789012"),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Routing Profile"),
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
import com.pulumi.aws.connect.RoutingProfile;
import com.pulumi.aws.connect.RoutingProfileArgs;
import com.pulumi.aws.connect.inputs.RoutingProfileMediaConcurrencyArgs;
import com.pulumi.aws.connect.inputs.RoutingProfileQueueConfigArgs;
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
        var example = new RoutingProfile("example", RoutingProfileArgs.builder()        
            .defaultOutboundQueueId("12345678-1234-1234-1234-123456789012")
            .description("example description")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .mediaConcurrencies(RoutingProfileMediaConcurrencyArgs.builder()
                .channel("VOICE")
                .concurrency(1)
                .build())
            .queueConfigs(RoutingProfileQueueConfigArgs.builder()
                .channel("VOICE")
                .delay(2)
                .priority(1)
                .queueId("12345678-1234-1234-1234-123456789012")
                .build())
            .tags(Map.of("Name", "Example Routing Profile"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:RoutingProfile
    properties:
      defaultOutboundQueueId: 12345678-1234-1234-1234-123456789012
      description: example description
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      mediaConcurrencies:
        - channel: VOICE
          concurrency: 1
      queueConfigs:
        - channel: VOICE
          delay: 2
          priority: 1
          queueId: 12345678-1234-1234-1234-123456789012
      tags:
        Name: Example Routing Profile
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Routing Profiles using the `instance_id` and `routing_profile_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/routingProfile:RoutingProfile example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 