Creates a replica of an existing EFS file system in the same or another region. Creating this resource causes the source EFS file system to be replicated to a new read-only destination EFS file system. Deleting this resource will cause the replication from source to destination to stop and the destination file system will no longer be read only.

> **NOTE:** Deleting this resource does **not** delete the destination file system that was created.

{{% examples %}}
## Example Usage
{{% example %}}

Will create a replica using regional storage in us-west-2 that will be encrypted by the default EFS KMS key `/aws/elasticfilesystem`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleFileSystem = new aws.efs.FileSystem("exampleFileSystem", {});
const exampleReplicationConfiguration = new aws.efs.ReplicationConfiguration("exampleReplicationConfiguration", {
    sourceFileSystemId: exampleFileSystem.id,
    destination: {
        region: "us-west-2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_file_system = aws.efs.FileSystem("exampleFileSystem")
example_replication_configuration = aws.efs.ReplicationConfiguration("exampleReplicationConfiguration",
    source_file_system_id=example_file_system.id,
    destination=aws.efs.ReplicationConfigurationDestinationArgs(
        region="us-west-2",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleFileSystem = new Aws.Efs.FileSystem("exampleFileSystem");

    var exampleReplicationConfiguration = new Aws.Efs.ReplicationConfiguration("exampleReplicationConfiguration", new()
    {
        SourceFileSystemId = exampleFileSystem.Id,
        Destination = new Aws.Efs.Inputs.ReplicationConfigurationDestinationArgs
        {
            Region = "us-west-2",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleFileSystem, err := efs.NewFileSystem(ctx, "exampleFileSystem", nil)
		if err != nil {
			return err
		}
		_, err = efs.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &efs.ReplicationConfigurationArgs{
			SourceFileSystemId: exampleFileSystem.ID(),
			Destination: &efs.ReplicationConfigurationDestinationArgs{
				Region: pulumi.String("us-west-2"),
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
import com.pulumi.aws.efs.FileSystem;
import com.pulumi.aws.efs.ReplicationConfiguration;
import com.pulumi.aws.efs.ReplicationConfigurationArgs;
import com.pulumi.aws.efs.inputs.ReplicationConfigurationDestinationArgs;
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
        var exampleFileSystem = new FileSystem("exampleFileSystem");

        var exampleReplicationConfiguration = new ReplicationConfiguration("exampleReplicationConfiguration", ReplicationConfigurationArgs.builder()        
            .sourceFileSystemId(exampleFileSystem.id())
            .destination(ReplicationConfigurationDestinationArgs.builder()
                .region("us-west-2")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleFileSystem:
    type: aws:efs:FileSystem
  exampleReplicationConfiguration:
    type: aws:efs:ReplicationConfiguration
    properties:
      sourceFileSystemId: ${exampleFileSystem.id}
      destination:
        region: us-west-2
```

Replica will be created as One Zone storage in the us-west-2b Availability Zone and encrypted with the specified KMS key.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleFileSystem = new aws.efs.FileSystem("exampleFileSystem", {});
const exampleReplicationConfiguration = new aws.efs.ReplicationConfiguration("exampleReplicationConfiguration", {
    sourceFileSystemId: exampleFileSystem.id,
    destination: {
        availabilityZoneName: "us-west-2b",
        kmsKeyId: "1234abcd-12ab-34cd-56ef-1234567890ab",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_file_system = aws.efs.FileSystem("exampleFileSystem")
example_replication_configuration = aws.efs.ReplicationConfiguration("exampleReplicationConfiguration",
    source_file_system_id=example_file_system.id,
    destination=aws.efs.ReplicationConfigurationDestinationArgs(
        availability_zone_name="us-west-2b",
        kms_key_id="1234abcd-12ab-34cd-56ef-1234567890ab",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleFileSystem = new Aws.Efs.FileSystem("exampleFileSystem");

    var exampleReplicationConfiguration = new Aws.Efs.ReplicationConfiguration("exampleReplicationConfiguration", new()
    {
        SourceFileSystemId = exampleFileSystem.Id,
        Destination = new Aws.Efs.Inputs.ReplicationConfigurationDestinationArgs
        {
            AvailabilityZoneName = "us-west-2b",
            KmsKeyId = "1234abcd-12ab-34cd-56ef-1234567890ab",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/efs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleFileSystem, err := efs.NewFileSystem(ctx, "exampleFileSystem", nil)
		if err != nil {
			return err
		}
		_, err = efs.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &efs.ReplicationConfigurationArgs{
			SourceFileSystemId: exampleFileSystem.ID(),
			Destination: &efs.ReplicationConfigurationDestinationArgs{
				AvailabilityZoneName: pulumi.String("us-west-2b"),
				KmsKeyId:             pulumi.String("1234abcd-12ab-34cd-56ef-1234567890ab"),
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
import com.pulumi.aws.efs.FileSystem;
import com.pulumi.aws.efs.ReplicationConfiguration;
import com.pulumi.aws.efs.ReplicationConfigurationArgs;
import com.pulumi.aws.efs.inputs.ReplicationConfigurationDestinationArgs;
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
        var exampleFileSystem = new FileSystem("exampleFileSystem");

        var exampleReplicationConfiguration = new ReplicationConfiguration("exampleReplicationConfiguration", ReplicationConfigurationArgs.builder()        
            .sourceFileSystemId(exampleFileSystem.id())
            .destination(ReplicationConfigurationDestinationArgs.builder()
                .availabilityZoneName("us-west-2b")
                .kmsKeyId("1234abcd-12ab-34cd-56ef-1234567890ab")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleFileSystem:
    type: aws:efs:FileSystem
  exampleReplicationConfiguration:
    type: aws:efs:ReplicationConfiguration
    properties:
      sourceFileSystemId: ${exampleFileSystem.id}
      destination:
        availabilityZoneName: us-west-2b
        kmsKeyId: 1234abcd-12ab-34cd-56ef-1234567890ab
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EFS Replication Configurations using the file system ID of either the source or destination file system. When importing, the `availability_zone_name` and `kms_key_id` attributes must __not__ be set in the configuration. The AWS API does not return these values when querying the replication configuration and their presence will therefore show as a diff in a subsequent plan. For example:

```sh
 $ pulumi import aws:efs/replicationConfiguration:ReplicationConfiguration example fs-id
```
 