Resource for managing an AWS SESv2 (Simple Email V2) Configuration Set.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.ConfigurationSet("example", {
    configurationSetName: "example",
    deliveryOptions: {
        tlsPolicy: "REQUIRE",
    },
    reputationOptions: {
        reputationMetricsEnabled: false,
    },
    sendingOptions: {
        sendingEnabled: true,
    },
    suppressionOptions: {
        suppressedReasons: [
            "BOUNCE",
            "COMPLAINT",
        ],
    },
    trackingOptions: {
        customRedirectDomain: "example.com",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.ConfigurationSet("example",
    configuration_set_name="example",
    delivery_options=aws.sesv2.ConfigurationSetDeliveryOptionsArgs(
        tls_policy="REQUIRE",
    ),
    reputation_options=aws.sesv2.ConfigurationSetReputationOptionsArgs(
        reputation_metrics_enabled=False,
    ),
    sending_options=aws.sesv2.ConfigurationSetSendingOptionsArgs(
        sending_enabled=True,
    ),
    suppression_options=aws.sesv2.ConfigurationSetSuppressionOptionsArgs(
        suppressed_reasons=[
            "BOUNCE",
            "COMPLAINT",
        ],
    ),
    tracking_options=aws.sesv2.ConfigurationSetTrackingOptionsArgs(
        custom_redirect_domain="example.com",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.ConfigurationSet("example", new()
    {
        ConfigurationSetName = "example",
        DeliveryOptions = new Aws.SesV2.Inputs.ConfigurationSetDeliveryOptionsArgs
        {
            TlsPolicy = "REQUIRE",
        },
        ReputationOptions = new Aws.SesV2.Inputs.ConfigurationSetReputationOptionsArgs
        {
            ReputationMetricsEnabled = false,
        },
        SendingOptions = new Aws.SesV2.Inputs.ConfigurationSetSendingOptionsArgs
        {
            SendingEnabled = true,
        },
        SuppressionOptions = new Aws.SesV2.Inputs.ConfigurationSetSuppressionOptionsArgs
        {
            SuppressedReasons = new[]
            {
                "BOUNCE",
                "COMPLAINT",
            },
        },
        TrackingOptions = new Aws.SesV2.Inputs.ConfigurationSetTrackingOptionsArgs
        {
            CustomRedirectDomain = "example.com",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewConfigurationSet(ctx, "example", &sesv2.ConfigurationSetArgs{
			ConfigurationSetName: pulumi.String("example"),
			DeliveryOptions: &sesv2.ConfigurationSetDeliveryOptionsArgs{
				TlsPolicy: pulumi.String("REQUIRE"),
			},
			ReputationOptions: &sesv2.ConfigurationSetReputationOptionsArgs{
				ReputationMetricsEnabled: pulumi.Bool(false),
			},
			SendingOptions: &sesv2.ConfigurationSetSendingOptionsArgs{
				SendingEnabled: pulumi.Bool(true),
			},
			SuppressionOptions: &sesv2.ConfigurationSetSuppressionOptionsArgs{
				SuppressedReasons: pulumi.StringArray{
					pulumi.String("BOUNCE"),
					pulumi.String("COMPLAINT"),
				},
			},
			TrackingOptions: &sesv2.ConfigurationSetTrackingOptionsArgs{
				CustomRedirectDomain: pulumi.String("example.com"),
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
import com.pulumi.aws.sesv2.ConfigurationSet;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetDeliveryOptionsArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetReputationOptionsArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetSendingOptionsArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetSuppressionOptionsArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetTrackingOptionsArgs;
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
        var example = new ConfigurationSet("example", ConfigurationSetArgs.builder()        
            .configurationSetName("example")
            .deliveryOptions(ConfigurationSetDeliveryOptionsArgs.builder()
                .tlsPolicy("REQUIRE")
                .build())
            .reputationOptions(ConfigurationSetReputationOptionsArgs.builder()
                .reputationMetricsEnabled(false)
                .build())
            .sendingOptions(ConfigurationSetSendingOptionsArgs.builder()
                .sendingEnabled(true)
                .build())
            .suppressionOptions(ConfigurationSetSuppressionOptionsArgs.builder()
                .suppressedReasons(                
                    "BOUNCE",
                    "COMPLAINT")
                .build())
            .trackingOptions(ConfigurationSetTrackingOptionsArgs.builder()
                .customRedirectDomain("example.com")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:ConfigurationSet
    properties:
      configurationSetName: example
      deliveryOptions:
        tlsPolicy: REQUIRE
      reputationOptions:
        reputationMetricsEnabled: false
      sendingOptions:
        sendingEnabled: true
      suppressionOptions:
        suppressedReasons:
          - BOUNCE
          - COMPLAINT
      trackingOptions:
        customRedirectDomain: example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Configuration Set using the `configuration_set_name`. For example:

```sh
 $ pulumi import aws:sesv2/configurationSet:ConfigurationSet example example
```
 