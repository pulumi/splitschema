Use this data source to get the id of a VPC Link in
API Gateway. To fetch the VPC Link you must provide a name to match against.
As there is no unique name constraint on API Gateway VPC Links this data source will
error if there is more than one match.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myApiGatewayVpcLink = aws.apigateway.getVpcLink({
    name: "my-vpc-link",
});
```
```python
import pulumi
import pulumi_aws as aws

my_api_gateway_vpc_link = aws.apigateway.get_vpc_link(name="my-vpc-link")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myApiGatewayVpcLink = Aws.ApiGateway.GetVpcLink.Invoke(new()
    {
        Name = "my-vpc-link",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.LookupVpcLink(ctx, &apigateway.LookupVpcLinkArgs{
			Name: "my-vpc-link",
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
import com.pulumi.aws.apigateway.ApigatewayFunctions;
import com.pulumi.aws.apigateway.inputs.GetVpcLinkArgs;
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
        final var myApiGatewayVpcLink = ApigatewayFunctions.getVpcLink(GetVpcLinkArgs.builder()
            .name("my-vpc-link")
            .build());

    }
}
```
```yaml
variables:
  myApiGatewayVpcLink:
    fn::invoke:
      Function: aws:apigateway:getVpcLink
      Arguments:
        name: my-vpc-link
```
{{% /example %}}
{{% /examples %}}