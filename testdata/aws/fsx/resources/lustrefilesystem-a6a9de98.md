Manages a FSx Lustre File System. See the [FSx Lustre Guide](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html) for more information.

> **NOTE:** `auto_import_policy`, `export_path`, `import_path` and `imported_file_chunk_size` are not supported with the `PERSISTENT_2` deployment type. Use `aws.fsx.DataRepositoryAssociation` instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.LustreFileSystem("example", {
    importPath: `s3://${aws_s3_bucket.example.bucket}`,
    storageCapacity: 1200,
    subnetIds: [aws_subnet.example.id],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fsx.LustreFileSystem("example",
    import_path=f"s3://{aws_s3_bucket['example']['bucket']}",
    storage_capacity=1200,
    subnet_ids=[aws_subnet["example"]["id"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Fsx.LustreFileSystem("example", new()
    {
        ImportPath = $"s3://{aws_s3_bucket.Example.Bucket}",
        StorageCapacity = 1200,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewLustreFileSystem(ctx, "example", &fsx.LustreFileSystemArgs{
			ImportPath:      pulumi.String(fmt.Sprintf("s3://%v", aws_s3_bucket.Example.Bucket)),
			StorageCapacity: pulumi.Int(1200),
			SubnetIds: pulumi.String{
				aws_subnet.Example.Id,
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
import com.pulumi.aws.fsx.LustreFileSystem;
import com.pulumi.aws.fsx.LustreFileSystemArgs;
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
        var example = new LustreFileSystem("example", LustreFileSystemArgs.builder()        
            .importPath(String.format("s3://%s", aws_s3_bucket.example().bucket()))
            .storageCapacity(1200)
            .subnetIds(aws_subnet.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:fsx:LustreFileSystem
    properties:
      importPath: s3://${aws_s3_bucket.example.bucket}
      storageCapacity: 1200
      subnetIds:
        - ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx File Systems using the `id`. For example:

```sh
 $ pulumi import aws:fsx/lustreFileSystem:LustreFileSystem example fs-543ab12b1ca672f33
```
 Certain resource arguments, like `security_group_ids`, do not have a FSx API method for reading the information after creation. If the argument is set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

