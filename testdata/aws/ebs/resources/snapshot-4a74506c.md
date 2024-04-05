Creates a Snapshot of an EBS Volume.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    size: 40,
    tags: {
        Name: "HelloWorld",
    },
});
const exampleSnapshot = new aws.ebs.Snapshot("exampleSnapshot", {
    volumeId: example.id,
    tags: {
        Name: "HelloWorld_snap",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=40,
    tags={
        "Name": "HelloWorld",
    })
example_snapshot = aws.ebs.Snapshot("exampleSnapshot",
    volume_id=example.id,
    tags={
        "Name": "HelloWorld_snap",
    })
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
        Tags = 
        {
            { "Name", "HelloWorld" },
        },
    });

    var exampleSnapshot = new Aws.Ebs.Snapshot("exampleSnapshot", new()
    {
        VolumeId = example.Id,
        Tags = 
        {
            { "Name", "HelloWorld_snap" },
        },
    });

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
		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(40),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ebs.NewSnapshot(ctx, "exampleSnapshot", &ebs.SnapshotArgs{
			VolumeId: example.ID(),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld_snap"),
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
import com.pulumi.aws.ebs.Volume;
import com.pulumi.aws.ebs.VolumeArgs;
import com.pulumi.aws.ebs.Snapshot;
import com.pulumi.aws.ebs.SnapshotArgs;
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
            .tags(Map.of("Name", "HelloWorld"))
            .build());

        var exampleSnapshot = new Snapshot("exampleSnapshot", SnapshotArgs.builder()        
            .volumeId(example.id())
            .tags(Map.of("Name", "HelloWorld_snap"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ebs:Volume
    properties:
      availabilityZone: us-west-2a
      size: 40
      tags:
        Name: HelloWorld
  exampleSnapshot:
    type: aws:ebs:Snapshot
    properties:
      volumeId: ${example.id}
      tags:
        Name: HelloWorld_snap
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EBS Snapshot using the `id`. For example:

```sh
 $ pulumi import aws:ebs/snapshot:Snapshot id snap-049df61146c4d7901
```
 