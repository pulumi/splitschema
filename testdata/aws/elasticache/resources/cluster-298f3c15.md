Provides an ElastiCache Cluster resource, which manages either a
[Memcached cluster](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/WhatIs.html), a
[single-node Redis instance](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/WhatIs.html), or a
[read replica in a Redis (Cluster Mode Enabled) replication group].

For working with Redis (Cluster Mode Enabled) replication groups, see the
`aws.elasticache.ReplicationGroup` resource.

> **Note:** When you change an attribute, such as `num_cache_nodes`, by default
it is applied in the next maintenance window. Because of this, this provider may report
a difference in its planning phase because the actual modification has not yet taken
place. You can use the `apply_immediately` flag to instruct the service to apply the
change immediately. Using `apply_immediately` can result in a brief downtime as the server reboots.
See the AWS Documentation on Modifying an ElastiCache Cache Cluster for
[ElastiCache for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/Clusters.Modify.html) or
[ElastiCache for Redis](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.Modify.html)
for more information.

> **Note:** Any attribute changes that re-create the resource will be applied immediately, regardless of the value of `apply_immediately`.

{{% examples %}}
## Example Usage
{{% example %}}
### Memcached Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.Cluster("example", {
    engine: "memcached",
    nodeType: "cache.m4.large",
    numCacheNodes: 2,
    parameterGroupName: "default.memcached1.4",
    port: 11211,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticache.Cluster("example",
    engine="memcached",
    node_type="cache.m4.large",
    num_cache_nodes=2,
    parameter_group_name="default.memcached1.4",
    port=11211)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElastiCache.Cluster("example", new()
    {
        Engine = "memcached",
        NodeType = "cache.m4.large",
        NumCacheNodes = 2,
        ParameterGroupName = "default.memcached1.4",
        Port = 11211,
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
		_, err := elasticache.NewCluster(ctx, "example", &elasticache.ClusterArgs{
			Engine:             pulumi.String("memcached"),
			NodeType:           pulumi.String("cache.m4.large"),
			NumCacheNodes:      pulumi.Int(2),
			ParameterGroupName: pulumi.String("default.memcached1.4"),
			Port:               pulumi.Int(11211),
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
import com.pulumi.aws.elasticache.Cluster;
import com.pulumi.aws.elasticache.ClusterArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .engine("memcached")
            .nodeType("cache.m4.large")
            .numCacheNodes(2)
            .parameterGroupName("default.memcached1.4")
            .port(11211)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:Cluster
    properties:
      engine: memcached
      nodeType: cache.m4.large
      numCacheNodes: 2
      parameterGroupName: default.memcached1.4
      port: 11211
```
{{% /example %}}
{{% example %}}
### Redis Instance

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.elasticache.Cluster("example", {
    engine: "redis",
    engineVersion: "3.2.10",
    nodeType: "cache.m4.large",
    numCacheNodes: 1,
    parameterGroupName: "default.redis3.2",
    port: 6379,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.elasticache.Cluster("example",
    engine="redis",
    engine_version="3.2.10",
    node_type="cache.m4.large",
    num_cache_nodes=1,
    parameter_group_name="default.redis3.2",
    port=6379)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ElastiCache.Cluster("example", new()
    {
        Engine = "redis",
        EngineVersion = "3.2.10",
        NodeType = "cache.m4.large",
        NumCacheNodes = 1,
        ParameterGroupName = "default.redis3.2",
        Port = 6379,
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
		_, err := elasticache.NewCluster(ctx, "example", &elasticache.ClusterArgs{
			Engine:             pulumi.String("redis"),
			EngineVersion:      pulumi.String("3.2.10"),
			NodeType:           pulumi.String("cache.m4.large"),
			NumCacheNodes:      pulumi.Int(1),
			ParameterGroupName: pulumi.String("default.redis3.2"),
			Port:               pulumi.Int(6379),
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
import com.pulumi.aws.elasticache.Cluster;
import com.pulumi.aws.elasticache.ClusterArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .engine("redis")
            .engineVersion("3.2.10")
            .nodeType("cache.m4.large")
            .numCacheNodes(1)
            .parameterGroupName("default.redis3.2")
            .port(6379)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:Cluster
    properties:
      engine: redis
      engineVersion: 3.2.10
      nodeType: cache.m4.large
      numCacheNodes: 1
      parameterGroupName: default.redis3.2
      port: 6379
```
{{% /example %}}
{{% example %}}
### Redis Cluster Mode Disabled Read Replica Instance

These inherit their settings from the replication group.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replica = new aws.elasticache.Cluster("replica", {replicationGroupId: aws_elasticache_replication_group.example.id});
```
```python
import pulumi
import pulumi_aws as aws

replica = aws.elasticache.Cluster("replica", replication_group_id=aws_elasticache_replication_group["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replica = new Aws.ElastiCache.Cluster("replica", new()
    {
        ReplicationGroupId = aws_elasticache_replication_group.Example.Id,
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
		_, err := elasticache.NewCluster(ctx, "replica", &elasticache.ClusterArgs{
			ReplicationGroupId: pulumi.Any(aws_elasticache_replication_group.Example.Id),
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
import com.pulumi.aws.elasticache.Cluster;
import com.pulumi.aws.elasticache.ClusterArgs;
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
        var replica = new Cluster("replica", ClusterArgs.builder()        
            .replicationGroupId(aws_elasticache_replication_group.example().id())
            .build());

    }
}
```
```yaml
resources:
  replica:
    type: aws:elasticache:Cluster
    properties:
      replicationGroupId: ${aws_elasticache_replication_group.example.id}
```
{{% /example %}}
{{% example %}}
### Redis Log Delivery configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.elasticache.Cluster("test", {
    engine: "redis",
    nodeType: "cache.t3.micro",
    numCacheNodes: 1,
    port: 6379,
    applyImmediately: true,
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

test = aws.elasticache.Cluster("test",
    engine="redis",
    node_type="cache.t3.micro",
    num_cache_nodes=1,
    port=6379,
    apply_immediately=True,
    log_delivery_configurations=[
        aws.elasticache.ClusterLogDeliveryConfigurationArgs(
            destination=aws_cloudwatch_log_group["example"]["name"],
            destination_type="cloudwatch-logs",
            log_format="text",
            log_type="slow-log",
        ),
        aws.elasticache.ClusterLogDeliveryConfigurationArgs(
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
    var test = new Aws.ElastiCache.Cluster("test", new()
    {
        Engine = "redis",
        NodeType = "cache.t3.micro",
        NumCacheNodes = 1,
        Port = 6379,
        ApplyImmediately = true,
        LogDeliveryConfigurations = new[]
        {
            new Aws.ElastiCache.Inputs.ClusterLogDeliveryConfigurationArgs
            {
                Destination = aws_cloudwatch_log_group.Example.Name,
                DestinationType = "cloudwatch-logs",
                LogFormat = "text",
                LogType = "slow-log",
            },
            new Aws.ElastiCache.Inputs.ClusterLogDeliveryConfigurationArgs
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
		_, err := elasticache.NewCluster(ctx, "test", &elasticache.ClusterArgs{
			Engine:           pulumi.String("redis"),
			NodeType:         pulumi.String("cache.t3.micro"),
			NumCacheNodes:    pulumi.Int(1),
			Port:             pulumi.Int(6379),
			ApplyImmediately: pulumi.Bool(true),
			LogDeliveryConfigurations: elasticache.ClusterLogDeliveryConfigurationArray{
				&elasticache.ClusterLogDeliveryConfigurationArgs{
					Destination:     pulumi.Any(aws_cloudwatch_log_group.Example.Name),
					DestinationType: pulumi.String("cloudwatch-logs"),
					LogFormat:       pulumi.String("text"),
					LogType:         pulumi.String("slow-log"),
				},
				&elasticache.ClusterLogDeliveryConfigurationArgs{
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
import com.pulumi.aws.elasticache.Cluster;
import com.pulumi.aws.elasticache.ClusterArgs;
import com.pulumi.aws.elasticache.inputs.ClusterLogDeliveryConfigurationArgs;
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
        var test = new Cluster("test", ClusterArgs.builder()        
            .engine("redis")
            .nodeType("cache.t3.micro")
            .numCacheNodes(1)
            .port(6379)
            .applyImmediately(true)
            .logDeliveryConfigurations(            
                ClusterLogDeliveryConfigurationArgs.builder()
                    .destination(aws_cloudwatch_log_group.example().name())
                    .destinationType("cloudwatch-logs")
                    .logFormat("text")
                    .logType("slow-log")
                    .build(),
                ClusterLogDeliveryConfigurationArgs.builder()
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
    type: aws:elasticache:Cluster
    properties:
      engine: redis
      nodeType: cache.t3.micro
      numCacheNodes: 1
      port: 6379
      applyImmediately: true
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
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache Clusters using the `cluster_id`. For example:

```sh
 $ pulumi import aws:elasticache/cluster:Cluster my_cluster my_cluster
```
 