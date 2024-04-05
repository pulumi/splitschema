Provides an ElastiCache Replication Group resource.

For working with a [Memcached cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/WhatIs.html) or a
[single-node Redis instance (Cluster Mode Disabled)](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/WhatIs.html),
see the `aws.elasticache.Cluster` resource.

> **Note:** When you change an attribute, such as `engine_version`, by
default the ElastiCache API applies it in the next maintenance window. Because
of this, this provider may report a difference in its planning phase because the
actual modification has not yet taken place. You can use the
`apply_immediately` flag to instruct the service to apply the change
immediately. Using `apply_immediately` can result in a brief downtime as
servers reboots.
See the AWS Documentation on
[Modifying an ElastiCache Cache Cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.Modify.html)
for more information.

> **Note:** Any attribute changes that re-create the resource will be applied immediately, regardless of the value of `apply_immediately`.

> **Note:** Be aware of the terminology collision around "cluster" for `aws.elasticache.ReplicationGroup`. For example, it is possible to create a ["Cluster Mode Disabled [Redis] Cluster"](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.Create.CON.Redis.html). With "Cluster Mode Enabled", the data will be stored in shards (called "node groups"). See [Redis Cluster Configuration](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/cluster-create-determine-requirements.html#redis-cluster-configuration) for a diagram of the differences. To enable cluster mode, use a parameter group that has cluster mode enabled. The default parameter groups provided by AWS end with ".cluster.on", for example `default.redis6.x.cluster.on`.

{{% examples %}}
## Example Usage
{{% example %}}
### Redis Cluster Mode Disabled

To create a single shard primary with single read replica:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.ReplicationGroup("example", {
    automaticFailoverEnabled: true,
    description: "example description",
    nodeType: "cache.m4.large",
    numCacheClusters: 2,
    parameterGroupName: "default.redis3.2",
    port: 6379,
    preferredCacheClusterAzs: [
        "us-west-2a",
        "us-west-2b",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticache.ReplicationGroup("example",
    automatic_failover_enabled=True,
    description="example description",
    node_type="cache.m4.large",
    num_cache_clusters=2,
    parameter_group_name="default.redis3.2",
    port=6379,
    preferred_cache_cluster_azs=[
        "us-west-2a",
        "us-west-2b",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElastiCache.ReplicationGroup("example", new()
    {
        AutomaticFailoverEnabled = true,
        Description = "example description",
        NodeType = "cache.m4.large",
        NumCacheClusters = 2,
        ParameterGroupName = "default.redis3.2",
        Port = 6379,
        PreferredCacheClusterAzs = new[]
        {
            "us-west-2a",
            "us-west-2b",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "example", &elasticache.ReplicationGroupArgs{
			AutomaticFailoverEnabled: pulumi.Bool(true),
			Description:              pulumi.String("example description"),
			NodeType:                 pulumi.String("cache.m4.large"),
			NumCacheClusters:         pulumi.Int(2),
			ParameterGroupName:       pulumi.String("default.redis3.2"),
			Port:                     pulumi.Int(6379),
			PreferredCacheClusterAzs: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
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
import com.pulumi.aws.elasticache.ReplicationGroup;
import com.pulumi.aws.elasticache.ReplicationGroupArgs;
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
        var example = new ReplicationGroup("example", ReplicationGroupArgs.builder()        
            .automaticFailoverEnabled(true)
            .description("example description")
            .nodeType("cache.m4.large")
            .numCacheClusters(2)
            .parameterGroupName("default.redis3.2")
            .port(6379)
            .preferredCacheClusterAzs(            
                "us-west-2a",
                "us-west-2b")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:ReplicationGroup
    properties:
      automaticFailoverEnabled: true
      description: example description
      nodeType: cache.m4.large
      numCacheClusters: 2
      parameterGroupName: default.redis3.2
      port: 6379
      preferredCacheClusterAzs:
        - us-west-2a
        - us-west-2b
```

You have two options for adjusting the number of replicas:

* Adjusting `num_cache_clusters` directly. This will attempt to automatically add or remove replicas, but provides no granular control (e.g., preferred availability zone, cache cluster ID) for the added or removed replicas. This also currently expects cache cluster IDs in the form of `replication_group_id-00#`.
* Otherwise for fine grained control of the underlying cache clusters, they can be added or removed with the `aws.elasticache.Cluster` resource and its `replication_group_id` attribute. In this situation, you will need to utilize [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to prevent perpetual differences with the `number_cache_cluster` attribute.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.ReplicationGroup("example", {
    automaticFailoverEnabled: true,
    preferredCacheClusterAzs: [
        "us-west-2a",
        "us-west-2b",
    ],
    description: "example description",
    nodeType: "cache.m4.large",
    numCacheClusters: 2,
    parameterGroupName: "default.redis3.2",
    port: 6379,
});
let replica: aws.elasticache.Cluster | undefined;
if (1 == true) {
    replica = new aws.elasticache.Cluster("replica", {replicationGroupId: example.id});
}
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticache.ReplicationGroup("example",
    automatic_failover_enabled=True,
    preferred_cache_cluster_azs=[
        "us-west-2a",
        "us-west-2b",
    ],
    description="example description",
    node_type="cache.m4.large",
    num_cache_clusters=2,
    parameter_group_name="default.redis3.2",
    port=6379)
replica = None
if 1 == True:
    replica = aws.elasticache.Cluster("replica", replication_group_id=example.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElastiCache.ReplicationGroup("example", new()
    {
        AutomaticFailoverEnabled = true,
        PreferredCacheClusterAzs = new[]
        {
            "us-west-2a",
            "us-west-2b",
        },
        Description = "example description",
        NodeType = "cache.m4.large",
        NumCacheClusters = 2,
        ParameterGroupName = "default.redis3.2",
        Port = 6379,
    });

    var replica = new List<Aws.ElastiCache.Cluster>();
    for (var rangeIndex = 0; rangeIndex < (1 == true); rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        replica.Add(new Aws.ElastiCache.Cluster($"replica-{range.Value}", new()
        {
            ReplicationGroupId = example.Id,
        }));
    }
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := elasticache.NewReplicationGroup(ctx, "example", &elasticache.ReplicationGroupArgs{
			AutomaticFailoverEnabled: pulumi.Bool(true),
			PreferredCacheClusterAzs: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
			},
			Description:        pulumi.String("example description"),
			NodeType:           pulumi.String("cache.m4.large"),
			NumCacheClusters:   pulumi.Int(2),
			ParameterGroupName: pulumi.String("default.redis3.2"),
			Port:               pulumi.Int(6379),
		})
		if err != nil {
			return err
		}
		var replica []*elasticache.Cluster
		for index := 0; index < 1 == true; index++ {
			key0 := index
			_ := index
			__res, err := elasticache.NewCluster(ctx, fmt.Sprintf("replica-%v", key0), &elasticache.ClusterArgs{
				ReplicationGroupId: example.ID(),
			})
			if err != nil {
				return err
			}
			replica = append(replica, __res)
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
import com.pulumi.aws.elasticache.ReplicationGroup;
import com.pulumi.aws.elasticache.ReplicationGroupArgs;
import com.pulumi.aws.elasticache.Cluster;
import com.pulumi.aws.elasticache.ClusterArgs;
import com.pulumi.codegen.internal.KeyedValue;
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
        var example = new ReplicationGroup("example", ReplicationGroupArgs.builder()        
            .automaticFailoverEnabled(true)
            .preferredCacheClusterAzs(            
                "us-west-2a",
                "us-west-2b")
            .description("example description")
            .nodeType("cache.m4.large")
            .numCacheClusters(2)
            .parameterGroupName("default.redis3.2")
            .port(6379)
            .build());

        for (var i = 0; i < (1 == true); i++) {
            new Cluster("replica-" + i, ClusterArgs.builder()            
                .replicationGroupId(example.id())
                .build());

        
}
    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:ReplicationGroup
    properties:
      automaticFailoverEnabled: true
      preferredCacheClusterAzs:
        - us-west-2a
        - us-west-2b
      description: example description
      nodeType: cache.m4.large
      numCacheClusters: 2
      parameterGroupName: default.redis3.2
      port: 6379
  replica:
    type: aws:elasticache:Cluster
    properties:
      replicationGroupId: ${example.id}
    options: {}
```
{{% /example %}}
{{% example %}}
### Redis Cluster Mode Enabled

To create two shards with a primary and a single read replica each:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const baz = new aws.elasticache.ReplicationGroup("baz", {
    automaticFailoverEnabled: true,
    description: "example description",
    nodeType: "cache.t2.small",
    numNodeGroups: 2,
    parameterGroupName: "default.redis3.2.cluster.on",
    port: 6379,
    replicasPerNodeGroup: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

baz = aws.elasticache.ReplicationGroup("baz",
    automatic_failover_enabled=True,
    description="example description",
    node_type="cache.t2.small",
    num_node_groups=2,
    parameter_group_name="default.redis3.2.cluster.on",
    port=6379,
    replicas_per_node_group=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var baz = new Aws.ElastiCache.ReplicationGroup("baz", new()
    {
        AutomaticFailoverEnabled = true,
        Description = "example description",
        NodeType = "cache.t2.small",
        NumNodeGroups = 2,
        ParameterGroupName = "default.redis3.2.cluster.on",
        Port = 6379,
        ReplicasPerNodeGroup = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "baz", &elasticache.ReplicationGroupArgs{
			AutomaticFailoverEnabled: pulumi.Bool(true),
			Description:              pulumi.String("example description"),
			NodeType:                 pulumi.String("cache.t2.small"),
			NumNodeGroups:            pulumi.Int(2),
			ParameterGroupName:       pulumi.String("default.redis3.2.cluster.on"),
			Port:                     pulumi.Int(6379),
			ReplicasPerNodeGroup:     pulumi.Int(1),
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
import com.pulumi.aws.elasticache.ReplicationGroup;
import com.pulumi.aws.elasticache.ReplicationGroupArgs;
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
        var baz = new ReplicationGroup("baz", ReplicationGroupArgs.builder()        
            .automaticFailoverEnabled(true)
            .description("example description")
            .nodeType("cache.t2.small")
            .numNodeGroups(2)
            .parameterGroupName("default.redis3.2.cluster.on")
            .port(6379)
            .replicasPerNodeGroup(1)
            .build());

    }
}
```
```yaml
resources:
  baz:
    type: aws:elasticache:ReplicationGroup
    properties:
      automaticFailoverEnabled: true
      description: example description
      nodeType: cache.t2.small
      numNodeGroups: 2
      parameterGroupName: default.redis3.2.cluster.on
      port: 6379
      replicasPerNodeGroup: 1
```
{{% /example %}}
{{% example %}}
### Redis Log Delivery configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.elasticache.ReplicationGroup("test", {
    description: "test description",
    nodeType: "cache.t3.small",
    port: 6379,
    applyImmediately: true,
    autoMinorVersionUpgrade: false,
    maintenanceWindow: "tue:06:30-tue:07:30",
    snapshotWindow: "01:00-02:00",
    logDeliveryConfigurations: [
        {
            destination: aws_cloudwatch_log_group.example.name,
            destinationType: "cloudwatch-logs",
            logFormat: "text",
            logType: "slow-log",
        },
        {
            destination: aws_kinesis_firehose_delivery_stream.example.name,
            destinationType: "kinesis-firehose",
            logFormat: "json",
            logType: "engine-log",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.elasticache.ReplicationGroup("test",
    description="test description",
    node_type="cache.t3.small",
    port=6379,
    apply_immediately=True,
    auto_minor_version_upgrade=False,
    maintenance_window="tue:06:30-tue:07:30",
    snapshot_window="01:00-02:00",
    log_delivery_configurations=[
        aws.elasticache.ReplicationGroupLogDeliveryConfigurationArgs(
            destination=aws_cloudwatch_log_group["example"]["name"],
            destination_type="cloudwatch-logs",
            log_format="text",
            log_type="slow-log",
        ),
        aws.elasticache.ReplicationGroupLogDeliveryConfigurationArgs(
            destination=aws_kinesis_firehose_delivery_stream["example"]["name"],
            destination_type="kinesis-firehose",
            log_format="json",
            log_type="engine-log",
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
    var test = new Aws.ElastiCache.ReplicationGroup("test", new()
    {
        Description = "test description",
        NodeType = "cache.t3.small",
        Port = 6379,
        ApplyImmediately = true,
        AutoMinorVersionUpgrade = false,
        MaintenanceWindow = "tue:06:30-tue:07:30",
        SnapshotWindow = "01:00-02:00",
        LogDeliveryConfigurations = new[]
        {
            new Aws.ElastiCache.Inputs.ReplicationGroupLogDeliveryConfigurationArgs
            {
                Destination = aws_cloudwatch_log_group.Example.Name,
                DestinationType = "cloudwatch-logs",
                LogFormat = "text",
                LogType = "slow-log",
            },
            new Aws.ElastiCache.Inputs.ReplicationGroupLogDeliveryConfigurationArgs
            {
                Destination = aws_kinesis_firehose_delivery_stream.Example.Name,
                DestinationType = "kinesis-firehose",
                LogFormat = "json",
                LogType = "engine-log",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "test", &elasticache.ReplicationGroupArgs{
			Description:             pulumi.String("test description"),
			NodeType:                pulumi.String("cache.t3.small"),
			Port:                    pulumi.Int(6379),
			ApplyImmediately:        pulumi.Bool(true),
			AutoMinorVersionUpgrade: pulumi.Bool(false),
			MaintenanceWindow:       pulumi.String("tue:06:30-tue:07:30"),
			SnapshotWindow:          pulumi.String("01:00-02:00"),
			LogDeliveryConfigurations: elasticache.ReplicationGroupLogDeliveryConfigurationArray{
				&elasticache.ReplicationGroupLogDeliveryConfigurationArgs{
					Destination:     pulumi.Any(aws_cloudwatch_log_group.Example.Name),
					DestinationType: pulumi.String("cloudwatch-logs"),
					LogFormat:       pulumi.String("text"),
					LogType:         pulumi.String("slow-log"),
				},
				&elasticache.ReplicationGroupLogDeliveryConfigurationArgs{
					Destination:     pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Name),
					DestinationType: pulumi.String("kinesis-firehose"),
					LogFormat:       pulumi.String("json"),
					LogType:         pulumi.String("engine-log"),
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
import com.pulumi.aws.elasticache.ReplicationGroup;
import com.pulumi.aws.elasticache.ReplicationGroupArgs;
import com.pulumi.aws.elasticache.inputs.ReplicationGroupLogDeliveryConfigurationArgs;
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
        var test = new ReplicationGroup("test", ReplicationGroupArgs.builder()        
            .description("test description")
            .nodeType("cache.t3.small")
            .port(6379)
            .applyImmediately(true)
            .autoMinorVersionUpgrade(false)
            .maintenanceWindow("tue:06:30-tue:07:30")
            .snapshotWindow("01:00-02:00")
            .logDeliveryConfigurations(            
                ReplicationGroupLogDeliveryConfigurationArgs.builder()
                    .destination(aws_cloudwatch_log_group.example().name())
                    .destinationType("cloudwatch-logs")
                    .logFormat("text")
                    .logType("slow-log")
                    .build(),
                ReplicationGroupLogDeliveryConfigurationArgs.builder()
                    .destination(aws_kinesis_firehose_delivery_stream.example().name())
                    .destinationType("kinesis-firehose")
                    .logFormat("json")
                    .logType("engine-log")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: test description
      nodeType: cache.t3.small
      port: 6379
      applyImmediately: true
      autoMinorVersionUpgrade: false
      maintenanceWindow: tue:06:30-tue:07:30
      snapshotWindow: 01:00-02:00
      logDeliveryConfigurations:
        - destination: ${aws_cloudwatch_log_group.example.name}
          destinationType: cloudwatch-logs
          logFormat: text
          logType: slow-log
        - destination: ${aws_kinesis_firehose_delivery_stream.example.name}
          destinationType: kinesis-firehose
          logFormat: json
          logType: engine-log
```

> **Note:** We currently do not support passing a `primary_cluster_id` in order to create the Replication Group.

> **Note:** Automatic Failover is unavailable for Redis versions earlier than 2.8.6,
and unavailable on T1 node types. For T2 node types, it is only available on Redis version 3.2.4 or later with cluster mode enabled. See the [High Availability Using Replication Groups](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Replication.html) guide
for full details on using Replication Groups.
{{% /example %}}
{{% example %}}
### Creating a secondary replication group for a global replication group

A Global Replication Group can have one one two secondary Replication Groups in different regions. These are added to an existing Global Replication Group.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.elasticache.ReplicationGroup("primary", {
    description: "primary replication group",
    engine: "redis",
    engineVersion: "5.0.6",
    nodeType: "cache.m5.large",
    numCacheClusters: 1,
}, {
    provider: aws.other_region,
});
const example = new aws.elasticache.GlobalReplicationGroup("example", {
    globalReplicationGroupIdSuffix: "example",
    primaryReplicationGroupId: primary.id,
}, {
    provider: aws.other_region,
});
const secondary = new aws.elasticache.ReplicationGroup("secondary", {
    description: "secondary replication group",
    globalReplicationGroupId: example.globalReplicationGroupId,
    numCacheClusters: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.elasticache.ReplicationGroup("primary",
    description="primary replication group",
    engine="redis",
    engine_version="5.0.6",
    node_type="cache.m5.large",
    num_cache_clusters=1,
    opts=pulumi.ResourceOptions(provider=aws["other_region"]))
example = aws.elasticache.GlobalReplicationGroup("example",
    global_replication_group_id_suffix="example",
    primary_replication_group_id=primary.id,
    opts=pulumi.ResourceOptions(provider=aws["other_region"]))
secondary = aws.elasticache.ReplicationGroup("secondary",
    description="secondary replication group",
    global_replication_group_id=example.global_replication_group_id,
    num_cache_clusters=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primary = new Aws.ElastiCache.ReplicationGroup("primary", new()
    {
        Description = "primary replication group",
        Engine = "redis",
        EngineVersion = "5.0.6",
        NodeType = "cache.m5.large",
        NumCacheClusters = 1,
    }, new CustomResourceOptions
    {
        Provider = aws.Other_region,
    });

    var example = new Aws.ElastiCache.GlobalReplicationGroup("example", new()
    {
        GlobalReplicationGroupIdSuffix = "example",
        PrimaryReplicationGroupId = primary.Id,
    }, new CustomResourceOptions
    {
        Provider = aws.Other_region,
    });

    var secondary = new Aws.ElastiCache.ReplicationGroup("secondary", new()
    {
        Description = "secondary replication group",
        GlobalReplicationGroupId = example.GlobalReplicationGroupId,
        NumCacheClusters = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		primary, err := elasticache.NewReplicationGroup(ctx, "primary", &elasticache.ReplicationGroupArgs{
			Description:      pulumi.String("primary replication group"),
			Engine:           pulumi.String("redis"),
			EngineVersion:    pulumi.String("5.0.6"),
			NodeType:         pulumi.String("cache.m5.large"),
			NumCacheClusters: pulumi.Int(1),
		}, pulumi.Provider(aws.Other_region))
		if err != nil {
			return err
		}
		example, err := elasticache.NewGlobalReplicationGroup(ctx, "example", &elasticache.GlobalReplicationGroupArgs{
			GlobalReplicationGroupIdSuffix: pulumi.String("example"),
			PrimaryReplicationGroupId:      primary.ID(),
		}, pulumi.Provider(aws.Other_region))
		if err != nil {
			return err
		}
		_, err = elasticache.NewReplicationGroup(ctx, "secondary", &elasticache.ReplicationGroupArgs{
			Description:              pulumi.String("secondary replication group"),
			GlobalReplicationGroupId: example.GlobalReplicationGroupId,
			NumCacheClusters:         pulumi.Int(1),
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
import com.pulumi.aws.elasticache.ReplicationGroup;
import com.pulumi.aws.elasticache.ReplicationGroupArgs;
import com.pulumi.aws.elasticache.GlobalReplicationGroup;
import com.pulumi.aws.elasticache.GlobalReplicationGroupArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var primary = new ReplicationGroup("primary", ReplicationGroupArgs.builder()        
            .description("primary replication group")
            .engine("redis")
            .engineVersion("5.0.6")
            .nodeType("cache.m5.large")
            .numCacheClusters(1)
            .build(), CustomResourceOptions.builder()
                .provider(aws.other_region())
                .build());

        var example = new GlobalReplicationGroup("example", GlobalReplicationGroupArgs.builder()        
            .globalReplicationGroupIdSuffix("example")
            .primaryReplicationGroupId(primary.id())
            .build(), CustomResourceOptions.builder()
                .provider(aws.other_region())
                .build());

        var secondary = new ReplicationGroup("secondary", ReplicationGroupArgs.builder()        
            .description("secondary replication group")
            .globalReplicationGroupId(example.globalReplicationGroupId())
            .numCacheClusters(1)
            .build());

    }
}
```
```yaml
resources:
  secondary:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: secondary replication group
      globalReplicationGroupId: ${example.globalReplicationGroupId}
      numCacheClusters: 1
  example:
    type: aws:elasticache:GlobalReplicationGroup
    properties:
      globalReplicationGroupIdSuffix: example
      primaryReplicationGroupId: ${primary.id}
    options:
      provider: ${aws.other_region}
  primary:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: primary replication group
      engine: redis
      engineVersion: 5.0.6
      nodeType: cache.m5.large
      numCacheClusters: 1
    options:
      provider: ${aws.other_region}
```
{{% /example %}}
{{% example %}}
### Redis AUTH and In-Transit Encryption Enabled

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.ReplicationGroup("example", {
    description: "example with authentication",
    nodeType: "cache.t2.micro",
    numCacheClusters: 1,
    port: 6379,
    subnetGroupName: aws_elasticache_subnet_group.example.name,
    securityGroupIds: [aws_security_group.example.id],
    parameterGroupName: "default.redis5.0",
    engineVersion: "5.0.6",
    transitEncryptionEnabled: true,
    authToken: "abcdefgh1234567890",
    authTokenUpdateStrategy: "ROTATE",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticache.ReplicationGroup("example",
    description="example with authentication",
    node_type="cache.t2.micro",
    num_cache_clusters=1,
    port=6379,
    subnet_group_name=aws_elasticache_subnet_group["example"]["name"],
    security_group_ids=[aws_security_group["example"]["id"]],
    parameter_group_name="default.redis5.0",
    engine_version="5.0.6",
    transit_encryption_enabled=True,
    auth_token="abcdefgh1234567890",
    auth_token_update_strategy="ROTATE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElastiCache.ReplicationGroup("example", new()
    {
        Description = "example with authentication",
        NodeType = "cache.t2.micro",
        NumCacheClusters = 1,
        Port = 6379,
        SubnetGroupName = aws_elasticache_subnet_group.Example.Name,
        SecurityGroupIds = new[]
        {
            aws_security_group.Example.Id,
        },
        ParameterGroupName = "default.redis5.0",
        EngineVersion = "5.0.6",
        TransitEncryptionEnabled = true,
        AuthToken = "abcdefgh1234567890",
        AuthTokenUpdateStrategy = "ROTATE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "example", &elasticache.ReplicationGroupArgs{
			Description:      pulumi.String("example with authentication"),
			NodeType:         pulumi.String("cache.t2.micro"),
			NumCacheClusters: pulumi.Int(1),
			Port:             pulumi.Int(6379),
			SubnetGroupName:  pulumi.Any(aws_elasticache_subnet_group.Example.Name),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Example.Id,
			},
			ParameterGroupName:       pulumi.String("default.redis5.0"),
			EngineVersion:            pulumi.String("5.0.6"),
			TransitEncryptionEnabled: pulumi.Bool(true),
			AuthToken:                pulumi.String("abcdefgh1234567890"),
			AuthTokenUpdateStrategy:  pulumi.String("ROTATE"),
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
import com.pulumi.aws.elasticache.ReplicationGroup;
import com.pulumi.aws.elasticache.ReplicationGroupArgs;
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
        var example = new ReplicationGroup("example", ReplicationGroupArgs.builder()        
            .description("example with authentication")
            .nodeType("cache.t2.micro")
            .numCacheClusters(1)
            .port(6379)
            .subnetGroupName(aws_elasticache_subnet_group.example().name())
            .securityGroupIds(aws_security_group.example().id())
            .parameterGroupName("default.redis5.0")
            .engineVersion("5.0.6")
            .transitEncryptionEnabled(true)
            .authToken("abcdefgh1234567890")
            .authTokenUpdateStrategy("ROTATE")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: example with authentication
      nodeType: cache.t2.micro
      numCacheClusters: 1
      port: 6379
      subnetGroupName: ${aws_elasticache_subnet_group.example.name}
      securityGroupIds:
        - ${aws_security_group.example.id}
      parameterGroupName: default.redis5.0
      engineVersion: 5.0.6
      transitEncryptionEnabled: true
      authToken: abcdefgh1234567890
      authTokenUpdateStrategy: ROTATE
```

> When adding a new `auth_token` to a previously passwordless replication group, using the `ROTATE` update strategy will result in support for **both** the new token and passwordless authentication. To immediately require authorization when adding the initial token, use the `SET` strategy instead. See the [Authenticating with the Redis AUTH command](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/auth.html) guide for additional details.
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache Replication Groups using the `replication_group_id`. For example:

```sh
 $ pulumi import aws:elasticache/replicationGroup:ReplicationGroup my_replication_group replication-group-1
```
 