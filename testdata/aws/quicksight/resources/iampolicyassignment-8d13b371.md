Resource for managing an AWS QuickSight IAM Policy Assignment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.IamPolicyAssignment("example", {
    assignmentName: "example",
    assignmentStatus: "ENABLED",
    policyArn: aws_iam_policy.example.arn,
    identities: {
        users: [aws_quicksight_user.example.user_name],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.IamPolicyAssignment("example",
    assignment_name="example",
    assignment_status="ENABLED",
    policy_arn=aws_iam_policy["example"]["arn"],
    identities=aws.quicksight.IamPolicyAssignmentIdentitiesArgs(
        users=[aws_quicksight_user["example"]["user_name"]],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.IamPolicyAssignment("example", new()
    {
        AssignmentName = "example",
        AssignmentStatus = "ENABLED",
        PolicyArn = aws_iam_policy.Example.Arn,
        Identities = new Aws.Quicksight.Inputs.IamPolicyAssignmentIdentitiesArgs
        {
            Users = new[]
            {
                aws_quicksight_user.Example.User_name,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewIamPolicyAssignment(ctx, "example", &quicksight.IamPolicyAssignmentArgs{
			AssignmentName:   pulumi.String("example"),
			AssignmentStatus: pulumi.String("ENABLED"),
			PolicyArn:        pulumi.Any(aws_iam_policy.Example.Arn),
			Identities: &quicksight.IamPolicyAssignmentIdentitiesArgs{
				Users: pulumi.StringArray{
					aws_quicksight_user.Example.User_name,
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
import com.pulumi.aws.quicksight.IamPolicyAssignment;
import com.pulumi.aws.quicksight.IamPolicyAssignmentArgs;
import com.pulumi.aws.quicksight.inputs.IamPolicyAssignmentIdentitiesArgs;
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
        var example = new IamPolicyAssignment("example", IamPolicyAssignmentArgs.builder()        
            .assignmentName("example")
            .assignmentStatus("ENABLED")
            .policyArn(aws_iam_policy.example().arn())
            .identities(IamPolicyAssignmentIdentitiesArgs.builder()
                .users(aws_quicksight_user.example().user_name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:IamPolicyAssignment
    properties:
      assignmentName: example
      assignmentStatus: ENABLED
      policyArn: ${aws_iam_policy.example.arn}
      identities:
        users:
          - ${aws_quicksight_user.example.user_name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight IAM Policy Assignment using the AWS account ID, namespace, and assignment name separated by commas (`,`). For example:

```sh
 $ pulumi import aws:quicksight/iamPolicyAssignment:IamPolicyAssignment example 123456789012,default,example
```
 