Associates an SSM Document to an instance or EC2 tag.

{{% examples %}}
## Example Usage
{{% example %}}
### Create an association for a specific instance

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.Association("example", {targets: [{
    key: "InstanceIds",
    values: [aws_instance.example.id],
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.Association("example", targets=[aws.ssm.AssociationTargetArgs(
    key="InstanceIds",
    values=[aws_instance["example"]["id"]],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ssm.Association("example", new()
    {
        Targets = new[]
        {
            new Aws.Ssm.Inputs.AssociationTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    aws_instance.Example.Id,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewAssociation(ctx, "example", &ssm.AssociationArgs{
			Targets: ssm.AssociationTargetArray{
				&ssm.AssociationTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						aws_instance.Example.Id,
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
import com.pulumi.aws.ssm.Association;
import com.pulumi.aws.ssm.AssociationArgs;
import com.pulumi.aws.ssm.inputs.AssociationTargetArgs;
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
        var example = new Association("example", AssociationArgs.builder()        
            .targets(AssociationTargetArgs.builder()
                .key("InstanceIds")
                .values(aws_instance.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:Association
    properties:
      targets:
        - key: InstanceIds
          values:
            - ${aws_instance.example.id}
```
{{% /example %}}
{{% example %}}
### Create an association for all managed instances in an AWS account

To target all managed instances in an AWS account, set the `key` as `"InstanceIds"` with `values` set as `["*"]`. This example also illustrates how to use an Amazon owned SSM document named `AmazonCloudWatch-ManageAgent`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.Association("example", {targets: [{
    key: "InstanceIds",
    values: ["*"],
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.Association("example", targets=[aws.ssm.AssociationTargetArgs(
    key="InstanceIds",
    values=["*"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ssm.Association("example", new()
    {
        Targets = new[]
        {
            new Aws.Ssm.Inputs.AssociationTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    "*",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewAssociation(ctx, "example", &ssm.AssociationArgs{
			Targets: ssm.AssociationTargetArray{
				&ssm.AssociationTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						pulumi.String("*"),
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
import com.pulumi.aws.ssm.Association;
import com.pulumi.aws.ssm.AssociationArgs;
import com.pulumi.aws.ssm.inputs.AssociationTargetArgs;
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
        var example = new Association("example", AssociationArgs.builder()        
            .targets(AssociationTargetArgs.builder()
                .key("InstanceIds")
                .values("*")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:Association
    properties:
      targets:
        - key: InstanceIds
          values:
            - '*'
```
{{% /example %}}
{{% example %}}
### Create an association for a specific tag

This example shows how to target all managed instances that are assigned a tag key of `Environment` and value of `Development`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.Association("example", {targets: [{
    key: "tag:Environment",
    values: ["Development"],
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.Association("example", targets=[aws.ssm.AssociationTargetArgs(
    key="tag:Environment",
    values=["Development"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ssm.Association("example", new()
    {
        Targets = new[]
        {
            new Aws.Ssm.Inputs.AssociationTargetArgs
            {
                Key = "tag:Environment",
                Values = new[]
                {
                    "Development",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewAssociation(ctx, "example", &ssm.AssociationArgs{
			Targets: ssm.AssociationTargetArray{
				&ssm.AssociationTargetArgs{
					Key: pulumi.String("tag:Environment"),
					Values: pulumi.StringArray{
						pulumi.String("Development"),
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
import com.pulumi.aws.ssm.Association;
import com.pulumi.aws.ssm.AssociationArgs;
import com.pulumi.aws.ssm.inputs.AssociationTargetArgs;
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
        var example = new Association("example", AssociationArgs.builder()        
            .targets(AssociationTargetArgs.builder()
                .key("tag:Environment")
                .values("Development")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:Association
    properties:
      targets:
        - key: tag:Environment
          values:
            - Development
```
{{% /example %}}
{{% example %}}
### Create an association with a specific schedule

This example shows how to schedule an association in various ways.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssm.Association("example", {
    scheduleExpression: "cron(0 2 ? * SUN *)",
    targets: [{
        key: "InstanceIds",
        values: [aws_instance.example.id],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssm.Association("example",
    schedule_expression="cron(0 2 ? * SUN *)",
    targets=[aws.ssm.AssociationTargetArgs(
        key="InstanceIds",
        values=[aws_instance["example"]["id"]],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ssm.Association("example", new()
    {
        ScheduleExpression = "cron(0 2 ? * SUN *)",
        Targets = new[]
        {
            new Aws.Ssm.Inputs.AssociationTargetArgs
            {
                Key = "InstanceIds",
                Values = new[]
                {
                    aws_instance.Example.Id,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.NewAssociation(ctx, "example", &ssm.AssociationArgs{
			ScheduleExpression: pulumi.String("cron(0 2 ? * SUN *)"),
			Targets: ssm.AssociationTargetArray{
				&ssm.AssociationTargetArgs{
					Key: pulumi.String("InstanceIds"),
					Values: pulumi.StringArray{
						aws_instance.Example.Id,
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
import com.pulumi.aws.ssm.Association;
import com.pulumi.aws.ssm.AssociationArgs;
import com.pulumi.aws.ssm.inputs.AssociationTargetArgs;
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
        var example = new Association("example", AssociationArgs.builder()        
            .scheduleExpression("cron(0 2 ? * SUN *)")
            .targets(AssociationTargetArgs.builder()
                .key("InstanceIds")
                .values(aws_instance.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssm:Association
    properties:
      # Cron expression example
      scheduleExpression: cron(0 2 ? * SUN *) # Rate expression example
      #   # schedule_expression = "rate(7 days)"
      targets:
        - key: InstanceIds
          values:
            - ${aws_instance.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM associations using the `association_id`. For example:

```sh
 $ pulumi import aws:ssm/association:Association test-association 10abcdef-0abc-1234-5678-90abcdef123456
```
 