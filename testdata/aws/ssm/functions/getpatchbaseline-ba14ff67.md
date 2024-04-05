Provides an SSM Patch Baseline data source. Useful if you wish to reuse the default baselines provided.

{{% examples %}}
## Example Usage
{{% example %}}

To retrieve a baseline provided by AWS:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const centos = aws.ssm.getPatchBaseline({
    namePrefix: "AWS-",
    operatingSystem: "CENTOS",
    owner: "AWS",
});
```
```python
import pulumi
import pulumi_aws as aws

centos = aws.ssm.get_patch_baseline(name_prefix="AWS-",
    operating_system="CENTOS",
    owner="AWS")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var centos = Aws.Ssm.GetPatchBaseline.Invoke(new()
    {
        NamePrefix = "AWS-",
        OperatingSystem = "CENTOS",
        Owner = "AWS",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
			NamePrefix:      pulumi.StringRef("AWS-"),
			OperatingSystem: pulumi.StringRef("CENTOS"),
			Owner:           "AWS",
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
import com.pulumi.aws.ssm.SsmFunctions;
import com.pulumi.aws.ssm.inputs.GetPatchBaselineArgs;
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
        final var centos = SsmFunctions.getPatchBaseline(GetPatchBaselineArgs.builder()
            .namePrefix("AWS-")
            .operatingSystem("CENTOS")
            .owner("AWS")
            .build());

    }
}
```
```yaml
variables:
  centos:
    fn::invoke:
      Function: aws:ssm:getPatchBaseline
      Arguments:
        namePrefix: AWS-
        operatingSystem: CENTOS
        owner: AWS
```

To retrieve a baseline on your account:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCustom = aws.ssm.getPatchBaseline({
    defaultBaseline: true,
    namePrefix: "MyCustomBaseline",
    operatingSystem: "WINDOWS",
    owner: "Self",
});
```
```python
import pulumi
import pulumi_aws as aws

default_custom = aws.ssm.get_patch_baseline(default_baseline=True,
    name_prefix="MyCustomBaseline",
    operating_system="WINDOWS",
    owner="Self")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultCustom = Aws.Ssm.GetPatchBaseline.Invoke(new()
    {
        DefaultBaseline = true,
        NamePrefix = "MyCustomBaseline",
        OperatingSystem = "WINDOWS",
        Owner = "Self",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.LookupPatchBaseline(ctx, &ssm.LookupPatchBaselineArgs{
			DefaultBaseline: pulumi.BoolRef(true),
			NamePrefix:      pulumi.StringRef("MyCustomBaseline"),
			OperatingSystem: pulumi.StringRef("WINDOWS"),
			Owner:           "Self",
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
import com.pulumi.aws.ssm.SsmFunctions;
import com.pulumi.aws.ssm.inputs.GetPatchBaselineArgs;
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
        final var defaultCustom = SsmFunctions.getPatchBaseline(GetPatchBaselineArgs.builder()
            .defaultBaseline(true)
            .namePrefix("MyCustomBaseline")
            .operatingSystem("WINDOWS")
            .owner("Self")
            .build());

    }
}
```
```yaml
variables:
  defaultCustom:
    fn::invoke:
      Function: aws:ssm:getPatchBaseline
      Arguments:
        defaultBaseline: true
        namePrefix: MyCustomBaseline
        operatingSystem: WINDOWS
        owner: Self
```
{{% /example %}}
{{% /examples %}}