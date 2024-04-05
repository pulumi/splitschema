Provides an Amazon Connect Instance Storage Config resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}
### Storage Config Kinesis Firehose Config

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.InstanceStorageConfig("example", {
    instanceId: aws_connect_instance.example.id,
    resourceType: "CONTACT_TRACE_RECORDS",
    storageConfig: {
        kinesisFirehoseConfig: {
            firehoseArn: aws_kinesis_firehose_delivery_stream.example.arn,
        },
        storageType: "KINESIS_FIREHOSE",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.InstanceStorageConfig("example",
    instance_id=aws_connect_instance["example"]["id"],
    resource_type="CONTACT_TRACE_RECORDS",
    storage_config=aws.connect.InstanceStorageConfigStorageConfigArgs(
        kinesis_firehose_config=aws.connect.InstanceStorageConfigStorageConfigKinesisFirehoseConfigArgs(
            firehose_arn=aws_kinesis_firehose_delivery_stream["example"]["arn"],
        ),
        storage_type="KINESIS_FIREHOSE",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.InstanceStorageConfig("example", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        ResourceType = "CONTACT_TRACE_RECORDS",
        StorageConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigArgs
        {
            KinesisFirehoseConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigKinesisFirehoseConfigArgs
            {
                FirehoseArn = aws_kinesis_firehose_delivery_stream.Example.Arn,
            },
            StorageType = "KINESIS_FIREHOSE",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstanceStorageConfig(ctx, "example", &connect.InstanceStorageConfigArgs{
			InstanceId:   pulumi.Any(aws_connect_instance.Example.Id),
			ResourceType: pulumi.String("CONTACT_TRACE_RECORDS"),
			StorageConfig: &connect.InstanceStorageConfigStorageConfigArgs{
				KinesisFirehoseConfig: &connect.InstanceStorageConfigStorageConfigKinesisFirehoseConfigArgs{
					FirehoseArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
				},
				StorageType: pulumi.String("KINESIS_FIREHOSE"),
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
import com.pulumi.aws.connect.InstanceStorageConfig;
import com.pulumi.aws.connect.InstanceStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigKinesisFirehoseConfigArgs;
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
        var example = new InstanceStorageConfig("example", InstanceStorageConfigArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .resourceType("CONTACT_TRACE_RECORDS")
            .storageConfig(InstanceStorageConfigStorageConfigArgs.builder()
                .kinesisFirehoseConfig(InstanceStorageConfigStorageConfigKinesisFirehoseConfigArgs.builder()
                    .firehoseArn(aws_kinesis_firehose_delivery_stream.example().arn())
                    .build())
                .storageType("KINESIS_FIREHOSE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:InstanceStorageConfig
    properties:
      instanceId: ${aws_connect_instance.example.id}
      resourceType: CONTACT_TRACE_RECORDS
      storageConfig:
        kinesisFirehoseConfig:
          firehoseArn: ${aws_kinesis_firehose_delivery_stream.example.arn}
        storageType: KINESIS_FIREHOSE
```
{{% /example %}}
{{% example %}}
### Storage Config Kinesis Stream Config

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.InstanceStorageConfig("example", {
    instanceId: aws_connect_instance.example.id,
    resourceType: "CONTACT_TRACE_RECORDS",
    storageConfig: {
        kinesisStreamConfig: {
            streamArn: aws_kinesis_stream.example.arn,
        },
        storageType: "KINESIS_STREAM",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.InstanceStorageConfig("example",
    instance_id=aws_connect_instance["example"]["id"],
    resource_type="CONTACT_TRACE_RECORDS",
    storage_config=aws.connect.InstanceStorageConfigStorageConfigArgs(
        kinesis_stream_config=aws.connect.InstanceStorageConfigStorageConfigKinesisStreamConfigArgs(
            stream_arn=aws_kinesis_stream["example"]["arn"],
        ),
        storage_type="KINESIS_STREAM",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.InstanceStorageConfig("example", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        ResourceType = "CONTACT_TRACE_RECORDS",
        StorageConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigArgs
        {
            KinesisStreamConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigKinesisStreamConfigArgs
            {
                StreamArn = aws_kinesis_stream.Example.Arn,
            },
            StorageType = "KINESIS_STREAM",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstanceStorageConfig(ctx, "example", &connect.InstanceStorageConfigArgs{
			InstanceId:   pulumi.Any(aws_connect_instance.Example.Id),
			ResourceType: pulumi.String("CONTACT_TRACE_RECORDS"),
			StorageConfig: &connect.InstanceStorageConfigStorageConfigArgs{
				KinesisStreamConfig: &connect.InstanceStorageConfigStorageConfigKinesisStreamConfigArgs{
					StreamArn: pulumi.Any(aws_kinesis_stream.Example.Arn),
				},
				StorageType: pulumi.String("KINESIS_STREAM"),
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
import com.pulumi.aws.connect.InstanceStorageConfig;
import com.pulumi.aws.connect.InstanceStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigKinesisStreamConfigArgs;
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
        var example = new InstanceStorageConfig("example", InstanceStorageConfigArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .resourceType("CONTACT_TRACE_RECORDS")
            .storageConfig(InstanceStorageConfigStorageConfigArgs.builder()
                .kinesisStreamConfig(InstanceStorageConfigStorageConfigKinesisStreamConfigArgs.builder()
                    .streamArn(aws_kinesis_stream.example().arn())
                    .build())
                .storageType("KINESIS_STREAM")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:InstanceStorageConfig
    properties:
      instanceId: ${aws_connect_instance.example.id}
      resourceType: CONTACT_TRACE_RECORDS
      storageConfig:
        kinesisStreamConfig:
          streamArn: ${aws_kinesis_stream.example.arn}
        storageType: KINESIS_STREAM
```
{{% /example %}}
{{% example %}}
### Storage Config Kinesis Video Stream Config

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.InstanceStorageConfig("example", {
    instanceId: aws_connect_instance.example.id,
    resourceType: "MEDIA_STREAMS",
    storageConfig: {
        kinesisVideoStreamConfig: {
            prefix: "example",
            retentionPeriodHours: 3,
            encryptionConfig: {
                encryptionType: "KMS",
                keyId: aws_kms_key.example.arn,
            },
        },
        storageType: "KINESIS_VIDEO_STREAM",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.InstanceStorageConfig("example",
    instance_id=aws_connect_instance["example"]["id"],
    resource_type="MEDIA_STREAMS",
    storage_config=aws.connect.InstanceStorageConfigStorageConfigArgs(
        kinesis_video_stream_config=aws.connect.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigArgs(
            prefix="example",
            retention_period_hours=3,
            encryption_config=aws.connect.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigArgs(
                encryption_type="KMS",
                key_id=aws_kms_key["example"]["arn"],
            ),
        ),
        storage_type="KINESIS_VIDEO_STREAM",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.InstanceStorageConfig("example", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        ResourceType = "MEDIA_STREAMS",
        StorageConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigArgs
        {
            KinesisVideoStreamConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigArgs
            {
                Prefix = "example",
                RetentionPeriodHours = 3,
                EncryptionConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigArgs
                {
                    EncryptionType = "KMS",
                    KeyId = aws_kms_key.Example.Arn,
                },
            },
            StorageType = "KINESIS_VIDEO_STREAM",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstanceStorageConfig(ctx, "example", &connect.InstanceStorageConfigArgs{
			InstanceId:   pulumi.Any(aws_connect_instance.Example.Id),
			ResourceType: pulumi.String("MEDIA_STREAMS"),
			StorageConfig: &connect.InstanceStorageConfigStorageConfigArgs{
				KinesisVideoStreamConfig: &connect.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigArgs{
					Prefix:               pulumi.String("example"),
					RetentionPeriodHours: pulumi.Int(3),
					EncryptionConfig: &connect.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigArgs{
						EncryptionType: pulumi.String("KMS"),
						KeyId:          pulumi.Any(aws_kms_key.Example.Arn),
					},
				},
				StorageType: pulumi.String("KINESIS_VIDEO_STREAM"),
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
import com.pulumi.aws.connect.InstanceStorageConfig;
import com.pulumi.aws.connect.InstanceStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigArgs;
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
        var example = new InstanceStorageConfig("example", InstanceStorageConfigArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .resourceType("MEDIA_STREAMS")
            .storageConfig(InstanceStorageConfigStorageConfigArgs.builder()
                .kinesisVideoStreamConfig(InstanceStorageConfigStorageConfigKinesisVideoStreamConfigArgs.builder()
                    .prefix("example")
                    .retentionPeriodHours(3)
                    .encryptionConfig(InstanceStorageConfigStorageConfigKinesisVideoStreamConfigEncryptionConfigArgs.builder()
                        .encryptionType("KMS")
                        .keyId(aws_kms_key.example().arn())
                        .build())
                    .build())
                .storageType("KINESIS_VIDEO_STREAM")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:InstanceStorageConfig
    properties:
      instanceId: ${aws_connect_instance.example.id}
      resourceType: MEDIA_STREAMS
      storageConfig:
        kinesisVideoStreamConfig:
          prefix: example
          retentionPeriodHours: 3
          encryptionConfig:
            encryptionType: KMS
            keyId: ${aws_kms_key.example.arn}
        storageType: KINESIS_VIDEO_STREAM
```
{{% /example %}}
{{% example %}}
### Storage Config S3 Config

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.InstanceStorageConfig("example", {
    instanceId: aws_connect_instance.example.id,
    resourceType: "CHAT_TRANSCRIPTS",
    storageConfig: {
        s3Config: {
            bucketName: aws_s3_bucket.example.id,
            bucketPrefix: "example",
        },
        storageType: "S3",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.InstanceStorageConfig("example",
    instance_id=aws_connect_instance["example"]["id"],
    resource_type="CHAT_TRANSCRIPTS",
    storage_config=aws.connect.InstanceStorageConfigStorageConfigArgs(
        s3_config=aws.connect.InstanceStorageConfigStorageConfigS3ConfigArgs(
            bucket_name=aws_s3_bucket["example"]["id"],
            bucket_prefix="example",
        ),
        storage_type="S3",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.InstanceStorageConfig("example", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        ResourceType = "CHAT_TRANSCRIPTS",
        StorageConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigArgs
        {
            S3Config = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigS3ConfigArgs
            {
                BucketName = aws_s3_bucket.Example.Id,
                BucketPrefix = "example",
            },
            StorageType = "S3",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstanceStorageConfig(ctx, "example", &connect.InstanceStorageConfigArgs{
			InstanceId:   pulumi.Any(aws_connect_instance.Example.Id),
			ResourceType: pulumi.String("CHAT_TRANSCRIPTS"),
			StorageConfig: &connect.InstanceStorageConfigStorageConfigArgs{
				S3Config: &connect.InstanceStorageConfigStorageConfigS3ConfigArgs{
					BucketName:   pulumi.Any(aws_s3_bucket.Example.Id),
					BucketPrefix: pulumi.String("example"),
				},
				StorageType: pulumi.String("S3"),
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
import com.pulumi.aws.connect.InstanceStorageConfig;
import com.pulumi.aws.connect.InstanceStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigS3ConfigArgs;
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
        var example = new InstanceStorageConfig("example", InstanceStorageConfigArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .resourceType("CHAT_TRANSCRIPTS")
            .storageConfig(InstanceStorageConfigStorageConfigArgs.builder()
                .s3Config(InstanceStorageConfigStorageConfigS3ConfigArgs.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .bucketPrefix("example")
                    .build())
                .storageType("S3")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:InstanceStorageConfig
    properties:
      instanceId: ${aws_connect_instance.example.id}
      resourceType: CHAT_TRANSCRIPTS
      storageConfig:
        s3Config:
          bucketName: ${aws_s3_bucket.example.id}
          bucketPrefix: example
        storageType: S3
```
{{% /example %}}
{{% example %}}
### Storage Config S3 Config with Encryption Config

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.InstanceStorageConfig("example", {
    instanceId: aws_connect_instance.example.id,
    resourceType: "CHAT_TRANSCRIPTS",
    storageConfig: {
        s3Config: {
            bucketName: aws_s3_bucket.example.id,
            bucketPrefix: "example",
            encryptionConfig: {
                encryptionType: "KMS",
                keyId: aws_kms_key.example.arn,
            },
        },
        storageType: "S3",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.InstanceStorageConfig("example",
    instance_id=aws_connect_instance["example"]["id"],
    resource_type="CHAT_TRANSCRIPTS",
    storage_config=aws.connect.InstanceStorageConfigStorageConfigArgs(
        s3_config=aws.connect.InstanceStorageConfigStorageConfigS3ConfigArgs(
            bucket_name=aws_s3_bucket["example"]["id"],
            bucket_prefix="example",
            encryption_config=aws.connect.InstanceStorageConfigStorageConfigS3ConfigEncryptionConfigArgs(
                encryption_type="KMS",
                key_id=aws_kms_key["example"]["arn"],
            ),
        ),
        storage_type="S3",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.InstanceStorageConfig("example", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        ResourceType = "CHAT_TRANSCRIPTS",
        StorageConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigArgs
        {
            S3Config = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigS3ConfigArgs
            {
                BucketName = aws_s3_bucket.Example.Id,
                BucketPrefix = "example",
                EncryptionConfig = new Aws.Connect.Inputs.InstanceStorageConfigStorageConfigS3ConfigEncryptionConfigArgs
                {
                    EncryptionType = "KMS",
                    KeyId = aws_kms_key.Example.Arn,
                },
            },
            StorageType = "S3",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstanceStorageConfig(ctx, "example", &connect.InstanceStorageConfigArgs{
			InstanceId:   pulumi.Any(aws_connect_instance.Example.Id),
			ResourceType: pulumi.String("CHAT_TRANSCRIPTS"),
			StorageConfig: &connect.InstanceStorageConfigStorageConfigArgs{
				S3Config: &connect.InstanceStorageConfigStorageConfigS3ConfigArgs{
					BucketName:   pulumi.Any(aws_s3_bucket.Example.Id),
					BucketPrefix: pulumi.String("example"),
					EncryptionConfig: &connect.InstanceStorageConfigStorageConfigS3ConfigEncryptionConfigArgs{
						EncryptionType: pulumi.String("KMS"),
						KeyId:          pulumi.Any(aws_kms_key.Example.Arn),
					},
				},
				StorageType: pulumi.String("S3"),
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
import com.pulumi.aws.connect.InstanceStorageConfig;
import com.pulumi.aws.connect.InstanceStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigS3ConfigArgs;
import com.pulumi.aws.connect.inputs.InstanceStorageConfigStorageConfigS3ConfigEncryptionConfigArgs;
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
        var example = new InstanceStorageConfig("example", InstanceStorageConfigArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .resourceType("CHAT_TRANSCRIPTS")
            .storageConfig(InstanceStorageConfigStorageConfigArgs.builder()
                .s3Config(InstanceStorageConfigStorageConfigS3ConfigArgs.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .bucketPrefix("example")
                    .encryptionConfig(InstanceStorageConfigStorageConfigS3ConfigEncryptionConfigArgs.builder()
                        .encryptionType("KMS")
                        .keyId(aws_kms_key.example().arn())
                        .build())
                    .build())
                .storageType("S3")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:InstanceStorageConfig
    properties:
      instanceId: ${aws_connect_instance.example.id}
      resourceType: CHAT_TRANSCRIPTS
      storageConfig:
        s3Config:
          bucketName: ${aws_s3_bucket.example.id}
          bucketPrefix: example
          encryptionConfig:
            encryptionType: KMS
            keyId: ${aws_kms_key.example.arn}
        storageType: S3
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Instance Storage Configs using the `instance_id`, `association_id`, and `resource_type` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/instanceStorageConfig:InstanceStorageConfig example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5:CHAT_TRANSCRIPTS
```
 