Resource for managing QuickSight Group

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Group("example", {groupName: "tf-example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.Group("example", group_name="tf-example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.Group("example", new()
    {
        GroupName = "tf-example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewGroup(ctx, "example", &quicksight.GroupArgs{
			GroupName: pulumi.String("tf-example"),
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
import com.pulumi.aws.quicksight.Group;
import com.pulumi.aws.quicksight.GroupArgs;
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
        var example = new Group("example", GroupArgs.builder()        
            .groupName("tf-example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Group
    properties:
      groupName: tf-example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight Group using the aws account id, namespace and group name separated by `/`. For example:

```sh
 $ pulumi import aws:quicksight/group:Group example 123456789123/default/tf-example
```
 