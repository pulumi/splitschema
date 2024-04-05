Provides a CloudWatch Log Group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const yada = new aws.cloudwatch.LogGroup("yada", {tags: {
    Application: "serviceA",
    Environment: "production",
}});
```
```python
import pulumi
import pulumi_aws as aws

yada = aws.cloudwatch.LogGroup("yada", tags={
    "Application": "serviceA",
    "Environment": "production",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var yada = new Aws.CloudWatch.LogGroup("yada", new()
    {
        Tags = 
        {
            { "Application", "serviceA" },
            { "Environment", "production" },
        },
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
		_, err := cloudwatch.NewLogGroup(ctx, "yada", &cloudwatch.LogGroupArgs{
			Tags: pulumi.StringMap{
				"Application": pulumi.String("serviceA"),
				"Environment": pulumi.String("production"),
			},
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogGroupArgs;
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
        var yada = new LogGroup("yada", LogGroupArgs.builder()        
            .tags(Map.ofEntries(
                Map.entry("Application", "serviceA"),
                Map.entry("Environment", "production")
            ))
            .build());

    }
}
```
```yaml
resources:
  yada:
    type: aws:cloudwatch:LogGroup
    properties:
      tags:
        Application: serviceA
        Environment: production
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudwatch Log Groups using the `name`. For example:

```sh
 $ pulumi import aws:cloudwatch/logGroup:LogGroup test_group yada
```
 