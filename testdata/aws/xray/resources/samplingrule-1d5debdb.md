Creates and manages an AWS XRay Sampling Rule.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.xray.SamplingRule("example", {
    attributes: {
        Hello: "Tris",
    },
    fixedRate: 0.05,
    host: "*",
    httpMethod: "*",
    priority: 9999,
    reservoirSize: 1,
    resourceArn: "*",
    ruleName: "example",
    serviceName: "*",
    serviceType: "*",
    urlPath: "*",
    version: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.xray.SamplingRule("example",
    attributes={
        "Hello": "Tris",
    },
    fixed_rate=0.05,
    host="*",
    http_method="*",
    priority=9999,
    reservoir_size=1,
    resource_arn="*",
    rule_name="example",
    service_name="*",
    service_type="*",
    url_path="*",
    version=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Xray.SamplingRule("example", new()
    {
        Attributes = 
        {
            { "Hello", "Tris" },
        },
        FixedRate = 0.05,
        Host = "*",
        HttpMethod = "*",
        Priority = 9999,
        ReservoirSize = 1,
        ResourceArn = "*",
        RuleName = "example",
        ServiceName = "*",
        ServiceType = "*",
        UrlPath = "*",
        Version = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/xray"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := xray.NewSamplingRule(ctx, "example", &xray.SamplingRuleArgs{
			Attributes: pulumi.StringMap{
				"Hello": pulumi.String("Tris"),
			},
			FixedRate:     pulumi.Float64(0.05),
			Host:          pulumi.String("*"),
			HttpMethod:    pulumi.String("*"),
			Priority:      pulumi.Int(9999),
			ReservoirSize: pulumi.Int(1),
			ResourceArn:   pulumi.String("*"),
			RuleName:      pulumi.String("example"),
			ServiceName:   pulumi.String("*"),
			ServiceType:   pulumi.String("*"),
			UrlPath:       pulumi.String("*"),
			Version:       pulumi.Int(1),
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
import com.pulumi.aws.xray.SamplingRule;
import com.pulumi.aws.xray.SamplingRuleArgs;
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
        var example = new SamplingRule("example", SamplingRuleArgs.builder()        
            .attributes(Map.of("Hello", "Tris"))
            .fixedRate(0.05)
            .host("*")
            .httpMethod("*")
            .priority(9999)
            .reservoirSize(1)
            .resourceArn("*")
            .ruleName("example")
            .serviceName("*")
            .serviceType("*")
            .urlPath("*")
            .version(1)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:xray:SamplingRule
    properties:
      attributes:
        Hello: Tris
      fixedRate: 0.05
      host: '*'
      httpMethod: '*'
      priority: 9999
      reservoirSize: 1
      resourceArn: '*'
      ruleName: example
      serviceName: '*'
      serviceType: '*'
      urlPath: '*'
      version: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import XRay Sampling Rules using the name. For example:

```sh
 $ pulumi import aws:xray/samplingRule:SamplingRule example example
```
 