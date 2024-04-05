Provides a Route 53 Resolver query logging configuration association resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.ResolverQueryLogConfigAssociation("example", {
    resolverQueryLogConfigId: aws_route53_resolver_query_log_config.example.id,
    resourceId: aws_vpc.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.ResolverQueryLogConfigAssociation("example",
    resolver_query_log_config_id=aws_route53_resolver_query_log_config["example"]["id"],
    resource_id=aws_vpc["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53.ResolverQueryLogConfigAssociation("example", new()
    {
        ResolverQueryLogConfigId = aws_route53_resolver_query_log_config.Example.Id,
        ResourceId = aws_vpc.Example.Id,
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
		_, err := route53.NewResolverQueryLogConfigAssociation(ctx, "example", &route53.ResolverQueryLogConfigAssociationArgs{
			ResolverQueryLogConfigId: pulumi.Any(aws_route53_resolver_query_log_config.Example.Id),
			ResourceId:               pulumi.Any(aws_vpc.Example.Id),
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
import com.pulumi.aws.route53.ResolverQueryLogConfigAssociation;
import com.pulumi.aws.route53.ResolverQueryLogConfigAssociationArgs;
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
        var example = new ResolverQueryLogConfigAssociation("example", ResolverQueryLogConfigAssociationArgs.builder()        
            .resolverQueryLogConfigId(aws_route53_resolver_query_log_config.example().id())
            .resourceId(aws_vpc.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53:ResolverQueryLogConfigAssociation
    properties:
      resolverQueryLogConfigId: ${aws_route53_resolver_query_log_config.example.id}
      resourceId: ${aws_vpc.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import

Route 53 Resolver query logging configuration associations using the Route 53 Resolver query logging configuration association ID. For example:

```sh
 $ pulumi import aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation example rqlca-b320624fef3c4d70
```
 