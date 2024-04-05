Provides an Traffic mirror filter rule.  
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
const ruleout = new aws.ec2.TrafficMirrorFilterRule("ruleout", {
    description: "test rule",
    trafficMirrorFilterId: filter.id,
    destinationCidrBlock: "10.0.0.0/8",
    sourceCidrBlock: "10.0.0.0/8",
    ruleNumber: 1,
    ruleAction: "accept",
    trafficDirection: "egress",
});
const rulein = new aws.ec2.TrafficMirrorFilterRule("rulein", {
    description: "test rule",
    trafficMirrorFilterId: filter.id,
    destinationCidrBlock: "10.0.0.0/8",
    sourceCidrBlock: "10.0.0.0/8",
    ruleNumber: 1,
    ruleAction: "accept",
    trafficDirection: "ingress",
    protocol: 6,
    destinationPortRange: {
        fromPort: 22,
        toPort: 53,
    },
    sourcePortRange: {
        fromPort: 0,
        toPort: 10,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

filter = aws.ec2.TrafficMirrorFilter("filter",
    description="traffic mirror filter - example",
    network_services=["amazon-dns"])
ruleout = aws.ec2.TrafficMirrorFilterRule("ruleout",
    description="test rule",
    traffic_mirror_filter_id=filter.id,
    destination_cidr_block="10.0.0.0/8",
    source_cidr_block="10.0.0.0/8",
    rule_number=1,
    rule_action="accept",
    traffic_direction="egress")
rulein = aws.ec2.TrafficMirrorFilterRule("rulein",
    description="test rule",
    traffic_mirror_filter_id=filter.id,
    destination_cidr_block="10.0.0.0/8",
    source_cidr_block="10.0.0.0/8",
    rule_number=1,
    rule_action="accept",
    traffic_direction="ingress",
    protocol=6,
    destination_port_range=aws.ec2.TrafficMirrorFilterRuleDestinationPortRangeArgs(
        from_port=22,
        to_port=53,
    ),
    source_port_range=aws.ec2.TrafficMirrorFilterRuleSourcePortRangeArgs(
        from_port=0,
        to_port=10,
    ))
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

    var ruleout = new Aws.Ec2.TrafficMirrorFilterRule("ruleout", new()
    {
        Description = "test rule",
        TrafficMirrorFilterId = filter.Id,
        DestinationCidrBlock = "10.0.0.0/8",
        SourceCidrBlock = "10.0.0.0/8",
        RuleNumber = 1,
        RuleAction = "accept",
        TrafficDirection = "egress",
    });

    var rulein = new Aws.Ec2.TrafficMirrorFilterRule("rulein", new()
    {
        Description = "test rule",
        TrafficMirrorFilterId = filter.Id,
        DestinationCidrBlock = "10.0.0.0/8",
        SourceCidrBlock = "10.0.0.0/8",
        RuleNumber = 1,
        RuleAction = "accept",
        TrafficDirection = "ingress",
        Protocol = 6,
        DestinationPortRange = new Aws.Ec2.Inputs.TrafficMirrorFilterRuleDestinationPortRangeArgs
        {
            FromPort = 22,
            ToPort = 53,
        },
        SourcePortRange = new Aws.Ec2.Inputs.TrafficMirrorFilterRuleSourcePortRangeArgs
        {
            FromPort = 0,
            ToPort = 10,
        },
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
		_, err = ec2.NewTrafficMirrorFilterRule(ctx, "ruleout", &ec2.TrafficMirrorFilterRuleArgs{
			Description:           pulumi.String("test rule"),
			TrafficMirrorFilterId: filter.ID(),
			DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
			SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
			RuleNumber:            pulumi.Int(1),
			RuleAction:            pulumi.String("accept"),
			TrafficDirection:      pulumi.String("egress"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewTrafficMirrorFilterRule(ctx, "rulein", &ec2.TrafficMirrorFilterRuleArgs{
			Description:           pulumi.String("test rule"),
			TrafficMirrorFilterId: filter.ID(),
			DestinationCidrBlock:  pulumi.String("10.0.0.0/8"),
			SourceCidrBlock:       pulumi.String("10.0.0.0/8"),
			RuleNumber:            pulumi.Int(1),
			RuleAction:            pulumi.String("accept"),
			TrafficDirection:      pulumi.String("ingress"),
			Protocol:              pulumi.Int(6),
			DestinationPortRange: &ec2.TrafficMirrorFilterRuleDestinationPortRangeArgs{
				FromPort: pulumi.Int(22),
				ToPort:   pulumi.Int(53),
			},
			SourcePortRange: &ec2.TrafficMirrorFilterRuleSourcePortRangeArgs{
				FromPort: pulumi.Int(0),
				ToPort:   pulumi.Int(10),
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
import com.pulumi.aws.ec2.TrafficMirrorFilter;
import com.pulumi.aws.ec2.TrafficMirrorFilterArgs;
import com.pulumi.aws.ec2.TrafficMirrorFilterRule;
import com.pulumi.aws.ec2.TrafficMirrorFilterRuleArgs;
import com.pulumi.aws.ec2.inputs.TrafficMirrorFilterRuleDestinationPortRangeArgs;
import com.pulumi.aws.ec2.inputs.TrafficMirrorFilterRuleSourcePortRangeArgs;
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

        var ruleout = new TrafficMirrorFilterRule("ruleout", TrafficMirrorFilterRuleArgs.builder()        
            .description("test rule")
            .trafficMirrorFilterId(filter.id())
            .destinationCidrBlock("10.0.0.0/8")
            .sourceCidrBlock("10.0.0.0/8")
            .ruleNumber(1)
            .ruleAction("accept")
            .trafficDirection("egress")
            .build());

        var rulein = new TrafficMirrorFilterRule("rulein", TrafficMirrorFilterRuleArgs.builder()        
            .description("test rule")
            .trafficMirrorFilterId(filter.id())
            .destinationCidrBlock("10.0.0.0/8")
            .sourceCidrBlock("10.0.0.0/8")
            .ruleNumber(1)
            .ruleAction("accept")
            .trafficDirection("ingress")
            .protocol(6)
            .destinationPortRange(TrafficMirrorFilterRuleDestinationPortRangeArgs.builder()
                .fromPort(22)
                .toPort(53)
                .build())
            .sourcePortRange(TrafficMirrorFilterRuleSourcePortRangeArgs.builder()
                .fromPort(0)
                .toPort(10)
                .build())
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
  ruleout:
    type: aws:ec2:TrafficMirrorFilterRule
    properties:
      description: test rule
      trafficMirrorFilterId: ${filter.id}
      destinationCidrBlock: 10.0.0.0/8
      sourceCidrBlock: 10.0.0.0/8
      ruleNumber: 1
      ruleAction: accept
      trafficDirection: egress
  rulein:
    type: aws:ec2:TrafficMirrorFilterRule
    properties:
      description: test rule
      trafficMirrorFilterId: ${filter.id}
      destinationCidrBlock: 10.0.0.0/8
      sourceCidrBlock: 10.0.0.0/8
      ruleNumber: 1
      ruleAction: accept
      trafficDirection: ingress
      protocol: 6
      destinationPortRange:
        fromPort: 22
        toPort: 53
      sourcePortRange:
        fromPort: 0
        toPort: 10
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import traffic mirror rules using the `traffic_mirror_filter_id` and `id` separated by `:`. For example:

```sh
 $ pulumi import aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule rule tmf-0fbb93ddf38198f64:tmfr-05a458f06445d0aee
```
 