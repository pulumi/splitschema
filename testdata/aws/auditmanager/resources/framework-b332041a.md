Resource for managing an AWS Audit Manager Framework.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.auditmanager.Framework("test", {controlSets: [{
    name: "example",
    controls: [{
        id: aws_auditmanager_control.test.id,
    }],
}]});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.auditmanager.Framework("test", control_sets=[aws.auditmanager.FrameworkControlSetArgs(
    name="example",
    controls=[aws.auditmanager.FrameworkControlSetControlArgs(
        id=aws_auditmanager_control["test"]["id"],
    )],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Auditmanager.Framework("test", new()
    {
        ControlSets = new[]
        {
            new Aws.Auditmanager.Inputs.FrameworkControlSetArgs
            {
                Name = "example",
                Controls = new[]
                {
                    new Aws.Auditmanager.Inputs.FrameworkControlSetControlArgs
                    {
                        Id = aws_auditmanager_control.Test.Id,
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewFramework(ctx, "test", &auditmanager.FrameworkArgs{
			ControlSets: auditmanager.FrameworkControlSetArray{
				&auditmanager.FrameworkControlSetArgs{
					Name: pulumi.String("example"),
					Controls: auditmanager.FrameworkControlSetControlArray{
						&auditmanager.FrameworkControlSetControlArgs{
							Id: pulumi.Any(aws_auditmanager_control.Test.Id),
						},
					},
				},
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
import com.pulumi.aws.auditmanager.Framework;
import com.pulumi.aws.auditmanager.FrameworkArgs;
import com.pulumi.aws.auditmanager.inputs.FrameworkControlSetArgs;
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
        var test = new Framework("test", FrameworkArgs.builder()        
            .controlSets(FrameworkControlSetArgs.builder()
                .name("example")
                .controls(FrameworkControlSetControlArgs.builder()
                    .id(aws_auditmanager_control.test().id())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:auditmanager:Framework
    properties:
      controlSets:
        - name: example
          controls:
            - id: ${aws_auditmanager_control.test.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Framework using the framework `id`. For example:

```sh
 $ pulumi import aws:auditmanager/framework:Framework example abc123-de45
```
 