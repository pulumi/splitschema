Provides an Amazon MSK Connect Connector resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mskconnect.Connector("example", {
    kafkaconnectVersion: "2.7.1",
    capacity: {
        autoscaling: {
            mcuCount: 1,
            minWorkerCount: 1,
            maxWorkerCount: 2,
            scaleInPolicy: {
                cpuUtilizationPercentage: 20,
            },
            scaleOutPolicy: {
                cpuUtilizationPercentage: 80,
            },
        },
    },
    connectorConfiguration: {
        "connector.class": "com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector",
        "tasks.max": "1",
        topics: "example",
    },
    kafkaCluster: {
        apacheKafkaCluster: {
            bootstrapServers: aws_msk_cluster.example.bootstrap_brokers_tls,
            vpc: {
                securityGroups: [aws_security_group.example.id],
                subnets: [
                    aws_subnet.example1.id,
                    aws_subnet.example2.id,
                    aws_subnet.example3.id,
                ],
            },
        },
    },
    kafkaClusterClientAuthentication: {
        authenticationType: "NONE",
    },
    kafkaClusterEncryptionInTransit: {
        encryptionType: "TLS",
    },
    plugins: [{
        customPlugin: {
            arn: aws_mskconnect_custom_plugin.example.arn,
            revision: aws_mskconnect_custom_plugin.example.latest_revision,
        },
    }],
    serviceExecutionRoleArn: aws_iam_role.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mskconnect.Connector("example",
    kafkaconnect_version="2.7.1",
    capacity=aws.mskconnect.ConnectorCapacityArgs(
        autoscaling=aws.mskconnect.ConnectorCapacityAutoscalingArgs(
            mcu_count=1,
            min_worker_count=1,
            max_worker_count=2,
            scale_in_policy=aws.mskconnect.ConnectorCapacityAutoscalingScaleInPolicyArgs(
                cpu_utilization_percentage=20,
            ),
            scale_out_policy=aws.mskconnect.ConnectorCapacityAutoscalingScaleOutPolicyArgs(
                cpu_utilization_percentage=80,
            ),
        ),
    ),
    connector_configuration={
        "connector.class": "com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector",
        "tasks.max": "1",
        "topics": "example",
    },
    kafka_cluster=aws.mskconnect.ConnectorKafkaClusterArgs(
        apache_kafka_cluster=aws.mskconnect.ConnectorKafkaClusterApacheKafkaClusterArgs(
            bootstrap_servers=aws_msk_cluster["example"]["bootstrap_brokers_tls"],
            vpc=aws.mskconnect.ConnectorKafkaClusterApacheKafkaClusterVpcArgs(
                security_groups=[aws_security_group["example"]["id"]],
                subnets=[
                    aws_subnet["example1"]["id"],
                    aws_subnet["example2"]["id"],
                    aws_subnet["example3"]["id"],
                ],
            ),
        ),
    ),
    kafka_cluster_client_authentication=aws.mskconnect.ConnectorKafkaClusterClientAuthenticationArgs(
        authentication_type="NONE",
    ),
    kafka_cluster_encryption_in_transit=aws.mskconnect.ConnectorKafkaClusterEncryptionInTransitArgs(
        encryption_type="TLS",
    ),
    plugins=[aws.mskconnect.ConnectorPluginArgs(
        custom_plugin=aws.mskconnect.ConnectorPluginCustomPluginArgs(
            arn=aws_mskconnect_custom_plugin["example"]["arn"],
            revision=aws_mskconnect_custom_plugin["example"]["latest_revision"],
        ),
    )],
    service_execution_role_arn=aws_iam_role["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.MskConnect.Connector("example", new()
    {
        KafkaconnectVersion = "2.7.1",
        Capacity = new Aws.MskConnect.Inputs.ConnectorCapacityArgs
        {
            Autoscaling = new Aws.MskConnect.Inputs.ConnectorCapacityAutoscalingArgs
            {
                McuCount = 1,
                MinWorkerCount = 1,
                MaxWorkerCount = 2,
                ScaleInPolicy = new Aws.MskConnect.Inputs.ConnectorCapacityAutoscalingScaleInPolicyArgs
                {
                    CpuUtilizationPercentage = 20,
                },
                ScaleOutPolicy = new Aws.MskConnect.Inputs.ConnectorCapacityAutoscalingScaleOutPolicyArgs
                {
                    CpuUtilizationPercentage = 80,
                },
            },
        },
        ConnectorConfiguration = 
        {
            { "connector.class", "com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector" },
            { "tasks.max", "1" },
            { "topics", "example" },
        },
        KafkaCluster = new Aws.MskConnect.Inputs.ConnectorKafkaClusterArgs
        {
            ApacheKafkaCluster = new Aws.MskConnect.Inputs.ConnectorKafkaClusterApacheKafkaClusterArgs
            {
                BootstrapServers = aws_msk_cluster.Example.Bootstrap_brokers_tls,
                Vpc = new Aws.MskConnect.Inputs.ConnectorKafkaClusterApacheKafkaClusterVpcArgs
                {
                    SecurityGroups = new[]
                    {
                        aws_security_group.Example.Id,
                    },
                    Subnets = new[]
                    {
                        aws_subnet.Example1.Id,
                        aws_subnet.Example2.Id,
                        aws_subnet.Example3.Id,
                    },
                },
            },
        },
        KafkaClusterClientAuthentication = new Aws.MskConnect.Inputs.ConnectorKafkaClusterClientAuthenticationArgs
        {
            AuthenticationType = "NONE",
        },
        KafkaClusterEncryptionInTransit = new Aws.MskConnect.Inputs.ConnectorKafkaClusterEncryptionInTransitArgs
        {
            EncryptionType = "TLS",
        },
        Plugins = new[]
        {
            new Aws.MskConnect.Inputs.ConnectorPluginArgs
            {
                CustomPlugin = new Aws.MskConnect.Inputs.ConnectorPluginCustomPluginArgs
                {
                    Arn = aws_mskconnect_custom_plugin.Example.Arn,
                    Revision = aws_mskconnect_custom_plugin.Example.Latest_revision,
                },
            },
        },
        ServiceExecutionRoleArn = aws_iam_role.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mskconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mskconnect.NewConnector(ctx, "example", &mskconnect.ConnectorArgs{
			KafkaconnectVersion: pulumi.String("2.7.1"),
			Capacity: &mskconnect.ConnectorCapacityArgs{
				Autoscaling: &mskconnect.ConnectorCapacityAutoscalingArgs{
					McuCount:       pulumi.Int(1),
					MinWorkerCount: pulumi.Int(1),
					MaxWorkerCount: pulumi.Int(2),
					ScaleInPolicy: &mskconnect.ConnectorCapacityAutoscalingScaleInPolicyArgs{
						CpuUtilizationPercentage: pulumi.Int(20),
					},
					ScaleOutPolicy: &mskconnect.ConnectorCapacityAutoscalingScaleOutPolicyArgs{
						CpuUtilizationPercentage: pulumi.Int(80),
					},
				},
			},
			ConnectorConfiguration: pulumi.StringMap{
				"connector.class": pulumi.String("com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector"),
				"tasks.max":       pulumi.String("1"),
				"topics":          pulumi.String("example"),
			},
			KafkaCluster: &mskconnect.ConnectorKafkaClusterArgs{
				ApacheKafkaCluster: &mskconnect.ConnectorKafkaClusterApacheKafkaClusterArgs{
					BootstrapServers: pulumi.Any(aws_msk_cluster.Example.Bootstrap_brokers_tls),
					Vpc: &mskconnect.ConnectorKafkaClusterApacheKafkaClusterVpcArgs{
						SecurityGroups: pulumi.StringArray{
							aws_security_group.Example.Id,
						},
						Subnets: pulumi.StringArray{
							aws_subnet.Example1.Id,
							aws_subnet.Example2.Id,
							aws_subnet.Example3.Id,
						},
					},
				},
			},
			KafkaClusterClientAuthentication: &mskconnect.ConnectorKafkaClusterClientAuthenticationArgs{
				AuthenticationType: pulumi.String("NONE"),
			},
			KafkaClusterEncryptionInTransit: &mskconnect.ConnectorKafkaClusterEncryptionInTransitArgs{
				EncryptionType: pulumi.String("TLS"),
			},
			Plugins: mskconnect.ConnectorPluginArray{
				&mskconnect.ConnectorPluginArgs{
					CustomPlugin: &mskconnect.ConnectorPluginCustomPluginArgs{
						Arn:      pulumi.Any(aws_mskconnect_custom_plugin.Example.Arn),
						Revision: pulumi.Any(aws_mskconnect_custom_plugin.Example.Latest_revision),
					},
				},
			},
			ServiceExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.mskconnect.Connector;
import com.pulumi.aws.mskconnect.ConnectorArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorCapacityArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorCapacityAutoscalingArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorCapacityAutoscalingScaleInPolicyArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorCapacityAutoscalingScaleOutPolicyArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorKafkaClusterArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorKafkaClusterApacheKafkaClusterArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorKafkaClusterApacheKafkaClusterVpcArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorKafkaClusterClientAuthenticationArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorKafkaClusterEncryptionInTransitArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorPluginArgs;
import com.pulumi.aws.mskconnect.inputs.ConnectorPluginCustomPluginArgs;
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
        var example = new Connector("example", ConnectorArgs.builder()        
            .kafkaconnectVersion("2.7.1")
            .capacity(ConnectorCapacityArgs.builder()
                .autoscaling(ConnectorCapacityAutoscalingArgs.builder()
                    .mcuCount(1)
                    .minWorkerCount(1)
                    .maxWorkerCount(2)
                    .scaleInPolicy(ConnectorCapacityAutoscalingScaleInPolicyArgs.builder()
                        .cpuUtilizationPercentage(20)
                        .build())
                    .scaleOutPolicy(ConnectorCapacityAutoscalingScaleOutPolicyArgs.builder()
                        .cpuUtilizationPercentage(80)
                        .build())
                    .build())
                .build())
            .connectorConfiguration(Map.ofEntries(
                Map.entry("connector.class", "com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector"),
                Map.entry("tasks.max", "1"),
                Map.entry("topics", "example")
            ))
            .kafkaCluster(ConnectorKafkaClusterArgs.builder()
                .apacheKafkaCluster(ConnectorKafkaClusterApacheKafkaClusterArgs.builder()
                    .bootstrapServers(aws_msk_cluster.example().bootstrap_brokers_tls())
                    .vpc(ConnectorKafkaClusterApacheKafkaClusterVpcArgs.builder()
                        .securityGroups(aws_security_group.example().id())
                        .subnets(                        
                            aws_subnet.example1().id(),
                            aws_subnet.example2().id(),
                            aws_subnet.example3().id())
                        .build())
                    .build())
                .build())
            .kafkaClusterClientAuthentication(ConnectorKafkaClusterClientAuthenticationArgs.builder()
                .authenticationType("NONE")
                .build())
            .kafkaClusterEncryptionInTransit(ConnectorKafkaClusterEncryptionInTransitArgs.builder()
                .encryptionType("TLS")
                .build())
            .plugins(ConnectorPluginArgs.builder()
                .customPlugin(ConnectorPluginCustomPluginArgs.builder()
                    .arn(aws_mskconnect_custom_plugin.example().arn())
                    .revision(aws_mskconnect_custom_plugin.example().latest_revision())
                    .build())
                .build())
            .serviceExecutionRoleArn(aws_iam_role.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:mskconnect:Connector
    properties:
      kafkaconnectVersion: 2.7.1
      capacity:
        autoscaling:
          mcuCount: 1
          minWorkerCount: 1
          maxWorkerCount: 2
          scaleInPolicy:
            cpuUtilizationPercentage: 20
          scaleOutPolicy:
            cpuUtilizationPercentage: 80
      connectorConfiguration:
        connector.class: com.github.jcustenborder.kafka.connect.simulator.SimulatorSinkConnector
        tasks.max: '1'
        topics: example
      kafkaCluster:
        apacheKafkaCluster:
          bootstrapServers: ${aws_msk_cluster.example.bootstrap_brokers_tls}
          vpc:
            securityGroups:
              - ${aws_security_group.example.id}
            subnets:
              - ${aws_subnet.example1.id}
              - ${aws_subnet.example2.id}
              - ${aws_subnet.example3.id}
      kafkaClusterClientAuthentication:
        authenticationType: NONE
      kafkaClusterEncryptionInTransit:
        encryptionType: TLS
      plugins:
        - customPlugin:
            arn: ${aws_mskconnect_custom_plugin.example.arn}
            revision: ${aws_mskconnect_custom_plugin.example.latest_revision}
      serviceExecutionRoleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MSK Connect Connector using the connector's `arn`. For example:

```sh
 $ pulumi import aws:mskconnect/connector:Connector example 'arn:aws:kafkaconnect:eu-central-1:123456789012:connector/example/264edee4-17a3-412e-bd76-6681cfc93805-3'
```
 