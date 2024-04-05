Provides a Direct Connect Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.Gateway("example", {amazonSideAsn: "64512"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.Gateway("example", amazon_side_asn="64512")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DirectConnect.Gateway("example", new()
    {
        AmazonSideAsn = "64512",
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
		_, err := directconnect.NewGateway(ctx, "example", &directconnect.GatewayArgs{
			AmazonSideAsn: pulumi.String("64512"),
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
import com.pulumi.aws.directconnect.Gateway;
import com.pulumi.aws.directconnect.GatewayArgs;
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
        var example = new Gateway("example", GatewayArgs.builder()        
            .amazonSideAsn("64512")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directconnect:Gateway
    properties:
      amazonSideAsn: '64512'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect Gateways using the gateway `id`. For example:

```sh
 $ pulumi import aws:directconnect/gateway:Gateway test abcd1234-dcba-5678-be23-cdef9876ab45
```
 