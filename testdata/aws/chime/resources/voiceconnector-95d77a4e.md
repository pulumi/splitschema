Enables you to connect your phone system to the telephone network at a substantial cost savings by using SIP trunking.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.chime.VoiceConnector("test", {
    awsRegion: "us-east-1",
    requireEncryption: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.chime.VoiceConnector("test",
    aws_region="us-east-1",
    require_encryption=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Chime.VoiceConnector("test", new()
    {
        AwsRegion = "us-east-1",
        RequireEncryption = true,
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
		_, err := chime.NewVoiceConnector(ctx, "test", &chime.VoiceConnectorArgs{
			AwsRegion:         pulumi.String("us-east-1"),
			RequireEncryption: pulumi.Bool(true),
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
        var test = new VoiceConnector("test", VoiceConnectorArgs.builder()        
            .awsRegion("us-east-1")
            .requireEncryption(true)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:chime:VoiceConnector
    properties:
      awsRegion: us-east-1
      requireEncryption: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Configuration Recorder using the name. For example:

```sh
 $ pulumi import aws:chime/voiceConnector:VoiceConnector test example
```
 