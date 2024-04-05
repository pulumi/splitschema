Provides an AWS Route 53 Recovery Readiness Recovery Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53recoveryreadiness.RecoveryGroup("example", {recoveryGroupName: "my-high-availability-app"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53recoveryreadiness.RecoveryGroup("example", recovery_group_name="my-high-availability-app")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53RecoveryReadiness.RecoveryGroup("example", new()
    {
        RecoveryGroupName = "my-high-availability-app",
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
		_, err := route53recoveryreadiness.NewRecoveryGroup(ctx, "example", &route53recoveryreadiness.RecoveryGroupArgs{
			RecoveryGroupName: pulumi.String("my-high-availability-app"),
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
import com.pulumi.aws.route53recoveryreadiness.RecoveryGroup;
import com.pulumi.aws.route53recoveryreadiness.RecoveryGroupArgs;
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
        var example = new RecoveryGroup("example", RecoveryGroupArgs.builder()        
            .recoveryGroupName("my-high-availability-app")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53recoveryreadiness:RecoveryGroup
    properties:
      recoveryGroupName: my-high-availability-app
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Recovery Readiness recovery groups using the recovery group name. For example:

```sh
 $ pulumi import aws:route53recoveryreadiness/recoveryGroup:RecoveryGroup my-high-availability-app my-high-availability-app
```
 