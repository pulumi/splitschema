Provides an AppSync API Key.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGraphQLApi = new aws.appsync.GraphQLApi("exampleGraphQLApi", {authenticationType: "API_KEY"});
const exampleApiKey = new aws.appsync.ApiKey("exampleApiKey", {
    apiId: exampleGraphQLApi.id,
    expires: "2018-05-03T04:00:00Z",
});
```
```python
import pulumi
import pulumi_aws as aws

example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
example_api_key = aws.appsync.ApiKey("exampleApiKey",
    api_id=example_graph_ql_api.id,
    expires="2018-05-03T04:00:00Z")
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

    var exampleApiKey = new Aws.AppSync.ApiKey("exampleApiKey", new()
    {
        ApiId = exampleGraphQLApi.Id,
        Expires = "2018-05-03T04:00:00Z",
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
		_, err = appsync.NewApiKey(ctx, "exampleApiKey", &appsync.ApiKeyArgs{
			ApiId:   exampleGraphQLApi.ID(),
			Expires: pulumi.String("2018-05-03T04:00:00Z"),
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
import com.pulumi.aws.appsync.ApiKey;
import com.pulumi.aws.appsync.ApiKeyArgs;
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

        var exampleApiKey = new ApiKey("exampleApiKey", ApiKeyArgs.builder()        
            .apiId(exampleGraphQLApi.id())
            .expires("2018-05-03T04:00:00Z")
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
  exampleApiKey:
    type: aws:appsync:ApiKey
    properties:
      apiId: ${exampleGraphQLApi.id}
      expires: 2018-05-03T04:00:00Z
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appsync_api_key` using the AppSync API ID and key separated by `:`. For example:

```sh
 $ pulumi import aws:appsync/apiKey:ApiKey example xxxxx:yyyyy
```
 