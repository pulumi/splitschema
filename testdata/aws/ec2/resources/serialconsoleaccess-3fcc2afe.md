Provides a resource to manage whether serial console access is enabled for your AWS account in the current AWS region.

> **NOTE:** Removing this resource disables serial console access.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SerialConsoleAccess("example", {enabled: true});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SerialConsoleAccess("example", enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SerialConsoleAccess("example", new()
    {
        Enabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewSerialConsoleAccess(ctx, "example", &ec2.SerialConsoleAccessArgs{
			Enabled: pulumi.Bool(true),
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
import com.pulumi.aws.ec2.SerialConsoleAccess;
import com.pulumi.aws.ec2.SerialConsoleAccessArgs;
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
        var example = new SerialConsoleAccess("example", SerialConsoleAccessArgs.builder()        
            .enabled(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SerialConsoleAccess
    properties:
      enabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import serial console access state. For example:

```sh
 $ pulumi import aws:ec2/serialConsoleAccess:SerialConsoleAccess example default
```
 