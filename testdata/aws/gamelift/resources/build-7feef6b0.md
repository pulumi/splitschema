Provides an GameLift Build resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.gamelift.Build("test", {
    operatingSystem: "WINDOWS_2012",
    storageLocation: {
        bucket: aws_s3_bucket.test.id,
        key: aws_s3_object.test.key,
        roleArn: aws_iam_role.test.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.gamelift.Build("test",
    operating_system="WINDOWS_2012",
    storage_location=aws.gamelift.BuildStorageLocationArgs(
        bucket=aws_s3_bucket["test"]["id"],
        key=aws_s3_object["test"]["key"],
        role_arn=aws_iam_role["test"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.GameLift.Build("test", new()
    {
        OperatingSystem = "WINDOWS_2012",
        StorageLocation = new Aws.GameLift.Inputs.BuildStorageLocationArgs
        {
            Bucket = aws_s3_bucket.Test.Id,
            Key = aws_s3_object.Test.Key,
            RoleArn = aws_iam_role.Test.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewBuild(ctx, "test", &gamelift.BuildArgs{
			OperatingSystem: pulumi.String("WINDOWS_2012"),
			StorageLocation: &gamelift.BuildStorageLocationArgs{
				Bucket:  pulumi.Any(aws_s3_bucket.Test.Id),
				Key:     pulumi.Any(aws_s3_object.Test.Key),
				RoleArn: pulumi.Any(aws_iam_role.Test.Arn),
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
import com.pulumi.aws.gamelift.Build;
import com.pulumi.aws.gamelift.BuildArgs;
import com.pulumi.aws.gamelift.inputs.BuildStorageLocationArgs;
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
        var test = new Build("test", BuildArgs.builder()        
            .operatingSystem("WINDOWS_2012")
            .storageLocation(BuildStorageLocationArgs.builder()
                .bucket(aws_s3_bucket.test().id())
                .key(aws_s3_object.test().key())
                .roleArn(aws_iam_role.test().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:gamelift:Build
    properties:
      operatingSystem: WINDOWS_2012
      storageLocation:
        bucket: ${aws_s3_bucket.test.id}
        key: ${aws_s3_object.test.key}
        roleArn: ${aws_iam_role.test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GameLift Builds using the ID. For example:

```sh
 $ pulumi import aws:gamelift/build:Build example <build-id>
```
 