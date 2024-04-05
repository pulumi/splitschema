Provides an API Gateway Usage Plan Key.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.apigateway.RestApi("test", {});
// ...
const myusageplan = new aws.apigateway.UsagePlan("myusageplan", {apiStages: [{
    apiId: test.id,
    stage: aws_api_gateway_stage.foo.stage_name,
}]});
const mykey = new aws.apigateway.ApiKey("mykey", {});
const main = new aws.apigateway.UsagePlanKey("main", {
    keyId: mykey.id,
    keyType: "API_KEY",
    usagePlanId: myusageplan.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.apigateway.RestApi("test")
# ...
myusageplan = aws.apigateway.UsagePlan("myusageplan", api_stages=[aws.apigateway.UsagePlanApiStageArgs(
    api_id=test.id,
    stage=aws_api_gateway_stage["foo"]["stage_name"],
)])
mykey = aws.apigateway.ApiKey("mykey")
main = aws.apigateway.UsagePlanKey("main",
    key_id=mykey.id,
    key_type="API_KEY",
    usage_plan_id=myusageplan.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.ApiGateway.RestApi("test");

    // ...
    var myusageplan = new Aws.ApiGateway.UsagePlan("myusageplan", new()
    {
        ApiStages = new[]
        {
            new Aws.ApiGateway.Inputs.UsagePlanApiStageArgs
            {
                ApiId = test.Id,
                Stage = aws_api_gateway_stage.Foo.Stage_name,
            },
        },
    });

    var mykey = new Aws.ApiGateway.ApiKey("mykey");

    var main = new Aws.ApiGateway.UsagePlanKey("main", new()
    {
        KeyId = mykey.Id,
        KeyType = "API_KEY",
        UsagePlanId = myusageplan.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := apigateway.NewRestApi(ctx, "test", nil)
		if err != nil {
			return err
		}
		myusageplan, err := apigateway.NewUsagePlan(ctx, "myusageplan", &apigateway.UsagePlanArgs{
			ApiStages: apigateway.UsagePlanApiStageArray{
				&apigateway.UsagePlanApiStageArgs{
					ApiId: test.ID(),
					Stage: pulumi.Any(aws_api_gateway_stage.Foo.Stage_name),
				},
			},
		})
		if err != nil {
			return err
		}
		mykey, err := apigateway.NewApiKey(ctx, "mykey", nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewUsagePlanKey(ctx, "main", &apigateway.UsagePlanKeyArgs{
			KeyId:       mykey.ID(),
			KeyType:     pulumi.String("API_KEY"),
			UsagePlanId: myusageplan.ID(),
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.apigateway.UsagePlan;
import com.pulumi.aws.apigateway.UsagePlanArgs;
import com.pulumi.aws.apigateway.inputs.UsagePlanApiStageArgs;
import com.pulumi.aws.apigateway.ApiKey;
import com.pulumi.aws.apigateway.UsagePlanKey;
import com.pulumi.aws.apigateway.UsagePlanKeyArgs;
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
        var test = new RestApi("test");

        var myusageplan = new UsagePlan("myusageplan", UsagePlanArgs.builder()        
            .apiStages(UsagePlanApiStageArgs.builder()
                .apiId(test.id())
                .stage(aws_api_gateway_stage.foo().stage_name())
                .build())
            .build());

        var mykey = new ApiKey("mykey");

        var main = new UsagePlanKey("main", UsagePlanKeyArgs.builder()        
            .keyId(mykey.id())
            .keyType("API_KEY")
            .usagePlanId(myusageplan.id())
            .build());

    }
}
```
```yaml
resources:
  test: # ...
    type: aws:apigateway:RestApi
  myusageplan:
    type: aws:apigateway:UsagePlan
    properties:
      apiStages:
        - apiId: ${test.id}
          stage: ${aws_api_gateway_stage.foo.stage_name}
  mykey:
    type: aws:apigateway:ApiKey
  main:
    type: aws:apigateway:UsagePlanKey
    properties:
      keyId: ${mykey.id}
      keyType: API_KEY
      usagePlanId: ${myusageplan.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS API Gateway Usage Plan Key using the `USAGE-PLAN-ID/USAGE-PLAN-KEY-ID`. For example:

```sh
 $ pulumi import aws:apigateway/usagePlanKey:UsagePlanKey key 12345abcde/zzz
```
 