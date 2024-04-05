Provides information about an Elastic File System (EFS) File System.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const fileSystemId = config.get("fileSystemId") || "";
const byId = aws.efs.getFileSystem({
    fileSystemId: fileSystemId,
});
const byTag = aws.efs.getFileSystem({
    tags: {
        Environment: "dev",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
file_system_id = config.get("fileSystemId")
if file_system_id is None:
    file_system_id = ""
by_id = aws.efs.get_file_system(file_system_id=file_system_id)
by_tag = aws.efs.get_file_system(tags={
    "Environment": "dev",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var fileSystemId = config.Get("fileSystemId") ?? "";
    var byId = Aws.Efs.GetFileSystem.Invoke(new()
    {
        FileSystemId = fileSystemId,
    });

    var byTag = Aws.Efs.GetFileSystem.Invoke(new()
    {
        Tags = 
        {
            { "Environment", "dev" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		fileSystemId := ""
		if param := cfg.Get("fileSystemId"); param != "" {
			fileSystemId = param
		}
		_, err := efs.LookupFileSystem(ctx, &efs.LookupFileSystemArgs{
			FileSystemId: pulumi.StringRef(fileSystemId),
		}, nil)
		if err != nil {
			return err
		}
		_, err = efs.LookupFileSystem(ctx, &efs.LookupFileSystemArgs{
			Tags: map[string]interface{}{
				"Environment": "dev",
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
import com.pulumi.aws.efs.EfsFunctions;
import com.pulumi.aws.efs.inputs.GetFileSystemArgs;
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
        final var config = ctx.config();
        final var fileSystemId = config.get("fileSystemId").orElse("");
        final var byId = EfsFunctions.getFileSystem(GetFileSystemArgs.builder()
            .fileSystemId(fileSystemId)
            .build());

        final var byTag = EfsFunctions.getFileSystem(GetFileSystemArgs.builder()
            .tags(Map.of("Environment", "dev"))
            .build());

    }
}
```
```yaml
configuration:
  fileSystemId:
    type: string
    default:
variables:
  byId:
    fn::invoke:
      Function: aws:efs:getFileSystem
      Arguments:
        fileSystemId: ${fileSystemId}
  byTag:
    fn::invoke:
      Function: aws:efs:getFileSystem
      Arguments:
        tags:
          Environment: dev
```
{{% /example %}}
{{% /examples %}}