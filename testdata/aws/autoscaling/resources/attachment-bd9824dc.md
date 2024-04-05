Attaches a load balancer to an Auto Scaling group.

> **NOTE on Auto Scaling Groups, Attachments and Traffic Source Attachments:** Pulumi provides standalone Attachment (for attaching Classic Load Balancers and Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target groups) and Traffic Source Attachment (for attaching Load Balancers and VPC Lattice target groups) resources and an Auto Scaling Group resource with `load_balancers`, `target_group_arns` and `traffic_source` attributes. Do not use the same traffic source in more than one of these resources. Doing so will cause a conflict of attachments. A `lifecycle` configuration block can be used to suppress differences if necessary.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new load balancer attachment
const example = new aws.autoscaling.Attachment("example", {
    autoscalingGroupName: aws_autoscaling_group.example.id,
    elb: aws_elb.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new load balancer attachment
example = aws.autoscaling.Attachment("example",
    autoscaling_group_name=aws_autoscaling_group["example"]["id"],
    elb=aws_elb["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new load balancer attachment
    var example = new Aws.AutoScaling.Attachment("example", new()
    {
        AutoscalingGroupName = aws_autoscaling_group.Example.Id,
        Elb = aws_elb.Example.Id,
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
		_, err := autoscaling.NewAttachment(ctx, "example", &autoscaling.AttachmentArgs{
			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Example.Id),
			Elb:                  pulumi.Any(aws_elb.Example.Id),
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
import com.pulumi.aws.autoscaling.Attachment;
import com.pulumi.aws.autoscaling.AttachmentArgs;
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
        var example = new Attachment("example", AttachmentArgs.builder()        
            .autoscalingGroupName(aws_autoscaling_group.example().id())
            .elb(aws_elb.example().id())
            .build());

    }
}
```
```yaml
resources:
  # Create a new load balancer attachment
  example:
    type: aws:autoscaling:Attachment
    properties:
      autoscalingGroupName: ${aws_autoscaling_group.example.id}
      elb: ${aws_elb.example.id}
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new ALB Target Group attachment
const example = new aws.autoscaling.Attachment("example", {
    autoscalingGroupName: aws_autoscaling_group.example.id,
    lbTargetGroupArn: aws_lb_target_group.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new ALB Target Group attachment
example = aws.autoscaling.Attachment("example",
    autoscaling_group_name=aws_autoscaling_group["example"]["id"],
    lb_target_group_arn=aws_lb_target_group["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new ALB Target Group attachment
    var example = new Aws.AutoScaling.Attachment("example", new()
    {
        AutoscalingGroupName = aws_autoscaling_group.Example.Id,
        LbTargetGroupArn = aws_lb_target_group.Example.Arn,
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
		_, err := autoscaling.NewAttachment(ctx, "example", &autoscaling.AttachmentArgs{
			AutoscalingGroupName: pulumi.Any(aws_autoscaling_group.Example.Id),
			LbTargetGroupArn:     pulumi.Any(aws_lb_target_group.Example.Arn),
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
import com.pulumi.aws.autoscaling.Attachment;
import com.pulumi.aws.autoscaling.AttachmentArgs;
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
        var example = new Attachment("example", AttachmentArgs.builder()        
            .autoscalingGroupName(aws_autoscaling_group.example().id())
            .lbTargetGroupArn(aws_lb_target_group.example().arn())
            .build());

    }
}
```
```yaml
resources:
  # Create a new ALB Target Group attachment
  example:
    type: aws:autoscaling:Attachment
    properties:
      autoscalingGroupName: ${aws_autoscaling_group.example.id}
      lbTargetGroupArn: ${aws_lb_target_group.example.arn}
```
{{% /example %}}
{{% /examples %}}