List of Control Tower controls applied to an OU.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const thisOrganization = aws.organizations.getOrganization({});
const thisOrganizationalUnits = thisOrganization.then(thisOrganization => aws.organizations.getOrganizationalUnits({
    parentId: thisOrganization.roots?.[0]?.id,
}));
const thisControls = thisOrganizationalUnits.then(thisOrganizationalUnits => aws.controltower.getControls({
    targetIdentifier: .filter(x => x.name == "Security").map(x => (x.arn))[0],
}));
```
```python
import pulumi
import pulumi_aws as aws

this_organization = aws.organizations.get_organization()
this_organizational_units = aws.organizations.get_organizational_units(parent_id=this_organization.roots[0].id)
this_controls = aws.controltower.get_controls(target_identifier=[x.arn for x in this_organizational_units.children if x.name == "Security"][0])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var thisOrganization = Aws.Organizations.GetOrganization.Invoke();

    var thisOrganizationalUnits = Aws.Organizations.GetOrganizationalUnits.Invoke(new()
    {
        ParentId = thisOrganization.Apply(getOrganizationResult => getOrganizationResult.Roots[0]?.Id),
    });

    var thisControls = Aws.ControlTower.GetControls.Invoke(new()
    {
        TargetIdentifier = [0],
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/controltower"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		thisOrganization, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		thisOrganizationalUnits, err := organizations.GetOrganizationalUnits(ctx, &organizations.GetOrganizationalUnitsArgs{
			ParentId: thisOrganization.Roots[0].Id,
		}, nil)
		if err != nil {
			return err
		}
		_, err = controltower.GetControls(ctx, &controltower.GetControlsArgs{
			TargetIdentifier: "TODO: For expression"[0],
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}