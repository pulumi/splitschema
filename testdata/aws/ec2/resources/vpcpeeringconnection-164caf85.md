Provides a resource to manage a VPC peering connection.

> **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
both a standalone VPC Peering Connection Options and a VPC Peering Connection
resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
Doing so will cause a conflict of options and will overwrite the options.
Using a VPC Peering Connection Options resource decouples management of the connection options from
management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.

> **Note:** For cross-account (requester's AWS account differs from the accepter's AWS account) or inter-region
VPC Peering Connections use the `aws.ec2.VpcPeeringConnection` resource to manage the requester's side of the
connection and use the `aws.ec2.VpcPeeringConnectionAccepter` resource to manage the accepter's side of the connection.

> **Note:** Creating multiple `aws.ec2.VpcPeeringConnection` resources with the same `peer_vpc_id` and `vpc_id` will not produce an error. Instead, AWS will return the connection `id` that already exists, resulting in multiple `aws.ec2.VpcPeeringConnection` resources with the same `id`.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.VpcPeeringConnection("foo", {
    peerOwnerId: _var.peer_owner_id,
    peerVpcId: aws_vpc.bar.id,
    vpcId: aws_vpc.foo.id,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.VpcPeeringConnection("foo",
    peer_owner_id=var["peer_owner_id"],
    peer_vpc_id=aws_vpc["bar"]["id"],
    vpc_id=aws_vpc["foo"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.VpcPeeringConnection("foo", new()
    {
        PeerOwnerId = @var.Peer_owner_id,
        PeerVpcId = aws_vpc.Bar.Id,
        VpcId = aws_vpc.Foo.Id,
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
		_, err := ec2.NewVpcPeeringConnection(ctx, "foo", &ec2.VpcPeeringConnectionArgs{
			PeerOwnerId: pulumi.Any(_var.Peer_owner_id),
			PeerVpcId:   pulumi.Any(aws_vpc.Bar.Id),
			VpcId:       pulumi.Any(aws_vpc.Foo.Id),
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
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
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
        var foo = new VpcPeeringConnection("foo", VpcPeeringConnectionArgs.builder()        
            .peerOwnerId(var_.peer_owner_id())
            .peerVpcId(aws_vpc.bar().id())
            .vpcId(aws_vpc.foo().id())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ec2:VpcPeeringConnection
    properties:
      peerOwnerId: ${var.peer_owner_id}
      peerVpcId: ${aws_vpc.bar.id}
      vpcId: ${aws_vpc.foo.id}
```

Basic usage with connection options:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.VpcPeeringConnection("foo", {
    peerOwnerId: _var.peer_owner_id,
    peerVpcId: aws_vpc.bar.id,
    vpcId: aws_vpc.foo.id,
    accepter: {
        allowRemoteVpcDnsResolution: true,
    },
    requester: {
        allowRemoteVpcDnsResolution: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.VpcPeeringConnection("foo",
    peer_owner_id=var["peer_owner_id"],
    peer_vpc_id=aws_vpc["bar"]["id"],
    vpc_id=aws_vpc["foo"]["id"],
    accepter=aws.ec2.VpcPeeringConnectionAccepterArgs(
        allow_remote_vpc_dns_resolution=True,
    ),
    requester=aws.ec2.VpcPeeringConnectionRequesterArgs(
        allow_remote_vpc_dns_resolution=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.VpcPeeringConnection("foo", new()
    {
        PeerOwnerId = @var.Peer_owner_id,
        PeerVpcId = aws_vpc.Bar.Id,
        VpcId = aws_vpc.Foo.Id,
        Accepter = new Aws.Ec2.Inputs.VpcPeeringConnectionAccepterArgs
        {
            AllowRemoteVpcDnsResolution = true,
        },
        Requester = new Aws.Ec2.Inputs.VpcPeeringConnectionRequesterArgs
        {
            AllowRemoteVpcDnsResolution = true,
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
		_, err := ec2.NewVpcPeeringConnection(ctx, "foo", &ec2.VpcPeeringConnectionArgs{
			PeerOwnerId: pulumi.Any(_var.Peer_owner_id),
			PeerVpcId:   pulumi.Any(aws_vpc.Bar.Id),
			VpcId:       pulumi.Any(aws_vpc.Foo.Id),
			Accepter: &ec2.VpcPeeringConnectionAccepterTypeArgs{
				AllowRemoteVpcDnsResolution: pulumi.Bool(true),
			},
			Requester: &ec2.VpcPeeringConnectionRequesterArgs{
				AllowRemoteVpcDnsResolution: pulumi.Bool(true),
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
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionAccepterArgs;
import com.pulumi.aws.ec2.inputs.VpcPeeringConnectionRequesterArgs;
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
        var foo = new VpcPeeringConnection("foo", VpcPeeringConnectionArgs.builder()        
            .peerOwnerId(var_.peer_owner_id())
            .peerVpcId(aws_vpc.bar().id())
            .vpcId(aws_vpc.foo().id())
            .accepter(VpcPeeringConnectionAccepterArgs.builder()
                .allowRemoteVpcDnsResolution(true)
                .build())
            .requester(VpcPeeringConnectionRequesterArgs.builder()
                .allowRemoteVpcDnsResolution(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ec2:VpcPeeringConnection
    properties:
      peerOwnerId: ${var.peer_owner_id}
      peerVpcId: ${aws_vpc.bar.id}
      vpcId: ${aws_vpc.foo.id}
      accepter:
        allowRemoteVpcDnsResolution: true
      requester:
        allowRemoteVpcDnsResolution: true
```

Basic usage with tags:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"});
const bar = new aws.ec2.Vpc("bar", {cidrBlock: "10.2.0.0/16"});
const fooVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection", {
    peerOwnerId: _var.peer_owner_id,
    peerVpcId: bar.id,
    vpcId: fooVpc.id,
    autoAccept: true,
    tags: {
        Name: "VPC Peering between foo and bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16")
bar = aws.ec2.Vpc("bar", cidr_block="10.2.0.0/16")
foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
    peer_owner_id=var["peer_owner_id"],
    peer_vpc_id=bar.id,
    vpc_id=foo_vpc.id,
    auto_accept=True,
    tags={
        "Name": "VPC Peering between foo and bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooVpc = new Aws.Ec2.Vpc("fooVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var bar = new Aws.Ec2.Vpc("bar", new()
    {
        CidrBlock = "10.2.0.0/16",
    });

    var fooVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("fooVpcPeeringConnection", new()
    {
        PeerOwnerId = @var.Peer_owner_id,
        PeerVpcId = bar.Id,
        VpcId = fooVpc.Id,
        AutoAccept = true,
        Tags = 
        {
            { "Name", "VPC Peering between foo and bar" },
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
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		bar, err := ec2.NewVpc(ctx, "bar", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.2.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcPeeringConnection(ctx, "fooVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
			PeerOwnerId: pulumi.Any(_var.Peer_owner_id),
			PeerVpcId:   bar.ID(),
			VpcId:       fooVpc.ID(),
			AutoAccept:  pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("VPC Peering between foo and bar"),
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
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
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
        var fooVpc = new Vpc("fooVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        var bar = new Vpc("bar", VpcArgs.builder()        
            .cidrBlock("10.2.0.0/16")
            .build());

        var fooVpcPeeringConnection = new VpcPeeringConnection("fooVpcPeeringConnection", VpcPeeringConnectionArgs.builder()        
            .peerOwnerId(var_.peer_owner_id())
            .peerVpcId(bar.id())
            .vpcId(fooVpc.id())
            .autoAccept(true)
            .tags(Map.of("Name", "VPC Peering between foo and bar"))
            .build());

    }
}
```
```yaml
resources:
  fooVpcPeeringConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      peerOwnerId: ${var.peer_owner_id}
      peerVpcId: ${bar.id}
      vpcId: ${fooVpc.id}
      autoAccept: true
      tags:
        Name: VPC Peering between foo and bar
  fooVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  bar:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.2.0.0/16
```

Basic usage with region:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"}, {
    provider: aws["us-west-2"],
});
const bar = new aws.ec2.Vpc("bar", {cidrBlock: "10.2.0.0/16"}, {
    provider: aws["us-east-1"],
});
const fooVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection", {
    peerOwnerId: _var.peer_owner_id,
    peerVpcId: bar.id,
    vpcId: fooVpc.id,
    peerRegion: "us-east-1",
});
```
```python
import pulumi
import pulumi_aws as aws

foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.1.0.0/16",
opts=pulumi.ResourceOptions(provider=aws["us-west-2"]))
bar = aws.ec2.Vpc("bar", cidr_block="10.2.0.0/16",
opts=pulumi.ResourceOptions(provider=aws["us-east-1"]))
foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
    peer_owner_id=var["peer_owner_id"],
    peer_vpc_id=bar.id,
    vpc_id=foo_vpc.id,
    peer_region="us-east-1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooVpc = new Aws.Ec2.Vpc("fooVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = aws.Us_west_2,
    });

    var bar = new Aws.Ec2.Vpc("bar", new()
    {
        CidrBlock = "10.2.0.0/16",
    }, new CustomResourceOptions
    {
        Provider = aws.Us_east_1,
    });

    var fooVpcPeeringConnection = new Aws.Ec2.VpcPeeringConnection("fooVpcPeeringConnection", new()
    {
        PeerOwnerId = @var.Peer_owner_id,
        PeerVpcId = bar.Id,
        VpcId = fooVpc.Id,
        PeerRegion = "us-east-1",
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
		fooVpc, err := ec2.NewVpc(ctx, "fooVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		}, pulumi.Provider(aws.UsWest2))
		if err != nil {
			return err
		}
		bar, err := ec2.NewVpc(ctx, "bar", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.2.0.0/16"),
		}, pulumi.Provider(aws.UsEast1))
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcPeeringConnection(ctx, "fooVpcPeeringConnection", &ec2.VpcPeeringConnectionArgs{
			PeerOwnerId: pulumi.Any(_var.Peer_owner_id),
			PeerVpcId:   bar.ID(),
			VpcId:       fooVpc.ID(),
			PeerRegion:  pulumi.String("us-east-1"),
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
import com.pulumi.aws.ec2.VpcPeeringConnection;
import com.pulumi.aws.ec2.VpcPeeringConnectionArgs;
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
        var fooVpc = new Vpc("fooVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-west-2())
                .build());

        var bar = new Vpc("bar", VpcArgs.builder()        
            .cidrBlock("10.2.0.0/16")
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-east-1())
                .build());

        var fooVpcPeeringConnection = new VpcPeeringConnection("fooVpcPeeringConnection", VpcPeeringConnectionArgs.builder()        
            .peerOwnerId(var_.peer_owner_id())
            .peerVpcId(bar.id())
            .vpcId(fooVpc.id())
            .peerRegion("us-east-1")
            .build());

    }
}
```
```yaml
resources:
  fooVpcPeeringConnection:
    type: aws:ec2:VpcPeeringConnection
    properties:
      peerOwnerId: ${var.peer_owner_id}
      peerVpcId: ${bar.id}
      vpcId: ${fooVpc.id}
      peerRegion: us-east-1
  fooVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
    options:
      provider: ${aws"us-west-2"[%!s(MISSING)]}
  bar:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.2.0.0/16
    options:
      provider: ${aws"us-east-1"[%!s(MISSING)]}
```
{{% /example %}}
{{% /examples %}}
## Notes

If both VPCs are not in the same AWS account and region do not enable the `auto_accept` attribute.
The accepter can manage its side of the connection using the `aws.ec2.VpcPeeringConnectionAccepter` resource
or accept the connection manually using the AWS Management Console, AWS CLI, through SDKs, etc.


## Import

Using `pulumi import`, import VPC Peering resources using the VPC peering `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcPeeringConnection:VpcPeeringConnection test_connection pcx-111aaa111
```
 