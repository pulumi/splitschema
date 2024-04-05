Use this data source to get information about an AWS Cloudwatch Log Group

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.cloudwatch.getLogGroup({
    name: "MyImportantLogs",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudwatch.get_log_group(name="MyImportantLogs")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CloudWatch.GetLogGroup.Invoke(new()
    {
        Name = "MyImportantLogs",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.LookupLogGroup(ctx, &cloudwatch.LookupLogGroupArgs{
			Name: "MyImportantLogs",
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
import com.pulumi.aws.cloudwatch.CloudwatchFunctions;
import com.pulumi.aws.cloudwatch.inputs.GetLogGroupArgs;
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
        final var example = CloudwatchFunctions.getLogGroup(GetLogGroupArgs.builder()
            .name("MyImportantLogs")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:cloudwatch:getLogGroup
      Arguments:
        name: MyImportantLogs
```
{{% /example %}}
{{% /examples %}}