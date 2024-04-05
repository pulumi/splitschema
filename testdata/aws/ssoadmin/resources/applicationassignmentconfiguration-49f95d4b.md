Resource for managing an AWS SSO Admin Application Assignment Configuration.

By default, applications will require users to have an explicit assignment in order to access an application.
This resource can be used to adjust this default behavior if necessary.

> Deleting this resource will return the assignment configuration for the application to the default AWS behavior (ie. `assignment_required = true`).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssoadmin.ApplicationAssignmentConfiguration("example", {
    applicationArn: aws_ssoadmin_application.example.application_arn,
    assignmentRequired: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssoadmin.ApplicationAssignmentConfiguration("example",
    application_arn=aws_ssoadmin_application["example"]["application_arn"],
    assignment_required=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsoAdmin.ApplicationAssignmentConfiguration("example", new()
    {
        ApplicationArn = aws_ssoadmin_application.Example.Application_arn,
        AssignmentRequired = true,
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
		_, err := ssoadmin.NewApplicationAssignmentConfiguration(ctx, "example", &ssoadmin.ApplicationAssignmentConfigurationArgs{
			ApplicationArn:     pulumi.Any(aws_ssoadmin_application.Example.Application_arn),
			AssignmentRequired: pulumi.Bool(true),
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
import com.pulumi.aws.ssoadmin.ApplicationAssignmentConfiguration;
import com.pulumi.aws.ssoadmin.ApplicationAssignmentConfigurationArgs;
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
        var example = new ApplicationAssignmentConfiguration("example", ApplicationAssignmentConfigurationArgs.builder()        
            .applicationArn(aws_ssoadmin_application.example().application_arn())
            .assignmentRequired(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssoadmin:ApplicationAssignmentConfiguration
    properties:
      applicationArn: ${aws_ssoadmin_application.example.application_arn}
      assignmentRequired: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSO Admin Application Assignment Configuration using the `id`. For example:

```sh
 $ pulumi import aws:ssoadmin/applicationAssignmentConfiguration:ApplicationAssignmentConfiguration example arn:aws:sso::012345678901:application/id-12345678
```
 