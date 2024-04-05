Manages an Image Builder Container Recipe.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.ContainerRecipe("example", {
    version: "1.0.0",
    containerType: "DOCKER",
    parentImage: "arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x",
    targetRepository: {
        repositoryName: aws_ecr_repository.example.name,
        service: "ECR",
    },
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
    dockerfileTemplateData: `FROM {{{ imagebuilder:parentImage }}}
{{{ imagebuilder:environments }}}
{{{ imagebuilder:components }}}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.ContainerRecipe("example",
    version="1.0.0",
    container_type="DOCKER",
    parent_image="arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x",
    target_repository=aws.imagebuilder.ContainerRecipeTargetRepositoryArgs(
        repository_name=aws_ecr_repository["example"]["name"],
        service="ECR",
    ),
    components=[aws.imagebuilder.ContainerRecipeComponentArgs(
        component_arn=aws_imagebuilder_component["example"]["arn"],
        parameters=[
            aws.imagebuilder.ContainerRecipeComponentParameterArgs(
                name="Parameter1",
                value="Value1",
            ),
            aws.imagebuilder.ContainerRecipeComponentParameterArgs(
                name="Parameter2",
                value="Value2",
            ),
        ],
    )],
    dockerfile_template_data="""FROM {{{ imagebuilder:parentImage }}}
{{{ imagebuilder:environments }}}
{{{ imagebuilder:components }}}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.ContainerRecipe("example", new()
    {
        Version = "1.0.0",
        ContainerType = "DOCKER",
        ParentImage = "arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x",
        TargetRepository = new Aws.ImageBuilder.Inputs.ContainerRecipeTargetRepositoryArgs
        {
            RepositoryName = aws_ecr_repository.Example.Name,
            Service = "ECR",
        },
        Components = new[]
        {
            new Aws.ImageBuilder.Inputs.ContainerRecipeComponentArgs
            {
                ComponentArn = aws_imagebuilder_component.Example.Arn,
                Parameters = new[]
                {
                    new Aws.ImageBuilder.Inputs.ContainerRecipeComponentParameterArgs
                    {
                        Name = "Parameter1",
                        Value = "Value1",
                    },
                    new Aws.ImageBuilder.Inputs.ContainerRecipeComponentParameterArgs
                    {
                        Name = "Parameter2",
                        Value = "Value2",
                    },
                },
            },
        },
        DockerfileTemplateData = @"FROM {{{ imagebuilder:parentImage }}}
{{{ imagebuilder:environments }}}
{{{ imagebuilder:components }}}
",
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
		_, err := imagebuilder.NewContainerRecipe(ctx, "example", &imagebuilder.ContainerRecipeArgs{
			Version:       pulumi.String("1.0.0"),
			ContainerType: pulumi.String("DOCKER"),
			ParentImage:   pulumi.String("arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x"),
			TargetRepository: &imagebuilder.ContainerRecipeTargetRepositoryArgs{
				RepositoryName: pulumi.Any(aws_ecr_repository.Example.Name),
				Service:        pulumi.String("ECR"),
			},
			Components: imagebuilder.ContainerRecipeComponentArray{
				&imagebuilder.ContainerRecipeComponentArgs{
					ComponentArn: pulumi.Any(aws_imagebuilder_component.Example.Arn),
					Parameters: imagebuilder.ContainerRecipeComponentParameterArray{
						&imagebuilder.ContainerRecipeComponentParameterArgs{
							Name:  pulumi.String("Parameter1"),
							Value: pulumi.String("Value1"),
						},
						&imagebuilder.ContainerRecipeComponentParameterArgs{
							Name:  pulumi.String("Parameter2"),
							Value: pulumi.String("Value2"),
						},
					},
				},
			},
			DockerfileTemplateData: pulumi.String("FROM {{{ imagebuilder:parentImage }}}\n{{{ imagebuilder:environments }}}\n{{{ imagebuilder:components }}}\n"),
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
import com.pulumi.aws.imagebuilder.ContainerRecipe;
import com.pulumi.aws.imagebuilder.ContainerRecipeArgs;
import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeTargetRepositoryArgs;
import com.pulumi.aws.imagebuilder.inputs.ContainerRecipeComponentArgs;
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
        var example = new ContainerRecipe("example", ContainerRecipeArgs.builder()        
            .version("1.0.0")
            .containerType("DOCKER")
            .parentImage("arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x")
            .targetRepository(ContainerRecipeTargetRepositoryArgs.builder()
                .repositoryName(aws_ecr_repository.example().name())
                .service("ECR")
                .build())
            .components(ContainerRecipeComponentArgs.builder()
                .componentArn(aws_imagebuilder_component.example().arn())
                .parameters(                
                    ContainerRecipeComponentParameterArgs.builder()
                        .name("Parameter1")
                        .value("Value1")
                        .build(),
                    ContainerRecipeComponentParameterArgs.builder()
                        .name("Parameter2")
                        .value("Value2")
                        .build())
                .build())
            .dockerfileTemplateData("""
FROM {{{ imagebuilder:parentImage }}}
{{{ imagebuilder:environments }}}
{{{ imagebuilder:components }}}
            """)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:ContainerRecipe
    properties:
      version: 1.0.0
      containerType: DOCKER
      parentImage: arn:aws:imagebuilder:eu-central-1:aws:image/amazon-linux-x86-latest/x.x.x
      targetRepository:
        repositoryName: ${aws_ecr_repository.example.name}
        service: ECR
      components:
        - componentArn: ${aws_imagebuilder_component.example.arn}
          parameters:
            - name: Parameter1
              value: Value1
            - name: Parameter2
              value: Value2
      dockerfileTemplateData: |
        FROM {{{ imagebuilder:parentImage }}}
        {{{ imagebuilder:environments }}}
        {{{ imagebuilder:components }}}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_container_recipe` resources using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/containerRecipe:ContainerRecipe example arn:aws:imagebuilder:us-east-1:123456789012:container-recipe/example/1.0.0
```
 