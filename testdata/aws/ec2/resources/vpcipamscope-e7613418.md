Creates a scope for AWS IPAM.

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
const exampleVpcIpamScope = new aws.ec2.VpcIpamScope("exampleVpcIpamScope", {
    ipamId: exampleVpcIpam.id,
    description: "Another Scope",
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example_vpc_ipam = aws.ec2.VpcIpam("exampleVpcIpam", operating_regions=[aws.ec2.VpcIpamOperatingRegionArgs(
    region_name=current.name,
)])
example_vpc_ipam_scope = aws.ec2.VpcIpamScope("exampleVpcIpamScope",
    ipam_id=example_vpc_ipam.id,
    description="Another Scope")
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

    var exampleVpcIpamScope = new Aws.Ec2.VpcIpamScope("exampleVpcIpamScope", new()
    {
        IpamId = exampleVpcIpam.Id,
        Description = "Another Scope",
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
		_, err = ec2.NewVpcIpamScope(ctx, "exampleVpcIpamScope", &ec2.VpcIpamScopeArgs{
			IpamId:      exampleVpcIpam.ID(),
			Description: pulumi.String("Another Scope"),
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
import com.pulumi.aws.ec2.VpcIpamScope;
import com.pulumi.aws.ec2.VpcIpamScopeArgs;
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

        var exampleVpcIpamScope = new VpcIpamScope("exampleVpcIpamScope", VpcIpamScopeArgs.builder()        
            .ipamId(exampleVpcIpam.id())
            .description("Another Scope")
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
  exampleVpcIpamScope:
    type: aws:ec2:VpcIpamScope
    properties:
      ipamId: ${exampleVpcIpam.id}
      description: Another Scope
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IPAMs using the `scope_id`. For example:

```sh
 $ pulumi import aws:ec2/vpcIpamScope:VpcIpamScope example ipam-scope-0513c69f283d11dfb
```
 