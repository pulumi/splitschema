Retrieve information about a Network Firewall resource policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getResourcePolicy({
    resourceArn: _var.resource_policy_arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_resource_policy(resource_arn=var["resource_policy_arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetResourcePolicy.Invoke(new()
    {
        ResourceArn = @var.Resource_policy_arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.LookupResourcePolicy(ctx, &networkfirewall.LookupResourcePolicyArgs{
			ResourceArn: _var.Resource_policy_arn,
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
import com.pulumi.aws.networkfirewall.NetworkfirewallFunctions;
import com.pulumi.aws.networkfirewall.inputs.GetResourcePolicyArgs;
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
        final var example = NetworkfirewallFunctions.getResourcePolicy(GetResourcePolicyArgs.builder()
            .resourceArn(var_.resource_policy_arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getResourcePolicy
      Arguments:
        resourceArn: ${var.resource_policy_arn}
```
{{% /example %}}
{{% /examples %}}