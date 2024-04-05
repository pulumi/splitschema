Allocates (reserves) a CIDR from an IPAM address pool, preventing usage by IPAM. Only works for private IPv4.

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
const exampleVpcIpamPoolCidr = new aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr", {
    ipamPoolId: exampleVpcIpamPool.id,
    cidr: "172.20.0.0/16",
});
const exampleVpcIpamPoolCidrAllocation = new aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation", {
    ipamPoolId: exampleVpcIpamPool.id,
    cidr: "172.20.0.0/24",
}, {
    dependsOn: [exampleVpcIpamPoolCidr],
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
example_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr",
    ipam_pool_id=example_vpc_ipam_pool.id,
    cidr="172.20.0.0/16")
example_vpc_ipam_pool_cidr_allocation = aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation",
    ipam_pool_id=example_vpc_ipam_pool.id,
    cidr="172.20.0.0/24",
    opts=pulumi.ResourceOptions(depends_on=[example_vpc_ipam_pool_cidr]))
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

    var exampleVpcIpamPoolCidr = new Aws.Ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr", new()
    {
        IpamPoolId = exampleVpcIpamPool.Id,
        Cidr = "172.20.0.0/16",
    });

    var exampleVpcIpamPoolCidrAllocation = new Aws.Ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation", new()
    {
        IpamPoolId = exampleVpcIpamPool.Id,
        Cidr = "172.20.0.0/24",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleVpcIpamPoolCidr,
        },
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
		exampleVpcIpamPool, err := ec2.NewVpcIpamPool(ctx, "exampleVpcIpamPool", &ec2.VpcIpamPoolArgs{
			AddressFamily: pulumi.String("ipv4"),
			IpamScopeId:   exampleVpcIpam.PrivateDefaultScopeId,
			Locale:        *pulumi.String(current.Name),
		})
		if err != nil {
			return err
		}
		exampleVpcIpamPoolCidr, err := ec2.NewVpcIpamPoolCidr(ctx, "exampleVpcIpamPoolCidr", &ec2.VpcIpamPoolCidrArgs{
			IpamPoolId: exampleVpcIpamPool.ID(),
			Cidr:       pulumi.String("172.20.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpamPoolCidrAllocation(ctx, "exampleVpcIpamPoolCidrAllocation", &ec2.VpcIpamPoolCidrAllocationArgs{
			IpamPoolId: exampleVpcIpamPool.ID(),
			Cidr:       pulumi.String("172.20.0.0/24"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleVpcIpamPoolCidr,
		}))
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
import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocation;
import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocationArgs;
import com.pulumi.resources.CustomResourceOptions;
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

        var exampleVpcIpamPoolCidr = new VpcIpamPoolCidr("exampleVpcIpamPoolCidr", VpcIpamPoolCidrArgs.builder()        
            .ipamPoolId(exampleVpcIpamPool.id())
            .cidr("172.20.0.0/16")
            .build());

        var exampleVpcIpamPoolCidrAllocation = new VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation", VpcIpamPoolCidrAllocationArgs.builder()        
            .ipamPoolId(exampleVpcIpamPool.id())
            .cidr("172.20.0.0/24")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleVpcIpamPoolCidr)
                .build());

    }
}
```
```yaml
resources:
  exampleVpcIpamPoolCidrAllocation:
    type: aws:ec2:VpcIpamPoolCidrAllocation
    properties:
      ipamPoolId: ${exampleVpcIpamPool.id}
      cidr: 172.20.0.0/24
    options:
      dependson:
        - ${exampleVpcIpamPoolCidr}
  exampleVpcIpamPoolCidr:
    type: aws:ec2:VpcIpamPoolCidr
    properties:
      ipamPoolId: ${exampleVpcIpamPool.id}
      cidr: 172.20.0.0/16
  exampleVpcIpamPool:
    type: aws:ec2:VpcIpamPool
    properties:
      addressFamily: ipv4
      ipamScopeId: ${exampleVpcIpam.privateDefaultScopeId}
      locale: ${current.name}
  exampleVpcIpam:
    type: aws:ec2:VpcIpam
    properties:
      operatingRegions:
        - regionName: ${current.name}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```

With the `disallowed_cidrs` attribute:

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
const exampleVpcIpamPoolCidr = new aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr", {
    ipamPoolId: exampleVpcIpamPool.id,
    cidr: "172.20.0.0/16",
});
const exampleVpcIpamPoolCidrAllocation = new aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation", {
    ipamPoolId: exampleVpcIpamPool.id,
    netmaskLength: 28,
    disallowedCidrs: ["172.20.0.0/28"],
}, {
    dependsOn: [exampleVpcIpamPoolCidr],
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
example_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr",
    ipam_pool_id=example_vpc_ipam_pool.id,
    cidr="172.20.0.0/16")
example_vpc_ipam_pool_cidr_allocation = aws.ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation",
    ipam_pool_id=example_vpc_ipam_pool.id,
    netmask_length=28,
    disallowed_cidrs=["172.20.0.0/28"],
    opts=pulumi.ResourceOptions(depends_on=[example_vpc_ipam_pool_cidr]))
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

    var exampleVpcIpamPoolCidr = new Aws.Ec2.VpcIpamPoolCidr("exampleVpcIpamPoolCidr", new()
    {
        IpamPoolId = exampleVpcIpamPool.Id,
        Cidr = "172.20.0.0/16",
    });

    var exampleVpcIpamPoolCidrAllocation = new Aws.Ec2.VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation", new()
    {
        IpamPoolId = exampleVpcIpamPool.Id,
        NetmaskLength = 28,
        DisallowedCidrs = new[]
        {
            "172.20.0.0/28",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleVpcIpamPoolCidr,
        },
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
		exampleVpcIpamPool, err := ec2.NewVpcIpamPool(ctx, "exampleVpcIpamPool", &ec2.VpcIpamPoolArgs{
			AddressFamily: pulumi.String("ipv4"),
			IpamScopeId:   exampleVpcIpam.PrivateDefaultScopeId,
			Locale:        *pulumi.String(current.Name),
		})
		if err != nil {
			return err
		}
		exampleVpcIpamPoolCidr, err := ec2.NewVpcIpamPoolCidr(ctx, "exampleVpcIpamPoolCidr", &ec2.VpcIpamPoolCidrArgs{
			IpamPoolId: exampleVpcIpamPool.ID(),
			Cidr:       pulumi.String("172.20.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpamPoolCidrAllocation(ctx, "exampleVpcIpamPoolCidrAllocation", &ec2.VpcIpamPoolCidrAllocationArgs{
			IpamPoolId:    exampleVpcIpamPool.ID(),
			NetmaskLength: pulumi.Int(28),
			DisallowedCidrs: pulumi.StringArray{
				pulumi.String("172.20.0.0/28"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleVpcIpamPoolCidr,
		}))
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
import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocation;
import com.pulumi.aws.ec2.VpcIpamPoolCidrAllocationArgs;
import com.pulumi.resources.CustomResourceOptions;
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

        var exampleVpcIpamPoolCidr = new VpcIpamPoolCidr("exampleVpcIpamPoolCidr", VpcIpamPoolCidrArgs.builder()        
            .ipamPoolId(exampleVpcIpamPool.id())
            .cidr("172.20.0.0/16")
            .build());

        var exampleVpcIpamPoolCidrAllocation = new VpcIpamPoolCidrAllocation("exampleVpcIpamPoolCidrAllocation", VpcIpamPoolCidrAllocationArgs.builder()        
            .ipamPoolId(exampleVpcIpamPool.id())
            .netmaskLength(28)
            .disallowedCidrs("172.20.0.0/28")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleVpcIpamPoolCidr)
                .build());

    }
}
```
```yaml
resources:
  exampleVpcIpamPoolCidrAllocation:
    type: aws:ec2:VpcIpamPoolCidrAllocation
    properties:
      ipamPoolId: ${exampleVpcIpamPool.id}
      netmaskLength: 28
      disallowedCidrs:
        - 172.20.0.0/28
    options:
      dependson:
        - ${exampleVpcIpamPoolCidr}
  exampleVpcIpamPoolCidr:
    type: aws:ec2:VpcIpamPoolCidr
    properties:
      ipamPoolId: ${exampleVpcIpamPool.id}
      cidr: 172.20.0.0/16
  exampleVpcIpamPool:
    type: aws:ec2:VpcIpamPool
    properties:
      addressFamily: ipv4
      ipamScopeId: ${exampleVpcIpam.privateDefaultScopeId}
      locale: ${current.name}
  exampleVpcIpam:
    type: aws:ec2:VpcIpam
    properties:
      operatingRegions:
        - regionName: ${current.name}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IPAM allocations using the allocation `id` and `pool id`, separated by `_`. For example:

```sh
 $ pulumi import aws:ec2/vpcIpamPoolCidrAllocation:VpcIpamPoolCidrAllocation example ipam-pool-alloc-0dc6d196509c049ba8b549ff99f639736_ipam-pool-07cfb559e0921fcbe
```
 