Provides an AppSync Domain Name.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appsync.DomainName("example", {
    domainName: "api.example.com",
    certificateArn: aws_acm_certificate.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appsync.DomainName("example",
    domain_name="api.example.com",
    certificate_arn=aws_acm_certificate["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppSync.DomainName("example", new()
    {
        Name = "api.example.com",
        CertificateArn = aws_acm_certificate.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsync.NewDomainName(ctx, "example", &appsync.DomainNameArgs{
			DomainName:     pulumi.String("api.example.com"),
			CertificateArn: pulumi.Any(aws_acm_certificate.Example.Arn),
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
import com.pulumi.aws.appsync.DomainName;
import com.pulumi.aws.appsync.DomainNameArgs;
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
        var example = new DomainName("example", DomainNameArgs.builder()        
            .domainName("api.example.com")
            .certificateArn(aws_acm_certificate.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appsync:DomainName
    properties:
      domainName: api.example.com
      certificateArn: ${aws_acm_certificate.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appsync_domain_name` using the AppSync domain name. For example:

```sh
 $ pulumi import aws:appsync/domainName:DomainName example example.com
```
 