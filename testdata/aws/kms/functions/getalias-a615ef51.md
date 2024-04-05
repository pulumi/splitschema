Use this data source to get the ARN of a KMS key alias.
By using this data source, you can reference key alias
without having to hard code the ARN as input.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3 = aws.kms.getAlias({
    name: "alias/aws/s3",
});
```
```python
import pulumi
import pulumi_aws as aws

s3 = aws.kms.get_alias(name="alias/aws/s3")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3 = Aws.Kms.GetAlias.Invoke(new()
    {
        Name = "alias/aws/s3",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kms.LookupAlias(ctx, &kms.LookupAliasArgs{
			Name: "alias/aws/s3",
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
import com.pulumi.aws.kms.KmsFunctions;
import com.pulumi.aws.kms.inputs.GetAliasArgs;
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
        final var s3 = KmsFunctions.getAlias(GetAliasArgs.builder()
            .name("alias/aws/s3")
            .build());

    }
}
```
```yaml
variables:
  s3:
    fn::invoke:
      Function: aws:kms:getAlias
      Arguments:
        name: alias/aws/s3
```
{{% /example %}}
{{% /examples %}}