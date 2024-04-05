Manages an individual ECS resource tag. This resource should only be used in cases where ECS resources are created outside the provider (e.g., ECS Clusters implicitly created by Batch Compute Environments).

> **NOTE:** This tagging resource should not be combined with the resource for managing the parent resource. For example, using `aws.ecs.Cluster` and `aws.ecs.Tag` to manage tags of the same ECS Cluster will cause a perpetual difference where the `aws.ecs.Cluster` resource will try to remove the tag being added by the `aws.ecs.Tag` resource.

> **NOTE:** This tagging resource does not use the provider `ignore_tags` configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleComputeEnvironment = new aws.batch.ComputeEnvironment("exampleComputeEnvironment", {
    computeEnvironmentName: "example",
    serviceRole: aws_iam_role.example.arn,
    type: "UNMANAGED",
});
const exampleTag = new aws.ecs.Tag("exampleTag", {
    resourceArn: exampleComputeEnvironment.ecsClusterArn,
    key: "Name",
    value: "Hello World",
});
```
```python
import pulumi
import pulumi_aws as aws

example_compute_environment = aws.batch.ComputeEnvironment("exampleComputeEnvironment",
    compute_environment_name="example",
    service_role=aws_iam_role["example"]["arn"],
    type="UNMANAGED")
example_tag = aws.ecs.Tag("exampleTag",
    resource_arn=example_compute_environment.ecs_cluster_arn,
    key="Name",
    value="Hello World")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleComputeEnvironment = new Aws.Batch.ComputeEnvironment("exampleComputeEnvironment", new()
    {
        ComputeEnvironmentName = "example",
        ServiceRole = aws_iam_role.Example.Arn,
        Type = "UNMANAGED",
    });

    var exampleTag = new Aws.Ecs.Tag("exampleTag", new()
    {
        ResourceArn = exampleComputeEnvironment.EcsClusterArn,
        Key = "Name",
        Value = "Hello World",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleComputeEnvironment, err := batch.NewComputeEnvironment(ctx, "exampleComputeEnvironment", &batch.ComputeEnvironmentArgs{
			ComputeEnvironmentName: pulumi.String("example"),
			ServiceRole:            pulumi.Any(aws_iam_role.Example.Arn),
			Type:                   pulumi.String("UNMANAGED"),
		})
		if err != nil {
			return err
		}
		_, err = ecs.NewTag(ctx, "exampleTag", &ecs.TagArgs{
			ResourceArn: exampleComputeEnvironment.EcsClusterArn,
			Key:         pulumi.String("Name"),
			Value:       pulumi.String("Hello World"),
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
import com.pulumi.aws.batch.ComputeEnvironment;
import com.pulumi.aws.batch.ComputeEnvironmentArgs;
import com.pulumi.aws.ecs.Tag;
import com.pulumi.aws.ecs.TagArgs;
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
        var exampleComputeEnvironment = new ComputeEnvironment("exampleComputeEnvironment", ComputeEnvironmentArgs.builder()        
            .computeEnvironmentName("example")
            .serviceRole(aws_iam_role.example().arn())
            .type("UNMANAGED")
            .build());

        var exampleTag = new Tag("exampleTag", TagArgs.builder()        
            .resourceArn(exampleComputeEnvironment.ecsClusterArn())
            .key("Name")
            .value("Hello World")
            .build());

    }
}
```
```yaml
resources:
  exampleComputeEnvironment:
    type: aws:batch:ComputeEnvironment
    properties:
      computeEnvironmentName: example
      serviceRole: ${aws_iam_role.example.arn}
      type: UNMANAGED
  exampleTag:
    type: aws:ecs:Tag
    properties:
      resourceArn: ${exampleComputeEnvironment.ecsClusterArn}
      key: Name
      value: Hello World
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ecs_tag` using the ECS resource identifier and key, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:ecs/tag:Tag example arn:aws:ecs:us-east-1:123456789012:cluster/example,Name
```
 