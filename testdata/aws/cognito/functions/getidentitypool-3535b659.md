Data source for managing an AWS Cognito Identity Pool.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.cognito.getIdentityPool({
    identityPoolName: "test pool",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cognito.get_identity_pool(identity_pool_name="test pool")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Cognito.GetIdentityPool.Invoke(new()
    {
        IdentityPoolName = "test pool",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognito.LookupIdentityPool(ctx, &cognito.LookupIdentityPoolArgs{
			IdentityPoolName: "test pool",
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
import com.pulumi.aws.cognito.CognitoFunctions;
import com.pulumi.aws.cognito.inputs.GetIdentityPoolArgs;
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
        final var example = CognitoFunctions.getIdentityPool(GetIdentityPoolArgs.builder()
            .identityPoolName("test pool")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:cognito:getIdentityPool
      Arguments:
        identityPoolName: test pool
```
{{% /example %}}
{{% /examples %}}