Manages a Service Catalog self-service action.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.ServiceAction("example", {
    definition: {
        name: "AWS-RestartEC2Instance",
    },
    description: "Motor generator unit",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.ServiceAction("example",
    definition=aws.servicecatalog.ServiceActionDefinitionArgs(
        name="AWS-RestartEC2Instance",
    ),
    description="Motor generator unit")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.ServiceAction("example", new()
    {
        Definition = new Aws.ServiceCatalog.Inputs.ServiceActionDefinitionArgs
        {
            Name = "AWS-RestartEC2Instance",
        },
        Description = "Motor generator unit",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicecatalog.NewServiceAction(ctx, "example", &servicecatalog.ServiceActionArgs{
			Definition: &servicecatalog.ServiceActionDefinitionArgs{
				Name: pulumi.String("AWS-RestartEC2Instance"),
			},
			Description: pulumi.String("Motor generator unit"),
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
import com.pulumi.aws.servicecatalog.ServiceAction;
import com.pulumi.aws.servicecatalog.ServiceActionArgs;
import com.pulumi.aws.servicecatalog.inputs.ServiceActionDefinitionArgs;
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
        var example = new ServiceAction("example", ServiceActionArgs.builder()        
            .definition(ServiceActionDefinitionArgs.builder()
                .name("AWS-RestartEC2Instance")
                .build())
            .description("Motor generator unit")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:ServiceAction
    properties:
      definition:
        name: AWS-RestartEC2Instance
      description: Motor generator unit
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_service_action` using the service action ID. For example:

```sh
 $ pulumi import aws:servicecatalog/serviceAction:ServiceAction example act-f1w12eperfslh
```
 