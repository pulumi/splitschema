Provides an AWS Backup Global Settings resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.backup.GlobalSettings("test", {globalSettings: {
    isCrossAccountBackupEnabled: "true",
}});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.backup.GlobalSettings("test", global_settings={
    "isCrossAccountBackupEnabled": "true",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Backup.GlobalSettings("test", new()
    {
        GlobalSettingsList = 
        {
            { "isCrossAccountBackupEnabled", "true" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewGlobalSettings(ctx, "test", &backup.GlobalSettingsArgs{
			GlobalSettings: pulumi.StringMap{
				"isCrossAccountBackupEnabled": pulumi.String("true"),
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
import com.pulumi.aws.backup.GlobalSettings;
import com.pulumi.aws.backup.GlobalSettingsArgs;
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
        var test = new GlobalSettings("test", GlobalSettingsArgs.builder()        
            .globalSettings(Map.of("isCrossAccountBackupEnabled", "true"))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:backup:GlobalSettings
    properties:
      globalSettings:
        isCrossAccountBackupEnabled: 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup Global Settings using the `id`. For example:

```sh
 $ pulumi import aws:backup/globalSettings:GlobalSettings example 123456789012
```
 