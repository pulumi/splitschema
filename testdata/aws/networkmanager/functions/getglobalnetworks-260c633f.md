Retrieve information about global networks.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkmanager.getGlobalNetworks({
    tags: {
        Env: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.get_global_networks(tags={
    "Env": "test",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkManager.GetGlobalNetworks.Invoke(new()
    {
        Tags = 
        {
            { "Env", "test" },
        },
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
		_, err := networkmanager.GetGlobalNetworks(ctx, &networkmanager.GetGlobalNetworksArgs{
			Tags: map[string]interface{}{
				"Env": "test",
			},
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
import com.pulumi.aws.networkmanager.NetworkmanagerFunctions;
import com.pulumi.aws.networkmanager.inputs.GetGlobalNetworksArgs;
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
        final var example = NetworkmanagerFunctions.getGlobalNetworks(GetGlobalNetworksArgs.builder()
            .tags(Map.of("Env", "test"))
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkmanager:getGlobalNetworks
      Arguments:
        tags:
          Env: test
```
{{% /example %}}
{{% /examples %}}