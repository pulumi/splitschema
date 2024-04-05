Provides an IP address pool resource for IPAM.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const exampleVpcIpam = new aws.ec2.VpcIpam("exampleVpcIpam", {operatingRegions: [{
    regionName: current.then(current => current.name),
}]});
const exampleVpcIpamPool = new aws.ec2.VpcIpamPool("exampleVpcIpamPool", {
    addressFamily: "ipv4",
    ipamScopeId: exampleVpcIpam.privateDefaultScopeId,
    locale: current.then(current => current.name),
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example_vpc_ipam = aws.ec2.VpcIpam("exampleVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
    region_name=current.name,
)])
example_vpc_ipam_pool = aws.ec2.VpcIpamPool("exampleVpcIpamPool",
    address_family="ipv4",
    ipam_scope_id=example_vpc_ipam.private_default_scope_id,
    locale=current.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var exampleVpcIpam = new Aws.Ec2.VpcIpam("exampleVpcIpam", new()
    {
        OperatingRegions = new[]
        {
            new Aws.Ec2.Inputs.VpcIpamOperatingRegionArgs
            {
                RegionName = current.Apply(getRegionResult => getRegionResult.Name),
            },
        },
    });

    var exampleVpcIpamPool = new Aws.Ec2.VpcIpamPool("exampleVpcIpamPool", new()
    {
        AddressFamily = "ipv4",
        IpamScopeId = exampleVpcIpam.PrivateDefaultScopeId,
        Locale = current.Apply(getRegionResult => getRegionResult.Name),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleVpcIpam, err := ec2.NewVpcIpam(ctx, "exampleVpcIpam", &ec2.VpcIpamArgs{
			OperatingRegions: ec2.VpcIpamOperatingRegionArray{
				&ec2.VpcIpamOperatingRegionArgs{
					RegionName: *pulumi.String(current.Name),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpamPool(ctx, "exampleVpcIpamPool", &ec2.VpcIpamPoolArgs{
			AddressFamily: pulumi.String("ipv4"),
			IpamScopeId:   exampleVpcIpam.PrivateDefaultScopeId,
			Locale:        *pulumi.String(current.Name),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.ec2.VpcIpam;
import com.pulumi.aws.ec2.VpcIpamArgs;
import com.pulumi.aws.ec2.inputs.VpcIpamOperatingRegionArgs;
import com.pulumi.aws.ec2.VpcIpamPool;
import com.pulumi.aws.ec2.VpcIpamPoolArgs;
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
        final var current = AwsFunctions.getRegion();

        var exampleVpcIpam = new VpcIpam("exampleVpcIpam", VpcIpamArgs.builder()        
            .operatingRegions(VpcIpamOperatingRegionArgs.builder()
                .regionName(current.applyValue(getRegionResult -> getRegionResult.name()))
                .build())
            .build());

        var exampleVpcIpamPool = new VpcIpamPool("exampleVpcIpamPool", VpcIpamPoolArgs.builder()        
            .addressFamily("ipv4")
            .ipamScopeId(exampleVpcIpam.privateDefaultScopeId())
            .locale(current.applyValue(getRegionResult -> getRegionResult.name()))
            .build());

    }
}
```
```yaml
resources:
  exampleVpcIpam:
    type: aws:ec2:VpcIpam
    properties:
      operatingRegions:
        - regionName: ${current.name}
  exampleVpcIpamPool:
    type: aws:ec2:VpcIpamPool
    properties:
      addressFamily: ipv4
      ipamScopeId: ${exampleVpcIpam.privateDefaultScopeId}
      locale: ${current.name}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```

Nested Pools:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const example = new aws.ec2.VpcIpam("example", {operatingRegions: [{
    regionName: current.then(current => current.name),
}]});
const parent = new aws.ec2.VpcIpamPool("parent", {
    addressFamily: "ipv4",
    ipamScopeId: example.privateDefaultScopeId,
});
const parentTest = new aws.ec2.VpcIpamPoolCidr("parentTest", {
    ipamPoolId: parent.id,
    cidr: "172.20.0.0/16",
});
const child = new aws.ec2.VpcIpamPool("child", {
    addressFamily: "ipv4",
    ipamScopeId: example.privateDefaultScopeId,
    locale: current.then(current => current.name),
    sourceIpamPoolId: parent.id,
});
const childTest = new aws.ec2.VpcIpamPoolCidr("childTest", {
    ipamPoolId: child.id,
    cidr: "172.20.0.0/24",
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example = aws.ec2.VpcIpam("example", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
    region_name=current.name,
)])
parent = aws.ec2.VpcIpamPool("parent",
    address_family="ipv4",
    ipam_scope_id=example.private_default_scope_id)
parent_test = aws.ec2.VpcIpamPoolCidr("parentTest",
    ipam_pool_id=parent.id,
    cidr="172.20.0.0/16")
child = aws.ec2.VpcIpamPool("child",
    address_family="ipv4",
    ipam_scope_id=example.private_default_scope_id,
    locale=current.name,
    source_ipam_pool_id=parent.id)
child_test = aws.ec2.VpcIpamPoolCidr("childTest",
    ipam_pool_id=child.id,
    cidr="172.20.0.0/24")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var example = new Aws.Ec2.VpcIpam("example", new()
    {
        OperatingRegions = new[]
        {
            new Aws.Ec2.Inputs.VpcIpamOperatingRegionArgs
            {
                RegionName = current.Apply(getRegionResult => getRegionResult.Name),
            },
        },
    });

    var parent = new Aws.Ec2.VpcIpamPool("parent", new()
    {
        AddressFamily = "ipv4",
        IpamScopeId = example.PrivateDefaultScopeId,
    });

    var parentTest = new Aws.Ec2.VpcIpamPoolCidr("parentTest", new()
    {
        IpamPoolId = parent.Id,
        Cidr = "172.20.0.0/16",
    });

    var child = new Aws.Ec2.VpcIpamPool("child", new()
    {
        AddressFamily = "ipv4",
        IpamScopeId = example.PrivateDefaultScopeId,
        Locale = current.Apply(getRegionResult => getRegionResult.Name),
        SourceIpamPoolId = parent.Id,
    });

    var childTest = new Aws.Ec2.VpcIpamPoolCidr("childTest", new()
    {
        IpamPoolId = child.Id,
        Cidr = "172.20.0.0/24",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		example, err := ec2.NewVpcIpam(ctx, "example", &ec2.VpcIpamArgs{
			OperatingRegions: ec2.VpcIpamOperatingRegionArray{
				&ec2.VpcIpamOperatingRegionArgs{
					RegionName: *pulumi.String(current.Name),
				},
			},
		})
		if err != nil {
			return err
		}
		parent, err := ec2.NewVpcIpamPool(ctx, "parent", &ec2.VpcIpamPoolArgs{
			AddressFamily: pulumi.String("ipv4"),
			IpamScopeId:   example.PrivateDefaultScopeId,
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpamPoolCidr(ctx, "parentTest", &ec2.VpcIpamPoolCidrArgs{
			IpamPoolId: parent.ID(),
			Cidr:       pulumi.String("172.20.0.0/16"),
		})
		if err != nil {
			return err
		}
		child, err := ec2.NewVpcIpamPool(ctx, "child", &ec2.VpcIpamPoolArgs{
			AddressFamily:    pulumi.String("ipv4"),
			IpamScopeId:      example.PrivateDefaultScopeId,
			Locale:           *pulumi.String(current.Name),
			SourceIpamPoolId: parent.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpamPoolCidr(ctx, "childTest", &ec2.VpcIpamPoolCidrArgs{
			IpamPoolId: child.ID(),
			Cidr:       pulumi.String("172.20.0.0/24"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.ec2.VpcIpam;
import com.pulumi.aws.ec2.VpcIpamArgs;
import com.pulumi.aws.ec2.inputs.VpcIpamOperatingRegionArgs;
import com.pulumi.aws.ec2.VpcIpamPool;
import com.pulumi.aws.ec2.VpcIpamPoolArgs;
import com.pulumi.aws.ec2.VpcIpamPoolCidr;
import com.pulumi.aws.ec2.VpcIpamPoolCidrArgs;
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
        final var current = AwsFunctions.getRegion();

        var example = new VpcIpam("example", VpcIpamArgs.builder()        
            .operatingRegions(VpcIpamOperatingRegionArgs.builder()
                .regionName(current.applyValue(getRegionResult -> getRegionResult.name()))
                .build())
            .build());

        var parent = new VpcIpamPool("parent", VpcIpamPoolArgs.builder()        
            .addressFamily("ipv4")
            .ipamScopeId(example.privateDefaultScopeId())
            .build());

        var parentTest = new VpcIpamPoolCidr("parentTest", VpcIpamPoolCidrArgs.builder()        
            .ipamPoolId(parent.id())
            .cidr("172.20.0.0/16")
            .build());

        var child = new VpcIpamPool("child", VpcIpamPoolArgs.builder()        
            .addressFamily("ipv4")
            .ipamScopeId(example.privateDefaultScopeId())
            .locale(current.applyValue(getRegionResult -> getRegionResult.name()))
            .sourceIpamPoolId(parent.id())
            .build());

        var childTest = new VpcIpamPoolCidr("childTest", VpcIpamPoolCidrArgs.builder()        
            .ipamPoolId(child.id())
            .cidr("172.20.0.0/24")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpcIpam
    properties:
      operatingRegions:
        - regionName: ${current.name}
  parent:
    type: aws:ec2:VpcIpamPool
    properties:
      addressFamily: ipv4
      ipamScopeId: ${example.privateDefaultScopeId}
  parentTest:
    type: aws:ec2:VpcIpamPoolCidr
    properties:
      ipamPoolId: ${parent.id}
      cidr: 172.20.0.0/16
  child:
    type: aws:ec2:VpcIpamPool
    properties:
      addressFamily: ipv4
      ipamScopeId: ${example.privateDefaultScopeId}
      locale: ${current.name}
      sourceIpamPoolId: ${parent.id}
  childTest:
    type: aws:ec2:VpcIpamPoolCidr
    properties:
      ipamPoolId: ${child.id}
      cidr: 172.20.0.0/24
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IPAMs using the IPAM pool `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcIpamPool:VpcIpamPool example ipam-pool-0958f95207d978e1e
```
 