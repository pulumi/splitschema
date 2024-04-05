`aws.route53.ResolverEndpoint` provides details about a specific Route53 Resolver Endpoint.

This data source allows to find a list of IPaddresses associated with a specific Route53 Resolver Endpoint.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverEndpoint({
    resolverEndpointId: "rslvr-in-1abc2345ef678g91h",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_endpoint(resolver_endpoint_id="rslvr-in-1abc2345ef678g91h")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverEndpoint.Invoke(new()
    {
        ResolverEndpointId = "rslvr-in-1abc2345ef678g91h",
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
		_, err := route53.LookupResolverEndpoint(ctx, &route53.LookupResolverEndpointArgs{
			ResolverEndpointId: pulumi.StringRef("rslvr-in-1abc2345ef678g91h"),
		}, nil)
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetResolverEndpointArgs;
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
        final var example = Route53Functions.getResolverEndpoint(GetResolverEndpointArgs.builder()
            .resolverEndpointId("rslvr-in-1abc2345ef678g91h")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverEndpoint
      Arguments:
        resolverEndpointId: rslvr-in-1abc2345ef678g91h
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverEndpoint({
    filters: [{
        name: "NAME",
        values: ["MyResolverExampleName"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_endpoint(filters=[aws.route53.GetResolverEndpointFilterArgs(
    name="NAME",
    values=["MyResolverExampleName"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverEndpoint.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Route53.Inputs.GetResolverEndpointFilterInputArgs
            {
                Name = "NAME",
                Values = new[]
                {
                    "MyResolverExampleName",
                },
            },
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
		_, err := route53.LookupResolverEndpoint(ctx, &route53.LookupResolverEndpointArgs{
			Filters: []route53.GetResolverEndpointFilter{
				{
					Name: "NAME",
					Values: []string{
						"MyResolverExampleName",
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetResolverEndpointArgs;
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
        final var example = Route53Functions.getResolverEndpoint(GetResolverEndpointArgs.builder()
            .filters(GetResolverEndpointFilterArgs.builder()
                .name("NAME")
                .values("MyResolverExampleName")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverEndpoint
      Arguments:
        filters:
          - name: NAME
            values:
              - MyResolverExampleName
```
{{% /example %}}
{{% /examples %}}