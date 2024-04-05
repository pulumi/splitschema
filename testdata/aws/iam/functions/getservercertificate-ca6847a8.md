Use this data source to lookup information about IAM Server Certificates.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const my-domain = aws.iam.getServerCertificate({
    namePrefix: "my-domain.org",
    latest: true,
});
const elb = new aws.elb.LoadBalancer("elb", {listeners: [{
    instancePort: 8000,
    instanceProtocol: "https",
    lbPort: 443,
    lbProtocol: "https",
    sslCertificateId: my_domain.then(my_domain => my_domain.arn),
}]});
```
```python
import pulumi
import pulumi_aws as aws

my_domain = aws.iam.get_server_certificate(name_prefix="my-domain.org",
    latest=True)
elb = aws.elb.LoadBalancer("elb", listeners=[aws.elb.LoadBalancerListenerArgs(
    instance_port=8000,
    instance_protocol="https",
    lb_port=443,
    lb_protocol="https",
    ssl_certificate_id=my_domain.arn,
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var my_domain = Aws.Iam.GetServerCertificate.Invoke(new()
    {
        NamePrefix = "my-domain.org",
        Latest = true,
    });

    var elb = new Aws.Elb.LoadBalancer("elb", new()
    {
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 8000,
                InstanceProtocol = "https",
                LbPort = 443,
                LbProtocol = "https",
                SslCertificateId = my_domain.Apply(my_domain => my_domain.Apply(getServerCertificateResult => getServerCertificateResult.Arn)),
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		my_domain, err := iam.LookupServerCertificate(ctx, &iam.LookupServerCertificateArgs{
			NamePrefix: pulumi.StringRef("my-domain.org"),
			Latest:     pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancer(ctx, "elb", &elb.LoadBalancerArgs{
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("https"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: *pulumi.String(my_domain.Arn),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetServerCertificateArgs;
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
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
        final var my-domain = IamFunctions.getServerCertificate(GetServerCertificateArgs.builder()
            .namePrefix("my-domain.org")
            .latest(true)
            .build());

        var elb = new LoadBalancer("elb", LoadBalancerArgs.builder()        
            .listeners(LoadBalancerListenerArgs.builder()
                .instancePort(8000)
                .instanceProtocol("https")
                .lbPort(443)
                .lbProtocol("https")
                .sslCertificateId(my_domain.arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  elb:
    type: aws:elb:LoadBalancer
    properties:
      listeners:
        - instancePort: 8000
          instanceProtocol: https
          lbPort: 443
          lbProtocol: https
          sslCertificateId: ${["my-domain"].arn}
variables:
  my-domain:
    fn::invoke:
      Function: aws:iam:getServerCertificate
      Arguments:
        namePrefix: my-domain.org
        latest: true
```
{{% /example %}}
{{% /examples %}}