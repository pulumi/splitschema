Manages a FSx Windows File System. See the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/what-is.html) for more information.

> **NOTE:** Either the `active_directory_id` argument or `self_managed_active_directory` configuration block must be specified.

{{% examples %}}
## Example Usage
{{% example %}}
### Using AWS Directory Service

Additional information for using AWS Directory Service with Windows File Systems can be found in the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/fsx-aws-managed-ad.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.WindowsFileSystem("example", {
    activeDirectoryId: aws_directory_service_directory.example.id,
    kmsKeyId: aws_kms_key.example.arn,
    storageCapacity: 300,
    subnetIds: [aws_subnet.example.id],
    throughputCapacity: 1024,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fsx.WindowsFileSystem("example",
    active_directory_id=aws_directory_service_directory["example"]["id"],
    kms_key_id=aws_kms_key["example"]["arn"],
    storage_capacity=300,
    subnet_ids=[aws_subnet["example"]["id"]],
    throughput_capacity=1024)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Fsx.WindowsFileSystem("example", new()
    {
        ActiveDirectoryId = aws_directory_service_directory.Example.Id,
        KmsKeyId = aws_kms_key.Example.Arn,
        StorageCapacity = 300,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        ThroughputCapacity = 1024,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewWindowsFileSystem(ctx, "example", &fsx.WindowsFileSystemArgs{
			ActiveDirectoryId: pulumi.Any(aws_directory_service_directory.Example.Id),
			KmsKeyId:          pulumi.Any(aws_kms_key.Example.Arn),
			StorageCapacity:   pulumi.Int(300),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			ThroughputCapacity: pulumi.Int(1024),
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
import com.pulumi.aws.fsx.WindowsFileSystem;
import com.pulumi.aws.fsx.WindowsFileSystemArgs;
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
        var example = new WindowsFileSystem("example", WindowsFileSystemArgs.builder()        
            .activeDirectoryId(aws_directory_service_directory.example().id())
            .kmsKeyId(aws_kms_key.example().arn())
            .storageCapacity(300)
            .subnetIds(aws_subnet.example().id())
            .throughputCapacity(1024)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:fsx:WindowsFileSystem
    properties:
      activeDirectoryId: ${aws_directory_service_directory.example.id}
      kmsKeyId: ${aws_kms_key.example.arn}
      storageCapacity: 300
      subnetIds:
        - ${aws_subnet.example.id}
      throughputCapacity: 1024
```
{{% /example %}}
{{% example %}}
### Using a Self-Managed Microsoft Active Directory

Additional information for using AWS Directory Service with Windows File Systems can be found in the [FSx Windows Guide](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/self-managed-AD.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.fsx.WindowsFileSystem("example", {
    kmsKeyId: aws_kms_key.example.arn,
    storageCapacity: 300,
    subnetIds: [aws_subnet.example.id],
    throughputCapacity: 1024,
    selfManagedActiveDirectory: {
        dnsIps: [
            "10.0.0.111",
            "10.0.0.222",
        ],
        domainName: "corp.example.com",
        password: "avoid-plaintext-passwords",
        username: "Admin",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fsx.WindowsFileSystem("example",
    kms_key_id=aws_kms_key["example"]["arn"],
    storage_capacity=300,
    subnet_ids=[aws_subnet["example"]["id"]],
    throughput_capacity=1024,
    self_managed_active_directory=aws.fsx.WindowsFileSystemSelfManagedActiveDirectoryArgs(
        dns_ips=[
            "10.0.0.111",
            "10.0.0.222",
        ],
        domain_name="corp.example.com",
        password="avoid-plaintext-passwords",
        username="Admin",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Fsx.WindowsFileSystem("example", new()
    {
        KmsKeyId = aws_kms_key.Example.Arn,
        StorageCapacity = 300,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        ThroughputCapacity = 1024,
        SelfManagedActiveDirectory = new Aws.Fsx.Inputs.WindowsFileSystemSelfManagedActiveDirectoryArgs
        {
            DnsIps = new[]
            {
                "10.0.0.111",
                "10.0.0.222",
            },
            DomainName = "corp.example.com",
            Password = "avoid-plaintext-passwords",
            Username = "Admin",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.NewWindowsFileSystem(ctx, "example", &fsx.WindowsFileSystemArgs{
			KmsKeyId:        pulumi.Any(aws_kms_key.Example.Arn),
			StorageCapacity: pulumi.Int(300),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			ThroughputCapacity: pulumi.Int(1024),
			SelfManagedActiveDirectory: &fsx.WindowsFileSystemSelfManagedActiveDirectoryArgs{
				DnsIps: pulumi.StringArray{
					pulumi.String("10.0.0.111"),
					pulumi.String("10.0.0.222"),
				},
				DomainName: pulumi.String("corp.example.com"),
				Password:   pulumi.String("avoid-plaintext-passwords"),
				Username:   pulumi.String("Admin"),
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
import com.pulumi.aws.fsx.WindowsFileSystem;
import com.pulumi.aws.fsx.WindowsFileSystemArgs;
import com.pulumi.aws.fsx.inputs.WindowsFileSystemSelfManagedActiveDirectoryArgs;
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
        var example = new WindowsFileSystem("example", WindowsFileSystemArgs.builder()        
            .kmsKeyId(aws_kms_key.example().arn())
            .storageCapacity(300)
            .subnetIds(aws_subnet.example().id())
            .throughputCapacity(1024)
            .selfManagedActiveDirectory(WindowsFileSystemSelfManagedActiveDirectoryArgs.builder()
                .dnsIps(                
                    "10.0.0.111",
                    "10.0.0.222")
                .domainName("corp.example.com")
                .password("avoid-plaintext-passwords")
                .username("Admin")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:fsx:WindowsFileSystem
    properties:
      kmsKeyId: ${aws_kms_key.example.arn}
      storageCapacity: 300
      subnetIds:
        - ${aws_subnet.example.id}
      throughputCapacity: 1024
      selfManagedActiveDirectory:
        dnsIps:
          - 10.0.0.111
          - 10.0.0.222
        domainName: corp.example.com
        password: avoid-plaintext-passwords
        username: Admin
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx File Systems using the `id`. For example:

```sh
 $ pulumi import aws:fsx/windowsFileSystem:WindowsFileSystem example fs-543ab12b1ca672f33
```
 Certain resource arguments, like `security_group_ids` and the `self_managed_active_directory` configuation block `password`, do not have a FSx API method for reading the information after creation. If these arguments are set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

