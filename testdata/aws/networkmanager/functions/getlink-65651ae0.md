Retrieve information about a link.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkmanager.getLink({
    globalNetworkId: _var.global_network_id,
    linkId: _var.link_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkmanager.get_link(global_network_id=var["global_network_id"],
    link_id=var["link_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkManager.GetLink.Invoke(new()
    {
        GlobalNetworkId = @var.Global_network_id,
        LinkId = @var.Link_id,
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
		_, err := networkmanager.LookupLink(ctx, &networkmanager.LookupLinkArgs{
			GlobalNetworkId: _var.Global_network_id,
			LinkId:          _var.Link_id,
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
import com.pulumi.aws.networkmanager.inputs.GetLinkArgs;
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
        final var example = NetworkmanagerFunctions.getLink(GetLinkArgs.builder()
            .globalNetworkId(var_.global_network_id())
            .linkId(var_.link_id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkmanager:getLink
      Arguments:
        globalNetworkId: ${var.global_network_id}
        linkId: ${var.link_id}
```
{{% /example %}}
{{% /examples %}}