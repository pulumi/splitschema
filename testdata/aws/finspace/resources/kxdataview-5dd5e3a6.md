Resource for managing an AWS FinSpace Kx Dataview.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.finspace.KxDataview("example", {
    environmentId: aws_finspace_kx_environment.example.id,
    databaseName: aws_finspace_kx_database.example.name,
    availabilityZoneId: "use1-az2",
    description: "Terraform managed Kx Dataview",
    azMode: "SINGLE",
    autoUpdate: true,
    segmentConfigurations: [{
        volumeName: aws_finspace_kx_volume.example.name,
        dbPaths: ["/*"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.finspace.KxDataview("example",
    environment_id=aws_finspace_kx_environment["example"]["id"],
    database_name=aws_finspace_kx_database["example"]["name"],
    availability_zone_id="use1-az2",
    description="Terraform managed Kx Dataview",
    az_mode="SINGLE",
    auto_update=True,
    segment_configurations=[aws.finspace.KxDataviewSegmentConfigurationArgs(
        volume_name=aws_finspace_kx_volume["example"]["name"],
        db_paths=["/*"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.FinSpace.KxDataview("example", new()
    {
        EnvironmentId = aws_finspace_kx_environment.Example.Id,
        DatabaseName = aws_finspace_kx_database.Example.Name,
        AvailabilityZoneId = "use1-az2",
        Description = "Terraform managed Kx Dataview",
        AzMode = "SINGLE",
        AutoUpdate = true,
        SegmentConfigurations = new[]
        {
            new Aws.FinSpace.Inputs.KxDataviewSegmentConfigurationArgs
            {
                VolumeName = aws_finspace_kx_volume.Example.Name,
                DbPaths = new[]
                {
                    "/*",
                },
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
		_, err := finspace.NewKxDataview(ctx, "example", &finspace.KxDataviewArgs{
			EnvironmentId:      pulumi.Any(aws_finspace_kx_environment.Example.Id),
			DatabaseName:       pulumi.Any(aws_finspace_kx_database.Example.Name),
			AvailabilityZoneId: pulumi.String("use1-az2"),
			Description:        pulumi.String("Terraform managed Kx Dataview"),
			AzMode:             pulumi.String("SINGLE"),
			AutoUpdate:         pulumi.Bool(true),
			SegmentConfigurations: finspace.KxDataviewSegmentConfigurationArray{
				&finspace.KxDataviewSegmentConfigurationArgs{
					VolumeName: pulumi.Any(aws_finspace_kx_volume.Example.Name),
					DbPaths: pulumi.StringArray{
						pulumi.String("/*"),
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
import com.pulumi.aws.finspace.KxDataview;
import com.pulumi.aws.finspace.KxDataviewArgs;
import com.pulumi.aws.finspace.inputs.KxDataviewSegmentConfigurationArgs;
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
        var example = new KxDataview("example", KxDataviewArgs.builder()        
            .environmentId(aws_finspace_kx_environment.example().id())
            .databaseName(aws_finspace_kx_database.example().name())
            .availabilityZoneId("use1-az2")
            .description("Terraform managed Kx Dataview")
            .azMode("SINGLE")
            .autoUpdate(true)
            .segmentConfigurations(KxDataviewSegmentConfigurationArgs.builder()
                .volumeName(aws_finspace_kx_volume.example().name())
                .dbPaths("/*")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:finspace:KxDataview
    properties:
      environmentId: ${aws_finspace_kx_environment.example.id}
      databaseName: ${aws_finspace_kx_database.example.name}
      availabilityZoneId: use1-az2
      description: Terraform managed Kx Dataview
      azMode: SINGLE
      autoUpdate: true
      segmentConfigurations:
        - volumeName: ${aws_finspace_kx_volume.example.name}
          dbPaths:
            - /*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx Cluster using the `id` (environment ID and cluster name, comma-delimited). For example:

```sh
 $ pulumi import aws:finspace/kxDataview:KxDataview example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-database,my-tf-kx-dataview
```
 