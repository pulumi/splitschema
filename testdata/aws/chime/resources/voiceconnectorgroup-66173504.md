Creates an Amazon Chime Voice Connector group under the administrator's AWS account. You can associate Amazon Chime Voice Connectors with the Amazon Chime Voice Connector group by including VoiceConnectorItems in the request.

You can include Amazon Chime Voice Connectors from different AWS Regions in your group. This creates a fault tolerant mechanism for fallback in case of availability events.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vc1 = new aws.chime.VoiceConnector("vc1", {
    requireEncryption: true,
    awsRegion: "us-east-1",
});
const vc2 = new aws.chime.VoiceConnector("vc2", {
    requireEncryption: true,
    awsRegion: "us-west-2",
});
const group = new aws.chime.VoiceConnectorGroup("group", {connectors: [
    {
        voiceConnectorId: vc1.id,
        priority: 1,
    },
    {
        voiceConnectorId: vc2.id,
        priority: 3,
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

vc1 = aws.chime.VoiceConnector("vc1",
    require_encryption=True,
    aws_region="us-east-1")
vc2 = aws.chime.VoiceConnector("vc2",
    require_encryption=True,
    aws_region="us-west-2")
group = aws.chime.VoiceConnectorGroup("group", connectors=[
    aws.chime.VoiceConnectorGroupConnectorArgs(
        voice_connector_id=vc1.id,
        priority=1,
    ),
    aws.chime.VoiceConnectorGroupConnectorArgs(
        voice_connector_id=vc2.id,
        priority=3,
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
    var vc1 = new Aws.Chime.VoiceConnector("vc1", new()
    {
        RequireEncryption = true,
        AwsRegion = "us-east-1",
    });

    var vc2 = new Aws.Chime.VoiceConnector("vc2", new()
    {
        RequireEncryption = true,
        AwsRegion = "us-west-2",
    });

    var @group = new Aws.Chime.VoiceConnectorGroup("group", new()
    {
        Connectors = new[]
        {
            new Aws.Chime.Inputs.VoiceConnectorGroupConnectorArgs
            {
                VoiceConnectorId = vc1.Id,
                Priority = 1,
            },
            new Aws.Chime.Inputs.VoiceConnectorGroupConnectorArgs
            {
                VoiceConnectorId = vc2.Id,
                Priority = 3,
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
		vc1, err := chime.NewVoiceConnector(ctx, "vc1", &chime.VoiceConnectorArgs{
			RequireEncryption: pulumi.Bool(true),
			AwsRegion:         pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		vc2, err := chime.NewVoiceConnector(ctx, "vc2", &chime.VoiceConnectorArgs{
			RequireEncryption: pulumi.Bool(true),
			AwsRegion:         pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = chime.NewVoiceConnectorGroup(ctx, "group", &chime.VoiceConnectorGroupArgs{
			Connectors: chime.VoiceConnectorGroupConnectorArray{
				&chime.VoiceConnectorGroupConnectorArgs{
					VoiceConnectorId: vc1.ID(),
					Priority:         pulumi.Int(1),
				},
				&chime.VoiceConnectorGroupConnectorArgs{
					VoiceConnectorId: vc2.ID(),
					Priority:         pulumi.Int(3),
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
import com.pulumi.aws.chime.VoiceConnectorGroup;
import com.pulumi.aws.chime.VoiceConnectorGroupArgs;
import com.pulumi.aws.chime.inputs.VoiceConnectorGroupConnectorArgs;
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
        var vc1 = new VoiceConnector("vc1", VoiceConnectorArgs.builder()        
            .requireEncryption(true)
            .awsRegion("us-east-1")
            .build());

        var vc2 = new VoiceConnector("vc2", VoiceConnectorArgs.builder()        
            .requireEncryption(true)
            .awsRegion("us-west-2")
            .build());

        var group = new VoiceConnectorGroup("group", VoiceConnectorGroupArgs.builder()        
            .connectors(            
                VoiceConnectorGroupConnectorArgs.builder()
                    .voiceConnectorId(vc1.id())
                    .priority(1)
                    .build(),
                VoiceConnectorGroupConnectorArgs.builder()
                    .voiceConnectorId(vc2.id())
                    .priority(3)
                    .build())
            .build());

    }
}
```
```yaml
resources:
  vc1:
    type: aws:chime:VoiceConnector
    properties:
      requireEncryption: true
      awsRegion: us-east-1
  vc2:
    type: aws:chime:VoiceConnector
    properties:
      requireEncryption: true
      awsRegion: us-west-2
  group:
    type: aws:chime:VoiceConnectorGroup
    properties:
      connectors:
        - voiceConnectorId: ${vc1.id}
          priority: 1
        - voiceConnectorId: ${vc2.id}
          priority: 3
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Configuration Recorder using the name. For example:

```sh
 $ pulumi import aws:chime/voiceConnectorGroup:VoiceConnectorGroup default example
```
 