Provides an AppConfig Application resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.Application("example", {
    description: "Example AppConfig Application",
    tags: {
        Type: "AppConfig Application",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appconfig.Application("example",
    description="Example AppConfig Application",
    tags={
        "Type": "AppConfig Application",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppConfig.Application("example", new()
    {
        Description = "Example AppConfig Application",
        Tags = 
        {
            { "Type", "AppConfig Application" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfig.NewApplication(ctx, "example", &appconfig.ApplicationArgs{
			Description: pulumi.String("Example AppConfig Application"),
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Application"),
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
import com.pulumi.aws.appconfig.Application;
import com.pulumi.aws.appconfig.ApplicationArgs;
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
        var example = new Application("example", ApplicationArgs.builder()        
            .description("Example AppConfig Application")
            .tags(Map.of("Type", "AppConfig Application"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:Application
    properties:
      description: Example AppConfig Application
      tags:
        Type: AppConfig Application
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Applications using their application ID. For example:

```sh
 $ pulumi import aws:appconfig/application:Application example 71rxuzt
```
 