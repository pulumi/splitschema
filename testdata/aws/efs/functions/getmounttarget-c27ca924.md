Provides information about an Elastic File System Mount Target (EFS).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const mountTargetId = config.get("mountTargetId") || "";
const byId = aws.efs.getMountTarget({
    mountTargetId: mountTargetId,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
mount_target_id = config.get("mountTargetId")
if mount_target_id is None:
    mount_target_id = ""
by_id = aws.efs.get_mount_target(mount_target_id=mount_target_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var mountTargetId = config.Get("mountTargetId") ?? "";
    var byId = Aws.Efs.GetMountTarget.Invoke(new()
    {
        MountTargetId = mountTargetId,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		mountTargetId := ""
		if param := cfg.Get("mountTargetId"); param != "" {
			mountTargetId = param
		}
		_, err := efs.LookupMountTarget(ctx, &efs.LookupMountTargetArgs{
			MountTargetId: pulumi.StringRef(mountTargetId),
		}, nil)
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
import com.pulumi.aws.efs.EfsFunctions;
import com.pulumi.aws.efs.inputs.GetMountTargetArgs;
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
        final var config = ctx.config();
        final var mountTargetId = config.get("mountTargetId").orElse("");
        final var byId = EfsFunctions.getMountTarget(GetMountTargetArgs.builder()
            .mountTargetId(mountTargetId)
            .build());

    }
}
```
```yaml
configuration:
  mountTargetId:
    type: string
    default:
variables:
  byId:
    fn::invoke:
      Function: aws:efs:getMountTarget
      Arguments:
        mountTargetId: ${mountTargetId}
```
{{% /example %}}
{{% /examples %}}