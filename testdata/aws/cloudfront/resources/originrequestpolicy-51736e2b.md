{{% examples %}}
## Example Usage
{{% example %}}

The following example below creates a CloudFront origin request policy.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.OriginRequestPolicy("example", {
    comment: "example comment",
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
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.OriginRequestPolicy("example",
    comment="example comment",
    cookies_config=aws.cloudfront.OriginRequestPolicyCookiesConfigArgs(
        cookie_behavior="whitelist",
        cookies=aws.cloudfront.OriginRequestPolicyCookiesConfigCookiesArgs(
            items=["example"],
        ),
    ),
    headers_config=aws.cloudfront.OriginRequestPolicyHeadersConfigArgs(
        header_behavior="whitelist",
        headers=aws.cloudfront.OriginRequestPolicyHeadersConfigHeadersArgs(
            items=["example"],
        ),
    ),
    query_strings_config=aws.cloudfront.OriginRequestPolicyQueryStringsConfigArgs(
        query_string_behavior="whitelist",
        query_strings=aws.cloudfront.OriginRequestPolicyQueryStringsConfigQueryStringsArgs(
            items=["example"],
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
    var example = new Aws.CloudFront.OriginRequestPolicy("example", new()
    {
        Comment = "example comment",
        CookiesConfig = new Aws.CloudFront.Inputs.OriginRequestPolicyCookiesConfigArgs
        {
            CookieBehavior = "whitelist",
            Cookies = new Aws.CloudFront.Inputs.OriginRequestPolicyCookiesConfigCookiesArgs
            {
                Items = new[]
                {
                    "example",
                },
            },
        },
        HeadersConfig = new Aws.CloudFront.Inputs.OriginRequestPolicyHeadersConfigArgs
        {
            HeaderBehavior = "whitelist",
            Headers = new Aws.CloudFront.Inputs.OriginRequestPolicyHeadersConfigHeadersArgs
            {
                Items = new[]
                {
                    "example",
                },
            },
        },
        QueryStringsConfig = new Aws.CloudFront.Inputs.OriginRequestPolicyQueryStringsConfigArgs
        {
            QueryStringBehavior = "whitelist",
            QueryStrings = new Aws.CloudFront.Inputs.OriginRequestPolicyQueryStringsConfigQueryStringsArgs
            {
                Items = new[]
                {
                    "example",
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
		_, err := cloudfront.NewOriginRequestPolicy(ctx, "example", &cloudfront.OriginRequestPolicyArgs{
			Comment: pulumi.String("example comment"),
			CookiesConfig: &cloudfront.OriginRequestPolicyCookiesConfigArgs{
				CookieBehavior: pulumi.String("whitelist"),
				Cookies: &cloudfront.OriginRequestPolicyCookiesConfigCookiesArgs{
					Items: pulumi.StringArray{
						pulumi.String("example"),
					},
				},
			},
			HeadersConfig: &cloudfront.OriginRequestPolicyHeadersConfigArgs{
				HeaderBehavior: pulumi.String("whitelist"),
				Headers: &cloudfront.OriginRequestPolicyHeadersConfigHeadersArgs{
					Items: pulumi.StringArray{
						pulumi.String("example"),
					},
				},
			},
			QueryStringsConfig: &cloudfront.OriginRequestPolicyQueryStringsConfigArgs{
				QueryStringBehavior: pulumi.String("whitelist"),
				QueryStrings: &cloudfront.OriginRequestPolicyQueryStringsConfigQueryStringsArgs{
					Items: pulumi.StringArray{
						pulumi.String("example"),
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
import com.pulumi.aws.cloudfront.OriginRequestPolicy;
import com.pulumi.aws.cloudfront.OriginRequestPolicyArgs;
import com.pulumi.aws.cloudfront.inputs.OriginRequestPolicyCookiesConfigArgs;
import com.pulumi.aws.cloudfront.inputs.OriginRequestPolicyCookiesConfigCookiesArgs;
import com.pulumi.aws.cloudfront.inputs.OriginRequestPolicyHeadersConfigArgs;
import com.pulumi.aws.cloudfront.inputs.OriginRequestPolicyHeadersConfigHeadersArgs;
import com.pulumi.aws.cloudfront.inputs.OriginRequestPolicyQueryStringsConfigArgs;
import com.pulumi.aws.cloudfront.inputs.OriginRequestPolicyQueryStringsConfigQueryStringsArgs;
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
        var example = new OriginRequestPolicy("example", OriginRequestPolicyArgs.builder()        
            .comment("example comment")
            .cookiesConfig(OriginRequestPolicyCookiesConfigArgs.builder()
                .cookieBehavior("whitelist")
                .cookies(OriginRequestPolicyCookiesConfigCookiesArgs.builder()
                    .items("example")
                    .build())
                .build())
            .headersConfig(OriginRequestPolicyHeadersConfigArgs.builder()
                .headerBehavior("whitelist")
                .headers(OriginRequestPolicyHeadersConfigHeadersArgs.builder()
                    .items("example")
                    .build())
                .build())
            .queryStringsConfig(OriginRequestPolicyQueryStringsConfigArgs.builder()
                .queryStringBehavior("whitelist")
                .queryStrings(OriginRequestPolicyQueryStringsConfigQueryStringsArgs.builder()
                    .items("example")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:OriginRequestPolicy
    properties:
      comment: example comment
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

Using `pulumi import`, import Cloudfront Origin Request Policies using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/originRequestPolicy:OriginRequestPolicy policy ccca32ef-dce3-4df3-80df-1bd3000bc4d3
```
 