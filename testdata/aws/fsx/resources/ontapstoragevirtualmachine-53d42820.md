Manages a FSx Storage Virtual Machine.
See the [FSx ONTAP User Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/managing-svms.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OntapStorageVirtualMachine("test", {fileSystemId: aws_fsx_ontap_file_system.test.id});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OntapStorageVirtualMachine("test", file_system_id=aws_fsx_ontap_file_system["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OntapStorageVirtualMachine("test", new()
    {
        FileSystemId = aws_fsx_ontap_file_system.Test.Id,
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
		_, err := fsx.NewOntapStorageVirtualMachine(ctx, "test", &fsx.OntapStorageVirtualMachineArgs{
			FileSystemId: pulumi.Any(aws_fsx_ontap_file_system.Test.Id),
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
import com.pulumi.aws.fsx.OntapStorageVirtualMachine;
import com.pulumi.aws.fsx.OntapStorageVirtualMachineArgs;
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
        var test = new OntapStorageVirtualMachine("test", OntapStorageVirtualMachineArgs.builder()        
            .fileSystemId(aws_fsx_ontap_file_system.test().id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OntapStorageVirtualMachine
    properties:
      fileSystemId: ${aws_fsx_ontap_file_system.test.id}
```
{{% /example %}}
{{% example %}}
### Using a Self-Managed Microsoft Active Directory

Additional information for using AWS Directory Service with ONTAP File Systems can be found in the [FSx ONTAP Guide](https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/self-managed-AD.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.fsx.OntapStorageVirtualMachine("test", {
    fileSystemId: aws_fsx_ontap_file_system.test.id,
    activeDirectoryConfiguration: {
        netbiosName: "mysvm",
        selfManagedActiveDirectoryConfiguration: {
            dnsIps: [
                "10.0.0.111",
                "10.0.0.222",
            ],
            domainName: "corp.example.com",
            password: "avoid-plaintext-passwords",
            username: "Admin",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.fsx.OntapStorageVirtualMachine("test",
    file_system_id=aws_fsx_ontap_file_system["test"]["id"],
    active_directory_configuration=aws.fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs(
        netbios_name="mysvm",
        self_managed_active_directory_configuration=aws.fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs(
            dns_ips=[
                "10.0.0.111",
                "10.0.0.222",
            ],
            domain_name="corp.example.com",
            password="avoid-plaintext-passwords",
            username="Admin",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Fsx.OntapStorageVirtualMachine("test", new()
    {
        FileSystemId = aws_fsx_ontap_file_system.Test.Id,
        ActiveDirectoryConfiguration = new Aws.Fsx.Inputs.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs
        {
            NetbiosName = "mysvm",
            SelfManagedActiveDirectoryConfiguration = new Aws.Fsx.Inputs.OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs
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
		_, err := fsx.NewOntapStorageVirtualMachine(ctx, "test", &fsx.OntapStorageVirtualMachineArgs{
			FileSystemId: pulumi.Any(aws_fsx_ontap_file_system.Test.Id),
			ActiveDirectoryConfiguration: &fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs{
				NetbiosName: pulumi.String("mysvm"),
				SelfManagedActiveDirectoryConfiguration: &fsx.OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs{
					DnsIps: pulumi.StringArray{
						pulumi.String("10.0.0.111"),
						pulumi.String("10.0.0.222"),
					},
					DomainName: pulumi.String("corp.example.com"),
					Password:   pulumi.String("avoid-plaintext-passwords"),
					Username:   pulumi.String("Admin"),
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
import com.pulumi.aws.fsx.OntapStorageVirtualMachine;
import com.pulumi.aws.fsx.OntapStorageVirtualMachineArgs;
import com.pulumi.aws.fsx.inputs.OntapStorageVirtualMachineActiveDirectoryConfigurationArgs;
import com.pulumi.aws.fsx.inputs.OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs;
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
        var test = new OntapStorageVirtualMachine("test", OntapStorageVirtualMachineArgs.builder()        
            .fileSystemId(aws_fsx_ontap_file_system.test().id())
            .activeDirectoryConfiguration(OntapStorageVirtualMachineActiveDirectoryConfigurationArgs.builder()
                .netbiosName("mysvm")
                .selfManagedActiveDirectoryConfiguration(OntapStorageVirtualMachineActiveDirectoryConfigurationSelfManagedActiveDirectoryConfigurationArgs.builder()
                    .dnsIps(                    
                        "10.0.0.111",
                        "10.0.0.222")
                    .domainName("corp.example.com")
                    .password("avoid-plaintext-passwords")
                    .username("Admin")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:fsx:OntapStorageVirtualMachine
    properties:
      fileSystemId: ${aws_fsx_ontap_file_system.test.id}
      activeDirectoryConfiguration:
        netbiosName: mysvm
        selfManagedActiveDirectoryConfiguration:
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

Using `pulumi import`, import FSx Storage Virtual Machine using the `id`. For example:

```sh
 $ pulumi import aws:fsx/ontapStorageVirtualMachine:OntapStorageVirtualMachine example svm-12345678abcdef123
```
 Certain resource arguments, like `svm_admin_password` and the `self_managed_active_directory` configuation block `password`, do not have a FSx API method for reading the information after creation. If these arguments are set in the Pulumi program on an imported resource, Pulumi will always show a difference. To workaround this behavior, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

