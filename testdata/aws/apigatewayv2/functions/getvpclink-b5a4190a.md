Data source for managing an AWS API Gateway V2 VPC Link.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.apigatewayv2.getVpcLink({
    vpcLinkId: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.get_vpc_link(vpc_link_id="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ApiGatewayV2.GetVpcLink.Invoke(new()
    {
        VpcLinkId = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.LookupVpcLink(ctx, &apigatewayv2.LookupVpcLinkArgs{
			VpcLinkId: "example",
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
import com.pulumi.aws.apigatewayv2.Apigatewayv2Functions;
import com.pulumi.aws.apigatewayv2.inputs.GetVpcLinkArgs;
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
        final var example = Apigatewayv2Functions.getVpcLink(GetVpcLinkArgs.builder()
            .vpcLinkId("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:apigatewayv2:getVpcLink
      Arguments:
        vpcLinkId: example
```
{{% /example %}}
{{% /examples %}}