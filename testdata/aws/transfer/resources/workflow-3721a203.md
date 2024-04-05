Provides a AWS Transfer Workflow resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic single step example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Workflow("example", {steps: [{
    deleteStepDetails: {
        name: "example",
        sourceFileLocation: "${original.file}",
    },
    type: "DELETE",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Workflow("example", steps=[aws.transfer.WorkflowStepArgs(
    delete_step_details=aws.transfer.WorkflowStepDeleteStepDetailsArgs(
        name="example",
        source_file_location="${original.file}",
    ),
    type="DELETE",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Workflow("example", new()
    {
        Steps = new[]
        {
            new Aws.Transfer.Inputs.WorkflowStepArgs
            {
                DeleteStepDetails = new Aws.Transfer.Inputs.WorkflowStepDeleteStepDetailsArgs
                {
                    Name = "example",
                    SourceFileLocation = "${original.file}",
                },
                Type = "DELETE",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewWorkflow(ctx, "example", &transfer.WorkflowArgs{
			Steps: transfer.WorkflowStepArray{
				&transfer.WorkflowStepArgs{
					DeleteStepDetails: &transfer.WorkflowStepDeleteStepDetailsArgs{
						Name:               pulumi.String("example"),
						SourceFileLocation: pulumi.String("${original.file}"),
					},
					Type: pulumi.String("DELETE"),
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
import com.pulumi.aws.transfer.Workflow;
import com.pulumi.aws.transfer.WorkflowArgs;
import com.pulumi.aws.transfer.inputs.WorkflowStepArgs;
import com.pulumi.aws.transfer.inputs.WorkflowStepDeleteStepDetailsArgs;
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
        var example = new Workflow("example", WorkflowArgs.builder()        
            .steps(WorkflowStepArgs.builder()
                .deleteStepDetails(WorkflowStepDeleteStepDetailsArgs.builder()
                    .name("example")
                    .sourceFileLocation("${original.file}")
                    .build())
                .type("DELETE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Workflow
    properties:
      steps:
        - deleteStepDetails:
            name: example
            sourceFileLocation: ${original.file}
          type: DELETE
```
{{% /example %}}
{{% example %}}
### Multistep example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Workflow("example", {steps: [
    {
        customStepDetails: {
            name: "example",
            sourceFileLocation: "${original.file}",
            target: aws_lambda_function.example.arn,
            timeoutSeconds: 60,
        },
        type: "CUSTOM",
    },
    {
        tagStepDetails: {
            name: "example",
            sourceFileLocation: "${original.file}",
            tags: [{
                key: "Name",
                value: "Hello World",
            }],
        },
        type: "TAG",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Workflow("example", steps=[
    aws.transfer.WorkflowStepArgs(
        custom_step_details=aws.transfer.WorkflowStepCustomStepDetailsArgs(
            name="example",
            source_file_location="${original.file}",
            target=aws_lambda_function["example"]["arn"],
            timeout_seconds=60,
        ),
        type="CUSTOM",
    ),
    aws.transfer.WorkflowStepArgs(
        tag_step_details=aws.transfer.WorkflowStepTagStepDetailsArgs(
            name="example",
            source_file_location="${original.file}",
            tags=[aws.transfer.WorkflowStepTagStepDetailsTagArgs(
                key="Name",
                value="Hello World",
            )],
        ),
        type="TAG",
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
    var example = new Aws.Transfer.Workflow("example", new()
    {
        Steps = new[]
        {
            new Aws.Transfer.Inputs.WorkflowStepArgs
            {
                CustomStepDetails = new Aws.Transfer.Inputs.WorkflowStepCustomStepDetailsArgs
                {
                    Name = "example",
                    SourceFileLocation = "${original.file}",
                    Target = aws_lambda_function.Example.Arn,
                    TimeoutSeconds = 60,
                },
                Type = "CUSTOM",
            },
            new Aws.Transfer.Inputs.WorkflowStepArgs
            {
                TagStepDetails = new Aws.Transfer.Inputs.WorkflowStepTagStepDetailsArgs
                {
                    Name = "example",
                    SourceFileLocation = "${original.file}",
                    Tags = new[]
                    {
                        new Aws.Transfer.Inputs.WorkflowStepTagStepDetailsTagArgs
                        {
                            Key = "Name",
                            Value = "Hello World",
                        },
                    },
                },
                Type = "TAG",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewWorkflow(ctx, "example", &transfer.WorkflowArgs{
			Steps: transfer.WorkflowStepArray{
				&transfer.WorkflowStepArgs{
					CustomStepDetails: &transfer.WorkflowStepCustomStepDetailsArgs{
						Name:               pulumi.String("example"),
						SourceFileLocation: pulumi.String("${original.file}"),
						Target:             pulumi.Any(aws_lambda_function.Example.Arn),
						TimeoutSeconds:     pulumi.Int(60),
					},
					Type: pulumi.String("CUSTOM"),
				},
				&transfer.WorkflowStepArgs{
					TagStepDetails: &transfer.WorkflowStepTagStepDetailsArgs{
						Name:               pulumi.String("example"),
						SourceFileLocation: pulumi.String("${original.file}"),
						Tags: transfer.WorkflowStepTagStepDetailsTagArray{
							&transfer.WorkflowStepTagStepDetailsTagArgs{
								Key:   pulumi.String("Name"),
								Value: pulumi.String("Hello World"),
							},
						},
					},
					Type: pulumi.String("TAG"),
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
import com.pulumi.aws.transfer.Workflow;
import com.pulumi.aws.transfer.WorkflowArgs;
import com.pulumi.aws.transfer.inputs.WorkflowStepArgs;
import com.pulumi.aws.transfer.inputs.WorkflowStepCustomStepDetailsArgs;
import com.pulumi.aws.transfer.inputs.WorkflowStepTagStepDetailsArgs;
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
        var example = new Workflow("example", WorkflowArgs.builder()        
            .steps(            
                WorkflowStepArgs.builder()
                    .customStepDetails(WorkflowStepCustomStepDetailsArgs.builder()
                        .name("example")
                        .sourceFileLocation("${original.file}")
                        .target(aws_lambda_function.example().arn())
                        .timeoutSeconds(60)
                        .build())
                    .type("CUSTOM")
                    .build(),
                WorkflowStepArgs.builder()
                    .tagStepDetails(WorkflowStepTagStepDetailsArgs.builder()
                        .name("example")
                        .sourceFileLocation("${original.file}")
                        .tags(WorkflowStepTagStepDetailsTagArgs.builder()
                            .key("Name")
                            .value("Hello World")
                            .build())
                        .build())
                    .type("TAG")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Workflow
    properties:
      steps:
        - customStepDetails:
            name: example
            sourceFileLocation: ${original.file}
            target: ${aws_lambda_function.example.arn}
            timeoutSeconds: 60
          type: CUSTOM
        - tagStepDetails:
            name: example
            sourceFileLocation: ${original.file}
            tags:
              - key: Name
                value: Hello World
          type: TAG
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer Workflows using the `worflow_id`. For example:

```sh
 $ pulumi import aws:transfer/workflow:Workflow example example
```
 