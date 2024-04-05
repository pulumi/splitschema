Provides details about an Image builder Container Recipe.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.imagebuilder.getContainerRecipe({
    arn: "arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.get_container_recipe(arn="arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ImageBuilder.GetContainerRecipe.Invoke(new()
    {
        Arn = "arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0",
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
		_, err := imagebuilder.LookupContainerRecipe(ctx, &imagebuilder.LookupContainerRecipeArgs{
			Arn: "arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0",
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
import com.pulumi.aws.imagebuilder.inputs.GetContainerRecipeArgs;
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
        final var example = ImagebuilderFunctions.getContainerRecipe(GetContainerRecipeArgs.builder()
            .arn("arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:imagebuilder:getContainerRecipe
      Arguments:
        arn: arn:aws:imagebuilder:us-east-1:aws:container-recipe/example/1.0.0
```
{{% /example %}}
{{% /examples %}}