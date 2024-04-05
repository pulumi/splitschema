Provides an Elastic Beanstalk Application Version Resource. Elastic Beanstalk allows
you to deploy and manage applications in the AWS cloud without worrying about
the infrastructure that runs those applications.

This resource creates a Beanstalk Application Version that can be deployed to a Beanstalk
Environment.

> **NOTE on Application Version Resource:**  When using the Application Version resource with multiple
Elastic Beanstalk Environments it is possible that an error may be returned
when attempting to delete an Application Version while it is still in use by a different environment.
To work around this you can either create each environment in a separate AWS account or create your `aws.elasticbeanstalk.ApplicationVersion` resources with a unique names in your Elastic Beanstalk Application. For example &lt;revision&gt;-&lt;environment&gt;.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultBucketV2 = new aws.s3.BucketV2("defaultBucketV2", {});
const defaultBucketObjectv2 = new aws.s3.BucketObjectv2("defaultBucketObjectv2", {
    bucket: defaultBucketV2.id,
    key: "beanstalk/go-v1.zip",
    source: new pulumi.asset.FileAsset("go-v1.zip"),
});
const defaultApplication = new aws.elasticbeanstalk.Application("defaultApplication", {description: "tf-test-desc"});
const defaultApplicationVersion = new aws.elasticbeanstalk.ApplicationVersion("defaultApplicationVersion", {
    application: "tf-test-name",
    description: "application version",
    bucket: defaultBucketV2.id,
    key: defaultBucketObjectv2.id,
});
```
```python
import pulumi
import pulumi_aws as aws

default_bucket_v2 = aws.s3.BucketV2("defaultBucketV2")
default_bucket_objectv2 = aws.s3.BucketObjectv2("defaultBucketObjectv2",
    bucket=default_bucket_v2.id,
    key="beanstalk/go-v1.zip",
    source=pulumi.FileAsset("go-v1.zip"))
default_application = aws.elasticbeanstalk.Application("defaultApplication", description="tf-test-desc")
default_application_version = aws.elasticbeanstalk.ApplicationVersion("defaultApplicationVersion",
    application="tf-test-name",
    description="application version",
    bucket=default_bucket_v2.id,
    key=default_bucket_objectv2.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultBucketV2 = new Aws.S3.BucketV2("defaultBucketV2");

    var defaultBucketObjectv2 = new Aws.S3.BucketObjectv2("defaultBucketObjectv2", new()
    {
        Bucket = defaultBucketV2.Id,
        Key = "beanstalk/go-v1.zip",
        Source = new FileAsset("go-v1.zip"),
    });

    var defaultApplication = new Aws.ElasticBeanstalk.Application("defaultApplication", new()
    {
        Description = "tf-test-desc",
    });

    var defaultApplicationVersion = new Aws.ElasticBeanstalk.ApplicationVersion("defaultApplicationVersion", new()
    {
        Application = "tf-test-name",
        Description = "application version",
        Bucket = defaultBucketV2.Id,
        Key = defaultBucketObjectv2.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultBucketV2, err := s3.NewBucketV2(ctx, "defaultBucketV2", nil)
		if err != nil {
			return err
		}
		defaultBucketObjectv2, err := s3.NewBucketObjectv2(ctx, "defaultBucketObjectv2", &s3.BucketObjectv2Args{
			Bucket: defaultBucketV2.ID(),
			Key:    pulumi.String("beanstalk/go-v1.zip"),
			Source: pulumi.NewFileAsset("go-v1.zip"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewApplication(ctx, "defaultApplication", &elasticbeanstalk.ApplicationArgs{
			Description: pulumi.String("tf-test-desc"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewApplicationVersion(ctx, "defaultApplicationVersion", &elasticbeanstalk.ApplicationVersionArgs{
			Application: pulumi.Any("tf-test-name"),
			Description: pulumi.String("application version"),
			Bucket:      defaultBucketV2.ID(),
			Key:         defaultBucketObjectv2.ID(),
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
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.elasticbeanstalk.Application;
import com.pulumi.aws.elasticbeanstalk.ApplicationArgs;
import com.pulumi.aws.elasticbeanstalk.ApplicationVersion;
import com.pulumi.aws.elasticbeanstalk.ApplicationVersionArgs;
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
        var defaultBucketV2 = new BucketV2("defaultBucketV2");

        var defaultBucketObjectv2 = new BucketObjectv2("defaultBucketObjectv2", BucketObjectv2Args.builder()        
            .bucket(defaultBucketV2.id())
            .key("beanstalk/go-v1.zip")
            .source(new FileAsset("go-v1.zip"))
            .build());

        var defaultApplication = new Application("defaultApplication", ApplicationArgs.builder()        
            .description("tf-test-desc")
            .build());

        var defaultApplicationVersion = new ApplicationVersion("defaultApplicationVersion", ApplicationVersionArgs.builder()        
            .application("tf-test-name")
            .description("application version")
            .bucket(defaultBucketV2.id())
            .key(defaultBucketObjectv2.id())
            .build());

    }
}
```
```yaml
resources:
  defaultBucketV2:
    type: aws:s3:BucketV2
  defaultBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${defaultBucketV2.id}
      key: beanstalk/go-v1.zip
      source:
        fn::FileAsset: go-v1.zip
  defaultApplication:
    type: aws:elasticbeanstalk:Application
    properties:
      description: tf-test-desc
  defaultApplicationVersion:
    type: aws:elasticbeanstalk:ApplicationVersion
    properties:
      application: tf-test-name
      description: application version
      bucket: ${defaultBucketV2.id}
      key: ${defaultBucketObjectv2.id}
```
{{% /example %}}
{{% /examples %}}