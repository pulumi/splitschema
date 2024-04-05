Provides an AppStream Directory Config.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appstream.DirectoryConfig("example", {
    directoryName: "NAME OF DIRECTORY",
    organizationalUnitDistinguishedNames: ["DISTINGUISHED NAME"],
    serviceAccountCredentials: {
        accountName: "NAME OF ACCOUNT",
        accountPassword: "PASSWORD OF ACCOUNT",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appstream.DirectoryConfig("example",
    directory_name="NAME OF DIRECTORY",
    organizational_unit_distinguished_names=["DISTINGUISHED NAME"],
    service_account_credentials=aws.appstream.DirectoryConfigServiceAccountCredentialsArgs(
        account_name="NAME OF ACCOUNT",
        account_password="PASSWORD OF ACCOUNT",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppStream.DirectoryConfig("example", new()
    {
        DirectoryName = "NAME OF DIRECTORY",
        OrganizationalUnitDistinguishedNames = new[]
        {
            "DISTINGUISHED NAME",
        },
        ServiceAccountCredentials = new Aws.AppStream.Inputs.DirectoryConfigServiceAccountCredentialsArgs
        {
            AccountName = "NAME OF ACCOUNT",
            AccountPassword = "PASSWORD OF ACCOUNT",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appstream"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appstream.NewDirectoryConfig(ctx, "example", &appstream.DirectoryConfigArgs{
			DirectoryName: pulumi.String("NAME OF DIRECTORY"),
			OrganizationalUnitDistinguishedNames: pulumi.StringArray{
				pulumi.String("DISTINGUISHED NAME"),
			},
			ServiceAccountCredentials: &appstream.DirectoryConfigServiceAccountCredentialsArgs{
				AccountName:     pulumi.String("NAME OF ACCOUNT"),
				AccountPassword: pulumi.String("PASSWORD OF ACCOUNT"),
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
import com.pulumi.aws.appstream.DirectoryConfig;
import com.pulumi.aws.appstream.DirectoryConfigArgs;
import com.pulumi.aws.appstream.inputs.DirectoryConfigServiceAccountCredentialsArgs;
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
        var example = new DirectoryConfig("example", DirectoryConfigArgs.builder()        
            .directoryName("NAME OF DIRECTORY")
            .organizationalUnitDistinguishedNames("DISTINGUISHED NAME")
            .serviceAccountCredentials(DirectoryConfigServiceAccountCredentialsArgs.builder()
                .accountName("NAME OF ACCOUNT")
                .accountPassword("PASSWORD OF ACCOUNT")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appstream:DirectoryConfig
    properties:
      directoryName: NAME OF DIRECTORY
      organizationalUnitDistinguishedNames:
        - DISTINGUISHED NAME
      serviceAccountCredentials:
        accountName: NAME OF ACCOUNT
        accountPassword: PASSWORD OF ACCOUNT
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appstream_directory_config` using the id. For example:

```sh
 $ pulumi import aws:appstream/directoryConfig:DirectoryConfig example directoryNameExample
```
 