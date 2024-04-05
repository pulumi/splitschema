Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.

> **NOTE:** Changing most arguments will stop the task if it is running. You can set `start_replication_task` to resume the task afterwards.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new replication task
const test = new aws.dms.ReplicationTask("test", {
    cdcStartTime: "1484346880",
    migrationType: "full-load",
    replicationInstanceArn: aws_dms_replication_instance["test-dms-replication-instance-tf"].replication_instance_arn,
    replicationTaskId: "test-dms-replication-task-tf",
    replicationTaskSettings: "...",
    sourceEndpointArn: aws_dms_endpoint["test-dms-source-endpoint-tf"].endpoint_arn,
    tableMappings: "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
    tags: {
        Name: "test",
    },
    targetEndpointArn: aws_dms_endpoint["test-dms-target-endpoint-tf"].endpoint_arn,
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new replication task
test = aws.dms.ReplicationTask("test",
    cdc_start_time="1484346880",
    migration_type="full-load",
    replication_instance_arn=aws_dms_replication_instance["test-dms-replication-instance-tf"]["replication_instance_arn"],
    replication_task_id="test-dms-replication-task-tf",
    replication_task_settings="...",
    source_endpoint_arn=aws_dms_endpoint["test-dms-source-endpoint-tf"]["endpoint_arn"],
    table_mappings="{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
    tags={
        "Name": "test",
    },
    target_endpoint_arn=aws_dms_endpoint["test-dms-target-endpoint-tf"]["endpoint_arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new replication task
    var test = new Aws.Dms.ReplicationTask("test", new()
    {
        CdcStartTime = "1484346880",
        MigrationType = "full-load",
        ReplicationInstanceArn = aws_dms_replication_instance.Test_dms_replication_instance_tf.Replication_instance_arn,
        ReplicationTaskId = "test-dms-replication-task-tf",
        ReplicationTaskSettings = "...",
        SourceEndpointArn = aws_dms_endpoint.Test_dms_source_endpoint_tf.Endpoint_arn,
        TableMappings = "{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}",
        Tags = 
        {
            { "Name", "test" },
        },
        TargetEndpointArn = aws_dms_endpoint.Test_dms_target_endpoint_tf.Endpoint_arn,
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
		_, err := dms.NewReplicationTask(ctx, "test", &dms.ReplicationTaskArgs{
			CdcStartTime:            pulumi.String("1484346880"),
			MigrationType:           pulumi.String("full-load"),
			ReplicationInstanceArn:  pulumi.Any(aws_dms_replication_instance.TestDmsReplicationInstanceTf.Replication_instance_arn),
			ReplicationTaskId:       pulumi.String("test-dms-replication-task-tf"),
			ReplicationTaskSettings: pulumi.String("..."),
			SourceEndpointArn:       pulumi.Any(aws_dms_endpoint.TestDmsSourceEndpointTf.Endpoint_arn),
			TableMappings:           pulumi.String("{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
			TargetEndpointArn: pulumi.Any(aws_dms_endpoint.TestDmsTargetEndpointTf.Endpoint_arn),
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
import com.pulumi.aws.dms.ReplicationTask;
import com.pulumi.aws.dms.ReplicationTaskArgs;
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
        var test = new ReplicationTask("test", ReplicationTaskArgs.builder()        
            .cdcStartTime(1484346880)
            .migrationType("full-load")
            .replicationInstanceArn(aws_dms_replication_instance.test-dms-replication-instance-tf().replication_instance_arn())
            .replicationTaskId("test-dms-replication-task-tf")
            .replicationTaskSettings("...")
            .sourceEndpointArn(aws_dms_endpoint.test-dms-source-endpoint-tf().endpoint_arn())
            .tableMappings("{\"rules\":[{\"rule-type\":\"selection\",\"rule-id\":\"1\",\"rule-name\":\"1\",\"object-locator\":{\"schema-name\":\"%\",\"table-name\":\"%\"},\"rule-action\":\"include\"}]}")
            .tags(Map.of("Name", "test"))
            .targetEndpointArn(aws_dms_endpoint.test-dms-target-endpoint-tf().endpoint_arn())
            .build());

    }
}
```
```yaml
resources:
  # Create a new replication task
  test:
    type: aws:dms:ReplicationTask
    properties:
      cdcStartTime: 1.48434688e+09
      migrationType: full-load
      replicationInstanceArn: ${aws_dms_replication_instance"test-dms-replication-instance-tf"[%!s(MISSING)].replication_instance_arn}
      replicationTaskId: test-dms-replication-task-tf
      replicationTaskSettings: '...'
      sourceEndpointArn: ${aws_dms_endpoint"test-dms-source-endpoint-tf"[%!s(MISSING)].endpoint_arn}
      tableMappings: '{"rules":[{"rule-type":"selection","rule-id":"1","rule-name":"1","object-locator":{"schema-name":"%","table-name":"%"},"rule-action":"include"}]}'
      tags:
        Name: test
      targetEndpointArn: ${aws_dms_endpoint"test-dms-target-endpoint-tf"[%!s(MISSING)].endpoint_arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import replication tasks using the `replication_task_id`. For example:

```sh
 $ pulumi import aws:dms/replicationTask:ReplicationTask test test-dms-replication-task-tf
```
 