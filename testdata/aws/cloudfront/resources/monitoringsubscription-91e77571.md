Provides a CloudFront real-time log configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.MonitoringSubscription("example", {
    distributionId: aws_cloudfront_distribution.example.id,
    monitoringSubscription: {
        realtimeMetricsSubscriptionConfig: {
            realtimeMetricsSubscriptionStatus: "Enabled",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.MonitoringSubscription("example",
    distribution_id=aws_cloudfront_distribution["example"]["id"],
    monitoring_subscription=aws.cloudfront.MonitoringSubscriptionMonitoringSubscriptionArgs(
        realtime_metrics_subscription_config=aws.cloudfront.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs(
            realtime_metrics_subscription_status="Enabled",
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
    var example = new Aws.CloudFront.MonitoringSubscription("example", new()
    {
        DistributionId = aws_cloudfront_distribution.Example.Id,
        MonitoringSubscriptionDetails = new Aws.CloudFront.Inputs.MonitoringSubscriptionMonitoringSubscriptionArgs
        {
            RealtimeMetricsSubscriptionConfig = new Aws.CloudFront.Inputs.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs
            {
                RealtimeMetricsSubscriptionStatus = "Enabled",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewMonitoringSubscription(ctx, "example", &cloudfront.MonitoringSubscriptionArgs{
			DistributionId: pulumi.Any(aws_cloudfront_distribution.Example.Id),
			MonitoringSubscription: &cloudfront.MonitoringSubscriptionMonitoringSubscriptionArgs{
				RealtimeMetricsSubscriptionConfig: &cloudfront.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs{
					RealtimeMetricsSubscriptionStatus: pulumi.String("Enabled"),
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
import com.pulumi.aws.cloudfront.MonitoringSubscription;
import com.pulumi.aws.cloudfront.MonitoringSubscriptionArgs;
import com.pulumi.aws.cloudfront.inputs.MonitoringSubscriptionMonitoringSubscriptionArgs;
import com.pulumi.aws.cloudfront.inputs.MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs;
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
        var example = new MonitoringSubscription("example", MonitoringSubscriptionArgs.builder()        
            .distributionId(aws_cloudfront_distribution.example().id())
            .monitoringSubscription(MonitoringSubscriptionMonitoringSubscriptionArgs.builder()
                .realtimeMetricsSubscriptionConfig(MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfigArgs.builder()
                    .realtimeMetricsSubscriptionStatus("Enabled")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:MonitoringSubscription
    properties:
      distributionId: ${aws_cloudfront_distribution.example.id}
      monitoringSubscription:
        realtimeMetricsSubscriptionConfig:
          realtimeMetricsSubscriptionStatus: Enabled
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront monitoring subscription using the id. For example:

```sh
 $ pulumi import aws:cloudfront/monitoringSubscription:MonitoringSubscription example E3QYSUHO4VYRGB
```
 