The S3 account public access block data source returns account-level public access block configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.s3.getAccountPublicAccessBlock({});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.get_account_public_access_block()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.S3.GetAccountPublicAccessBlock.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.LookupAccountPublicAccessBlock(ctx, nil, nil)
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetAccountPublicAccessBlockArgs;
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
        final var example = S3Functions.getAccountPublicAccessBlock();

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:s3:getAccountPublicAccessBlock
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}