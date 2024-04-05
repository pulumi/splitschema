Provides details about a specific Amazon Connect Contact Flow.

{{% examples %}}
## Example Usage
{{% example %}}

By name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.connect.getContactFlow({
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    name: "Test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.get_contact_flow(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    name="Test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Connect.GetContactFlow.Invoke(new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Name = "Test",
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
		_, err := connect.LookupContactFlow(ctx, &connect.LookupContactFlowArgs{
			InstanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
			Name:       pulumi.StringRef("Test"),
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
import com.pulumi.aws.connect.inputs.GetContactFlowArgs;
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
        final var test = ConnectFunctions.getContactFlow(GetContactFlowArgs.builder()
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .name("Test")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:connect:getContactFlow
      Arguments:
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
        name: Test
```

By contact_flow_id

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.connect.getContactFlow({
    contactFlowId: "cccccccc-bbbb-cccc-dddd-111111111111",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.get_contact_flow(contact_flow_id="cccccccc-bbbb-cccc-dddd-111111111111",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Connect.GetContactFlow.Invoke(new()
    {
        ContactFlowId = "cccccccc-bbbb-cccc-dddd-111111111111",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
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
		_, err := connect.LookupContactFlow(ctx, &connect.LookupContactFlowArgs{
			ContactFlowId: pulumi.StringRef("cccccccc-bbbb-cccc-dddd-111111111111"),
			InstanceId:    "aaaaaaaa-bbbb-cccc-dddd-111111111111",
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
import com.pulumi.aws.connect.inputs.GetContactFlowArgs;
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
        final var test = ConnectFunctions.getContactFlow(GetContactFlowArgs.builder()
            .contactFlowId("cccccccc-bbbb-cccc-dddd-111111111111")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:connect:getContactFlow
      Arguments:
        contactFlowId: cccccccc-bbbb-cccc-dddd-111111111111
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
```
{{% /example %}}
{{% /examples %}}