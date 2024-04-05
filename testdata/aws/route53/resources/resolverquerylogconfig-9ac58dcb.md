Provides a Route 53 Resolver query logging configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.ResolverQueryLogConfig("example", {
    destinationArn: aws_s3_bucket.example.arn,
    tags: {
        Environment: "Prod",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.ResolverQueryLogConfig("example",
    destination_arn=aws_s3_bucket["example"]["arn"],
    tags={
        "Environment": "Prod",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53.ResolverQueryLogConfig("example", new()
    {
        DestinationArn = aws_s3_bucket.Example.Arn,
        Tags = 
        {
            { "Environment", "Prod" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewResolverQueryLogConfig(ctx, "example", &route53.ResolverQueryLogConfigArgs{
			DestinationArn: pulumi.Any(aws_s3_bucket.Example.Arn),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Prod"),
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
import com.pulumi.aws.route53.ResolverQueryLogConfig;
import com.pulumi.aws.route53.ResolverQueryLogConfigArgs;
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
        var example = new ResolverQueryLogConfig("example", ResolverQueryLogConfigArgs.builder()        
            .destinationArn(aws_s3_bucket.example().arn())
            .tags(Map.of("Environment", "Prod"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53:ResolverQueryLogConfig
    properties:
      destinationArn: ${aws_s3_bucket.example.arn}
      tags:
        Environment: Prod
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import

Route 53 Resolver query logging configurations using the Route 53 Resolver query logging configuration ID. For example:

```sh
 $ pulumi import aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig example rqlc-92edc3b1838248bf
```
 