Use this data source to get the custom domain name for use with AWS API Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.apigateway.getDomainName({
    domainName: "api.example.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigateway.get_domain_name(domain_name="api.example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ApiGateway.GetDomainName.Invoke(new()
    {
        DomainName = "api.example.com",
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
		_, err := apigateway.LookupDomainName(ctx, &apigateway.LookupDomainNameArgs{
			DomainName: "api.example.com",
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
import com.pulumi.aws.apigateway.inputs.GetDomainNameArgs;
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
        final var example = ApigatewayFunctions.getDomainName(GetDomainNameArgs.builder()
            .domainName("api.example.com")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:apigateway:getDomainName
      Arguments:
        domainName: api.example.com
```
{{% /example %}}
{{% /examples %}}