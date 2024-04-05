Provides a Route53 traffic policy instance resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.route53.TrafficPolicyInstance("test", {
    hostedZoneId: "Z033120931TAQO548OGJC",
    trafficPolicyId: "b3gb108f-ea6f-45a5-baab-9d112d8b4037",
    trafficPolicyVersion: 1,
    ttl: 360,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.route53.TrafficPolicyInstance("test",
    hosted_zone_id="Z033120931TAQO548OGJC",
    traffic_policy_id="b3gb108f-ea6f-45a5-baab-9d112d8b4037",
    traffic_policy_version=1,
    ttl=360)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Route53.TrafficPolicyInstance("test", new()
    {
        HostedZoneId = "Z033120931TAQO548OGJC",
        TrafficPolicyId = "b3gb108f-ea6f-45a5-baab-9d112d8b4037",
        TrafficPolicyVersion = 1,
        Ttl = 360,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewTrafficPolicyInstance(ctx, "test", &route53.TrafficPolicyInstanceArgs{
			HostedZoneId:         pulumi.String("Z033120931TAQO548OGJC"),
			TrafficPolicyId:      pulumi.String("b3gb108f-ea6f-45a5-baab-9d112d8b4037"),
			TrafficPolicyVersion: pulumi.Int(1),
			Ttl:                  pulumi.Int(360),
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
import com.pulumi.aws.route53.TrafficPolicyInstance;
import com.pulumi.aws.route53.TrafficPolicyInstanceArgs;
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
        var test = new TrafficPolicyInstance("test", TrafficPolicyInstanceArgs.builder()        
            .hostedZoneId("Z033120931TAQO548OGJC")
            .trafficPolicyId("b3gb108f-ea6f-45a5-baab-9d112d8b4037")
            .trafficPolicyVersion(1)
            .ttl(360)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:route53:TrafficPolicyInstance
    properties:
      hostedZoneId: Z033120931TAQO548OGJC
      trafficPolicyId: b3gb108f-ea6f-45a5-baab-9d112d8b4037
      trafficPolicyVersion: 1
      ttl: 360
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 traffic policy instance using its id. For example:

```sh
 $ pulumi import aws:route53/trafficPolicyInstance:TrafficPolicyInstance test df579d9a-6396-410e-ac22-e7ad60cf9e7e
```
 