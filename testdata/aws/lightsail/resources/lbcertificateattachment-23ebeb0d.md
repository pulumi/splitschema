Attaches a Lightsail Load Balancer Certificate to a Lightsail Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLb = new aws.lightsail.Lb("testLb", {
    healthCheckPath: "/",
    instancePort: 80,
    tags: {
        foo: "bar",
    },
});
const testLbCertificate = new aws.lightsail.LbCertificate("testLbCertificate", {
    lbName: testLb.id,
    domainName: "test.com",
});
const testLbCertificateAttachment = new aws.lightsail.LbCertificateAttachment("testLbCertificateAttachment", {
    lbName: testLb.name,
    certificateName: testLbCertificate.name,
});
```
```python
import pulumi
import pulumi_aws as aws

test_lb = aws.lightsail.Lb("testLb",
    health_check_path="/",
    instance_port=80,
    tags={
        "foo": "bar",
    })
test_lb_certificate = aws.lightsail.LbCertificate("testLbCertificate",
    lb_name=test_lb.id,
    domain_name="test.com")
test_lb_certificate_attachment = aws.lightsail.LbCertificateAttachment("testLbCertificateAttachment",
    lb_name=test_lb.name,
    certificate_name=test_lb_certificate.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testLb = new Aws.LightSail.Lb("testLb", new()
    {
        HealthCheckPath = "/",
        InstancePort = 80,
        Tags = 
        {
            { "foo", "bar" },
        },
    });

    var testLbCertificate = new Aws.LightSail.LbCertificate("testLbCertificate", new()
    {
        LbName = testLb.Id,
        DomainName = "test.com",
    });

    var testLbCertificateAttachment = new Aws.LightSail.LbCertificateAttachment("testLbCertificateAttachment", new()
    {
        LbName = testLb.Name,
        CertificateName = testLbCertificate.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testLb, err := lightsail.NewLb(ctx, "testLb", &lightsail.LbArgs{
			HealthCheckPath: pulumi.String("/"),
			InstancePort:    pulumi.Int(80),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
		})
		if err != nil {
			return err
		}
		testLbCertificate, err := lightsail.NewLbCertificate(ctx, "testLbCertificate", &lightsail.LbCertificateArgs{
			LbName:     testLb.ID(),
			DomainName: pulumi.String("test.com"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewLbCertificateAttachment(ctx, "testLbCertificateAttachment", &lightsail.LbCertificateAttachmentArgs{
			LbName:          testLb.Name,
			CertificateName: testLbCertificate.Name,
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
import com.pulumi.aws.lightsail.Lb;
import com.pulumi.aws.lightsail.LbArgs;
import com.pulumi.aws.lightsail.LbCertificate;
import com.pulumi.aws.lightsail.LbCertificateArgs;
import com.pulumi.aws.lightsail.LbCertificateAttachment;
import com.pulumi.aws.lightsail.LbCertificateAttachmentArgs;
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
        var testLb = new Lb("testLb", LbArgs.builder()        
            .healthCheckPath("/")
            .instancePort("80")
            .tags(Map.of("foo", "bar"))
            .build());

        var testLbCertificate = new LbCertificate("testLbCertificate", LbCertificateArgs.builder()        
            .lbName(testLb.id())
            .domainName("test.com")
            .build());

        var testLbCertificateAttachment = new LbCertificateAttachment("testLbCertificateAttachment", LbCertificateAttachmentArgs.builder()        
            .lbName(testLb.name())
            .certificateName(testLbCertificate.name())
            .build());

    }
}
```
```yaml
resources:
  testLb:
    type: aws:lightsail:Lb
    properties:
      healthCheckPath: /
      instancePort: '80'
      tags:
        foo: bar
  testLbCertificate:
    type: aws:lightsail:LbCertificate
    properties:
      lbName: ${testLb.id}
      domainName: test.com
  testLbCertificateAttachment:
    type: aws:lightsail:LbCertificateAttachment
    properties:
      lbName: ${testLb.name}
      certificateName: ${testLbCertificate.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_lightsail_lb_certificate_attachment` using the name attribute. For example:

```sh
 $ pulumi import aws:lightsail/lbCertificateAttachment:LbCertificateAttachment test example-load-balancer,example-certificate
```
 