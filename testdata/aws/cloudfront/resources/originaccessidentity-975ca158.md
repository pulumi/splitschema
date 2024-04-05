Creates an Amazon CloudFront origin access identity.

For information about CloudFront distributions, see the
[Amazon CloudFront Developer Guide](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html). For more information on generating
origin access identities, see
[Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content][2].

{{% examples %}}
## Example Usage
{{% example %}}

The following example below creates a CloudFront origin access identity.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.OriginAccessIdentity("example", {comment: "Some comment"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.OriginAccessIdentity("example", comment="Some comment")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.OriginAccessIdentity("example", new()
    {
        Comment = "Some comment",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewOriginAccessIdentity(ctx, "example", &cloudfront.OriginAccessIdentityArgs{
			Comment: pulumi.String("Some comment"),
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
import com.pulumi.aws.cloudfront.OriginAccessIdentity;
import com.pulumi.aws.cloudfront.OriginAccessIdentityArgs;
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
        var example = new OriginAccessIdentity("example", OriginAccessIdentityArgs.builder()        
            .comment("Some comment")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:OriginAccessIdentity
    properties:
      comment: Some comment
```
{{% /example %}}
{{% /examples %}}
## Using With CloudFront

Normally, when referencing an origin access identity in CloudFront, you need to
prefix the ID with the `origin-access-identity/cloudfront/` special path.
The `cloudfront_access_identity_path` allows this to be circumvented.
The below snippet demonstrates use with the `s3_origin_config` structure for the
`aws.cloudfront.Distribution` resource:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const example = new aws.cloudfront.Distribution("example", {origins: [{
    s3OriginConfig: {
        originAccessIdentity: aws_cloudfront_origin_access_identity.example.cloudfront_access_identity_path,
    },
}]});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example = aws.cloudfront.Distribution("example", origins=[aws.cloudfront.DistributionOriginArgs(
    s3_origin_config=aws.cloudfront.DistributionOriginS3OriginConfigArgs(
        origin_access_identity=aws_cloudfront_origin_access_identity["example"]["cloudfront_access_identity_path"],
    ),
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configuration ...
    var example = new Aws.CloudFront.Distribution("example", new()
    {
        Origins = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOriginArgs
            {
                S3OriginConfig = new Aws.CloudFront.Inputs.DistributionOriginS3OriginConfigArgs
                {
                    OriginAccessIdentity = aws_cloudfront_origin_access_identity.Example.Cloudfront_access_identity_path,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewDistribution(ctx, "example", &cloudfront.DistributionArgs{
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					S3OriginConfig: &cloudfront.DistributionOriginS3OriginConfigArgs{
						OriginAccessIdentity: pulumi.Any(aws_cloudfront_origin_access_identity.Example.Cloudfront_access_identity_path),
					},
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
import com.pulumi.aws.cloudfront.Distribution;
import com.pulumi.aws.cloudfront.DistributionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginS3OriginConfigArgs;
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
        var example = new Distribution("example", DistributionArgs.builder()        
            .origins(DistributionOriginArgs.builder()
                .s3OriginConfig(DistributionOriginS3OriginConfigArgs.builder()
                    .originAccessIdentity(aws_cloudfront_origin_access_identity.example().cloudfront_access_identity_path())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:Distribution
    properties:
      origins:
        - s3OriginConfig:
            originAccessIdentity: ${aws_cloudfront_origin_access_identity.example.cloudfront_access_identity_path}
```

### Updating your bucket policy

Note that the AWS API may translate the `s3_canonical_user_id` `CanonicalUser`
principal into an `AWS` IAM ARN principal when supplied in an
`aws.s3.BucketV2` bucket policy, causing spurious diffs. If
you see this behaviour, use the `iam_arn` instead:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3Policy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["s3:GetObject"],
        resources: [`${aws_s3_bucket.example.arn}/*`],
        principals: [{
            type: "AWS",
            identifiers: [aws_cloudfront_origin_access_identity.example.iam_arn],
        }],
    }],
});
const example = new aws.s3.BucketPolicy("example", {
    bucket: aws_s3_bucket.example.id,
    policy: s3Policy.then(s3Policy => s3Policy.json),
});
```
```python
import pulumi
import pulumi_aws as aws

s3_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["s3:GetObject"],
    resources=[f"{aws_s3_bucket['example']['arn']}/*"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[aws_cloudfront_origin_access_identity["example"]["iam_arn"]],
    )],
)])
example = aws.s3.BucketPolicy("example",
    bucket=aws_s3_bucket["example"]["id"],
    policy=s3_policy.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3Policy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "s3:GetObject",
                },
                Resources = new[]
                {
                    $"{aws_s3_bucket.Example.Arn}/*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            aws_cloudfront_origin_access_identity.Example.Iam_arn,
                        },
                    },
                },
            },
        },
    });

    var example = new Aws.S3.BucketPolicy("example", new()
    {
        Bucket = aws_s3_bucket.Example.Id,
        Policy = s3Policy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
s3Policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Actions: []string{
"s3:GetObject",
},
Resources: []string{
fmt.Sprintf("%v/*", aws_s3_bucket.Example.Arn),
},
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
aws_cloudfront_origin_access_identity.Example.Iam_arn,
},
},
},
},
},
}, nil);
if err != nil {
return err
}
_, err = s3.NewBucketPolicy(ctx, "example", &s3.BucketPolicyArgs{
Bucket: pulumi.Any(aws_s3_bucket.Example.Id),
Policy: *pulumi.String(s3Policy.Json),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
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
        final var s3Policy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("s3:GetObject")
                .resources(String.format("%s/*", aws_s3_bucket.example().arn()))
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(aws_cloudfront_origin_access_identity.example().iam_arn())
                    .build())
                .build())
            .build());

        var example = new BucketPolicy("example", BucketPolicyArgs.builder()        
            .bucket(aws_s3_bucket.example().id())
            .policy(s3Policy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${aws_s3_bucket.example.id}
      policy: ${s3Policy.json}
variables:
  s3Policy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - s3:GetObject
            resources:
              - ${aws_s3_bucket.example.arn}/*
            principals:
              - type: AWS
                identifiers:
                  - ${aws_cloudfront_origin_access_identity.example.iam_arn}
```

[1]: http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html
[2]: http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html


## Import

Using `pulumi import`, import Cloudfront Origin Access Identities using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/originAccessIdentity:OriginAccessIdentity origin_access E74FTE3AEXAMPLE
```
 