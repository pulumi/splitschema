This data source can be used to fetch information about a specific
IAM policy.

{{% examples %}}
## Example Usage
{{% example %}}
### By ARN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.iam.getPolicy({
    arn: "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.get_policy(arn="arn:aws:iam::123456789012:policy/UsersManageOwnCredentials")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iam.GetPolicy.Invoke(new()
    {
        Arn = "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
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
		_, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
			Arn: pulumi.StringRef("arn:aws:iam::123456789012:policy/UsersManageOwnCredentials"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyArgs;
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
        final var example = IamFunctions.getPolicy(GetPolicyArgs.builder()
            .arn("arn:aws:iam::123456789012:policy/UsersManageOwnCredentials")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:iam:getPolicy
      Arguments:
        arn: arn:aws:iam::123456789012:policy/UsersManageOwnCredentials
```
{{% /example %}}
{{% example %}}
### By Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.iam.getPolicy({
    name: "test_policy",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.get_policy(name="test_policy")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iam.GetPolicy.Invoke(new()
    {
        Name = "test_policy",
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
		_, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
			Name: pulumi.StringRef("test_policy"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyArgs;
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
        final var example = IamFunctions.getPolicy(GetPolicyArgs.builder()
            .name("test_policy")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:iam:getPolicy
      Arguments:
        name: test_policy
```
{{% /example %}}
{{% /examples %}}