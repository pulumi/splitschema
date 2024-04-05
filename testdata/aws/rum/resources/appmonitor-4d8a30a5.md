Provides a CloudWatch RUM App Monitor resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rum.AppMonitor("example", {domain: "localhost"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rum.AppMonitor("example", domain="localhost")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rum.AppMonitor("example", new()
    {
        Domain = "localhost",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rum"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rum.NewAppMonitor(ctx, "example", &rum.AppMonitorArgs{
			Domain: pulumi.String("localhost"),
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
import com.pulumi.aws.rum.AppMonitor;
import com.pulumi.aws.rum.AppMonitorArgs;
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
        var example = new AppMonitor("example", AppMonitorArgs.builder()        
            .domain("localhost")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rum:AppMonitor
    properties:
      domain: localhost
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudwatch RUM App Monitor using the `name`. For example:

```sh
 $ pulumi import aws:rum/appMonitor:AppMonitor example example
```
 