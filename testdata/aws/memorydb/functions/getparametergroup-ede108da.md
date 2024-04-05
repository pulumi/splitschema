Provides information about a MemoryDB Parameter Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.memorydb.getParameterGroup({
    name: "my-parameter-group",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.get_parameter_group(name="my-parameter-group")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.MemoryDb.GetParameterGroup.Invoke(new()
    {
        Name = "my-parameter-group",
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
		_, err := memorydb.LookupParameterGroup(ctx, &memorydb.LookupParameterGroupArgs{
			Name: "my-parameter-group",
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
import com.pulumi.aws.memorydb.inputs.GetParameterGroupArgs;
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
        final var example = MemorydbFunctions.getParameterGroup(GetParameterGroupArgs.builder()
            .name("my-parameter-group")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:memorydb:getParameterGroup
      Arguments:
        name: my-parameter-group
```
{{% /example %}}
{{% /examples %}}