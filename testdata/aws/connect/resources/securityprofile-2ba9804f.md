Provides an Amazon Connect Security Profile resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.SecurityProfile("example", {
    description: "example description",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    permissions: [
        "BasicAgentAccess",
        "OutboundCallAccess",
    ],
    tags: {
        Name: "Example Security Profile",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.SecurityProfile("example",
    description="example description",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    permissions=[
        "BasicAgentAccess",
        "OutboundCallAccess",
    ],
    tags={
        "Name": "Example Security Profile",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.SecurityProfile("example", new()
    {
        Description = "example description",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Permissions = new[]
        {
            "BasicAgentAccess",
            "OutboundCallAccess",
        },
        Tags = 
        {
            { "Name", "Example Security Profile" },
        },
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
		_, err := connect.NewSecurityProfile(ctx, "example", &connect.SecurityProfileArgs{
			Description: pulumi.String("example description"),
			InstanceId:  pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			Permissions: pulumi.StringArray{
				pulumi.String("BasicAgentAccess"),
				pulumi.String("OutboundCallAccess"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Security Profile"),
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
import com.pulumi.aws.connect.SecurityProfile;
import com.pulumi.aws.connect.SecurityProfileArgs;
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
        var example = new SecurityProfile("example", SecurityProfileArgs.builder()        
            .description("example description")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .permissions(            
                "BasicAgentAccess",
                "OutboundCallAccess")
            .tags(Map.of("Name", "Example Security Profile"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:SecurityProfile
    properties:
      description: example description
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      permissions:
        - BasicAgentAccess
        - OutboundCallAccess
      tags:
        Name: Example Security Profile
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Security Profiles using the `instance_id` and `security_profile_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/securityProfile:SecurityProfile example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 