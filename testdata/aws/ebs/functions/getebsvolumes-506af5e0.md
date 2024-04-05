`aws.ebs.getEbsVolumes` provides identifying information for EBS volumes matching given criteria.

This data source can be useful for getting a list of volume IDs with (for example) matching tags.

{{% examples %}}
## Example Usage
{{% example %}}

The following demonstrates obtaining a map of availability zone to EBS volume ID for volumes with a given tag value.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleEbsVolumes = aws.ebs.getEbsVolumes({
    tags: {
        VolumeSet: "TestVolumeSet",
    },
});
const exampleVolume = exampleEbsVolumes.then(exampleEbsVolumes => .map(([, ]) => (aws.ebs.getVolume({
    filters: [{
        name: "volume-id",
        values: [each.value],
    }],
}))));
export const availabilityZoneToVolumeId = exampleVolume.reduce((__obj, s) => ({ ...__obj, [s.id]: s.availabilityZone }));
```
```python
import pulumi
import pulumi_aws as aws

example_ebs_volumes = aws.ebs.get_ebs_volumes(tags={
    "VolumeSet": "TestVolumeSet",
})
example_volume = [aws.ebs.get_volume(filters=[aws.ebs.GetVolumeFilterArgs(
    name="volume-id",
    values=[each["value"]],
)]) for __key, __value in example_ebs_volumes.ids]
pulumi.export("availabilityZoneToVolumeId", {s.id: s.availability_zone for s in example_volume})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleEbsVolumes = Aws.Ebs.GetEbsVolumes.Invoke(new()
    {
        Tags = 
        {
            { "VolumeSet", "TestVolumeSet" },
        },
    });

    var exampleVolume = .Select(__value => 
    {
        return Aws.Ebs.GetVolume.Invoke(new()
        {
            Filters = new[]
            {
                new Aws.Ebs.Inputs.GetVolumeFilterInputArgs
                {
                    Name = "volume-id",
                    Values = new[]
                    {
                        each.Value,
                    },
                },
            },
        });
    }).ToList();

    return new Dictionary<string, object?>
    {
        ["availabilityZoneToVolumeId"] = exampleVolume,
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleEbsVolumes, err := ebs.GetEbsVolumes(ctx, &ebs.GetEbsVolumesArgs{
			Tags: map[string]interface{}{
				"VolumeSet": "TestVolumeSet",
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleVolume := "TODO: For expression"
		ctx.Export("availabilityZoneToVolumeId", "TODO: For expression")
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}