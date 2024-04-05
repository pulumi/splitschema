Data source for retrieving Router Configuration instructions for a given AWS Direct Connect Virtual Interface and Router Type.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.directconnect.getRouterConfiguration({
    routerTypeIdentifier: "CiscoSystemsInc-2900SeriesRouters-IOS124",
    virtualInterfaceId: "dxvif-abcde123",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.get_router_configuration(router_type_identifier="CiscoSystemsInc-2900SeriesRouters-IOS124",
    virtual_interface_id="dxvif-abcde123")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.DirectConnect.GetRouterConfiguration.Invoke(new()
    {
        RouterTypeIdentifier = "CiscoSystemsInc-2900SeriesRouters-IOS124",
        VirtualInterfaceId = "dxvif-abcde123",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.GetRouterConfiguration(ctx, &directconnect.GetRouterConfigurationArgs{
			RouterTypeIdentifier: "CiscoSystemsInc-2900SeriesRouters-IOS124",
			VirtualInterfaceId:   "dxvif-abcde123",
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
import com.pulumi.aws.directconnect.DirectconnectFunctions;
import com.pulumi.aws.directconnect.inputs.GetRouterConfigurationArgs;
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
        final var example = DirectconnectFunctions.getRouterConfiguration(GetRouterConfigurationArgs.builder()
            .routerTypeIdentifier("CiscoSystemsInc-2900SeriesRouters-IOS124")
            .virtualInterfaceId("dxvif-abcde123")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:directconnect:getRouterConfiguration
      Arguments:
        routerTypeIdentifier: CiscoSystemsInc-2900SeriesRouters-IOS124
        virtualInterfaceId: dxvif-abcde123
```
{{% /example %}}
{{% /examples %}}