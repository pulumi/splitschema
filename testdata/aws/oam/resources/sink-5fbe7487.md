Resource for managing an AWS CloudWatch Observability Access Manager Sink.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.oam.Sink("example", {tags: {
    Env: "prod",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.oam.Sink("example", tags={
    "Env": "prod",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Oam.Sink("example", new()
    {
        Tags = 
        {
            { "Env", "prod" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/oam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := oam.NewSink(ctx, "example", &oam.SinkArgs{
			Tags: pulumi.StringMap{
				"Env": pulumi.String("prod"),
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
import com.pulumi.aws.oam.Sink;
import com.pulumi.aws.oam.SinkArgs;
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
        var example = new Sink("example", SinkArgs.builder()        
            .tags(Map.of("Env", "prod"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:oam:Sink
    properties:
      tags:
        Env: prod
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Observability Access Manager Sink using the `arn`. For example:

```sh
 $ pulumi import aws:oam/sink:Sink example arn:aws:oam:us-west-2:123456789012:sink/sink-id
```
 