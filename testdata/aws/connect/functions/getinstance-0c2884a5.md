Provides details about a specific Amazon Connect Instance.

{{% examples %}}
## Example Usage
{{% example %}}

By instance_alias

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = aws.connect.getInstance({
    instanceAlias: "foo",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.connect.get_instance(instance_alias="foo")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = Aws.Connect.GetInstance.Invoke(new()
    {
        InstanceAlias = "foo",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.LookupInstance(ctx, &connect.LookupInstanceArgs{
			InstanceAlias: pulumi.StringRef("foo"),
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
import com.pulumi.aws.connect.ConnectFunctions;
import com.pulumi.aws.connect.inputs.GetInstanceArgs;
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
        final var foo = ConnectFunctions.getInstance(GetInstanceArgs.builder()
            .instanceAlias("foo")
            .build());

    }
}
```
```yaml
variables:
  foo:
    fn::invoke:
      Function: aws:connect:getInstance
      Arguments:
        instanceAlias: foo
```

By instance_id

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = aws.connect.getInstance({
    instanceId: "97afc98d-101a-ba98-ab97-ae114fc115ec",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.connect.get_instance(instance_id="97afc98d-101a-ba98-ab97-ae114fc115ec")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = Aws.Connect.GetInstance.Invoke(new()
    {
        InstanceId = "97afc98d-101a-ba98-ab97-ae114fc115ec",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.LookupInstance(ctx, &connect.LookupInstanceArgs{
			InstanceId: pulumi.StringRef("97afc98d-101a-ba98-ab97-ae114fc115ec"),
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
import com.pulumi.aws.connect.ConnectFunctions;
import com.pulumi.aws.connect.inputs.GetInstanceArgs;
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
        final var foo = ConnectFunctions.getInstance(GetInstanceArgs.builder()
            .instanceId("97afc98d-101a-ba98-ab97-ae114fc115ec")
            .build());

    }
}
```
```yaml
variables:
  foo:
    fn::invoke:
      Function: aws:connect:getInstance
      Arguments:
        instanceId: 97afc98d-101a-ba98-ab97-ae114fc115ec
```
{{% /example %}}
{{% /examples %}}