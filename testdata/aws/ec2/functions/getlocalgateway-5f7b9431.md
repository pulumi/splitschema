Provides details about an EC2 Local Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how one might accept a local gateway id as a variable.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const localGatewayId = config.requireObject("localGatewayId");
const selected = aws.ec2.getLocalGateway({
    id: localGatewayId,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
local_gateway_id = config.require_object("localGatewayId")
selected = aws.ec2.get_local_gateway(id=local_gateway_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var localGatewayId = config.RequireObject<dynamic>("localGatewayId");
    var selected = Aws.Ec2.GetLocalGateway.Invoke(new()
    {
        Id = localGatewayId,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		localGatewayId := cfg.RequireObject("localGatewayId")
		_, err := ec2.GetLocalGateway(ctx, &ec2.GetLocalGatewayArgs{
			Id: pulumi.StringRef(localGatewayId),
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
import com.pulumi.aws.ec2.inputs.GetLocalGatewayArgs;
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
        final var config = ctx.config();
        final var localGatewayId = config.get("localGatewayId");
        final var selected = Ec2Functions.getLocalGateway(GetLocalGatewayArgs.builder()
            .id(localGatewayId)
            .build());

    }
}
```
```yaml
configuration:
  localGatewayId:
    type: dynamic
variables:
  selected:
    fn::invoke:
      Function: aws:ec2:getLocalGateway
      Arguments:
        id: ${localGatewayId}
```
{{% /example %}}
{{% /examples %}}