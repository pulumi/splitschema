Provides a global network resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkmanager.GlobalNetwork("example", {description: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.GlobalNetwork("example", description="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkManager.GlobalNetwork("example", new()
    {
        Description = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkmanager.NewGlobalNetwork(ctx, "example", &networkmanager.GlobalNetworkArgs{
			Description: pulumi.String("example"),
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
import com.pulumi.aws.networkmanager.GlobalNetwork;
import com.pulumi.aws.networkmanager.GlobalNetworkArgs;
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
        var example = new GlobalNetwork("example", GlobalNetworkArgs.builder()        
            .description("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkmanager:GlobalNetwork
    properties:
      description: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_networkmanager_global_network` using the global network ID. For example:

```sh
 $ pulumi import aws:networkmanager/globalNetwork:GlobalNetwork example global-network-0d47f6t230mz46dy4
```
 