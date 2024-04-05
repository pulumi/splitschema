Use the `aws.pinpoint.SmsChannel` resource to manage Pinpoint SMS Channels.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const sms = new aws.pinpoint.SmsChannel("sms", {applicationId: app.applicationId});
```
```python
import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
sms = aws.pinpoint.SmsChannel("sms", application_id=app.application_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var app = new Aws.Pinpoint.App("app");

    var sms = new Aws.Pinpoint.SmsChannel("sms", new()
    {
        ApplicationId = app.ApplicationId,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := pinpoint.NewApp(ctx, "app", nil)
		if err != nil {
			return err
		}
		_, err = pinpoint.NewSmsChannel(ctx, "sms", &pinpoint.SmsChannelArgs{
			ApplicationId: app.ApplicationId,
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
import com.pulumi.aws.pinpoint.App;
import com.pulumi.aws.pinpoint.SmsChannel;
import com.pulumi.aws.pinpoint.SmsChannelArgs;
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
        var app = new App("app");

        var sms = new SmsChannel("sms", SmsChannelArgs.builder()        
            .applicationId(app.applicationId())
            .build());

    }
}
```
```yaml
resources:
  sms:
    type: aws:pinpoint:SmsChannel
    properties:
      applicationId: ${app.applicationId}
  app:
    type: aws:pinpoint:App
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the Pinpoint SMS Channel using the `application_id`. For example:

```sh
 $ pulumi import aws:pinpoint/smsChannel:SmsChannel sms application-id
```
 