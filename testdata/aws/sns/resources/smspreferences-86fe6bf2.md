Provides a way to set SNS SMS preferences.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const updateSmsPrefs = new aws.sns.SmsPreferences("updateSmsPrefs", {});
```
```python
import pulumi
import pulumi_aws as aws

update_sms_prefs = aws.sns.SmsPreferences("updateSmsPrefs")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var updateSmsPrefs = new Aws.Sns.SmsPreferences("updateSmsPrefs");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.NewSmsPreferences(ctx, "updateSmsPrefs", nil)
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
import com.pulumi.aws.sns.SmsPreferences;
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
        var updateSmsPrefs = new SmsPreferences("updateSmsPrefs");

    }
}
```
```yaml
resources:
  updateSmsPrefs:
    type: aws:sns:SmsPreferences
```
{{% /example %}}
{{% /examples %}}