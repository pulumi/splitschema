Provides an AppSync API Cache.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGraphQLApi = new aws.appsync.GraphQLApi("exampleGraphQLApi", {authenticationType: "API_KEY"});
const exampleApiCache = new aws.appsync.ApiCache("exampleApiCache", {
    apiId: exampleGraphQLApi.id,
    apiCachingBehavior: "FULL_REQUEST_CACHING",
    type: "LARGE",
    ttl: 900,
});
```
```python
import pulumi
import pulumi_aws as aws

example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
example_api_cache = aws.appsync.ApiCache("exampleApiCache",
    api_id=example_graph_ql_api.id,
    api_caching_behavior="FULL_REQUEST_CACHING",
    type="LARGE",
    ttl=900)
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

    var exampleApiCache = new Aws.AppSync.ApiCache("exampleApiCache", new()
    {
        ApiId = exampleGraphQLApi.Id,
        ApiCachingBehavior = "FULL_REQUEST_CACHING",
        Type = "LARGE",
        Ttl = 900,
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
		_, err = appsync.NewApiCache(ctx, "exampleApiCache", &appsync.ApiCacheArgs{
			ApiId:              exampleGraphQLApi.ID(),
			ApiCachingBehavior: pulumi.String("FULL_REQUEST_CACHING"),
			Type:               pulumi.String("LARGE"),
			Ttl:                pulumi.Int(900),
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
import com.pulumi.aws.appsync.ApiCache;
import com.pulumi.aws.appsync.ApiCacheArgs;
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

        var exampleApiCache = new ApiCache("exampleApiCache", ApiCacheArgs.builder()        
            .apiId(exampleGraphQLApi.id())
            .apiCachingBehavior("FULL_REQUEST_CACHING")
            .type("LARGE")
            .ttl(900)
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
  exampleApiCache:
    type: aws:appsync:ApiCache
    properties:
      apiId: ${exampleGraphQLApi.id}
      apiCachingBehavior: FULL_REQUEST_CACHING
      type: LARGE
      ttl: 900
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appsync_api_cache` using the AppSync API ID. For example:

```sh
 $ pulumi import aws:appsync/apiCache:ApiCache example xxxxx
```
 