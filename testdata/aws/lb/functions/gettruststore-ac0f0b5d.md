> **Note:** `aws_alb_trust_store` is known as `aws.lb.TrustStore`. The functionality is identical.

Provides information about a Load Balancer Trust Store.

This data source can prove useful when a module accepts an LB Trust Store as an
input variable and needs to know its attributes. It can also be used to get the ARN of
an LB Trust Store for use in other resources, given LB Trust Store name.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const lbTsArn = config.get("lbTsArn") || "";
const lbTsName = config.get("lbTsName") || "";
const test = aws.lb.getTrustStore({
    arn: lbTsArn,
    name: lbTsName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_ts_arn = config.get("lbTsArn")
if lb_ts_arn is None:
    lb_ts_arn = ""
lb_ts_name = config.get("lbTsName")
if lb_ts_name is None:
    lb_ts_name = ""
test = aws.lb.get_trust_store(arn=lb_ts_arn,
    name=lb_ts_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var lbTsArn = config.Get("lbTsArn") ?? "";
    var lbTsName = config.Get("lbTsName") ?? "";
    var test = Aws.LB.GetTrustStore.Invoke(new()
    {
        Arn = lbTsArn,
        Name = lbTsName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		lbTsArn := ""
		if param := cfg.Get("lbTsArn"); param != "" {
			lbTsArn = param
		}
		lbTsName := ""
		if param := cfg.Get("lbTsName"); param != "" {
			lbTsName = param
		}
		_, err := lb.LookupTrustStore(ctx, &lb.LookupTrustStoreArgs{
			Arn:  pulumi.StringRef(lbTsArn),
			Name: pulumi.StringRef(lbTsName),
		}, nil)
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
import com.pulumi.aws.lb.LbFunctions;
import com.pulumi.aws.lb.inputs.GetTrustStoreArgs;
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
        final var config = ctx.config();
        final var lbTsArn = config.get("lbTsArn").orElse("");
        final var lbTsName = config.get("lbTsName").orElse("");
        final var test = LbFunctions.getTrustStore(GetTrustStoreArgs.builder()
            .arn(lbTsArn)
            .name(lbTsName)
            .build());

    }
}
```
```yaml
configuration:
  lbTsArn:
    type: string
    default:
  lbTsName:
    type: string
    default:
variables:
  test:
    fn::invoke:
      Function: aws:lb:getTrustStore
      Arguments:
        arn: ${lbTsArn}
        name: ${lbTsName}
```
{{% /example %}}
{{% /examples %}}