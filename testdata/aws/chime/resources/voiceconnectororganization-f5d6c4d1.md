Enable origination settings to control inbound calling to your SIP infrastructure.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultVoiceConnector = new aws.chime.VoiceConnector("defaultVoiceConnector", {requireEncryption: true});
const defaultVoiceConnectorOrganization = new aws.chime.VoiceConnectorOrganization("defaultVoiceConnectorOrganization", {
    disabled: false,
    voiceConnectorId: defaultVoiceConnector.id,
    routes: [
        {
            host: "127.0.0.1",
            port: 8081,
            protocol: "TCP",
            priority: 1,
            weight: 1,
        },
        {
            host: "127.0.0.2",
            port: 8082,
            protocol: "TCP",
            priority: 2,
            weight: 10,
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

default_voice_connector = aws.chime.VoiceConnector("defaultVoiceConnector", require_encryption=True)
default_voice_connector_organization = aws.chime.VoiceConnectorOrganization("defaultVoiceConnectorOrganization",
    disabled=False,
    voice_connector_id=default_voice_connector.id,
    routes=[
        aws.chime.VoiceConnectorOrganizationRouteArgs(
            host="127.0.0.1",
            port=8081,
            protocol="TCP",
            priority=1,
            weight=1,
        ),
        aws.chime.VoiceConnectorOrganizationRouteArgs(
            host="127.0.0.2",
            port=8082,
            protocol="TCP",
            priority=2,
            weight=10,
        ),
    ])
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

    var defaultVoiceConnectorOrganization = new Aws.Chime.VoiceConnectorOrganization("defaultVoiceConnectorOrganization", new()
    {
        Disabled = false,
        VoiceConnectorId = defaultVoiceConnector.Id,
        Routes = new[]
        {
            new Aws.Chime.Inputs.VoiceConnectorOrganizationRouteArgs
            {
                Host = "127.0.0.1",
                Port = 8081,
                Protocol = "TCP",
                Priority = 1,
                Weight = 1,
            },
            new Aws.Chime.Inputs.VoiceConnectorOrganizationRouteArgs
            {
                Host = "127.0.0.2",
                Port = 8082,
                Protocol = "TCP",
                Priority = 2,
                Weight = 10,
            },
        },
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
		_, err = chime.NewVoiceConnectorOrganization(ctx, "defaultVoiceConnectorOrganization", &chime.VoiceConnectorOrganizationArgs{
			Disabled:         pulumi.Bool(false),
			VoiceConnectorId: defaultVoiceConnector.ID(),
			Routes: chime.VoiceConnectorOrganizationRouteArray{
				&chime.VoiceConnectorOrganizationRouteArgs{
					Host:     pulumi.String("127.0.0.1"),
					Port:     pulumi.Int(8081),
					Protocol: pulumi.String("TCP"),
					Priority: pulumi.Int(1),
					Weight:   pulumi.Int(1),
				},
				&chime.VoiceConnectorOrganizationRouteArgs{
					Host:     pulumi.String("127.0.0.2"),
					Port:     pulumi.Int(8082),
					Protocol: pulumi.String("TCP"),
					Priority: pulumi.Int(2),
					Weight:   pulumi.Int(10),
				},
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
import com.pulumi.aws.chime.VoiceConnector;
import com.pulumi.aws.chime.VoiceConnectorArgs;
import com.pulumi.aws.chime.VoiceConnectorOrganization;
import com.pulumi.aws.chime.VoiceConnectorOrganizationArgs;
import com.pulumi.aws.chime.inputs.VoiceConnectorOrganizationRouteArgs;
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

        var defaultVoiceConnectorOrganization = new VoiceConnectorOrganization("defaultVoiceConnectorOrganization", VoiceConnectorOrganizationArgs.builder()        
            .disabled(false)
            .voiceConnectorId(defaultVoiceConnector.id())
            .routes(            
                VoiceConnectorOrganizationRouteArgs.builder()
                    .host("127.0.0.1")
                    .port(8081)
                    .protocol("TCP")
                    .priority(1)
                    .weight(1)
                    .build(),
                VoiceConnectorOrganizationRouteArgs.builder()
                    .host("127.0.0.2")
                    .port(8082)
                    .protocol("TCP")
                    .priority(2)
                    .weight(10)
                    .build())
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
  defaultVoiceConnectorOrganization:
    type: aws:chime:VoiceConnectorOrganization
    properties:
      disabled: false
      voiceConnectorId: ${defaultVoiceConnector.id}
      routes:
        - host: 127.0.0.1
          port: 8081
          protocol: TCP
          priority: 1
          weight: 1
        - host: 127.0.0.2
          port: 8082
          protocol: TCP
          priority: 2
          weight: 10
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Chime Voice Connector Origination using the `voice_connector_id`. For example:

```sh
 $ pulumi import aws:chime/voiceConnectorOrganization:VoiceConnectorOrganization default abcdef1ghij2klmno3pqr4
```
 