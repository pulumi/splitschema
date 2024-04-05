Information about hardware assets in an Outpost.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getAssets({
    arn: data.aws_outposts_outpost.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.outposts.get_assets(arn=data["aws_outposts_outpost"]["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Outposts.GetAssets.Invoke(new()
    {
        Arn = data.Aws_outposts_outpost.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := outposts.GetAssets(ctx, &outposts.GetAssetsArgs{
			Arn: data.Aws_outposts_outpost.Example.Arn,
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
import com.pulumi.aws.outposts.OutpostsFunctions;
import com.pulumi.aws.outposts.inputs.GetAssetsArgs;
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
        final var example = OutpostsFunctions.getAssets(GetAssetsArgs.builder()
            .arn(data.aws_outposts_outpost().example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:outposts:getAssets
      Arguments:
        arn: ${data.aws_outposts_outpost.example.arn}
```
{{% /example %}}
{{% example %}}
### With Host ID Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getAssets({
    arn: data.aws_outposts_outpost.example.arn,
    hostIdFilters: ["h-x38g5n0yd2a0ueb61"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.outposts.get_assets(arn=data["aws_outposts_outpost"]["example"]["arn"],
    host_id_filters=["h-x38g5n0yd2a0ueb61"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Outposts.GetAssets.Invoke(new()
    {
        Arn = data.Aws_outposts_outpost.Example.Arn,
        HostIdFilters = new[]
        {
            "h-x38g5n0yd2a0ueb61",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := outposts.GetAssets(ctx, &outposts.GetAssetsArgs{
			Arn: data.Aws_outposts_outpost.Example.Arn,
			HostIdFilters: []string{
				"h-x38g5n0yd2a0ueb61",
			},
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
import com.pulumi.aws.outposts.OutpostsFunctions;
import com.pulumi.aws.outposts.inputs.GetAssetsArgs;
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
        final var example = OutpostsFunctions.getAssets(GetAssetsArgs.builder()
            .arn(data.aws_outposts_outpost().example().arn())
            .hostIdFilters("h-x38g5n0yd2a0ueb61")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:outposts:getAssets
      Arguments:
        arn: ${data.aws_outposts_outpost.example.arn}
        hostIdFilters:
          - h-x38g5n0yd2a0ueb61
```
{{% /example %}}
{{% example %}}
### With Status ID Filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.outposts.getAssets({
    arn: data.aws_outposts_outpost.example.arn,
    statusIdFilters: ["ACTIVE"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.outposts.get_assets(arn=data["aws_outposts_outpost"]["example"]["arn"],
    status_id_filters=["ACTIVE"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Outposts.GetAssets.Invoke(new()
    {
        Arn = data.Aws_outposts_outpost.Example.Arn,
        StatusIdFilters = new[]
        {
            "ACTIVE",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/outposts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := outposts.GetAssets(ctx, &outposts.GetAssetsArgs{
			Arn: data.Aws_outposts_outpost.Example.Arn,
			StatusIdFilters: []string{
				"ACTIVE",
			},
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
import com.pulumi.aws.outposts.OutpostsFunctions;
import com.pulumi.aws.outposts.inputs.GetAssetsArgs;
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
        final var example = OutpostsFunctions.getAssets(GetAssetsArgs.builder()
            .arn(data.aws_outposts_outpost().example().arn())
            .statusIdFilters("ACTIVE")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:outposts:getAssets
      Arguments:
        arn: ${data.aws_outposts_outpost.example.arn}
        statusIdFilters:
          - ACTIVE
```
{{% /example %}}
{{% /examples %}}