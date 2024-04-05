Provides a Pinpoint APNs VoIP Channel resource.

> **Note:** All arguments, including certificates and tokens, will be stored in the raw state as plain-text.
{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const app = new aws.pinpoint.App("app", {});
const apnsVoip = new aws.pinpoint.ApnsVoipChannel("apnsVoip", {
    applicationId: app.applicationId,
    certificate: fs.readFileSync("./certificate.pem", "utf8"),
    privateKey: fs.readFileSync("./private_key.key", "utf8"),
});
```
```python
import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
apns_voip = aws.pinpoint.ApnsVoipChannel("apnsVoip",
    application_id=app.application_id,
    certificate=(lambda path: open(path).read())("./certificate.pem"),
    private_key=(lambda path: open(path).read())("./private_key.key"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var app = new Aws.Pinpoint.App("app");

    var apnsVoip = new Aws.Pinpoint.ApnsVoipChannel("apnsVoip", new()
    {
        ApplicationId = app.ApplicationId,
        Certificate = File.ReadAllText("./certificate.pem"),
        PrivateKey = File.ReadAllText("./private_key.key"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := pinpoint.NewApp(ctx, "app", nil)
		if err != nil {
			return err
		}
		_, err = pinpoint.NewApnsVoipChannel(ctx, "apnsVoip", &pinpoint.ApnsVoipChannelArgs{
			ApplicationId: app.ApplicationId,
			Certificate:   readFileOrPanic("./certificate.pem"),
			PrivateKey:    readFileOrPanic("./private_key.key"),
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
import com.pulumi.aws.pinpoint.ApnsVoipChannel;
import com.pulumi.aws.pinpoint.ApnsVoipChannelArgs;
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

        var apnsVoip = new ApnsVoipChannel("apnsVoip", ApnsVoipChannelArgs.builder()        
            .applicationId(app.applicationId())
            .certificate(Files.readString(Paths.get("./certificate.pem")))
            .privateKey(Files.readString(Paths.get("./private_key.key")))
            .build());

    }
}
```
```yaml
resources:
  apnsVoip:
    type: aws:pinpoint:ApnsVoipChannel
    properties:
      applicationId: ${app.applicationId}
      certificate:
        fn::readFile: ./certificate.pem
      privateKey:
        fn::readFile: ./private_key.key
  app:
    type: aws:pinpoint:App
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Pinpoint APNs VoIP Channel using the `application-id`. For example:

```sh
 $ pulumi import aws:pinpoint/apnsVoipChannel:ApnsVoipChannel apns_voip application-id
```
 