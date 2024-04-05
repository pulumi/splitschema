Provides an Elastic MapReduce Studio.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emr.Studio("example", {
    authMode: "SSO",
    defaultS3Location: `s3://${aws_s3_bucket.test.bucket}/test`,
    engineSecurityGroupId: aws_security_group.test.id,
    serviceRole: aws_iam_role.test.arn,
    subnetIds: [aws_subnet.test.id],
    userRole: aws_iam_role.test.arn,
    vpcId: aws_vpc.test.id,
    workspaceSecurityGroupId: aws_security_group.test.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emr.Studio("example",
    auth_mode="SSO",
    default_s3_location=f"s3://{aws_s3_bucket['test']['bucket']}/test",
    engine_security_group_id=aws_security_group["test"]["id"],
    service_role=aws_iam_role["test"]["arn"],
    subnet_ids=[aws_subnet["test"]["id"]],
    user_role=aws_iam_role["test"]["arn"],
    vpc_id=aws_vpc["test"]["id"],
    workspace_security_group_id=aws_security_group["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Emr.Studio("example", new()
    {
        AuthMode = "SSO",
        DefaultS3Location = $"s3://{aws_s3_bucket.Test.Bucket}/test",
        EngineSecurityGroupId = aws_security_group.Test.Id,
        ServiceRole = aws_iam_role.Test.Arn,
        SubnetIds = new[]
        {
            aws_subnet.Test.Id,
        },
        UserRole = aws_iam_role.Test.Arn,
        VpcId = aws_vpc.Test.Id,
        WorkspaceSecurityGroupId = aws_security_group.Test.Id,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewStudio(ctx, "example", &emr.StudioArgs{
			AuthMode:              pulumi.String("SSO"),
			DefaultS3Location:     pulumi.String(fmt.Sprintf("s3://%v/test", aws_s3_bucket.Test.Bucket)),
			EngineSecurityGroupId: pulumi.Any(aws_security_group.Test.Id),
			ServiceRole:           pulumi.Any(aws_iam_role.Test.Arn),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Test.Id,
			},
			UserRole:                 pulumi.Any(aws_iam_role.Test.Arn),
			VpcId:                    pulumi.Any(aws_vpc.Test.Id),
			WorkspaceSecurityGroupId: pulumi.Any(aws_security_group.Test.Id),
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
import com.pulumi.aws.emr.Studio;
import com.pulumi.aws.emr.StudioArgs;
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
        var example = new Studio("example", StudioArgs.builder()        
            .authMode("SSO")
            .defaultS3Location(String.format("s3://%s/test", aws_s3_bucket.test().bucket()))
            .engineSecurityGroupId(aws_security_group.test().id())
            .serviceRole(aws_iam_role.test().arn())
            .subnetIds(aws_subnet.test().id())
            .userRole(aws_iam_role.test().arn())
            .vpcId(aws_vpc.test().id())
            .workspaceSecurityGroupId(aws_security_group.test().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emr:Studio
    properties:
      authMode: SSO
      defaultS3Location: s3://${aws_s3_bucket.test.bucket}/test
      engineSecurityGroupId: ${aws_security_group.test.id}
      serviceRole: ${aws_iam_role.test.arn}
      subnetIds:
        - ${aws_subnet.test.id}
      userRole: ${aws_iam_role.test.arn}
      vpcId: ${aws_vpc.test.id}
      workspaceSecurityGroupId: ${aws_security_group.test.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR studios using the `id`. For example:

```sh
 $ pulumi import aws:emr/studio:Studio studio es-123456ABCDEF
```
 