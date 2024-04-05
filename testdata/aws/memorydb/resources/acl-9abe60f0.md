Provides a MemoryDB ACL.

More information about users and ACL-s can be found in the [MemoryDB User Guide](https://docs.aws.amazon.com/memorydb/latest/devguide/clusters.acls.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.memorydb.Acl("example", {userNames: [
    "my-user-1",
    "my-user-2",
]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.memorydb.Acl("example", user_names=[
    "my-user-1",
    "my-user-2",
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.MemoryDb.Acl("example", new()
    {
        UserNames = new[]
        {
            "my-user-1",
            "my-user-2",
        },
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
		_, err := memorydb.NewAcl(ctx, "example", &memorydb.AclArgs{
			UserNames: pulumi.StringArray{
				pulumi.String("my-user-1"),
				pulumi.String("my-user-2"),
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
import com.pulumi.aws.memorydb.Acl;
import com.pulumi.aws.memorydb.AclArgs;
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
        var example = new Acl("example", AclArgs.builder()        
            .userNames(            
                "my-user-1",
                "my-user-2")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:memorydb:Acl
    properties:
      userNames:
        - my-user-1
        - my-user-2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an ACL using the `name`. For example:

```sh
 $ pulumi import aws:memorydb/acl:Acl example my-acl
```
 