Manages an inbound (ingress) rule for a security group.

When specifying an inbound rule for your security group in a VPC, the configuration must include a source for the traffic.

> **NOTE on Security Groups and Security Group Rules:** this provider currently provides a Security Group resource with `ingress` and `egress` rules defined in-line and a Security Group Rule resource which manages one or more `ingress` or
`egress` rules. Both of these resource were added before AWS assigned a [security group rule unique ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules.html), and they do not work well in all scenarios using the`description` and `tags` attributes, which rely on the unique ID.
The `aws.vpc.SecurityGroupIngressRule` resource has been added to address these limitations and should be used for all new security group rules.
You should not use the `aws.vpc.SecurityGroupIngressRule` resource in conjunction with an `aws.ec2.SecurityGroup` resource with in-line rules or with `aws.ec2.SecurityGroupRule` resources defined for the same Security Group, as rule conflicts may occur and rules will be overwritten.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSecurityGroup = new aws.ec2.SecurityGroup("exampleSecurityGroup", {
    description: "example",
    vpcId: aws_vpc.main.id,
    tags: {
        Name: "example",
    },
});
const exampleSecurityGroupIngressRule = new aws.vpc.SecurityGroupIngressRule("exampleSecurityGroupIngressRule", {
    securityGroupId: exampleSecurityGroup.id,
    cidrIpv4: "10.0.0.0/8",
    fromPort: 80,
    ipProtocol: "tcp",
    toPort: 80,
});
```
```python
import pulumi
import pulumi_aws as aws

example_security_group = aws.ec2.SecurityGroup("exampleSecurityGroup",
    description="example",
    vpc_id=aws_vpc["main"]["id"],
    tags={
        "Name": "example",
    })
example_security_group_ingress_rule = aws.vpc.SecurityGroupIngressRule("exampleSecurityGroupIngressRule",
    security_group_id=example_security_group.id,
    cidr_ipv4="10.0.0.0/8",
    from_port=80,
    ip_protocol="tcp",
    to_port=80)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSecurityGroup = new Aws.Ec2.SecurityGroup("exampleSecurityGroup", new()
    {
        Description = "example",
        VpcId = aws_vpc.Main.Id,
        Tags = 
        {
            { "Name", "example" },
        },
    });

    var exampleSecurityGroupIngressRule = new Aws.Vpc.SecurityGroupIngressRule("exampleSecurityGroupIngressRule", new()
    {
        SecurityGroupId = exampleSecurityGroup.Id,
        CidrIpv4 = "10.0.0.0/8",
        FromPort = 80,
        IpProtocol = "tcp",
        ToPort = 80,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSecurityGroup, err := ec2.NewSecurityGroup(ctx, "exampleSecurityGroup", &ec2.SecurityGroupArgs{
			Description: pulumi.String("example"),
			VpcId:       pulumi.Any(aws_vpc.Main.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
			},
		})
		if err != nil {
			return err
		}
		_, err = vpc.NewSecurityGroupIngressRule(ctx, "exampleSecurityGroupIngressRule", &vpc.SecurityGroupIngressRuleArgs{
			SecurityGroupId: exampleSecurityGroup.ID(),
			CidrIpv4:        pulumi.String("10.0.0.0/8"),
			FromPort:        pulumi.Int(80),
			IpProtocol:      pulumi.String("tcp"),
			ToPort:          pulumi.Int(80),
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
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.vpc.SecurityGroupIngressRule;
import com.pulumi.aws.vpc.SecurityGroupIngressRuleArgs;
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
        var exampleSecurityGroup = new SecurityGroup("exampleSecurityGroup", SecurityGroupArgs.builder()        
            .description("example")
            .vpcId(aws_vpc.main().id())
            .tags(Map.of("Name", "example"))
            .build());

        var exampleSecurityGroupIngressRule = new SecurityGroupIngressRule("exampleSecurityGroupIngressRule", SecurityGroupIngressRuleArgs.builder()        
            .securityGroupId(exampleSecurityGroup.id())
            .cidrIpv4("10.0.0.0/8")
            .fromPort(80)
            .ipProtocol("tcp")
            .toPort(80)
            .build());

    }
}
```
```yaml
resources:
  exampleSecurityGroup:
    type: aws:ec2:SecurityGroup
    properties:
      description: example
      vpcId: ${aws_vpc.main.id}
      tags:
        Name: example
  exampleSecurityGroupIngressRule:
    type: aws:vpc:SecurityGroupIngressRule
    properties:
      securityGroupId: ${exampleSecurityGroup.id}
      cidrIpv4: 10.0.0.0/8
      fromPort: 80
      ipProtocol: tcp
      toPort: 80
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import security group ingress rules using the `security_group_rule_id`. For example:

```sh
 $ pulumi import aws:vpc/securityGroupIngressRule:SecurityGroupIngressRule example sgr-02108b27edd666983
```
 