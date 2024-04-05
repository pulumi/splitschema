Creates a new Amazon Redshift Serverless Namespace.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshiftserverless.Namespace("example", {namespaceName: "concurrency-scaling"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshiftserverless.Namespace("example", namespace_name="concurrency-scaling")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedshiftServerless.Namespace("example", new()
    {
        NamespaceName = "concurrency-scaling",
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
		_, err := redshiftserverless.NewNamespace(ctx, "example", &redshiftserverless.NamespaceArgs{
			NamespaceName: pulumi.String("concurrency-scaling"),
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
import com.pulumi.aws.redshiftserverless.Namespace;
import com.pulumi.aws.redshiftserverless.NamespaceArgs;
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
            .namespaceName("concurrency-scaling")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshiftserverless:Namespace
    properties:
      namespaceName: concurrency-scaling
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Serverless Namespaces using the `namespace_name`. For example:

```sh
 $ pulumi import aws:redshiftserverless/namespace:Namespace example example
```
 