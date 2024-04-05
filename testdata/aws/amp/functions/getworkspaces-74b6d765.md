Provides the aliases, ARNs, and workspace IDs of Amazon Prometheus workspaces.

{{% examples %}}
## Example Usage
{{% example %}}

The following example returns all of the workspaces in a region:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.amp.getWorkspaces({});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.amp.get_workspaces()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Amp.GetWorkspaces.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := amp.GetWorkspaces(ctx, nil, nil)
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
import com.pulumi.aws.amp.AmpFunctions;
import com.pulumi.aws.amp.inputs.GetWorkspacesArgs;
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
        final var example = AmpFunctions.getWorkspaces();

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:amp:getWorkspaces
      Arguments: {}
```

The following example filters the workspaces by alias. Only the workspaces with
aliases that begin with the value of `alias_prefix` will be returned:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.amp.getWorkspaces({
    aliasPrefix: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.amp.get_workspaces(alias_prefix="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Amp.GetWorkspaces.Invoke(new()
    {
        AliasPrefix = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := amp.GetWorkspaces(ctx, &amp.GetWorkspacesArgs{
			AliasPrefix: pulumi.StringRef("example"),
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
import com.pulumi.aws.amp.AmpFunctions;
import com.pulumi.aws.amp.inputs.GetWorkspacesArgs;
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
        final var example = AmpFunctions.getWorkspaces(GetWorkspacesArgs.builder()
            .aliasPrefix("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:amp:getWorkspaces
      Arguments:
        aliasPrefix: example
```
{{% /example %}}
{{% /examples %}}