Provides an independent configuration resource for S3 bucket [lifecycle configuration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html).

An S3 Lifecycle configuration consists of one or more Lifecycle rules. Each rule consists of the following:

* Rule metadata (`id` and `status`)
* Filter identifying objects to which the rule applies
* One or more transition or expiration actions

For more information see the Amazon S3 User Guide on [`Lifecycle Configuration Elements`](https://docs.aws.amazon.com/AmazonS3/latest/userguide/intro-lifecycle-rules.html).

> **NOTE:** S3 Buckets only support a single lifecycle configuration. Declaring multiple `aws.s3.BucketLifecycleConfigurationV2` resources to the same S3 Bucket will cause a perpetual difference in configuration.

> **NOTE:** Lifecycle configurations may take some time to fully propagate to all AWS S3 systems.
Running Pulumi operations shortly after creating a lifecycle configuration may result in changes that affect configuration idempotence.
See the Amazon S3 User Guide on [setting lifecycle configuration on a bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/how-to-set-lifecycle-configuration-intro.html).

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### With neither a filter nor prefix specified

The Lifecycle rule applies to a subset of objects based on the key name prefix (`""`).

This configuration is intended to replicate the default behavior of the `lifecycle_rule`
parameter in the AWS Provider `aws.s3.BucketV2` resource prior to `v4.0`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id:     pulumi.String("rule-1"),
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying an empty filter

The Lifecycle rule applies to all objects in the bucket.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {},
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = null,
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id:     pulumi.String("rule-1"),
					Filter: nil,
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter()
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter: {}
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying a filter using key prefixes

The Lifecycle rule applies to a subset of objects based on the key name prefix (`logs/`).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {
            prefix: "logs/",
        },
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            prefix="logs/",
        ),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    Prefix = "logs/",
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						Prefix: pulumi.String("logs/"),
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .prefix("logs/")
                    .build())
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            prefix: logs/
          status: Enabled
```

If you want to apply a Lifecycle action to a subset of objects based on different key name prefixes, specify separate rules.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [
        {
            id: "rule-1",
            filter: {
                prefix: "logs/",
            },
            status: "Enabled",
        },
        {
            id: "rule-2",
            filter: {
                prefix: "tmp/",
            },
            status: "Enabled",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[
        aws.s3.BucketLifecycleConfigurationV2RuleArgs(
            id="rule-1",
            filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
                prefix="logs/",
            ),
            status="Enabled",
        ),
        aws.s3.BucketLifecycleConfigurationV2RuleArgs(
            id="rule-2",
            filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
                prefix="tmp/",
            ),
            status="Enabled",
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
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    Prefix = "logs/",
                },
                Status = "Enabled",
            },
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-2",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    Prefix = "tmp/",
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						Prefix: pulumi.String("logs/"),
					},
					Status: pulumi.String("Enabled"),
				},
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-2"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						Prefix: pulumi.String("tmp/"),
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(            
                BucketLifecycleConfigurationV2RuleArgs.builder()
                    .id("rule-1")
                    .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                        .prefix("logs/")
                        .build())
                    .status("Enabled")
                    .build(),
                BucketLifecycleConfigurationV2RuleArgs.builder()
                    .id("rule-2")
                    .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                        .prefix("tmp/")
                        .build())
                    .status("Enabled")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            prefix: logs/
          status: Enabled
        - id: rule-2
          filter:
            prefix: tmp/
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying a filter based on an object tag

The Lifecycle rule specifies a filter based on a tag key and value. The rule then applies only to a subset of objects with the specific tag.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {
            tag: {
                key: "Name",
                value: "Staging",
            },
        },
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            tag=aws.s3.BucketLifecycleConfigurationV2RuleFilterTagArgs(
                key="Name",
                value="Staging",
            ),
        ),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    Tag = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterTagArgs
                    {
                        Key = "Name",
                        Value = "Staging",
                    },
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						Tag: &s3.BucketLifecycleConfigurationV2RuleFilterTagArgs{
							Key:   pulumi.String("Name"),
							Value: pulumi.String("Staging"),
						},
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterTagArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .tag(BucketLifecycleConfigurationV2RuleFilterTagArgs.builder()
                        .key("Name")
                        .value("Staging")
                        .build())
                    .build())
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            tag:
              key: Name
              value: Staging
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying a filter based on multiple tags

The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with two tags (with the specific tag keys and values). Notice `tags` is wrapped in the `and` configuration block.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {
            and: {
                tags: {
                    Key1: "Value1",
                    Key2: "Value2",
                },
            },
        },
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            and_=aws.s3.BucketLifecycleConfigurationV2RuleFilterAndArgs(
                tags={
                    "Key1": "Value1",
                    "Key2": "Value2",
                },
            ),
        ),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
                    {
                        Tags = 
                        {
                            { "Key1", "Value1" },
                            { "Key2", "Value2" },
                        },
                    },
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
							Tags: pulumi.StringMap{
								"Key1": pulumi.String("Value1"),
								"Key2": pulumi.String("Value2"),
							},
						},
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .and(BucketLifecycleConfigurationV2RuleFilterAndArgs.builder()
                        .tags(Map.ofEntries(
                            Map.entry("Key1", "Value1"),
                            Map.entry("Key2", "Value2")
                        ))
                        .build())
                    .build())
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            and:
              tags:
                Key1: Value1
                Key2: Value2
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying a filter based on both prefix and one or more tags

The Lifecycle rule directs Amazon S3 to perform lifecycle actions on objects with the specified prefix and two tags (with the specific tag keys and values). Notice both `prefix` and `tags` are wrapped in the `and` configuration block.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {
            and: {
                prefix: "logs/",
                tags: {
                    Key1: "Value1",
                    Key2: "Value2",
                },
            },
        },
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            and_=aws.s3.BucketLifecycleConfigurationV2RuleFilterAndArgs(
                prefix="logs/",
                tags={
                    "Key1": "Value1",
                    "Key2": "Value2",
                },
            ),
        ),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
                    {
                        Prefix = "logs/",
                        Tags = 
                        {
                            { "Key1", "Value1" },
                            { "Key2", "Value2" },
                        },
                    },
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
							Prefix: pulumi.String("logs/"),
							Tags: pulumi.StringMap{
								"Key1": pulumi.String("Value1"),
								"Key2": pulumi.String("Value2"),
							},
						},
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .and(BucketLifecycleConfigurationV2RuleFilterAndArgs.builder()
                        .prefix("logs/")
                        .tags(Map.ofEntries(
                            Map.entry("Key1", "Value1"),
                            Map.entry("Key2", "Value2")
                        ))
                        .build())
                    .build())
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            and:
              prefix: logs/
              tags:
                Key1: Value1
                Key2: Value2
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying a filter based on object size

Object size values are in bytes. Maximum filter size is 5TB. Some storage classes have minimum object size limitations, for more information, see [Comparing the Amazon S3 storage classes](https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html#sc-compare).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {
            objectSizeGreaterThan: "500",
        },
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            object_size_greater_than="500",
        ),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    ObjectSizeGreaterThan = "500",
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						ObjectSizeGreaterThan: pulumi.String("500"),
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .objectSizeGreaterThan(500)
                    .build())
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            objectSizeGreaterThan: 500
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Specifying a filter based on object size range and prefix

The `object_size_greater_than` must be less than the `object_size_less_than`. Notice both the object size range and prefix are wrapped in the `and` configuration block.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketLifecycleConfigurationV2("example", {
    bucket: aws_s3_bucket.bucket.id,
    rules: [{
        id: "rule-1",
        filter: {
            and: {
                prefix: "logs/",
                objectSizeGreaterThan: 500,
                objectSizeLessThan: 64000,
            },
        },
        status: "Enabled",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketLifecycleConfigurationV2("example",
    bucket=aws_s3_bucket["bucket"]["id"],
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="rule-1",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            and_=aws.s3.BucketLifecycleConfigurationV2RuleFilterAndArgs(
                prefix="logs/",
                object_size_greater_than=500,
                object_size_less_than=64000,
            ),
        ),
        status="Enabled",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketLifecycleConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "rule-1",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
                    {
                        Prefix = "logs/",
                        ObjectSizeGreaterThan = 500,
                        ObjectSizeLessThan = 64000,
                    },
                },
                Status = "Enabled",
            },
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
		_, err := s3.NewBucketLifecycleConfigurationV2(ctx, "example", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Bucket.Id),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("rule-1"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
							Prefix:                pulumi.String("logs/"),
							ObjectSizeGreaterThan: pulumi.Int(500),
							ObjectSizeLessThan:    pulumi.Int(64000),
						},
					},
					Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs;
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
        var example = new BucketLifecycleConfigurationV2("example", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.bucket().id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("rule-1")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .and(BucketLifecycleConfigurationV2RuleFilterAndArgs.builder()
                        .prefix("logs/")
                        .objectSizeGreaterThan(500)
                        .objectSizeLessThan(64000)
                        .build())
                    .build())
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.bucket.id}
      rules:
        - id: rule-1
          filter:
            and:
              prefix: logs/
              objectSizeGreaterThan: 500
              objectSizeLessThan: 64000
          status: Enabled
```
{{% /example %}}
{{% example %}}
### Creating a Lifecycle Configuration for a bucket with versioning

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucket = new aws.s3.BucketV2("bucket", {});
const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
    bucket: bucket.id,
    acl: "private",
});
const bucket_config = new aws.s3.BucketLifecycleConfigurationV2("bucket-config", {
    bucket: bucket.id,
    rules: [
        {
            id: "log",
            expiration: {
                days: 90,
            },
            filter: {
                and: {
                    prefix: "log/",
                    tags: {
                        rule: "log",
                        autoclean: "true",
                    },
                },
            },
            status: "Enabled",
            transitions: [
                {
                    days: 30,
                    storageClass: "STANDARD_IA",
                },
                {
                    days: 60,
                    storageClass: "GLACIER",
                },
            ],
        },
        {
            id: "tmp",
            filter: {
                prefix: "tmp/",
            },
            expiration: {
                date: "2023-01-13T00:00:00Z",
            },
            status: "Enabled",
        },
    ],
});
const versioningBucket = new aws.s3.BucketV2("versioningBucket", {});
const versioningBucketAcl = new aws.s3.BucketAclV2("versioningBucketAcl", {
    bucket: versioningBucket.id,
    acl: "private",
});
const versioning = new aws.s3.BucketVersioningV2("versioning", {
    bucket: versioningBucket.id,
    versioningConfiguration: {
        status: "Enabled",
    },
});
const versioning_bucket_config = new aws.s3.BucketLifecycleConfigurationV2("versioning-bucket-config", {
    bucket: versioningBucket.id,
    rules: [{
        id: "config",
        filter: {
            prefix: "config/",
        },
        noncurrentVersionExpiration: {
            noncurrentDays: 90,
        },
        noncurrentVersionTransitions: [
            {
                noncurrentDays: 30,
                storageClass: "STANDARD_IA",
            },
            {
                noncurrentDays: 60,
                storageClass: "GLACIER",
            },
        ],
        status: "Enabled",
    }],
}, {
    dependsOn: [versioning],
});
```
```python
import pulumi
import pulumi_aws as aws

bucket = aws.s3.BucketV2("bucket")
bucket_acl = aws.s3.BucketAclV2("bucketAcl",
    bucket=bucket.id,
    acl="private")
bucket_config = aws.s3.BucketLifecycleConfigurationV2("bucket-config",
    bucket=bucket.id,
    rules=[
        aws.s3.BucketLifecycleConfigurationV2RuleArgs(
            id="log",
            expiration=aws.s3.BucketLifecycleConfigurationV2RuleExpirationArgs(
                days=90,
            ),
            filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
                and_=aws.s3.BucketLifecycleConfigurationV2RuleFilterAndArgs(
                    prefix="log/",
                    tags={
                        "rule": "log",
                        "autoclean": "true",
                    },
                ),
            ),
            status="Enabled",
            transitions=[
                aws.s3.BucketLifecycleConfigurationV2RuleTransitionArgs(
                    days=30,
                    storage_class="STANDARD_IA",
                ),
                aws.s3.BucketLifecycleConfigurationV2RuleTransitionArgs(
                    days=60,
                    storage_class="GLACIER",
                ),
            ],
        ),
        aws.s3.BucketLifecycleConfigurationV2RuleArgs(
            id="tmp",
            filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
                prefix="tmp/",
            ),
            expiration=aws.s3.BucketLifecycleConfigurationV2RuleExpirationArgs(
                date="2023-01-13T00:00:00Z",
            ),
            status="Enabled",
        ),
    ])
versioning_bucket = aws.s3.BucketV2("versioningBucket")
versioning_bucket_acl = aws.s3.BucketAclV2("versioningBucketAcl",
    bucket=versioning_bucket.id,
    acl="private")
versioning = aws.s3.BucketVersioningV2("versioning",
    bucket=versioning_bucket.id,
    versioning_configuration=aws.s3.BucketVersioningV2VersioningConfigurationArgs(
        status="Enabled",
    ))
versioning_bucket_config = aws.s3.BucketLifecycleConfigurationV2("versioning-bucket-config",
    bucket=versioning_bucket.id,
    rules=[aws.s3.BucketLifecycleConfigurationV2RuleArgs(
        id="config",
        filter=aws.s3.BucketLifecycleConfigurationV2RuleFilterArgs(
            prefix="config/",
        ),
        noncurrent_version_expiration=aws.s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs(
            noncurrent_days=90,
        ),
        noncurrent_version_transitions=[
            aws.s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs(
                noncurrent_days=30,
                storage_class="STANDARD_IA",
            ),
            aws.s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs(
                noncurrent_days=60,
                storage_class="GLACIER",
            ),
        ],
        status="Enabled",
    )],
    opts=pulumi.ResourceOptions(depends_on=[versioning]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucket = new Aws.S3.BucketV2("bucket");

    var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    {
        Bucket = bucket.Id,
        Acl = "private",
    });

    var bucket_config = new Aws.S3.BucketLifecycleConfigurationV2("bucket-config", new()
    {
        Bucket = bucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "log",
                Expiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleExpirationArgs
                {
                    Days = 90,
                },
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    And = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs
                    {
                        Prefix = "log/",
                        Tags = 
                        {
                            { "rule", "log" },
                            { "autoclean", "true" },
                        },
                    },
                },
                Status = "Enabled",
                Transitions = new[]
                {
                    new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleTransitionArgs
                    {
                        Days = 30,
                        StorageClass = "STANDARD_IA",
                    },
                    new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleTransitionArgs
                    {
                        Days = 60,
                        StorageClass = "GLACIER",
                    },
                },
            },
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "tmp",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    Prefix = "tmp/",
                },
                Expiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleExpirationArgs
                {
                    Date = "2023-01-13T00:00:00Z",
                },
                Status = "Enabled",
            },
        },
    });

    var versioningBucket = new Aws.S3.BucketV2("versioningBucket");

    var versioningBucketAcl = new Aws.S3.BucketAclV2("versioningBucketAcl", new()
    {
        Bucket = versioningBucket.Id,
        Acl = "private",
    });

    var versioning = new Aws.S3.BucketVersioningV2("versioning", new()
    {
        Bucket = versioningBucket.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningV2VersioningConfigurationArgs
        {
            Status = "Enabled",
        },
    });

    var versioning_bucket_config = new Aws.S3.BucketLifecycleConfigurationV2("versioning-bucket-config", new()
    {
        Bucket = versioningBucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleArgs
            {
                Id = "config",
                Filter = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleFilterArgs
                {
                    Prefix = "config/",
                },
                NoncurrentVersionExpiration = new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs
                {
                    NoncurrentDays = 90,
                },
                NoncurrentVersionTransitions = new[]
                {
                    new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs
                    {
                        NoncurrentDays = 30,
                        StorageClass = "STANDARD_IA",
                    },
                    new Aws.S3.Inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs
                    {
                        NoncurrentDays = 60,
                        StorageClass = "GLACIER",
                    },
                },
                Status = "Enabled",
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            versioning,
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
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "bucketAcl", &s3.BucketAclV2Args{
			Bucket: bucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketLifecycleConfigurationV2(ctx, "bucket-config", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: bucket.ID(),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("log"),
					Expiration: &s3.BucketLifecycleConfigurationV2RuleExpirationArgs{
						Days: pulumi.Int(90),
					},
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						And: &s3.BucketLifecycleConfigurationV2RuleFilterAndArgs{
							Prefix: pulumi.String("log/"),
							Tags: pulumi.StringMap{
								"rule":      pulumi.String("log"),
								"autoclean": pulumi.String("true"),
							},
						},
					},
					Status: pulumi.String("Enabled"),
					Transitions: s3.BucketLifecycleConfigurationV2RuleTransitionArray{
						&s3.BucketLifecycleConfigurationV2RuleTransitionArgs{
							Days:         pulumi.Int(30),
							StorageClass: pulumi.String("STANDARD_IA"),
						},
						&s3.BucketLifecycleConfigurationV2RuleTransitionArgs{
							Days:         pulumi.Int(60),
							StorageClass: pulumi.String("GLACIER"),
						},
					},
				},
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("tmp"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						Prefix: pulumi.String("tmp/"),
					},
					Expiration: &s3.BucketLifecycleConfigurationV2RuleExpirationArgs{
						Date: pulumi.String("2023-01-13T00:00:00Z"),
					},
					Status: pulumi.String("Enabled"),
				},
			},
		})
		if err != nil {
			return err
		}
		versioningBucket, err := s3.NewBucketV2(ctx, "versioningBucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "versioningBucketAcl", &s3.BucketAclV2Args{
			Bucket: versioningBucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		versioning, err := s3.NewBucketVersioningV2(ctx, "versioning", &s3.BucketVersioningV2Args{
			Bucket: versioningBucket.ID(),
			VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketLifecycleConfigurationV2(ctx, "versioning-bucket-config", &s3.BucketLifecycleConfigurationV2Args{
			Bucket: versioningBucket.ID(),
			Rules: s3.BucketLifecycleConfigurationV2RuleArray{
				&s3.BucketLifecycleConfigurationV2RuleArgs{
					Id: pulumi.String("config"),
					Filter: &s3.BucketLifecycleConfigurationV2RuleFilterArgs{
						Prefix: pulumi.String("config/"),
					},
					NoncurrentVersionExpiration: &s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs{
						NoncurrentDays: pulumi.Int(90),
					},
					NoncurrentVersionTransitions: s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArray{
						&s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs{
							NoncurrentDays: pulumi.Int(30),
							StorageClass:   pulumi.String("STANDARD_IA"),
						},
						&s3.BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs{
							NoncurrentDays: pulumi.Int(60),
							StorageClass:   pulumi.String("GLACIER"),
						},
					},
					Status: pulumi.String("Enabled"),
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			versioning,
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
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2;
import com.pulumi.aws.s3.BucketLifecycleConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleExpirationArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleFilterAndArgs;
import com.pulumi.aws.s3.BucketVersioningV2;
import com.pulumi.aws.s3.BucketVersioningV2Args;
import com.pulumi.aws.s3.inputs.BucketVersioningV2VersioningConfigurationArgs;
import com.pulumi.aws.s3.inputs.BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs;
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
        var bucket = new BucketV2("bucket");

        var bucketAcl = new BucketAclV2("bucketAcl", BucketAclV2Args.builder()        
            .bucket(bucket.id())
            .acl("private")
            .build());

        var bucket_config = new BucketLifecycleConfigurationV2("bucket-config", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(bucket.id())
            .rules(            
                BucketLifecycleConfigurationV2RuleArgs.builder()
                    .id("log")
                    .expiration(BucketLifecycleConfigurationV2RuleExpirationArgs.builder()
                        .days(90)
                        .build())
                    .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                        .and(BucketLifecycleConfigurationV2RuleFilterAndArgs.builder()
                            .prefix("log/")
                            .tags(Map.ofEntries(
                                Map.entry("rule", "log"),
                                Map.entry("autoclean", "true")
                            ))
                            .build())
                        .build())
                    .status("Enabled")
                    .transitions(                    
                        BucketLifecycleConfigurationV2RuleTransitionArgs.builder()
                            .days(30)
                            .storageClass("STANDARD_IA")
                            .build(),
                        BucketLifecycleConfigurationV2RuleTransitionArgs.builder()
                            .days(60)
                            .storageClass("GLACIER")
                            .build())
                    .build(),
                BucketLifecycleConfigurationV2RuleArgs.builder()
                    .id("tmp")
                    .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                        .prefix("tmp/")
                        .build())
                    .expiration(BucketLifecycleConfigurationV2RuleExpirationArgs.builder()
                        .date("2023-01-13T00:00:00Z")
                        .build())
                    .status("Enabled")
                    .build())
            .build());

        var versioningBucket = new BucketV2("versioningBucket");

        var versioningBucketAcl = new BucketAclV2("versioningBucketAcl", BucketAclV2Args.builder()        
            .bucket(versioningBucket.id())
            .acl("private")
            .build());

        var versioning = new BucketVersioningV2("versioning", BucketVersioningV2Args.builder()        
            .bucket(versioningBucket.id())
            .versioningConfiguration(BucketVersioningV2VersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build());

        var versioning_bucket_config = new BucketLifecycleConfigurationV2("versioning-bucket-config", BucketLifecycleConfigurationV2Args.builder()        
            .bucket(versioningBucket.id())
            .rules(BucketLifecycleConfigurationV2RuleArgs.builder()
                .id("config")
                .filter(BucketLifecycleConfigurationV2RuleFilterArgs.builder()
                    .prefix("config/")
                    .build())
                .noncurrentVersionExpiration(BucketLifecycleConfigurationV2RuleNoncurrentVersionExpirationArgs.builder()
                    .noncurrentDays(90)
                    .build())
                .noncurrentVersionTransitions(                
                    BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs.builder()
                        .noncurrentDays(30)
                        .storageClass("STANDARD_IA")
                        .build(),
                    BucketLifecycleConfigurationV2RuleNoncurrentVersionTransitionArgs.builder()
                        .noncurrentDays(60)
                        .storageClass("GLACIER")
                        .build())
                .status("Enabled")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(versioning)
                .build());

    }
}
```
```yaml
resources:
  bucket:
    type: aws:s3:BucketV2
  bucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${bucket.id}
      acl: private
  bucket-config:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${bucket.id}
      rules:
        - id: log
          expiration:
            days: 90
          filter:
            and:
              prefix: log/
              tags:
                rule: log
                autoclean: 'true'
          status: Enabled
          transitions:
            - days: 30
              storageClass: STANDARD_IA
            - days: 60
              storageClass: GLACIER
        - id: tmp
          filter:
            prefix: tmp/
          expiration:
            date: 2023-01-13T00:00:00Z
          status: Enabled
  versioningBucket:
    type: aws:s3:BucketV2
  versioningBucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${versioningBucket.id}
      acl: private
  versioning:
    type: aws:s3:BucketVersioningV2
    properties:
      bucket: ${versioningBucket.id}
      versioningConfiguration:
        status: Enabled
  versioning-bucket-config:
    type: aws:s3:BucketLifecycleConfigurationV2
    properties:
      bucket: ${versioningBucket.id}
      rules:
        - id: config
          filter:
            prefix: config/
          noncurrentVersionExpiration:
            noncurrentDays: 90
          noncurrentVersionTransitions:
            - noncurrentDays: 30
              storageClass: STANDARD_IA
            - noncurrentDays: 60
              storageClass: GLACIER
          status: Enabled
    options:
      dependson:
        - ${versioning}
```
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

__Using `pulumi import` to import__ S3 bucket lifecycle configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:

If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2 example bucket-name
```
 If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketLifecycleConfigurationV2:BucketLifecycleConfigurationV2 example bucket-name,123456789012
```
 