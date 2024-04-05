Manages an Image Builder Image Recipe.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.ImageRecipe("example", {
    blockDeviceMappings: [{
        deviceName: "/dev/xvdb",
        ebs: {
            deleteOnTermination: "true",
            volumeSize: 100,
            volumeType: "gp2",
        },
    }],
    components: [{
        componentArn: aws_imagebuilder_component.example.arn,
        parameters: [
            {
                name: "Parameter1",
                value: "Value1",
            },
            {
                name: "Parameter2",
                value: "Value2",
            },
        ],
    }],
    parentImage: `arn:${data.aws_partition.current.partition}:imagebuilder:${data.aws_region.current.name}:aws:image/amazon-linux-2-x86/x.x.x`,
    version: "1.0.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.ImageRecipe("example",
    block_device_mappings=[aws.imagebuilder.ImageRecipeBlockDeviceMappingArgs(
        device_name="/dev/xvdb",
        ebs=aws.imagebuilder.ImageRecipeBlockDeviceMappingEbsArgs(
            delete_on_termination="true",
            volume_size=100,
            volume_type="gp2",
        ),
    )],
    components=[aws.imagebuilder.ImageRecipeComponentArgs(
        component_arn=aws_imagebuilder_component["example"]["arn"],
        parameters=[
            aws.imagebuilder.ImageRecipeComponentParameterArgs(
                name="Parameter1",
                value="Value1",
            ),
            aws.imagebuilder.ImageRecipeComponentParameterArgs(
                name="Parameter2",
                value="Value2",
            ),
        ],
    )],
    parent_image=f"arn:{data['aws_partition']['current']['partition']}:imagebuilder:{data['aws_region']['current']['name']}:aws:image/amazon-linux-2-x86/x.x.x",
    version="1.0.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.ImageRecipe("example", new()
    {
        BlockDeviceMappings = new[]
        {
            new Aws.ImageBuilder.Inputs.ImageRecipeBlockDeviceMappingArgs
            {
                DeviceName = "/dev/xvdb",
                Ebs = new Aws.ImageBuilder.Inputs.ImageRecipeBlockDeviceMappingEbsArgs
                {
                    DeleteOnTermination = "true",
                    VolumeSize = 100,
                    VolumeType = "gp2",
                },
            },
        },
        Components = new[]
        {
            new Aws.ImageBuilder.Inputs.ImageRecipeComponentArgs
            {
                ComponentArn = aws_imagebuilder_component.Example.Arn,
                Parameters = new[]
                {
                    new Aws.ImageBuilder.Inputs.ImageRecipeComponentParameterArgs
                    {
                        Name = "Parameter1",
                        Value = "Value1",
                    },
                    new Aws.ImageBuilder.Inputs.ImageRecipeComponentParameterArgs
                    {
                        Name = "Parameter2",
                        Value = "Value2",
                    },
                },
            },
        },
        ParentImage = $"arn:{data.Aws_partition.Current.Partition}:imagebuilder:{data.Aws_region.Current.Name}:aws:image/amazon-linux-2-x86/x.x.x",
        Version = "1.0.0",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := imagebuilder.NewImageRecipe(ctx, "example", &imagebuilder.ImageRecipeArgs{
			BlockDeviceMappings: imagebuilder.ImageRecipeBlockDeviceMappingArray{
				&imagebuilder.ImageRecipeBlockDeviceMappingArgs{
					DeviceName: pulumi.String("/dev/xvdb"),
					Ebs: &imagebuilder.ImageRecipeBlockDeviceMappingEbsArgs{
						DeleteOnTermination: pulumi.String("true"),
						VolumeSize:          pulumi.Int(100),
						VolumeType:          pulumi.String("gp2"),
					},
				},
			},
			Components: imagebuilder.ImageRecipeComponentArray{
				&imagebuilder.ImageRecipeComponentArgs{
					ComponentArn: pulumi.Any(aws_imagebuilder_component.Example.Arn),
					Parameters: imagebuilder.ImageRecipeComponentParameterArray{
						&imagebuilder.ImageRecipeComponentParameterArgs{
							Name:  pulumi.String("Parameter1"),
							Value: pulumi.String("Value1"),
						},
						&imagebuilder.ImageRecipeComponentParameterArgs{
							Name:  pulumi.String("Parameter2"),
							Value: pulumi.String("Value2"),
						},
					},
				},
			},
			ParentImage: pulumi.String(fmt.Sprintf("arn:%v:imagebuilder:%v:aws:image/amazon-linux-2-x86/x.x.x", data.Aws_partition.Current.Partition, data.Aws_region.Current.Name)),
			Version:     pulumi.String("1.0.0"),
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
import com.pulumi.aws.imagebuilder.ImageRecipe;
import com.pulumi.aws.imagebuilder.ImageRecipeArgs;
import com.pulumi.aws.imagebuilder.inputs.ImageRecipeBlockDeviceMappingArgs;
import com.pulumi.aws.imagebuilder.inputs.ImageRecipeBlockDeviceMappingEbsArgs;
import com.pulumi.aws.imagebuilder.inputs.ImageRecipeComponentArgs;
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
        var example = new ImageRecipe("example", ImageRecipeArgs.builder()        
            .blockDeviceMappings(ImageRecipeBlockDeviceMappingArgs.builder()
                .deviceName("/dev/xvdb")
                .ebs(ImageRecipeBlockDeviceMappingEbsArgs.builder()
                    .deleteOnTermination(true)
                    .volumeSize(100)
                    .volumeType("gp2")
                    .build())
                .build())
            .components(ImageRecipeComponentArgs.builder()
                .componentArn(aws_imagebuilder_component.example().arn())
                .parameters(                
                    ImageRecipeComponentParameterArgs.builder()
                        .name("Parameter1")
                        .value("Value1")
                        .build(),
                    ImageRecipeComponentParameterArgs.builder()
                        .name("Parameter2")
                        .value("Value2")
                        .build())
                .build())
            .parentImage(String.format("arn:%s:imagebuilder:%s:aws:image/amazon-linux-2-x86/x.x.x", data.aws_partition().current().partition(),data.aws_region().current().name()))
            .version("1.0.0")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:ImageRecipe
    properties:
      blockDeviceMappings:
        - deviceName: /dev/xvdb
          ebs:
            deleteOnTermination: true
            volumeSize: 100
            volumeType: gp2
      components:
        - componentArn: ${aws_imagebuilder_component.example.arn}
          parameters:
            - name: Parameter1
              value: Value1
            - name: Parameter2
              value: Value2
      parentImage: arn:${data.aws_partition.current.partition}:imagebuilder:${data.aws_region.current.name}:aws:image/amazon-linux-2-x86/x.x.x
      version: 1.0.0
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_image_recipe` resources using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/imageRecipe:ImageRecipe example arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/example/1.0.0
```
 