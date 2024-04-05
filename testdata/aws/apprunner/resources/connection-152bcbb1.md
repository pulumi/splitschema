Manages an App Runner Connection.

> **NOTE:** After creation, you must complete the authentication handshake using the App Runner console.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apprunner.Connection("example", {
    connectionName: "example",
    providerType: "GITHUB",
    tags: {
        Name: "example-apprunner-connection",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apprunner.Connection("example",
    connection_name="example",
    provider_type="GITHUB",
    tags={
        "Name": "example-apprunner-connection",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppRunner.Connection("example", new()
    {
        ConnectionName = "example",
        ProviderType = "GITHUB",
        Tags = 
        {
            { "Name", "example-apprunner-connection" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apprunner"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apprunner.NewConnection(ctx, "example", &apprunner.ConnectionArgs{
			ConnectionName: pulumi.String("example"),
			ProviderType:   pulumi.String("GITHUB"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-apprunner-connection"),
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
import com.pulumi.aws.apprunner.Connection;
import com.pulumi.aws.apprunner.ConnectionArgs;
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
        var example = new Connection("example", ConnectionArgs.builder()        
            .connectionName("example")
            .providerType("GITHUB")
            .tags(Map.of("Name", "example-apprunner-connection"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apprunner:Connection
    properties:
      connectionName: example
      providerType: GITHUB
      tags:
        Name: example-apprunner-connection
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner Connections using the `connection_name`. For example:

```sh
 $ pulumi import aws:apprunner/connection:Connection example example
```
 