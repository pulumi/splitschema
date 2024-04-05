Creates and manages an AWS XRay Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.xray.Group("example", {
    filterExpression: "responsetime > 5",
    groupName: "example",
    insightsConfiguration: {
        insightsEnabled: true,
        notificationsEnabled: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.xray.Group("example",
    filter_expression="responsetime > 5",
    group_name="example",
    insights_configuration=aws.xray.GroupInsightsConfigurationArgs(
        insights_enabled=True,
        notifications_enabled=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Xray.Group("example", new()
    {
        FilterExpression = "responsetime > 5",
        GroupName = "example",
        InsightsConfiguration = new Aws.Xray.Inputs.GroupInsightsConfigurationArgs
        {
            InsightsEnabled = true,
            NotificationsEnabled = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/xray"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := xray.NewGroup(ctx, "example", &xray.GroupArgs{
			FilterExpression: pulumi.String("responsetime > 5"),
			GroupName:        pulumi.String("example"),
			InsightsConfiguration: &xray.GroupInsightsConfigurationArgs{
				InsightsEnabled:      pulumi.Bool(true),
				NotificationsEnabled: pulumi.Bool(true),
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
import com.pulumi.aws.xray.Group;
import com.pulumi.aws.xray.GroupArgs;
import com.pulumi.aws.xray.inputs.GroupInsightsConfigurationArgs;
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
        var example = new Group("example", GroupArgs.builder()        
            .filterExpression("responsetime > 5")
            .groupName("example")
            .insightsConfiguration(GroupInsightsConfigurationArgs.builder()
                .insightsEnabled(true)
                .notificationsEnabled(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:xray:Group
    properties:
      filterExpression: responsetime > 5
      groupName: example
      insightsConfiguration:
        insightsEnabled: true
        notificationsEnabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import XRay Groups using the ARN. For example:

```sh
 $ pulumi import aws:xray/group:Group example arn:aws:xray:us-west-2:1234567890:group/example-group/TNGX7SW5U6QY36T4ZMOUA3HVLBYCZTWDIOOXY3CJAXTHSS3YCWUA
```
 