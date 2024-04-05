Provides a DMS Serverless replication config resource.

> **NOTE:** Changing most arguments will stop the replication if it is running. You can set `start_replication` to resume the replication afterwards.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const name = new aws.dms.ReplicationConfig("name", {
    replicationConfigIdentifier: "test-dms-serverless-replication-tf",
    resourceIdentifier: "test-dms-serverless-replication-tf",
    replicationType: "cdc",
    sourceEndpointArn: aws_dms_endpoint.source.endpoint_arn,
    targetEndpointArn: aws_dms_endpoint.target.endpoint_arn,
    tableMappings: `  {
    "rules":[{"rule-type":"selection","rule-id":"1","rule-name":"1","rule-action":"include","object-locator":{"schema-name":"%%","table-name":"%%"}}]
  }
`,
    startReplication: true,
    computeConfig: {
        replicationSubnetGroupId: aws_dms_replication_subnet_group["default"].replication_subnet_group_id,
        maxCapacityUnits: 64,
        minCapacityUnits: 2,
        preferredMaintenanceWindow: "sun:23:45-mon:00:30",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

name = aws.dms.ReplicationConfig("name",
    replication_config_identifier="test-dms-serverless-replication-tf",
    resource_identifier="test-dms-serverless-replication-tf",
    replication_type="cdc",
    source_endpoint_arn=aws_dms_endpoint["source"]["endpoint_arn"],
    target_endpoint_arn=aws_dms_endpoint["target"]["endpoint_arn"],
    table_mappings="""  {
    "rules":[{"rule-type":"selection","rule-id":"1","rule-name":"1","rule-action":"include","object-locator":{"schema-name":"%%","table-name":"%%"}}]
  }
""",
    start_replication=True,
    compute_config=aws.dms.ReplicationConfigComputeConfigArgs(
        replication_subnet_group_id=aws_dms_replication_subnet_group["default"]["replication_subnet_group_id"],
        max_capacity_units=64,
        min_capacity_units=2,
        preferred_maintenance_window="sun:23:45-mon:00:30",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var name = new Aws.Dms.ReplicationConfig("name", new()
    {
        ReplicationConfigIdentifier = "test-dms-serverless-replication-tf",
        ResourceIdentifier = "test-dms-serverless-replication-tf",
        ReplicationType = "cdc",
        SourceEndpointArn = aws_dms_endpoint.Source.Endpoint_arn,
        TargetEndpointArn = aws_dms_endpoint.Target.Endpoint_arn,
        TableMappings = @"  {
    ""rules"":[{""rule-type"":""selection"",""rule-id"":""1"",""rule-name"":""1"",""rule-action"":""include"",""object-locator"":{""schema-name"":""%%"",""table-name"":""%%""}}]
  }
",
        StartReplication = true,
        ComputeConfig = new Aws.Dms.Inputs.ReplicationConfigComputeConfigArgs
        {
            ReplicationSubnetGroupId = aws_dms_replication_subnet_group.Default.Replication_subnet_group_id,
            MaxCapacityUnits = 64,
            MinCapacityUnits = 2,
            PreferredMaintenanceWindow = "sun:23:45-mon:00:30",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewReplicationConfig(ctx, "name", &dms.ReplicationConfigArgs{
			ReplicationConfigIdentifier: pulumi.String("test-dms-serverless-replication-tf"),
			ResourceIdentifier:          pulumi.String("test-dms-serverless-replication-tf"),
			ReplicationType:             pulumi.String("cdc"),
			SourceEndpointArn:           pulumi.Any(aws_dms_endpoint.Source.Endpoint_arn),
			TargetEndpointArn:           pulumi.Any(aws_dms_endpoint.Target.Endpoint_arn),
			TableMappings:               pulumi.String("  {\n    \"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"rule-action\":\"include\",\"object-locator\":{\"schema-name\":\"%%\",\"table-name\":\"%%\"}}]\n  }\n"),
			StartReplication:            pulumi.Bool(true),
			ComputeConfig: &dms.ReplicationConfigComputeConfigArgs{
				ReplicationSubnetGroupId:   pulumi.Any(aws_dms_replication_subnet_group.Default.Replication_subnet_group_id),
				MaxCapacityUnits:           pulumi.Int(64),
				MinCapacityUnits:           pulumi.Int(2),
				PreferredMaintenanceWindow: pulumi.String("sun:23:45-mon:00:30"),
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
import com.pulumi.aws.dms.ReplicationConfig;
import com.pulumi.aws.dms.ReplicationConfigArgs;
import com.pulumi.aws.dms.inputs.ReplicationConfigComputeConfigArgs;
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
        var name = new ReplicationConfig("name", ReplicationConfigArgs.builder()        
            .replicationConfigIdentifier("test-dms-serverless-replication-tf")
            .resourceIdentifier("test-dms-serverless-replication-tf")
            .replicationType("cdc")
            .sourceEndpointArn(aws_dms_endpoint.source().endpoint_arn())
            .targetEndpointArn(aws_dms_endpoint.target().endpoint_arn())
            .tableMappings("""
  {
    "rules":[{"rule-type":"selection","rule-id":"1","rule-name":"1","rule-action":"include","object-locator":{"schema-name":"%%","table-name":"%%"}}]
  }
            """)
            .startReplication(true)
            .computeConfig(ReplicationConfigComputeConfigArgs.builder()
                .replicationSubnetGroupId(aws_dms_replication_subnet_group.default().replication_subnet_group_id())
                .maxCapacityUnits("64")
                .minCapacityUnits("2")
                .preferredMaintenanceWindow("sun:23:45-mon:00:30")
                .build())
            .build());

    }
}
```
```yaml
resources:
  name:
    type: aws:dms:ReplicationConfig
    properties:
      replicationConfigIdentifier: test-dms-serverless-replication-tf
      resourceIdentifier: test-dms-serverless-replication-tf
      replicationType: cdc
      sourceEndpointArn: ${aws_dms_endpoint.source.endpoint_arn}
      targetEndpointArn: ${aws_dms_endpoint.target.endpoint_arn}
      tableMappings: |2
          {
            "rules":[{"rule-type":"selection","rule-id":"1","rule-name":"1","rule-action":"include","object-locator":{"schema-name":"%%","table-name":"%%"}}]
          }
      startReplication: true
      computeConfig:
        replicationSubnetGroupId: ${aws_dms_replication_subnet_group.default.replication_subnet_group_id}
        maxCapacityUnits: '64'
        minCapacityUnits: '2'
        preferredMaintenanceWindow: sun:23:45-mon:00:30
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a replication config using the `arn`. For example:

```sh
 $ pulumi import aws:dms/replicationConfig:ReplicationConfig example arn:aws:dms:us-east-1:123456789012:replication-config:UX6OL6MHMMJKFFOXE3H7LLJCMEKBDUG4ZV7DRSI
```
 