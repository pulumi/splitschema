Creates a new Amazon Redshift Serverless Usage Limit.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleWorkgroup = new aws.redshiftserverless.Workgroup("exampleWorkgroup", {
    namespaceName: aws_redshiftserverless_namespace.example.namespace_name,
    workgroupName: "example",
});
const exampleUsageLimit = new aws.redshiftserverless.UsageLimit("exampleUsageLimit", {
    resourceArn: exampleWorkgroup.arn,
    usageType: "serverless-compute",
    amount: 60,
});
```
```python
import pulumi
import pulumi_aws as aws

example_workgroup = aws.redshiftserverless.Workgroup("exampleWorkgroup",
    namespace_name=aws_redshiftserverless_namespace["example"]["namespace_name"],
    workgroup_name="example")
example_usage_limit = aws.redshiftserverless.UsageLimit("exampleUsageLimit",
    resource_arn=example_workgroup.arn,
    usage_type="serverless-compute",
    amount=60)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleWorkgroup = new Aws.RedshiftServerless.Workgroup("exampleWorkgroup", new()
    {
        NamespaceName = aws_redshiftserverless_namespace.Example.Namespace_name,
        WorkgroupName = "example",
    });

    var exampleUsageLimit = new Aws.RedshiftServerless.UsageLimit("exampleUsageLimit", new()
    {
        ResourceArn = exampleWorkgroup.Arn,
        UsageType = "serverless-compute",
        Amount = 60,
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
		exampleWorkgroup, err := redshiftserverless.NewWorkgroup(ctx, "exampleWorkgroup", &redshiftserverless.WorkgroupArgs{
			NamespaceName: pulumi.Any(aws_redshiftserverless_namespace.Example.Namespace_name),
			WorkgroupName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = redshiftserverless.NewUsageLimit(ctx, "exampleUsageLimit", &redshiftserverless.UsageLimitArgs{
			ResourceArn: exampleWorkgroup.Arn,
			UsageType:   pulumi.String("serverless-compute"),
			Amount:      pulumi.Int(60),
		})
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
import com.pulumi.aws.redshiftserverless.Workgroup;
import com.pulumi.aws.redshiftserverless.WorkgroupArgs;
import com.pulumi.aws.redshiftserverless.UsageLimit;
import com.pulumi.aws.redshiftserverless.UsageLimitArgs;
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
        var exampleWorkgroup = new Workgroup("exampleWorkgroup", WorkgroupArgs.builder()        
            .namespaceName(aws_redshiftserverless_namespace.example().namespace_name())
            .workgroupName("example")
            .build());

        var exampleUsageLimit = new UsageLimit("exampleUsageLimit", UsageLimitArgs.builder()        
            .resourceArn(exampleWorkgroup.arn())
            .usageType("serverless-compute")
            .amount(60)
            .build());

    }
}
```
```yaml
resources:
  exampleWorkgroup:
    type: aws:redshiftserverless:Workgroup
    properties:
      namespaceName: ${aws_redshiftserverless_namespace.example.namespace_name}
      workgroupName: example
  exampleUsageLimit:
    type: aws:redshiftserverless:UsageLimit
    properties:
      resourceArn: ${exampleWorkgroup.arn}
      usageType: serverless-compute
      amount: 60
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Serverless Usage Limits using the `id`. For example:

```sh
 $ pulumi import aws:redshiftserverless/usageLimit:UsageLimit example example-id
```
 