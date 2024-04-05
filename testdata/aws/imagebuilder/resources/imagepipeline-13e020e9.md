Manages an Image Builder Image Pipeline.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.ImagePipeline("example", {
    imageRecipeArn: aws_imagebuilder_image_recipe.example.arn,
    infrastructureConfigurationArn: aws_imagebuilder_infrastructure_configuration.example.arn,
    schedule: {
        scheduleExpression: "cron(0 0 * * ? *)",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.ImagePipeline("example",
    image_recipe_arn=aws_imagebuilder_image_recipe["example"]["arn"],
    infrastructure_configuration_arn=aws_imagebuilder_infrastructure_configuration["example"]["arn"],
    schedule=aws.imagebuilder.ImagePipelineScheduleArgs(
        schedule_expression="cron(0 0 * * ? *)",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.ImagePipeline("example", new()
    {
        ImageRecipeArn = aws_imagebuilder_image_recipe.Example.Arn,
        InfrastructureConfigurationArn = aws_imagebuilder_infrastructure_configuration.Example.Arn,
        Schedule = new Aws.ImageBuilder.Inputs.ImagePipelineScheduleArgs
        {
            ScheduleExpression = "cron(0 0 * * ? *)",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := imagebuilder.NewImagePipeline(ctx, "example", &imagebuilder.ImagePipelineArgs{
			ImageRecipeArn:                 pulumi.Any(aws_imagebuilder_image_recipe.Example.Arn),
			InfrastructureConfigurationArn: pulumi.Any(aws_imagebuilder_infrastructure_configuration.Example.Arn),
			Schedule: &imagebuilder.ImagePipelineScheduleArgs{
				ScheduleExpression: pulumi.String("cron(0 0 * * ? *)"),
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
import com.pulumi.aws.imagebuilder.ImagePipeline;
import com.pulumi.aws.imagebuilder.ImagePipelineArgs;
import com.pulumi.aws.imagebuilder.inputs.ImagePipelineScheduleArgs;
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
        var example = new ImagePipeline("example", ImagePipelineArgs.builder()        
            .imageRecipeArn(aws_imagebuilder_image_recipe.example().arn())
            .infrastructureConfigurationArn(aws_imagebuilder_infrastructure_configuration.example().arn())
            .schedule(ImagePipelineScheduleArgs.builder()
                .scheduleExpression("cron(0 0 * * ? *)")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:ImagePipeline
    properties:
      imageRecipeArn: ${aws_imagebuilder_image_recipe.example.arn}
      infrastructureConfigurationArn: ${aws_imagebuilder_infrastructure_configuration.example.arn}
      schedule:
        scheduleExpression: cron(0 0 * * ? *)
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_image_pipeline` resources using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/imagePipeline:ImagePipeline example arn:aws:imagebuilder:us-east-1:123456789012:image-pipeline/example
```
 