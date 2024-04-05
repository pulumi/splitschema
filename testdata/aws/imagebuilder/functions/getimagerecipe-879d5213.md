Provides details about an Image Builder Image Recipe.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.imagebuilder.getImageRecipe({
    arn: "arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.get_image_recipe(arn="arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ImageBuilder.GetImageRecipe.Invoke(new()
    {
        Arn = "arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0",
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
		_, err := imagebuilder.LookupImageRecipe(ctx, &imagebuilder.LookupImageRecipeArgs{
			Arn: "arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0",
		}, nil)
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
import com.pulumi.aws.imagebuilder.ImagebuilderFunctions;
import com.pulumi.aws.imagebuilder.inputs.GetImageRecipeArgs;
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
        final var example = ImagebuilderFunctions.getImageRecipe(GetImageRecipeArgs.builder()
            .arn("arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:imagebuilder:getImageRecipe
      Arguments:
        arn: arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0
```
{{% /example %}}
{{% /examples %}}