Provides a resource to manage the default customer master key (CMK) that your AWS account uses to encrypt EBS volumes.

Your AWS account has an AWS-managed default CMK that is used for encrypting an EBS volume when no CMK is specified in the API call that creates the volume.
By using the `aws.ebs.DefaultKmsKey` resource, you can specify a customer-managed CMK to use in place of the AWS-managed default CMK.

> **NOTE:** Creating an `aws.ebs.DefaultKmsKey` resource does not enable default EBS encryption. Use the `aws.ebs.EncryptionByDefault` to enable default EBS encryption.

> **NOTE:** Destroying this resource will reset the default CMK to the account's AWS-managed default CMK for EBS.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ebs.DefaultKmsKey("example", {keyArn: aws_kms_key.example.arn});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ebs.DefaultKmsKey("example", key_arn=aws_kms_key["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ebs.DefaultKmsKey("example", new()
    {
        KeyArn = aws_kms_key.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ebs.NewDefaultKmsKey(ctx, "example", &ebs.DefaultKmsKeyArgs{
			KeyArn: pulumi.Any(aws_kms_key.Example.Arn),
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
import com.pulumi.aws.ebs.DefaultKmsKey;
import com.pulumi.aws.ebs.DefaultKmsKeyArgs;
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
        var example = new DefaultKmsKey("example", DefaultKmsKeyArgs.builder()        
            .keyArn(aws_kms_key.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ebs:DefaultKmsKey
    properties:
      keyArn: ${aws_kms_key.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the EBS default KMS CMK using the KMS key ARN. For example:

```sh
 $ pulumi import aws:ebs/defaultKmsKey:DefaultKmsKey example arn:aws:kms:us-east-1:123456789012:key/abcd-1234
```
 