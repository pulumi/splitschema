This setting defines how a user interacts with or uses a service or a feature of a service.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testSetting = new aws.ssm.ServiceSetting("testSetting", {
    settingId: "arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled",
    settingValue: "true",
});
```
```python
import pulumi
import pulumi_aws as aws

test_setting = aws.ssm.ServiceSetting("testSetting",
    setting_id="arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled",
    setting_value="true")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testSetting = new Aws.Ssm.ServiceSetting("testSetting", new()
    {
        SettingId = "arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled",
        SettingValue = "true",
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
		_, err := ssm.NewServiceSetting(ctx, "testSetting", &ssm.ServiceSettingArgs{
			SettingId:    pulumi.String("arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled"),
			SettingValue: pulumi.String("true"),
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
import com.pulumi.aws.ssm.ServiceSetting;
import com.pulumi.aws.ssm.ServiceSettingArgs;
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
        var testSetting = new ServiceSetting("testSetting", ServiceSettingArgs.builder()        
            .settingId("arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled")
            .settingValue("true")
            .build());

    }
}
```
```yaml
resources:
  testSetting:
    type: aws:ssm:ServiceSetting
    properties:
      settingId: arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled
      settingValue: 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS SSM Service Setting using the `setting_id`. For example:

```sh
 $ pulumi import aws:ssm/serviceSetting:ServiceSetting example arn:aws:ssm:us-east-1:123456789012:servicesetting/ssm/parameter-store/high-throughput-enabled
```
 