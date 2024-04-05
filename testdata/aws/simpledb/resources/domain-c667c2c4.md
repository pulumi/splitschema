Provides a SimpleDB domain resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const users = new aws.simpledb.Domain("users", {});
```
```python
import pulumi
import pulumi_aws as aws

users = aws.simpledb.Domain("users")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var users = new Aws.SimpleDB.Domain("users");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/simpledb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := simpledb.NewDomain(ctx, "users", nil)
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
import com.pulumi.aws.simpledb.Domain;
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
        var users = new Domain("users");

    }
}
```
```yaml
resources:
  users:
    type: aws:simpledb:Domain
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SimpleDB Domains using the `name`. For example:

```sh
 $ pulumi import aws:simpledb/domain:Domain users users
```
 