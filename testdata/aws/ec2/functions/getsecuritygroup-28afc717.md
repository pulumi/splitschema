`aws.ec2.SecurityGroup` provides details about a specific Security Group.

This resource can prove useful when a module accepts a Security Group id as
an input variable and needs to, for example, determine the id of the
VPC that the security group belongs to.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how one might accept a Security Group id as a variable
and use this data source to obtain the data necessary to create a subnet.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const securityGroupId = config.requireObject("securityGroupId");
const selected = aws.ec2.getSecurityGroup({
    id: securityGroupId,
});
const subnet = new aws.ec2.Subnet("subnet", {
    vpcId: selected.then(selected => selected.vpcId),
    cidrBlock: "10.0.1.0/24",
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
security_group_id = config.require_object("securityGroupId")
selected = aws.ec2.get_security_group(id=security_group_id)
subnet = aws.ec2.Subnet("subnet",
    vpc_id=selected.vpc_id,
    cidr_block="10.0.1.0/24")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var securityGroupId = config.RequireObject<dynamic>("securityGroupId");
    var selected = Aws.Ec2.GetSecurityGroup.Invoke(new()
    {
        Id = securityGroupId,
    });

    var subnet = new Aws.Ec2.Subnet("subnet", new()
    {
        VpcId = selected.Apply(getSecurityGroupResult => getSecurityGroupResult.VpcId),
        CidrBlock = "10.0.1.0/24",
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
		securityGroupId := cfg.RequireObject("securityGroupId")
		selected, err := ec2.LookupSecurityGroup(ctx, &ec2.LookupSecurityGroupArgs{
			Id: pulumi.StringRef(securityGroupId),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSubnet(ctx, "subnet", &ec2.SubnetArgs{
			VpcId:     *pulumi.String(selected.VpcId),
			CidrBlock: pulumi.String("10.0.1.0/24"),
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
import com.pulumi.aws.ec2.inputs.GetSecurityGroupArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
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
        final var securityGroupId = config.get("securityGroupId");
        final var selected = Ec2Functions.getSecurityGroup(GetSecurityGroupArgs.builder()
            .id(securityGroupId)
            .build());

        var subnet = new Subnet("subnet", SubnetArgs.builder()        
            .vpcId(selected.applyValue(getSecurityGroupResult -> getSecurityGroupResult.vpcId()))
            .cidrBlock("10.0.1.0/24")
            .build());

    }
}
```
```yaml
configuration:
  securityGroupId:
    type: dynamic
resources:
  subnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${selected.vpcId}
      cidrBlock: 10.0.1.0/24
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getSecurityGroup
      Arguments:
        id: ${securityGroupId}
```
{{% /example %}}
{{% /examples %}}