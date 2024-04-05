Provides an Traffic mirror session.  
Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

{{% examples %}}
## Example Usage
{{% example %}}

To create a basic traffic mirror session

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const filter = new aws.ec2.TrafficMirrorFilter("filter", {
    description: "traffic mirror filter - example",
    networkServices: ["amazon-dns"],
});
const target = new aws.ec2.TrafficMirrorTarget("target", {networkLoadBalancerArn: aws_lb.lb.arn});
const session = new aws.ec2.TrafficMirrorSession("session", {
    description: "traffic mirror session - example",
    networkInterfaceId: aws_instance.test.primary_network_interface_id,
    sessionNumber: 1,
    trafficMirrorFilterId: filter.id,
    trafficMirrorTargetId: target.id,
});
```
```python
import pulumi
import pulumi_aws as aws

filter = aws.ec2.TrafficMirrorFilter("filter",
    description="traffic mirror filter - example",
    network_services=["amazon-dns"])
target = aws.ec2.TrafficMirrorTarget("target", network_load_balancer_arn=aws_lb["lb"]["arn"])
session = aws.ec2.TrafficMirrorSession("session",
    description="traffic mirror session - example",
    network_interface_id=aws_instance["test"]["primary_network_interface_id"],
    session_number=1,
    traffic_mirror_filter_id=filter.id,
    traffic_mirror_target_id=target.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var filter = new Aws.Ec2.TrafficMirrorFilter("filter", new()
    {
        Description = "traffic mirror filter - example",
        NetworkServices = new[]
        {
            "amazon-dns",
        },
    });

    var target = new Aws.Ec2.TrafficMirrorTarget("target", new()
    {
        NetworkLoadBalancerArn = aws_lb.Lb.Arn,
    });

    var session = new Aws.Ec2.TrafficMirrorSession("session", new()
    {
        Description = "traffic mirror session - example",
        NetworkInterfaceId = aws_instance.Test.Primary_network_interface_id,
        SessionNumber = 1,
        TrafficMirrorFilterId = filter.Id,
        TrafficMirrorTargetId = target.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		filter, err := ec2.NewTrafficMirrorFilter(ctx, "filter", &ec2.TrafficMirrorFilterArgs{
			Description: pulumi.String("traffic mirror filter - example"),
			NetworkServices: pulumi.StringArray{
				pulumi.String("amazon-dns"),
			},
		})
		if err != nil {
			return err
		}
		target, err := ec2.NewTrafficMirrorTarget(ctx, "target", &ec2.TrafficMirrorTargetArgs{
			NetworkLoadBalancerArn: pulumi.Any(aws_lb.Lb.Arn),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewTrafficMirrorSession(ctx, "session", &ec2.TrafficMirrorSessionArgs{
			Description:           pulumi.String("traffic mirror session - example"),
			NetworkInterfaceId:    pulumi.Any(aws_instance.Test.Primary_network_interface_id),
			SessionNumber:         pulumi.Int(1),
			TrafficMirrorFilterId: filter.ID(),
			TrafficMirrorTargetId: target.ID(),
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
import com.pulumi.aws.ec2.TrafficMirrorFilter;
import com.pulumi.aws.ec2.TrafficMirrorFilterArgs;
import com.pulumi.aws.ec2.TrafficMirrorTarget;
import com.pulumi.aws.ec2.TrafficMirrorTargetArgs;
import com.pulumi.aws.ec2.TrafficMirrorSession;
import com.pulumi.aws.ec2.TrafficMirrorSessionArgs;
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
        var filter = new TrafficMirrorFilter("filter", TrafficMirrorFilterArgs.builder()        
            .description("traffic mirror filter - example")
            .networkServices("amazon-dns")
            .build());

        var target = new TrafficMirrorTarget("target", TrafficMirrorTargetArgs.builder()        
            .networkLoadBalancerArn(aws_lb.lb().arn())
            .build());

        var session = new TrafficMirrorSession("session", TrafficMirrorSessionArgs.builder()        
            .description("traffic mirror session - example")
            .networkInterfaceId(aws_instance.test().primary_network_interface_id())
            .sessionNumber(1)
            .trafficMirrorFilterId(filter.id())
            .trafficMirrorTargetId(target.id())
            .build());

    }
}
```
```yaml
resources:
  filter:
    type: aws:ec2:TrafficMirrorFilter
    properties:
      description: traffic mirror filter - example
      networkServices:
        - amazon-dns
  target:
    type: aws:ec2:TrafficMirrorTarget
    properties:
      networkLoadBalancerArn: ${aws_lb.lb.arn}
  session:
    type: aws:ec2:TrafficMirrorSession
    properties:
      description: traffic mirror session - example
      networkInterfaceId: ${aws_instance.test.primary_network_interface_id}
      sessionNumber: 1
      trafficMirrorFilterId: ${filter.id}
      trafficMirrorTargetId: ${target.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import traffic mirror sessions using the `id`. For example:

```sh
 $ pulumi import aws:ec2/trafficMirrorSession:TrafficMirrorSession session tms-0d8aa3ca35897b82e
```
 