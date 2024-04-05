Provides an SSM Patch Group resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = new aws.ssm.PatchBaseline("production", {approvedPatches: ["KB123456"]});
const patchgroup = new aws.ssm.PatchGroup("patchgroup", {
    baselineId: production.id,
    patchGroup: "patch-group-name",
});
```
```python
import pulumi
import pulumi_aws as aws

production = aws.ssm.PatchBaseline("production", approved_patches=["KB123456"])
patchgroup = aws.ssm.PatchGroup("patchgroup",
    baseline_id=production.id,
    patch_group="patch-group-name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var production = new Aws.Ssm.PatchBaseline("production", new()
    {
        ApprovedPatches = new[]
        {
            "KB123456",
        },
    });

    var patchgroup = new Aws.Ssm.PatchGroup("patchgroup", new()
    {
        BaselineId = production.Id,
        PatchGroupName = "patch-group-name",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		production, err := ssm.NewPatchBaseline(ctx, "production", &ssm.PatchBaselineArgs{
			ApprovedPatches: pulumi.StringArray{
				pulumi.String("KB123456"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewPatchGroup(ctx, "patchgroup", &ssm.PatchGroupArgs{
			BaselineId: production.ID(),
			PatchGroup: pulumi.String("patch-group-name"),
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
import com.pulumi.aws.ssm.PatchBaseline;
import com.pulumi.aws.ssm.PatchBaselineArgs;
import com.pulumi.aws.ssm.PatchGroup;
import com.pulumi.aws.ssm.PatchGroupArgs;
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
        var production = new PatchBaseline("production", PatchBaselineArgs.builder()        
            .approvedPatches("KB123456")
            .build());

        var patchgroup = new PatchGroup("patchgroup", PatchGroupArgs.builder()        
            .baselineId(production.id())
            .patchGroup("patch-group-name")
            .build());

    }
}
```
```yaml
resources:
  production:
    type: aws:ssm:PatchBaseline
    properties:
      approvedPatches:
        - KB123456
  patchgroup:
    type: aws:ssm:PatchGroup
    properties:
      baselineId: ${production.id}
      patchGroup: patch-group-name
```
{{% /example %}}
{{% /examples %}}