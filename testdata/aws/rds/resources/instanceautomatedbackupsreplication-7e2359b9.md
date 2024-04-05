Manage cross-region replication of automated backups to a different AWS Region. Documentation for cross-region automated backup replication can be found at:

* [Replicating automated backups to another AWS Region](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReplicateBackups.html)

> **Note:** This resource has to be created in the destination region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.InstanceAutomatedBackupsReplication("default", {
    retentionPeriod: 14,
    sourceDbInstanceArn: "arn:aws:rds:us-west-2:123456789012:db:mydatabase",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.InstanceAutomatedBackupsReplication("default",
    retention_period=14,
    source_db_instance_arn="arn:aws:rds:us-west-2:123456789012:db:mydatabase")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.InstanceAutomatedBackupsReplication("default", new()
    {
        RetentionPeriod = 14,
        SourceDbInstanceArn = "arn:aws:rds:us-west-2:123456789012:db:mydatabase",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewInstanceAutomatedBackupsReplication(ctx, "default", &rds.InstanceAutomatedBackupsReplicationArgs{
			RetentionPeriod:     pulumi.Int(14),
			SourceDbInstanceArn: pulumi.String("arn:aws:rds:us-west-2:123456789012:db:mydatabase"),
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
import com.pulumi.aws.rds.InstanceAutomatedBackupsReplication;
import com.pulumi.aws.rds.InstanceAutomatedBackupsReplicationArgs;
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
        var default_ = new InstanceAutomatedBackupsReplication("default", InstanceAutomatedBackupsReplicationArgs.builder()        
            .retentionPeriod(14)
            .sourceDbInstanceArn("arn:aws:rds:us-west-2:123456789012:db:mydatabase")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:InstanceAutomatedBackupsReplication
    properties:
      retentionPeriod: 14
      sourceDbInstanceArn: arn:aws:rds:us-west-2:123456789012:db:mydatabase
```
{{% /example %}}
{{% /examples %}}
## Encrypting the automated backup with KMS

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.rds.InstanceAutomatedBackupsReplication("default", {
    kmsKeyId: "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    sourceDbInstanceArn: "arn:aws:rds:us-west-2:123456789012:db:mydatabase",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.rds.InstanceAutomatedBackupsReplication("default",
    kms_key_id="arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    source_db_instance_arn="arn:aws:rds:us-west-2:123456789012:db:mydatabase")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Rds.InstanceAutomatedBackupsReplication("default", new()
    {
        KmsKeyId = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
        SourceDbInstanceArn = "arn:aws:rds:us-west-2:123456789012:db:mydatabase",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewInstanceAutomatedBackupsReplication(ctx, "default", &rds.InstanceAutomatedBackupsReplicationArgs{
			KmsKeyId:            pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
			SourceDbInstanceArn: pulumi.String("arn:aws:rds:us-west-2:123456789012:db:mydatabase"),
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
import com.pulumi.aws.rds.InstanceAutomatedBackupsReplication;
import com.pulumi.aws.rds.InstanceAutomatedBackupsReplicationArgs;
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
        var default_ = new InstanceAutomatedBackupsReplication("default", InstanceAutomatedBackupsReplicationArgs.builder()        
            .kmsKeyId("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")
            .sourceDbInstanceArn("arn:aws:rds:us-west-2:123456789012:db:mydatabase")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:rds:InstanceAutomatedBackupsReplication
    properties:
      kmsKeyId: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
      sourceDbInstanceArn: arn:aws:rds:us-west-2:123456789012:db:mydatabase
```

## Example including a RDS DB instance

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replica = new aws.Provider("replica", {region: "us-west-2"});
const defaultInstance = new aws.rds.Instance("defaultInstance", {
    allocatedStorage: 10,
    identifier: "mydb",
    engine: "postgres",
    engineVersion: "13.4",
    instanceClass: "db.t3.micro",
    dbName: "mydb",
    username: "masterusername",
    password: "mustbeeightcharacters",
    backupRetentionPeriod: 7,
    storageEncrypted: true,
    skipFinalSnapshot: true,
});
const defaultKey = new aws.kms.Key("defaultKey", {description: "Encryption key for automated backups"}, {
    provider: aws.replica,
});
const defaultInstanceAutomatedBackupsReplication = new aws.rds.InstanceAutomatedBackupsReplication("defaultInstanceAutomatedBackupsReplication", {
    sourceDbInstanceArn: defaultInstance.arn,
    kmsKeyId: defaultKey.arn,
}, {
    provider: aws.replica,
});
```
```python
import pulumi
import pulumi_aws as aws

replica = aws.Provider("replica", region="us-west-2")
default_instance = aws.rds.Instance("defaultInstance",
    allocated_storage=10,
    identifier="mydb",
    engine="postgres",
    engine_version="13.4",
    instance_class="db.t3.micro",
    db_name="mydb",
    username="masterusername",
    password="mustbeeightcharacters",
    backup_retention_period=7,
    storage_encrypted=True,
    skip_final_snapshot=True)
default_key = aws.kms.Key("defaultKey", description="Encryption key for automated backups",
opts=pulumi.ResourceOptions(provider=aws["replica"]))
default_instance_automated_backups_replication = aws.rds.InstanceAutomatedBackupsReplication("defaultInstanceAutomatedBackupsReplication",
    source_db_instance_arn=default_instance.arn,
    kms_key_id=default_key.arn,
    opts=pulumi.ResourceOptions(provider=aws["replica"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replica = new Aws.Provider("replica", new()
    {
        Region = "us-west-2",
    });

    var defaultInstance = new Aws.Rds.Instance("defaultInstance", new()
    {
        AllocatedStorage = 10,
        Identifier = "mydb",
        Engine = "postgres",
        EngineVersion = "13.4",
        InstanceClass = "db.t3.micro",
        DbName = "mydb",
        Username = "masterusername",
        Password = "mustbeeightcharacters",
        BackupRetentionPeriod = 7,
        StorageEncrypted = true,
        SkipFinalSnapshot = true,
    });

    var defaultKey = new Aws.Kms.Key("defaultKey", new()
    {
        Description = "Encryption key for automated backups",
    }, new CustomResourceOptions
    {
        Provider = aws.Replica,
    });

    var defaultInstanceAutomatedBackupsReplication = new Aws.Rds.InstanceAutomatedBackupsReplication("defaultInstanceAutomatedBackupsReplication", new()
    {
        SourceDbInstanceArn = defaultInstance.Arn,
        KmsKeyId = defaultKey.Arn,
    }, new CustomResourceOptions
    {
        Provider = aws.Replica,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "replica", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		defaultInstance, err := rds.NewInstance(ctx, "defaultInstance", &rds.InstanceArgs{
			AllocatedStorage:      pulumi.Int(10),
			Identifier:            pulumi.String("mydb"),
			Engine:                pulumi.String("postgres"),
			EngineVersion:         pulumi.String("13.4"),
			InstanceClass:         pulumi.String("db.t3.micro"),
			DbName:                pulumi.String("mydb"),
			Username:              pulumi.String("masterusername"),
			Password:              pulumi.String("mustbeeightcharacters"),
			BackupRetentionPeriod: pulumi.Int(7),
			StorageEncrypted:      pulumi.Bool(true),
			SkipFinalSnapshot:     pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		defaultKey, err := kms.NewKey(ctx, "defaultKey", &kms.KeyArgs{
			Description: pulumi.String("Encryption key for automated backups"),
		}, pulumi.Provider(aws.Replica))
		if err != nil {
			return err
		}
		_, err = rds.NewInstanceAutomatedBackupsReplication(ctx, "defaultInstanceAutomatedBackupsReplication", &rds.InstanceAutomatedBackupsReplicationArgs{
			SourceDbInstanceArn: defaultInstance.Arn,
			KmsKeyId:            defaultKey.Arn,
		}, pulumi.Provider(aws.Replica))
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.rds.InstanceAutomatedBackupsReplication;
import com.pulumi.aws.rds.InstanceAutomatedBackupsReplicationArgs;
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
        var replica = new Provider("replica", ProviderArgs.builder()        
            .region("us-west-2")
            .build());

        var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()        
            .allocatedStorage(10)
            .identifier("mydb")
            .engine("postgres")
            .engineVersion("13.4")
            .instanceClass("db.t3.micro")
            .dbName("mydb")
            .username("masterusername")
            .password("mustbeeightcharacters")
            .backupRetentionPeriod(7)
            .storageEncrypted(true)
            .skipFinalSnapshot(true)
            .build());

        var defaultKey = new Key("defaultKey", KeyArgs.builder()        
            .description("Encryption key for automated backups")
            .build(), CustomResourceOptions.builder()
                .provider(aws.replica())
                .build());

        var defaultInstanceAutomatedBackupsReplication = new InstanceAutomatedBackupsReplication("defaultInstanceAutomatedBackupsReplication", InstanceAutomatedBackupsReplicationArgs.builder()        
            .sourceDbInstanceArn(defaultInstance.arn())
            .kmsKeyId(defaultKey.arn())
            .build(), CustomResourceOptions.builder()
                .provider(aws.replica())
                .build());

    }
}
```
```yaml
resources:
  replica:
    type: pulumi:providers:aws
    properties:
      region: us-west-2
  defaultInstance:
    type: aws:rds:Instance
    properties:
      allocatedStorage: 10
      identifier: mydb
      engine: postgres
      engineVersion: '13.4'
      instanceClass: db.t3.micro
      dbName: mydb
      username: masterusername
      password: mustbeeightcharacters
      backupRetentionPeriod: 7
      storageEncrypted: true
      skipFinalSnapshot: true
  defaultKey:
    type: aws:kms:Key
    properties:
      description: Encryption key for automated backups
    options:
      provider: ${aws.replica}
  defaultInstanceAutomatedBackupsReplication:
    type: aws:rds:InstanceAutomatedBackupsReplication
    properties:
      sourceDbInstanceArn: ${defaultInstance.arn}
      kmsKeyId: ${defaultKey.arn}
    options:
      provider: ${aws.replica}
```


## Import

Using `pulumi import`, import RDS instance automated backups replication using the `arn`. For example:

```sh
 $ pulumi import aws:rds/instanceAutomatedBackupsReplication:InstanceAutomatedBackupsReplication default arn:aws:rds:us-east-1:123456789012:auto-backup:ab-faaa2mgdj1vmp4xflr7yhsrmtbtob7ltrzzz2my
```
 