Provides an Elastic Beanstalk Configuration Template, which are associated with
a specific application and are used to deploy different versions of the
application with the same configuration settings.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {description: "tf-test-desc"});
const myTemplate = new aws.elasticbeanstalk.ConfigurationTemplate("myTemplate", {
    application: tftest.name,
    solutionStackName: "64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4",
});
```
```python
import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
my_template = aws.elasticbeanstalk.ConfigurationTemplate("myTemplate",
    application=tftest.name,
    solution_stack_name="64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4")
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
    });

    var myTemplate = new Aws.ElasticBeanstalk.ConfigurationTemplate("myTemplate", new()
    {
        Application = tftest.Name,
        SolutionStackName = "64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4",
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
		tftest, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
			Description: pulumi.String("tf-test-desc"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewConfigurationTemplate(ctx, "myTemplate", &elasticbeanstalk.ConfigurationTemplateArgs{
			Application:       tftest.Name,
			SolutionStackName: pulumi.String("64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4"),
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
import com.pulumi.aws.elasticbeanstalk.ConfigurationTemplate;
import com.pulumi.aws.elasticbeanstalk.ConfigurationTemplateArgs;
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
            .build());

        var myTemplate = new ConfigurationTemplate("myTemplate", ConfigurationTemplateArgs.builder()        
            .application(tftest.name())
            .solutionStackName("64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4")
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
  myTemplate:
    type: aws:elasticbeanstalk:ConfigurationTemplate
    properties:
      application: ${tftest.name}
      solutionStackName: 64bit Amazon Linux 2015.09 v2.0.8 running Go 1.4
```
{{% /example %}}
{{% /examples %}}
## Option Settings

The `setting` field supports the following format:

* `namespace` - unique namespace identifying the option's associated AWS resource
* `name` - name of the configuration option
* `value` - value for the configuration option
* `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)
