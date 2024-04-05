Resource for managing a Verified Access Logging Configuration.

{{% examples %}}
## Example Usage
{{% example %}}
### With CloudWatch Logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.InstanceLoggingConfiguration("example", {
    accessLogs: {
        cloudwatchLogs: {
            enabled: true,
            logGroup: aws_cloudwatch_log_group.example.id,
        },
    },
    verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.InstanceLoggingConfiguration("example",
    access_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs(
        cloudwatch_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs(
            enabled=True,
            log_group=aws_cloudwatch_log_group["example"]["id"],
        ),
    ),
    verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.InstanceLoggingConfiguration("example", new()
    {
        AccessLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsArgs
        {
            CloudwatchLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs
            {
                Enabled = true,
                LogGroup = aws_cloudwatch_log_group.Example.Id,
            },
        },
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstanceLoggingConfiguration(ctx, "example", &verifiedaccess.InstanceLoggingConfigurationArgs{
			AccessLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs{
				CloudwatchLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs{
					Enabled:  pulumi.Bool(true),
					LogGroup: pulumi.Any(aws_cloudwatch_log_group.Example.Id),
				},
			},
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfiguration;
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfigurationArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs;
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
        var example = new InstanceLoggingConfiguration("example", InstanceLoggingConfigurationArgs.builder()        
            .accessLogs(InstanceLoggingConfigurationAccessLogsArgs.builder()
                .cloudwatchLogs(InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs.builder()
                    .enabled(true)
                    .logGroup(aws_cloudwatch_log_group.example().id())
                    .build())
                .build())
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:InstanceLoggingConfiguration
    properties:
      accessLogs:
        cloudwatchLogs:
          enabled: true
          logGroup: ${aws_cloudwatch_log_group.example.id}
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% example %}}
### With Kinesis Data Firehose Logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.InstanceLoggingConfiguration("example", {
    accessLogs: {
        kinesisDataFirehose: {
            deliveryStream: aws_kinesis_firehose_delivery_stream.example.name,
            enabled: true,
        },
    },
    verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.InstanceLoggingConfiguration("example",
    access_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs(
        kinesis_data_firehose=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs(
            delivery_stream=aws_kinesis_firehose_delivery_stream["example"]["name"],
            enabled=True,
        ),
    ),
    verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.InstanceLoggingConfiguration("example", new()
    {
        AccessLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsArgs
        {
            KinesisDataFirehose = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs
            {
                DeliveryStream = aws_kinesis_firehose_delivery_stream.Example.Name,
                Enabled = true,
            },
        },
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstanceLoggingConfiguration(ctx, "example", &verifiedaccess.InstanceLoggingConfigurationArgs{
			AccessLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs{
				KinesisDataFirehose: &verifiedaccess.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs{
					DeliveryStream: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Name),
					Enabled:        pulumi.Bool(true),
				},
			},
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfiguration;
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfigurationArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs;
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
        var example = new InstanceLoggingConfiguration("example", InstanceLoggingConfigurationArgs.builder()        
            .accessLogs(InstanceLoggingConfigurationAccessLogsArgs.builder()
                .kinesisDataFirehose(InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs.builder()
                    .deliveryStream(aws_kinesis_firehose_delivery_stream.example().name())
                    .enabled(true)
                    .build())
                .build())
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:InstanceLoggingConfiguration
    properties:
      accessLogs:
        kinesisDataFirehose:
          deliveryStream: ${aws_kinesis_firehose_delivery_stream.example.name}
          enabled: true
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% example %}}
### With S3 logging

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.InstanceLoggingConfiguration("example", {
    accessLogs: {
        s3: {
            bucketName: aws_s3_bucket.example.id,
            enabled: true,
            prefix: "example",
        },
    },
    verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.InstanceLoggingConfiguration("example",
    access_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs(
        s3=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsS3Args(
            bucket_name=aws_s3_bucket["example"]["id"],
            enabled=True,
            prefix="example",
        ),
    ),
    verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.InstanceLoggingConfiguration("example", new()
    {
        AccessLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsArgs
        {
            S3 = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsS3Args
            {
                BucketName = aws_s3_bucket.Example.Id,
                Enabled = true,
                Prefix = "example",
            },
        },
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstanceLoggingConfiguration(ctx, "example", &verifiedaccess.InstanceLoggingConfigurationArgs{
			AccessLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs{
				S3: &verifiedaccess.InstanceLoggingConfigurationAccessLogsS3Args{
					BucketName: pulumi.Any(aws_s3_bucket.Example.Id),
					Enabled:    pulumi.Bool(true),
					Prefix:     pulumi.String("example"),
				},
			},
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfiguration;
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfigurationArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsS3Args;
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
        var example = new InstanceLoggingConfiguration("example", InstanceLoggingConfigurationArgs.builder()        
            .accessLogs(InstanceLoggingConfigurationAccessLogsArgs.builder()
                .s3(InstanceLoggingConfigurationAccessLogsS3Args.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .enabled(true)
                    .prefix("example")
                    .build())
                .build())
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:InstanceLoggingConfiguration
    properties:
      accessLogs:
        s3:
          bucketName: ${aws_s3_bucket.example.id}
          enabled: true
          prefix: example
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% example %}}
### With all three logging options

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.InstanceLoggingConfiguration("example", {
    accessLogs: {
        cloudwatchLogs: {
            enabled: true,
            logGroup: aws_cloudwatch_log_group.example.id,
        },
        kinesisDataFirehose: {
            deliveryStream: aws_kinesis_firehose_delivery_stream.example.name,
            enabled: true,
        },
        s3: {
            bucketName: aws_s3_bucket.example.id,
            enabled: true,
        },
    },
    verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.InstanceLoggingConfiguration("example",
    access_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs(
        cloudwatch_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs(
            enabled=True,
            log_group=aws_cloudwatch_log_group["example"]["id"],
        ),
        kinesis_data_firehose=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs(
            delivery_stream=aws_kinesis_firehose_delivery_stream["example"]["name"],
            enabled=True,
        ),
        s3=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsS3Args(
            bucket_name=aws_s3_bucket["example"]["id"],
            enabled=True,
        ),
    ),
    verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.InstanceLoggingConfiguration("example", new()
    {
        AccessLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsArgs
        {
            CloudwatchLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs
            {
                Enabled = true,
                LogGroup = aws_cloudwatch_log_group.Example.Id,
            },
            KinesisDataFirehose = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs
            {
                DeliveryStream = aws_kinesis_firehose_delivery_stream.Example.Name,
                Enabled = true,
            },
            S3 = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsS3Args
            {
                BucketName = aws_s3_bucket.Example.Id,
                Enabled = true,
            },
        },
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstanceLoggingConfiguration(ctx, "example", &verifiedaccess.InstanceLoggingConfigurationArgs{
			AccessLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs{
				CloudwatchLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs{
					Enabled:  pulumi.Bool(true),
					LogGroup: pulumi.Any(aws_cloudwatch_log_group.Example.Id),
				},
				KinesisDataFirehose: &verifiedaccess.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs{
					DeliveryStream: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Name),
					Enabled:        pulumi.Bool(true),
				},
				S3: &verifiedaccess.InstanceLoggingConfigurationAccessLogsS3Args{
					BucketName: pulumi.Any(aws_s3_bucket.Example.Id),
					Enabled:    pulumi.Bool(true),
				},
			},
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfiguration;
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfigurationArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsS3Args;
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
        var example = new InstanceLoggingConfiguration("example", InstanceLoggingConfigurationArgs.builder()        
            .accessLogs(InstanceLoggingConfigurationAccessLogsArgs.builder()
                .cloudwatchLogs(InstanceLoggingConfigurationAccessLogsCloudwatchLogsArgs.builder()
                    .enabled(true)
                    .logGroup(aws_cloudwatch_log_group.example().id())
                    .build())
                .kinesisDataFirehose(InstanceLoggingConfigurationAccessLogsKinesisDataFirehoseArgs.builder()
                    .deliveryStream(aws_kinesis_firehose_delivery_stream.example().name())
                    .enabled(true)
                    .build())
                .s3(InstanceLoggingConfigurationAccessLogsS3Args.builder()
                    .bucketName(aws_s3_bucket.example().id())
                    .enabled(true)
                    .build())
                .build())
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:InstanceLoggingConfiguration
    properties:
      accessLogs:
        cloudwatchLogs:
          enabled: true
          logGroup: ${aws_cloudwatch_log_group.example.id}
        kinesisDataFirehose:
          deliveryStream: ${aws_kinesis_firehose_delivery_stream.example.name}
          enabled: true
        s3:
          bucketName: ${aws_s3_bucket.example.id}
          enabled: true
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% example %}}
### With `include_trust_context`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.InstanceLoggingConfiguration("example", {
    accessLogs: {
        includeTrustContext: true,
    },
    verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.InstanceLoggingConfiguration("example",
    access_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs(
        include_trust_context=True,
    ),
    verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.InstanceLoggingConfiguration("example", new()
    {
        AccessLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsArgs
        {
            IncludeTrustContext = true,
        },
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstanceLoggingConfiguration(ctx, "example", &verifiedaccess.InstanceLoggingConfigurationArgs{
			AccessLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs{
				IncludeTrustContext: pulumi.Bool(true),
			},
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfiguration;
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfigurationArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsArgs;
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
        var example = new InstanceLoggingConfiguration("example", InstanceLoggingConfigurationArgs.builder()        
            .accessLogs(InstanceLoggingConfigurationAccessLogsArgs.builder()
                .includeTrustContext(true)
                .build())
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:InstanceLoggingConfiguration
    properties:
      accessLogs:
        includeTrustContext: true
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% example %}}
### With `log_version`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.InstanceLoggingConfiguration("example", {
    accessLogs: {
        logVersion: "ocsf-1.0.0-rc.2",
    },
    verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.InstanceLoggingConfiguration("example",
    access_logs=aws.verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs(
        log_version="ocsf-1.0.0-rc.2",
    ),
    verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.InstanceLoggingConfiguration("example", new()
    {
        AccessLogs = new Aws.VerifiedAccess.Inputs.InstanceLoggingConfigurationAccessLogsArgs
        {
            LogVersion = "ocsf-1.0.0-rc.2",
        },
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewInstanceLoggingConfiguration(ctx, "example", &verifiedaccess.InstanceLoggingConfigurationArgs{
			AccessLogs: &verifiedaccess.InstanceLoggingConfigurationAccessLogsArgs{
				LogVersion: pulumi.String("ocsf-1.0.0-rc.2"),
			},
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfiguration;
import com.pulumi.aws.verifiedaccess.InstanceLoggingConfigurationArgs;
import com.pulumi.aws.verifiedaccess.inputs.InstanceLoggingConfigurationAccessLogsArgs;
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
        var example = new InstanceLoggingConfiguration("example", InstanceLoggingConfigurationArgs.builder()        
            .accessLogs(InstanceLoggingConfigurationAccessLogsArgs.builder()
                .logVersion("ocsf-1.0.0-rc.2")
                .build())
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:InstanceLoggingConfiguration
    properties:
      accessLogs:
        logVersion: ocsf-1.0.0-rc.2
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Verified Access Logging Configuration using the Verified Access Instance `id`. For example:

```sh
 $ pulumi import aws:verifiedaccess/instanceLoggingConfiguration:InstanceLoggingConfiguration example vai-1234567890abcdef0
```
 