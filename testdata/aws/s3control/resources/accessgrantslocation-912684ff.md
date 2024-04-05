Provides a resource to manage an S3 Access Grants location.
A location is an S3 resource (bucket or prefix) in a permission grant that the grantee can access.
The S3 data must be in the same Region as your S3 Access Grants instance.
When you register a location, you must include the IAM role that has permission to manage the S3 location that you are registering.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccessGrantsInstance = new aws.s3control.AccessGrantsInstance("exampleAccessGrantsInstance", {});
const exampleAccessGrantsLocation = new aws.s3control.AccessGrantsLocation("exampleAccessGrantsLocation", {
    iamRoleArn: aws_iam_role.example.arn,
    locationScope: "s3://",
}, {
    dependsOn: [exampleAccessGrantsInstance],
});
// Default scope.
```
```python
import pulumi
import pulumi_aws as aws

example_access_grants_instance = aws.s3control.AccessGrantsInstance("exampleAccessGrantsInstance")
example_access_grants_location = aws.s3control.AccessGrantsLocation("exampleAccessGrantsLocation",
    iam_role_arn=aws_iam_role["example"]["arn"],
    location_scope="s3://",
    opts=pulumi.ResourceOptions(depends_on=[example_access_grants_instance]))
# Default scope.
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccessGrantsInstance = new Aws.S3Control.AccessGrantsInstance("exampleAccessGrantsInstance");

    var exampleAccessGrantsLocation = new Aws.S3Control.AccessGrantsLocation("exampleAccessGrantsLocation", new()
    {
        IamRoleArn = aws_iam_role.Example.Arn,
        LocationScope = "s3://",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccessGrantsInstance,
        },
    });

    // Default scope.
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccessGrantsInstance, err := s3control.NewAccessGrantsInstance(ctx, "exampleAccessGrantsInstance", nil)
		if err != nil {
			return err
		}
		_, err = s3control.NewAccessGrantsLocation(ctx, "exampleAccessGrantsLocation", &s3control.AccessGrantsLocationArgs{
			IamRoleArn:    pulumi.Any(aws_iam_role.Example.Arn),
			LocationScope: pulumi.String("s3://"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccessGrantsInstance,
		}))
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
import com.pulumi.aws.s3control.AccessGrantsInstance;
import com.pulumi.aws.s3control.AccessGrantsLocation;
import com.pulumi.aws.s3control.AccessGrantsLocationArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccessGrantsInstance = new AccessGrantsInstance("exampleAccessGrantsInstance");

        var exampleAccessGrantsLocation = new AccessGrantsLocation("exampleAccessGrantsLocation", AccessGrantsLocationArgs.builder()        
            .iamRoleArn(aws_iam_role.example().arn())
            .locationScope("s3://")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccessGrantsInstance)
                .build());

    }
}
```
```yaml
resources:
  exampleAccessGrantsInstance:
    type: aws:s3control:AccessGrantsInstance
  exampleAccessGrantsLocation:
    type: aws:s3control:AccessGrantsLocation
    properties:
      iamRoleArn: ${aws_iam_role.example.arn}
      locationScope: s3://
    options:
      dependson:
        - ${exampleAccessGrantsInstance}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Access Grants locations using the `account_id` and `access_grants_location_id`, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:s3control/accessGrantsLocation:AccessGrantsLocation example 123456789012,default
```
 