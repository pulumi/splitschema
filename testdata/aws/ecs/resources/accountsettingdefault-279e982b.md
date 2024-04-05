Provides an ECS default account setting for a specific ECS Resource name within a specific region. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html).

> **NOTE:** The AWS API does not delete this resource. When you run `destroy`, the provider will attempt to disable the setting.

> **NOTE:** Your AWS account may not support disabling `containerInstanceLongArnFormat`, `serviceLongArnFormat`, and `taskLongArnFormat`. If your account does not support disabling these, "destroying" this resource will not disable the setting nor cause a provider error. However, the AWS Provider will log an AWS error: `InvalidParameterException: You can no longer disable Long Arn settings`.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ecs.AccountSettingDefault("test", {value: "enabled"});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ecs.AccountSettingDefault("test", value="enabled")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ecs.AccountSettingDefault("test", new()
    {
        Value = "enabled",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewAccountSettingDefault(ctx, "test", &ecs.AccountSettingDefaultArgs{
			Value: pulumi.String("enabled"),
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
import com.pulumi.aws.ecs.AccountSettingDefault;
import com.pulumi.aws.ecs.AccountSettingDefaultArgs;
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
        var test = new AccountSettingDefault("test", AccountSettingDefaultArgs.builder()        
            .value("enabled")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ecs:AccountSettingDefault
    properties:
      value: enabled
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS Account Setting defaults using the `name`. For example:

```sh
 $ pulumi import aws:ecs/accountSettingDefault:AccountSettingDefault example taskLongArnFormat
```
 