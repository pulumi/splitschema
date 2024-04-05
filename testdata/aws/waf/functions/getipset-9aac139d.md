`aws.waf.IpSet` Retrieves a WAF IP Set Resource Id.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.waf.getIpset({
    name: "tfWAFIPSet",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.waf.get_ipset(name="tfWAFIPSet")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Waf.GetIpset.Invoke(new()
    {
        Name = "tfWAFIPSet",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/waf"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := waf.GetIpset(ctx, &waf.GetIpsetArgs{
			Name: "tfWAFIPSet",
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
import com.pulumi.aws.waf.WafFunctions;
import com.pulumi.aws.waf.inputs.GetIpsetArgs;
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
        final var example = WafFunctions.getIpset(GetIpsetArgs.builder()
            .name("tfWAFIPSet")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:waf:getIpset
      Arguments:
        name: tfWAFIPSet
```
{{% /example %}}
{{% /examples %}}