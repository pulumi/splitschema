Resource for managing an AWS SSO Admin Application Assignment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssoadmin.ApplicationAssignment("example", {
    applicationArn: aws_ssoadmin_application.example.application_arn,
    principalId: aws_identitystore_user.example.user_id,
    principalType: "USER",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssoadmin.ApplicationAssignment("example",
    application_arn=aws_ssoadmin_application["example"]["application_arn"],
    principal_id=aws_identitystore_user["example"]["user_id"],
    principal_type="USER")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsoAdmin.ApplicationAssignment("example", new()
    {
        ApplicationArn = aws_ssoadmin_application.Example.Application_arn,
        PrincipalId = aws_identitystore_user.Example.User_id,
        PrincipalType = "USER",
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
		_, err := ssoadmin.NewApplicationAssignment(ctx, "example", &ssoadmin.ApplicationAssignmentArgs{
			ApplicationArn: pulumi.Any(aws_ssoadmin_application.Example.Application_arn),
			PrincipalId:    pulumi.Any(aws_identitystore_user.Example.User_id),
			PrincipalType:  pulumi.String("USER"),
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
import com.pulumi.aws.ssoadmin.ApplicationAssignment;
import com.pulumi.aws.ssoadmin.ApplicationAssignmentArgs;
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
        var example = new ApplicationAssignment("example", ApplicationAssignmentArgs.builder()        
            .applicationArn(aws_ssoadmin_application.example().application_arn())
            .principalId(aws_identitystore_user.example().user_id())
            .principalType("USER")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssoadmin:ApplicationAssignment
    properties:
      applicationArn: ${aws_ssoadmin_application.example.application_arn}
      principalId: ${aws_identitystore_user.example.user_id}
      principalType: USER
```
{{% /example %}}
{{% example %}}
### Group Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssoadmin.ApplicationAssignment("example", {
    applicationArn: aws_ssoadmin_application.example.application_arn,
    principalId: aws_identitystore_group.example.group_id,
    principalType: "GROUP",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssoadmin.ApplicationAssignment("example",
    application_arn=aws_ssoadmin_application["example"]["application_arn"],
    principal_id=aws_identitystore_group["example"]["group_id"],
    principal_type="GROUP")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsoAdmin.ApplicationAssignment("example", new()
    {
        ApplicationArn = aws_ssoadmin_application.Example.Application_arn,
        PrincipalId = aws_identitystore_group.Example.Group_id,
        PrincipalType = "GROUP",
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
		_, err := ssoadmin.NewApplicationAssignment(ctx, "example", &ssoadmin.ApplicationAssignmentArgs{
			ApplicationArn: pulumi.Any(aws_ssoadmin_application.Example.Application_arn),
			PrincipalId:    pulumi.Any(aws_identitystore_group.Example.Group_id),
			PrincipalType:  pulumi.String("GROUP"),
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
import com.pulumi.aws.ssoadmin.ApplicationAssignment;
import com.pulumi.aws.ssoadmin.ApplicationAssignmentArgs;
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
        var example = new ApplicationAssignment("example", ApplicationAssignmentArgs.builder()        
            .applicationArn(aws_ssoadmin_application.example().application_arn())
            .principalId(aws_identitystore_group.example().group_id())
            .principalType("GROUP")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssoadmin:ApplicationAssignment
    properties:
      applicationArn: ${aws_ssoadmin_application.example.application_arn}
      principalId: ${aws_identitystore_group.example.group_id}
      principalType: GROUP
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSO Admin Application Assignment using the `id`. For example:

```sh
 $ pulumi import aws:ssoadmin/applicationAssignment:ApplicationAssignment example arn:aws:sso::012345678901:application/id-12345678,abcd1234,USER
```
 