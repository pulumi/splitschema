Provides an Elastic File System (EFS) File System resource.

{{% examples %}}
## Example Usage
{{% example %}}
### EFS File System w/ tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.efs.FileSystem("foo", {tags: {
    Name: "MyProduct",
}});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.efs.FileSystem("foo", tags={
    "Name": "MyProduct",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Efs.FileSystem("foo", new()
    {
        Tags = 
        {
            { "Name", "MyProduct" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.NewFileSystem(ctx, "foo", &efs.FileSystemArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("MyProduct"),
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
import com.pulumi.aws.efs.FileSystem;
import com.pulumi.aws.efs.FileSystemArgs;
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
        var foo = new FileSystem("foo", FileSystemArgs.builder()        
            .tags(Map.of("Name", "MyProduct"))
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:efs:FileSystem
    properties:
      tags:
        Name: MyProduct
```
{{% /example %}}
{{% example %}}
### Using lifecycle policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooWithLifecylePolicy = new aws.efs.FileSystem("fooWithLifecylePolicy", {lifecyclePolicies: [{
    transitionToIa: "AFTER_30_DAYS",
}]});
```
```python
import pulumi
import pulumi_aws as aws

foo_with_lifecyle_policy = aws.efs.FileSystem("fooWithLifecylePolicy", lifecycle_policies=[aws.efs.FileSystemLifecyclePolicyArgs(
    transition_to_ia="AFTER_30_DAYS",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooWithLifecylePolicy = new Aws.Efs.FileSystem("fooWithLifecylePolicy", new()
    {
        LifecyclePolicies = new[]
        {
            new Aws.Efs.Inputs.FileSystemLifecyclePolicyArgs
            {
                TransitionToIa = "AFTER_30_DAYS",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := efs.NewFileSystem(ctx, "fooWithLifecylePolicy", &efs.FileSystemArgs{
			LifecyclePolicies: efs.FileSystemLifecyclePolicyArray{
				&efs.FileSystemLifecyclePolicyArgs{
					TransitionToIa: pulumi.String("AFTER_30_DAYS"),
				},
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
import com.pulumi.aws.efs.FileSystem;
import com.pulumi.aws.efs.FileSystemArgs;
import com.pulumi.aws.efs.inputs.FileSystemLifecyclePolicyArgs;
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
        var fooWithLifecylePolicy = new FileSystem("fooWithLifecylePolicy", FileSystemArgs.builder()        
            .lifecyclePolicies(FileSystemLifecyclePolicyArgs.builder()
                .transitionToIa("AFTER_30_DAYS")
                .build())
            .build());

    }
}
```
```yaml
resources:
  fooWithLifecylePolicy:
    type: aws:efs:FileSystem
    properties:
      lifecyclePolicies:
        - transitionToIa: AFTER_30_DAYS
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the EFS file systems using the `id`. For example:

```sh
 $ pulumi import aws:efs/fileSystem:FileSystem foo fs-6fa144c6
```
 