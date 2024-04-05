Use this data source to get information on an AWS Cost and Usage Report Definition.

> *NOTE:* The AWS Cost and Usage Report service is only available in `us-east-1` currently.

> *NOTE:* If AWS Organizations is enabled, only the master account can use this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const reportDefinition = aws.cur.getReportDefinition({
    reportName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

report_definition = aws.cur.get_report_definition(report_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var reportDefinition = Aws.Cur.GetReportDefinition.Invoke(new()
    {
        ReportName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cur"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cur.LookupReportDefinition(ctx, &cur.LookupReportDefinitionArgs{
			ReportName: "example",
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
import com.pulumi.aws.cur.CurFunctions;
import com.pulumi.aws.cur.inputs.GetReportDefinitionArgs;
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
        final var reportDefinition = CurFunctions.getReportDefinition(GetReportDefinitionArgs.builder()
            .reportName("example")
            .build());

    }
}
```
```yaml
variables:
  reportDefinition:
    fn::invoke:
      Function: aws:cur:getReportDefinition
      Arguments:
        reportName: example
```
{{% /example %}}
{{% /examples %}}