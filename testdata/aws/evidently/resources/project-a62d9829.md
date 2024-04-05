Provides a CloudWatch Evidently Project resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Project("example", {
    description: "Example Description",
    tags: {
        Key1: "example Project",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Project("example",
    description="Example Description",
    tags={
        "Key1": "example Project",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Project("example", new()
    {
        Description = "Example Description",
        Tags = 
        {
            { "Key1", "example Project" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewProject(ctx, "example", &evidently.ProjectArgs{
			Description: pulumi.String("Example Description"),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("example Project"),
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
import com.pulumi.aws.evidently.Project;
import com.pulumi.aws.evidently.ProjectArgs;
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
        var example = new Project("example", ProjectArgs.builder()        
            .description("Example Description")
            .tags(Map.of("Key1", "example Project"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Project
    properties:
      description: Example Description
      tags:
        Key1: example Project
```
{{% /example %}}
{{% example %}}
### Store evaluation events in a CloudWatch Log Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Project("example", {
    dataDelivery: {
        cloudwatchLogs: {
            logGroup: "example-log-group-name",
        },
    },
    description: "Example Description",
    tags: {
        Key1: "example Project",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Project("example",
    data_delivery=aws.evidently.ProjectDataDeliveryArgs(
        cloudwatch_logs=aws.evidently.ProjectDataDeliveryCloudwatchLogsArgs(
            log_group="example-log-group-name",
        ),
    ),
    description="Example Description",
    tags={
        "Key1": "example Project",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Project("example", new()
    {
        DataDelivery = new Aws.Evidently.Inputs.ProjectDataDeliveryArgs
        {
            CloudwatchLogs = new Aws.Evidently.Inputs.ProjectDataDeliveryCloudwatchLogsArgs
            {
                LogGroup = "example-log-group-name",
            },
        },
        Description = "Example Description",
        Tags = 
        {
            { "Key1", "example Project" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewProject(ctx, "example", &evidently.ProjectArgs{
			DataDelivery: &evidently.ProjectDataDeliveryArgs{
				CloudwatchLogs: &evidently.ProjectDataDeliveryCloudwatchLogsArgs{
					LogGroup: pulumi.String("example-log-group-name"),
				},
			},
			Description: pulumi.String("Example Description"),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("example Project"),
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
import com.pulumi.aws.evidently.Project;
import com.pulumi.aws.evidently.ProjectArgs;
import com.pulumi.aws.evidently.inputs.ProjectDataDeliveryArgs;
import com.pulumi.aws.evidently.inputs.ProjectDataDeliveryCloudwatchLogsArgs;
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
        var example = new Project("example", ProjectArgs.builder()        
            .dataDelivery(ProjectDataDeliveryArgs.builder()
                .cloudwatchLogs(ProjectDataDeliveryCloudwatchLogsArgs.builder()
                    .logGroup("example-log-group-name")
                    .build())
                .build())
            .description("Example Description")
            .tags(Map.of("Key1", "example Project"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Project
    properties:
      dataDelivery:
        cloudwatchLogs:
          logGroup: example-log-group-name
      description: Example Description
      tags:
        Key1: example Project
```
{{% /example %}}
{{% example %}}
### Store evaluation events in an S3 bucket

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.evidently.Project("example", {
    dataDelivery: {
        s3Destination: {
            bucket: "example-bucket-name",
            prefix: "example",
        },
    },
    description: "Example Description",
    tags: {
        Key1: "example Project",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.evidently.Project("example",
    data_delivery=aws.evidently.ProjectDataDeliveryArgs(
        s3_destination=aws.evidently.ProjectDataDeliveryS3DestinationArgs(
            bucket="example-bucket-name",
            prefix="example",
        ),
    ),
    description="Example Description",
    tags={
        "Key1": "example Project",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Evidently.Project("example", new()
    {
        DataDelivery = new Aws.Evidently.Inputs.ProjectDataDeliveryArgs
        {
            S3Destination = new Aws.Evidently.Inputs.ProjectDataDeliveryS3DestinationArgs
            {
                Bucket = "example-bucket-name",
                Prefix = "example",
            },
        },
        Description = "Example Description",
        Tags = 
        {
            { "Key1", "example Project" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/evidently"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := evidently.NewProject(ctx, "example", &evidently.ProjectArgs{
			DataDelivery: &evidently.ProjectDataDeliveryArgs{
				S3Destination: &evidently.ProjectDataDeliveryS3DestinationArgs{
					Bucket: pulumi.String("example-bucket-name"),
					Prefix: pulumi.String("example"),
				},
			},
			Description: pulumi.String("Example Description"),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("example Project"),
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
import com.pulumi.aws.evidently.Project;
import com.pulumi.aws.evidently.ProjectArgs;
import com.pulumi.aws.evidently.inputs.ProjectDataDeliveryArgs;
import com.pulumi.aws.evidently.inputs.ProjectDataDeliveryS3DestinationArgs;
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
        var example = new Project("example", ProjectArgs.builder()        
            .dataDelivery(ProjectDataDeliveryArgs.builder()
                .s3Destination(ProjectDataDeliveryS3DestinationArgs.builder()
                    .bucket("example-bucket-name")
                    .prefix("example")
                    .build())
                .build())
            .description("Example Description")
            .tags(Map.of("Key1", "example Project"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:evidently:Project
    properties:
      dataDelivery:
        s3Destination:
          bucket: example-bucket-name
          prefix: example
      description: Example Description
      tags:
        Key1: example Project
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Evidently Project using the `arn`. For example:

```sh
 $ pulumi import aws:evidently/project:Project example arn:aws:evidently:us-east-1:123456789012:segment/example
```
 