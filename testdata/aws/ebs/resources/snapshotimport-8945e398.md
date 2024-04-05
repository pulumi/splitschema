Imports a disk image from S3 as a Snapshot.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.SnapshotImport("example", {
    diskContainer: {
        format: "VHD",
        userBucket: {
            s3Bucket: "disk-images",
            s3Key: "source.vhd",
        },
    },
    roleName: "disk-image-import",
    tags: {
        Name: "HelloWorld",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ebs.SnapshotImport("example",
    disk_container=aws.ebs.SnapshotImportDiskContainerArgs(
        format="VHD",
        user_bucket=aws.ebs.SnapshotImportDiskContainerUserBucketArgs(
            s3_bucket="disk-images",
            s3_key="source.vhd",
        ),
    ),
    role_name="disk-image-import",
    tags={
        "Name": "HelloWorld",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ebs.SnapshotImport("example", new()
    {
        DiskContainer = new Aws.Ebs.Inputs.SnapshotImportDiskContainerArgs
        {
            Format = "VHD",
            UserBucket = new Aws.Ebs.Inputs.SnapshotImportDiskContainerUserBucketArgs
            {
                S3Bucket = "disk-images",
                S3Key = "source.vhd",
            },
        },
        RoleName = "disk-image-import",
        Tags = 
        {
            { "Name", "HelloWorld" },
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
		_, err := ebs.NewSnapshotImport(ctx, "example", &ebs.SnapshotImportArgs{
			DiskContainer: &ebs.SnapshotImportDiskContainerArgs{
				Format: pulumi.String("VHD"),
				UserBucket: &ebs.SnapshotImportDiskContainerUserBucketArgs{
					S3Bucket: pulumi.String("disk-images"),
					S3Key:    pulumi.String("source.vhd"),
				},
			},
			RoleName: pulumi.String("disk-image-import"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
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
import com.pulumi.aws.ebs.SnapshotImport;
import com.pulumi.aws.ebs.SnapshotImportArgs;
import com.pulumi.aws.ebs.inputs.SnapshotImportDiskContainerArgs;
import com.pulumi.aws.ebs.inputs.SnapshotImportDiskContainerUserBucketArgs;
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
        var example = new SnapshotImport("example", SnapshotImportArgs.builder()        
            .diskContainer(SnapshotImportDiskContainerArgs.builder()
                .format("VHD")
                .userBucket(SnapshotImportDiskContainerUserBucketArgs.builder()
                    .s3Bucket("disk-images")
                    .s3Key("source.vhd")
                    .build())
                .build())
            .roleName("disk-image-import")
            .tags(Map.of("Name", "HelloWorld"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ebs:SnapshotImport
    properties:
      diskContainer:
        format: VHD
        userBucket:
          s3Bucket: disk-images
          s3Key: source.vhd
      roleName: disk-image-import
      tags:
        Name: HelloWorld
```
{{% /example %}}
{{% /examples %}}