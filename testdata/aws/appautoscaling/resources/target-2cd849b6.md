Provides an Application AutoScaling ScalableTarget resource. To manage policies which get attached to the target, see the `aws.appautoscaling.Policy` resource.

> **NOTE:** Scalable targets created before 2023-03-20 may not have an assigned `arn`. These resource cannot use `tags` or participate in `default_tags`. To prevent `pulumi preview` showing differences that can never be reconciled, use the `lifecycle.ignore_changes` meta-argument. See the example below.

> **NOTE:** The [Application Auto Scaling service automatically attempts to manage IAM Service-Linked Roles](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) when registering certain service namespaces for the first time. To manually manage this role, see the `aws.iam.ServiceLinkedRole` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### DynamoDB Table Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodbTableReadTarget", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: `table/${aws_dynamodb_table.example.name}`,
    scalableDimension: "dynamodb:table:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});
```
```python
import pulumi
import pulumi_aws as aws

dynamodb_table_read_target = aws.appautoscaling.Target("dynamodbTableReadTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id=f"table/{aws_dynamodb_table['example']['name']}",
    scalable_dimension="dynamodb:table:ReadCapacityUnits",
    service_namespace="dynamodb")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dynamodbTableReadTarget = new Aws.AppAutoScaling.Target("dynamodbTableReadTarget", new()
    {
        MaxCapacity = 100,
        MinCapacity = 5,
        ResourceId = $"table/{aws_dynamodb_table.Example.Name}",
        ScalableDimension = "dynamodb:table:ReadCapacityUnits",
        ServiceNamespace = "dynamodb",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "dynamodbTableReadTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String(fmt.Sprintf("table/%v", aws_dynamodb_table.Example.Name)),
			ScalableDimension: pulumi.String("dynamodb:table:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
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
        var dynamodbTableReadTarget = new Target("dynamodbTableReadTarget", TargetArgs.builder()        
            .maxCapacity(100)
            .minCapacity(5)
            .resourceId(String.format("table/%s", aws_dynamodb_table.example().name()))
            .scalableDimension("dynamodb:table:ReadCapacityUnits")
            .serviceNamespace("dynamodb")
            .build());

    }
}
```
```yaml
resources:
  dynamodbTableReadTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 100
      minCapacity: 5
      resourceId: table/${aws_dynamodb_table.example.name}
      scalableDimension: dynamodb:table:ReadCapacityUnits
      serviceNamespace: dynamodb
```
{{% /example %}}
{{% example %}}
### DynamoDB Index Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dynamodbIndexReadTarget = new aws.appautoscaling.Target("dynamodbIndexReadTarget", {
    maxCapacity: 100,
    minCapacity: 5,
    resourceId: `table/${aws_dynamodb_table.example.name}/index/${_var.index_name}`,
    scalableDimension: "dynamodb:index:ReadCapacityUnits",
    serviceNamespace: "dynamodb",
});
```
```python
import pulumi
import pulumi_aws as aws

dynamodb_index_read_target = aws.appautoscaling.Target("dynamodbIndexReadTarget",
    max_capacity=100,
    min_capacity=5,
    resource_id=f"table/{aws_dynamodb_table['example']['name']}/index/{var['index_name']}",
    scalable_dimension="dynamodb:index:ReadCapacityUnits",
    service_namespace="dynamodb")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dynamodbIndexReadTarget = new Aws.AppAutoScaling.Target("dynamodbIndexReadTarget", new()
    {
        MaxCapacity = 100,
        MinCapacity = 5,
        ResourceId = $"table/{aws_dynamodb_table.Example.Name}/index/{@var.Index_name}",
        ScalableDimension = "dynamodb:index:ReadCapacityUnits",
        ServiceNamespace = "dynamodb",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "dynamodbIndexReadTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(100),
			MinCapacity:       pulumi.Int(5),
			ResourceId:        pulumi.String(fmt.Sprintf("table/%v/index/%v", aws_dynamodb_table.Example.Name, _var.Index_name)),
			ScalableDimension: pulumi.String("dynamodb:index:ReadCapacityUnits"),
			ServiceNamespace:  pulumi.String("dynamodb"),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
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
        var dynamodbIndexReadTarget = new Target("dynamodbIndexReadTarget", TargetArgs.builder()        
            .maxCapacity(100)
            .minCapacity(5)
            .resourceId(String.format("table/%s/index/%s", aws_dynamodb_table.example().name(),var_.index_name()))
            .scalableDimension("dynamodb:index:ReadCapacityUnits")
            .serviceNamespace("dynamodb")
            .build());

    }
}
```
```yaml
resources:
  dynamodbIndexReadTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 100
      minCapacity: 5
      resourceId: table/${aws_dynamodb_table.example.name}/index/${var.index_name}
      scalableDimension: dynamodb:index:ReadCapacityUnits
      serviceNamespace: dynamodb
```
{{% /example %}}
{{% example %}}
### ECS Service Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsTarget = new aws.appautoscaling.Target("ecsTarget", {
    maxCapacity: 4,
    minCapacity: 1,
    resourceId: `service/${aws_ecs_cluster.example.name}/${aws_ecs_service.example.name}`,
    scalableDimension: "ecs:service:DesiredCount",
    serviceNamespace: "ecs",
});
```
```python
import pulumi
import pulumi_aws as aws

ecs_target = aws.appautoscaling.Target("ecsTarget",
    max_capacity=4,
    min_capacity=1,
    resource_id=f"service/{aws_ecs_cluster['example']['name']}/{aws_ecs_service['example']['name']}",
    scalable_dimension="ecs:service:DesiredCount",
    service_namespace="ecs")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ecsTarget = new Aws.AppAutoScaling.Target("ecsTarget", new()
    {
        MaxCapacity = 4,
        MinCapacity = 1,
        ResourceId = $"service/{aws_ecs_cluster.Example.Name}/{aws_ecs_service.Example.Name}",
        ScalableDimension = "ecs:service:DesiredCount",
        ServiceNamespace = "ecs",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "ecsTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(4),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String(fmt.Sprintf("service/%v/%v", aws_ecs_cluster.Example.Name, aws_ecs_service.Example.Name)),
			ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
			ServiceNamespace:  pulumi.String("ecs"),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
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
        var ecsTarget = new Target("ecsTarget", TargetArgs.builder()        
            .maxCapacity(4)
            .minCapacity(1)
            .resourceId(String.format("service/%s/%s", aws_ecs_cluster.example().name(),aws_ecs_service.example().name()))
            .scalableDimension("ecs:service:DesiredCount")
            .serviceNamespace("ecs")
            .build());

    }
}
```
```yaml
resources:
  ecsTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 4
      minCapacity: 1
      resourceId: service/${aws_ecs_cluster.example.name}/${aws_ecs_service.example.name}
      scalableDimension: ecs:service:DesiredCount
      serviceNamespace: ecs
```
{{% /example %}}
{{% example %}}
### Aurora Read Replica Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicas = new aws.appautoscaling.Target("replicas", {
    maxCapacity: 15,
    minCapacity: 1,
    resourceId: `cluster:${aws_rds_cluster.example.id}`,
    scalableDimension: "rds:cluster:ReadReplicaCount",
    serviceNamespace: "rds",
});
```
```python
import pulumi
import pulumi_aws as aws

replicas = aws.appautoscaling.Target("replicas",
    max_capacity=15,
    min_capacity=1,
    resource_id=f"cluster:{aws_rds_cluster['example']['id']}",
    scalable_dimension="rds:cluster:ReadReplicaCount",
    service_namespace="rds")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replicas = new Aws.AppAutoScaling.Target("replicas", new()
    {
        MaxCapacity = 15,
        MinCapacity = 1,
        ResourceId = $"cluster:{aws_rds_cluster.Example.Id}",
        ScalableDimension = "rds:cluster:ReadReplicaCount",
        ServiceNamespace = "rds",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "replicas", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(15),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String(fmt.Sprintf("cluster:%v", aws_rds_cluster.Example.Id)),
			ScalableDimension: pulumi.String("rds:cluster:ReadReplicaCount"),
			ServiceNamespace:  pulumi.String("rds"),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
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
        var replicas = new Target("replicas", TargetArgs.builder()        
            .maxCapacity(15)
            .minCapacity(1)
            .resourceId(String.format("cluster:%s", aws_rds_cluster.example().id()))
            .scalableDimension("rds:cluster:ReadReplicaCount")
            .serviceNamespace("rds")
            .build());

    }
}
```
```yaml
resources:
  replicas:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 15
      minCapacity: 1
      resourceId: cluster:${aws_rds_cluster.example.id}
      scalableDimension: rds:cluster:ReadReplicaCount
      serviceNamespace: rds
```
{{% /example %}}
{{% example %}}
### Suppressing `tags_all` Differences For Older Resources

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ecsTarget = new aws.appautoscaling.Target("ecsTarget", {
    maxCapacity: 4,
    minCapacity: 1,
    resourceId: `service/${aws_ecs_cluster.example.name}/${aws_ecs_service.example.name}`,
    scalableDimension: "ecs:service:DesiredCount",
    serviceNamespace: "ecs",
});
```
```python
import pulumi
import pulumi_aws as aws

ecs_target = aws.appautoscaling.Target("ecsTarget",
    max_capacity=4,
    min_capacity=1,
    resource_id=f"service/{aws_ecs_cluster['example']['name']}/{aws_ecs_service['example']['name']}",
    scalable_dimension="ecs:service:DesiredCount",
    service_namespace="ecs")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ecsTarget = new Aws.AppAutoScaling.Target("ecsTarget", new()
    {
        MaxCapacity = 4,
        MinCapacity = 1,
        ResourceId = $"service/{aws_ecs_cluster.Example.Name}/{aws_ecs_service.Example.Name}",
        ScalableDimension = "ecs:service:DesiredCount",
        ServiceNamespace = "ecs",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "ecsTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(4),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.String(fmt.Sprintf("service/%v/%v", aws_ecs_cluster.Example.Name, aws_ecs_service.Example.Name)),
			ScalableDimension: pulumi.String("ecs:service:DesiredCount"),
			ServiceNamespace:  pulumi.String("ecs"),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
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
        var ecsTarget = new Target("ecsTarget", TargetArgs.builder()        
            .maxCapacity(4)
            .minCapacity(1)
            .resourceId(String.format("service/%s/%s", aws_ecs_cluster.example().name(),aws_ecs_service.example().name()))
            .scalableDimension("ecs:service:DesiredCount")
            .serviceNamespace("ecs")
            .build());

    }
}
```
```yaml
resources:
  ecsTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 4
      minCapacity: 1
      resourceId: service/${aws_ecs_cluster.example.name}/${aws_ecs_service.example.name}
      scalableDimension: ecs:service:DesiredCount
      serviceNamespace: ecs
```
{{% /example %}}
{{% example %}}
### MSK / Kafka Autoscaling

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mskTarget = new aws.appautoscaling.Target("mskTarget", {
    maxCapacity: 8,
    minCapacity: 1,
    resourceId: aws_msk_cluster.example.arn,
    scalableDimension: "kafka:broker-storage:VolumeSize",
    serviceNamespace: "kafka",
});
```
```python
import pulumi
import pulumi_aws as aws

msk_target = aws.appautoscaling.Target("mskTarget",
    max_capacity=8,
    min_capacity=1,
    resource_id=aws_msk_cluster["example"]["arn"],
    scalable_dimension="kafka:broker-storage:VolumeSize",
    service_namespace="kafka")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mskTarget = new Aws.AppAutoScaling.Target("mskTarget", new()
    {
        MaxCapacity = 8,
        MinCapacity = 1,
        ResourceId = aws_msk_cluster.Example.Arn,
        ScalableDimension = "kafka:broker-storage:VolumeSize",
        ServiceNamespace = "kafka",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appautoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appautoscaling.NewTarget(ctx, "mskTarget", &appautoscaling.TargetArgs{
			MaxCapacity:       pulumi.Int(8),
			MinCapacity:       pulumi.Int(1),
			ResourceId:        pulumi.Any(aws_msk_cluster.Example.Arn),
			ScalableDimension: pulumi.String("kafka:broker-storage:VolumeSize"),
			ServiceNamespace:  pulumi.String("kafka"),
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
import com.pulumi.aws.appautoscaling.Target;
import com.pulumi.aws.appautoscaling.TargetArgs;
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
        var mskTarget = new Target("mskTarget", TargetArgs.builder()        
            .maxCapacity(8)
            .minCapacity(1)
            .resourceId(aws_msk_cluster.example().arn())
            .scalableDimension("kafka:broker-storage:VolumeSize")
            .serviceNamespace("kafka")
            .build());

    }
}
```
```yaml
resources:
  mskTarget:
    type: aws:appautoscaling:Target
    properties:
      maxCapacity: 8
      minCapacity: 1
      resourceId: ${aws_msk_cluster.example.arn}
      scalableDimension: kafka:broker-storage:VolumeSize
      serviceNamespace: kafka
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Application AutoScaling Target using the `service-namespace` , `resource-id` and `scalable-dimension` separated by `/`. For example:

```sh
 $ pulumi import aws:appautoscaling/target:Target test-target service-namespace/resource-id/scalable-dimension
```
 