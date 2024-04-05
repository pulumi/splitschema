Retrieve information about an AWS WorkSpaces directory.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.workspaces.getDirectory({
    directoryId: "d-9067783251",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.workspaces.get_directory(directory_id="d-9067783251")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Workspaces.GetDirectory.Invoke(new()
    {
        DirectoryId = "d-9067783251",
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
		_, err := workspaces.LookupDirectory(ctx, &workspaces.LookupDirectoryArgs{
			DirectoryId: "d-9067783251",
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
import com.pulumi.aws.workspaces.inputs.GetDirectoryArgs;
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
        final var example = WorkspacesFunctions.getDirectory(GetDirectoryArgs.builder()
            .directoryId("d-9067783251")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:workspaces:getDirectory
      Arguments:
        directoryId: d-9067783251
```
{{% /example %}}
{{% /examples %}}