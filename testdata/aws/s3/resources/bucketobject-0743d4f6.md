Provides an S3 object resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Encrypting with KMS Key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplekms = new aws.kms.Key("examplekms", {
    description: "KMS key 1",
    deletionWindowInDays: 7,
});
const examplebucket = new aws.s3.BucketV2("examplebucket", {});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: examplebucket.id,
    acl: "private",
});
const exampleBucketObject = new aws.s3.BucketObject("exampleBucketObject", {
    key: "someobject",
    bucket: examplebucket.id,
    source: new pulumi.asset.FileAsset("index.html"),
    kmsKeyId: examplekms.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

examplekms = aws.kms.Key("examplekms",
    description="KMS key 1",
    deletion_window_in_days=7)
examplebucket = aws.s3.BucketV2("examplebucket")
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=examplebucket.id,
    acl="private")
example_bucket_object = aws.s3.BucketObject("exampleBucketObject",
    key="someobject",
    bucket=examplebucket.id,
    source=pulumi.FileAsset("index.html"),
    kms_key_id=examplekms.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplekms = new Aws.Kms.Key("examplekms", new()
    {
        Description = "KMS key 1",
        DeletionWindowInDays = 7,
    });

    var examplebucket = new Aws.S3.BucketV2("examplebucket");

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = examplebucket.Id,
        Acl = "private",
    });

    var exampleBucketObject = new Aws.S3.BucketObject("exampleBucketObject", new()
    {
        Key = "someobject",
        Bucket = examplebucket.Id,
        Source = new FileAsset("index.html"),
        KmsKeyId = examplekms.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplekms, err := kms.NewKey(ctx, "examplekms", &kms.KeyArgs{
			Description:          pulumi.String("KMS key 1"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		examplebucket, err := s3.NewBucketV2(ctx, "examplebucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: examplebucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "exampleBucketObject", &s3.BucketObjectArgs{
			Key:      pulumi.String("someobject"),
			Bucket:   examplebucket.ID(),
			Source:   pulumi.NewFileAsset("index.html"),
			KmsKeyId: examplekms.Arn,
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketObject;
import com.pulumi.aws.s3.BucketObjectArgs;
import com.pulumi.asset.FileAsset;
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
        var examplekms = new Key("examplekms", KeyArgs.builder()        
            .description("KMS key 1")
            .deletionWindowInDays(7)
            .build());

        var examplebucket = new BucketV2("examplebucket");

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(examplebucket.id())
            .acl("private")
            .build());

        var exampleBucketObject = new BucketObject("exampleBucketObject", BucketObjectArgs.builder()        
            .key("someobject")
            .bucket(examplebucket.id())
            .source(new FileAsset("index.html"))
            .kmsKeyId(examplekms.arn())
            .build());

    }
}
```
```yaml
resources:
  examplekms:
    type: aws:kms:Key
    properties:
      description: KMS key 1
      deletionWindowInDays: 7
  examplebucket:
    type: aws:s3:BucketV2
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${examplebucket.id}
      acl: private
  exampleBucketObject:
    type: aws:s3:BucketObject
    properties:
      key: someobject
      bucket: ${examplebucket.id}
      source:
        fn::FileAsset: index.html
      kmsKeyId: ${examplekms.arn}
```
{{% /example %}}
{{% example %}}
### Server Side Encryption with S3 Default Master Key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplebucket = new aws.s3.BucketV2("examplebucket", {});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: examplebucket.id,
    acl: "private",
});
const exampleBucketObject = new aws.s3.BucketObject("exampleBucketObject", {
    key: "someobject",
    bucket: examplebucket.id,
    source: new pulumi.asset.FileAsset("index.html"),
    serverSideEncryption: "aws:kms",
});
```
```python
import pulumi
import pulumi_aws as aws

examplebucket = aws.s3.BucketV2("examplebucket")
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=examplebucket.id,
    acl="private")
example_bucket_object = aws.s3.BucketObject("exampleBucketObject",
    key="someobject",
    bucket=examplebucket.id,
    source=pulumi.FileAsset("index.html"),
    server_side_encryption="aws:kms")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplebucket = new Aws.S3.BucketV2("examplebucket");

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = examplebucket.Id,
        Acl = "private",
    });

    var exampleBucketObject = new Aws.S3.BucketObject("exampleBucketObject", new()
    {
        Key = "someobject",
        Bucket = examplebucket.Id,
        Source = new FileAsset("index.html"),
        ServerSideEncryption = "aws:kms",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplebucket, err := s3.NewBucketV2(ctx, "examplebucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: examplebucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "exampleBucketObject", &s3.BucketObjectArgs{
			Key:                  pulumi.String("someobject"),
			Bucket:               examplebucket.ID(),
			Source:               pulumi.NewFileAsset("index.html"),
			ServerSideEncryption: pulumi.String("aws:kms"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketObject;
import com.pulumi.aws.s3.BucketObjectArgs;
import com.pulumi.asset.FileAsset;
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
        var examplebucket = new BucketV2("examplebucket");

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(examplebucket.id())
            .acl("private")
            .build());

        var exampleBucketObject = new BucketObject("exampleBucketObject", BucketObjectArgs.builder()        
            .key("someobject")
            .bucket(examplebucket.id())
            .source(new FileAsset("index.html"))
            .serverSideEncryption("aws:kms")
            .build());

    }
}
```
```yaml
resources:
  examplebucket:
    type: aws:s3:BucketV2
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${examplebucket.id}
      acl: private
  exampleBucketObject:
    type: aws:s3:BucketObject
    properties:
      key: someobject
      bucket: ${examplebucket.id}
      source:
        fn::FileAsset: index.html
      serverSideEncryption: aws:kms
```
{{% /example %}}
{{% example %}}
### Server Side Encryption with AWS-Managed Key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplebucket = new aws.s3.BucketV2("examplebucket", {});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: examplebucket.id,
    acl: "private",
});
const exampleBucketObject = new aws.s3.BucketObject("exampleBucketObject", {
    key: "someobject",
    bucket: examplebucket.id,
    source: new pulumi.asset.FileAsset("index.html"),
    serverSideEncryption: "AES256",
});
```
```python
import pulumi
import pulumi_aws as aws

examplebucket = aws.s3.BucketV2("examplebucket")
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=examplebucket.id,
    acl="private")
example_bucket_object = aws.s3.BucketObject("exampleBucketObject",
    key="someobject",
    bucket=examplebucket.id,
    source=pulumi.FileAsset("index.html"),
    server_side_encryption="AES256")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplebucket = new Aws.S3.BucketV2("examplebucket");

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = examplebucket.Id,
        Acl = "private",
    });

    var exampleBucketObject = new Aws.S3.BucketObject("exampleBucketObject", new()
    {
        Key = "someobject",
        Bucket = examplebucket.Id,
        Source = new FileAsset("index.html"),
        ServerSideEncryption = "AES256",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplebucket, err := s3.NewBucketV2(ctx, "examplebucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: examplebucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "exampleBucketObject", &s3.BucketObjectArgs{
			Key:                  pulumi.String("someobject"),
			Bucket:               examplebucket.ID(),
			Source:               pulumi.NewFileAsset("index.html"),
			ServerSideEncryption: pulumi.String("AES256"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketObject;
import com.pulumi.aws.s3.BucketObjectArgs;
import com.pulumi.asset.FileAsset;
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
        var examplebucket = new BucketV2("examplebucket");

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(examplebucket.id())
            .acl("private")
            .build());

        var exampleBucketObject = new BucketObject("exampleBucketObject", BucketObjectArgs.builder()        
            .key("someobject")
            .bucket(examplebucket.id())
            .source(new FileAsset("index.html"))
            .serverSideEncryption("AES256")
            .build());

    }
}
```
```yaml
resources:
  examplebucket:
    type: aws:s3:BucketV2
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${examplebucket.id}
      acl: private
  exampleBucketObject:
    type: aws:s3:BucketObject
    properties:
      key: someobject
      bucket: ${examplebucket.id}
      source:
        fn::FileAsset: index.html
      serverSideEncryption: AES256
```
{{% /example %}}
{{% example %}}
### S3 Object Lock

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplebucket = new aws.s3.BucketV2("examplebucket", {objectLockEnabled: true});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: examplebucket.id,
    acl: "private",
});
const exampleBucketVersioningV2 = new aws.s3.BucketVersioningV2("exampleBucketVersioningV2", {
    bucket: examplebucket.id,
    versioningConfiguration: {
        status: "Enabled",
    },
});
const exampleBucketObject = new aws.s3.BucketObject("exampleBucketObject", {
    key: "someobject",
    bucket: examplebucket.id,
    source: new pulumi.asset.FileAsset("important.txt"),
    objectLockLegalHoldStatus: "ON",
    objectLockMode: "GOVERNANCE",
    objectLockRetainUntilDate: "2021-12-31T23:59:60Z",
    forceDestroy: true,
}, {
    dependsOn: [exampleBucketVersioningV2],
});
```
```python
import pulumi
import pulumi_aws as aws

examplebucket = aws.s3.BucketV2("examplebucket", object_lock_enabled=True)
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=examplebucket.id,
    acl="private")
example_bucket_versioning_v2 = aws.s3.BucketVersioningV2("exampleBucketVersioningV2",
    bucket=examplebucket.id,
    versioning_configuration=aws.s3.BucketVersioningV2VersioningConfigurationArgs(
        status="Enabled",
    ))
example_bucket_object = aws.s3.BucketObject("exampleBucketObject",
    key="someobject",
    bucket=examplebucket.id,
    source=pulumi.FileAsset("important.txt"),
    object_lock_legal_hold_status="ON",
    object_lock_mode="GOVERNANCE",
    object_lock_retain_until_date="2021-12-31T23:59:60Z",
    force_destroy=True,
    opts=pulumi.ResourceOptions(depends_on=[example_bucket_versioning_v2]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplebucket = new Aws.S3.BucketV2("examplebucket", new()
    {
        ObjectLockEnabled = true,
    });

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = examplebucket.Id,
        Acl = "private",
    });

    var exampleBucketVersioningV2 = new Aws.S3.BucketVersioningV2("exampleBucketVersioningV2", new()
    {
        Bucket = examplebucket.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningV2VersioningConfigurationArgs
        {
            Status = "Enabled",
        },
    });

    var exampleBucketObject = new Aws.S3.BucketObject("exampleBucketObject", new()
    {
        Key = "someobject",
        Bucket = examplebucket.Id,
        Source = new FileAsset("important.txt"),
        ObjectLockLegalHoldStatus = "ON",
        ObjectLockMode = "GOVERNANCE",
        ObjectLockRetainUntilDate = "2021-12-31T23:59:60Z",
        ForceDestroy = true,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleBucketVersioningV2,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplebucket, err := s3.NewBucketV2(ctx, "examplebucket", &s3.BucketV2Args{
			ObjectLockEnabled: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: examplebucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		exampleBucketVersioningV2, err := s3.NewBucketVersioningV2(ctx, "exampleBucketVersioningV2", &s3.BucketVersioningV2Args{
			Bucket: examplebucket.ID(),
			VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObject(ctx, "exampleBucketObject", &s3.BucketObjectArgs{
			Key:                       pulumi.String("someobject"),
			Bucket:                    examplebucket.ID(),
			Source:                    pulumi.NewFileAsset("important.txt"),
			ObjectLockLegalHoldStatus: pulumi.String("ON"),
			ObjectLockMode:            pulumi.String("GOVERNANCE"),
			ObjectLockRetainUntilDate: pulumi.String("2021-12-31T23:59:60Z"),
			ForceDestroy:              pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleBucketVersioningV2,
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketVersioningV2;
import com.pulumi.aws.s3.BucketVersioningV2Args;
import com.pulumi.aws.s3.inputs.BucketVersioningV2VersioningConfigurationArgs;
import com.pulumi.aws.s3.BucketObject;
import com.pulumi.aws.s3.BucketObjectArgs;
import com.pulumi.resources.CustomResourceOptions;
import com.pulumi.asset.FileAsset;
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
        var examplebucket = new BucketV2("examplebucket", BucketV2Args.builder()        
            .objectLockEnabled(true)
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(examplebucket.id())
            .acl("private")
            .build());

        var exampleBucketVersioningV2 = new BucketVersioningV2("exampleBucketVersioningV2", BucketVersioningV2Args.builder()        
            .bucket(examplebucket.id())
            .versioningConfiguration(BucketVersioningV2VersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build());

        var exampleBucketObject = new BucketObject("exampleBucketObject", BucketObjectArgs.builder()        
            .key("someobject")
            .bucket(examplebucket.id())
            .source(new FileAsset("important.txt"))
            .objectLockLegalHoldStatus("ON")
            .objectLockMode("GOVERNANCE")
            .objectLockRetainUntilDate("2021-12-31T23:59:60Z")
            .forceDestroy(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleBucketVersioningV2)
                .build());

    }
}
```
```yaml
resources:
  examplebucket:
    type: aws:s3:BucketV2
    properties:
      objectLockEnabled: true
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${examplebucket.id}
      acl: private
  exampleBucketVersioningV2:
    type: aws:s3:BucketVersioningV2
    properties:
      bucket: ${examplebucket.id}
      versioningConfiguration:
        status: Enabled
  exampleBucketObject:
    type: aws:s3:BucketObject
    properties:
      key: someobject
      bucket: ${examplebucket.id}
      source:
        fn::FileAsset: important.txt
      objectLockLegalHoldStatus: ON
      objectLockMode: GOVERNANCE
      objectLockRetainUntilDate: 2021-12-31T23:59:60Z
      forceDestroy: true
    options:
      dependson:
        - ${exampleBucketVersioningV2}
```
{{% /example %}}
{{% /examples %}}

## Import

Import using S3 URL syntax:

__Using `pulumi import` to import__ objects using the `id` or S3 URL. For example:

Import using the `id`, which is the bucket name and the key together:

```sh
 $ pulumi import aws:s3/bucketObject:BucketObject example some-bucket-name/some/key.txt
```
 Import using S3 URL syntax:

```sh
 $ pulumi import aws:s3/bucketObject:BucketObject example s3://some-bucket-name/some/key.txt
```
 