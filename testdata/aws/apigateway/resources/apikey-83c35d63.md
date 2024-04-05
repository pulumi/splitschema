Provides an API Gateway API Key.

> **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigateway.ApiKey("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigateway.ApiKey("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGateway.ApiKey("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.NewApiKey(ctx, "example", nil)
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
import com.pulumi.aws.apigateway.ApiKey;
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
        var example = new ApiKey("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:apigateway:ApiKey
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway Keys using the `id`. For example:

```sh
 $ pulumi import aws:apigateway/apiKey:ApiKey example 8bklk8bl1k3sB38D9B3l0enyWT8c09B30lkq0blk
```
 