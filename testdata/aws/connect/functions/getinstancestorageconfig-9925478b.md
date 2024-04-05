Provides details about a specific Amazon Connect Instance Storage Config.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getInstanceStorageConfig({
    associationId: "1234567890123456789012345678901234567890123456789012345678901234",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    resourceType: "CONTACT_TRACE_RECORDS",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_instance_storage_config(association_id="1234567890123456789012345678901234567890123456789012345678901234",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    resource_type="CONTACT_TRACE_RECORDS")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetInstanceStorageConfig.Invoke(new()
    {
        AssociationId = "1234567890123456789012345678901234567890123456789012345678901234",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        ResourceType = "CONTACT_TRACE_RECORDS",
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
		_, err := connect.LookupInstanceStorageConfig(ctx, &connect.LookupInstanceStorageConfigArgs{
			AssociationId: "1234567890123456789012345678901234567890123456789012345678901234",
			InstanceId:    "aaaaaaaa-bbbb-cccc-dddd-111111111111",
			ResourceType:  "CONTACT_TRACE_RECORDS",
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
import com.pulumi.aws.connect.inputs.GetInstanceStorageConfigArgs;
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
        final var example = ConnectFunctions.getInstanceStorageConfig(GetInstanceStorageConfigArgs.builder()
            .associationId("1234567890123456789012345678901234567890123456789012345678901234")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .resourceType("CONTACT_TRACE_RECORDS")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getInstanceStorageConfig
      Arguments:
        associationId: 1234567890123456789012345678901234567890123456789012345678901234
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
        resourceType: CONTACT_TRACE_RECORDS
```
{{% /example %}}
{{% /examples %}}