Provides an AWS Network Firewall Firewall Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.Firewall("example", {
    firewallPolicyArn: aws_networkfirewall_firewall_policy.example.arn,
    vpcId: aws_vpc.example.id,
    subnetMappings: [{
        subnetId: aws_subnet.example.id,
    }],
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.Firewall("example",
    firewall_policy_arn=aws_networkfirewall_firewall_policy["example"]["arn"],
    vpc_id=aws_vpc["example"]["id"],
    subnet_mappings=[aws.networkfirewall.FirewallSubnetMappingArgs(
        subnet_id=aws_subnet["example"]["id"],
    )],
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.Firewall("example", new()
    {
        FirewallPolicyArn = aws_networkfirewall_firewall_policy.Example.Arn,
        VpcId = aws_vpc.Example.Id,
        SubnetMappings = new[]
        {
            new Aws.NetworkFirewall.Inputs.FirewallSubnetMappingArgs
            {
                SubnetId = aws_subnet.Example.Id,
            },
        },
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.NewFirewall(ctx, "example", &networkfirewall.FirewallArgs{
			FirewallPolicyArn: pulumi.Any(aws_networkfirewall_firewall_policy.Example.Arn),
			VpcId:             pulumi.Any(aws_vpc.Example.Id),
			SubnetMappings: networkfirewall.FirewallSubnetMappingArray{
				&networkfirewall.FirewallSubnetMappingArgs{
					SubnetId: pulumi.Any(aws_subnet.Example.Id),
				},
			},
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
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
import com.pulumi.aws.networkfirewall.Firewall;
import com.pulumi.aws.networkfirewall.FirewallArgs;
import com.pulumi.aws.networkfirewall.inputs.FirewallSubnetMappingArgs;
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
        var example = new Firewall("example", FirewallArgs.builder()        
            .firewallPolicyArn(aws_networkfirewall_firewall_policy.example().arn())
            .vpcId(aws_vpc.example().id())
            .subnetMappings(FirewallSubnetMappingArgs.builder()
                .subnetId(aws_subnet.example().id())
                .build())
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:Firewall
    properties:
      firewallPolicyArn: ${aws_networkfirewall_firewall_policy.example.arn}
      vpcId: ${aws_vpc.example.id}
      subnetMappings:
        - subnetId: ${aws_subnet.example.id}
      tags:
        Tag1: Value1
        Tag2: Value2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Firewall Firewalls using their `arn`. For example:

```sh
 $ pulumi import aws:networkfirewall/firewall:Firewall example arn:aws:network-firewall:us-west-1:123456789012:firewall/example
```
 