Manages an Image Builder Image.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.Image("example", {
    distributionConfigurationArn: aws_imagebuilder_distribution_configuration.example.arn,
    imageRecipeArn: aws_imagebuilder_image_recipe.example.arn,
    infrastructureConfigurationArn: aws_imagebuilder_infrastructure_configuration.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.Image("example",
    distribution_configuration_arn=aws_imagebuilder_distribution_configuration["example"]["arn"],
    image_recipe_arn=aws_imagebuilder_image_recipe["example"]["arn"],
    infrastructure_configuration_arn=aws_imagebuilder_infrastructure_configuration["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.Image("example", new()
    {
        DistributionConfigurationArn = aws_imagebuilder_distribution_configuration.Example.Arn,
        ImageRecipeArn = aws_imagebuilder_image_recipe.Example.Arn,
        InfrastructureConfigurationArn = aws_imagebuilder_infrastructure_configuration.Example.Arn,
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
		_, err := imagebuilder.NewImage(ctx, "example", &imagebuilder.ImageArgs{
			DistributionConfigurationArn:   pulumi.Any(aws_imagebuilder_distribution_configuration.Example.Arn),
			ImageRecipeArn:                 pulumi.Any(aws_imagebuilder_image_recipe.Example.Arn),
			InfrastructureConfigurationArn: pulumi.Any(aws_imagebuilder_infrastructure_configuration.Example.Arn),
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
import com.pulumi.aws.imagebuilder.Image;
import com.pulumi.aws.imagebuilder.ImageArgs;
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
        var example = new Image("example", ImageArgs.builder()        
            .distributionConfigurationArn(aws_imagebuilder_distribution_configuration.example().arn())
            .imageRecipeArn(aws_imagebuilder_image_recipe.example().arn())
            .infrastructureConfigurationArn(aws_imagebuilder_infrastructure_configuration.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:Image
    properties:
      distributionConfigurationArn: ${aws_imagebuilder_distribution_configuration.example.arn}
      imageRecipeArn: ${aws_imagebuilder_image_recipe.example.arn}
      infrastructureConfigurationArn: ${aws_imagebuilder_infrastructure_configuration.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_image` resources using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/image:Image example arn:aws:imagebuilder:us-east-1:123456789012:image/example/1.0.0/1
```
 