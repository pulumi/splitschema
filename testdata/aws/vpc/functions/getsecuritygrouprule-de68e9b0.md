`aws.vpc.getSecurityGroupRule` provides details about a specific security group rule.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.vpc.getSecurityGroupRule({
    securityGroupRuleId: _var.security_group_rule_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpc.get_security_group_rule(security_group_rule_id=var["security_group_rule_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Vpc.GetSecurityGroupRule.Invoke(new()
    {
        SecurityGroupRuleId = @var.Security_group_rule_id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpc"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpc.GetSecurityGroupRule(ctx, &vpc.GetSecurityGroupRuleArgs{
			SecurityGroupRuleId: pulumi.StringRef(_var.Security_group_rule_id),
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
import com.pulumi.aws.vpc.VpcFunctions;
import com.pulumi.aws.vpc.inputs.GetSecurityGroupRuleArgs;
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
        final var example = VpcFunctions.getSecurityGroupRule(GetSecurityGroupRuleArgs.builder()
            .securityGroupRuleId(var_.security_group_rule_id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:vpc:getSecurityGroupRule
      Arguments:
        securityGroupRuleId: ${var.security_group_rule_id}
```
{{% /example %}}
{{% /examples %}}