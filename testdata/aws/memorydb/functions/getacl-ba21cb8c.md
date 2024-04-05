Provides information about a MemoryDB ACL.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.memorydb.getAcl({
    name: "my-acl",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.get_acl(name="my-acl")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.MemoryDb.GetAcl.Invoke(new()
    {
        Name = "my-acl",
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
		_, err := memorydb.LookupAcl(ctx, &memorydb.LookupAclArgs{
			Name: "my-acl",
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
import com.pulumi.aws.memorydb.inputs.GetAclArgs;
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
        final var example = MemorydbFunctions.getAcl(GetAclArgs.builder()
            .name("my-acl")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:memorydb:getAcl
      Arguments:
        name: my-acl
```
{{% /example %}}
{{% /examples %}}