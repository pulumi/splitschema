Manages an AppStream User Stack association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStack = new aws.appstream.Stack("testStack", {});
const testUser = new aws.appstream.User("testUser", {
    authenticationType: "USERPOOL",
    userName: "EMAIL",
});
const testUserStackAssociation = new aws.appstream.UserStackAssociation("testUserStackAssociation", {
    authenticationType: testUser.authenticationType,
    stackName: testStack.name,
    userName: testUser.userName,
});
```
```python
import pulumi
import pulumi_aws as aws

test_stack = aws.appstream.Stack("testStack")
test_user = aws.appstream.User("testUser",
    authentication_type="USERPOOL",
    user_name="EMAIL")
test_user_stack_association = aws.appstream.UserStackAssociation("testUserStackAssociation",
    authentication_type=test_user.authentication_type,
    stack_name=test_stack.name,
    user_name=test_user.user_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testStack = new Aws.AppStream.Stack("testStack");

    var testUser = new Aws.AppStream.User("testUser", new()
    {
        AuthenticationType = "USERPOOL",
        UserName = "EMAIL",
    });

    var testUserStackAssociation = new Aws.AppStream.UserStackAssociation("testUserStackAssociation", new()
    {
        AuthenticationType = testUser.AuthenticationType,
        StackName = testStack.Name,
        UserName = testUser.UserName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appstream"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testStack, err := appstream.NewStack(ctx, "testStack", nil)
		if err != nil {
			return err
		}
		testUser, err := appstream.NewUser(ctx, "testUser", &appstream.UserArgs{
			AuthenticationType: pulumi.String("USERPOOL"),
			UserName:           pulumi.String("EMAIL"),
		})
		if err != nil {
			return err
		}
		_, err = appstream.NewUserStackAssociation(ctx, "testUserStackAssociation", &appstream.UserStackAssociationArgs{
			AuthenticationType: testUser.AuthenticationType,
			StackName:          testStack.Name,
			UserName:           testUser.UserName,
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
import com.pulumi.aws.appstream.Stack;
import com.pulumi.aws.appstream.User;
import com.pulumi.aws.appstream.UserArgs;
import com.pulumi.aws.appstream.UserStackAssociation;
import com.pulumi.aws.appstream.UserStackAssociationArgs;
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
        var testStack = new Stack("testStack");

        var testUser = new User("testUser", UserArgs.builder()        
            .authenticationType("USERPOOL")
            .userName("EMAIL")
            .build());

        var testUserStackAssociation = new UserStackAssociation("testUserStackAssociation", UserStackAssociationArgs.builder()        
            .authenticationType(testUser.authenticationType())
            .stackName(testStack.name())
            .userName(testUser.userName())
            .build());

    }
}
```
```yaml
resources:
  testStack:
    type: aws:appstream:Stack
  testUser:
    type: aws:appstream:User
    properties:
      authenticationType: USERPOOL
      userName: EMAIL
  testUserStackAssociation:
    type: aws:appstream:UserStackAssociation
    properties:
      authenticationType: ${testUser.authenticationType}
      stackName: ${testStack.name}
      userName: ${testUser.userName}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppStream User Stack Association using the `user_name`, `authentication_type`, and `stack_name`, separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:appstream/userStackAssociation:UserStackAssociation example userName/auhtenticationType/stackName
```
 