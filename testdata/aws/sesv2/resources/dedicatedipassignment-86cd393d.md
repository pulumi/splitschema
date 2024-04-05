Resource for managing an AWS SESv2 (Simple Email V2) Dedicated IP Assignment.

This resource is used with "Standard" dedicated IP addresses. This includes addresses [requested and relinquished manually](https://docs.aws.amazon.com/ses/latest/dg/dedicated-ip-case.html) via an AWS support case, or [Bring Your Own IP](https://docs.aws.amazon.com/ses/latest/dg/dedicated-ip-byo.html) addresses. Once no longer assigned, this resource returns the IP to the [`ses-default-dedicated-pool`](https://docs.aws.amazon.com/ses/latest/dg/managing-ip-pools.html), managed by AWS.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.DedicatedIpAssignment("example", {
    destinationPoolName: "my-pool",
    ip: "0.0.0.0",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.DedicatedIpAssignment("example",
    destination_pool_name="my-pool",
    ip="0.0.0.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.DedicatedIpAssignment("example", new()
    {
        DestinationPoolName = "my-pool",
        Ip = "0.0.0.0",
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
		_, err := sesv2.NewDedicatedIpAssignment(ctx, "example", &sesv2.DedicatedIpAssignmentArgs{
			DestinationPoolName: pulumi.String("my-pool"),
			Ip:                  pulumi.String("0.0.0.0"),
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
import com.pulumi.aws.sesv2.DedicatedIpAssignment;
import com.pulumi.aws.sesv2.DedicatedIpAssignmentArgs;
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
        var example = new DedicatedIpAssignment("example", DedicatedIpAssignmentArgs.builder()        
            .destinationPoolName("my-pool")
            .ip("0.0.0.0")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:DedicatedIpAssignment
    properties:
      destinationPoolName: my-pool
      ip: 0.0.0.0
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Dedicated IP Assignment using the `id`, which is a comma-separated string made up of `ip` and `destination_pool_name`. For example:

```sh
 $ pulumi import aws:sesv2/dedicatedIpAssignment:DedicatedIpAssignment example "0.0.0.0,my-pool"
```
 