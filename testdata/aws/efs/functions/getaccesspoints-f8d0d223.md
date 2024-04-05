Provides information about multiple Elastic File System (EFS) Access Points.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.efs.getAccessPoints({
    fileSystemId: "fs-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.efs.get_access_points(file_system_id="fs-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Efs.GetAccessPoints.Invoke(new()
    {
        FileSystemId = "fs-12345678",
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
		_, err := efs.GetAccessPoints(ctx, &efs.GetAccessPointsArgs{
			FileSystemId: "fs-12345678",
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
import com.pulumi.aws.efs.inputs.GetAccessPointsArgs;
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
        final var test = EfsFunctions.getAccessPoints(GetAccessPointsArgs.builder()
            .fileSystemId("fs-12345678")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:efs:getAccessPoints
      Arguments:
        fileSystemId: fs-12345678
```
{{% /example %}}
{{% /examples %}}