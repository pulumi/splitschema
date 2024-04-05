Resource for managing an AWS Audit Manager Assessment Delegation.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.auditmanager.AssessmentDelegation("example", {
    assessmentId: aws_auditmanager_assessment.example.id,
    roleArn: aws_iam_role.example.arn,
    roleType: "RESOURCE_OWNER",
    controlSetId: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.AssessmentDelegation("example",
    assessment_id=aws_auditmanager_assessment["example"]["id"],
    role_arn=aws_iam_role["example"]["arn"],
    role_type="RESOURCE_OWNER",
    control_set_id="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Auditmanager.AssessmentDelegation("example", new()
    {
        AssessmentId = aws_auditmanager_assessment.Example.Id,
        RoleArn = aws_iam_role.Example.Arn,
        RoleType = "RESOURCE_OWNER",
        ControlSetId = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewAssessmentDelegation(ctx, "example", &auditmanager.AssessmentDelegationArgs{
			AssessmentId: pulumi.Any(aws_auditmanager_assessment.Example.Id),
			RoleArn:      pulumi.Any(aws_iam_role.Example.Arn),
			RoleType:     pulumi.String("RESOURCE_OWNER"),
			ControlSetId: pulumi.String("example"),
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
import com.pulumi.aws.auditmanager.AssessmentDelegation;
import com.pulumi.aws.auditmanager.AssessmentDelegationArgs;
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
        var example = new AssessmentDelegation("example", AssessmentDelegationArgs.builder()        
            .assessmentId(aws_auditmanager_assessment.example().id())
            .roleArn(aws_iam_role.example().arn())
            .roleType("RESOURCE_OWNER")
            .controlSetId("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:auditmanager:AssessmentDelegation
    properties:
      assessmentId: ${aws_auditmanager_assessment.example.id}
      roleArn: ${aws_iam_role.example.arn}
      roleType: RESOURCE_OWNER
      controlSetId: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Assessment Delegation using the `id`. For example:

```sh
 $ pulumi import aws:auditmanager/assessmentDelegation:AssessmentDelegation example abcdef-123456,arn:aws:iam::012345678901:role/example,example
```
 