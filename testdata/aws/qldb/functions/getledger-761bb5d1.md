Use this data source to fetch information about a Quantum Ledger Database.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.qldb.getLedger({
    name: "an_example_ledger",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.qldb.get_ledger(name="an_example_ledger")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Qldb.GetLedger.Invoke(new()
    {
        Name = "an_example_ledger",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/qldb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := qldb.LookupLedger(ctx, &qldb.LookupLedgerArgs{
			Name: "an_example_ledger",
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
import com.pulumi.aws.qldb.QldbFunctions;
import com.pulumi.aws.qldb.inputs.GetLedgerArgs;
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
        final var example = QldbFunctions.getLedger(GetLedgerArgs.builder()
            .name("an_example_ledger")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:qldb:getLedger
      Arguments:
        name: an_example_ledger
```
{{% /example %}}
{{% /examples %}}