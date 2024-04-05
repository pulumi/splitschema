Provides an environment member to an AWS Cloud9 development environment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testEnvironmentEC2 = new aws.cloud9.EnvironmentEC2("testEnvironmentEC2", {instanceType: "t2.micro"});
const testUser = new aws.iam.User("testUser", {});
const testEnvironmentMembership = new aws.cloud9.EnvironmentMembership("testEnvironmentMembership", {
    environmentId: testEnvironmentEC2.id,
    permissions: "read-only",
    userArn: testUser.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test_environment_ec2 = aws.cloud9.EnvironmentEC2("testEnvironmentEC2", instance_type="t2.micro")
test_user = aws.iam.User("testUser")
test_environment_membership = aws.cloud9.EnvironmentMembership("testEnvironmentMembership",
    environment_id=test_environment_ec2.id,
    permissions="read-only",
    user_arn=test_user.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testEnvironmentEC2 = new Aws.Cloud9.EnvironmentEC2("testEnvironmentEC2", new()
    {
        InstanceType = "t2.micro",
    });

    var testUser = new Aws.Iam.User("testUser");

    var testEnvironmentMembership = new Aws.Cloud9.EnvironmentMembership("testEnvironmentMembership", new()
    {
        EnvironmentId = testEnvironmentEC2.Id,
        Permissions = "read-only",
        UserArn = testUser.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloud9"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testEnvironmentEC2, err := cloud9.NewEnvironmentEC2(ctx, "testEnvironmentEC2", &cloud9.EnvironmentEC2Args{
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		testUser, err := iam.NewUser(ctx, "testUser", nil)
		if err != nil {
			return err
		}
		_, err = cloud9.NewEnvironmentMembership(ctx, "testEnvironmentMembership", &cloud9.EnvironmentMembershipArgs{
			EnvironmentId: testEnvironmentEC2.ID(),
			Permissions:   pulumi.String("read-only"),
			UserArn:       testUser.Arn,
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
import com.pulumi.aws.cloud9.EnvironmentEC2;
import com.pulumi.aws.cloud9.EnvironmentEC2Args;
import com.pulumi.aws.iam.User;
import com.pulumi.aws.cloud9.EnvironmentMembership;
import com.pulumi.aws.cloud9.EnvironmentMembershipArgs;
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
        var testEnvironmentEC2 = new EnvironmentEC2("testEnvironmentEC2", EnvironmentEC2Args.builder()        
            .instanceType("t2.micro")
            .build());

        var testUser = new User("testUser");

        var testEnvironmentMembership = new EnvironmentMembership("testEnvironmentMembership", EnvironmentMembershipArgs.builder()        
            .environmentId(testEnvironmentEC2.id())
            .permissions("read-only")
            .userArn(testUser.arn())
            .build());

    }
}
```
```yaml
resources:
  testEnvironmentEC2:
    type: aws:cloud9:EnvironmentEC2
    properties:
      instanceType: t2.micro
  testUser:
    type: aws:iam:User
  testEnvironmentMembership:
    type: aws:cloud9:EnvironmentMembership
    properties:
      environmentId: ${testEnvironmentEC2.id}
      permissions: read-only
      userArn: ${testUser.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloud9 environment membership using the `environment-id#user-arn`. For example:

```sh
 $ pulumi import aws:cloud9/environmentMembership:EnvironmentMembership test environment-id#user-arn
```
 