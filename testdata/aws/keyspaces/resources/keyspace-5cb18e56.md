Provides a Keyspaces Keyspace.

More information about keyspaces can be found in the [Keyspaces User Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/what-is-keyspaces.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.keyspaces.Keyspace("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.keyspaces.Keyspace("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Keyspaces.Keyspace("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/keyspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := keyspaces.NewKeyspace(ctx, "example", nil)
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
import com.pulumi.aws.keyspaces.Keyspace;
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
        var example = new Keyspace("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:keyspaces:Keyspace
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a keyspace using the `name`. For example:

```sh
 $ pulumi import aws:keyspaces/keyspace:Keyspace example my_keyspace
```
 