Data source for managing an AWS Audit Manager Framework.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.auditmanager.getFramework({
    frameworkType: "Standard",
    name: "Essential Eight",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.get_framework(framework_type="Standard",
    name="Essential Eight")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Auditmanager.GetFramework.Invoke(new()
    {
        FrameworkType = "Standard",
        Name = "Essential Eight",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.LookupFramework(ctx, &auditmanager.LookupFrameworkArgs{
			FrameworkType: "Standard",
			Name:          "Essential Eight",
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
import com.pulumi.aws.auditmanager.AuditmanagerFunctions;
import com.pulumi.aws.auditmanager.inputs.GetFrameworkArgs;
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
        final var example = AuditmanagerFunctions.getFramework(GetFrameworkArgs.builder()
            .frameworkType("Standard")
            .name("Essential Eight")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:auditmanager:getFramework
      Arguments:
        frameworkType: Standard
        name: Essential Eight
```
{{% /example %}}
{{% /examples %}}