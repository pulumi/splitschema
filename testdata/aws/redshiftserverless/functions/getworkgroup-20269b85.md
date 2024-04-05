Data source for managing an AWS Redshift Serverless Workgroup.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.redshiftserverless.getWorkgroup({
    workgroupName: aws_redshiftserverless_workgroup.example.workgroup_name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshiftserverless.get_workgroup(workgroup_name=aws_redshiftserverless_workgroup["example"]["workgroup_name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.RedshiftServerless.GetWorkgroup.Invoke(new()
    {
        WorkgroupName = aws_redshiftserverless_workgroup.Example.Workgroup_name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshiftserverless"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshiftserverless.LookupWorkgroup(ctx, &redshiftserverless.LookupWorkgroupArgs{
			WorkgroupName: aws_redshiftserverless_workgroup.Example.Workgroup_name,
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
import com.pulumi.aws.redshiftserverless.RedshiftserverlessFunctions;
import com.pulumi.aws.redshiftserverless.inputs.GetWorkgroupArgs;
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
        final var example = RedshiftserverlessFunctions.getWorkgroup(GetWorkgroupArgs.builder()
            .workgroupName(aws_redshiftserverless_workgroup.example().workgroup_name())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:redshiftserverless:getWorkgroup
      Arguments:
        workgroupName: ${aws_redshiftserverless_workgroup.example.workgroup_name}
```
{{% /example %}}
{{% /examples %}}