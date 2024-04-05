Manages an AWS Config Aggregate Authorization

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cfg.AggregateAuthorization("example", {
    accountId: "123456789012",
    region: "eu-west-2",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cfg.AggregateAuthorization("example",
    account_id="123456789012",
    region="eu-west-2")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cfg.AggregateAuthorization("example", new()
    {
        AccountId = "123456789012",
        Region = "eu-west-2",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewAggregateAuthorization(ctx, "example", &cfg.AggregateAuthorizationArgs{
			AccountId: pulumi.String("123456789012"),
			Region:    pulumi.String("eu-west-2"),
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
import com.pulumi.aws.cfg.AggregateAuthorization;
import com.pulumi.aws.cfg.AggregateAuthorizationArgs;
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
        var example = new AggregateAuthorization("example", AggregateAuthorizationArgs.builder()        
            .accountId("123456789012")
            .region("eu-west-2")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cfg:AggregateAuthorization
    properties:
      accountId: '123456789012'
      region: eu-west-2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Config aggregate authorizations using `account_id:region`. For example:

```sh
 $ pulumi import aws:cfg/aggregateAuthorization:AggregateAuthorization example 123456789012:us-east-1
```
 