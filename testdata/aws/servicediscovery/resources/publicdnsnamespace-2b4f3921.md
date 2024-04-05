Provides a Service Discovery Public DNS Namespace resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicediscovery.PublicDnsNamespace("example", {description: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicediscovery.PublicDnsNamespace("example", description="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceDiscovery.PublicDnsNamespace("example", new()
    {
        Description = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicediscovery"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicediscovery.NewPublicDnsNamespace(ctx, "example", &servicediscovery.PublicDnsNamespaceArgs{
			Description: pulumi.String("example"),
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
import com.pulumi.aws.servicediscovery.PublicDnsNamespace;
import com.pulumi.aws.servicediscovery.PublicDnsNamespaceArgs;
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
        var example = new PublicDnsNamespace("example", PublicDnsNamespaceArgs.builder()        
            .description("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicediscovery:PublicDnsNamespace
    properties:
      description: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Discovery Public DNS Namespace using the namespace ID. For example:

```sh
 $ pulumi import aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace example 0123456789
```
 