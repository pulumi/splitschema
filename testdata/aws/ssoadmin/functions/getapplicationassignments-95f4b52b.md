Data source for managing AWS SSO Admin Application Assignments.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ssoadmin.getApplicationAssignments({
    applicationArn: aws_ssoadmin_application.example.application_arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssoadmin.get_application_assignments(application_arn=aws_ssoadmin_application["example"]["application_arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SsoAdmin.GetApplicationAssignments.Invoke(new()
    {
        ApplicationArn = aws_ssoadmin_application.Example.Application_arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssoadmin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssoadmin.GetApplicationAssignments(ctx, &ssoadmin.GetApplicationAssignmentsArgs{
			ApplicationArn: aws_ssoadmin_application.Example.Application_arn,
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
import com.pulumi.aws.ssoadmin.SsoadminFunctions;
import com.pulumi.aws.ssoadmin.inputs.GetApplicationAssignmentsArgs;
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
        final var example = SsoadminFunctions.getApplicationAssignments(GetApplicationAssignmentsArgs.builder()
            .applicationArn(aws_ssoadmin_application.example().application_arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ssoadmin:getApplicationAssignments
      Arguments:
        applicationArn: ${aws_ssoadmin_application.example.application_arn}
```
{{% /example %}}
{{% /examples %}}