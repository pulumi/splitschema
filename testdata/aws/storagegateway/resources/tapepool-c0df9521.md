Manages an AWS Storage Gateway Tape Pool.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.TapePool("example", {
    poolName: "example",
    storageClass: "GLACIER",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.TapePool("example",
    pool_name="example",
    storage_class="GLACIER")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.TapePool("example", new()
    {
        PoolName = "example",
        StorageClass = "GLACIER",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewTapePool(ctx, "example", &storagegateway.TapePoolArgs{
			PoolName:     pulumi.String("example"),
			StorageClass: pulumi.String("GLACIER"),
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
import com.pulumi.aws.storagegateway.TapePool;
import com.pulumi.aws.storagegateway.TapePoolArgs;
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
        var example = new TapePool("example", TapePoolArgs.builder()        
            .poolName("example")
            .storageClass("GLACIER")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:TapePool
    properties:
      poolName: example
      storageClass: GLACIER
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_tape_pool` using the volume Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/tapePool:TapePool example arn:aws:storagegateway:us-east-1:123456789012:tapepool/pool-12345678
```
 