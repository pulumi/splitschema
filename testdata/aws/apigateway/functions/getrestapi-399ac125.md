Use this data source to get the id and root_resource_id of a REST API in
API Gateway. To fetch the REST API you must provide a name to match against.
As there is no unique name constraint on REST APIs this data source will
error if there is more than one match.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myRestApi = aws.apigateway.getRestApi({
    name: "my-rest-api",
});
```
```python
import pulumi
import pulumi_aws as aws

my_rest_api = aws.apigateway.get_rest_api(name="my-rest-api")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myRestApi = Aws.ApiGateway.GetRestApi.Invoke(new()
    {
        Name = "my-rest-api",
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
		_, err := apigateway.LookupRestApi(ctx, &apigateway.LookupRestApiArgs{
			Name: "my-rest-api",
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
import com.pulumi.aws.apigateway.inputs.GetRestApiArgs;
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
        final var myRestApi = ApigatewayFunctions.getRestApi(GetRestApiArgs.builder()
            .name("my-rest-api")
            .build());

    }
}
```
```yaml
variables:
  myRestApi:
    fn::invoke:
      Function: aws:apigateway:getRestApi
      Arguments:
        name: my-rest-api
```
{{% /example %}}
{{% /examples %}}