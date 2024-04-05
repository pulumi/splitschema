Use this data source to get information about a Workspaces image.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.workspaces.getImage({
    imageId: "wsi-ten5h0y19",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.workspaces.get_image(image_id="wsi-ten5h0y19")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Workspaces.GetImage.Invoke(new()
    {
        ImageId = "wsi-ten5h0y19",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workspaces.GetImage(ctx, &workspaces.GetImageArgs{
			ImageId: "wsi-ten5h0y19",
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
import com.pulumi.aws.workspaces.WorkspacesFunctions;
import com.pulumi.aws.workspaces.inputs.GetImageArgs;
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
        final var example = WorkspacesFunctions.getImage(GetImageArgs.builder()
            .imageId("wsi-ten5h0y19")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:workspaces:getImage
      Arguments:
        imageId: wsi-ten5h0y19
```
{{% /example %}}
{{% /examples %}}