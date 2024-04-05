`aws.ec2.Subnet` provides details about a specific VPC subnet.

This resource can prove useful when a module accepts a subnet ID as an input variable and needs to, for example, determine the ID of the VPC that the subnet belongs to.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how one might accept a subnet ID as a variable and use this data source to obtain the data necessary to create a security group that allows connections from hosts in that subnet.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const subnetId = config.requireObject("subnetId");
const selected = aws.ec2.getSubnet({
    id: subnetId,
});
const subnet = new aws.ec2.SecurityGroup("subnet", {
    vpcId: selected.then(selected => selected.vpcId),
    ingress: [{
        cidrBlocks: [selected.then(selected => selected.cidrBlock)],
        fromPort: 80,
        toPort: 80,
        protocol: "tcp",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
subnet_id = config.require_object("subnetId")
selected = aws.ec2.get_subnet(id=subnet_id)
subnet = aws.ec2.SecurityGroup("subnet",
    vpc_id=selected.vpc_id,
    ingress=[aws.ec2.SecurityGroupIngressArgs(
        cidr_blocks=[selected.cidr_block],
        from_port=80,
        to_port=80,
        protocol="tcp",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var subnetId = config.RequireObject<dynamic>("subnetId");
    var selected = Aws.Ec2.GetSubnet.Invoke(new()
    {
        Id = subnetId,
    });

    var subnet = new Aws.Ec2.SecurityGroup("subnet", new()
    {
        VpcId = selected.Apply(getSubnetResult => getSubnetResult.VpcId),
        Ingress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupIngressArgs
            {
                CidrBlocks = new[]
                {
                    selected.Apply(getSubnetResult => getSubnetResult.CidrBlock),
                },
                FromPort = 80,
                ToPort = 80,
                Protocol = "tcp",
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
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		subnetId := cfg.RequireObject("subnetId")
		selected, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Id: pulumi.StringRef(subnetId),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSecurityGroup(ctx, "subnet", &ec2.SecurityGroupArgs{
			VpcId: *pulumi.String(selected.VpcId),
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					CidrBlocks: pulumi.StringArray{
						*pulumi.String(selected.CidrBlock),
					},
					FromPort: pulumi.Int(80),
					ToPort:   pulumi.Int(80),
					Protocol: pulumi.String("tcp"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetSubnetArgs;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupIngressArgs;
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
        final var config = ctx.config();
        final var subnetId = config.get("subnetId");
        final var selected = Ec2Functions.getSubnet(GetSubnetArgs.builder()
            .id(subnetId)
            .build());

        var subnet = new SecurityGroup("subnet", SecurityGroupArgs.builder()        
            .vpcId(selected.applyValue(getSubnetResult -> getSubnetResult.vpcId()))
            .ingress(SecurityGroupIngressArgs.builder()
                .cidrBlocks(selected.applyValue(getSubnetResult -> getSubnetResult.cidrBlock()))
                .fromPort(80)
                .toPort(80)
                .protocol("tcp")
                .build())
            .build());

    }
}
```
```yaml
configuration:
  subnetId:
    type: dynamic
resources:
  subnet:
    type: aws:ec2:SecurityGroup
    properties:
      vpcId: ${selected.vpcId}
      ingress:
        - cidrBlocks:
            - ${selected.cidrBlock}
          fromPort: 80
          toPort: 80
          protocol: tcp
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getSubnet
      Arguments:
        id: ${subnetId}
```
{{% /example %}}
{{% example %}}
### Filter Example

If you want to match against tag `Name`, use:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = aws.ec2.getSubnet({
    filters: [{
        name: "tag:Name",
        values: ["yakdriver"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

selected = aws.ec2.get_subnet(filters=[aws.ec2.GetSubnetFilterArgs(
    name="tag:Name",
    values=["yakdriver"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var selected = Aws.Ec2.GetSubnet.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetSubnetFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "yakdriver",
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
		_, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Filters: []ec2.GetSubnetFilter{
				{
					Name: "tag:Name",
					Values: []string{
						"yakdriver",
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
import com.pulumi.aws.ec2.inputs.GetSubnetArgs;
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
        final var selected = Ec2Functions.getSubnet(GetSubnetArgs.builder()
            .filters(GetSubnetFilterArgs.builder()
                .name("tag:Name")
                .values("yakdriver")
                .build())
            .build());

    }
}
```
```yaml
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getSubnet
      Arguments:
        filters:
          - name: tag:Name
            values:
              - yakdriver
```
{{% /example %}}
{{% /examples %}}