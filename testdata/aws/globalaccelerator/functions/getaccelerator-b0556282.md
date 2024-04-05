Provides information about a Global Accelerator accelerator.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const acceleratorArn = config.get("acceleratorArn") || "";
const acceleratorName = config.get("acceleratorName") || "";
const example = aws.globalaccelerator.getAccelerator({
    arn: acceleratorArn,
    name: acceleratorName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
accelerator_arn = config.get("acceleratorArn")
if accelerator_arn is None:
    accelerator_arn = ""
accelerator_name = config.get("acceleratorName")
if accelerator_name is None:
    accelerator_name = ""
example = aws.globalaccelerator.get_accelerator(arn=accelerator_arn,
    name=accelerator_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var acceleratorArn = config.Get("acceleratorArn") ?? "";
    var acceleratorName = config.Get("acceleratorName") ?? "";
    var example = Aws.GlobalAccelerator.GetAccelerator.Invoke(new()
    {
        Arn = acceleratorArn,
        Name = acceleratorName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/globalaccelerator"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		acceleratorArn := ""
		if param := cfg.Get("acceleratorArn"); param != "" {
			acceleratorArn = param
		}
		acceleratorName := ""
		if param := cfg.Get("acceleratorName"); param != "" {
			acceleratorName = param
		}
		_, err := globalaccelerator.LookupAccelerator(ctx, &globalaccelerator.LookupAcceleratorArgs{
			Arn:  pulumi.StringRef(acceleratorArn),
			Name: pulumi.StringRef(acceleratorName),
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
import com.pulumi.aws.globalaccelerator.GlobalacceleratorFunctions;
import com.pulumi.aws.globalaccelerator.inputs.GetAcceleratorArgs;
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
        final var config = ctx.config();
        final var acceleratorArn = config.get("acceleratorArn").orElse("");
        final var acceleratorName = config.get("acceleratorName").orElse("");
        final var example = GlobalacceleratorFunctions.getAccelerator(GetAcceleratorArgs.builder()
            .arn(acceleratorArn)
            .name(acceleratorName)
            .build());

    }
}
```
```yaml
configuration:
  acceleratorArn:
    type: string
    default:
  acceleratorName:
    type: string
    default:
variables:
  example:
    fn::invoke:
      Function: aws:globalaccelerator:getAccelerator
      Arguments:
        arn: ${acceleratorArn}
        name: ${acceleratorName}
```
{{% /example %}}
{{% /examples %}}