Provides an Elastic File System (EFS) mount target.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.Vpc("foo", {cidrBlock: "10.0.0.0/16"});
const alphaSubnet = new aws.ec2.Subnet("alphaSubnet", {
    vpcId: foo.id,
    availabilityZone: "us-west-2a",
    cidrBlock: "10.0.1.0/24",
});
const alphaMountTarget = new aws.efs.MountTarget("alphaMountTarget", {
    fileSystemId: aws_efs_file_system.foo.id,
    subnetId: alphaSubnet.id,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.Vpc("foo", cidr_block="10.0.0.0/16")
alpha_subnet = aws.ec2.Subnet("alphaSubnet",
    vpc_id=foo.id,
    availability_zone="us-west-2a",
    cidr_block="10.0.1.0/24")
alpha_mount_target = aws.efs.MountTarget("alphaMountTarget",
    file_system_id=aws_efs_file_system["foo"]["id"],
    subnet_id=alpha_subnet.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.Vpc("foo", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var alphaSubnet = new Aws.Ec2.Subnet("alphaSubnet", new()
    {
        VpcId = foo.Id,
        AvailabilityZone = "us-west-2a",
        CidrBlock = "10.0.1.0/24",
    });

    var alphaMountTarget = new Aws.Efs.MountTarget("alphaMountTarget", new()
    {
        FileSystemId = aws_efs_file_system.Foo.Id,
        SubnetId = alphaSubnet.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ec2.NewVpc(ctx, "foo", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		alphaSubnet, err := ec2.NewSubnet(ctx, "alphaSubnet", &ec2.SubnetArgs{
			VpcId:            foo.ID(),
			AvailabilityZone: pulumi.String("us-west-2a"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
		})
		if err != nil {
			return err
		}
		_, err = efs.NewMountTarget(ctx, "alphaMountTarget", &efs.MountTargetArgs{
			FileSystemId: pulumi.Any(aws_efs_file_system.Foo.Id),
			SubnetId:     alphaSubnet.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.efs.MountTarget;
import com.pulumi.aws.efs.MountTargetArgs;
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
        var foo = new Vpc("foo", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var alphaSubnet = new Subnet("alphaSubnet", SubnetArgs.builder()        
            .vpcId(foo.id())
            .availabilityZone("us-west-2a")
            .cidrBlock("10.0.1.0/24")
            .build());

        var alphaMountTarget = new MountTarget("alphaMountTarget", MountTargetArgs.builder()        
            .fileSystemId(aws_efs_file_system.foo().id())
            .subnetId(alphaSubnet.id())
            .build());

    }
}
```
```yaml
resources:
  alphaMountTarget:
    type: aws:efs:MountTarget
    properties:
      fileSystemId: ${aws_efs_file_system.foo.id}
      subnetId: ${alphaSubnet.id}
  foo:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  alphaSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${foo.id}
      availabilityZone: us-west-2a
      cidrBlock: 10.0.1.0/24
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the EFS mount targets using the `id`. For example:

```sh
 $ pulumi import aws:efs/mountTarget:MountTarget alpha fsmt-52a643fb
```
 