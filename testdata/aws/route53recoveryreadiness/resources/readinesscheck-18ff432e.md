Provides an AWS Route 53 Recovery Readiness Readiness Check.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53recoveryreadiness.ReadinessCheck("example", {
    readinessCheckName: my_cw_alarm_check,
    resourceSetName: my_cw_alarm_set,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53recoveryreadiness.ReadinessCheck("example",
    readiness_check_name=my_cw_alarm_check,
    resource_set_name=my_cw_alarm_set)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53RecoveryReadiness.ReadinessCheck("example", new()
    {
        ReadinessCheckName = my_cw_alarm_check,
        ResourceSetName = my_cw_alarm_set,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53recoveryreadiness"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53recoveryreadiness.NewReadinessCheck(ctx, "example", &route53recoveryreadiness.ReadinessCheckArgs{
			ReadinessCheckName: pulumi.Any(my_cw_alarm_check),
			ResourceSetName:    pulumi.Any(my_cw_alarm_set),
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
import com.pulumi.aws.route53recoveryreadiness.ReadinessCheck;
import com.pulumi.aws.route53recoveryreadiness.ReadinessCheckArgs;
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
        var example = new ReadinessCheck("example", ReadinessCheckArgs.builder()        
            .readinessCheckName(my_cw_alarm_check)
            .resourceSetName(my_cw_alarm_set)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53recoveryreadiness:ReadinessCheck
    properties:
      readinessCheckName: ${["my-cw-alarm-check"]}
      resourceSetName: ${["my-cw-alarm-set"]}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Recovery Readiness readiness checks using the readiness check name. For example:

```sh
 $ pulumi import aws:route53recoveryreadiness/readinessCheck:ReadinessCheck my-cw-alarm-check example
```
 