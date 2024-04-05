Use this data source to get the ARNs and names of Image Builder Components matching the specified criteria.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.imagebuilder.getComponents({
    filters: [{
        name: "platform",
        values: ["Linux"],
    }],
    owner: "Self",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.get_components(filters=[aws.imagebuilder.GetComponentsFilterArgs(
        name="platform",
        values=["Linux"],
    )],
    owner="Self")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ImageBuilder.GetComponents.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.ImageBuilder.Inputs.GetComponentsFilterInputArgs
            {
                Name = "platform",
                Values = new[]
                {
                    "Linux",
                },
            },
        },
        Owner = "Self",
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
		_, err := imagebuilder.GetComponents(ctx, &imagebuilder.GetComponentsArgs{
			Filters: []imagebuilder.GetComponentsFilter{
				{
					Name: "platform",
					Values: []string{
						"Linux",
					},
				},
			},
			Owner: pulumi.StringRef("Self"),
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
import com.pulumi.aws.imagebuilder.inputs.GetComponentsArgs;
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
        final var example = ImagebuilderFunctions.getComponents(GetComponentsArgs.builder()
            .filters(GetComponentsFilterArgs.builder()
                .name("platform")
                .values("Linux")
                .build())
            .owner("Self")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:imagebuilder:getComponents
      Arguments:
        filters:
          - name: platform
            values:
              - Linux
        owner: Self
```
{{% /example %}}
{{% /examples %}}