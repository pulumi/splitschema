Use this data source to get the ARNs and names of Image Builder Infrastructure Configurations matching the specified criteria.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.imagebuilder.getInfrastructureConfigurations({
    filters: [{
        name: "name",
        values: ["example"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.get_infrastructure_configurations(filters=[aws.imagebuilder.GetInfrastructureConfigurationsFilterArgs(
    name="name",
    values=["example"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ImageBuilder.GetInfrastructureConfigurations.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.ImageBuilder.Inputs.GetInfrastructureConfigurationsFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "example",
                },
            },
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
		_, err := imagebuilder.GetInfrastructureConfigurations(ctx, &imagebuilder.GetInfrastructureConfigurationsArgs{
			Filters: []imagebuilder.GetInfrastructureConfigurationsFilter{
				{
					Name: "name",
					Values: []string{
						"example",
					},
				},
			},
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
import com.pulumi.aws.imagebuilder.inputs.GetInfrastructureConfigurationsArgs;
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
        final var example = ImagebuilderFunctions.getInfrastructureConfigurations(GetInfrastructureConfigurationsArgs.builder()
            .filters(GetInfrastructureConfigurationsFilterArgs.builder()
                .name("name")
                .values("example")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:imagebuilder:getInfrastructureConfigurations
      Arguments:
        filters:
          - name: name
            values:
              - example
```
{{% /example %}}
{{% /examples %}}