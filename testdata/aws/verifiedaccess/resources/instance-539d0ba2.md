Resource for managing a Verified Access Instance.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.Instance("example", {
    description: "example",
    tags: {
        Name: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.Instance("example",
    description="example",
    tags={
        "Name": "example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.Instance("example", new()
    {
        Description = "example",
        Tags = 
        {
            { "Name", "example" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstance(ctx, "example", &verifiedaccess.InstanceArgs{
			Description: pulumi.String("example"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
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
import com.pulumi.aws.verifiedaccess.Instance;
import com.pulumi.aws.verifiedaccess.InstanceArgs;
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
        var example = new Instance("example", InstanceArgs.builder()        
            .description("example")
            .tags(Map.of("Name", "example"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:Instance
    properties:
      description: example
      tags:
        Name: example
```
{{% /example %}}
{{% example %}}
### With `fips_enabled`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.Instance("example", {fipsEnabled: true});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.Instance("example", fips_enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.Instance("example", new()
    {
        FipsEnabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstance(ctx, "example", &verifiedaccess.InstanceArgs{
			FipsEnabled: pulumi.Bool(true),
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
import com.pulumi.aws.verifiedaccess.Instance;
import com.pulumi.aws.verifiedaccess.InstanceArgs;
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
        var example = new Instance("example", InstanceArgs.builder()        
            .fipsEnabled(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:Instance
    properties:
      fipsEnabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Verified Access Instances using the

`id`. For example:

```sh
 $ pulumi import aws:verifiedaccess/instance:Instance example vai-1234567890abcdef0
```
 