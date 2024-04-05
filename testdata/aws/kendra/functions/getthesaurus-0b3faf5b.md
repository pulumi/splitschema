Provides details about a specific Amazon Kendra Thesaurus.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.kendra.getThesaurus({
    indexId: "12345678-1234-1234-1234-123456789123",
    thesaurusId: "87654321-1234-4321-4321-321987654321",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.get_thesaurus(index_id="12345678-1234-1234-1234-123456789123",
    thesaurus_id="87654321-1234-4321-4321-321987654321")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Kendra.GetThesaurus.Invoke(new()
    {
        IndexId = "12345678-1234-1234-1234-123456789123",
        ThesaurusId = "87654321-1234-4321-4321-321987654321",
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
		_, err := kendra.LookupThesaurus(ctx, &kendra.LookupThesaurusArgs{
			IndexId:     "12345678-1234-1234-1234-123456789123",
			ThesaurusId: "87654321-1234-4321-4321-321987654321",
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
import com.pulumi.aws.kendra.inputs.GetThesaurusArgs;
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
        final var example = KendraFunctions.getThesaurus(GetThesaurusArgs.builder()
            .indexId("12345678-1234-1234-1234-123456789123")
            .thesaurusId("87654321-1234-4321-4321-321987654321")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:kendra:getThesaurus
      Arguments:
        indexId: 12345678-1234-1234-1234-123456789123
        thesaurusId: 87654321-1234-4321-4321-321987654321
```
{{% /example %}}
{{% /examples %}}