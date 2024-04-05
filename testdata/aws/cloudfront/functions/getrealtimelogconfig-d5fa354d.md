Provides a CloudFront real-time log configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.cloudfront.getRealtimeLogConfig({
    name: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.get_realtime_log_config(name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CloudFront.GetRealtimeLogConfig.Invoke(new()
    {
        Name = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.LookupRealtimeLogConfig(ctx, &cloudfront.LookupRealtimeLogConfigArgs{
			Name: "example",
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
import com.pulumi.aws.cloudfront.CloudfrontFunctions;
import com.pulumi.aws.cloudfront.inputs.GetRealtimeLogConfigArgs;
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
        final var example = CloudfrontFunctions.getRealtimeLogConfig(GetRealtimeLogConfigArgs.builder()
            .name("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:cloudfront:getRealtimeLogConfig
      Arguments:
        name: example
```
{{% /example %}}
{{% /examples %}}