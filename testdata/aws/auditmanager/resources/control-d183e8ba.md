Resource for managing an AWS Audit Manager Control.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.auditmanager.Control("example", {controlMappingSources: [{
    sourceName: "example",
    sourceSetUpOption: "Procedural_Controls_Mapping",
    sourceType: "MANUAL",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.Control("example", control_mapping_sources=[aws.auditmanager.ControlControlMappingSourceArgs(
    source_name="example",
    source_set_up_option="Procedural_Controls_Mapping",
    source_type="MANUAL",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Auditmanager.Control("example", new()
    {
        ControlMappingSources = new[]
        {
            new Aws.Auditmanager.Inputs.ControlControlMappingSourceArgs
            {
                SourceName = "example",
                SourceSetUpOption = "Procedural_Controls_Mapping",
                SourceType = "MANUAL",
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
		_, err := auditmanager.NewControl(ctx, "example", &auditmanager.ControlArgs{
			ControlMappingSources: auditmanager.ControlControlMappingSourceArray{
				&auditmanager.ControlControlMappingSourceArgs{
					SourceName:        pulumi.String("example"),
					SourceSetUpOption: pulumi.String("Procedural_Controls_Mapping"),
					SourceType:        pulumi.String("MANUAL"),
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
import com.pulumi.aws.auditmanager.Control;
import com.pulumi.aws.auditmanager.ControlArgs;
import com.pulumi.aws.auditmanager.inputs.ControlControlMappingSourceArgs;
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
        var example = new Control("example", ControlArgs.builder()        
            .controlMappingSources(ControlControlMappingSourceArgs.builder()
                .sourceName("example")
                .sourceSetUpOption("Procedural_Controls_Mapping")
                .sourceType("MANUAL")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:auditmanager:Control
    properties:
      controlMappingSources:
        - sourceName: example
          sourceSetUpOption: Procedural_Controls_Mapping
          sourceType: MANUAL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an Audit Manager Control using the `id`. For example:

```sh
 $ pulumi import aws:auditmanager/control:Control example abc123-de45
```
 