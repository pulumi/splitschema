Provides an API Gateway Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
const myDemoResource = new aws.apigateway.Resource("myDemoResource", {
    restApi: myDemoAPI.id,
    parentId: myDemoAPI.rootResourceId,
    pathPart: "mydemoresource",
});
```
```python
import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
my_demo_resource = aws.apigateway.Resource("myDemoResource",
    rest_api=my_demo_api.id,
    parent_id=my_demo_api.root_resource_id,
    path_part="mydemoresource")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new()
    {
        Description = "This is my API for demonstration purposes",
    });

    var myDemoResource = new Aws.ApiGateway.Resource("myDemoResource", new()
    {
        RestApi = myDemoAPI.Id,
        ParentId = myDemoAPI.RootResourceId,
        PathPart = "mydemoresource",
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
		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewResource(ctx, "myDemoResource", &apigateway.ResourceArgs{
			RestApi:  myDemoAPI.ID(),
			ParentId: myDemoAPI.RootResourceId,
			PathPart: pulumi.String("mydemoresource"),
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.apigateway.RestApiArgs;
import com.pulumi.aws.apigateway.Resource;
import com.pulumi.aws.apigateway.ResourceArgs;
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
        var myDemoAPI = new RestApi("myDemoAPI", RestApiArgs.builder()        
            .description("This is my API for demonstration purposes")
            .build());

        var myDemoResource = new Resource("myDemoResource", ResourceArgs.builder()        
            .restApi(myDemoAPI.id())
            .parentId(myDemoAPI.rootResourceId())
            .pathPart("mydemoresource")
            .build());

    }
}
```
```yaml
resources:
  myDemoAPI:
    type: aws:apigateway:RestApi
    properties:
      description: This is my API for demonstration purposes
  myDemoResource:
    type: aws:apigateway:Resource
    properties:
      restApi: ${myDemoAPI.id}
      parentId: ${myDemoAPI.rootResourceId}
      pathPart: mydemoresource
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_api_gateway_resource` using `REST-API-ID/RESOURCE-ID`. For example:

```sh
 $ pulumi import aws:apigateway/resource:Resource example 12345abcde/67890fghij
```
 