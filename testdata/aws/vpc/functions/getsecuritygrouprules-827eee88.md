This resource can be useful for getting back a set of security group rule IDs.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.vpc.getSecurityGroupRules({
    filters: [{
        name: "group-id",
        values: [_var.security_group_id],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpc.get_security_group_rules(filters=[aws.vpc.GetSecurityGroupRulesFilterArgs(
    name="group-id",
    values=[var["security_group_id"]],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Vpc.GetSecurityGroupRules.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Vpc.Inputs.GetSecurityGroupRulesFilterInputArgs
            {
                Name = "group-id",
                Values = new[]
                {
                    @var.Security_group_id,
                },
            },
        },
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
_, err := vpc.GetSecurityGroupRules(ctx, &vpc.GetSecurityGroupRulesArgs{
Filters: []vpc.GetSecurityGroupRulesFilter{
{
Name: "group-id",
Values: interface{}{
_var.Security_group_id,
},
},
},
}, nil);
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
import com.pulumi.aws.vpc.inputs.GetSecurityGroupRulesArgs;
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
        final var example = VpcFunctions.getSecurityGroupRules(GetSecurityGroupRulesArgs.builder()
            .filters(GetSecurityGroupRulesFilterArgs.builder()
                .name("group-id")
                .values(var_.security_group_id())
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:vpc:getSecurityGroupRules
      Arguments:
        filters:
          - name: group-id
            values:
              - ${var.security_group_id}
```
{{% /example %}}
{{% /examples %}}