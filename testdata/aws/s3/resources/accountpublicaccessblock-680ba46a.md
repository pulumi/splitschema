Manages S3 account-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

> **NOTE:** Each AWS account may only have one S3 Public Access Block configuration. Multiple configurations of the resource against the same AWS account will cause a perpetual difference.

> Advanced usage: To use a custom API endpoint for this resource, use the `s3control` endpoint provider configuration, not the `s3` endpoint provider configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.AccountPublicAccessBlock("example", {
    blockPublicAcls: true,
    blockPublicPolicy: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.AccountPublicAccessBlock("example",
    block_public_acls=True,
    block_public_policy=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.AccountPublicAccessBlock("example", new()
    {
        BlockPublicAcls = true,
        BlockPublicPolicy = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewAccountPublicAccessBlock(ctx, "example", &s3.AccountPublicAccessBlockArgs{
			BlockPublicAcls:   pulumi.Bool(true),
			BlockPublicPolicy: pulumi.Bool(true),
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
import com.pulumi.aws.s3.AccountPublicAccessBlock;
import com.pulumi.aws.s3.AccountPublicAccessBlockArgs;
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
        var example = new AccountPublicAccessBlock("example", AccountPublicAccessBlockArgs.builder()        
            .blockPublicAcls(true)
            .blockPublicPolicy(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:AccountPublicAccessBlock
    properties:
      blockPublicAcls: true
      blockPublicPolicy: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_s3_account_public_access_block` using the AWS account ID. For example:

```sh
 $ pulumi import aws:s3/accountPublicAccessBlock:AccountPublicAccessBlock example 123456789012
```
 