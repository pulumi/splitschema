Resource for managing an AWS FinSpace Kx Volume.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.finspace.KxVolume("example", {
    environmentId: aws_finspace_kx_environment.example.id,
    availabilityZones: "use1-az2",
    azMode: "SINGLE",
    type: "NAS_1",
    nas1Configurations: [{
        size: 1200,
        type: "SSD_250",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.finspace.KxVolume("example",
    environment_id=aws_finspace_kx_environment["example"]["id"],
    availability_zones="use1-az2",
    az_mode="SINGLE",
    type="NAS_1",
    nas1_configurations=[aws.finspace.KxVolumeNas1ConfigurationArgs(
        size=1200,
        type="SSD_250",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.FinSpace.KxVolume("example", new()
    {
        EnvironmentId = aws_finspace_kx_environment.Example.Id,
        AvailabilityZones = "use1-az2",
        AzMode = "SINGLE",
        Type = "NAS_1",
        Nas1Configurations = new[]
        {
            new Aws.FinSpace.Inputs.KxVolumeNas1ConfigurationArgs
            {
                Size = 1200,
                Type = "SSD_250",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := finspace.NewKxVolume(ctx, "example", &finspace.KxVolumeArgs{
			EnvironmentId:     pulumi.Any(aws_finspace_kx_environment.Example.Id),
			AvailabilityZones: pulumi.StringArray("use1-az2"),
			AzMode:            pulumi.String("SINGLE"),
			Type:              pulumi.String("NAS_1"),
			Nas1Configurations: finspace.KxVolumeNas1ConfigurationArray{
				&finspace.KxVolumeNas1ConfigurationArgs{
					Size: pulumi.Int(1200),
					Type: pulumi.String("SSD_250"),
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
import com.pulumi.aws.finspace.KxVolume;
import com.pulumi.aws.finspace.KxVolumeArgs;
import com.pulumi.aws.finspace.inputs.KxVolumeNas1ConfigurationArgs;
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
        var example = new KxVolume("example", KxVolumeArgs.builder()        
            .environmentId(aws_finspace_kx_environment.example().id())
            .availabilityZones("use1-az2")
            .azMode("SINGLE")
            .type("NAS_1")
            .nas1Configurations(KxVolumeNas1ConfigurationArgs.builder()
                .size(1200)
                .type("SSD_250")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:finspace:KxVolume
    properties:
      environmentId: ${aws_finspace_kx_environment.example.id}
      availabilityZones: use1-az2
      azMode: SINGLE
      type: NAS_1
      nas1Configurations:
        - size: 1200
          type: SSD_250
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx Volume using the `id` (environment ID and volume name, comma-delimited). For example:

```sh
 $ pulumi import aws:finspace/kxVolume:KxVolume example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-volume
```
 