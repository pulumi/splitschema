Manages an outbound (egress) rule for a security group.

When specifying an outbound rule for your security group in a VPC, the configuration must include a destination for the traffic.

> **NOTE on Security Groups and Security Group Rules:** this provider currently provides a Security Group resource with `ingress` and `egress` rules defined in-line and a Security Group Rule resource which manages one or more `ingress` or
`egress` rules. Both of these resource were added before AWS assigned a [security group rule unique ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules.html), and they do not work well in all scenarios using the`description` and `tags` attributes, which rely on the unique ID.
The `aws.vpc.SecurityGroupEgressRule` resource has been added to address these limitations and should be used for all new security group rules.
You should not use the `aws.vpc.SecurityGroupEgressRule` resource in conjunction with an `aws.ec2.SecurityGroup` resource with in-line rules or with `aws.ec2.SecurityGroupRule` resources defined for the same Security Group, as rule conflicts may occur and rules will be overwritten.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpc.SecurityGroupEgressRule("example", {
    securityGroupId: aws_security_group.example.id,
    cidrIpv4: "10.0.0.0/8",
    fromPort: 80,
    ipProtocol: "tcp",
    toPort: 80,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpc.SecurityGroupEgressRule("example",
    security_group_id=aws_security_group["example"]["id"],
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
    var example = new Aws.Vpc.SecurityGroupEgressRule("example", new()
    {
        SecurityGroupId = aws_security_group.Example.Id,
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpc.NewSecurityGroupEgressRule(ctx, "example", &vpc.SecurityGroupEgressRuleArgs{
			SecurityGroupId: pulumi.Any(aws_security_group.Example.Id),
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
import com.pulumi.aws.vpc.SecurityGroupEgressRule;
import com.pulumi.aws.vpc.SecurityGroupEgressRuleArgs;
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
        var example = new SecurityGroupEgressRule("example", SecurityGroupEgressRuleArgs.builder()        
            .securityGroupId(aws_security_group.example().id())
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
  example:
    type: aws:vpc:SecurityGroupEgressRule
    properties:
      securityGroupId: ${aws_security_group.example.id}
      cidrIpv4: 10.0.0.0/8
      fromPort: 80
      ipProtocol: tcp
      toPort: 80
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import security group egress rules using the `security_group_rule_id`. For example:

```sh
 $ pulumi import aws:vpc/securityGroupEgressRule:SecurityGroupEgressRule example sgr-02108b27edd666983
```
 