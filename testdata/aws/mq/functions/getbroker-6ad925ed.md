Provides information about a MQ Broker.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const brokerId = config.get("brokerId") || "";
const brokerName = config.get("brokerName") || "";
const byId = aws.mq.getBroker({
    brokerId: brokerId,
});
const byName = aws.mq.getBroker({
    brokerName: brokerName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
broker_id = config.get("brokerId")
if broker_id is None:
    broker_id = ""
broker_name = config.get("brokerName")
if broker_name is None:
    broker_name = ""
by_id = aws.mq.get_broker(broker_id=broker_id)
by_name = aws.mq.get_broker(broker_name=broker_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var brokerId = config.Get("brokerId") ?? "";
    var brokerName = config.Get("brokerName") ?? "";
    var byId = Aws.Mq.GetBroker.Invoke(new()
    {
        BrokerId = brokerId,
    });

    var byName = Aws.Mq.GetBroker.Invoke(new()
    {
        BrokerName = brokerName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		brokerId := ""
		if param := cfg.Get("brokerId"); param != "" {
			brokerId = param
		}
		brokerName := ""
		if param := cfg.Get("brokerName"); param != "" {
			brokerName = param
		}
		_, err := mq.LookupBroker(ctx, &mq.LookupBrokerArgs{
			BrokerId: pulumi.StringRef(brokerId),
		}, nil)
		if err != nil {
			return err
		}
		_, err = mq.LookupBroker(ctx, &mq.LookupBrokerArgs{
			BrokerName: pulumi.StringRef(brokerName),
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
import com.pulumi.aws.mq.MqFunctions;
import com.pulumi.aws.mq.inputs.GetBrokerArgs;
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
        final var brokerId = config.get("brokerId").orElse("");
        final var brokerName = config.get("brokerName").orElse("");
        final var byId = MqFunctions.getBroker(GetBrokerArgs.builder()
            .brokerId(brokerId)
            .build());

        final var byName = MqFunctions.getBroker(GetBrokerArgs.builder()
            .brokerName(brokerName)
            .build());

    }
}
```
```yaml
configuration:
  brokerId:
    type: string
    default:
  brokerName:
    type: string
    default:
variables:
  byId:
    fn::invoke:
      Function: aws:mq:getBroker
      Arguments:
        brokerId: ${brokerId}
  byName:
    fn::invoke:
      Function: aws:mq:getBroker
      Arguments:
        brokerName: ${brokerName}
```
{{% /example %}}
{{% /examples %}}