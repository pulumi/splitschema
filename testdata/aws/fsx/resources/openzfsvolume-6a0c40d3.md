Manages an Amazon FSx for OpenZFS volume.
See the [FSx OpenZFS User Guide](https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/what-is-fsx.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OpenZfsVolume("test", {parentVolumeId: aws_fsx_openzfs_file_system.test.root_volume_id});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OpenZfsVolume("test", parent_volume_id=aws_fsx_openzfs_file_system["test"]["root_volume_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OpenZfsVolume("test", new()
    {
        ParentVolumeId = aws_fsx_openzfs_file_system.Test.Root_volume_id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewOpenZfsVolume(ctx, "test", &fsx.OpenZfsVolumeArgs{
			ParentVolumeId: pulumi.Any(aws_fsx_openzfs_file_system.Test.Root_volume_id),
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
import com.pulumi.aws.fsx.OpenZfsVolume;
import com.pulumi.aws.fsx.OpenZfsVolumeArgs;
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
        var test = new OpenZfsVolume("test", OpenZfsVolumeArgs.builder()        
            .parentVolumeId(aws_fsx_openzfs_file_system.test().root_volume_id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OpenZfsVolume
    properties:
      parentVolumeId: ${aws_fsx_openzfs_file_system.test.root_volume_id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx Volumes using the `id`. For example:

```sh
 $ pulumi import aws:fsx/openZfsVolume:OpenZfsVolume example fsvol-543ab12b1ca672f33
```
 