Provides a resource to manage an API Gateway Documentation Version.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {});
const exampleDocumentationPart = new aws.apigateway.DocumentationPart("exampleDocumentationPart", {
    location: {
        type: "API",
    },
    properties: "{\"description\":\"Example\"}",
    restApiId: exampleRestApi.id,
});
const exampleDocumentationVersion = new aws.apigateway.DocumentationVersion("exampleDocumentationVersion", {
    version: "example_version",
    restApiId: exampleRestApi.id,
    description: "Example description",
}, {
    dependsOn: [exampleDocumentationPart],
});
```
```python
import pulumi
import pulumi_aws as aws

example_rest_api = aws.apigateway.RestApi("exampleRestApi")
example_documentation_part = aws.apigateway.DocumentationPart("exampleDocumentationPart",
    location=aws.apigateway.DocumentationPartLocationArgs(
        type="API",
    ),
    properties="{\"description\":\"Example\"}",
    rest_api_id=example_rest_api.id)
example_documentation_version = aws.apigateway.DocumentationVersion("exampleDocumentationVersion",
    version="example_version",
    rest_api_id=example_rest_api.id,
    description="Example description",
    opts=pulumi.ResourceOptions(depends_on=[example_documentation_part]))
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
            Type = "API",
        },
        Properties = "{\"description\":\"Example\"}",
        RestApiId = exampleRestApi.Id,
    });

    var exampleDocumentationVersion = new Aws.ApiGateway.DocumentationVersion("exampleDocumentationVersion", new()
    {
        Version = "example_version",
        RestApiId = exampleRestApi.Id,
        Description = "Example description",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleDocumentationPart,
        },
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
		exampleDocumentationPart, err := apigateway.NewDocumentationPart(ctx, "exampleDocumentationPart", &apigateway.DocumentationPartArgs{
			Location: &apigateway.DocumentationPartLocationArgs{
				Type: pulumi.String("API"),
			},
			Properties: pulumi.String("{\"description\":\"Example\"}"),
			RestApiId:  exampleRestApi.ID(),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewDocumentationVersion(ctx, "exampleDocumentationVersion", &apigateway.DocumentationVersionArgs{
			Version:     pulumi.String("example_version"),
			RestApiId:   exampleRestApi.ID(),
			Description: pulumi.String("Example description"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleDocumentationPart,
		}))
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
import com.pulumi.aws.apigateway.DocumentationVersion;
import com.pulumi.aws.apigateway.DocumentationVersionArgs;
import com.pulumi.resources.CustomResourceOptions;
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
                .type("API")
                .build())
            .properties("{\"description\":\"Example\"}")
            .restApiId(exampleRestApi.id())
            .build());

        var exampleDocumentationVersion = new DocumentationVersion("exampleDocumentationVersion", DocumentationVersionArgs.builder()        
            .version("example_version")
            .restApiId(exampleRestApi.id())
            .description("Example description")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleDocumentationPart)
                .build());

    }
}
```
```yaml
resources:
  exampleDocumentationVersion:
    type: aws:apigateway:DocumentationVersion
    properties:
      version: example_version
      restApiId: ${exampleRestApi.id}
      description: Example description
    options:
      dependson:
        - ${exampleDocumentationPart}
  exampleRestApi:
    type: aws:apigateway:RestApi
  exampleDocumentationPart:
    type: aws:apigateway:DocumentationPart
    properties:
      location:
        type: API
      properties: '{"description":"Example"}'
      restApiId: ${exampleRestApi.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway documentation versions using `REST-API-ID/VERSION`. For example:

```sh
 $ pulumi import aws:apigateway/documentationVersion:DocumentationVersion example 5i4e1ko720/example-version
```
 