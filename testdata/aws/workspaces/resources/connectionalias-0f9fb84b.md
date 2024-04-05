Resource for managing an AWS WorkSpaces Connection Alias.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.workspaces.ConnectionAlias("example", {connectionString: "testdomain.test"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.workspaces.ConnectionAlias("example", connection_string="testdomain.test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Workspaces.ConnectionAlias("example", new()
    {
        ConnectionString = "testdomain.test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workspaces.NewConnectionAlias(ctx, "example", &workspaces.ConnectionAliasArgs{
			ConnectionString: pulumi.String("testdomain.test"),
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
import com.pulumi.aws.workspaces.ConnectionAlias;
import com.pulumi.aws.workspaces.ConnectionAliasArgs;
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
        var example = new ConnectionAlias("example", ConnectionAliasArgs.builder()        
            .connectionString("testdomain.test")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:workspaces:ConnectionAlias
    properties:
      connectionString: testdomain.test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WorkSpaces Connection Alias using the connection alias ID. For example:

```sh
 $ pulumi import aws:workspaces/connectionAlias:ConnectionAlias example rft-8012925589
```
 