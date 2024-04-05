Provides a resource to create an EventBridge Global Endpoint.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _this = new aws.cloudwatch.EventEndpoint("this", {
    roleArn: aws_iam_role.replication.arn,
    eventBuses: [
        {
            eventBusArn: aws_cloudwatch_event_bus.primary.arn,
        },
        {
            eventBusArn: aws_cloudwatch_event_bus.secondary.arn,
        },
    ],
    replicationConfig: {
        state: "DISABLED",
    },
    routingConfig: {
        failoverConfig: {
            primary: {
                healthCheck: aws_route53_health_check.primary.arn,
            },
            secondary: {
                route: "us-east-2",
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

this = aws.cloudwatch.EventEndpoint("this",
    role_arn=aws_iam_role["replication"]["arn"],
    event_buses=[
        aws.cloudwatch.EventEndpointEventBusArgs(
            event_bus_arn=aws_cloudwatch_event_bus["primary"]["arn"],
        ),
        aws.cloudwatch.EventEndpointEventBusArgs(
            event_bus_arn=aws_cloudwatch_event_bus["secondary"]["arn"],
        ),
    ],
    replication_config=aws.cloudwatch.EventEndpointReplicationConfigArgs(
        state="DISABLED",
    ),
    routing_config=aws.cloudwatch.EventEndpointRoutingConfigArgs(
        failover_config=aws.cloudwatch.EventEndpointRoutingConfigFailoverConfigArgs(
            primary=aws.cloudwatch.EventEndpointRoutingConfigFailoverConfigPrimaryArgs(
                health_check=aws_route53_health_check["primary"]["arn"],
            ),
            secondary=aws.cloudwatch.EventEndpointRoutingConfigFailoverConfigSecondaryArgs(
                route="us-east-2",
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
    var @this = new Aws.CloudWatch.EventEndpoint("this", new()
    {
        RoleArn = aws_iam_role.Replication.Arn,
        EventBuses = new[]
        {
            new Aws.CloudWatch.Inputs.EventEndpointEventBusArgs
            {
                EventBusArn = aws_cloudwatch_event_bus.Primary.Arn,
            },
            new Aws.CloudWatch.Inputs.EventEndpointEventBusArgs
            {
                EventBusArn = aws_cloudwatch_event_bus.Secondary.Arn,
            },
        },
        ReplicationConfig = new Aws.CloudWatch.Inputs.EventEndpointReplicationConfigArgs
        {
            State = "DISABLED",
        },
        RoutingConfig = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigArgs
        {
            FailoverConfig = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigFailoverConfigArgs
            {
                Primary = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigFailoverConfigPrimaryArgs
                {
                    HealthCheck = aws_route53_health_check.Primary.Arn,
                },
                Secondary = new Aws.CloudWatch.Inputs.EventEndpointRoutingConfigFailoverConfigSecondaryArgs
                {
                    Route = "us-east-2",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventEndpoint(ctx, "this", &cloudwatch.EventEndpointArgs{
			RoleArn: pulumi.Any(aws_iam_role.Replication.Arn),
			EventBuses: cloudwatch.EventEndpointEventBusArray{
				&cloudwatch.EventEndpointEventBusArgs{
					EventBusArn: pulumi.Any(aws_cloudwatch_event_bus.Primary.Arn),
				},
				&cloudwatch.EventEndpointEventBusArgs{
					EventBusArn: pulumi.Any(aws_cloudwatch_event_bus.Secondary.Arn),
				},
			},
			ReplicationConfig: &cloudwatch.EventEndpointReplicationConfigArgs{
				State: pulumi.String("DISABLED"),
			},
			RoutingConfig: &cloudwatch.EventEndpointRoutingConfigArgs{
				FailoverConfig: &cloudwatch.EventEndpointRoutingConfigFailoverConfigArgs{
					Primary: &cloudwatch.EventEndpointRoutingConfigFailoverConfigPrimaryArgs{
						HealthCheck: pulumi.Any(aws_route53_health_check.Primary.Arn),
					},
					Secondary: &cloudwatch.EventEndpointRoutingConfigFailoverConfigSecondaryArgs{
						Route: pulumi.String("us-east-2"),
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
import com.pulumi.aws.cloudwatch.EventEndpoint;
import com.pulumi.aws.cloudwatch.EventEndpointArgs;
import com.pulumi.aws.cloudwatch.inputs.EventEndpointEventBusArgs;
import com.pulumi.aws.cloudwatch.inputs.EventEndpointReplicationConfigArgs;
import com.pulumi.aws.cloudwatch.inputs.EventEndpointRoutingConfigArgs;
import com.pulumi.aws.cloudwatch.inputs.EventEndpointRoutingConfigFailoverConfigArgs;
import com.pulumi.aws.cloudwatch.inputs.EventEndpointRoutingConfigFailoverConfigPrimaryArgs;
import com.pulumi.aws.cloudwatch.inputs.EventEndpointRoutingConfigFailoverConfigSecondaryArgs;
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
        var this_ = new EventEndpoint("this", EventEndpointArgs.builder()        
            .roleArn(aws_iam_role.replication().arn())
            .eventBuses(            
                EventEndpointEventBusArgs.builder()
                    .eventBusArn(aws_cloudwatch_event_bus.primary().arn())
                    .build(),
                EventEndpointEventBusArgs.builder()
                    .eventBusArn(aws_cloudwatch_event_bus.secondary().arn())
                    .build())
            .replicationConfig(EventEndpointReplicationConfigArgs.builder()
                .state("DISABLED")
                .build())
            .routingConfig(EventEndpointRoutingConfigArgs.builder()
                .failoverConfig(EventEndpointRoutingConfigFailoverConfigArgs.builder()
                    .primary(EventEndpointRoutingConfigFailoverConfigPrimaryArgs.builder()
                        .healthCheck(aws_route53_health_check.primary().arn())
                        .build())
                    .secondary(EventEndpointRoutingConfigFailoverConfigSecondaryArgs.builder()
                        .route("us-east-2")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  this:
    type: aws:cloudwatch:EventEndpoint
    properties:
      roleArn: ${aws_iam_role.replication.arn}
      eventBuses:
        - eventBusArn: ${aws_cloudwatch_event_bus.primary.arn}
        - eventBusArn: ${aws_cloudwatch_event_bus.secondary.arn}
      replicationConfig:
        state: DISABLED
      routingConfig:
        failoverConfig:
          primary:
            healthCheck: ${aws_route53_health_check.primary.arn}
          secondary:
            route: us-east-2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge Global Endpoints using the `name`. For example:

```sh
 $ pulumi import aws:cloudwatch/eventEndpoint:EventEndpoint imported_endpoint example-endpoint
```
 