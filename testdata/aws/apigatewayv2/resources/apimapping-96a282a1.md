Manages an Amazon API Gateway Version 2 API mapping.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.ApiMapping("example", {
    apiId: aws_apigatewayv2_api.example.id,
    domainName: aws_apigatewayv2_domain_name.example.id,
    stage: aws_apigatewayv2_stage.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.ApiMapping("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    domain_name=aws_apigatewayv2_domain_name["example"]["id"],
    stage=aws_apigatewayv2_stage["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.ApiMapping("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        DomainName = aws_apigatewayv2_domain_name.Example.Id,
        Stage = aws_apigatewayv2_stage.Example.Id,
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
		_, err := apigatewayv2.NewApiMapping(ctx, "example", &apigatewayv2.ApiMappingArgs{
			ApiId:      pulumi.Any(aws_apigatewayv2_api.Example.Id),
			DomainName: pulumi.Any(aws_apigatewayv2_domain_name.Example.Id),
			Stage:      pulumi.Any(aws_apigatewayv2_stage.Example.Id),
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
import com.pulumi.aws.apigatewayv2.ApiMapping;
import com.pulumi.aws.apigatewayv2.ApiMappingArgs;
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
        var example = new ApiMapping("example", ApiMappingArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .domainName(aws_apigatewayv2_domain_name.example().id())
            .stage(aws_apigatewayv2_stage.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:ApiMapping
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      domainName: ${aws_apigatewayv2_domain_name.example.id}
      stage: ${aws_apigatewayv2_stage.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_api_mapping` using the API mapping identifier and domain name. For example:

```sh
 $ pulumi import aws:apigatewayv2/apiMapping:ApiMapping example 1122334/ws-api.example.com
```
 