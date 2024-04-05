Resource for managing an AWS SESv2 (Simple Email V2) Configuration Set Event Destination.

{{% examples %}}
## Example Usage
{{% example %}}
### Cloud Watch Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConfigurationSet = new aws.sesv2.ConfigurationSet("exampleConfigurationSet", {configurationSetName: "example"});
const exampleConfigurationSetEventDestination = new aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", {
    configurationSetName: exampleConfigurationSet.configurationSetName,
    eventDestinationName: "example",
    eventDestination: {
        cloudWatchDestination: {
            dimensionConfigurations: [{
                defaultDimensionValue: "example",
                dimensionName: "example",
                dimensionValueSource: "MESSAGE_TAG",
            }],
        },
        enabled: true,
        matchingEventTypes: ["SEND"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
example_configuration_set_event_destination = aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination",
    configuration_set_name=example_configuration_set.configuration_set_name,
    event_destination_name="example",
    event_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationArgs(
        cloud_watch_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationArgs(
            dimension_configurations=[aws.sesv2.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfigurationArgs(
                default_dimension_value="example",
                dimension_name="example",
                dimension_value_source="MESSAGE_TAG",
            )],
        ),
        enabled=True,
        matching_event_types=["SEND"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConfigurationSet = new Aws.SesV2.ConfigurationSet("exampleConfigurationSet", new()
    {
        ConfigurationSetName = "example",
    });

    var exampleConfigurationSetEventDestination = new Aws.SesV2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", new()
    {
        ConfigurationSetName = exampleConfigurationSet.ConfigurationSetName,
        EventDestinationName = "example",
        EventDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationArgs
        {
            CloudWatchDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationArgs
            {
                DimensionConfigurations = new[]
                {
                    new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfigurationArgs
                    {
                        DefaultDimensionValue = "example",
                        DimensionName = "example",
                        DimensionValueSource = "MESSAGE_TAG",
                    },
                },
            },
            Enabled = true,
            MatchingEventTypes = new[]
            {
                "SEND",
            },
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
		exampleConfigurationSet, err := sesv2.NewConfigurationSet(ctx, "exampleConfigurationSet", &sesv2.ConfigurationSetArgs{
			ConfigurationSetName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewConfigurationSetEventDestination(ctx, "exampleConfigurationSetEventDestination", &sesv2.ConfigurationSetEventDestinationArgs{
			ConfigurationSetName: exampleConfigurationSet.ConfigurationSetName,
			EventDestinationName: pulumi.String("example"),
			EventDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationArgs{
				CloudWatchDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationArgs{
					DimensionConfigurations: sesv2.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfigurationArray{
						&sesv2.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfigurationArgs{
							DefaultDimensionValue: pulumi.String("example"),
							DimensionName:         pulumi.String("example"),
							DimensionValueSource:  pulumi.String("MESSAGE_TAG"),
						},
					},
				},
				Enabled: pulumi.Bool(true),
				MatchingEventTypes: pulumi.StringArray{
					pulumi.String("SEND"),
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
import com.pulumi.aws.sesv2.ConfigurationSet;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestination;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationArgs;
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
        var exampleConfigurationSet = new ConfigurationSet("exampleConfigurationSet", ConfigurationSetArgs.builder()        
            .configurationSetName("example")
            .build());

        var exampleConfigurationSetEventDestination = new ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", ConfigurationSetEventDestinationArgs.builder()        
            .configurationSetName(exampleConfigurationSet.configurationSetName())
            .eventDestinationName("example")
            .eventDestination(ConfigurationSetEventDestinationEventDestinationArgs.builder()
                .cloudWatchDestination(ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationArgs.builder()
                    .dimensionConfigurations(ConfigurationSetEventDestinationEventDestinationCloudWatchDestinationDimensionConfigurationArgs.builder()
                        .defaultDimensionValue("example")
                        .dimensionName("example")
                        .dimensionValueSource("MESSAGE_TAG")
                        .build())
                    .build())
                .enabled(true)
                .matchingEventTypes("SEND")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleConfigurationSet:
    type: aws:sesv2:ConfigurationSet
    properties:
      configurationSetName: example
  exampleConfigurationSetEventDestination:
    type: aws:sesv2:ConfigurationSetEventDestination
    properties:
      configurationSetName: ${exampleConfigurationSet.configurationSetName}
      eventDestinationName: example
      eventDestination:
        cloudWatchDestination:
          dimensionConfigurations:
            - defaultDimensionValue: example
              dimensionName: example
              dimensionValueSource: MESSAGE_TAG
        enabled: true
        matchingEventTypes:
          - SEND
```
{{% /example %}}
{{% example %}}
### Kinesis Firehose Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConfigurationSet = new aws.sesv2.ConfigurationSet("exampleConfigurationSet", {configurationSetName: "example"});
const exampleConfigurationSetEventDestination = new aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", {
    configurationSetName: exampleConfigurationSet.configurationSetName,
    eventDestinationName: "example",
    eventDestination: {
        kinesisFirehoseDestination: {
            deliveryStreamArn: aws_kinesis_firehose_delivery_stream.example.arn,
            iamRoleArn: aws_iam_role.example.arn,
        },
        enabled: true,
        matchingEventTypes: ["SEND"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
example_configuration_set_event_destination = aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination",
    configuration_set_name=example_configuration_set.configuration_set_name,
    event_destination_name="example",
    event_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationArgs(
        kinesis_firehose_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationArgs(
            delivery_stream_arn=aws_kinesis_firehose_delivery_stream["example"]["arn"],
            iam_role_arn=aws_iam_role["example"]["arn"],
        ),
        enabled=True,
        matching_event_types=["SEND"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConfigurationSet = new Aws.SesV2.ConfigurationSet("exampleConfigurationSet", new()
    {
        ConfigurationSetName = "example",
    });

    var exampleConfigurationSetEventDestination = new Aws.SesV2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", new()
    {
        ConfigurationSetName = exampleConfigurationSet.ConfigurationSetName,
        EventDestinationName = "example",
        EventDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationArgs
        {
            KinesisFirehoseDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationArgs
            {
                DeliveryStreamArn = aws_kinesis_firehose_delivery_stream.Example.Arn,
                IamRoleArn = aws_iam_role.Example.Arn,
            },
            Enabled = true,
            MatchingEventTypes = new[]
            {
                "SEND",
            },
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
		exampleConfigurationSet, err := sesv2.NewConfigurationSet(ctx, "exampleConfigurationSet", &sesv2.ConfigurationSetArgs{
			ConfigurationSetName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewConfigurationSetEventDestination(ctx, "exampleConfigurationSetEventDestination", &sesv2.ConfigurationSetEventDestinationArgs{
			ConfigurationSetName: exampleConfigurationSet.ConfigurationSetName,
			EventDestinationName: pulumi.String("example"),
			EventDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationArgs{
				KinesisFirehoseDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationArgs{
					DeliveryStreamArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
					IamRoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
				},
				Enabled: pulumi.Bool(true),
				MatchingEventTypes: pulumi.StringArray{
					pulumi.String("SEND"),
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
import com.pulumi.aws.sesv2.ConfigurationSet;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestination;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationArgs;
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
        var exampleConfigurationSet = new ConfigurationSet("exampleConfigurationSet", ConfigurationSetArgs.builder()        
            .configurationSetName("example")
            .build());

        var exampleConfigurationSetEventDestination = new ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", ConfigurationSetEventDestinationArgs.builder()        
            .configurationSetName(exampleConfigurationSet.configurationSetName())
            .eventDestinationName("example")
            .eventDestination(ConfigurationSetEventDestinationEventDestinationArgs.builder()
                .kinesisFirehoseDestination(ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationArgs.builder()
                    .deliveryStreamArn(aws_kinesis_firehose_delivery_stream.example().arn())
                    .iamRoleArn(aws_iam_role.example().arn())
                    .build())
                .enabled(true)
                .matchingEventTypes("SEND")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleConfigurationSet:
    type: aws:sesv2:ConfigurationSet
    properties:
      configurationSetName: example
  exampleConfigurationSetEventDestination:
    type: aws:sesv2:ConfigurationSetEventDestination
    properties:
      configurationSetName: ${exampleConfigurationSet.configurationSetName}
      eventDestinationName: example
      eventDestination:
        kinesisFirehoseDestination:
          deliveryStreamArn: ${aws_kinesis_firehose_delivery_stream.example.arn}
          iamRoleArn: ${aws_iam_role.example.arn}
        enabled: true
        matchingEventTypes:
          - SEND
```
{{% /example %}}
{{% example %}}
### Pinpoint Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConfigurationSet = new aws.sesv2.ConfigurationSet("exampleConfigurationSet", {configurationSetName: "example"});
const exampleConfigurationSetEventDestination = new aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", {
    configurationSetName: exampleConfigurationSet.configurationSetName,
    eventDestinationName: "example",
    eventDestination: {
        pinpointDestination: {
            applicationArn: aws_pinpoint_app.example.arn,
        },
        enabled: true,
        matchingEventTypes: ["SEND"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
example_configuration_set_event_destination = aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination",
    configuration_set_name=example_configuration_set.configuration_set_name,
    event_destination_name="example",
    event_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationArgs(
        pinpoint_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationPinpointDestinationArgs(
            application_arn=aws_pinpoint_app["example"]["arn"],
        ),
        enabled=True,
        matching_event_types=["SEND"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConfigurationSet = new Aws.SesV2.ConfigurationSet("exampleConfigurationSet", new()
    {
        ConfigurationSetName = "example",
    });

    var exampleConfigurationSetEventDestination = new Aws.SesV2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", new()
    {
        ConfigurationSetName = exampleConfigurationSet.ConfigurationSetName,
        EventDestinationName = "example",
        EventDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationArgs
        {
            PinpointDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationPinpointDestinationArgs
            {
                ApplicationArn = aws_pinpoint_app.Example.Arn,
            },
            Enabled = true,
            MatchingEventTypes = new[]
            {
                "SEND",
            },
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
		exampleConfigurationSet, err := sesv2.NewConfigurationSet(ctx, "exampleConfigurationSet", &sesv2.ConfigurationSetArgs{
			ConfigurationSetName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewConfigurationSetEventDestination(ctx, "exampleConfigurationSetEventDestination", &sesv2.ConfigurationSetEventDestinationArgs{
			ConfigurationSetName: exampleConfigurationSet.ConfigurationSetName,
			EventDestinationName: pulumi.String("example"),
			EventDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationArgs{
				PinpointDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationPinpointDestinationArgs{
					ApplicationArn: pulumi.Any(aws_pinpoint_app.Example.Arn),
				},
				Enabled: pulumi.Bool(true),
				MatchingEventTypes: pulumi.StringArray{
					pulumi.String("SEND"),
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
import com.pulumi.aws.sesv2.ConfigurationSet;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestination;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationPinpointDestinationArgs;
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
        var exampleConfigurationSet = new ConfigurationSet("exampleConfigurationSet", ConfigurationSetArgs.builder()        
            .configurationSetName("example")
            .build());

        var exampleConfigurationSetEventDestination = new ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", ConfigurationSetEventDestinationArgs.builder()        
            .configurationSetName(exampleConfigurationSet.configurationSetName())
            .eventDestinationName("example")
            .eventDestination(ConfigurationSetEventDestinationEventDestinationArgs.builder()
                .pinpointDestination(ConfigurationSetEventDestinationEventDestinationPinpointDestinationArgs.builder()
                    .applicationArn(aws_pinpoint_app.example().arn())
                    .build())
                .enabled(true)
                .matchingEventTypes("SEND")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleConfigurationSet:
    type: aws:sesv2:ConfigurationSet
    properties:
      configurationSetName: example
  exampleConfigurationSetEventDestination:
    type: aws:sesv2:ConfigurationSetEventDestination
    properties:
      configurationSetName: ${exampleConfigurationSet.configurationSetName}
      eventDestinationName: example
      eventDestination:
        pinpointDestination:
          applicationArn: ${aws_pinpoint_app.example.arn}
        enabled: true
        matchingEventTypes:
          - SEND
```
{{% /example %}}
{{% example %}}
### SNS Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConfigurationSet = new aws.sesv2.ConfigurationSet("exampleConfigurationSet", {configurationSetName: "example"});
const exampleConfigurationSetEventDestination = new aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", {
    configurationSetName: exampleConfigurationSet.configurationSetName,
    eventDestinationName: "example",
    eventDestination: {
        snsDestination: {
            topicArn: aws_sns_topic.example.arn,
        },
        enabled: true,
        matchingEventTypes: ["SEND"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
example_configuration_set_event_destination = aws.sesv2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination",
    configuration_set_name=example_configuration_set.configuration_set_name,
    event_destination_name="example",
    event_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationArgs(
        sns_destination=aws.sesv2.ConfigurationSetEventDestinationEventDestinationSnsDestinationArgs(
            topic_arn=aws_sns_topic["example"]["arn"],
        ),
        enabled=True,
        matching_event_types=["SEND"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConfigurationSet = new Aws.SesV2.ConfigurationSet("exampleConfigurationSet", new()
    {
        ConfigurationSetName = "example",
    });

    var exampleConfigurationSetEventDestination = new Aws.SesV2.ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", new()
    {
        ConfigurationSetName = exampleConfigurationSet.ConfigurationSetName,
        EventDestinationName = "example",
        EventDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationArgs
        {
            SnsDestination = new Aws.SesV2.Inputs.ConfigurationSetEventDestinationEventDestinationSnsDestinationArgs
            {
                TopicArn = aws_sns_topic.Example.Arn,
            },
            Enabled = true,
            MatchingEventTypes = new[]
            {
                "SEND",
            },
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
		exampleConfigurationSet, err := sesv2.NewConfigurationSet(ctx, "exampleConfigurationSet", &sesv2.ConfigurationSetArgs{
			ConfigurationSetName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewConfigurationSetEventDestination(ctx, "exampleConfigurationSetEventDestination", &sesv2.ConfigurationSetEventDestinationArgs{
			ConfigurationSetName: exampleConfigurationSet.ConfigurationSetName,
			EventDestinationName: pulumi.String("example"),
			EventDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationArgs{
				SnsDestination: &sesv2.ConfigurationSetEventDestinationEventDestinationSnsDestinationArgs{
					TopicArn: pulumi.Any(aws_sns_topic.Example.Arn),
				},
				Enabled: pulumi.Bool(true),
				MatchingEventTypes: pulumi.StringArray{
					pulumi.String("SEND"),
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
import com.pulumi.aws.sesv2.ConfigurationSet;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestination;
import com.pulumi.aws.sesv2.ConfigurationSetEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationArgs;
import com.pulumi.aws.sesv2.inputs.ConfigurationSetEventDestinationEventDestinationSnsDestinationArgs;
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
        var exampleConfigurationSet = new ConfigurationSet("exampleConfigurationSet", ConfigurationSetArgs.builder()        
            .configurationSetName("example")
            .build());

        var exampleConfigurationSetEventDestination = new ConfigurationSetEventDestination("exampleConfigurationSetEventDestination", ConfigurationSetEventDestinationArgs.builder()        
            .configurationSetName(exampleConfigurationSet.configurationSetName())
            .eventDestinationName("example")
            .eventDestination(ConfigurationSetEventDestinationEventDestinationArgs.builder()
                .snsDestination(ConfigurationSetEventDestinationEventDestinationSnsDestinationArgs.builder()
                    .topicArn(aws_sns_topic.example().arn())
                    .build())
                .enabled(true)
                .matchingEventTypes("SEND")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleConfigurationSet:
    type: aws:sesv2:ConfigurationSet
    properties:
      configurationSetName: example
  exampleConfigurationSetEventDestination:
    type: aws:sesv2:ConfigurationSetEventDestination
    properties:
      configurationSetName: ${exampleConfigurationSet.configurationSetName}
      eventDestinationName: example
      eventDestination:
        snsDestination:
          topicArn: ${aws_sns_topic.example.arn}
        enabled: true
        matchingEventTypes:
          - SEND
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Configuration Set Event Destination using the `id` (`configuration_set_name|event_destination_name`). For example:

```sh
 $ pulumi import aws:sesv2/configurationSetEventDestination:ConfigurationSetEventDestination example example_configuration_set|example_event_destination
```
 