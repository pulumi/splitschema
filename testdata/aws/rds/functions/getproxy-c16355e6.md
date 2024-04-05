Use this data source to get information about a DB Proxy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const proxy = aws.rds.getProxy({
    name: "my-test-db-proxy",
});
```
```python
import pulumi
import pulumi_aws as aws

proxy = aws.rds.get_proxy(name="my-test-db-proxy")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var proxy = Aws.Rds.GetProxy.Invoke(new()
    {
        Name = "my-test-db-proxy",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.LookupProxy(ctx, &rds.LookupProxyArgs{
			Name: "my-test-db-proxy",
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetProxyArgs;
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
        final var proxy = RdsFunctions.getProxy(GetProxyArgs.builder()
            .name("my-test-db-proxy")
            .build());

    }
}
```
```yaml
variables:
  proxy:
    fn::invoke:
      Function: aws:rds:getProxy
      Arguments:
        name: my-test-db-proxy
```
{{% /example %}}
{{% /examples %}}