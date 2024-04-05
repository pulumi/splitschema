Manages an AWS DataSync FSx Windows Location.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.LocationFsxWindows("example", {
    fsxFilesystemArn: aws_fsx_windows_file_system.example.arn,
    user: "SomeUser",
    password: "SuperSecretPassw0rd",
    securityGroupArns: [aws_security_group.example.arn],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.LocationFsxWindows("example",
    fsx_filesystem_arn=aws_fsx_windows_file_system["example"]["arn"],
    user="SomeUser",
    password="SuperSecretPassw0rd",
    security_group_arns=[aws_security_group["example"]["arn"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.LocationFsxWindows("example", new()
    {
        FsxFilesystemArn = aws_fsx_windows_file_system.Example.Arn,
        User = "SomeUser",
        Password = "SuperSecretPassw0rd",
        SecurityGroupArns = new[]
        {
            aws_security_group.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewLocationFsxWindows(ctx, "example", &datasync.LocationFsxWindowsArgs{
			FsxFilesystemArn: pulumi.Any(aws_fsx_windows_file_system.Example.Arn),
			User:             pulumi.String("SomeUser"),
			Password:         pulumi.String("SuperSecretPassw0rd"),
			SecurityGroupArns: pulumi.StringArray{
				aws_security_group.Example.Arn,
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
import com.pulumi.aws.datasync.LocationFsxWindows;
import com.pulumi.aws.datasync.LocationFsxWindowsArgs;
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
        var example = new LocationFsxWindows("example", LocationFsxWindowsArgs.builder()        
            .fsxFilesystemArn(aws_fsx_windows_file_system.example().arn())
            .user("SomeUser")
            .password("SuperSecretPassw0rd")
            .securityGroupArns(aws_security_group.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:LocationFsxWindows
    properties:
      fsxFilesystemArn: ${aws_fsx_windows_file_system.example.arn}
      user: SomeUser
      password: SuperSecretPassw0rd
      securityGroupArns:
        - ${aws_security_group.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_fsx_windows_file_system` using the `DataSync-ARN#FSx-Windows-ARN`. For example:

```sh
 $ pulumi import aws:datasync/locationFsxWindows:LocationFsxWindows example arn:aws:datasync:us-west-2:123456789012:location/loc-12345678901234567#arn:aws:fsx:us-west-2:476956259333:file-system/fs-08e04cd442c1bb94a
```
 