Manages an AWS DataSync FSx OpenZfs Location.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.FsxOpenZfsFileSystem("example", {
    fsxFilesystemArn: aws_fsx_openzfs_file_system.example.arn,
    securityGroupArns: [aws_security_group.example.arn],
    protocol: {
        nfs: {
            mountOptions: {
                version: "AUTOMATIC",
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.FsxOpenZfsFileSystem("example",
    fsx_filesystem_arn=aws_fsx_openzfs_file_system["example"]["arn"],
    security_group_arns=[aws_security_group["example"]["arn"]],
    protocol=aws.datasync.FsxOpenZfsFileSystemProtocolArgs(
        nfs=aws.datasync.FsxOpenZfsFileSystemProtocolNfsArgs(
            mount_options=aws.datasync.FsxOpenZfsFileSystemProtocolNfsMountOptionsArgs(
                version="AUTOMATIC",
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.FsxOpenZfsFileSystem("example", new()
    {
        FsxFilesystemArn = aws_fsx_openzfs_file_system.Example.Arn,
        SecurityGroupArns = new[]
        {
            aws_security_group.Example.Arn,
        },
        Protocol = new Aws.DataSync.Inputs.FsxOpenZfsFileSystemProtocolArgs
        {
            Nfs = new Aws.DataSync.Inputs.FsxOpenZfsFileSystemProtocolNfsArgs
            {
                MountOptions = new Aws.DataSync.Inputs.FsxOpenZfsFileSystemProtocolNfsMountOptionsArgs
                {
                    Version = "AUTOMATIC",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewFsxOpenZfsFileSystem(ctx, "example", &datasync.FsxOpenZfsFileSystemArgs{
			FsxFilesystemArn: pulumi.Any(aws_fsx_openzfs_file_system.Example.Arn),
			SecurityGroupArns: pulumi.StringArray{
				aws_security_group.Example.Arn,
			},
			Protocol: &datasync.FsxOpenZfsFileSystemProtocolArgs{
				Nfs: &datasync.FsxOpenZfsFileSystemProtocolNfsArgs{
					MountOptions: &datasync.FsxOpenZfsFileSystemProtocolNfsMountOptionsArgs{
						Version: pulumi.String("AUTOMATIC"),
					},
				},
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
import com.pulumi.aws.datasync.FsxOpenZfsFileSystem;
import com.pulumi.aws.datasync.FsxOpenZfsFileSystemArgs;
import com.pulumi.aws.datasync.inputs.FsxOpenZfsFileSystemProtocolArgs;
import com.pulumi.aws.datasync.inputs.FsxOpenZfsFileSystemProtocolNfsArgs;
import com.pulumi.aws.datasync.inputs.FsxOpenZfsFileSystemProtocolNfsMountOptionsArgs;
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
        var example = new FsxOpenZfsFileSystem("example", FsxOpenZfsFileSystemArgs.builder()        
            .fsxFilesystemArn(aws_fsx_openzfs_file_system.example().arn())
            .securityGroupArns(aws_security_group.example().arn())
            .protocol(FsxOpenZfsFileSystemProtocolArgs.builder()
                .nfs(FsxOpenZfsFileSystemProtocolNfsArgs.builder()
                    .mountOptions(FsxOpenZfsFileSystemProtocolNfsMountOptionsArgs.builder()
                        .version("AUTOMATIC")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:FsxOpenZfsFileSystem
    properties:
      fsxFilesystemArn: ${aws_fsx_openzfs_file_system.example.arn}
      securityGroupArns:
        - ${aws_security_group.example.arn}
      protocol:
        nfs:
          mountOptions:
            version: AUTOMATIC
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_fsx_openzfs_file_system` using the `DataSync-ARN#FSx-openzfs-ARN`. For example:

```sh
 $ pulumi import aws:datasync/fsxOpenZfsFileSystem:FsxOpenZfsFileSystem example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:123456789012:file-system/fs-08e04cd442c1bb94a
```
 