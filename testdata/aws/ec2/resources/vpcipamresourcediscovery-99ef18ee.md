Provides an IPAM Resource Discovery resource. IPAM Resource Discoveries are resources meant for multi-organization customers. If you wish to use a single IPAM across multiple orgs, a resource discovery can be created and shared from a subordinate organization to the management organizations IPAM delegated admin account. For a full deployment example, see `aws.ec2.VpcIpamResourceDiscoveryAssociation` resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const main = new aws.ec2.VpcIpamResourceDiscovery("main", {
    description: "My IPAM Resource Discovery",
    operatingRegions: [{
        regionName: current.then(current => current.name),
    }],
    tags: {
        Test: "Main",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
main = aws.ec2.VpcIpamResourceDiscovery("main",
    description="My IPAM Resource Discovery",
    operating_regions=[aws.ec2.VpcIpamResourceDiscoveryOperatingRegionArgs(
        region_name=current.name,
    )],
    tags={
        "Test": "Main",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var main = new Aws.Ec2.VpcIpamResourceDiscovery("main", new()
    {
        Description = "My IPAM Resource Discovery",
        OperatingRegions = new[]
        {
            new Aws.Ec2.Inputs.VpcIpamResourceDiscoveryOperatingRegionArgs
            {
                RegionName = current.Apply(getRegionResult => getRegionResult.Name),
            },
        },
        Tags = 
        {
            { "Test", "Main" },
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
		_, err = ec2.NewVpcIpamResourceDiscovery(ctx, "main", &ec2.VpcIpamResourceDiscoveryArgs{
			Description: pulumi.String("My IPAM Resource Discovery"),
			OperatingRegions: ec2.VpcIpamResourceDiscoveryOperatingRegionArray{
				&ec2.VpcIpamResourceDiscoveryOperatingRegionArgs{
					RegionName: *pulumi.String(current.Name),
				},
			},
			Tags: pulumi.StringMap{
				"Test": pulumi.String("Main"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.ec2.VpcIpamResourceDiscovery;
import com.pulumi.aws.ec2.VpcIpamResourceDiscoveryArgs;
import com.pulumi.aws.ec2.inputs.VpcIpamResourceDiscoveryOperatingRegionArgs;
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

        var main = new VpcIpamResourceDiscovery("main", VpcIpamResourceDiscoveryArgs.builder()        
            .description("My IPAM Resource Discovery")
            .operatingRegions(VpcIpamResourceDiscoveryOperatingRegionArgs.builder()
                .regionName(current.applyValue(getRegionResult -> getRegionResult.name()))
                .build())
            .tags(Map.of("Test", "Main"))
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:VpcIpamResourceDiscovery
    properties:
      description: My IPAM Resource Discovery
      operatingRegions:
        - regionName: ${current.name}
      tags:
        Test: Main
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IPAMs using the IPAM resource discovery `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcIpamResourceDiscovery:VpcIpamResourceDiscovery example ipam-res-disco-0178368ad2146a492
```
 