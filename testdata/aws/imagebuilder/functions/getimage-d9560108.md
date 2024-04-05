Provides details about an Image Builder Image.

{{% examples %}}
## Example Usage
{{% example %}}
### Latest

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.imagebuilder.getImage({
    arn: "arn:aws:imagebuilder:us-west-2:aws:image/amazon-linux-2-x86/x.x.x",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.get_image(arn="arn:aws:imagebuilder:us-west-2:aws:image/amazon-linux-2-x86/x.x.x")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ImageBuilder.GetImage.Invoke(new()
    {
        Arn = "arn:aws:imagebuilder:us-west-2:aws:image/amazon-linux-2-x86/x.x.x",
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
		_, err := imagebuilder.LookupImage(ctx, &imagebuilder.LookupImageArgs{
			Arn: "arn:aws:imagebuilder:us-west-2:aws:image/amazon-linux-2-x86/x.x.x",
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
import com.pulumi.aws.imagebuilder.inputs.GetImageArgs;
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
        final var example = ImagebuilderFunctions.getImage(GetImageArgs.builder()
            .arn("arn:aws:imagebuilder:us-west-2:aws:image/amazon-linux-2-x86/x.x.x")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:imagebuilder:getImage
      Arguments:
        arn: arn:aws:imagebuilder:us-west-2:aws:image/amazon-linux-2-x86/x.x.x
```
{{% /example %}}
{{% /examples %}}