Provides access to all Environments for an AppConfig Application. This will allow you to pass Environment IDs to another
resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.appconfig.getEnvironments({
    applicationId: "a1d3rpe",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appconfig.get_environments(application_id="a1d3rpe")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.AppConfig.GetEnvironments.Invoke(new()
    {
        ApplicationId = "a1d3rpe",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfig.GetEnvironments(ctx, &appconfig.GetEnvironmentsArgs{
			ApplicationId: "a1d3rpe",
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
import com.pulumi.aws.appconfig.AppconfigFunctions;
import com.pulumi.aws.appconfig.inputs.GetEnvironmentsArgs;
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
        final var example = AppconfigFunctions.getEnvironments(GetEnvironmentsArgs.builder()
            .applicationId("a1d3rpe")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:appconfig:getEnvironments
      Arguments:
        applicationId: a1d3rpe
```
{{% /example %}}
{{% /examples %}}