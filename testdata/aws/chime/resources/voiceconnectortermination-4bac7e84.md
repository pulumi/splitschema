Enable Termination settings to control outbound calling from your SIP infrastructure.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultVoiceConnector = new aws.chime.VoiceConnector("defaultVoiceConnector", {requireEncryption: true});
const defaultVoiceConnectorTermination = new aws.chime.VoiceConnectorTermination("defaultVoiceConnectorTermination", {
    disabled: false,
    cpsLimit: 1,
    cidrAllowLists: ["50.35.78.96/31"],
    callingRegions: [
        "US",
        "CA",
    ],
    voiceConnectorId: defaultVoiceConnector.id,
});
```
```python
import pulumi
import pulumi_aws as aws

default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
default_voice_connector_termination = aws.chime.VoiceConnectorTermination("defaultVoiceConnectorTermination",
    disabled=False,
    cps_limit=1,
    cidr_allow_lists=["50.35.78.96/31"],
    calling_regions=[
        "US",
        "CA",
    ],
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

    var defaultVoiceConnectorTermination = new Aws.Chime.VoiceConnectorTermination("defaultVoiceConnectorTermination", new()
    {
        Disabled = false,
        CpsLimit = 1,
        CidrAllowLists = new[]
        {
            "50.35.78.96/31",
        },
        CallingRegions = new[]
        {
            "US",
            "CA",
        },
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
		_, err = chime.NewVoiceConnectorTermination(ctx, "defaultVoiceConnectorTermination", &chime.VoiceConnectorTerminationArgs{
			Disabled: pulumi.Bool(false),
			CpsLimit: pulumi.Int(1),
			CidrAllowLists: pulumi.StringArray{
				pulumi.String("50.35.78.96/31"),
			},
			CallingRegions: pulumi.StringArray{
				pulumi.String("US"),
				pulumi.String("CA"),
			},
			VoiceConnectorId: defaultVoiceConnector.ID(),
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
import com.pulumi.aws.chime.VoiceConnectorTermination;
import com.pulumi.aws.chime.VoiceConnectorTerminationArgs;
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

        var defaultVoiceConnectorTermination = new VoiceConnectorTermination("defaultVoiceConnectorTermination", VoiceConnectorTerminationArgs.builder()        
            .disabled(false)
            .cpsLimit(1)
            .cidrAllowLists("50.35.78.96/31")
            .callingRegions(            
                "US",
                "CA")
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
  defaultVoiceConnectorTermination:
    type: aws:chime:VoiceConnectorTermination
    properties:
      disabled: false
      cpsLimit: 1
      cidrAllowLists:
        - 50.35.78.96/31
      callingRegions:
        - US
        - CA
      voiceConnectorId: ${defaultVoiceConnector.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Chime Voice Connector Termination using the `voice_connector_id`. For example:

```sh
 $ pulumi import aws:chime/voiceConnectorTermination:VoiceConnectorTermination default abcdef1ghij2klmno3pqr4
```
 