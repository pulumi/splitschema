Provides information about a MemoryDB Snapshot.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.memorydb.getSnapshot({
    name: "my-snapshot",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.get_snapshot(name="my-snapshot")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.MemoryDb.GetSnapshot.Invoke(new()
    {
        Name = "my-snapshot",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/memorydb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := memorydb.LookupSnapshot(ctx, &memorydb.LookupSnapshotArgs{
			Name: "my-snapshot",
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
import com.pulumi.aws.memorydb.MemorydbFunctions;
import com.pulumi.aws.memorydb.inputs.GetSnapshotArgs;
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
        final var example = MemorydbFunctions.getSnapshot(GetSnapshotArgs.builder()
            .name("my-snapshot")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:memorydb:getSnapshot
      Arguments:
        name: my-snapshot
```
{{% /example %}}
{{% /examples %}}