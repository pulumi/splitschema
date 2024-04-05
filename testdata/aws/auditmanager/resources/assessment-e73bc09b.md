Resource for managing an AWS Audit Manager Assessment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.auditmanager.Assessment("test", {
    assessmentReportsDestination: {
        destination: `s3://${aws_s3_bucket.test.id}`,
        destinationType: "S3",
    },
    frameworkId: aws_auditmanager_framework.test.id,
    roles: [{
        roleArn: aws_iam_role.test.arn,
        roleType: "PROCESS_OWNER",
    }],
    scope: {
        awsAccounts: [{
            id: data.aws_caller_identity.current.account_id,
        }],
        awsServices: [{
            serviceName: "S3",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.auditmanager.Assessment("test",
    assessment_reports_destination=aws.auditmanager.AssessmentAssessmentReportsDestinationArgs(
        destination=f"s3://{aws_s3_bucket['test']['id']}",
        destination_type="S3",
    ),
    framework_id=aws_auditmanager_framework["test"]["id"],
    roles=[aws.auditmanager.AssessmentRoleArgs(
        role_arn=aws_iam_role["test"]["arn"],
        role_type="PROCESS_OWNER",
    )],
    scope=aws.auditmanager.AssessmentScopeArgs(
        aws_accounts=[aws.auditmanager.AssessmentScopeAwsAccountArgs(
            id=data["aws_caller_identity"]["current"]["account_id"],
        )],
        aws_services=[aws.auditmanager.AssessmentScopeAwsServiceArgs(
            service_name="S3",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Auditmanager.Assessment("test", new()
    {
        AssessmentReportsDestination = new Aws.Auditmanager.Inputs.AssessmentAssessmentReportsDestinationArgs
        {
            Destination = $"s3://{aws_s3_bucket.Test.Id}",
            DestinationType = "S3",
        },
        FrameworkId = aws_auditmanager_framework.Test.Id,
        Roles = new[]
        {
            new Aws.Auditmanager.Inputs.AssessmentRoleArgs
            {
                RoleArn = aws_iam_role.Test.Arn,
                RoleType = "PROCESS_OWNER",
            },
        },
        Scope = new Aws.Auditmanager.Inputs.AssessmentScopeArgs
        {
            AwsAccounts = new[]
            {
                new Aws.Auditmanager.Inputs.AssessmentScopeAwsAccountArgs
                {
                    Id = data.Aws_caller_identity.Current.Account_id,
                },
            },
            AwsServices = new[]
            {
                new Aws.Auditmanager.Inputs.AssessmentScopeAwsServiceArgs
                {
                    ServiceName = "S3",
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewAssessment(ctx, "test", &auditmanager.AssessmentArgs{
			AssessmentReportsDestination: &auditmanager.AssessmentAssessmentReportsDestinationArgs{
				Destination:     pulumi.String(fmt.Sprintf("s3://%v", aws_s3_bucket.Test.Id)),
				DestinationType: pulumi.String("S3"),
			},
			FrameworkId: pulumi.Any(aws_auditmanager_framework.Test.Id),
			Roles: auditmanager.AssessmentRoleArray{
				&auditmanager.AssessmentRoleArgs{
					RoleArn:  pulumi.Any(aws_iam_role.Test.Arn),
					RoleType: pulumi.String("PROCESS_OWNER"),
				},
			},
			Scope: &auditmanager.AssessmentScopeArgs{
				AwsAccounts: auditmanager.AssessmentScopeAwsAccountArray{
					&auditmanager.AssessmentScopeAwsAccountArgs{
						Id: pulumi.Any(data.Aws_caller_identity.Current.Account_id),
					},
				},
				AwsServices: auditmanager.AssessmentScopeAwsServiceArray{
					&auditmanager.AssessmentScopeAwsServiceArgs{
						ServiceName: pulumi.String("S3"),
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
import com.pulumi.aws.auditmanager.Assessment;
import com.pulumi.aws.auditmanager.AssessmentArgs;
import com.pulumi.aws.auditmanager.inputs.AssessmentAssessmentReportsDestinationArgs;
import com.pulumi.aws.auditmanager.inputs.AssessmentRoleArgs;
import com.pulumi.aws.auditmanager.inputs.AssessmentScopeArgs;
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
        var test = new Assessment("test", AssessmentArgs.builder()        
            .assessmentReportsDestination(AssessmentAssessmentReportsDestinationArgs.builder()
                .destination(String.format("s3://%s", aws_s3_bucket.test().id()))
                .destinationType("S3")
                .build())
            .frameworkId(aws_auditmanager_framework.test().id())
            .roles(AssessmentRoleArgs.builder()
                .roleArn(aws_iam_role.test().arn())
                .roleType("PROCESS_OWNER")
                .build())
            .scope(AssessmentScopeArgs.builder()
                .awsAccounts(AssessmentScopeAwsAccountArgs.builder()
                    .id(data.aws_caller_identity().current().account_id())
                    .build())
                .awsServices(AssessmentScopeAwsServiceArgs.builder()
                    .serviceName("S3")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:auditmanager:Assessment
    properties:
      assessmentReportsDestination:
        destination: s3://${aws_s3_bucket.test.id}
        destinationType: S3
      frameworkId: ${aws_auditmanager_framework.test.id}
      roles:
        - roleArn: ${aws_iam_role.test.arn}
          roleType: PROCESS_OWNER
      scope:
        awsAccounts:
          - id: ${data.aws_caller_identity.current.account_id}
        awsServices:
          - serviceName: S3
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Assessments using the assessment `id`. For example:

```sh
 $ pulumi import aws:auditmanager/assessment:Assessment example abc123-de45
```
 