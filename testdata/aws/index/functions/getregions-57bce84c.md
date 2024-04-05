Provides information about AWS Regions. Can be used to filter regions i.e., by Opt-In status or only regions enabled for current account. To get details like endpoint and description of each region the data source can be combined with the `aws.getRegion` data source.

{{% examples %}}
## Example Usage
{{% example %}}

Enabled AWS Regions:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegions({});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_regions()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegions.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetRegions(ctx, nil, nil)
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionsArgs;
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
        final var current = AwsFunctions.getRegions();

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getRegions
      Arguments: {}
```

All the regions regardless of the availability

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegions({
    allRegions: true,
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_regions(all_regions=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegions.Invoke(new()
    {
        AllRegions = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetRegions(ctx, &aws.GetRegionsArgs{
			AllRegions: pulumi.BoolRef(true),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionsArgs;
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
        final var current = AwsFunctions.getRegions(GetRegionsArgs.builder()
            .allRegions(true)
            .build());

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getRegions
      Arguments:
        allRegions: true
```

To see regions that are filtered by `"not-opted-in"`, the `all_regions` argument needs to be set to `true` or no results will be returned.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegions({
    allRegions: true,
    filters: [{
        name: "opt-in-status",
        values: ["not-opted-in"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_regions(all_regions=True,
    filters=[aws.GetRegionsFilterArgs(
        name="opt-in-status",
        values=["not-opted-in"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegions.Invoke(new()
    {
        AllRegions = true,
        Filters = new[]
        {
            new Aws.Inputs.GetRegionsFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "not-opted-in",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.GetRegions(ctx, &aws.GetRegionsArgs{
			AllRegions: pulumi.BoolRef(true),
			Filters: []aws.GetRegionsFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"not-opted-in",
					},
				},
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionsArgs;
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
        final var current = AwsFunctions.getRegions(GetRegionsArgs.builder()
            .allRegions(true)
            .filters(GetRegionsFilterArgs.builder()
                .name("opt-in-status")
                .values("not-opted-in")
                .build())
            .build());

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getRegions
      Arguments:
        allRegions: true
        filters:
          - name: opt-in-status
            values:
              - not-opted-in
```
{{% /example %}}
{{% /examples %}}