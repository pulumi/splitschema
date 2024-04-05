Provides an S3 bucket Object Lock configuration resource. For more information about Object Locking, go to [Using S3 Object Lock](https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock.html) in the Amazon S3 User Guide.

> **NOTE:** This resource **does not enable** Object Lock for **new** buckets. It configures a default retention period for objects placed in the specified bucket.
Thus, to **enable** Object Lock for a **new** bucket, see the Using object lock configuration section in  the `aws.s3.BucketV2` resource or the Object Lock configuration for a new bucket example below.
If you want to **enable** Object Lock for an **existing** bucket, contact AWS Support and see the Object Lock configuration for an existing bucket example below.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Object Lock configuration for a new bucket

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {objectLockEnabled: true});
const exampleBucketObjectLockConfigurationV2 = new aws.s3.BucketObjectLockConfigurationV2("exampleBucketObjectLockConfigurationV2", {
    bucket: exampleBucketV2.id,
    rule: {
        defaultRetention: {
            mode: "COMPLIANCE",
            days: 5,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", object_lock_enabled=True)
example_bucket_object_lock_configuration_v2 = aws.s3.BucketObjectLockConfigurationV2("exampleBucketObjectLockConfigurationV2",
    bucket=example_bucket_v2.id,
    rule=aws.s3.BucketObjectLockConfigurationV2RuleArgs(
        default_retention=aws.s3.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs(
            mode="COMPLIANCE",
            days=5,
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        ObjectLockEnabled = true,
    });

    var exampleBucketObjectLockConfigurationV2 = new Aws.S3.BucketObjectLockConfigurationV2("exampleBucketObjectLockConfigurationV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Rule = new Aws.S3.Inputs.BucketObjectLockConfigurationV2RuleArgs
        {
            DefaultRetention = new Aws.S3.Inputs.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs
            {
                Mode = "COMPLIANCE",
                Days = 5,
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
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			ObjectLockEnabled: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObjectLockConfigurationV2(ctx, "exampleBucketObjectLockConfigurationV2", &s3.BucketObjectLockConfigurationV2Args{
			Bucket: exampleBucketV2.ID(),
			Rule: &s3.BucketObjectLockConfigurationV2RuleArgs{
				DefaultRetention: &s3.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs{
					Mode: pulumi.String("COMPLIANCE"),
					Days: pulumi.Int(5),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketObjectLockConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .objectLockEnabled(true)
            .build());

        var exampleBucketObjectLockConfigurationV2 = new BucketObjectLockConfigurationV2("exampleBucketObjectLockConfigurationV2", BucketObjectLockConfigurationV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .rule(BucketObjectLockConfigurationV2RuleArgs.builder()
                .defaultRetention(BucketObjectLockConfigurationV2RuleDefaultRetentionArgs.builder()
                    .mode("COMPLIANCE")
                    .days(5)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      objectLockEnabled: true
  exampleBucketObjectLockConfigurationV2:
    type: aws:s3:BucketObjectLockConfigurationV2
    properties:
      bucket: ${exampleBucketV2.id}
      rule:
        defaultRetention:
          mode: COMPLIANCE
          days: 5
```
{{% /example %}}
{{% example %}}
### Object Lock configuration for an existing bucket

This is a multistep process that requires AWS Support intervention.

1. Enable versioning on your S3 bucket, if you have not already done so.
Doing so will generate an "Object Lock token" in the back-end.

<!-- markdownlint-disable MD029 -->
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketVersioningV2 = new aws.s3.BucketVersioningV2("exampleBucketVersioningV2", {
    bucket: exampleBucketV2.id,
    versioningConfiguration: {
        status: "Enabled",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_versioning_v2 = aws.s3.BucketVersioningV2("exampleBucketVersioningV2",
    bucket=example_bucket_v2.id,
    versioning_configuration=aws.s3.BucketVersioningV2VersioningConfigurationArgs(
        status="Enabled",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketVersioningV2 = new Aws.S3.BucketVersioningV2("exampleBucketVersioningV2", new()
    {
        Bucket = exampleBucketV2.Id,
        VersioningConfiguration = new Aws.S3.Inputs.BucketVersioningV2VersioningConfigurationArgs
        {
            Status = "Enabled",
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
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketVersioningV2(ctx, "exampleBucketVersioningV2", &s3.BucketVersioningV2Args{
			Bucket: exampleBucketV2.ID(),
			VersioningConfiguration: &s3.BucketVersioningV2VersioningConfigurationArgs{
				Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketVersioningV2;
import com.pulumi.aws.s3.BucketVersioningV2Args;
import com.pulumi.aws.s3.inputs.BucketVersioningV2VersioningConfigurationArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketVersioningV2 = new BucketVersioningV2("exampleBucketVersioningV2", BucketVersioningV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .versioningConfiguration(BucketVersioningV2VersioningConfigurationArgs.builder()
                .status("Enabled")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketVersioningV2:
    type: aws:s3:BucketVersioningV2
    properties:
      bucket: ${exampleBucketV2.id}
      versioningConfiguration:
        status: Enabled
```
<!-- markdownlint-disable MD029 -->

2. Contact AWS Support to provide you with the "Object Lock token" for the specified bucket and use the token (or token ID) within your new `aws.s3.BucketObjectLockConfigurationV2` resource.
   Notice the `object_lock_enabled` argument does not need to be specified as it defaults to `Enabled`.

<!-- markdownlint-disable MD029 -->
```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketObjectLockConfigurationV2("example", {
    bucket: aws_s3_bucket.example.id,
    rule: {
        defaultRetention: {
            mode: "COMPLIANCE",
            days: 5,
        },
    },
    token: "NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w==",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketObjectLockConfigurationV2("example",
    bucket=aws_s3_bucket["example"]["id"],
    rule=aws.s3.BucketObjectLockConfigurationV2RuleArgs(
        default_retention=aws.s3.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs(
            mode="COMPLIANCE",
            days=5,
        ),
    ),
    token="NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w==")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketObjectLockConfigurationV2("example", new()
    {
        Bucket = aws_s3_bucket.Example.Id,
        Rule = new Aws.S3.Inputs.BucketObjectLockConfigurationV2RuleArgs
        {
            DefaultRetention = new Aws.S3.Inputs.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs
            {
                Mode = "COMPLIANCE",
                Days = 5,
            },
        },
        Token = "NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w==",
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
		_, err := s3.NewBucketObjectLockConfigurationV2(ctx, "example", &s3.BucketObjectLockConfigurationV2Args{
			Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
			Rule: &s3.BucketObjectLockConfigurationV2RuleArgs{
				DefaultRetention: &s3.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs{
					Mode: pulumi.String("COMPLIANCE"),
					Days: pulumi.Int(5),
				},
			},
			Token: pulumi.String("NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w=="),
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
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2;
import com.pulumi.aws.s3.BucketObjectLockConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketObjectLockConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketObjectLockConfigurationV2RuleDefaultRetentionArgs;
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
        var example = new BucketObjectLockConfigurationV2("example", BucketObjectLockConfigurationV2Args.builder()        
            .bucket(aws_s3_bucket.example().id())
            .rule(BucketObjectLockConfigurationV2RuleArgs.builder()
                .defaultRetention(BucketObjectLockConfigurationV2RuleDefaultRetentionArgs.builder()
                    .mode("COMPLIANCE")
                    .days(5)
                    .build())
                .build())
            .token("NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w==")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketObjectLockConfigurationV2
    properties:
      bucket: ${aws_s3_bucket.example.id}
      rule:
        defaultRetention:
          mode: COMPLIANCE
          days: 5
      token: NG2MKsfoLqV3A+aquXneSG4LOu/ekrlXkRXwIPFVfERT7XOPos+/k444d7RIH0E3W3p5QU6ml2exS2F/eYCFmMWHJ3hFZGk6al1sIJkmNhUMYmsv0jYVQyTTZNLM+DnfooA6SATt39mM1VW1yJh4E+XljMlWzaBwHKbss3/EjlGDjOmVhaSs4Z6427mMCaFD0RLwsYY7zX49gEc31YfOMJGxbXCXSeyNwAhhM/A8UH7gQf38RmjHjjAFbbbLtl8arsxTPW8F1IYohqwmKIr9DnotLLj8Tg44U2SPwujVaqmlKKP9s41rfgb4UbIm7khSafDBng0LGfxC4pMlT9Ny2w==
```
<!-- markdownlint-disable MD029 -->
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

import {

 to = aws_s3_bucket_object_lock_configuration.example

 id = "bucket-name,123456789012" }

__Using `pulumi import` to import__ S3 bucket Object Lock configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:

If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2 example bucket-name
```
 If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketObjectLockConfigurationV2:BucketObjectLockConfigurationV2 example bucket-name,123456789012
```
 