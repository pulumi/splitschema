Provides details about a specific Amazon Kendra Faq.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.kendra.getFaq({
    faqId: "87654321-1234-4321-4321-321987654321",
    indexId: "12345678-1234-1234-1234-123456789123",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.kendra.get_faq(faq_id="87654321-1234-4321-4321-321987654321",
    index_id="12345678-1234-1234-1234-123456789123")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Kendra.GetFaq.Invoke(new()
    {
        FaqId = "87654321-1234-4321-4321-321987654321",
        IndexId = "12345678-1234-1234-1234-123456789123",
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
		_, err := kendra.LookupFaq(ctx, &kendra.LookupFaqArgs{
			FaqId:   "87654321-1234-4321-4321-321987654321",
			IndexId: "12345678-1234-1234-1234-123456789123",
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
import com.pulumi.aws.kendra.inputs.GetFaqArgs;
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
        final var test = KendraFunctions.getFaq(GetFaqArgs.builder()
            .faqId("87654321-1234-4321-4321-321987654321")
            .indexId("12345678-1234-1234-1234-123456789123")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:kendra:getFaq
      Arguments:
        faqId: 87654321-1234-4321-4321-321987654321
        indexId: 12345678-1234-1234-1234-123456789123
```
{{% /example %}}
{{% /examples %}}