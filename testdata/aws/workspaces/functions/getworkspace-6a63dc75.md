Use this data source to get information about a workspace in [AWS Workspaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces.html) Service.

{{% examples %}}
## Example Usage
{{% example %}}
### Filter By Workspace ID

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.workspaces.getWorkspace({
    workspaceId: "ws-cj5xcxsz5",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.workspaces.get_workspace(workspace_id="ws-cj5xcxsz5")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Workspaces.GetWorkspace.Invoke(new()
    {
        WorkspaceId = "ws-cj5xcxsz5",
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
		_, err := workspaces.LookupWorkspace(ctx, &workspaces.LookupWorkspaceArgs{
			WorkspaceId: pulumi.StringRef("ws-cj5xcxsz5"),
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
import com.pulumi.aws.workspaces.inputs.GetWorkspaceArgs;
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
        final var example = WorkspacesFunctions.getWorkspace(GetWorkspaceArgs.builder()
            .workspaceId("ws-cj5xcxsz5")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:workspaces:getWorkspace
      Arguments:
        workspaceId: ws-cj5xcxsz5
```
{{% /example %}}
{{% example %}}
### Filter By Directory ID & User Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.workspaces.getWorkspace({
    directoryId: "d-9967252f57",
    userName: "Example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.workspaces.get_workspace(directory_id="d-9967252f57",
    user_name="Example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Workspaces.GetWorkspace.Invoke(new()
    {
        DirectoryId = "d-9967252f57",
        UserName = "Example",
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
		_, err := workspaces.LookupWorkspace(ctx, &workspaces.LookupWorkspaceArgs{
			DirectoryId: pulumi.StringRef("d-9967252f57"),
			UserName:    pulumi.StringRef("Example"),
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
import com.pulumi.aws.workspaces.inputs.GetWorkspaceArgs;
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
        final var example = WorkspacesFunctions.getWorkspace(GetWorkspaceArgs.builder()
            .directoryId("d-9967252f57")
            .userName("Example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:workspaces:getWorkspace
      Arguments:
        directoryId: d-9967252f57
        userName: Example
```
{{% /example %}}
{{% /examples %}}