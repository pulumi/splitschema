Provides a settings of an API Gateway Documentation Part.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {});
const exampleDocumentationPart = new aws.apigateway.DocumentationPart("exampleDocumentationPart", {
    location: {
        type: "METHOD",
        method: "GET",
        path: "/example",
    },
    properties: "{\"description\":\"Example description\"}",
    restApiId: exampleRestApi.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_rest_api = aws.apigateway.RestApi("exampleRestApi")
example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
    location=aws.apigateway.DocumentationPartLocationArgs(
        type="METHOD",
        method="GET",
        path="/example",
    ),
    properties="{\"description\":\"Example description\"}",
    rest_api_id=example_rest_api.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi");

    var exampleDocumentationPart = new Aws.ApiGateway.DocumentationPart("exampleDocumentationPart", new()
    {
        Location = new Aws.ApiGateway.Inputs.DocumentationPartLocationArgs
        {
            Type = "METHOD",
            Method = "GET",
            Path = "/example",
        },
        Properties = "{\"description\":\"Example description\"}",
        RestApiId = exampleRestApi.Id,
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
		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewDocumentationPart(ctx, "exampleDocumentationPart", &apigateway.DocumentationPartArgs{
			Location: &apigateway.DocumentationPartLocationArgs{
				Type:   pulumi.String("METHOD"),
				Method: pulumi.String("GET"),
				Path:   pulumi.String("/example"),
			},
			Properties: pulumi.String("{\"description\":\"Example description\"}"),
			RestApiId:  exampleRestApi.ID(),
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
import com.pulumi.aws.apigateway.DocumentationPart;
import com.pulumi.aws.apigateway.DocumentationPartArgs;
import com.pulumi.aws.apigateway.inputs.DocumentationPartLocationArgs;
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
        var exampleRestApi = new RestApi("exampleRestApi");

        var exampleDocumentationPart = new DocumentationPart("exampleDocumentationPart", DocumentationPartArgs.builder()        
            .location(DocumentationPartLocationArgs.builder()
                .type("METHOD")
                .method("GET")
                .path("/example")
                .build())
            .properties("{\"description\":\"Example description\"}")
            .restApiId(exampleRestApi.id())
            .build());

    }
}
```
```yaml
resources:
  exampleDocumentationPart:
    type: aws:apigateway:DocumentationPart
    properties:
      location:
        type: METHOD
        method: GET
        path: /example
      properties: '{"description":"Example description"}'
      restApiId: ${exampleRestApi.id}
  exampleRestApi:
    type: aws:apigateway:RestApi
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway documentation_parts using `REST-API-ID/DOC-PART-ID`. For example:

```sh
 $ pulumi import aws:apigateway/documentationPart:DocumentationPart example 5i4e1ko720/3oyy3t
```
 