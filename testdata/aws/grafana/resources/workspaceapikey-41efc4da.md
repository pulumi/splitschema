Provides an Amazon Managed Grafana workspace API Key resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const key = new aws.grafana.WorkspaceApiKey("key", {
    keyName: "test-key",
    keyRole: "VIEWER",
    secondsToLive: 3600,
    workspaceId: aws_grafana_workspace.test.id,
});
```
```python
import pulumi
import pulumi_aws as aws

key = aws.grafana.WorkspaceApiKey("key",
    key_name="test-key",
    key_role="VIEWER",
    seconds_to_live=3600,
    workspace_id=aws_grafana_workspace["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var key = new Aws.Grafana.WorkspaceApiKey("key", new()
    {
        KeyName = "test-key",
        KeyRole = "VIEWER",
        SecondsToLive = 3600,
        WorkspaceId = aws_grafana_workspace.Test.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/grafana"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := grafana.NewWorkspaceApiKey(ctx, "key", &grafana.WorkspaceApiKeyArgs{
			KeyName:       pulumi.String("test-key"),
			KeyRole:       pulumi.String("VIEWER"),
			SecondsToLive: pulumi.Int(3600),
			WorkspaceId:   pulumi.Any(aws_grafana_workspace.Test.Id),
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
import com.pulumi.aws.grafana.WorkspaceApiKey;
import com.pulumi.aws.grafana.WorkspaceApiKeyArgs;
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
        var key = new WorkspaceApiKey("key", WorkspaceApiKeyArgs.builder()        
            .keyName("test-key")
            .keyRole("VIEWER")
            .secondsToLive(3600)
            .workspaceId(aws_grafana_workspace.test().id())
            .build());

    }
}
```
```yaml
resources:
  key:
    type: aws:grafana:WorkspaceApiKey
    properties:
      keyName: test-key
      keyRole: VIEWER
      secondsToLive: 3600
      workspaceId: ${aws_grafana_workspace.test.id}
```
{{% /example %}}
{{% /examples %}}