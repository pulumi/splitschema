Provides a VPC resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
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
		_, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
```

Basic usage with tags:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {
    cidrBlock: "10.0.0.0/16",
    instanceTenancy: "default",
    tags: {
        Name: "main",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main",
    cidr_block="10.0.0.0/16",
    instance_tenancy="default",
    tags={
        "Name": "main",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
        InstanceTenancy = "default",
        Tags = 
        {
            { "Name", "main" },
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
		_, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock:       pulumi.String("10.0.0.0/16"),
			InstanceTenancy: pulumi.String("default"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("main"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .instanceTenancy("default")
            .tags(Map.of("Name", "main"))
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
      instanceTenancy: default
      tags:
        Name: main
```

VPC with CIDR from AWS IPAM:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const testVpcIpam = new aws.ec2.VpcIpam("testVpcIpam", {operatingRegions: [{
    regionName: current.then(current => current.name),
}]});
const testVpcIpamPool = new aws.ec2.VpcIpamPool("testVpcIpamPool", {
    addressFamily: "ipv4",
    ipamScopeId: testVpcIpam.privateDefaultScopeId,
    locale: current.then(current => current.name),
});
const testVpcIpamPoolCidr = new aws.ec2.VpcIpamPoolCidr("testVpcIpamPoolCidr", {
    ipamPoolId: testVpcIpamPool.id,
    cidr: "172.20.0.0/16",
});
const testVpc = new aws.ec2.Vpc("testVpc", {
    ipv4IpamPoolId: testVpcIpamPool.id,
    ipv4NetmaskLength: 28,
}, {
    dependsOn: [testVpcIpamPoolCidr],
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
test_vpc_ipam = aws.ec2.VpcIpam("testVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
    region_name=current.name,
)])
test_vpc_ipam_pool = aws.ec2.VpcIpamPool("testVpcIpamPool",
    address_family="ipv4",
    ipam_scope_id=test_vpc_ipam.private_default_scope_id,
    locale=current.name)
test_vpc_ipam_pool_cidr = aws.ec2.VpcIpamPoolCidr("testVpcIpamPoolCidr",
    ipam_pool_id=test_vpc_ipam_pool.id,
    cidr="172.20.0.0/16")
test_vpc = aws.ec2.Vpc("testVpc",
    ipv4_ipam_pool_id=test_vpc_ipam_pool.id,
    ipv4_netmask_length=28,
    opts=pulumi.ResourceOptions(depends_on=[test_vpc_ipam_pool_cidr]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var testVpcIpam = new Aws.Ec2.VpcIpam("testVpcIpam", new()
    {
        OperatingRegions = new[]
        {
            new Aws.Ec2.Inputs.VpcIpamOperatingRegionArgs
            {
                RegionName = current.Apply(getRegionResult => getRegionResult.Name),
            },
        },
    });

    var testVpcIpamPool = new Aws.Ec2.VpcIpamPool("testVpcIpamPool", new()
    {
        AddressFamily = "ipv4",
        IpamScopeId = testVpcIpam.PrivateDefaultScopeId,
        Locale = current.Apply(getRegionResult => getRegionResult.Name),
    });

    var testVpcIpamPoolCidr = new Aws.Ec2.VpcIpamPoolCidr("testVpcIpamPoolCidr", new()
    {
        IpamPoolId = testVpcIpamPool.Id,
        Cidr = "172.20.0.0/16",
    });

    var testVpc = new Aws.Ec2.Vpc("testVpc", new()
    {
        Ipv4IpamPoolId = testVpcIpamPool.Id,
        Ipv4NetmaskLength = 28,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            testVpcIpamPoolCidr,
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
		testVpcIpam, err := ec2.NewVpcIpam(ctx, "testVpcIpam", &ec2.VpcIpamArgs{
			OperatingRegions: ec2.VpcIpamOperatingRegionArray{
				&ec2.VpcIpamOperatingRegionArgs{
					RegionName: *pulumi.String(current.Name),
				},
			},
		})
		if err != nil {
			return err
		}
		testVpcIpamPool, err := ec2.NewVpcIpamPool(ctx, "testVpcIpamPool", &ec2.VpcIpamPoolArgs{
			AddressFamily: pulumi.String("ipv4"),
			IpamScopeId:   testVpcIpam.PrivateDefaultScopeId,
			Locale:        *pulumi.String(current.Name),
		})
		if err != nil {
			return err
		}
		testVpcIpamPoolCidr, err := ec2.NewVpcIpamPoolCidr(ctx, "testVpcIpamPoolCidr", &ec2.VpcIpamPoolCidrArgs{
			IpamPoolId: testVpcIpamPool.ID(),
			Cidr:       pulumi.String("172.20.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpc(ctx, "testVpc", &ec2.VpcArgs{
			Ipv4IpamPoolId:    testVpcIpamPool.ID(),
			Ipv4NetmaskLength: pulumi.Int(28),
		}, pulumi.DependsOn([]pulumi.Resource{
			testVpcIpamPoolCidr,
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
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

        var testVpcIpam = new VpcIpam("testVpcIpam", VpcIpamArgs.builder()        
            .operatingRegions(VpcIpamOperatingRegionArgs.builder()
                .regionName(current.applyValue(getRegionResult -> getRegionResult.name()))
                .build())
            .build());

        var testVpcIpamPool = new VpcIpamPool("testVpcIpamPool", VpcIpamPoolArgs.builder()        
            .addressFamily("ipv4")
            .ipamScopeId(testVpcIpam.privateDefaultScopeId())
            .locale(current.applyValue(getRegionResult -> getRegionResult.name()))
            .build());

        var testVpcIpamPoolCidr = new VpcIpamPoolCidr("testVpcIpamPoolCidr", VpcIpamPoolCidrArgs.builder()        
            .ipamPoolId(testVpcIpamPool.id())
            .cidr("172.20.0.0/16")
            .build());

        var testVpc = new Vpc("testVpc", VpcArgs.builder()        
            .ipv4IpamPoolId(testVpcIpamPool.id())
            .ipv4NetmaskLength(28)
            .build(), CustomResourceOptions.builder()
                .dependsOn(testVpcIpamPoolCidr)
                .build());

    }
}
```
```yaml
resources:
  testVpcIpam:
    type: aws:ec2:VpcIpam
    properties:
      operatingRegions:
        - regionName: ${current.name}
  testVpcIpamPool:
    type: aws:ec2:VpcIpamPool
    properties:
      addressFamily: ipv4
      ipamScopeId: ${testVpcIpam.privateDefaultScopeId}
      locale: ${current.name}
  testVpcIpamPoolCidr:
    type: aws:ec2:VpcIpamPoolCidr
    properties:
      ipamPoolId: ${testVpcIpamPool.id}
      cidr: 172.20.0.0/16
  testVpc:
    type: aws:ec2:Vpc
    properties:
      ipv4IpamPoolId: ${testVpcIpamPool.id}
      ipv4NetmaskLength: 28
    options:
      dependson:
        - ${testVpcIpamPoolCidr}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPCs using the VPC `id`. For example:

```sh
 $ pulumi import aws:ec2/vpc:Vpc test_vpc vpc-a01106c2
```
 