Provides an AWS Network Firewall Logging Configuration Resource

{{% examples %}}
## Example Usage
{{% example %}}
### Logging to S3

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.LoggingConfiguration("example", {
    firewallArn: aws_networkfirewall_firewall.example.arn,
    loggingConfiguration: {
        logDestinationConfigs: [{
            logDestination: {
                bucketName: aws_s3_bucket.example.bucket,
                prefix: "/example",
            },
            logDestinationType: "S3",
            logType: "FLOW",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.LoggingConfiguration("example",
    firewall_arn=aws_networkfirewall_firewall["example"]["arn"],
    logging_configuration=aws.networkfirewall.LoggingConfigurationLoggingConfigurationArgs(
        log_destination_configs=[aws.networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs(
            log_destination={
                "bucketName": aws_s3_bucket["example"]["bucket"],
                "prefix": "/example",
            },
            log_destination_type="S3",
            log_type="FLOW",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.LoggingConfiguration("example", new()
    {
        FirewallArn = aws_networkfirewall_firewall.Example.Arn,
        LoggingConfig = new Aws.NetworkFirewall.Inputs.LoggingConfigurationLoggingConfigurationArgs
        {
            LogDestinationConfigs = new[]
            {
                new Aws.NetworkFirewall.Inputs.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs
                {
                    LogDestination = 
                    {
                        { "bucketName", aws_s3_bucket.Example.Bucket },
                        { "prefix", "/example" },
                    },
                    LogDestinationType = "S3",
                    LogType = "FLOW",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.NewLoggingConfiguration(ctx, "example", &networkfirewall.LoggingConfigurationArgs{
			FirewallArn: pulumi.Any(aws_networkfirewall_firewall.Example.Arn),
			LoggingConfiguration: &networkfirewall.LoggingConfigurationLoggingConfigurationArgs{
				LogDestinationConfigs: networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArray{
					&networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs{
						LogDestination: pulumi.StringMap{
							"bucketName": pulumi.Any(aws_s3_bucket.Example.Bucket),
							"prefix":     pulumi.String("/example"),
						},
						LogDestinationType: pulumi.String("S3"),
						LogType:            pulumi.String("FLOW"),
					},
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
import com.pulumi.aws.networkfirewall.LoggingConfiguration;
import com.pulumi.aws.networkfirewall.LoggingConfigurationArgs;
import com.pulumi.aws.networkfirewall.inputs.LoggingConfigurationLoggingConfigurationArgs;
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
        var example = new LoggingConfiguration("example", LoggingConfigurationArgs.builder()        
            .firewallArn(aws_networkfirewall_firewall.example().arn())
            .loggingConfiguration(LoggingConfigurationLoggingConfigurationArgs.builder()
                .logDestinationConfigs(LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs.builder()
                    .logDestination(Map.ofEntries(
                        Map.entry("bucketName", aws_s3_bucket.example().bucket()),
                        Map.entry("prefix", "/example")
                    ))
                    .logDestinationType("S3")
                    .logType("FLOW")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:LoggingConfiguration
    properties:
      firewallArn: ${aws_networkfirewall_firewall.example.arn}
      loggingConfiguration:
        logDestinationConfigs:
          - logDestination:
              bucketName: ${aws_s3_bucket.example.bucket}
              prefix: /example
            logDestinationType: S3
            logType: FLOW
```
{{% /example %}}
{{% example %}}
### Logging to CloudWatch

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.LoggingConfiguration("example", {
    firewallArn: aws_networkfirewall_firewall.example.arn,
    loggingConfiguration: {
        logDestinationConfigs: [{
            logDestination: {
                logGroup: aws_cloudwatch_log_group.example.name,
            },
            logDestinationType: "CloudWatchLogs",
            logType: "ALERT",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.LoggingConfiguration("example",
    firewall_arn=aws_networkfirewall_firewall["example"]["arn"],
    logging_configuration=aws.networkfirewall.LoggingConfigurationLoggingConfigurationArgs(
        log_destination_configs=[aws.networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs(
            log_destination={
                "logGroup": aws_cloudwatch_log_group["example"]["name"],
            },
            log_destination_type="CloudWatchLogs",
            log_type="ALERT",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.LoggingConfiguration("example", new()
    {
        FirewallArn = aws_networkfirewall_firewall.Example.Arn,
        LoggingConfig = new Aws.NetworkFirewall.Inputs.LoggingConfigurationLoggingConfigurationArgs
        {
            LogDestinationConfigs = new[]
            {
                new Aws.NetworkFirewall.Inputs.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs
                {
                    LogDestination = 
                    {
                        { "logGroup", aws_cloudwatch_log_group.Example.Name },
                    },
                    LogDestinationType = "CloudWatchLogs",
                    LogType = "ALERT",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.NewLoggingConfiguration(ctx, "example", &networkfirewall.LoggingConfigurationArgs{
			FirewallArn: pulumi.Any(aws_networkfirewall_firewall.Example.Arn),
			LoggingConfiguration: &networkfirewall.LoggingConfigurationLoggingConfigurationArgs{
				LogDestinationConfigs: networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArray{
					&networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs{
						LogDestination: pulumi.StringMap{
							"logGroup": pulumi.Any(aws_cloudwatch_log_group.Example.Name),
						},
						LogDestinationType: pulumi.String("CloudWatchLogs"),
						LogType:            pulumi.String("ALERT"),
					},
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
import com.pulumi.aws.networkfirewall.LoggingConfiguration;
import com.pulumi.aws.networkfirewall.LoggingConfigurationArgs;
import com.pulumi.aws.networkfirewall.inputs.LoggingConfigurationLoggingConfigurationArgs;
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
        var example = new LoggingConfiguration("example", LoggingConfigurationArgs.builder()        
            .firewallArn(aws_networkfirewall_firewall.example().arn())
            .loggingConfiguration(LoggingConfigurationLoggingConfigurationArgs.builder()
                .logDestinationConfigs(LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs.builder()
                    .logDestination(Map.of("logGroup", aws_cloudwatch_log_group.example().name()))
                    .logDestinationType("CloudWatchLogs")
                    .logType("ALERT")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:LoggingConfiguration
    properties:
      firewallArn: ${aws_networkfirewall_firewall.example.arn}
      loggingConfiguration:
        logDestinationConfigs:
          - logDestination:
              logGroup: ${aws_cloudwatch_log_group.example.name}
            logDestinationType: CloudWatchLogs
            logType: ALERT
```
{{% /example %}}
{{% example %}}
### Logging to Kinesis Data Firehose

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.networkfirewall.LoggingConfiguration("example", {
    firewallArn: aws_networkfirewall_firewall.example.arn,
    loggingConfiguration: {
        logDestinationConfigs: [{
            logDestination: {
                deliveryStream: aws_kinesis_firehose_delivery_stream.example.name,
            },
            logDestinationType: "KinesisDataFirehose",
            logType: "ALERT",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.LoggingConfiguration("example",
    firewall_arn=aws_networkfirewall_firewall["example"]["arn"],
    logging_configuration=aws.networkfirewall.LoggingConfigurationLoggingConfigurationArgs(
        log_destination_configs=[aws.networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs(
            log_destination={
                "deliveryStream": aws_kinesis_firehose_delivery_stream["example"]["name"],
            },
            log_destination_type="KinesisDataFirehose",
            log_type="ALERT",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.NetworkFirewall.LoggingConfiguration("example", new()
    {
        FirewallArn = aws_networkfirewall_firewall.Example.Arn,
        LoggingConfig = new Aws.NetworkFirewall.Inputs.LoggingConfigurationLoggingConfigurationArgs
        {
            LogDestinationConfigs = new[]
            {
                new Aws.NetworkFirewall.Inputs.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs
                {
                    LogDestination = 
                    {
                        { "deliveryStream", aws_kinesis_firehose_delivery_stream.Example.Name },
                    },
                    LogDestinationType = "KinesisDataFirehose",
                    LogType = "ALERT",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.NewLoggingConfiguration(ctx, "example", &networkfirewall.LoggingConfigurationArgs{
			FirewallArn: pulumi.Any(aws_networkfirewall_firewall.Example.Arn),
			LoggingConfiguration: &networkfirewall.LoggingConfigurationLoggingConfigurationArgs{
				LogDestinationConfigs: networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArray{
					&networkfirewall.LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs{
						LogDestination: pulumi.StringMap{
							"deliveryStream": pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Name),
						},
						LogDestinationType: pulumi.String("KinesisDataFirehose"),
						LogType:            pulumi.String("ALERT"),
					},
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
import com.pulumi.aws.networkfirewall.LoggingConfiguration;
import com.pulumi.aws.networkfirewall.LoggingConfigurationArgs;
import com.pulumi.aws.networkfirewall.inputs.LoggingConfigurationLoggingConfigurationArgs;
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
        var example = new LoggingConfiguration("example", LoggingConfigurationArgs.builder()        
            .firewallArn(aws_networkfirewall_firewall.example().arn())
            .loggingConfiguration(LoggingConfigurationLoggingConfigurationArgs.builder()
                .logDestinationConfigs(LoggingConfigurationLoggingConfigurationLogDestinationConfigArgs.builder()
                    .logDestination(Map.of("deliveryStream", aws_kinesis_firehose_delivery_stream.example().name()))
                    .logDestinationType("KinesisDataFirehose")
                    .logType("ALERT")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:networkfirewall:LoggingConfiguration
    properties:
      firewallArn: ${aws_networkfirewall_firewall.example.arn}
      loggingConfiguration:
        logDestinationConfigs:
          - logDestination:
              deliveryStream: ${aws_kinesis_firehose_delivery_stream.example.name}
            logDestinationType: KinesisDataFirehose
            logType: ALERT
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Firewall Logging Configurations using the `firewall_arn`. For example:

```sh
 $ pulumi import aws:networkfirewall/loggingConfiguration:LoggingConfiguration example arn:aws:network-firewall:us-west-1:123456789012:firewall/example
```
 