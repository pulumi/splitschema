Provides an AppStream image builder.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testFleet = new aws.appstream.ImageBuilder("testFleet", {
    description: "Description of a ImageBuilder",
    displayName: "Display name of a ImageBuilder",
    enableDefaultInternetAccess: false,
    imageName: "AppStream-WinServer2019-10-05-2022",
    instanceType: "stream.standard.large",
    vpcConfig: {
        subnetIds: [aws_subnet.example.id],
    },
    tags: {
        Name: "Example Image Builder",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_fleet = aws.appstream.ImageBuilder("testFleet",
    description="Description of a ImageBuilder",
    display_name="Display name of a ImageBuilder",
    enable_default_internet_access=False,
    image_name="AppStream-WinServer2019-10-05-2022",
    instance_type="stream.standard.large",
    vpc_config=aws.appstream.ImageBuilderVpcConfigArgs(
        subnet_ids=[aws_subnet["example"]["id"]],
    ),
    tags={
        "Name": "Example Image Builder",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testFleet = new Aws.AppStream.ImageBuilder("testFleet", new()
    {
        Description = "Description of a ImageBuilder",
        DisplayName = "Display name of a ImageBuilder",
        EnableDefaultInternetAccess = false,
        ImageName = "AppStream-WinServer2019-10-05-2022",
        InstanceType = "stream.standard.large",
        VpcConfig = new Aws.AppStream.Inputs.ImageBuilderVpcConfigArgs
        {
            SubnetIds = new[]
            {
                aws_subnet.Example.Id,
            },
        },
        Tags = 
        {
            { "Name", "Example Image Builder" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appstream"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appstream.NewImageBuilder(ctx, "testFleet", &appstream.ImageBuilderArgs{
			Description:                 pulumi.String("Description of a ImageBuilder"),
			DisplayName:                 pulumi.String("Display name of a ImageBuilder"),
			EnableDefaultInternetAccess: pulumi.Bool(false),
			ImageName:                   pulumi.String("AppStream-WinServer2019-10-05-2022"),
			InstanceType:                pulumi.String("stream.standard.large"),
			VpcConfig: &appstream.ImageBuilderVpcConfigArgs{
				SubnetIds: pulumi.StringArray{
					aws_subnet.Example.Id,
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Image Builder"),
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
import com.pulumi.aws.appstream.ImageBuilder;
import com.pulumi.aws.appstream.ImageBuilderArgs;
import com.pulumi.aws.appstream.inputs.ImageBuilderVpcConfigArgs;
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
        var testFleet = new ImageBuilder("testFleet", ImageBuilderArgs.builder()        
            .description("Description of a ImageBuilder")
            .displayName("Display name of a ImageBuilder")
            .enableDefaultInternetAccess(false)
            .imageName("AppStream-WinServer2019-10-05-2022")
            .instanceType("stream.standard.large")
            .vpcConfig(ImageBuilderVpcConfigArgs.builder()
                .subnetIds(aws_subnet.example().id())
                .build())
            .tags(Map.of("Name", "Example Image Builder"))
            .build());

    }
}
```
```yaml
resources:
  testFleet:
    type: aws:appstream:ImageBuilder
    properties:
      description: Description of a ImageBuilder
      displayName: Display name of a ImageBuilder
      enableDefaultInternetAccess: false
      imageName: AppStream-WinServer2019-10-05-2022
      instanceType: stream.standard.large
      vpcConfig:
        subnetIds:
          - ${aws_subnet.example.id}
      tags:
        Name: Example Image Builder
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appstream_image_builder` using the `name`. For example:

```sh
 $ pulumi import aws:appstream/imageBuilder:ImageBuilder example imageBuilderExample
```
 