Provides an SES event destination

{{% examples %}}
## Example Usage
{{% example %}}
### CloudWatch Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cloudwatch = new aws.ses.EventDestination("cloudwatch", {
    configurationSetName: aws_ses_configuration_set.example.name,
    enabled: true,
    matchingTypes: [
        "bounce",
        "send",
    ],
    cloudwatchDestinations: [{
        defaultValue: "default",
        dimensionName: "dimension",
        valueSource: "emailHeader",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

cloudwatch = aws.ses.EventDestination("cloudwatch",
    configuration_set_name=aws_ses_configuration_set["example"]["name"],
    enabled=True,
    matching_types=[
        "bounce",
        "send",
    ],
    cloudwatch_destinations=[aws.ses.EventDestinationCloudwatchDestinationArgs(
        default_value="default",
        dimension_name="dimension",
        value_source="emailHeader",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cloudwatch = new Aws.Ses.EventDestination("cloudwatch", new()
    {
        ConfigurationSetName = aws_ses_configuration_set.Example.Name,
        Enabled = true,
        MatchingTypes = new[]
        {
            "bounce",
            "send",
        },
        CloudwatchDestinations = new[]
        {
            new Aws.Ses.Inputs.EventDestinationCloudwatchDestinationArgs
            {
                DefaultValue = "default",
                DimensionName = "dimension",
                ValueSource = "emailHeader",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEventDestination(ctx, "cloudwatch", &ses.EventDestinationArgs{
			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
			Enabled:              pulumi.Bool(true),
			MatchingTypes: pulumi.StringArray{
				pulumi.String("bounce"),
				pulumi.String("send"),
			},
			CloudwatchDestinations: ses.EventDestinationCloudwatchDestinationArray{
				&ses.EventDestinationCloudwatchDestinationArgs{
					DefaultValue:  pulumi.String("default"),
					DimensionName: pulumi.String("dimension"),
					ValueSource:   pulumi.String("emailHeader"),
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
import com.pulumi.aws.ses.EventDestination;
import com.pulumi.aws.ses.EventDestinationArgs;
import com.pulumi.aws.ses.inputs.EventDestinationCloudwatchDestinationArgs;
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
        var cloudwatch = new EventDestination("cloudwatch", EventDestinationArgs.builder()        
            .configurationSetName(aws_ses_configuration_set.example().name())
            .enabled(true)
            .matchingTypes(            
                "bounce",
                "send")
            .cloudwatchDestinations(EventDestinationCloudwatchDestinationArgs.builder()
                .defaultValue("default")
                .dimensionName("dimension")
                .valueSource("emailHeader")
                .build())
            .build());

    }
}
```
```yaml
resources:
  cloudwatch:
    type: aws:ses:EventDestination
    properties:
      configurationSetName: ${aws_ses_configuration_set.example.name}
      enabled: true
      matchingTypes:
        - bounce
        - send
      cloudwatchDestinations:
        - defaultValue: default
          dimensionName: dimension
          valueSource: emailHeader
```
{{% /example %}}
{{% example %}}
### Kinesis Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const kinesis = new aws.ses.EventDestination("kinesis", {
    configurationSetName: aws_ses_configuration_set.example.name,
    enabled: true,
    matchingTypes: [
        "bounce",
        "send",
    ],
    kinesisDestination: {
        streamArn: aws_kinesis_firehose_delivery_stream.example.arn,
        roleArn: aws_iam_role.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

kinesis = aws.ses.EventDestination("kinesis",
    configuration_set_name=aws_ses_configuration_set["example"]["name"],
    enabled=True,
    matching_types=[
        "bounce",
        "send",
    ],
    kinesis_destination=aws.ses.EventDestinationKinesisDestinationArgs(
        stream_arn=aws_kinesis_firehose_delivery_stream["example"]["arn"],
        role_arn=aws_iam_role["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var kinesis = new Aws.Ses.EventDestination("kinesis", new()
    {
        ConfigurationSetName = aws_ses_configuration_set.Example.Name,
        Enabled = true,
        MatchingTypes = new[]
        {
            "bounce",
            "send",
        },
        KinesisDestination = new Aws.Ses.Inputs.EventDestinationKinesisDestinationArgs
        {
            StreamArn = aws_kinesis_firehose_delivery_stream.Example.Arn,
            RoleArn = aws_iam_role.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEventDestination(ctx, "kinesis", &ses.EventDestinationArgs{
			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
			Enabled:              pulumi.Bool(true),
			MatchingTypes: pulumi.StringArray{
				pulumi.String("bounce"),
				pulumi.String("send"),
			},
			KinesisDestination: &ses.EventDestinationKinesisDestinationArgs{
				StreamArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
				RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.ses.EventDestination;
import com.pulumi.aws.ses.EventDestinationArgs;
import com.pulumi.aws.ses.inputs.EventDestinationKinesisDestinationArgs;
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
        var kinesis = new EventDestination("kinesis", EventDestinationArgs.builder()        
            .configurationSetName(aws_ses_configuration_set.example().name())
            .enabled(true)
            .matchingTypes(            
                "bounce",
                "send")
            .kinesisDestination(EventDestinationKinesisDestinationArgs.builder()
                .streamArn(aws_kinesis_firehose_delivery_stream.example().arn())
                .roleArn(aws_iam_role.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  kinesis:
    type: aws:ses:EventDestination
    properties:
      configurationSetName: ${aws_ses_configuration_set.example.name}
      enabled: true
      matchingTypes:
        - bounce
        - send
      kinesisDestination:
        streamArn: ${aws_kinesis_firehose_delivery_stream.example.arn}
        roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% example %}}
### SNS Destination

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sns = new aws.ses.EventDestination("sns", {
    configurationSetName: aws_ses_configuration_set.example.name,
    enabled: true,
    matchingTypes: [
        "bounce",
        "send",
    ],
    snsDestination: {
        topicArn: aws_sns_topic.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

sns = aws.ses.EventDestination("sns",
    configuration_set_name=aws_ses_configuration_set["example"]["name"],
    enabled=True,
    matching_types=[
        "bounce",
        "send",
    ],
    sns_destination=aws.ses.EventDestinationSnsDestinationArgs(
        topic_arn=aws_sns_topic["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sns = new Aws.Ses.EventDestination("sns", new()
    {
        ConfigurationSetName = aws_ses_configuration_set.Example.Name,
        Enabled = true,
        MatchingTypes = new[]
        {
            "bounce",
            "send",
        },
        SnsDestination = new Aws.Ses.Inputs.EventDestinationSnsDestinationArgs
        {
            TopicArn = aws_sns_topic.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewEventDestination(ctx, "sns", &ses.EventDestinationArgs{
			ConfigurationSetName: pulumi.Any(aws_ses_configuration_set.Example.Name),
			Enabled:              pulumi.Bool(true),
			MatchingTypes: pulumi.StringArray{
				pulumi.String("bounce"),
				pulumi.String("send"),
			},
			SnsDestination: &ses.EventDestinationSnsDestinationArgs{
				TopicArn: pulumi.Any(aws_sns_topic.Example.Arn),
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
import com.pulumi.aws.ses.EventDestination;
import com.pulumi.aws.ses.EventDestinationArgs;
import com.pulumi.aws.ses.inputs.EventDestinationSnsDestinationArgs;
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
        var sns = new EventDestination("sns", EventDestinationArgs.builder()        
            .configurationSetName(aws_ses_configuration_set.example().name())
            .enabled(true)
            .matchingTypes(            
                "bounce",
                "send")
            .snsDestination(EventDestinationSnsDestinationArgs.builder()
                .topicArn(aws_sns_topic.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  sns:
    type: aws:ses:EventDestination
    properties:
      configurationSetName: ${aws_ses_configuration_set.example.name}
      enabled: true
      matchingTypes:
        - bounce
        - send
      snsDestination:
        topicArn: ${aws_sns_topic.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES event destinations using `configuration_set_name` together with the event destination's `name`. For example:

```sh
 $ pulumi import aws:ses/eventDestination:EventDestination sns some-configuration-set-test/event-destination-sns
```
 