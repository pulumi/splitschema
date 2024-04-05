Provides an Elastic Beanstalk Environment Resource. Elastic Beanstalk allows
you to deploy and manage applications in the AWS cloud without worrying about
the infrastructure that runs those applications.

Environments are often things such as `development`, `integration`, or
`production`.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {description: "tf-test-desc"});
const tfenvtest = new aws.elasticbeanstalk.Environment("tfenvtest", {
    application: tftest.name,
    solutionStackName: "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
});
```
```python
import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
tfenvtest = aws.elasticbeanstalk.Environment("tfenvtest",
    application=tftest.name,
    solution_stack_name="64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4")
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

    var tfenvtest = new Aws.ElasticBeanstalk.Environment("tfenvtest", new()
    {
        Application = tftest.Name,
        SolutionStackName = "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
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
		_, err = elasticbeanstalk.NewEnvironment(ctx, "tfenvtest", &elasticbeanstalk.EnvironmentArgs{
			Application:       tftest.Name,
			SolutionStackName: pulumi.String("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4"),
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
import com.pulumi.aws.elasticbeanstalk.Environment;
import com.pulumi.aws.elasticbeanstalk.EnvironmentArgs;
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

        var tfenvtest = new Environment("tfenvtest", EnvironmentArgs.builder()        
            .application(tftest.name())
            .solutionStackName("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4")
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
  tfenvtest:
    type: aws:elasticbeanstalk:Environment
    properties:
      application: ${tftest.name}
      solutionStackName: 64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4
```
{{% /example %}}
{{% /examples %}}
## Option Settings

Some options can be stack-specific, check [AWS Docs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html)
for supported options and examples.

The `setting` and `all_settings` mappings support the following format:

* `namespace` - unique namespace identifying the option's associated AWS resource
* `name` - name of the configuration option
* `value` - value for the configuration option
* `resource` - (Optional) resource name for [scheduled action](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options-general.html#command-options-general-autoscalingscheduledaction)

### Example With Options

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tftest = new aws.elasticbeanstalk.Application("tftest", {description: "tf-test-desc"});
const tfenvtest = new aws.elasticbeanstalk.Environment("tfenvtest", {
    application: tftest.name,
    solutionStackName: "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
    settings: [
        {
            namespace: "aws:ec2:vpc",
            name: "VPCId",
            value: "vpc-xxxxxxxx",
        },
        {
            namespace: "aws:ec2:vpc",
            name: "Subnets",
            value: "subnet-xxxxxxxx",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

tftest = aws.elasticbeanstalk.Application("tftest", description="tf-test-desc")
tfenvtest = aws.elasticbeanstalk.Environment("tfenvtest",
    application=tftest.name,
    solution_stack_name="64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
    settings=[
        aws.elasticbeanstalk.EnvironmentSettingArgs(
            namespace="aws:ec2:vpc",
            name="VPCId",
            value="vpc-xxxxxxxx",
        ),
        aws.elasticbeanstalk.EnvironmentSettingArgs(
            namespace="aws:ec2:vpc",
            name="Subnets",
            value="subnet-xxxxxxxx",
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
    var tftest = new Aws.ElasticBeanstalk.Application("tftest", new()
    {
        Description = "tf-test-desc",
    });

    var tfenvtest = new Aws.ElasticBeanstalk.Environment("tfenvtest", new()
    {
        Application = tftest.Name,
        SolutionStackName = "64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4",
        Settings = new[]
        {
            new Aws.ElasticBeanstalk.Inputs.EnvironmentSettingArgs
            {
                Namespace = "aws:ec2:vpc",
                Name = "VPCId",
                Value = "vpc-xxxxxxxx",
            },
            new Aws.ElasticBeanstalk.Inputs.EnvironmentSettingArgs
            {
                Namespace = "aws:ec2:vpc",
                Name = "Subnets",
                Value = "subnet-xxxxxxxx",
            },
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
		tftest, err := elasticbeanstalk.NewApplication(ctx, "tftest", &elasticbeanstalk.ApplicationArgs{
			Description: pulumi.String("tf-test-desc"),
		})
		if err != nil {
			return err
		}
		_, err = elasticbeanstalk.NewEnvironment(ctx, "tfenvtest", &elasticbeanstalk.EnvironmentArgs{
			Application:       tftest.Name,
			SolutionStackName: pulumi.String("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4"),
			Settings: elasticbeanstalk.EnvironmentSettingArray{
				&elasticbeanstalk.EnvironmentSettingArgs{
					Namespace: pulumi.String("aws:ec2:vpc"),
					Name:      pulumi.String("VPCId"),
					Value:     pulumi.String("vpc-xxxxxxxx"),
				},
				&elasticbeanstalk.EnvironmentSettingArgs{
					Namespace: pulumi.String("aws:ec2:vpc"),
					Name:      pulumi.String("Subnets"),
					Value:     pulumi.String("subnet-xxxxxxxx"),
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
import com.pulumi.aws.elasticbeanstalk.Application;
import com.pulumi.aws.elasticbeanstalk.ApplicationArgs;
import com.pulumi.aws.elasticbeanstalk.Environment;
import com.pulumi.aws.elasticbeanstalk.EnvironmentArgs;
import com.pulumi.aws.elasticbeanstalk.inputs.EnvironmentSettingArgs;
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

        var tfenvtest = new Environment("tfenvtest", EnvironmentArgs.builder()        
            .application(tftest.name())
            .solutionStackName("64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4")
            .settings(            
                EnvironmentSettingArgs.builder()
                    .namespace("aws:ec2:vpc")
                    .name("VPCId")
                    .value("vpc-xxxxxxxx")
                    .build(),
                EnvironmentSettingArgs.builder()
                    .namespace("aws:ec2:vpc")
                    .name("Subnets")
                    .value("subnet-xxxxxxxx")
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
  tfenvtest:
    type: aws:elasticbeanstalk:Environment
    properties:
      application: ${tftest.name}
      solutionStackName: 64bit Amazon Linux 2015.03 v2.0.3 running Go 1.4
      settings:
        - namespace: aws:ec2:vpc
          name: VPCId
          value: vpc-xxxxxxxx
        - namespace: aws:ec2:vpc
          name: Subnets
          value: subnet-xxxxxxxx
```


## Import

Using `pulumi import`, import Elastic Beanstalk Environments using the `id`. For example:

```sh
 $ pulumi import aws:elasticbeanstalk/environment:Environment prodenv e-rpqsewtp2j
```
 