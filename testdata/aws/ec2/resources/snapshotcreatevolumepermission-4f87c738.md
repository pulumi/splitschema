Adds permission to create volumes off of a given EBS Snapshot.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    size: 40,
});
const exampleSnapshot = new aws.ebs.Snapshot("exampleSnapshot", {volumeId: example.id});
const examplePerm = new aws.ec2.SnapshotCreateVolumePermission("examplePerm", {
    snapshotId: exampleSnapshot.id,
    accountId: "12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=40)
example_snapshot = aws.ebs.Snapshot("exampleSnapshot", volume_id=example.id)
example_perm = aws.ec2.SnapshotCreateVolumePermission("examplePerm",
    snapshot_id=example_snapshot.id,
    account_id="12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ebs.Volume("example", new()
    {
        AvailabilityZone = "us-west-2a",
        Size = 40,
    });

    var exampleSnapshot = new Aws.Ebs.Snapshot("exampleSnapshot", new()
    {
        VolumeId = example.Id,
    });

    var examplePerm = new Aws.Ec2.SnapshotCreateVolumePermission("examplePerm", new()
    {
        SnapshotId = exampleSnapshot.Id,
        AccountId = "12345678",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(40),
		})
		if err != nil {
			return err
		}
		exampleSnapshot, err := ebs.NewSnapshot(ctx, "exampleSnapshot", &ebs.SnapshotArgs{
			VolumeId: example.ID(),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSnapshotCreateVolumePermission(ctx, "examplePerm", &ec2.SnapshotCreateVolumePermissionArgs{
			SnapshotId: exampleSnapshot.ID(),
			AccountId:  pulumi.String("12345678"),
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
import com.pulumi.aws.ebs.Volume;
import com.pulumi.aws.ebs.VolumeArgs;
import com.pulumi.aws.ebs.Snapshot;
import com.pulumi.aws.ebs.SnapshotArgs;
import com.pulumi.aws.ec2.SnapshotCreateVolumePermission;
import com.pulumi.aws.ec2.SnapshotCreateVolumePermissionArgs;
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
        var example = new Volume("example", VolumeArgs.builder()        
            .availabilityZone("us-west-2a")
            .size(40)
            .build());

        var exampleSnapshot = new Snapshot("exampleSnapshot", SnapshotArgs.builder()        
            .volumeId(example.id())
            .build());

        var examplePerm = new SnapshotCreateVolumePermission("examplePerm", SnapshotCreateVolumePermissionArgs.builder()        
            .snapshotId(exampleSnapshot.id())
            .accountId("12345678")
            .build());

    }
}
```
```yaml
resources:
  examplePerm:
    type: aws:ec2:SnapshotCreateVolumePermission
    properties:
      snapshotId: ${exampleSnapshot.id}
      accountId: '12345678'
  example:
    type: aws:ebs:Volume
    properties:
      availabilityZone: us-west-2a
      size: 40
  exampleSnapshot:
    type: aws:ebs:Snapshot
    properties:
      volumeId: ${example.id}
```
{{% /example %}}
{{% /examples %}}