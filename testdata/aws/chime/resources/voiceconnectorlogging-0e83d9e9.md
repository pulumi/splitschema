Adds a logging configuration for the specified Amazon Chime Voice Connector. The logging configuration specifies whether SIP message logs are enabled for sending to Amazon CloudWatch Logs.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultVoiceConnector = new aws.chime.VoiceConnector("defaultVoiceConnector", {requireEncryption: true});
const defaultVoiceConnectorLogging = new aws.chime.VoiceConnectorLogging("defaultVoiceConnectorLogging", {
    enableSipLogs: true,
    enableMediaMetricLogs: true,
    voiceConnectorId: defaultVoiceConnector.id,
});
```
```python
import pulumi
import pulumi_aws as aws

default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
default_voice_connector_logging = aws.chime.VoiceConnectorLogging("defaultVoiceConnectorLogging",
    enable_sip_logs=True,
    enable_media_metric_logs=True,
    voice_connector_id=default_voice_connector.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultVoiceConnector = new Aws.Chime.VoiceConnector("defaultVoiceConnector", new()
    {
        RequireEncryption = true,
    });

    var defaultVoiceConnectorLogging = new Aws.Chime.VoiceConnectorLogging("defaultVoiceConnectorLogging", new()
    {
        EnableSipLogs = true,
        EnableMediaMetricLogs = true,
        VoiceConnectorId = defaultVoiceConnector.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chime"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultVoiceConnector, err := chime.NewVoiceConnector(ctx, "defaultVoiceConnector", &chime.VoiceConnectorArgs{
			RequireEncryption: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = chime.NewVoiceConnectorLogging(ctx, "defaultVoiceConnectorLogging", &chime.VoiceConnectorLoggingArgs{
			EnableSipLogs:         pulumi.Bool(true),
			EnableMediaMetricLogs: pulumi.Bool(true),
			VoiceConnectorId:      defaultVoiceConnector.ID(),
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
import com.pulumi.aws.chime.VoiceConnector;
import com.pulumi.aws.chime.VoiceConnectorArgs;
import com.pulumi.aws.chime.VoiceConnectorLogging;
import com.pulumi.aws.chime.VoiceConnectorLoggingArgs;
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
        var defaultVoiceConnector = new VoiceConnector("defaultVoiceConnector", VoiceConnectorArgs.builder()        
            .requireEncryption(true)
            .build());

        var defaultVoiceConnectorLogging = new VoiceConnectorLogging("defaultVoiceConnectorLogging", VoiceConnectorLoggingArgs.builder()        
            .enableSipLogs(true)
            .enableMediaMetricLogs(true)
            .voiceConnectorId(defaultVoiceConnector.id())
            .build());

    }
}
```
```yaml
resources:
  defaultVoiceConnector:
    type: aws:chime:VoiceConnector
    properties:
      requireEncryption: true
  defaultVoiceConnectorLogging:
    type: aws:chime:VoiceConnectorLogging
    properties:
      enableSipLogs: true
      enableMediaMetricLogs: true
      voiceConnectorId: ${defaultVoiceConnector.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Chime Voice Connector Logging using the `voice_connector_id`. For example:

```sh
 $ pulumi import aws:chime/voiceConnectorLogging:VoiceConnectorLogging default abcdef1ghij2klmno3pqr4
```
 