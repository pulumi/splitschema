Provides a resource to manage an S3 Access Grants instance resource policy.
Use a resource policy to manage cross-account access to your S3 Access Grants instance.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccessGrantsInstance = new aws.s3control.AccessGrantsInstance("exampleAccessGrantsInstance", {});
const exampleAccessGrantsInstanceResourcePolicy = new aws.s3control.AccessGrantsInstanceResourcePolicy("exampleAccessGrantsInstanceResourcePolicy", {policy: pulumi.interpolate`{
  "Version": "2012-10-17",
  "Id": "S3AccessGrantsPolicy",
  "Statement": [{
    "Sid": "AllowAccessToS3AccessGrants",
    "Effect": "Allow",
    "Principal": {
      "AWS": "123456789456"
    },
    "Action": [
      "s3:ListAccessGrants",
      "s3:ListAccessGrantsLocations",
      "s3:GetDataAccess"
    ],
    "Resource": "${exampleAccessGrantsInstance.accessGrantsInstanceArn}"
  }]
}

`});
```
```python
import pulumi
import pulumi_aws as aws

example_access_grants_instance = aws.s3control.AccessGrantsInstance("exampleAccessGrantsInstance")
example_access_grants_instance_resource_policy = aws.s3control.AccessGrantsInstanceResourcePolicy("exampleAccessGrantsInstanceResourcePolicy", policy=example_access_grants_instance.access_grants_instance_arn.apply(lambda access_grants_instance_arn: f"""{{
  "Version": "2012-10-17",
  "Id": "S3AccessGrantsPolicy",
  "Statement": [{{
    "Sid": "AllowAccessToS3AccessGrants",
    "Effect": "Allow",
    "Principal": {{
      "AWS": "123456789456"
    }},
    "Action": [
      "s3:ListAccessGrants",
      "s3:ListAccessGrantsLocations",
      "s3:GetDataAccess"
    ],
    "Resource": "{access_grants_instance_arn}"
  }}]
}}

"""))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccessGrantsInstance = new Aws.S3Control.AccessGrantsInstance("exampleAccessGrantsInstance");

    var exampleAccessGrantsInstanceResourcePolicy = new Aws.S3Control.AccessGrantsInstanceResourcePolicy("exampleAccessGrantsInstanceResourcePolicy", new()
    {
        Policy = exampleAccessGrantsInstance.AccessGrantsInstanceArn.Apply(accessGrantsInstanceArn => @$"{{
  ""Version"": ""2012-10-17"",
  ""Id"": ""S3AccessGrantsPolicy"",
  ""Statement"": [{{
    ""Sid"": ""AllowAccessToS3AccessGrants"",
    ""Effect"": ""Allow"",
    ""Principal"": {{
      ""AWS"": ""123456789456""
    }},
    ""Action"": [
      ""s3:ListAccessGrants"",
      ""s3:ListAccessGrantsLocations"",
      ""s3:GetDataAccess""
    ],
    ""Resource"": ""{accessGrantsInstanceArn}""
  }}]
}}

"),
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
		_, err = s3control.NewAccessGrantsInstanceResourcePolicy(ctx, "exampleAccessGrantsInstanceResourcePolicy", &s3control.AccessGrantsInstanceResourcePolicyArgs{
			Policy: exampleAccessGrantsInstance.AccessGrantsInstanceArn.ApplyT(func(accessGrantsInstanceArn string) (string, error) {
				return fmt.Sprintf(`{
  "Version": "2012-10-17",
  "Id": "S3AccessGrantsPolicy",
  "Statement": [{
    "Sid": "AllowAccessToS3AccessGrants",
    "Effect": "Allow",
    "Principal": {
      "AWS": "123456789456"
    },
    "Action": [
      "s3:ListAccessGrants",
      "s3:ListAccessGrantsLocations",
      "s3:GetDataAccess"
    ],
    "Resource": "%v"
  }]
}

`, accessGrantsInstanceArn), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.s3control.AccessGrantsInstanceResourcePolicy;
import com.pulumi.aws.s3control.AccessGrantsInstanceResourcePolicyArgs;
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

        var exampleAccessGrantsInstanceResourcePolicy = new AccessGrantsInstanceResourcePolicy("exampleAccessGrantsInstanceResourcePolicy", AccessGrantsInstanceResourcePolicyArgs.builder()        
            .policy(exampleAccessGrantsInstance.accessGrantsInstanceArn().applyValue(accessGrantsInstanceArn -> """
{
  "Version": "2012-10-17",
  "Id": "S3AccessGrantsPolicy",
  "Statement": [{
    "Sid": "AllowAccessToS3AccessGrants",
    "Effect": "Allow",
    "Principal": {
      "AWS": "123456789456"
    },
    "Action": [
      "s3:ListAccessGrants",
      "s3:ListAccessGrantsLocations",
      "s3:GetDataAccess"
    ],
    "Resource": "%s"
  }]
}

", accessGrantsInstanceArn)))
            .build());

    }
}
```
```yaml
resources:
  exampleAccessGrantsInstance:
    type: aws:s3control:AccessGrantsInstance
  exampleAccessGrantsInstanceResourcePolicy:
    type: aws:s3control:AccessGrantsInstanceResourcePolicy
    properties:
      policy: |+
        {
          "Version": "2012-10-17",
          "Id": "S3AccessGrantsPolicy",
          "Statement": [{
            "Sid": "AllowAccessToS3AccessGrants",
            "Effect": "Allow",
            "Principal": {
              "AWS": "123456789456"
            },
            "Action": [
              "s3:ListAccessGrants",
              "s3:ListAccessGrantsLocations",
              "s3:GetDataAccess"
            ],
            "Resource": "${exampleAccessGrantsInstance.accessGrantsInstanceArn}"
          }]
        }
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Access Grants instance resource policies using the `account_id`. For example:

```sh
 $ pulumi import aws:s3control/accessGrantsInstanceResourcePolicy:AccessGrantsInstanceResourcePolicy example 123456789012
```
 