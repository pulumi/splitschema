Provides information about a Lambda Layer Version.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const layerName = config.require("layerName");
const existing = aws.lambda.getLayerVersion({
    layerName: layerName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
layer_name = config.require("layerName")
existing = aws.lambda.get_layer_version(layer_name=layer_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var layerName = config.Require("layerName");
    var existing = Aws.Lambda.GetLayerVersion.Invoke(new()
    {
        LayerName = layerName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		layerName := cfg.Require("layerName")
		_, err := lambda.LookupLayerVersion(ctx, &lambda.LookupLayerVersionArgs{
			LayerName: layerName,
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
import com.pulumi.aws.lambda.LambdaFunctions;
import com.pulumi.aws.lambda.inputs.GetLayerVersionArgs;
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
        final var config = ctx.config();
        final var layerName = config.get("layerName");
        final var existing = LambdaFunctions.getLayerVersion(GetLayerVersionArgs.builder()
            .layerName(layerName)
            .build());

    }
}
```
```yaml
configuration:
  layerName:
    type: string
variables:
  existing:
    fn::invoke:
      Function: aws:lambda:getLayerVersion
      Arguments:
        layerName: ${layerName}
```
{{% /example %}}
{{% /examples %}}