Data source for managing an AWS SESv2 (Simple Email V2) Dedicated IP Pool.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.sesv2.getDedicatedIpPool({
    poolName: "my-pool",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.get_dedicated_ip_pool(pool_name="my-pool")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SesV2.GetDedicatedIpPool.Invoke(new()
    {
        PoolName = "my-pool",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.LookupDedicatedIpPool(ctx, &sesv2.LookupDedicatedIpPoolArgs{
			PoolName: "my-pool",
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
import com.pulumi.aws.sesv2.Sesv2Functions;
import com.pulumi.aws.sesv2.inputs.GetDedicatedIpPoolArgs;
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
        final var example = Sesv2Functions.getDedicatedIpPool(GetDedicatedIpPoolArgs.builder()
            .poolName("my-pool")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:sesv2:getDedicatedIpPool
      Arguments:
        poolName: my-pool
```
{{% /example %}}
{{% /examples %}}