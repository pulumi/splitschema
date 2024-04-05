Provides information about a Launch Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = aws.ec2.getLaunchConfiguration({
    name: "test-launch-config",
});
```
```python
import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_launch_configuration(name="test-launch-config")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ubuntu = Aws.Ec2.GetLaunchConfiguration.Invoke(new()
    {
        Name = "test-launch-config",
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
		_, err := ec2.LookupLaunchConfiguration(ctx, &ec2.LookupLaunchConfigurationArgs{
			Name: "test-launch-config",
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetLaunchConfigurationArgs;
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
        final var ubuntu = Ec2Functions.getLaunchConfiguration(GetLaunchConfigurationArgs.builder()
            .name("test-launch-config")
            .build());

    }
}
```
```yaml
variables:
  ubuntu:
    fn::invoke:
      Function: aws:ec2:getLaunchConfiguration
      Arguments:
        name: test-launch-config
```
{{% /example %}}
{{% /examples %}}