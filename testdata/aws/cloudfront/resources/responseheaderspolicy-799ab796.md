Provides a CloudFront response headers policy resource.
A response headers policy contains information about a set of HTTP response headers and their values.
After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution.
When it’s attached to a cache behavior, CloudFront adds the headers in the policy to every response that it sends for requests that match the cache behavior.

{{% examples %}}
## Example Usage
{{% example %}}

The example below creates a CloudFront response headers policy.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.ResponseHeadersPolicy("example", {
    comment: "test comment",
    corsConfig: {
        accessControlAllowCredentials: true,
        accessControlAllowHeaders: {
            items: ["test"],
        },
        accessControlAllowMethods: {
            items: ["GET"],
        },
        accessControlAllowOrigins: {
            items: ["test.example.comtest"],
        },
        originOverride: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.ResponseHeadersPolicy("example",
    comment="test comment",
    cors_config=aws.cloudfront.ResponseHeadersPolicyCorsConfigArgs(
        access_control_allow_credentials=True,
        access_control_allow_headers=aws.cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowHeadersArgs(
            items=["test"],
        ),
        access_control_allow_methods=aws.cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowMethodsArgs(
            items=["GET"],
        ),
        access_control_allow_origins=aws.cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowOriginsArgs(
            items=["test.example.comtest"],
        ),
        origin_override=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.ResponseHeadersPolicy("example", new()
    {
        Comment = "test comment",
        CorsConfig = new Aws.CloudFront.Inputs.ResponseHeadersPolicyCorsConfigArgs
        {
            AccessControlAllowCredentials = true,
            AccessControlAllowHeaders = new Aws.CloudFront.Inputs.ResponseHeadersPolicyCorsConfigAccessControlAllowHeadersArgs
            {
                Items = new[]
                {
                    "test",
                },
            },
            AccessControlAllowMethods = new Aws.CloudFront.Inputs.ResponseHeadersPolicyCorsConfigAccessControlAllowMethodsArgs
            {
                Items = new[]
                {
                    "GET",
                },
            },
            AccessControlAllowOrigins = new Aws.CloudFront.Inputs.ResponseHeadersPolicyCorsConfigAccessControlAllowOriginsArgs
            {
                Items = new[]
                {
                    "test.example.comtest",
                },
            },
            OriginOverride = true,
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
		_, err := cloudfront.NewResponseHeadersPolicy(ctx, "example", &cloudfront.ResponseHeadersPolicyArgs{
			Comment: pulumi.String("test comment"),
			CorsConfig: &cloudfront.ResponseHeadersPolicyCorsConfigArgs{
				AccessControlAllowCredentials: pulumi.Bool(true),
				AccessControlAllowHeaders: &cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowHeadersArgs{
					Items: pulumi.StringArray{
						pulumi.String("test"),
					},
				},
				AccessControlAllowMethods: &cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowMethodsArgs{
					Items: pulumi.StringArray{
						pulumi.String("GET"),
					},
				},
				AccessControlAllowOrigins: &cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowOriginsArgs{
					Items: pulumi.StringArray{
						pulumi.String("test.example.comtest"),
					},
				},
				OriginOverride: pulumi.Bool(true),
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
import com.pulumi.aws.cloudfront.ResponseHeadersPolicy;
import com.pulumi.aws.cloudfront.ResponseHeadersPolicyArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyCorsConfigArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyCorsConfigAccessControlAllowHeadersArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyCorsConfigAccessControlAllowMethodsArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyCorsConfigAccessControlAllowOriginsArgs;
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
        var example = new ResponseHeadersPolicy("example", ResponseHeadersPolicyArgs.builder()        
            .comment("test comment")
            .corsConfig(ResponseHeadersPolicyCorsConfigArgs.builder()
                .accessControlAllowCredentials(true)
                .accessControlAllowHeaders(ResponseHeadersPolicyCorsConfigAccessControlAllowHeadersArgs.builder()
                    .items("test")
                    .build())
                .accessControlAllowMethods(ResponseHeadersPolicyCorsConfigAccessControlAllowMethodsArgs.builder()
                    .items("GET")
                    .build())
                .accessControlAllowOrigins(ResponseHeadersPolicyCorsConfigAccessControlAllowOriginsArgs.builder()
                    .items("test.example.comtest")
                    .build())
                .originOverride(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:ResponseHeadersPolicy
    properties:
      comment: test comment
      corsConfig:
        accessControlAllowCredentials: true
        accessControlAllowHeaders:
          items:
            - test
        accessControlAllowMethods:
          items:
            - GET
        accessControlAllowOrigins:
          items:
            - test.example.comtest
        originOverride: true
```

The example below creates a CloudFront response headers policy with a custom headers config.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.ResponseHeadersPolicy("example", {customHeadersConfig: {
    items: [
        {
            header: "X-Permitted-Cross-Domain-Policies",
            override: true,
            value: "none",
        },
        {
            header: "X-Test",
            override: true,
            value: "none",
        },
    ],
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.ResponseHeadersPolicy("example", custom_headers_config=aws.cloudfront.ResponseHeadersPolicyCustomHeadersConfigArgs(
    items=[
        aws.cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs(
            header="X-Permitted-Cross-Domain-Policies",
            override=True,
            value="none",
        ),
        aws.cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs(
            header="X-Test",
            override=True,
            value="none",
        ),
    ],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.ResponseHeadersPolicy("example", new()
    {
        CustomHeadersConfig = new Aws.CloudFront.Inputs.ResponseHeadersPolicyCustomHeadersConfigArgs
        {
            Items = new[]
            {
                new Aws.CloudFront.Inputs.ResponseHeadersPolicyCustomHeadersConfigItemArgs
                {
                    Header = "X-Permitted-Cross-Domain-Policies",
                    Override = true,
                    Value = "none",
                },
                new Aws.CloudFront.Inputs.ResponseHeadersPolicyCustomHeadersConfigItemArgs
                {
                    Header = "X-Test",
                    Override = true,
                    Value = "none",
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
		_, err := cloudfront.NewResponseHeadersPolicy(ctx, "example", &cloudfront.ResponseHeadersPolicyArgs{
			CustomHeadersConfig: &cloudfront.ResponseHeadersPolicyCustomHeadersConfigArgs{
				Items: cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArray{
					&cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs{
						Header:   pulumi.String("X-Permitted-Cross-Domain-Policies"),
						Override: pulumi.Bool(true),
						Value:    pulumi.String("none"),
					},
					&cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs{
						Header:   pulumi.String("X-Test"),
						Override: pulumi.Bool(true),
						Value:    pulumi.String("none"),
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
import com.pulumi.aws.cloudfront.ResponseHeadersPolicy;
import com.pulumi.aws.cloudfront.ResponseHeadersPolicyArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyCustomHeadersConfigArgs;
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
        var example = new ResponseHeadersPolicy("example", ResponseHeadersPolicyArgs.builder()        
            .customHeadersConfig(ResponseHeadersPolicyCustomHeadersConfigArgs.builder()
                .items(                
                    ResponseHeadersPolicyCustomHeadersConfigItemArgs.builder()
                        .header("X-Permitted-Cross-Domain-Policies")
                        .override(true)
                        .value("none")
                        .build(),
                    ResponseHeadersPolicyCustomHeadersConfigItemArgs.builder()
                        .header("X-Test")
                        .override(true)
                        .value("none")
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:ResponseHeadersPolicy
    properties:
      customHeadersConfig:
        items:
          - header: X-Permitted-Cross-Domain-Policies
            override: true
            value: none
          - header: X-Test
            override: true
            value: none
```

The example below creates a CloudFront response headers policy with a custom headers config and server timing headers config.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.ResponseHeadersPolicy("example", {
    customHeadersConfig: {
        items: [{
            header: "X-Permitted-Cross-Domain-Policies",
            override: true,
            value: "none",
        }],
    },
    serverTimingHeadersConfig: {
        enabled: true,
        samplingRate: 50,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.ResponseHeadersPolicy("example",
    custom_headers_config=aws.cloudfront.ResponseHeadersPolicyCustomHeadersConfigArgs(
        items=[aws.cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs(
            header="X-Permitted-Cross-Domain-Policies",
            override=True,
            value="none",
        )],
    ),
    server_timing_headers_config=aws.cloudfront.ResponseHeadersPolicyServerTimingHeadersConfigArgs(
        enabled=True,
        sampling_rate=50,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.ResponseHeadersPolicy("example", new()
    {
        CustomHeadersConfig = new Aws.CloudFront.Inputs.ResponseHeadersPolicyCustomHeadersConfigArgs
        {
            Items = new[]
            {
                new Aws.CloudFront.Inputs.ResponseHeadersPolicyCustomHeadersConfigItemArgs
                {
                    Header = "X-Permitted-Cross-Domain-Policies",
                    Override = true,
                    Value = "none",
                },
            },
        },
        ServerTimingHeadersConfig = new Aws.CloudFront.Inputs.ResponseHeadersPolicyServerTimingHeadersConfigArgs
        {
            Enabled = true,
            SamplingRate = 50,
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
		_, err := cloudfront.NewResponseHeadersPolicy(ctx, "example", &cloudfront.ResponseHeadersPolicyArgs{
			CustomHeadersConfig: &cloudfront.ResponseHeadersPolicyCustomHeadersConfigArgs{
				Items: cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArray{
					&cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs{
						Header:   pulumi.String("X-Permitted-Cross-Domain-Policies"),
						Override: pulumi.Bool(true),
						Value:    pulumi.String("none"),
					},
				},
			},
			ServerTimingHeadersConfig: &cloudfront.ResponseHeadersPolicyServerTimingHeadersConfigArgs{
				Enabled:      pulumi.Bool(true),
				SamplingRate: pulumi.Float64(50),
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
import com.pulumi.aws.cloudfront.ResponseHeadersPolicy;
import com.pulumi.aws.cloudfront.ResponseHeadersPolicyArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyCustomHeadersConfigArgs;
import com.pulumi.aws.cloudfront.inputs.ResponseHeadersPolicyServerTimingHeadersConfigArgs;
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
        var example = new ResponseHeadersPolicy("example", ResponseHeadersPolicyArgs.builder()        
            .customHeadersConfig(ResponseHeadersPolicyCustomHeadersConfigArgs.builder()
                .items(ResponseHeadersPolicyCustomHeadersConfigItemArgs.builder()
                    .header("X-Permitted-Cross-Domain-Policies")
                    .override(true)
                    .value("none")
                    .build())
                .build())
            .serverTimingHeadersConfig(ResponseHeadersPolicyServerTimingHeadersConfigArgs.builder()
                .enabled(true)
                .samplingRate(50)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:ResponseHeadersPolicy
    properties:
      customHeadersConfig:
        items:
          - header: X-Permitted-Cross-Domain-Policies
            override: true
            value: none
      serverTimingHeadersConfig:
        enabled: true
        samplingRate: 50
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudfront Response Headers Policies using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy policy 658327ea-f89d-4fab-a63d-7e88639e58f9
```
 