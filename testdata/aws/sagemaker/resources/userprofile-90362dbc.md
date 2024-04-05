Provides a SageMaker User Profile resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.UserProfile("example", {
    domainId: aws_sagemaker_domain.test.id,
    userProfileName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.UserProfile("example",
    domain_id=aws_sagemaker_domain["test"]["id"],
    user_profile_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.UserProfile("example", new()
    {
        DomainId = aws_sagemaker_domain.Test.Id,
        UserProfileName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewUserProfile(ctx, "example", &sagemaker.UserProfileArgs{
			DomainId:        pulumi.Any(aws_sagemaker_domain.Test.Id),
			UserProfileName: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.UserProfile;
import com.pulumi.aws.sagemaker.UserProfileArgs;
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
        var example = new UserProfile("example", UserProfileArgs.builder()        
            .domainId(aws_sagemaker_domain.test().id())
            .userProfileName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:UserProfile
    properties:
      domainId: ${aws_sagemaker_domain.test.id}
      userProfileName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker User Profiles using the `arn`. For example:

```sh
 $ pulumi import aws:sagemaker/userProfile:UserProfile test_user_profile arn:aws:sagemaker:us-west-2:123456789012:user-profile/domain-id/profile-name
```
 