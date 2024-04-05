Provides details about a specific AWS EC2 Public IPv4 Pool.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getPublicIpv4Pool({
    poolId: "ipv4pool-ec2-000df99cff0c1ec10",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_public_ipv4_pool(pool_id="ipv4pool-ec2-000df99cff0c1ec10")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetPublicIpv4Pool.Invoke(new()
    {
        PoolId = "ipv4pool-ec2-000df99cff0c1ec10",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetPublicIpv4Pool(ctx, &ec2.GetPublicIpv4PoolArgs{
			PoolId: "ipv4pool-ec2-000df99cff0c1ec10",
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetPublicIpv4PoolArgs;
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
        final var example = Ec2Functions.getPublicIpv4Pool(GetPublicIpv4PoolArgs.builder()
            .poolId("ipv4pool-ec2-000df99cff0c1ec10")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getPublicIpv4Pool
      Arguments:
        poolId: ipv4pool-ec2-000df99cff0c1ec10
```
{{% /example %}}
{{% /examples %}}