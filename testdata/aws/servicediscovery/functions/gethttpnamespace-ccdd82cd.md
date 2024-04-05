{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.servicediscovery.getHttpNamespace({
    name: "development",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicediscovery.get_http_namespace(name="development")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ServiceDiscovery.GetHttpNamespace.Invoke(new()
    {
        Name = "development",
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
		_, err := servicediscovery.LookupHttpNamespace(ctx, &servicediscovery.LookupHttpNamespaceArgs{
			Name: "development",
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
import com.pulumi.aws.servicediscovery.ServicediscoveryFunctions;
import com.pulumi.aws.servicediscovery.inputs.GetHttpNamespaceArgs;
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
        final var example = ServicediscoveryFunctions.getHttpNamespace(GetHttpNamespaceArgs.builder()
            .name("development")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:servicediscovery:getHttpNamespace
      Arguments:
        name: development
```
{{% /example %}}
{{% /examples %}}