Provides details about a specific Amazon Connect Contact Flow Module.

{{% examples %}}
## Example Usage
{{% example %}}

By `name`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getContactFlowModule({
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    name: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_contact_flow_module(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetContactFlowModule.Invoke(new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Name = "example",
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
		_, err := connect.LookupContactFlowModule(ctx, &connect.LookupContactFlowModuleArgs{
			InstanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
			Name:       pulumi.StringRef("example"),
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
import com.pulumi.aws.connect.inputs.GetContactFlowModuleArgs;
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
        final var example = ConnectFunctions.getContactFlowModule(GetContactFlowModuleArgs.builder()
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .name("example")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getContactFlowModule
      Arguments:
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
        name: example
```

By `contact_flow_module_id`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getContactFlowModule({
    contactFlowModuleId: "cccccccc-bbbb-cccc-dddd-111111111111",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_contact_flow_module(contact_flow_module_id="cccccccc-bbbb-cccc-dddd-111111111111",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetContactFlowModule.Invoke(new()
    {
        ContactFlowModuleId = "cccccccc-bbbb-cccc-dddd-111111111111",
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
		_, err := connect.LookupContactFlowModule(ctx, &connect.LookupContactFlowModuleArgs{
			ContactFlowModuleId: pulumi.StringRef("cccccccc-bbbb-cccc-dddd-111111111111"),
			InstanceId:          "aaaaaaaa-bbbb-cccc-dddd-111111111111",
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
import com.pulumi.aws.connect.inputs.GetContactFlowModuleArgs;
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
        final var example = ConnectFunctions.getContactFlowModule(GetContactFlowModuleArgs.builder()
            .contactFlowModuleId("cccccccc-bbbb-cccc-dddd-111111111111")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getContactFlowModule
      Arguments:
        contactFlowModuleId: cccccccc-bbbb-cccc-dddd-111111111111
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
```
{{% /example %}}
{{% /examples %}}