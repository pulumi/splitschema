Provides an AppSync Type.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGraphQLApi = new aws.appsync.GraphQLApi("exampleGraphQLApi", {authenticationType: "API_KEY"});
const exampleType = new aws.appsync.Type("exampleType", {
    apiId: exampleGraphQLApi.id,
    format: "SDL",
    definition: `type Mutation

{
putPost(id: ID!,title: String! ): Post

}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
example_type = aws.appsync.Type("exampleType",
    api_id=example_graph_ql_api.id,
    format="SDL",
    definition="""type Mutation

{
putPost(id: ID!,title: String! ): Post

}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new()
    {
        AuthenticationType = "API_KEY",
    });

    var exampleType = new Aws.AppSync.Type("exampleType", new()
    {
        ApiId = exampleGraphQLApi.Id,
        Format = "SDL",
        Definition = @"type Mutation

{
putPost(id: ID!,title: String! ): Post

}
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("API_KEY"),
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewType(ctx, "exampleType", &appsync.TypeArgs{
			ApiId:  exampleGraphQLApi.ID(),
			Format: pulumi.String("SDL"),
			Definition: pulumi.String(`type Mutation

{
putPost(id: ID!,title: String! ): Post

}
`),
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
import com.pulumi.aws.appsync.GraphQLApi;
import com.pulumi.aws.appsync.GraphQLApiArgs;
import com.pulumi.aws.appsync.Type;
import com.pulumi.aws.appsync.TypeArgs;
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
        var exampleGraphQLApi = new GraphQLApi("exampleGraphQLApi", GraphQLApiArgs.builder()        
            .authenticationType("API_KEY")
            .build());

        var exampleType = new Type("exampleType", TypeArgs.builder()        
            .apiId(exampleGraphQLApi.id())
            .format("SDL")
            .definition("""
type Mutation

{
putPost(id: ID!,title: String! ): Post

}
            """)
            .build());

    }
}
```
```yaml
resources:
  exampleGraphQLApi:
    type: aws:appsync:GraphQLApi
    properties:
      authenticationType: API_KEY
  exampleType:
    type: aws:appsync:Type
    properties:
      apiId: ${exampleGraphQLApi.id}
      format: SDL
      definition: |
        type Mutation

        {
        putPost(id: ID!,title: String! ): Post

        }
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Appsync Types using the `id`. For example:

```sh
 $ pulumi import aws:appsync/type:Type example api-id:format:name
```
 