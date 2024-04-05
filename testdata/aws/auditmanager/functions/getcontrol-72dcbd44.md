Data source for managing an AWS Audit Manager Control.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.auditmanager.getControl({
    name: "1. Risk Management",
    type: "Standard",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.get_control(name="1. Risk Management",
    type="Standard")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Auditmanager.GetControl.Invoke(new()
    {
        Name = "1. Risk Management",
        Type = "Standard",
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
		_, err := auditmanager.LookupControl(ctx, &auditmanager.LookupControlArgs{
			Name: "1. Risk Management",
			Type: "Standard",
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
import com.pulumi.aws.auditmanager.AuditmanagerFunctions;
import com.pulumi.aws.auditmanager.inputs.GetControlArgs;
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
        final var example = AuditmanagerFunctions.getControl(GetControlArgs.builder()
            .name("1. Risk Management")
            .type("Standard")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:auditmanager:getControl
      Arguments:
        name: 1. Risk Management
        type: Standard
```
{{% /example %}}
{{% example %}}
### With Framework Resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleControl = aws.auditmanager.getControl({
    name: "1. Risk Management",
    type: "Standard",
});
const example2 = aws.auditmanager.getControl({
    name: "2. Personnel",
    type: "Standard",
});
const exampleFramework = new aws.auditmanager.Framework("exampleFramework", {controlSets: [
    {
        name: "example",
        controls: [{
            id: exampleControl.then(exampleControl => exampleControl.id),
        }],
    },
    {
        name: "example2",
        controls: [{
            id: example2.then(example2 => example2.id),
        }],
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

example_control = aws.auditmanager.get_control(name="1. Risk Management",
    type="Standard")
example2 = aws.auditmanager.get_control(name="2. Personnel",
    type="Standard")
example_framework = aws.auditmanager.Framework("exampleFramework", control_sets=[
    aws.auditmanager.FrameworkControlSetArgs(
        name="example",
        controls=[aws.auditmanager.FrameworkControlSetControlArgs(
            id=example_control.id,
        )],
    ),
    aws.auditmanager.FrameworkControlSetArgs(
        name="example2",
        controls=[aws.auditmanager.FrameworkControlSetControlArgs(
            id=example2.id,
        )],
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleControl = Aws.Auditmanager.GetControl.Invoke(new()
    {
        Name = "1. Risk Management",
        Type = "Standard",
    });

    var example2 = Aws.Auditmanager.GetControl.Invoke(new()
    {
        Name = "2. Personnel",
        Type = "Standard",
    });

    var exampleFramework = new Aws.Auditmanager.Framework("exampleFramework", new()
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
                        Id = exampleControl.Apply(getControlResult => getControlResult.Id),
                    },
                },
            },
            new Aws.Auditmanager.Inputs.FrameworkControlSetArgs
            {
                Name = "example2",
                Controls = new[]
                {
                    new Aws.Auditmanager.Inputs.FrameworkControlSetControlArgs
                    {
                        Id = example2.Apply(getControlResult => getControlResult.Id),
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
		exampleControl, err := auditmanager.LookupControl(ctx, &auditmanager.LookupControlArgs{
			Name: "1. Risk Management",
			Type: "Standard",
		}, nil)
		if err != nil {
			return err
		}
		example2, err := auditmanager.LookupControl(ctx, &auditmanager.LookupControlArgs{
			Name: "2. Personnel",
			Type: "Standard",
		}, nil)
		if err != nil {
			return err
		}
		_, err = auditmanager.NewFramework(ctx, "exampleFramework", &auditmanager.FrameworkArgs{
			ControlSets: auditmanager.FrameworkControlSetArray{
				&auditmanager.FrameworkControlSetArgs{
					Name: pulumi.String("example"),
					Controls: auditmanager.FrameworkControlSetControlArray{
						&auditmanager.FrameworkControlSetControlArgs{
							Id: *pulumi.String(exampleControl.Id),
						},
					},
				},
				&auditmanager.FrameworkControlSetArgs{
					Name: pulumi.String("example2"),
					Controls: auditmanager.FrameworkControlSetControlArray{
						&auditmanager.FrameworkControlSetControlArgs{
							Id: *pulumi.String(example2.Id),
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
import com.pulumi.aws.auditmanager.AuditmanagerFunctions;
import com.pulumi.aws.auditmanager.inputs.GetControlArgs;
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
        final var exampleControl = AuditmanagerFunctions.getControl(GetControlArgs.builder()
            .name("1. Risk Management")
            .type("Standard")
            .build());

        final var example2 = AuditmanagerFunctions.getControl(GetControlArgs.builder()
            .name("2. Personnel")
            .type("Standard")
            .build());

        var exampleFramework = new Framework("exampleFramework", FrameworkArgs.builder()        
            .controlSets(            
                FrameworkControlSetArgs.builder()
                    .name("example")
                    .controls(FrameworkControlSetControlArgs.builder()
                        .id(exampleControl.applyValue(getControlResult -> getControlResult.id()))
                        .build())
                    .build(),
                FrameworkControlSetArgs.builder()
                    .name("example2")
                    .controls(FrameworkControlSetControlArgs.builder()
                        .id(example2.applyValue(getControlResult -> getControlResult.id()))
                        .build())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  exampleFramework:
    type: aws:auditmanager:Framework
    properties:
      controlSets:
        - name: example
          controls:
            - id: ${exampleControl.id}
        - name: example2
          controls:
            - id: ${example2.id}
variables:
  exampleControl:
    fn::invoke:
      Function: aws:auditmanager:getControl
      Arguments:
        name: 1. Risk Management
        type: Standard
  example2:
    fn::invoke:
      Function: aws:auditmanager:getControl
      Arguments:
        name: 2. Personnel
        type: Standard
```
{{% /example %}}
{{% /examples %}}