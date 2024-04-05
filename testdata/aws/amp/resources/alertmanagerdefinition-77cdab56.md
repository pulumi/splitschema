Manages an Amazon Managed Service for Prometheus (AMP) Alert Manager Definition

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const demoWorkspace = new aws.amp.Workspace("demoWorkspace", {});
const demoAlertManagerDefinition = new aws.amp.AlertManagerDefinition("demoAlertManagerDefinition", {
    workspaceId: demoWorkspace.id,
    definition: `alertmanager_config: |
  route:
    receiver: 'default'
  receivers:
    - name: 'default'
`,
});
```
```python
import pulumi
import pulumi_aws as aws

demo_workspace = aws.amp.Workspace("demoWorkspace")
demo_alert_manager_definition = aws.amp.AlertManagerDefinition("demoAlertManagerDefinition",
    workspace_id=demo_workspace.id,
    definition="""alertmanager_config: |
  route:
    receiver: 'default'
  receivers:
    - name: 'default'
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var demoWorkspace = new Aws.Amp.Workspace("demoWorkspace");

    var demoAlertManagerDefinition = new Aws.Amp.AlertManagerDefinition("demoAlertManagerDefinition", new()
    {
        WorkspaceId = demoWorkspace.Id,
        Definition = @"alertmanager_config: |
  route:
    receiver: 'default'
  receivers:
    - name: 'default'
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		demoWorkspace, err := amp.NewWorkspace(ctx, "demoWorkspace", nil)
		if err != nil {
			return err
		}
		_, err = amp.NewAlertManagerDefinition(ctx, "demoAlertManagerDefinition", &amp.AlertManagerDefinitionArgs{
			WorkspaceId: demoWorkspace.ID(),
			Definition: pulumi.String(`alertmanager_config: |
  route:
    receiver: 'default'
  receivers:
    - name: 'default'
`),
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
import com.pulumi.aws.amp.Workspace;
import com.pulumi.aws.amp.AlertManagerDefinition;
import com.pulumi.aws.amp.AlertManagerDefinitionArgs;
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
        var demoWorkspace = new Workspace("demoWorkspace");

        var demoAlertManagerDefinition = new AlertManagerDefinition("demoAlertManagerDefinition", AlertManagerDefinitionArgs.builder()        
            .workspaceId(demoWorkspace.id())
            .definition("""
alertmanager_config: |
  route:
    receiver: 'default'
  receivers:
    - name: 'default'
            """)
            .build());

    }
}
```
```yaml
resources:
  demoWorkspace:
    type: aws:amp:Workspace
  demoAlertManagerDefinition:
    type: aws:amp:AlertManagerDefinition
    properties:
      workspaceId: ${demoWorkspace.id}
      definition: |
        alertmanager_config: |
          route:
            receiver: 'default'
          receivers:
            - name: 'default'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the prometheus alert manager definition using the workspace identifier. For example:

```sh
 $ pulumi import aws:amp/alertManagerDefinition:AlertManagerDefinition demo ws-C6DCB907-F2D7-4D96-957B-66691F865D8B
```
 