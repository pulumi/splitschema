Manages an AWS Storage Gateway NFS File Share.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.NfsFileShare("example", {
    clientLists: ["0.0.0.0/0"],
    gatewayArn: aws_storagegateway_gateway.example.arn,
    locationArn: aws_s3_bucket.example.arn,
    roleArn: aws_iam_role.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.NfsFileShare("example",
    client_lists=["0.0.0.0/0"],
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
    var example = new Aws.StorageGateway.NfsFileShare("example", new()
    {
        ClientLists = new[]
        {
            "0.0.0.0/0",
        },
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
		_, err := storagegateway.NewNfsFileShare(ctx, "example", &storagegateway.NfsFileShareArgs{
			ClientLists: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
			GatewayArn:  pulumi.Any(aws_storagegateway_gateway.Example.Arn),
			LocationArn: pulumi.Any(aws_s3_bucket.Example.Arn),
			RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.storagegateway.NfsFileShare;
import com.pulumi.aws.storagegateway.NfsFileShareArgs;
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
        var example = new NfsFileShare("example", NfsFileShareArgs.builder()        
            .clientLists("0.0.0.0/0")
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
    type: aws:storagegateway:NfsFileShare
    properties:
      clientLists:
        - 0.0.0.0/0
      gatewayArn: ${aws_storagegateway_gateway.example.arn}
      locationArn: ${aws_s3_bucket.example.arn}
      roleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_nfs_file_share` using the NFS File Share Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/nfsFileShare:NfsFileShare example arn:aws:storagegateway:us-east-1:123456789012:share/share-12345678
```
 