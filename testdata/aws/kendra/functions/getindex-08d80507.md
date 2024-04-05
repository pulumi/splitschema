Provides details about a specific Amazon Kendra Index.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.kendra.getIndex({
    id: "12345678-1234-1234-1234-123456789123",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.get_index(id="12345678-1234-1234-1234-123456789123")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Kendra.GetIndex.Invoke(new()
    {
        Id = "12345678-1234-1234-1234-123456789123",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.LookupIndex(ctx, &kendra.LookupIndexArgs{
			Id: "12345678-1234-1234-1234-123456789123",
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
import com.pulumi.aws.kendra.KendraFunctions;
import com.pulumi.aws.kendra.inputs.GetIndexArgs;
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
        final var example = KendraFunctions.getIndex(GetIndexArgs.builder()
            .id("12345678-1234-1234-1234-123456789123")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:kendra:getIndex
      Arguments:
        id: 12345678-1234-1234-1234-123456789123
```
{{% /example %}}
{{% /examples %}}