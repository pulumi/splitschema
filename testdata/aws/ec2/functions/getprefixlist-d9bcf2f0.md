`aws.ec2.getPrefixList` provides details about a specific AWS prefix list (PL)
in the current region.

This can be used both to validate a prefix list given in a variable
and to obtain the CIDR blocks (IP address ranges) for the associated
AWS service. The latter may be useful e.g., for adding network ACL
rules.

The aws.ec2.ManagedPrefixList data source is normally more appropriate to use given it can return customer-managed prefix list info, as well as additional attributes.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const privateS3VpcEndpoint = new aws.ec2.VpcEndpoint("privateS3VpcEndpoint", {
    vpcId: aws_vpc.foo.id,
    serviceName: "com.amazonaws.us-west-2.s3",
});
const privateS3PrefixList = aws.ec2.getPrefixListOutput({
    prefixListId: privateS3VpcEndpoint.prefixListId,
});
const bar = new aws.ec2.NetworkAcl("bar", {vpcId: aws_vpc.foo.id});
const privateS3NetworkAclRule = new aws.ec2.NetworkAclRule("privateS3NetworkAclRule", {
    networkAclId: bar.id,
    ruleNumber: 200,
    egress: false,
    protocol: "tcp",
    ruleAction: "allow",
    cidrBlock: privateS3PrefixList.apply(privateS3PrefixList => privateS3PrefixList.cidrBlocks?.[0]),
    fromPort: 443,
    toPort: 443,
});
```
```python
import pulumi
import pulumi_aws as aws

private_s3_vpc_endpoint = aws.ec2.VpcEndpoint("privateS3VpcEndpoint",
    vpc_id=aws_vpc["foo"]["id"],
    service_name="com.amazonaws.us-west-2.s3")
private_s3_prefix_list = aws.ec2.get_prefix_list_output(prefix_list_id=private_s3_vpc_endpoint.prefix_list_id)
bar = aws.ec2.NetworkAcl("bar", vpc_id=aws_vpc["foo"]["id"])
private_s3_network_acl_rule = aws.ec2.NetworkAclRule("privateS3NetworkAclRule",
    network_acl_id=bar.id,
    rule_number=200,
    egress=False,
    protocol="tcp",
    rule_action="allow",
    cidr_block=private_s3_prefix_list.cidr_blocks[0],
    from_port=443,
    to_port=443)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var privateS3VpcEndpoint = new Aws.Ec2.VpcEndpoint("privateS3VpcEndpoint", new()
    {
        VpcId = aws_vpc.Foo.Id,
        ServiceName = "com.amazonaws.us-west-2.s3",
    });

    var privateS3PrefixList = Aws.Ec2.GetPrefixList.Invoke(new()
    {
        PrefixListId = privateS3VpcEndpoint.PrefixListId,
    });

    var bar = new Aws.Ec2.NetworkAcl("bar", new()
    {
        VpcId = aws_vpc.Foo.Id,
    });

    var privateS3NetworkAclRule = new Aws.Ec2.NetworkAclRule("privateS3NetworkAclRule", new()
    {
        NetworkAclId = bar.Id,
        RuleNumber = 200,
        Egress = false,
        Protocol = "tcp",
        RuleAction = "allow",
        CidrBlock = privateS3PrefixList.Apply(getPrefixListResult => getPrefixListResult.CidrBlocks[0]),
        FromPort = 443,
        ToPort = 443,
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
		privateS3VpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "privateS3VpcEndpoint", &ec2.VpcEndpointArgs{
			VpcId:       pulumi.Any(aws_vpc.Foo.Id),
			ServiceName: pulumi.String("com.amazonaws.us-west-2.s3"),
		})
		if err != nil {
			return err
		}
		privateS3PrefixList := ec2.GetPrefixListOutput(ctx, ec2.GetPrefixListOutputArgs{
			PrefixListId: privateS3VpcEndpoint.PrefixListId,
		}, nil)
		bar, err := ec2.NewNetworkAcl(ctx, "bar", &ec2.NetworkAclArgs{
			VpcId: pulumi.Any(aws_vpc.Foo.Id),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkAclRule(ctx, "privateS3NetworkAclRule", &ec2.NetworkAclRuleArgs{
			NetworkAclId: bar.ID(),
			RuleNumber:   pulumi.Int(200),
			Egress:       pulumi.Bool(false),
			Protocol:     pulumi.String("tcp"),
			RuleAction:   pulumi.String("allow"),
			CidrBlock: privateS3PrefixList.ApplyT(func(privateS3PrefixList ec2.GetPrefixListResult) (*string, error) {
				return &privateS3PrefixList.CidrBlocks[0], nil
			}).(pulumi.StringPtrOutput),
			FromPort: pulumi.Int(443),
			ToPort:   pulumi.Int(443),
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
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetPrefixListArgs;
import com.pulumi.aws.ec2.NetworkAcl;
import com.pulumi.aws.ec2.NetworkAclArgs;
import com.pulumi.aws.ec2.NetworkAclRule;
import com.pulumi.aws.ec2.NetworkAclRuleArgs;
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
        var privateS3VpcEndpoint = new VpcEndpoint("privateS3VpcEndpoint", VpcEndpointArgs.builder()        
            .vpcId(aws_vpc.foo().id())
            .serviceName("com.amazonaws.us-west-2.s3")
            .build());

        final var privateS3PrefixList = Ec2Functions.getPrefixList(GetPrefixListArgs.builder()
            .prefixListId(privateS3VpcEndpoint.prefixListId())
            .build());

        var bar = new NetworkAcl("bar", NetworkAclArgs.builder()        
            .vpcId(aws_vpc.foo().id())
            .build());

        var privateS3NetworkAclRule = new NetworkAclRule("privateS3NetworkAclRule", NetworkAclRuleArgs.builder()        
            .networkAclId(bar.id())
            .ruleNumber(200)
            .egress(false)
            .protocol("tcp")
            .ruleAction("allow")
            .cidrBlock(privateS3PrefixList.applyValue(getPrefixListResult -> getPrefixListResult).applyValue(privateS3PrefixList -> privateS3PrefixList.applyValue(getPrefixListResult -> getPrefixListResult.cidrBlocks()[0])))
            .fromPort(443)
            .toPort(443)
            .build());

    }
}
```
```yaml
resources:
  privateS3VpcEndpoint:
    type: aws:ec2:VpcEndpoint
    properties:
      vpcId: ${aws_vpc.foo.id}
      serviceName: com.amazonaws.us-west-2.s3
  bar:
    type: aws:ec2:NetworkAcl
    properties:
      vpcId: ${aws_vpc.foo.id}
  privateS3NetworkAclRule:
    type: aws:ec2:NetworkAclRule
    properties:
      networkAclId: ${bar.id}
      ruleNumber: 200
      egress: false
      protocol: tcp
      ruleAction: allow
      cidrBlock: ${privateS3PrefixList.cidrBlocks[0]}
      fromPort: 443
      toPort: 443
variables:
  privateS3PrefixList:
    fn::invoke:
      Function: aws:ec2:getPrefixList
      Arguments:
        prefixListId: ${privateS3VpcEndpoint.prefixListId}
```
{{% /example %}}
{{% example %}}
### Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ec2.getPrefixList({
    filters: [{
        name: "prefix-list-id",
        values: ["pl-68a54001"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.get_prefix_list(filters=[aws.ec2.GetPrefixListFilterArgs(
    name="prefix-list-id",
    values=["pl-68a54001"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ec2.GetPrefixList.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetPrefixListFilterInputArgs
            {
                Name = "prefix-list-id",
                Values = new[]
                {
                    "pl-68a54001",
                },
            },
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
		_, err := ec2.GetPrefixList(ctx, &ec2.GetPrefixListArgs{
			Filters: []ec2.GetPrefixListFilter{
				{
					Name: "prefix-list-id",
					Values: []string{
						"pl-68a54001",
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetPrefixListArgs;
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
        final var test = Ec2Functions.getPrefixList(GetPrefixListArgs.builder()
            .filters(GetPrefixListFilterArgs.builder()
                .name("prefix-list-id")
                .values("pl-68a54001")
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ec2:getPrefixList
      Arguments:
        filters:
          - name: prefix-list-id
            values:
              - pl-68a54001
```
{{% /example %}}
{{% /examples %}}