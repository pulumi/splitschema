Provides an Elastic Beanstalk Application Resource. Elastic Beanstalk allows
you to deploy and manage applications in the AWS cloud without worrying about
the infrastructure that runs those applications.

This resource creates an application that has one configuration template named
`default`, and no application versions

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {
    description: "tf-test-desc",
    appversionLifecycle: {
        serviceRole: aws_iam_role.beanstalk_service.arn,
        maxCount: 128,
        deleteSourceFromS3: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest",
    description="tf-test-desc",
    appversion_lifecycle=aws.elasticbeanstalk.ApplicationAppversionLifecycleArgs(
        service_role=aws_iam_role["beanstalk_service"]["arn"],
        max_count=128,
        delete_source_from_s3=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var tftest = new Aws.ElasticBeanstalk.Application("tftest", new()
    {
        Description = "tf-test-desc",
        AppversionLifecycle = new Aws.ElasticBeanstalk.Inputs.ApplicationAppversionLifecycleArgs
        {
            ServiceRole = aws_iam_role.Beanstalk_service.Arn,
            MaxCount = 128,
            DeleteSourceFromS3 = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
			Description: pulumi.String("tf-test-desc"),
			AppversionLifecycle: &elasticbeanstalk.ApplicationAppversionLifecycleArgs{
				ServiceRole:        pulumi.Any(aws_iam_role.Beanstalk_service.Arn),
				MaxCount:           pulumi.Int(128),
				DeleteSourceFromS3: pulumi.Bool(true),
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
import com.pulumi.aws.elasticbeanstalk.Application;
import com.pulumi.aws.elasticbeanstalk.ApplicationArgs;
import com.pulumi.aws.elasticbeanstalk.inputs.ApplicationAppversionLifecycleArgs;
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
        var tftest = new Application("tftest", ApplicationArgs.builder()        
            .description("tf-test-desc")
            .appversionLifecycle(ApplicationAppversionLifecycleArgs.builder()
                .serviceRole(aws_iam_role.beanstalk_service().arn())
                .maxCount(128)
                .deleteSourceFromS3(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  tftest:
    type: aws:elasticbeanstalk:Application
    properties:
      description: tf-test-desc
      appversionLifecycle:
        serviceRole: ${aws_iam_role.beanstalk_service.arn}
        maxCount: 128
        deleteSourceFromS3: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Elastic Beanstalk Applications using the `name`. For example:

```sh
 $ pulumi import aws:elasticbeanstalk/application:Application tf_test tf-test-name
```
 