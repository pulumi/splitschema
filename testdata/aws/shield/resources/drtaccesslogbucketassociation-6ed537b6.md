Resource for managing an AWS Shield DRT Access Log Bucket Association. Up to 10 log buckets can be associated for DRT Access sharing with the Shield Response Team (SRT).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testDrtAccessRoleArnAssociation = new aws.shield.DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", {roleArn: `arn:aws:iam:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${_var.shield_drt_access_role_name}`});
const testDrtAccessLogBucketAssociation = new aws.shield.DrtAccessLogBucketAssociation("testDrtAccessLogBucketAssociation", {
    logBucket: _var.shield_drt_access_log_bucket,
    roleArnAssociationId: testDrtAccessRoleArnAssociation.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test_drt_access_role_arn_association = aws.shield.DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", role_arn=f"arn:aws:iam:{data['aws_region']['current']['name']}:{data['aws_caller_identity']['current']['account_id']}:{var['shield_drt_access_role_name']}")
test_drt_access_log_bucket_association = aws.shield.DrtAccessLogBucketAssociation("testDrtAccessLogBucketAssociation",
    log_bucket=var["shield_drt_access_log_bucket"],
    role_arn_association_id=test_drt_access_role_arn_association.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testDrtAccessRoleArnAssociation = new Aws.Shield.DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", new()
    {
        RoleArn = $"arn:aws:iam:{data.Aws_region.Current.Name}:{data.Aws_caller_identity.Current.Account_id}:{@var.Shield_drt_access_role_name}",
    });

    var testDrtAccessLogBucketAssociation = new Aws.Shield.DrtAccessLogBucketAssociation("testDrtAccessLogBucketAssociation", new()
    {
        LogBucket = @var.Shield_drt_access_log_bucket,
        RoleArnAssociationId = testDrtAccessRoleArnAssociation.Id,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testDrtAccessRoleArnAssociation, err := shield.NewDrtAccessRoleArnAssociation(ctx, "testDrtAccessRoleArnAssociation", &shield.DrtAccessRoleArnAssociationArgs{
			RoleArn: pulumi.String(fmt.Sprintf("arn:aws:iam:%v:%v:%v", data.Aws_region.Current.Name, data.Aws_caller_identity.Current.Account_id, _var.Shield_drt_access_role_name)),
		})
		if err != nil {
			return err
		}
		_, err = shield.NewDrtAccessLogBucketAssociation(ctx, "testDrtAccessLogBucketAssociation", &shield.DrtAccessLogBucketAssociationArgs{
			LogBucket:            pulumi.Any(_var.Shield_drt_access_log_bucket),
			RoleArnAssociationId: testDrtAccessRoleArnAssociation.ID(),
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
import com.pulumi.aws.shield.DrtAccessRoleArnAssociation;
import com.pulumi.aws.shield.DrtAccessRoleArnAssociationArgs;
import com.pulumi.aws.shield.DrtAccessLogBucketAssociation;
import com.pulumi.aws.shield.DrtAccessLogBucketAssociationArgs;
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
        var testDrtAccessRoleArnAssociation = new DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", DrtAccessRoleArnAssociationArgs.builder()        
            .roleArn(String.format("arn:aws:iam:%s:%s:%s", data.aws_region().current().name(),data.aws_caller_identity().current().account_id(),var_.shield_drt_access_role_name()))
            .build());

        var testDrtAccessLogBucketAssociation = new DrtAccessLogBucketAssociation("testDrtAccessLogBucketAssociation", DrtAccessLogBucketAssociationArgs.builder()        
            .logBucket(var_.shield_drt_access_log_bucket())
            .roleArnAssociationId(testDrtAccessRoleArnAssociation.id())
            .build());

    }
}
```
```yaml
resources:
  testDrtAccessRoleArnAssociation:
    type: aws:shield:DrtAccessRoleArnAssociation
    properties:
      roleArn: arn:aws:iam:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${var.shield_drt_access_role_name}
  testDrtAccessLogBucketAssociation:
    type: aws:shield:DrtAccessLogBucketAssociation
    properties:
      logBucket: ${var.shield_drt_access_log_bucket}
      roleArnAssociationId: ${testDrtAccessRoleArnAssociation.id}
```
{{% /example %}}
{{% /examples %}}