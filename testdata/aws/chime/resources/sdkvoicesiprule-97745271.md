A SIP rule associates your SIP media application with a phone number or a Request URI hostname. You can associate a SIP rule with more than one SIP media application. Each application then runs only that rule.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.chime.SdkvoiceSipRule("example", {
    triggerType: "RequestUriHostname",
    triggerValue: aws_chime_voice_connector["example-voice-connector"].outbound_host_name,
    targetApplications: [{
        priority: 1,
        sipMediaApplicationId: aws_chimesdkvoice_sip_media_application["example-sma"].id,
        awsRegion: "us-east-1",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.chime.SdkvoiceSipRule("example",
    trigger_type="RequestUriHostname",
    trigger_value=aws_chime_voice_connector["example-voice-connector"]["outbound_host_name"],
    target_applications=[aws.chime.SdkvoiceSipRuleTargetApplicationArgs(
        priority=1,
        sip_media_application_id=aws_chimesdkvoice_sip_media_application["example-sma"]["id"],
        aws_region="us-east-1",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Chime.SdkvoiceSipRule("example", new()
    {
        TriggerType = "RequestUriHostname",
        TriggerValue = aws_chime_voice_connector.Example_voice_connector.Outbound_host_name,
        TargetApplications = new[]
        {
            new Aws.Chime.Inputs.SdkvoiceSipRuleTargetApplicationArgs
            {
                Priority = 1,
                SipMediaApplicationId = aws_chimesdkvoice_sip_media_application.Example_sma.Id,
                AwsRegion = "us-east-1",
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
		_, err := chime.NewSdkvoiceSipRule(ctx, "example", &chime.SdkvoiceSipRuleArgs{
			TriggerType:  pulumi.String("RequestUriHostname"),
			TriggerValue: pulumi.Any(aws_chime_voice_connector.ExampleVoiceConnector.Outbound_host_name),
			TargetApplications: chime.SdkvoiceSipRuleTargetApplicationArray{
				&chime.SdkvoiceSipRuleTargetApplicationArgs{
					Priority:              pulumi.Int(1),
					SipMediaApplicationId: pulumi.Any(aws_chimesdkvoice_sip_media_application.ExampleSma.Id),
					AwsRegion:             pulumi.String("us-east-1"),
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
import com.pulumi.aws.chime.SdkvoiceSipRule;
import com.pulumi.aws.chime.SdkvoiceSipRuleArgs;
import com.pulumi.aws.chime.inputs.SdkvoiceSipRuleTargetApplicationArgs;
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
        var example = new SdkvoiceSipRule("example", SdkvoiceSipRuleArgs.builder()        
            .triggerType("RequestUriHostname")
            .triggerValue(aws_chime_voice_connector.example-voice-connector().outbound_host_name())
            .targetApplications(SdkvoiceSipRuleTargetApplicationArgs.builder()
                .priority(1)
                .sipMediaApplicationId(aws_chimesdkvoice_sip_media_application.example-sma().id())
                .awsRegion("us-east-1")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:chime:SdkvoiceSipRule
    properties:
      triggerType: RequestUriHostname
      triggerValue: ${aws_chime_voice_connector"example-voice-connector"[%!s(MISSING)].outbound_host_name}
      targetApplications:
        - priority: 1
          sipMediaApplicationId: ${aws_chimesdkvoice_sip_media_application"example-sma"[%!s(MISSING)].id}
          awsRegion: us-east-1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a ChimeSDKVoice SIP Rule using the `id`. For example:

```sh
 $ pulumi import aws:chime/sdkvoiceSipRule:SdkvoiceSipRule example abcdef123456
```
 