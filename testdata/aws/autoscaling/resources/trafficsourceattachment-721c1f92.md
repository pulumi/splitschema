Attaches a traffic source to an Auto Scaling group.

> **NOTE on Auto Scaling Groups, Attachments and Traffic Source Attachments:** Pulumi provides standalone Attachment (for attaching Classic Load Balancers and Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target groups) and Traffic Source Attachment (for attaching Load Balancers and VPC Lattice target groups) resources and an Auto Scaling Group resource with `load_balancers`, `target_group_arns` and `traffic_source` attributes. Do not use the same traffic source in more than one of these resources. Doing so will cause a conflict of attachments. A `lifecycle` configuration block can be used to suppress differences if necessary.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.autoscaling.TrafficSourceAttachment("example", {
    autoscalingGroupName: aws_autoscaling_group.example.id,
    trafficSource: {
        identifier: aws_lb_target_group.example.arn,
        type: "elbv2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.autoscaling.TrafficSourceAttachment("example",
    autoscaling_group_name=aws_autoscaling_group["example"]["id"],
    traffic_source=aws.autoscaling.TrafficSourceAttachmentTrafficSourceArgs(
        identifier=aws_lb_target_group["example"]["arn"],
        type="elbv2",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AutoScaling.TrafficSourceAttachment("example", new()
    {
        AutoscalingGroupName = aws_autoscaling_group.Example.Id,
        TrafficSource = new Aws.AutoScaling.Inputs.TrafficSourceAttachmentTrafficSourceArgs
        {
            Identifier = aws_lb_target_group.Example.Arn,
            Type = "elbv2",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := autoscaling.NewTrafficSourceAttachment(ctx, "example", &autoscaling.TrafficSourceAttachmentArgs{
			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Example.Id),
			TrafficSource: &autoscaling.TrafficSourceAttachmentTrafficSourceArgs{
				Identifier: pulumi.Any(aws_lb_target_group.Example.Arn),
				Type:       pulumi.String("elbv2"),
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
import com.pulumi.aws.autoscaling.TrafficSourceAttachment;
import com.pulumi.aws.autoscaling.TrafficSourceAttachmentArgs;
import com.pulumi.aws.autoscaling.inputs.TrafficSourceAttachmentTrafficSourceArgs;
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
        var example = new TrafficSourceAttachment("example", TrafficSourceAttachmentArgs.builder()        
            .autoscalingGroupName(aws_autoscaling_group.example().id())
            .trafficSource(TrafficSourceAttachmentTrafficSourceArgs.builder()
                .identifier(aws_lb_target_group.example().arn())
                .type("elbv2")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:autoscaling:TrafficSourceAttachment
    properties:
      autoscalingGroupName: ${aws_autoscaling_group.example.id}
      trafficSource:
        identifier: ${aws_lb_target_group.example.arn}
        type: elbv2
```
{{% /example %}}
{{% /examples %}}