{{% examples %}}
## Example Usage
{{% example %}}

Use the `aws.cloudfront.CachePolicy` resource to create a cache policy for CloudFront.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.CachePolicy("example", {
    comment: "test comment",
    defaultTtl: 50,
    maxTtl: 100,
    minTtl: 1,
    parametersInCacheKeyAndForwardedToOrigin: {
        cookiesConfig: {
            cookieBehavior: "whitelist",
            cookies: {
                items: ["example"],
            },
        },
        headersConfig: {
            headerBehavior: "whitelist",
            headers: {
                items: ["example"],
            },
        },
        queryStringsConfig: {
            queryStringBehavior: "whitelist",
            queryStrings: {
                items: ["example"],
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.CachePolicy("example",
    comment="test comment",
    default_ttl=50,
    max_ttl=100,
    min_ttl=1,
    parameters_in_cache_key_and_forwarded_to_origin=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs(
        cookies_config=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigArgs(
            cookie_behavior="whitelist",
            cookies=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesArgs(
                items=["example"],
            ),
        ),
        headers_config=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigArgs(
            header_behavior="whitelist",
            headers=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersArgs(
                items=["example"],
            ),
        ),
        query_strings_config=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigArgs(
            query_string_behavior="whitelist",
            query_strings=aws.cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsArgs(
                items=["example"],
            ),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.CachePolicy("example", new()
    {
        Comment = "test comment",
        DefaultTtl = 50,
        MaxTtl = 100,
        MinTtl = 1,
        ParametersInCacheKeyAndForwardedToOrigin = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs
        {
            CookiesConfig = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigArgs
            {
                CookieBehavior = "whitelist",
                Cookies = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesArgs
                {
                    Items = new[]
                    {
                        "example",
                    },
                },
            },
            HeadersConfig = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigArgs
            {
                HeaderBehavior = "whitelist",
                Headers = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersArgs
                {
                    Items = new[]
                    {
                        "example",
                    },
                },
            },
            QueryStringsConfig = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigArgs
            {
                QueryStringBehavior = "whitelist",
                QueryStrings = new Aws.CloudFront.Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsArgs
                {
                    Items = new[]
                    {
                        "example",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewCachePolicy(ctx, "example", &cloudfront.CachePolicyArgs{
			Comment:    pulumi.String("test comment"),
			DefaultTtl: pulumi.Int(50),
			MaxTtl:     pulumi.Int(100),
			MinTtl:     pulumi.Int(1),
			ParametersInCacheKeyAndForwardedToOrigin: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs{
				CookiesConfig: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigArgs{
					CookieBehavior: pulumi.String("whitelist"),
					Cookies: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesArgs{
						Items: pulumi.StringArray{
							pulumi.String("example"),
						},
					},
				},
				HeadersConfig: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigArgs{
					HeaderBehavior: pulumi.String("whitelist"),
					Headers: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersArgs{
						Items: pulumi.StringArray{
							pulumi.String("example"),
						},
					},
				},
				QueryStringsConfig: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigArgs{
					QueryStringBehavior: pulumi.String("whitelist"),
					QueryStrings: &cloudfront.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsArgs{
						Items: pulumi.StringArray{
							pulumi.String("example"),
						},
					},
				},
			},
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
import com.pulumi.aws.cloudfront.CachePolicy;
import com.pulumi.aws.cloudfront.CachePolicyArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigArgs;
import com.pulumi.aws.cloudfront.inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsArgs;
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
        var example = new CachePolicy("example", CachePolicyArgs.builder()        
            .comment("test comment")
            .defaultTtl(50)
            .maxTtl(100)
            .minTtl(1)
            .parametersInCacheKeyAndForwardedToOrigin(CachePolicyParametersInCacheKeyAndForwardedToOriginArgs.builder()
                .cookiesConfig(CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigArgs.builder()
                    .cookieBehavior("whitelist")
                    .cookies(CachePolicyParametersInCacheKeyAndForwardedToOriginCookiesConfigCookiesArgs.builder()
                        .items("example")
                        .build())
                    .build())
                .headersConfig(CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigArgs.builder()
                    .headerBehavior("whitelist")
                    .headers(CachePolicyParametersInCacheKeyAndForwardedToOriginHeadersConfigHeadersArgs.builder()
                        .items("example")
                        .build())
                    .build())
                .queryStringsConfig(CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigArgs.builder()
                    .queryStringBehavior("whitelist")
                    .queryStrings(CachePolicyParametersInCacheKeyAndForwardedToOriginQueryStringsConfigQueryStringsArgs.builder()
                        .items("example")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:CachePolicy
    properties:
      comment: test comment
      defaultTtl: 50
      maxTtl: 100
      minTtl: 1
      parametersInCacheKeyAndForwardedToOrigin:
        cookiesConfig:
          cookieBehavior: whitelist
          cookies:
            items:
              - example
        headersConfig:
          headerBehavior: whitelist
          headers:
            items:
              - example
        queryStringsConfig:
          queryStringBehavior: whitelist
          queryStrings:
            items:
              - example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront cache policies using the `id` of the cache policy. For example:

```sh
 $ pulumi import aws:cloudfront/cachePolicy:CachePolicy policy 658327ea-f89d-4fab-a63d-7e88639e58f6
```
 