Manages RDS Aurora Cluster Database Activity Streams.

Database Activity Streams have some limits and requirements, refer to the [Monitoring Amazon Aurora using Database Activity Streams](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html) documentation for detailed limitations and requirements.

> **Note:** This resource always calls the RDS [`StartActivityStream`][2] API with the `ApplyImmediately` parameter set to `true`. This is because the provider needs the activity stream to be started in order for it to get the associated attributes.

> **Note:** This resource depends on having at least one `aws.rds.ClusterInstance` created. To avoid race conditions when all resources are being created together, add an explicit resource reference using the resource `depends_on` meta-argument.

> **Note:** This resource is available in all regions except the following: `cn-north-1`, `cn-northwest-1`, `us-gov-east-1`, `us-gov-west-1`

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultCluster = new aws.rds.Cluster("defaultCluster", {
    clusterIdentifier: "aurora-cluster-demo",
    availabilityZones: [
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    databaseName: "mydb",
    masterUsername: "foo",
    masterPassword: "mustbeeightcharaters",
    engine: "aurora-postgresql",
    engineVersion: "13.4",
});
const defaultClusterInstance = new aws.rds.ClusterInstance("defaultClusterInstance", {
    identifier: "aurora-instance-demo",
    clusterIdentifier: defaultCluster.clusterIdentifier,
    engine: defaultCluster.engine,
    instanceClass: "db.r6g.large",
});
const defaultKey = new aws.kms.Key("defaultKey", {description: "AWS KMS Key to encrypt Database Activity Stream"});
const defaultClusterActivityStream = new aws.rds.ClusterActivityStream("defaultClusterActivityStream", {
    resourceArn: defaultCluster.arn,
    mode: "async",
    kmsKeyId: defaultKey.keyId,
}, {
    dependsOn: [defaultClusterInstance],
});
```
```python
import pulumi
import pulumi_aws as aws

default_cluster = aws.rds.Cluster("defaultCluster",
    cluster_identifier="aurora-cluster-demo",
    availability_zones=[
        "us-west-2a",
        "us-west-2b",
        "us-west-2c",
    ],
    database_name="mydb",
    master_username="foo",
    master_password="mustbeeightcharaters",
    engine="aurora-postgresql",
    engine_version="13.4")
default_cluster_instance = aws.rds.ClusterInstance("defaultClusterInstance",
    identifier="aurora-instance-demo",
    cluster_identifier=default_cluster.cluster_identifier,
    engine=default_cluster.engine,
    instance_class="db.r6g.large")
default_key = aws.kms.Key("defaultKey", description="AWS KMS Key to encrypt Database Activity Stream")
default_cluster_activity_stream = aws.rds.ClusterActivityStream("defaultClusterActivityStream",
    resource_arn=default_cluster.arn,
    mode="async",
    kms_key_id=default_key.key_id,
    opts=pulumi.ResourceOptions(depends_on=[default_cluster_instance]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultCluster = new Aws.Rds.Cluster("defaultCluster", new()
    {
        ClusterIdentifier = "aurora-cluster-demo",
        AvailabilityZones = new[]
        {
            "us-west-2a",
            "us-west-2b",
            "us-west-2c",
        },
        DatabaseName = "mydb",
        MasterUsername = "foo",
        MasterPassword = "mustbeeightcharaters",
        Engine = "aurora-postgresql",
        EngineVersion = "13.4",
    });

    var defaultClusterInstance = new Aws.Rds.ClusterInstance("defaultClusterInstance", new()
    {
        Identifier = "aurora-instance-demo",
        ClusterIdentifier = defaultCluster.ClusterIdentifier,
        Engine = defaultCluster.Engine,
        InstanceClass = "db.r6g.large",
    });

    var defaultKey = new Aws.Kms.Key("defaultKey", new()
    {
        Description = "AWS KMS Key to encrypt Database Activity Stream",
    });

    var defaultClusterActivityStream = new Aws.Rds.ClusterActivityStream("defaultClusterActivityStream", new()
    {
        ResourceArn = defaultCluster.Arn,
        Mode = "async",
        KmsKeyId = defaultKey.KeyId,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            defaultClusterInstance,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultCluster, err := rds.NewCluster(ctx, "defaultCluster", &rds.ClusterArgs{
			ClusterIdentifier: pulumi.String("aurora-cluster-demo"),
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
				pulumi.String("us-west-2c"),
			},
			DatabaseName:   pulumi.String("mydb"),
			MasterUsername: pulumi.String("foo"),
			MasterPassword: pulumi.String("mustbeeightcharaters"),
			Engine:         pulumi.String("aurora-postgresql"),
			EngineVersion:  pulumi.String("13.4"),
		})
		if err != nil {
			return err
		}
		defaultClusterInstance, err := rds.NewClusterInstance(ctx, "defaultClusterInstance", &rds.ClusterInstanceArgs{
			Identifier:        pulumi.String("aurora-instance-demo"),
			ClusterIdentifier: defaultCluster.ClusterIdentifier,
			Engine:            defaultCluster.Engine,
			InstanceClass:     pulumi.String("db.r6g.large"),
		})
		if err != nil {
			return err
		}
		defaultKey, err := kms.NewKey(ctx, "defaultKey", &kms.KeyArgs{
			Description: pulumi.String("AWS KMS Key to encrypt Database Activity Stream"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewClusterActivityStream(ctx, "defaultClusterActivityStream", &rds.ClusterActivityStreamArgs{
			ResourceArn: defaultCluster.Arn,
			Mode:        pulumi.String("async"),
			KmsKeyId:    defaultKey.KeyId,
		}, pulumi.DependsOn([]pulumi.Resource{
			defaultClusterInstance,
		}))
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
import com.pulumi.aws.rds.Cluster;
import com.pulumi.aws.rds.ClusterArgs;
import com.pulumi.aws.rds.ClusterInstance;
import com.pulumi.aws.rds.ClusterInstanceArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.rds.ClusterActivityStream;
import com.pulumi.aws.rds.ClusterActivityStreamArgs;
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
        var defaultCluster = new Cluster("defaultCluster", ClusterArgs.builder()        
            .clusterIdentifier("aurora-cluster-demo")
            .availabilityZones(            
                "us-west-2a",
                "us-west-2b",
                "us-west-2c")
            .databaseName("mydb")
            .masterUsername("foo")
            .masterPassword("mustbeeightcharaters")
            .engine("aurora-postgresql")
            .engineVersion("13.4")
            .build());

        var defaultClusterInstance = new ClusterInstance("defaultClusterInstance", ClusterInstanceArgs.builder()        
            .identifier("aurora-instance-demo")
            .clusterIdentifier(defaultCluster.clusterIdentifier())
            .engine(defaultCluster.engine())
            .instanceClass("db.r6g.large")
            .build());

        var defaultKey = new Key("defaultKey", KeyArgs.builder()        
            .description("AWS KMS Key to encrypt Database Activity Stream")
            .build());

        var defaultClusterActivityStream = new ClusterActivityStream("defaultClusterActivityStream", ClusterActivityStreamArgs.builder()        
            .resourceArn(defaultCluster.arn())
            .mode("async")
            .kmsKeyId(defaultKey.keyId())
            .build(), CustomResourceOptions.builder()
                .dependsOn(defaultClusterInstance)
                .build());

    }
}
```
```yaml
resources:
  defaultCluster:
    type: aws:rds:Cluster
    properties:
      clusterIdentifier: aurora-cluster-demo
      availabilityZones:
        - us-west-2a
        - us-west-2b
        - us-west-2c
      databaseName: mydb
      masterUsername: foo
      masterPassword: mustbeeightcharaters
      engine: aurora-postgresql
      engineVersion: '13.4'
  defaultClusterInstance:
    type: aws:rds:ClusterInstance
    properties:
      identifier: aurora-instance-demo
      clusterIdentifier: ${defaultCluster.clusterIdentifier}
      engine: ${defaultCluster.engine}
      instanceClass: db.r6g.large
  defaultKey:
    type: aws:kms:Key
    properties:
      description: AWS KMS Key to encrypt Database Activity Stream
  defaultClusterActivityStream:
    type: aws:rds:ClusterActivityStream
    properties:
      resourceArn: ${defaultCluster.arn}
      mode: async
      kmsKeyId: ${defaultKey.keyId}
    options:
      dependson:
        - ${defaultClusterInstance}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RDS Aurora Cluster Database Activity Streams using the `resource_arn`. For example:

```sh
 $ pulumi import aws:rds/clusterActivityStream:ClusterActivityStream default arn:aws:rds:us-west-2:123456789012:cluster:aurora-cluster-demo
```
 