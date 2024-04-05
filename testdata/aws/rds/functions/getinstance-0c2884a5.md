Use this data source to get information about an RDS instance

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const database = aws.rds.getInstance({
    dbInstanceIdentifier: "my-test-database",
});
```
```python
import pulumi
import pulumi_aws as aws

database = aws.rds.get_instance(db_instance_identifier="my-test-database")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var database = Aws.Rds.GetInstance.Invoke(new()
    {
        DbInstanceIdentifier = "my-test-database",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.LookupInstance(ctx, &rds.LookupInstanceArgs{
			DbInstanceIdentifier: pulumi.StringRef("my-test-database"),
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetInstanceArgs;
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
        final var database = RdsFunctions.getInstance(GetInstanceArgs.builder()
            .dbInstanceIdentifier("my-test-database")
            .build());

    }
}
```
```yaml
variables:
  database:
    fn::invoke:
      Function: aws:rds:getInstance
      Arguments:
        dbInstanceIdentifier: my-test-database
```
{{% /example %}}
{{% /examples %}}