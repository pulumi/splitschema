Provides an IAM Service Specific Credential.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleUser = new aws.iam.User("exampleUser", {});
const exampleServiceSpecificCredential = new aws.iam.ServiceSpecificCredential("exampleServiceSpecificCredential", {
    serviceName: "codecommit.amazonaws.com",
    userName: exampleUser.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example_user = aws.iam.User("exampleUser")
example_service_specific_credential = aws.iam.ServiceSpecificCredential("exampleServiceSpecificCredential",
    service_name="codecommit.amazonaws.com",
    user_name=example_user.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleUser = new Aws.Iam.User("exampleUser");

    var exampleServiceSpecificCredential = new Aws.Iam.ServiceSpecificCredential("exampleServiceSpecificCredential", new()
    {
        ServiceName = "codecommit.amazonaws.com",
        UserName = exampleUser.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUser, err := iam.NewUser(ctx, "exampleUser", nil)
		if err != nil {
			return err
		}
		_, err = iam.NewServiceSpecificCredential(ctx, "exampleServiceSpecificCredential", &iam.ServiceSpecificCredentialArgs{
			ServiceName: pulumi.String("codecommit.amazonaws.com"),
			UserName:    exampleUser.Name,
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
import com.pulumi.aws.iam.User;
import com.pulumi.aws.iam.ServiceSpecificCredential;
import com.pulumi.aws.iam.ServiceSpecificCredentialArgs;
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
        var exampleUser = new User("exampleUser");

        var exampleServiceSpecificCredential = new ServiceSpecificCredential("exampleServiceSpecificCredential", ServiceSpecificCredentialArgs.builder()        
            .serviceName("codecommit.amazonaws.com")
            .userName(exampleUser.name())
            .build());

    }
}
```
```yaml
resources:
  exampleUser:
    type: aws:iam:User
  exampleServiceSpecificCredential:
    type: aws:iam:ServiceSpecificCredential
    properties:
      serviceName: codecommit.amazonaws.com
      userName: ${exampleUser.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Service Specific Credentials using the `service_name:user_name:service_specific_credential_id`. For example:

```sh
 $ pulumi import aws:iam/serviceSpecificCredential:ServiceSpecificCredential default `codecommit.amazonaws.com:example:some-id`
```
 