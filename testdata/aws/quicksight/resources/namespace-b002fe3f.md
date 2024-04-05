Resource for managing an AWS QuickSight Namespace.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Namespace("example", {namespace: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.Namespace("example", namespace="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.Namespace("example", new()
    {
        NameSpace = "example",
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
		_, err := quicksight.NewNamespace(ctx, "example", &quicksight.NamespaceArgs{
			Namespace: pulumi.String("example"),
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
import com.pulumi.aws.quicksight.Namespace;
import com.pulumi.aws.quicksight.NamespaceArgs;
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
        var example = new Namespace("example", NamespaceArgs.builder()        
            .namespace("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Namespace
    properties:
      namespace: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight Namespace using the AWS account ID and namespace separated by commas (`,`). For example:

```sh
 $ pulumi import aws:quicksight/namespace:Namespace example 123456789012,example
```
 