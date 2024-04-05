Retrieves the summary of a WAFv2 Web ACL.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.wafv2.getWebAcl({
    name: "some-web-acl",
    scope: "REGIONAL",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafv2.get_web_acl(name="some-web-acl",
    scope="REGIONAL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.WafV2.GetWebAcl.Invoke(new()
    {
        Name = "some-web-acl",
        Scope = "REGIONAL",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.LookupWebAcl(ctx, &wafv2.LookupWebAclArgs{
			Name:  "some-web-acl",
			Scope: "REGIONAL",
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
import com.pulumi.aws.wafv2.Wafv2Functions;
import com.pulumi.aws.wafv2.inputs.GetWebAclArgs;
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
        final var example = Wafv2Functions.getWebAcl(GetWebAclArgs.builder()
            .name("some-web-acl")
            .scope("REGIONAL")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:wafv2:getWebAcl
      Arguments:
        name: some-web-acl
        scope: REGIONAL
```
{{% /example %}}
{{% /examples %}}