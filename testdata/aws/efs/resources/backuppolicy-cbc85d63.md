Provides an Elastic File System (EFS) Backup Policy resource.
Backup policies turn automatic backups on or off for an existing file system.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fs = new aws.efs.FileSystem("fs", {});
const policy = new aws.efs.BackupPolicy("policy", {
    fileSystemId: fs.id,
    backupPolicy: {
        status: "ENABLED",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

fs = aws.efs.FileSystem("fs")
policy = aws.efs.BackupPolicy("policy",
    file_system_id=fs.id,
    backup_policy=aws.efs.BackupPolicyBackupPolicyArgs(
        status="ENABLED",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fs = new Aws.Efs.FileSystem("fs");

    var policy = new Aws.Efs.BackupPolicy("policy", new()
    {
        FileSystemId = fs.Id,
        BackupPolicyDetails = new Aws.Efs.Inputs.BackupPolicyBackupPolicyArgs
        {
            Status = "ENABLED",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		fs, err := efs.NewFileSystem(ctx, "fs", nil)
		if err != nil {
			return err
		}
		_, err = efs.NewBackupPolicy(ctx, "policy", &efs.BackupPolicyArgs{
			FileSystemId: fs.ID(),
			BackupPolicy: &efs.BackupPolicyBackupPolicyArgs{
				Status: pulumi.String("ENABLED"),
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
import com.pulumi.aws.efs.FileSystem;
import com.pulumi.aws.efs.BackupPolicy;
import com.pulumi.aws.efs.BackupPolicyArgs;
import com.pulumi.aws.efs.inputs.BackupPolicyBackupPolicyArgs;
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
        var fs = new FileSystem("fs");

        var policy = new BackupPolicy("policy", BackupPolicyArgs.builder()        
            .fileSystemId(fs.id())
            .backupPolicy(BackupPolicyBackupPolicyArgs.builder()
                .status("ENABLED")
                .build())
            .build());

    }
}
```
```yaml
resources:
  fs:
    type: aws:efs:FileSystem
  policy:
    type: aws:efs:BackupPolicy
    properties:
      fileSystemId: ${fs.id}
      backupPolicy:
        status: ENABLED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the EFS backup policies using the `id`. For example:

```sh
 $ pulumi import aws:efs/backupPolicy:BackupPolicy example fs-6fa144c6
```
 