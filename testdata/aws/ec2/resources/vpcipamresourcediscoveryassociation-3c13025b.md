Provides an association between an Amazon IP Address Manager (IPAM) and a IPAM Resource Discovery. IPAM Resource Discoveries are resources meant for multi-organization customers. If you wish to use a single IPAM across multiple orgs, a resource discovery can be created and shared from a subordinate organization to the management organizations IPAM delegated admin account.

Once an association is created between two organizations via IPAM & a IPAM Resource Discovery, IPAM Pools can be shared via Resource Access Manager (RAM) to accounts in the subordinate organization; these RAM shares must be accepted by the end user account. Pools can then also discover and monitor IPAM resources in the subordinate organization.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.VpcIpamResourceDiscoveryAssociation("test", {
    ipamId: aws_vpc_ipam.test.id,
    ipamResourceDiscoveryId: aws_vpc_ipam_resource_discovery.test.id,
    tags: {
        Name: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.VpcIpamResourceDiscoveryAssociation("test",
    ipam_id=aws_vpc_ipam["test"]["id"],
    ipam_resource_discovery_id=aws_vpc_ipam_resource_discovery["test"]["id"],
    tags={
        "Name": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ec2.VpcIpamResourceDiscoveryAssociation("test", new()
    {
        IpamId = aws_vpc_ipam.Test.Id,
        IpamResourceDiscoveryId = aws_vpc_ipam_resource_discovery.Test.Id,
        Tags = 
        {
            { "Name", "test" },
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
		_, err := ec2.NewVpcIpamResourceDiscoveryAssociation(ctx, "test", &ec2.VpcIpamResourceDiscoveryAssociationArgs{
			IpamId:                  pulumi.Any(aws_vpc_ipam.Test.Id),
			IpamResourceDiscoveryId: pulumi.Any(aws_vpc_ipam_resource_discovery.Test.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
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
import com.pulumi.aws.ec2.VpcIpamResourceDiscoveryAssociation;
import com.pulumi.aws.ec2.VpcIpamResourceDiscoveryAssociationArgs;
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
        var test = new VpcIpamResourceDiscoveryAssociation("test", VpcIpamResourceDiscoveryAssociationArgs.builder()        
            .ipamId(aws_vpc_ipam.test().id())
            .ipamResourceDiscoveryId(aws_vpc_ipam_resource_discovery.test().id())
            .tags(Map.of("Name", "test"))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ec2:VpcIpamResourceDiscoveryAssociation
    properties:
      ipamId: ${aws_vpc_ipam.test.id}
      ipamResourceDiscoveryId: ${aws_vpc_ipam_resource_discovery.test.id}
      tags:
        Name: test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IPAMs using the IPAM resource discovery association `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcIpamResourceDiscoveryAssociation:VpcIpamResourceDiscoveryAssociation example ipam-res-disco-assoc-0178368ad2146a492
```
 