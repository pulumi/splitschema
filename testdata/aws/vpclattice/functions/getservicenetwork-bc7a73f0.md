Data source for managing an AWS VPC Lattice Service Network.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.vpclattice.getServiceNetwork({
    serviceNetworkIdentifier: "snsa-01112223334445556",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.get_service_network(service_network_identifier="snsa-01112223334445556")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.VpcLattice.GetServiceNetwork.Invoke(new()
    {
        ServiceNetworkIdentifier = "snsa-01112223334445556",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.LookupServiceNetwork(ctx, &vpclattice.LookupServiceNetworkArgs{
			ServiceNetworkIdentifier: "snsa-01112223334445556",
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
import com.pulumi.aws.vpclattice.VpclatticeFunctions;
import com.pulumi.aws.vpclattice.inputs.GetServiceNetworkArgs;
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
        final var example = VpclatticeFunctions.getServiceNetwork(GetServiceNetworkArgs.builder()
            .serviceNetworkIdentifier("snsa-01112223334445556")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:vpclattice:getServiceNetwork
      Arguments:
        serviceNetworkIdentifier: snsa-01112223334445556
```
{{% /example %}}
{{% /examples %}}