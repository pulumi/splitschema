Provides an ElastiCache Global Replication Group resource, which manages replication between two or more Replication Groups in different regions. For more information, see the [ElastiCache User Guide](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Redis-Global-Datastore.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Global replication group with one secondary replication group

The global replication group depends on the primary group existing. Secondary replication groups depend on the global replication group. the provider dependency management will handle this transparently using resource value references.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.elasticache.ReplicationGroup("primary", {
    description: "primary replication group",
    engine: "redis",
    engineVersion: "5.0.6",
    nodeType: "cache.m5.large",
    numCacheClusters: 1,
});
const example = new aws.elasticache.GlobalReplicationGroup("example", {
    globalReplicationGroupIdSuffix: "example",
    primaryReplicationGroupId: primary.id,
});
const secondary = new aws.elasticache.ReplicationGroup("secondary", {
    description: "secondary replication group",
    globalReplicationGroupId: example.globalReplicationGroupId,
    numCacheClusters: 1,
}, {
    provider: aws.other_region,
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
    num_cache_clusters=1)
example = aws.elasticache.GlobalReplicationGroup("example",
    global_replication_group_id_suffix="example",
    primary_replication_group_id=primary.id)
secondary = aws.elasticache.ReplicationGroup("secondary",
    description="secondary replication group",
    global_replication_group_id=example.global_replication_group_id,
    num_cache_clusters=1,
    opts=pulumi.ResourceOptions(provider=aws["other_region"]))
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
    });

    var example = new Aws.ElastiCache.GlobalReplicationGroup("example", new()
    {
        GlobalReplicationGroupIdSuffix = "example",
        PrimaryReplicationGroupId = primary.Id,
    });

    var secondary = new Aws.ElastiCache.ReplicationGroup("secondary", new()
    {
        Description = "secondary replication group",
        GlobalReplicationGroupId = example.GlobalReplicationGroupId,
        NumCacheClusters = 1,
    }, new CustomResourceOptions
    {
        Provider = aws.Other_region,
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
		})
		if err != nil {
			return err
		}
		example, err := elasticache.NewGlobalReplicationGroup(ctx, "example", &elasticache.GlobalReplicationGroupArgs{
			GlobalReplicationGroupIdSuffix: pulumi.String("example"),
			PrimaryReplicationGroupId:      primary.ID(),
		})
		if err != nil {
			return err
		}
		_, err = elasticache.NewReplicationGroup(ctx, "secondary", &elasticache.ReplicationGroupArgs{
			Description:              pulumi.String("secondary replication group"),
			GlobalReplicationGroupId: example.GlobalReplicationGroupId,
			NumCacheClusters:         pulumi.Int(1),
		}, pulumi.Provider(aws.Other_region))
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
            .build());

        var example = new GlobalReplicationGroup("example", GlobalReplicationGroupArgs.builder()        
            .globalReplicationGroupIdSuffix("example")
            .primaryReplicationGroupId(primary.id())
            .build());

        var secondary = new ReplicationGroup("secondary", ReplicationGroupArgs.builder()        
            .description("secondary replication group")
            .globalReplicationGroupId(example.globalReplicationGroupId())
            .numCacheClusters(1)
            .build(), CustomResourceOptions.builder()
                .provider(aws.other_region())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:GlobalReplicationGroup
    properties:
      globalReplicationGroupIdSuffix: example
      primaryReplicationGroupId: ${primary.id}
  primary:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: primary replication group
      engine: redis
      engineVersion: 5.0.6
      nodeType: cache.m5.large
      numCacheClusters: 1
  secondary:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: secondary replication group
      globalReplicationGroupId: ${example.globalReplicationGroupId}
      numCacheClusters: 1
    options:
      provider: ${aws.other_region}
```
{{% /example %}}
{{% example %}}
### Managing Redis Engine Versions

The initial Redis version is determined by the version set on the primary replication group.
However, once it is part of a Global Replication Group,
the Global Replication Group manages the version of all member replication groups.

The member replication groups must have `lifecycle.ignore_changes[engine_version]` set,
or the provider will always return a diff.

In this example,
the primary replication group will be created with Redis 6.0,
and then upgraded to Redis 6.2 once added to the Global Replication Group.
The secondary replication group will be created with Redis 6.2.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.elasticache.ReplicationGroup("primary", {
    description: "primary replication group",
    engine: "redis",
    engineVersion: "6.0",
    nodeType: "cache.m5.large",
    numCacheClusters: 1,
});
const example = new aws.elasticache.GlobalReplicationGroup("example", {
    globalReplicationGroupIdSuffix: "example",
    primaryReplicationGroupId: primary.id,
    engineVersion: "6.2",
});
const secondary = new aws.elasticache.ReplicationGroup("secondary", {
    description: "secondary replication group",
    globalReplicationGroupId: example.globalReplicationGroupId,
    numCacheClusters: 1,
}, {
    provider: aws.other_region,
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.elasticache.ReplicationGroup("primary",
    description="primary replication group",
    engine="redis",
    engine_version="6.0",
    node_type="cache.m5.large",
    num_cache_clusters=1)
example = aws.elasticache.GlobalReplicationGroup("example",
    global_replication_group_id_suffix="example",
    primary_replication_group_id=primary.id,
    engine_version="6.2")
secondary = aws.elasticache.ReplicationGroup("secondary",
    description="secondary replication group",
    global_replication_group_id=example.global_replication_group_id,
    num_cache_clusters=1,
    opts=pulumi.ResourceOptions(provider=aws["other_region"]))
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
        EngineVersion = "6.0",
        NodeType = "cache.m5.large",
        NumCacheClusters = 1,
    });

    var example = new Aws.ElastiCache.GlobalReplicationGroup("example", new()
    {
        GlobalReplicationGroupIdSuffix = "example",
        PrimaryReplicationGroupId = primary.Id,
        EngineVersion = "6.2",
    });

    var secondary = new Aws.ElastiCache.ReplicationGroup("secondary", new()
    {
        Description = "secondary replication group",
        GlobalReplicationGroupId = example.GlobalReplicationGroupId,
        NumCacheClusters = 1,
    }, new CustomResourceOptions
    {
        Provider = aws.Other_region,
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
			EngineVersion:    pulumi.String("6.0"),
			NodeType:         pulumi.String("cache.m5.large"),
			NumCacheClusters: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		example, err := elasticache.NewGlobalReplicationGroup(ctx, "example", &elasticache.GlobalReplicationGroupArgs{
			GlobalReplicationGroupIdSuffix: pulumi.String("example"),
			PrimaryReplicationGroupId:      primary.ID(),
			EngineVersion:                  pulumi.String("6.2"),
		})
		if err != nil {
			return err
		}
		_, err = elasticache.NewReplicationGroup(ctx, "secondary", &elasticache.ReplicationGroupArgs{
			Description:              pulumi.String("secondary replication group"),
			GlobalReplicationGroupId: example.GlobalReplicationGroupId,
			NumCacheClusters:         pulumi.Int(1),
		}, pulumi.Provider(aws.Other_region))
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
            .engineVersion("6.0")
            .nodeType("cache.m5.large")
            .numCacheClusters(1)
            .build());

        var example = new GlobalReplicationGroup("example", GlobalReplicationGroupArgs.builder()        
            .globalReplicationGroupIdSuffix("example")
            .primaryReplicationGroupId(primary.id())
            .engineVersion("6.2")
            .build());

        var secondary = new ReplicationGroup("secondary", ReplicationGroupArgs.builder()        
            .description("secondary replication group")
            .globalReplicationGroupId(example.globalReplicationGroupId())
            .numCacheClusters(1)
            .build(), CustomResourceOptions.builder()
                .provider(aws.other_region())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:elasticache:GlobalReplicationGroup
    properties:
      globalReplicationGroupIdSuffix: example
      primaryReplicationGroupId: ${primary.id}
      engineVersion: '6.2'
  primary:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: primary replication group
      engine: redis
      engineVersion: '6.0'
      nodeType: cache.m5.large
      numCacheClusters: 1
  secondary:
    type: aws:elasticache:ReplicationGroup
    properties:
      description: secondary replication group
      globalReplicationGroupId: ${example.globalReplicationGroupId}
      numCacheClusters: 1
    options:
      provider: ${aws.other_region}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache Global Replication Groups using the `global_replication_group_id`. For example:

```sh
 $ pulumi import aws:elasticache/globalReplicationGroup:GlobalReplicationGroup my_global_replication_group okuqm-global-replication-group-1
```
 