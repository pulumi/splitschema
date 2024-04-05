Creates a snapshot copy grant that allows AWS Redshift to encrypt copied snapshots with a customer master key from AWS KMS in a destination region.

Note that the grant must exist in the destination region, and not in the region of the cluster.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testSnapshotCopyGrant = new aws.redshift.SnapshotCopyGrant("testSnapshotCopyGrant", {snapshotCopyGrantName: "my-grant"});
const testCluster = new aws.redshift.Cluster("testCluster", {snapshotCopy: {
    destinationRegion: "us-east-2",
    grantName: testSnapshotCopyGrant.snapshotCopyGrantName,
}});
```
```python
import pulumi
import pulumi_aws as aws

test_snapshot_copy_grant = aws.redshift.SnapshotCopyGrant("testSnapshotCopyGrant", snapshot_copy_grant_name="my-grant")
test_cluster = aws.redshift.Cluster("testCluster", snapshot_copy=aws.redshift.ClusterSnapshotCopyArgs(
    destination_region="us-east-2",
    grant_name=test_snapshot_copy_grant.snapshot_copy_grant_name,
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testSnapshotCopyGrant = new Aws.RedShift.SnapshotCopyGrant("testSnapshotCopyGrant", new()
    {
        SnapshotCopyGrantName = "my-grant",
    });

    var testCluster = new Aws.RedShift.Cluster("testCluster", new()
    {
        SnapshotCopy = new Aws.RedShift.Inputs.ClusterSnapshotCopyArgs
        {
            DestinationRegion = "us-east-2",
            GrantName = testSnapshotCopyGrant.SnapshotCopyGrantName,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testSnapshotCopyGrant, err := redshift.NewSnapshotCopyGrant(ctx, "testSnapshotCopyGrant", &redshift.SnapshotCopyGrantArgs{
			SnapshotCopyGrantName: pulumi.String("my-grant"),
		})
		if err != nil {
			return err
		}
		_, err = redshift.NewCluster(ctx, "testCluster", &redshift.ClusterArgs{
			SnapshotCopy: &redshift.ClusterSnapshotCopyArgs{
				DestinationRegion: pulumi.String("us-east-2"),
				GrantName:         testSnapshotCopyGrant.SnapshotCopyGrantName,
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
import com.pulumi.aws.redshift.SnapshotCopyGrant;
import com.pulumi.aws.redshift.SnapshotCopyGrantArgs;
import com.pulumi.aws.redshift.Cluster;
import com.pulumi.aws.redshift.ClusterArgs;
import com.pulumi.aws.redshift.inputs.ClusterSnapshotCopyArgs;
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
        var testSnapshotCopyGrant = new SnapshotCopyGrant("testSnapshotCopyGrant", SnapshotCopyGrantArgs.builder()        
            .snapshotCopyGrantName("my-grant")
            .build());

        var testCluster = new Cluster("testCluster", ClusterArgs.builder()        
            .snapshotCopy(ClusterSnapshotCopyArgs.builder()
                .destinationRegion("us-east-2")
                .grantName(testSnapshotCopyGrant.snapshotCopyGrantName())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testSnapshotCopyGrant:
    type: aws:redshift:SnapshotCopyGrant
    properties:
      snapshotCopyGrantName: my-grant
  testCluster:
    type: aws:redshift:Cluster
    properties:
      snapshotCopy:
        destinationRegion: us-east-2
        grantName: ${testSnapshotCopyGrant.snapshotCopyGrantName}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Snapshot Copy Grants by name. For example:

```sh
 $ pulumi import aws:redshift/snapshotCopyGrant:SnapshotCopyGrant test my-grant
```
 