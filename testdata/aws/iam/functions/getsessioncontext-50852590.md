This data source provides information on the IAM source role of an STS assumed role. For non-role ARNs, this data source simply passes the ARN through in `issuer_arn`.

For some AWS resources, multiple types of principals are allowed in the same argument (e.g., IAM users and IAM roles). However, these arguments often do not allow assumed-role (i.e., STS, temporary credential) principals. Given an STS ARN, this data source provides the ARN for the source IAM role.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.iam.getSessionContext({
    arn: "arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.get_session_context(arn="arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iam.GetSessionContext.Invoke(new()
    {
        Arn = "arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes",
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
		_, err := iam.GetSessionContext(ctx, &iam.GetSessionContextArgs{
			Arn: "arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes",
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
import com.pulumi.aws.iam.inputs.GetSessionContextArgs;
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
        final var example = IamFunctions.getSessionContext(GetSessionContextArgs.builder()
            .arn("arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:iam:getSessionContext
      Arguments:
        arn: arn:aws:sts::123456789012:assumed-role/Audien-Heaven/MatyNoyes
```
{{% /example %}}
{{% example %}}
### Find the Runner's Source Role

Combined with `aws.getCallerIdentity`, you can get the current user's source IAM role ARN (`issuer_arn`) if you're using an assumed role. If you're not using an assumed role, the caller's (e.g., an IAM user's) ARN will simply be passed through. In environments where both IAM users and individuals using assumed roles need to apply the same configurations, this data source enables seamless use.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const example = current.then(current => aws.iam.getSessionContext({
    arn: current.arn,
}));
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example = aws.iam.get_session_context(arn=current.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var example = Aws.Iam.GetSessionContext.Invoke(new()
    {
        Arn = current.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = iam.GetSessionContext(ctx, &iam.GetSessionContextArgs{
			Arn: current.Arn,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetSessionContextArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var example = IamFunctions.getSessionContext(GetSessionContextArgs.builder()
            .arn(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn()))
            .build());

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  example:
    fn::invoke:
      Function: aws:iam:getSessionContext
      Arguments:
        arn: ${current.arn}
```
{{% /example %}}
{{% /examples %}}