This data source can be used to fetch information about IAM access keys of a
specific IAM user.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.iam.getAccessKeys({
    user: "an_example_user_name",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.get_access_keys(user="an_example_user_name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iam.GetAccessKeys.Invoke(new()
    {
        User = "an_example_user_name",
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
		_, err := iam.GetAccessKeys(ctx, &iam.GetAccessKeysArgs{
			User: "an_example_user_name",
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
import com.pulumi.aws.iam.inputs.GetAccessKeysArgs;
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
        final var example = IamFunctions.getAccessKeys(GetAccessKeysArgs.builder()
            .user("an_example_user_name")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:iam:getAccessKeys
      Arguments:
        user: an_example_user_name
```
{{% /example %}}
{{% /examples %}}