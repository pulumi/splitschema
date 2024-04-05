Provides a resource to manage an S3 Access Grant.
Each access grant has its own ID and gives an IAM user or role or a directory user, or group (the grantee) access to a registered location. You determine the level of access, such as `READ` or `READWRITE`.
Before you can create a grant, you must have an S3 Access Grants instance in the same Region as the S3 data.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccessGrantsInstance = new aws.s3control.AccessGrantsInstance("exampleAccessGrantsInstance", {});
const exampleAccessGrantsLocation = new aws.s3control.AccessGrantsLocation("exampleAccessGrantsLocation", {
    iamRoleArn: aws_iam_role.example.arn,
    locationScope: `s3://${aws_s3_bucket.example.bucket}/prefixA*`,
}, {
    dependsOn: [exampleAccessGrantsInstance],
});
const exampleAccessGrant = new aws.s3control.AccessGrant("exampleAccessGrant", {
    accessGrantsLocationId: exampleAccessGrantsLocation.accessGrantsLocationId,
    permission: "READ",
    accessGrantsLocationConfiguration: {
        s3SubPrefix: "prefixB*",
    },
    grantee: {
        granteeType: "IAM",
        granteeIdentifier: aws_iam_user.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_access_grants_instance = aws.s3control.AccessGrantsInstance("exampleAccessGrantsInstance")
example_access_grants_location = aws.s3control.AccessGrantsLocation("exampleAccessGrantsLocation",
    iam_role_arn=aws_iam_role["example"]["arn"],
    location_scope=f"s3://{aws_s3_bucket['example']['bucket']}/prefixA*",
    opts=pulumi.ResourceOptions(depends_on=[example_access_grants_instance]))
example_access_grant = aws.s3control.AccessGrant("exampleAccessGrant",
    access_grants_location_id=example_access_grants_location.access_grants_location_id,
    permission="READ",
    access_grants_location_configuration=aws.s3control.AccessGrantAccessGrantsLocationConfigurationArgs(
        s3_sub_prefix="prefixB*",
    ),
    grantee=aws.s3control.AccessGrantGranteeArgs(
        grantee_type="IAM",
        grantee_identifier=aws_iam_user["example"]["arn"],
    ))
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
        LocationScope = $"s3://{aws_s3_bucket.Example.Bucket}/prefixA*",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccessGrantsInstance,
        },
    });

    var exampleAccessGrant = new Aws.S3Control.AccessGrant("exampleAccessGrant", new()
    {
        AccessGrantsLocationId = exampleAccessGrantsLocation.AccessGrantsLocationId,
        Permission = "READ",
        AccessGrantsLocationConfiguration = new Aws.S3Control.Inputs.AccessGrantAccessGrantsLocationConfigurationArgs
        {
            S3SubPrefix = "prefixB*",
        },
        Grantee = new Aws.S3Control.Inputs.AccessGrantGranteeArgs
        {
            GranteeType = "IAM",
            GranteeIdentifier = aws_iam_user.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccessGrantsInstance, err := s3control.NewAccessGrantsInstance(ctx, "exampleAccessGrantsInstance", nil)
		if err != nil {
			return err
		}
		exampleAccessGrantsLocation, err := s3control.NewAccessGrantsLocation(ctx, "exampleAccessGrantsLocation", &s3control.AccessGrantsLocationArgs{
			IamRoleArn:    pulumi.Any(aws_iam_role.Example.Arn),
			LocationScope: pulumi.String(fmt.Sprintf("s3://%v/prefixA*", aws_s3_bucket.Example.Bucket)),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccessGrantsInstance,
		}))
		if err != nil {
			return err
		}
		_, err = s3control.NewAccessGrant(ctx, "exampleAccessGrant", &s3control.AccessGrantArgs{
			AccessGrantsLocationId: exampleAccessGrantsLocation.AccessGrantsLocationId,
			Permission:             pulumi.String("READ"),
			AccessGrantsLocationConfiguration: &s3control.AccessGrantAccessGrantsLocationConfigurationArgs{
				S3SubPrefix: pulumi.String("prefixB*"),
			},
			Grantee: &s3control.AccessGrantGranteeArgs{
				GranteeType:       pulumi.String("IAM"),
				GranteeIdentifier: pulumi.Any(aws_iam_user.Example.Arn),
			},
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
import com.pulumi.aws.s3control.AccessGrantsInstance;
import com.pulumi.aws.s3control.AccessGrantsLocation;
import com.pulumi.aws.s3control.AccessGrantsLocationArgs;
import com.pulumi.aws.s3control.AccessGrant;
import com.pulumi.aws.s3control.AccessGrantArgs;
import com.pulumi.aws.s3control.inputs.AccessGrantAccessGrantsLocationConfigurationArgs;
import com.pulumi.aws.s3control.inputs.AccessGrantGranteeArgs;
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
            .locationScope(String.format("s3://%s/prefixA*", aws_s3_bucket.example().bucket()))
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccessGrantsInstance)
                .build());

        var exampleAccessGrant = new AccessGrant("exampleAccessGrant", AccessGrantArgs.builder()        
            .accessGrantsLocationId(exampleAccessGrantsLocation.accessGrantsLocationId())
            .permission("READ")
            .accessGrantsLocationConfiguration(AccessGrantAccessGrantsLocationConfigurationArgs.builder()
                .s3SubPrefix("prefixB*")
                .build())
            .grantee(AccessGrantGranteeArgs.builder()
                .granteeType("IAM")
                .granteeIdentifier(aws_iam_user.example().arn())
                .build())
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
      locationScope: s3://${aws_s3_bucket.example.bucket}/prefixA*
    options:
      dependson:
        - ${exampleAccessGrantsInstance}
  exampleAccessGrant:
    type: aws:s3control:AccessGrant
    properties:
      accessGrantsLocationId: ${exampleAccessGrantsLocation.accessGrantsLocationId}
      permission: READ
      accessGrantsLocationConfiguration:
        s3SubPrefix: prefixB*
      grantee:
        granteeType: IAM
        granteeIdentifier: ${aws_iam_user.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Access Grants using the `account_id` and `access_grant_id`, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:s3control/accessGrant:AccessGrant example 123456789012,04549c5e-2f3c-4a07-824d-2cafe720aa22
```
 