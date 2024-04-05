Provides details about a specific Amazon Kendra block list used for query suggestions for an index.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.kendra.getQuerySuggestionsBlockList({
    indexId: "12345678-1234-1234-1234-123456789123",
    querySuggestionsBlockListId: "87654321-1234-4321-4321-321987654321",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.get_query_suggestions_block_list(index_id="12345678-1234-1234-1234-123456789123",
    query_suggestions_block_list_id="87654321-1234-4321-4321-321987654321")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Kendra.GetQuerySuggestionsBlockList.Invoke(new()
    {
        IndexId = "12345678-1234-1234-1234-123456789123",
        QuerySuggestionsBlockListId = "87654321-1234-4321-4321-321987654321",
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
		_, err := kendra.LookupQuerySuggestionsBlockList(ctx, &kendra.LookupQuerySuggestionsBlockListArgs{
			IndexId:                     "12345678-1234-1234-1234-123456789123",
			QuerySuggestionsBlockListId: "87654321-1234-4321-4321-321987654321",
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
import com.pulumi.aws.kendra.inputs.GetQuerySuggestionsBlockListArgs;
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
        final var example = KendraFunctions.getQuerySuggestionsBlockList(GetQuerySuggestionsBlockListArgs.builder()
            .indexId("12345678-1234-1234-1234-123456789123")
            .querySuggestionsBlockListId("87654321-1234-4321-4321-321987654321")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:kendra:getQuerySuggestionsBlockList
      Arguments:
        indexId: 12345678-1234-1234-1234-123456789123
        querySuggestionsBlockListId: 87654321-1234-4321-4321-321987654321
```
{{% /example %}}
{{% /examples %}}