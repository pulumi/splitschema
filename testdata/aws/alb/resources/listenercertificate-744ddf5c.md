Provides a Load Balancer Listener Certificate resource.

This resource is for additional certificates and does not replace the default certificate on the listener.

> **Note:** `aws.alb.ListenerCertificate` is known as `aws.lb.ListenerCertificate`. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCertificate = new aws.acm.Certificate("exampleCertificate", {});
// ...
const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
// ...
const frontEndListener = new aws.lb.Listener("frontEndListener", {});
// ...
const exampleListenerCertificate = new aws.lb.ListenerCertificate("exampleListenerCertificate", {
    listenerArn: frontEndListener.arn,
    certificateArn: exampleCertificate.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example_certificate = aws.acm.Certificate("exampleCertificate")
# ...
front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
# ...
front_end_listener = aws.lb.Listener("frontEndListener")
# ...
example_listener_certificate = aws.lb.ListenerCertificate("exampleListenerCertificate",
    listener_arn=front_end_listener.arn,
    certificate_arn=example_certificate.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCertificate = new Aws.Acm.Certificate("exampleCertificate");

    // ...
    var frontEndLoadBalancer = new Aws.LB.LoadBalancer("frontEndLoadBalancer");

    // ...
    var frontEndListener = new Aws.LB.Listener("frontEndListener");

    // ...
    var exampleListenerCertificate = new Aws.LB.ListenerCertificate("exampleListenerCertificate", new()
    {
        ListenerArn = frontEndListener.Arn,
        CertificateArn = exampleCertificate.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCertificate, err := acm.NewCertificate(ctx, "exampleCertificate", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
		if err != nil {
			return err
		}
		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListenerCertificate(ctx, "exampleListenerCertificate", &lb.ListenerCertificateArgs{
			ListenerArn:    frontEndListener.Arn,
			CertificateArn: exampleCertificate.Arn,
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
import com.pulumi.aws.acm.Certificate;
import com.pulumi.aws.lb.LoadBalancer;
import com.pulumi.aws.lb.Listener;
import com.pulumi.aws.lb.ListenerCertificate;
import com.pulumi.aws.lb.ListenerCertificateArgs;
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
        var exampleCertificate = new Certificate("exampleCertificate");

        var frontEndLoadBalancer = new LoadBalancer("frontEndLoadBalancer");

        var frontEndListener = new Listener("frontEndListener");

        var exampleListenerCertificate = new ListenerCertificate("exampleListenerCertificate", ListenerCertificateArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .certificateArn(exampleCertificate.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleCertificate:
    type: aws:acm:Certificate
  frontEndLoadBalancer:
    type: aws:lb:LoadBalancer
  frontEndListener:
    type: aws:lb:Listener
  exampleListenerCertificate:
    type: aws:lb:ListenerCertificate
    properties:
      listenerArn: ${frontEndListener.arn}
      certificateArn: ${exampleCertificate.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Listener Certificates using the listener arn and certificate arn, separated by an underscore (`_`). For example:

```sh
 $ pulumi import aws:alb/listenerCertificate:ListenerCertificate example arn:aws:elasticloadbalancing:us-west-2:123456789012:listener/app/test/8e4497da625e2d8a/9ab28ade35828f96/67b3d2d36dd7c26b_arn:aws:iam::123456789012:server-certificate/tf-acc-test-6453083910015726063
```
 