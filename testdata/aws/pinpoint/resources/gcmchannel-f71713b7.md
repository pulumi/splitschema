Provides a Pinpoint GCM Channel resource.

> **Note:** Api Key argument will be stored in the raw state as plain-text.
{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const gcm = new aws.pinpoint.GcmChannel("gcm", {
    applicationId: app.applicationId,
    apiKey: "api_key",
});
```
```python
import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
gcm = aws.pinpoint.GcmChannel("gcm",
    application_id=app.application_id,
    api_key="api_key")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var app = new Aws.Pinpoint.App("app");

    var gcm = new Aws.Pinpoint.GcmChannel("gcm", new()
    {
        ApplicationId = app.ApplicationId,
        ApiKey = "api_key",
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
		_, err = pinpoint.NewGcmChannel(ctx, "gcm", &pinpoint.GcmChannelArgs{
			ApplicationId: app.ApplicationId,
			ApiKey:        pulumi.String("api_key"),
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
import com.pulumi.aws.pinpoint.GcmChannel;
import com.pulumi.aws.pinpoint.GcmChannelArgs;
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

        var gcm = new GcmChannel("gcm", GcmChannelArgs.builder()        
            .applicationId(app.applicationId())
            .apiKey("api_key")
            .build());

    }
}
```
```yaml
resources:
  gcm:
    type: aws:pinpoint:GcmChannel
    properties:
      applicationId: ${app.applicationId}
      apiKey: api_key
  app:
    type: aws:pinpoint:App
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Pinpoint GCM Channel using the `application-id`. For example:

```sh
 $ pulumi import aws:pinpoint/gcmChannel:GcmChannel gcm application-id
```
 