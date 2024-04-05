> **NOTE:** The `aws.s3.BucketObject` data source is DEPRECATED and will be removed in a future version! Use `aws.s3.BucketObjectv2` instead, where new features and fixes will be added.

The S3 object data source allows access to the metadata and
_optionally_ (see below) content of an object stored inside S3 bucket.

> **Note:** The content of an object (`body` field) is available only for objects which have a human-readable `Content-Type` (`text/*` and `application/json`). This is to prevent printing unsafe characters and potentially downloading large amount of data which would be thrown away in favour of metadata.

{{% examples %}}
## Example Usage
{{% example %}}

The following example retrieves a text object (which must have a `Content-Type`
value starting with `text/`) and uses it as the `user_data` for an EC2 instance:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bootstrapScript = aws.s3.getBucketObject({
    bucket: "ourcorp-deploy-config",
    key: "ec2-bootstrap-script.sh",
});
const example = new aws.ec2.Instance("example", {
    instanceType: "t2.micro",
    ami: "ami-2757f631",
    userData: bootstrapScript.then(bootstrapScript => bootstrapScript.body),
});
```
```python
import pulumi
import pulumi_aws as aws

bootstrap_script = aws.s3.get_bucket_object(bucket="ourcorp-deploy-config",
    key="ec2-bootstrap-script.sh")
example = aws.ec2.Instance("example",
    instance_type="t2.micro",
    ami="ami-2757f631",
    user_data=bootstrap_script.body)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bootstrapScript = Aws.S3.GetBucketObject.Invoke(new()
    {
        Bucket = "ourcorp-deploy-config",
        Key = "ec2-bootstrap-script.sh",
    });

    var example = new Aws.Ec2.Instance("example", new()
    {
        InstanceType = "t2.micro",
        Ami = "ami-2757f631",
        UserData = bootstrapScript.Apply(getBucketObjectResult => getBucketObjectResult.Body),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bootstrapScript, err := s3.LookupBucketObject(ctx, &s3.LookupBucketObjectArgs{
			Bucket: "ourcorp-deploy-config",
			Key:    "ec2-bootstrap-script.sh",
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "example", &ec2.InstanceArgs{
			InstanceType: pulumi.String("t2.micro"),
			Ami:          pulumi.String("ami-2757f631"),
			UserData:     *pulumi.String(bootstrapScript.Body),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketObjectArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
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
        final var bootstrapScript = S3Functions.getBucketObject(GetBucketObjectArgs.builder()
            .bucket("ourcorp-deploy-config")
            .key("ec2-bootstrap-script.sh")
            .build());

        var example = new Instance("example", InstanceArgs.builder()        
            .instanceType("t2.micro")
            .ami("ami-2757f631")
            .userData(bootstrapScript.applyValue(getBucketObjectResult -> getBucketObjectResult.body()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:Instance
    properties:
      instanceType: t2.micro
      ami: ami-2757f631
      userData: ${bootstrapScript.body}
variables:
  bootstrapScript:
    fn::invoke:
      Function: aws:s3:getBucketObject
      Arguments:
        bucket: ourcorp-deploy-config
        key: ec2-bootstrap-script.sh
```

The following, more-complex example retrieves only the metadata for a zip
file stored in S3, which is then used to pass the most recent `version_id`
to AWS Lambda for use as a function implementation. More information about
Lambda functions is available in the documentation for
`aws.lambda.Function`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lambda = aws.s3.getBucketObject({
    bucket: "ourcorp-lambda-functions",
    key: "hello-world.zip",
});
const testLambda = new aws.lambda.Function("testLambda", {
    s3Bucket: lambda.then(lambda => lambda.id),
    s3Key: lambda.then(lambda => lambda.key),
    s3ObjectVersion: lambda.then(lambda => lambda.versionId),
    role: aws_iam_role.iam_for_lambda.arn,
    handler: "exports.test",
});
```
```python
import pulumi
import pulumi_aws as aws

lambda_ = aws.s3.get_bucket_object(bucket="ourcorp-lambda-functions",
    key="hello-world.zip")
test_lambda = aws.lambda_.Function("testLambda",
    s3_bucket=lambda_.id,
    s3_key=lambda_.key,
    s3_object_version=lambda_.version_id,
    role=aws_iam_role["iam_for_lambda"]["arn"],
    handler="exports.test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lambda = Aws.S3.GetBucketObject.Invoke(new()
    {
        Bucket = "ourcorp-lambda-functions",
        Key = "hello-world.zip",
    });

    var testLambda = new Aws.Lambda.Function("testLambda", new()
    {
        S3Bucket = lambda.Apply(getBucketObjectResult => getBucketObjectResult.Id),
        S3Key = lambda.Apply(getBucketObjectResult => getBucketObjectResult.Key),
        S3ObjectVersion = lambda.Apply(getBucketObjectResult => getBucketObjectResult.VersionId),
        Role = aws_iam_role.Iam_for_lambda.Arn,
        Handler = "exports.test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		lambda, err := s3.LookupBucketObject(ctx, &s3.LookupBucketObjectArgs{
			Bucket: "ourcorp-lambda-functions",
			Key:    "hello-world.zip",
		}, nil)
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "testLambda", &lambda.FunctionArgs{
			S3Bucket:        *pulumi.String(lambda.Id),
			S3Key:           *pulumi.String(lambda.Key),
			S3ObjectVersion: *pulumi.String(lambda.VersionId),
			Role:            pulumi.Any(aws_iam_role.Iam_for_lambda.Arn),
			Handler:         pulumi.String("exports.test"),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketObjectArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
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
        final var lambda = S3Functions.getBucketObject(GetBucketObjectArgs.builder()
            .bucket("ourcorp-lambda-functions")
            .key("hello-world.zip")
            .build());

        var testLambda = new Function("testLambda", FunctionArgs.builder()        
            .s3Bucket(lambda.applyValue(getBucketObjectResult -> getBucketObjectResult.id()))
            .s3Key(lambda.applyValue(getBucketObjectResult -> getBucketObjectResult.key()))
            .s3ObjectVersion(lambda.applyValue(getBucketObjectResult -> getBucketObjectResult.versionId()))
            .role(aws_iam_role.iam_for_lambda().arn())
            .handler("exports.test")
            .build());

    }
}
```
```yaml
resources:
  testLambda:
    type: aws:lambda:Function
    properties:
      s3Bucket: ${lambda.id}
      s3Key: ${lambda.key}
      s3ObjectVersion: ${lambda.versionId}
      role: ${aws_iam_role.iam_for_lambda.arn}
      # (not shown)
      handler: exports.test
variables:
  lambda:
    fn::invoke:
      Function: aws:s3:getBucketObject
      Arguments:
        bucket: ourcorp-lambda-functions
        key: hello-world.zip
```
{{% /example %}}
{{% /examples %}}