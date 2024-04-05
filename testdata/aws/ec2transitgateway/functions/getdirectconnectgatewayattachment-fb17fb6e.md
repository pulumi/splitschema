Get information on an EC2 Transit Gateway's attachment to a Direct Connect Gateway.

{{% examples %}}
## Example Usage
{{% example %}}
### By Transit Gateway and Direct Connect Gateway Identifiers

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2transitgateway.getDirectConnectGatewayAttachment({
    transitGatewayId: aws_ec2_transit_gateway.example.id,
    dxGatewayId: aws_dx_gateway.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2transitgateway.get_direct_connect_gateway_attachment(transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
    dx_gateway_id=aws_dx_gateway["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2TransitGateway.GetDirectConnectGatewayAttachment.Invoke(new()
    {
        TransitGatewayId = aws_ec2_transit_gateway.Example.Id,
        DxGatewayId = aws_dx_gateway.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2transitgateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2transitgateway.GetDirectConnectGatewayAttachment(ctx, &ec2transitgateway.GetDirectConnectGatewayAttachmentArgs{
			TransitGatewayId: pulumi.StringRef(aws_ec2_transit_gateway.Example.Id),
			DxGatewayId:      pulumi.StringRef(aws_dx_gateway.Example.Id),
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
import com.pulumi.aws.ec2transitgateway.Ec2transitgatewayFunctions;
import com.pulumi.aws.ec2transitgateway.inputs.GetDirectConnectGatewayAttachmentArgs;
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
        final var example = Ec2transitgatewayFunctions.getDirectConnectGatewayAttachment(GetDirectConnectGatewayAttachmentArgs.builder()
            .transitGatewayId(aws_ec2_transit_gateway.example().id())
            .dxGatewayId(aws_dx_gateway.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2transitgateway:getDirectConnectGatewayAttachment
      Arguments:
        transitGatewayId: ${aws_ec2_transit_gateway.example.id}
        dxGatewayId: ${aws_dx_gateway.example.id}
```
{{% /example %}}
{{% /examples %}}