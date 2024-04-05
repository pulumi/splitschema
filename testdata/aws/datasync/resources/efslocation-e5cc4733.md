Manages an AWS DataSync EFS Location.

> **NOTE:** The EFS File System must have a mounted EFS Mount Target before creating this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.EfsLocation("example", {
    efsFileSystemArn: aws_efs_mount_target.example.file_system_arn,
    ec2Config: {
        securityGroupArns: [aws_security_group.example.arn],
        subnetArn: aws_subnet.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.EfsLocation("example",
    efs_file_system_arn=aws_efs_mount_target["example"]["file_system_arn"],
    ec2_config=aws.datasync.EfsLocationEc2ConfigArgs(
        security_group_arns=[aws_security_group["example"]["arn"]],
        subnet_arn=aws_subnet["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.EfsLocation("example", new()
    {
        EfsFileSystemArn = aws_efs_mount_target.Example.File_system_arn,
        Ec2Config = new Aws.DataSync.Inputs.EfsLocationEc2ConfigArgs
        {
            SecurityGroupArns = new[]
            {
                aws_security_group.Example.Arn,
            },
            SubnetArn = aws_subnet.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewEfsLocation(ctx, "example", &datasync.EfsLocationArgs{
			EfsFileSystemArn: pulumi.Any(aws_efs_mount_target.Example.File_system_arn),
			Ec2Config: &datasync.EfsLocationEc2ConfigArgs{
				SecurityGroupArns: pulumi.StringArray{
					aws_security_group.Example.Arn,
				},
				SubnetArn: pulumi.Any(aws_subnet.Example.Arn),
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
import com.pulumi.aws.datasync.EfsLocation;
import com.pulumi.aws.datasync.EfsLocationArgs;
import com.pulumi.aws.datasync.inputs.EfsLocationEc2ConfigArgs;
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
        var example = new EfsLocation("example", EfsLocationArgs.builder()        
            .efsFileSystemArn(aws_efs_mount_target.example().file_system_arn())
            .ec2Config(EfsLocationEc2ConfigArgs.builder()
                .securityGroupArns(aws_security_group.example().arn())
                .subnetArn(aws_subnet.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:EfsLocation
    properties:
      # The below example uses aws_efs_mount_target as a reference to ensure a mount target already exists when resource creation occurs.
      #   # You can accomplish the same behavior with depends_on or an aws_efs_mount_target data source reference.
      efsFileSystemArn: ${aws_efs_mount_target.example.file_system_arn}
      ec2Config:
        securityGroupArns:
          - ${aws_security_group.example.arn}
        subnetArn: ${aws_subnet.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_efs` using the DataSync Task Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/efsLocation:EfsLocation example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 