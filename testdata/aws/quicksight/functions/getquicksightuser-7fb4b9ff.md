This data source can be used to fetch information about a specific
QuickSight user. By using this data source, you can reference QuickSight user
properties without having to hard code ARNs or unique IDs as input.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.quicksight.getQuicksightUser({
    userName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.get_quicksight_user(user_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Quicksight.GetQuicksightUser.Invoke(new()
    {
        UserName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.GetQuicksightUser(ctx, &quicksight.GetQuicksightUserArgs{
			UserName: "example",
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
import com.pulumi.aws.quicksight.QuicksightFunctions;
import com.pulumi.aws.quicksight.inputs.GetQuicksightUserArgs;
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
        final var example = QuicksightFunctions.getQuicksightUser(GetQuicksightUserArgs.builder()
            .userName("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:quicksight:getQuicksightUser
      Arguments:
        userName: example
```
{{% /example %}}
{{% /examples %}}