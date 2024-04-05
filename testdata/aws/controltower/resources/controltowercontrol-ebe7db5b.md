Allows the application of pre-defined controls to organizational units. For more information on usage, please see the
[AWS Control Tower User Guide](https://docs.aws.amazon.com/controltower/latest/userguide/enable-guardrails.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const exampleOrganization = aws.organizations.getOrganization({});
const exampleOrganizationalUnits = exampleOrganization.then(exampleOrganization => aws.organizations.getOrganizationalUnits({
    parentId: exampleOrganization.roots?.[0]?.id,
}));
const exampleControlTowerControl = new aws.controltower.ControlTowerControl("exampleControlTowerControl", {
    controlIdentifier: current.then(current => `arn:aws:controltower:${current.name}::control/AWS-GR_EC2_VOLUME_INUSE_CHECK`),
    targetIdentifier: exampleOrganizationalUnits.then(exampleOrganizationalUnits => .filter(x => x.name == "Infrastructure").map(x => (x.arn))[0]),
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example_organization = aws.organizations.get_organization()
example_organizational_units = aws.organizations.get_organizational_units(parent_id=example_organization.roots[0].id)
example_control_tower_control = aws.controltower.ControlTowerControl("exampleControlTowerControl",
    control_identifier=f"arn:aws:controltower:{current.name}::control/AWS-GR_EC2_VOLUME_INUSE_CHECK",
    target_identifier=[x.arn for x in example_organizational_units.children if x.name == "Infrastructure"][0])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var exampleOrganization = Aws.Organizations.GetOrganization.Invoke();

    var exampleOrganizationalUnits = Aws.Organizations.GetOrganizationalUnits.Invoke(new()
    {
        ParentId = exampleOrganization.Apply(getOrganizationResult => getOrganizationResult.Roots[0]?.Id),
    });

    var exampleControlTowerControl = new Aws.ControlTower.ControlTowerControl("exampleControlTowerControl", new()
    {
        ControlIdentifier = $"arn:aws:controltower:{current.Apply(getRegionResult => getRegionResult.Name)}::control/AWS-GR_EC2_VOLUME_INUSE_CHECK",
        TargetIdentifier = .Where(x => x.Name == "Infrastructure").Select(x => 
        {
            return x.Arn;
        }).ToList()[0],
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/controltower"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleOrganization, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleOrganizationalUnits, err := organizations.GetOrganizationalUnits(ctx, &organizations.GetOrganizationalUnitsArgs{
			ParentId: exampleOrganization.Roots[0].Id,
		}, nil)
		if err != nil {
			return err
		}
		_, err = controltower.NewControlTowerControl(ctx, "exampleControlTowerControl", &controltower.ControlTowerControlArgs{
			ControlIdentifier: pulumi.String(fmt.Sprintf("arn:aws:controltower:%v::control/AWS-GR_EC2_VOLUME_INUSE_CHECK", current.Name)),
			TargetIdentifier:  "TODO: For expression"[0],
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Control Tower Controls using their `organizational_unit_arn/control_identifier`. For example:

```sh
 $ pulumi import aws:controltower/controlTowerControl:ControlTowerControl example arn:aws:organizations::123456789101:ou/o-qqaejywet/ou-qg5o-ufbhdtv3,arn:aws:controltower:us-east-1::control/WTDSMKDKDNLE
```
 