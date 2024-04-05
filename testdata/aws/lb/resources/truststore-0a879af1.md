Provides a ELBv2 Trust Store for use with Application Load Balancer Listener resources.

{{% examples %}}
## Example Usage
{{% example %}}
### Trust Store Load Balancer Listener

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lb.TrustStore("test", {
    caCertificatesBundleS3Bucket: "...",
    caCertificatesBundleS3Key: "...",
});
const example = new aws.lb.Listener("example", {
    loadBalancerArn: aws_lb.example.id,
    defaultActions: [{
        targetGroupArn: aws_lb_target_group.example.id,
        type: "forward",
    }],
    mutualAuthentication: {
        mode: "verify",
        trustStoreArn: test.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lb.TrustStore("test",
    ca_certificates_bundle_s3_bucket="...",
    ca_certificates_bundle_s3_key="...")
example = aws.lb.Listener("example",
    load_balancer_arn=aws_lb["example"]["id"],
    default_actions=[aws.lb.ListenerDefaultActionArgs(
        target_group_arn=aws_lb_target_group["example"]["id"],
        type="forward",
    )],
    mutual_authentication=aws.lb.ListenerMutualAuthenticationArgs(
        mode="verify",
        trust_store_arn=test.arn,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LB.TrustStore("test", new()
    {
        CaCertificatesBundleS3Bucket = "...",
        CaCertificatesBundleS3Key = "...",
    });

    var example = new Aws.LB.Listener("example", new()
    {
        LoadBalancerArn = aws_lb.Example.Id,
        DefaultActions = new[]
        {
            new Aws.LB.Inputs.ListenerDefaultActionArgs
            {
                TargetGroupArn = aws_lb_target_group.Example.Id,
                Type = "forward",
            },
        },
        MutualAuthentication = new Aws.LB.Inputs.ListenerMutualAuthenticationArgs
        {
            Mode = "verify",
            TrustStoreArn = test.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := lb.NewTrustStore(ctx, "test", &lb.TrustStoreArgs{
			CaCertificatesBundleS3Bucket: pulumi.String("..."),
			CaCertificatesBundleS3Key:    pulumi.String("..."),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListener(ctx, "example", &lb.ListenerArgs{
			LoadBalancerArn: pulumi.Any(aws_lb.Example.Id),
			DefaultActions: lb.ListenerDefaultActionArray{
				&lb.ListenerDefaultActionArgs{
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Example.Id),
					Type:           pulumi.String("forward"),
				},
			},
			MutualAuthentication: &lb.ListenerMutualAuthenticationArgs{
				Mode:          pulumi.String("verify"),
				TrustStoreArn: test.Arn,
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
import com.pulumi.aws.lb.TrustStore;
import com.pulumi.aws.lb.TrustStoreArgs;
import com.pulumi.aws.lb.Listener;
import com.pulumi.aws.lb.ListenerArgs;
import com.pulumi.aws.lb.inputs.ListenerDefaultActionArgs;
import com.pulumi.aws.lb.inputs.ListenerMutualAuthenticationArgs;
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
        var test = new TrustStore("test", TrustStoreArgs.builder()        
            .caCertificatesBundleS3Bucket("...")
            .caCertificatesBundleS3Key("...")
            .build());

        var example = new Listener("example", ListenerArgs.builder()        
            .loadBalancerArn(aws_lb.example().id())
            .defaultActions(ListenerDefaultActionArgs.builder()
                .targetGroupArn(aws_lb_target_group.example().id())
                .type("forward")
                .build())
            .mutualAuthentication(ListenerMutualAuthenticationArgs.builder()
                .mode("verify")
                .trustStoreArn(test.arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lb:TrustStore
    properties:
      caCertificatesBundleS3Bucket: '...'
      caCertificatesBundleS3Key: '...'
  example:
    type: aws:lb:Listener
    properties:
      loadBalancerArn: ${aws_lb.example.id}
      defaultActions:
        - targetGroupArn: ${aws_lb_target_group.example.id}
          type: forward
      mutualAuthentication:
        mode: verify
        trustStoreArn: ${test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Target Groups using their ARN. For example:

```sh
 $ pulumi import aws:lb/trustStore:TrustStore example arn:aws:elasticloadbalancing:us-west-2:187416307283:truststore/my-trust-store/20cfe21448b66314
```
 