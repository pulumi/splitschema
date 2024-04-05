Provides an Amazon Inspector Classic Resource Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.inspector.ResourceGroup("example", {tags: {
    Env: "bar",
    Name: "foo",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.inspector.ResourceGroup("example", tags={
    "Env": "bar",
    "Name": "foo",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Inspector.ResourceGroup("example", new()
    {
        Tags = 
        {
            { "Env", "bar" },
            { "Name", "foo" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/inspector"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := inspector.NewResourceGroup(ctx, "example", &inspector.ResourceGroupArgs{
			Tags: pulumi.StringMap{
				"Env":  pulumi.String("bar"),
				"Name": pulumi.String("foo"),
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
import com.pulumi.aws.inspector.ResourceGroup;
import com.pulumi.aws.inspector.ResourceGroupArgs;
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
        var example = new ResourceGroup("example", ResourceGroupArgs.builder()        
            .tags(Map.ofEntries(
                Map.entry("Env", "bar"),
                Map.entry("Name", "foo")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:inspector:ResourceGroup
    properties:
      tags:
        Env: bar
        Name: foo
```
{{% /example %}}
{{% /examples %}}