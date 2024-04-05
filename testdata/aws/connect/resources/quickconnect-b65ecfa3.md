Provides an Amazon Connect Quick Connect resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.connect.QuickConnect("test", {
    description: "quick connect phone number",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    quickConnectConfig: {
        phoneConfigs: [{
            phoneNumber: "+12345678912",
        }],
        quickConnectType: "PHONE_NUMBER",
    },
    tags: {
        Name: "Example Quick Connect",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.QuickConnect("test",
    description="quick connect phone number",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    quick_connect_config=aws.connect.QuickConnectQuickConnectConfigArgs(
        phone_configs=[aws.connect.QuickConnectQuickConnectConfigPhoneConfigArgs(
            phone_number="+12345678912",
        )],
        quick_connect_type="PHONE_NUMBER",
    ),
    tags={
        "Name": "Example Quick Connect",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Connect.QuickConnect("test", new()
    {
        Description = "quick connect phone number",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        QuickConnectConfig = new Aws.Connect.Inputs.QuickConnectQuickConnectConfigArgs
        {
            PhoneConfigs = new[]
            {
                new Aws.Connect.Inputs.QuickConnectQuickConnectConfigPhoneConfigArgs
                {
                    PhoneNumber = "+12345678912",
                },
            },
            QuickConnectType = "PHONE_NUMBER",
        },
        Tags = 
        {
            { "Name", "Example Quick Connect" },
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
		_, err := connect.NewQuickConnect(ctx, "test", &connect.QuickConnectArgs{
			Description: pulumi.String("quick connect phone number"),
			InstanceId:  pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			QuickConnectConfig: &connect.QuickConnectQuickConnectConfigArgs{
				PhoneConfigs: connect.QuickConnectQuickConnectConfigPhoneConfigArray{
					&connect.QuickConnectQuickConnectConfigPhoneConfigArgs{
						PhoneNumber: pulumi.String("+12345678912"),
					},
				},
				QuickConnectType: pulumi.String("PHONE_NUMBER"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example Quick Connect"),
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
import com.pulumi.aws.connect.QuickConnect;
import com.pulumi.aws.connect.QuickConnectArgs;
import com.pulumi.aws.connect.inputs.QuickConnectQuickConnectConfigArgs;
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
        var test = new QuickConnect("test", QuickConnectArgs.builder()        
            .description("quick connect phone number")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .quickConnectConfig(QuickConnectQuickConnectConfigArgs.builder()
                .phoneConfigs(QuickConnectQuickConnectConfigPhoneConfigArgs.builder()
                    .phoneNumber("+12345678912")
                    .build())
                .quickConnectType("PHONE_NUMBER")
                .build())
            .tags(Map.of("Name", "Example Quick Connect"))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:connect:QuickConnect
    properties:
      description: quick connect phone number
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      quickConnectConfig:
        phoneConfigs:
          - phoneNumber: '+12345678912'
        quickConnectType: PHONE_NUMBER
      tags:
        Name: Example Quick Connect
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Quick Connects using the `instance_id` and `quick_connect_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/quickConnect:QuickConnect example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 