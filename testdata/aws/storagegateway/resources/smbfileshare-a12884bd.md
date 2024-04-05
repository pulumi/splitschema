Manages an AWS Storage Gateway SMB File Share.

{{% examples %}}
## Example Usage
{{% example %}}
### Active Directory Authentication

> **NOTE:** The gateway must have already joined the Active Directory domain prior to SMB file share creationE.g., via "SMB Settings" in the AWS Storage Gateway console or `smb_active_directory_settings` in the `aws.storagegateway.Gateway` resource.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.SmbFileShare("example", {
    authentication: "ActiveDirectory",
    gatewayArn: aws_storagegateway_gateway.example.arn,
    locationArn: aws_s3_bucket.example.arn,
    roleArn: aws_iam_role.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.SmbFileShare("example",
    authentication="ActiveDirectory",
    gateway_arn=aws_storagegateway_gateway["example"]["arn"],
    location_arn=aws_s3_bucket["example"]["arn"],
    role_arn=aws_iam_role["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.SmbFileShare("example", new()
    {
        Authentication = "ActiveDirectory",
        GatewayArn = aws_storagegateway_gateway.Example.Arn,
        LocationArn = aws_s3_bucket.Example.Arn,
        RoleArn = aws_iam_role.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
			Authentication: pulumi.String("ActiveDirectory"),
			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.storagegateway.SmbFileShare;
import com.pulumi.aws.storagegateway.SmbFileShareArgs;
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
        var example = new SmbFileShare("example", SmbFileShareArgs.builder()        
            .authentication("ActiveDirectory")
            .gatewayArn(aws_storagegateway_gateway.example().arn())
            .locationArn(aws_s3_bucket.example().arn())
            .roleArn(aws_iam_role.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:SmbFileShare
    properties:
      authentication: ActiveDirectory
      gatewayArn: ${aws_storagegateway_gateway.example.arn}
      locationArn: ${aws_s3_bucket.example.arn}
      roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% example %}}
### Guest Authentication

> **NOTE:** The gateway must have already had the SMB guest password set prior to SMB file share creationE.g., via "SMB Settings" in the AWS Storage Gateway console or `smb_guest_password` in the `aws.storagegateway.Gateway` resource.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.SmbFileShare("example", {
    authentication: "GuestAccess",
    gatewayArn: aws_storagegateway_gateway.example.arn,
    locationArn: aws_s3_bucket.example.arn,
    roleArn: aws_iam_role.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.SmbFileShare("example",
    authentication="GuestAccess",
    gateway_arn=aws_storagegateway_gateway["example"]["arn"],
    location_arn=aws_s3_bucket["example"]["arn"],
    role_arn=aws_iam_role["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.SmbFileShare("example", new()
    {
        Authentication = "GuestAccess",
        GatewayArn = aws_storagegateway_gateway.Example.Arn,
        LocationArn = aws_s3_bucket.Example.Arn,
        RoleArn = aws_iam_role.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewSmbFileShare(ctx, "example", &storagegateway.SmbFileShareArgs{
			Authentication: pulumi.String("GuestAccess"),
			GatewayArn:     pulumi.Any(aws_storagegateway_gateway.Example.Arn),
			LocationArn:    pulumi.Any(aws_s3_bucket.Example.Arn),
			RoleArn:        pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.storagegateway.SmbFileShare;
import com.pulumi.aws.storagegateway.SmbFileShareArgs;
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
        var example = new SmbFileShare("example", SmbFileShareArgs.builder()        
            .authentication("GuestAccess")
            .gatewayArn(aws_storagegateway_gateway.example().arn())
            .locationArn(aws_s3_bucket.example().arn())
            .roleArn(aws_iam_role.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:SmbFileShare
    properties:
      authentication: GuestAccess
      gatewayArn: ${aws_storagegateway_gateway.example.arn}
      locationArn: ${aws_s3_bucket.example.arn}
      roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_smb_file_share` using the SMB File Share Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/smbFileShare:SmbFileShare example arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
```
 