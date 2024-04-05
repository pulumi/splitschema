Manages an AWS CloudFront Origin Access Control, which is used by CloudFront Distributions with an Amazon S3 bucket as the origin.

Read more about Origin Access Control in the [CloudFront Developer Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudfront.OriginAccessControl("example", {
    description: "Example Policy",
    originAccessControlOriginType: "s3",
    signingBehavior: "always",
    signingProtocol: "sigv4",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.OriginAccessControl("example",
    description="Example Policy",
    origin_access_control_origin_type="s3",
    signing_behavior="always",
    signing_protocol="sigv4")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFront.OriginAccessControl("example", new()
    {
        Description = "Example Policy",
        OriginAccessControlOriginType = "s3",
        SigningBehavior = "always",
        SigningProtocol = "sigv4",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewOriginAccessControl(ctx, "example", &cloudfront.OriginAccessControlArgs{
			Description:                   pulumi.String("Example Policy"),
			OriginAccessControlOriginType: pulumi.String("s3"),
			SigningBehavior:               pulumi.String("always"),
			SigningProtocol:               pulumi.String("sigv4"),
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
import com.pulumi.aws.cloudfront.OriginAccessControl;
import com.pulumi.aws.cloudfront.OriginAccessControlArgs;
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
        var example = new OriginAccessControl("example", OriginAccessControlArgs.builder()        
            .description("Example Policy")
            .originAccessControlOriginType("s3")
            .signingBehavior("always")
            .signingProtocol("sigv4")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudfront:OriginAccessControl
    properties:
      description: Example Policy
      originAccessControlOriginType: s3
      signingBehavior: always
      signingProtocol: sigv4
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront Origin Access Control using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/originAccessControl:OriginAccessControl example E327GJI25M56DG
```
 