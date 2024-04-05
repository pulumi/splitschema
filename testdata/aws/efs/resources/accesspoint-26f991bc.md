Provides an Elastic File System (EFS) access point.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.efs.AccessPoint("test", {fileSystemId: aws_efs_file_system.foo.id});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.efs.AccessPoint("test", file_system_id=aws_efs_file_system["foo"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Efs.AccessPoint("test", new()
    {
        FileSystemId = aws_efs_file_system.Foo.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.NewAccessPoint(ctx, "test", &efs.AccessPointArgs{
			FileSystemId: pulumi.Any(aws_efs_file_system.Foo.Id),
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
import com.pulumi.aws.efs.AccessPoint;
import com.pulumi.aws.efs.AccessPointArgs;
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
        var test = new AccessPoint("test", AccessPointArgs.builder()        
            .fileSystemId(aws_efs_file_system.foo().id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:efs:AccessPoint
    properties:
      fileSystemId: ${aws_efs_file_system.foo.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the EFS access points using the `id`. For example:

```sh
 $ pulumi import aws:efs/accessPoint:AccessPoint test fsap-52a643fb
```
 