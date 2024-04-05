Resource for managing an AWS SESv2 (Simple Email V2) Dedicated IP Pool.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.DedicatedIpPool("example", {poolName: "my-pool"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.DedicatedIpPool("example", pool_name="my-pool")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.DedicatedIpPool("example", new()
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
		_, err := sesv2.NewDedicatedIpPool(ctx, "example", &sesv2.DedicatedIpPoolArgs{
			PoolName: pulumi.String("my-pool"),
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
import com.pulumi.aws.sesv2.DedicatedIpPool;
import com.pulumi.aws.sesv2.DedicatedIpPoolArgs;
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
        var example = new DedicatedIpPool("example", DedicatedIpPoolArgs.builder()        
            .poolName("my-pool")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:DedicatedIpPool
    properties:
      poolName: my-pool
```
{{% /example %}}
{{% example %}}
### Managed Pool

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.DedicatedIpPool("example", {
    poolName: "my-managed-pool",
    scalingMode: "MANAGED",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.DedicatedIpPool("example",
    pool_name="my-managed-pool",
    scaling_mode="MANAGED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.DedicatedIpPool("example", new()
    {
        PoolName = "my-managed-pool",
        ScalingMode = "MANAGED",
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
		_, err := sesv2.NewDedicatedIpPool(ctx, "example", &sesv2.DedicatedIpPoolArgs{
			PoolName:    pulumi.String("my-managed-pool"),
			ScalingMode: pulumi.String("MANAGED"),
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
import com.pulumi.aws.sesv2.DedicatedIpPool;
import com.pulumi.aws.sesv2.DedicatedIpPoolArgs;
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
        var example = new DedicatedIpPool("example", DedicatedIpPoolArgs.builder()        
            .poolName("my-managed-pool")
            .scalingMode("MANAGED")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:DedicatedIpPool
    properties:
      poolName: my-managed-pool
      scalingMode: MANAGED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Dedicated IP Pool using the `pool_name`. For example:

```sh
 $ pulumi import aws:sesv2/dedicatedIpPool:DedicatedIpPool example my-pool
```
 