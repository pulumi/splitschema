Use this data source to get information about a Network Interface.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bar = aws.ec2.getNetworkInterface({
    id: "eni-01234567",
});
```
```python
import pulumi
import pulumi_aws as aws

bar = aws.ec2.get_network_interface(id="eni-01234567")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bar = Aws.Ec2.GetNetworkInterface.Invoke(new()
    {
        Id = "eni-01234567",
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
		_, err := ec2.LookupNetworkInterface(ctx, &ec2.LookupNetworkInterfaceArgs{
			Id: pulumi.StringRef("eni-01234567"),
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
import com.pulumi.aws.ec2.inputs.GetNetworkInterfaceArgs;
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
        final var bar = Ec2Functions.getNetworkInterface(GetNetworkInterfaceArgs.builder()
            .id("eni-01234567")
            .build());

    }
}
```
```yaml
variables:
  bar:
    fn::invoke:
      Function: aws:ec2:getNetworkInterface
      Arguments:
        id: eni-01234567
```
{{% /example %}}
{{% /examples %}}