Provides an Amazon Managed Prometheus workspace data source.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.amp.getWorkspace({
    workspaceId: "ws-41det8a1-2c67-6a1a-9381-9b83d3d78ef7",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.amp.get_workspace(workspace_id="ws-41det8a1-2c67-6a1a-9381-9b83d3d78ef7")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Amp.GetWorkspace.Invoke(new()
    {
        WorkspaceId = "ws-41det8a1-2c67-6a1a-9381-9b83d3d78ef7",
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
		_, err := amp.LookupWorkspace(ctx, &amp.LookupWorkspaceArgs{
			WorkspaceId: "ws-41det8a1-2c67-6a1a-9381-9b83d3d78ef7",
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
import com.pulumi.aws.amp.AmpFunctions;
import com.pulumi.aws.amp.inputs.GetWorkspaceArgs;
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
        final var example = AmpFunctions.getWorkspace(GetWorkspaceArgs.builder()
            .workspaceId("ws-41det8a1-2c67-6a1a-9381-9b83d3d78ef7")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:amp:getWorkspace
      Arguments:
        workspaceId: ws-41det8a1-2c67-6a1a-9381-9b83d3d78ef7
```
{{% /example %}}
{{% /examples %}}