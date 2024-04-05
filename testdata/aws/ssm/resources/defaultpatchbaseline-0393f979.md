Resource for registering an AWS Systems Manager Default Patch Baseline.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePatchBaseline = new aws.ssm.PatchBaseline("examplePatchBaseline", {approvedPatches: ["KB123456"]});
const exampleDefaultPatchBaseline = new aws.ssm.DefaultPatchBaseline("exampleDefaultPatchBaseline", {
    baselineId: examplePatchBaseline.id,
    operatingSystem: examplePatchBaseline.operatingSystem,
});
```
```python
import pulumi
import pulumi_aws as aws

example_patch_baseline = aws.ssm.PatchBaseline("examplePatchBaseline", approved_patches=["KB123456"])
example_default_patch_baseline = aws.ssm.DefaultPatchBaseline("exampleDefaultPatchBaseline",
    baseline_id=example_patch_baseline.id,
    operating_system=example_patch_baseline.operating_system)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePatchBaseline = new Aws.Ssm.PatchBaseline("examplePatchBaseline", new()
    {
        ApprovedPatches = new[]
        {
            "KB123456",
        },
    });

    var exampleDefaultPatchBaseline = new Aws.Ssm.DefaultPatchBaseline("exampleDefaultPatchBaseline", new()
    {
        BaselineId = examplePatchBaseline.Id,
        OperatingSystem = examplePatchBaseline.OperatingSystem,
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
		examplePatchBaseline, err := ssm.NewPatchBaseline(ctx, "examplePatchBaseline", &ssm.PatchBaselineArgs{
			ApprovedPatches: pulumi.StringArray{
				pulumi.String("KB123456"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewDefaultPatchBaseline(ctx, "exampleDefaultPatchBaseline", &ssm.DefaultPatchBaselineArgs{
			BaselineId:      examplePatchBaseline.ID(),
			OperatingSystem: examplePatchBaseline.OperatingSystem,
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
import com.pulumi.aws.ssm.DefaultPatchBaseline;
import com.pulumi.aws.ssm.DefaultPatchBaselineArgs;
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
        var examplePatchBaseline = new PatchBaseline("examplePatchBaseline", PatchBaselineArgs.builder()        
            .approvedPatches("KB123456")
            .build());

        var exampleDefaultPatchBaseline = new DefaultPatchBaseline("exampleDefaultPatchBaseline", DefaultPatchBaselineArgs.builder()        
            .baselineId(examplePatchBaseline.id())
            .operatingSystem(examplePatchBaseline.operatingSystem())
            .build());

    }
}
```
```yaml
resources:
  exampleDefaultPatchBaseline:
    type: aws:ssm:DefaultPatchBaseline
    properties:
      baselineId: ${examplePatchBaseline.id}
      operatingSystem: ${examplePatchBaseline.operatingSystem}
  examplePatchBaseline:
    type: aws:ssm:PatchBaseline
    properties:
      approvedPatches:
        - KB123456
```
{{% /example %}}
{{% /examples %}}

## Import

Using the patch baseline ARN:

Using the operating system value:

__Using `pulumi import` to import__ the Systems Manager Default Patch Baseline using the patch baseline ID, patch baseline ARN, or the operating system value. For example:

Using the patch baseline ID:

```sh
 $ pulumi import aws:ssm/defaultPatchBaseline:DefaultPatchBaseline example pb-1234567890abcdef1
```
 Using the patch baseline ARN:

```sh
 $ pulumi import aws:ssm/defaultPatchBaseline:DefaultPatchBaseline example arn:aws:ssm:us-west-2:123456789012:patchbaseline/pb-1234567890abcdef1
```
 Using the operating system value:

```sh
 $ pulumi import aws:ssm/defaultPatchBaseline:DefaultPatchBaseline example CENTOS
```
 