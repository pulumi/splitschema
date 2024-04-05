Provides information about a MQ Broker Instance Offerings.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const empty = aws.mq.getInstanceTypeOfferings({});
const engine = aws.mq.getInstanceTypeOfferings({
    engineType: "ACTIVEMQ",
});
const storage = aws.mq.getInstanceTypeOfferings({
    storageType: "EBS",
});
const instance = aws.mq.getInstanceTypeOfferings({
    hostInstanceType: "mq.m5.large",
});
const all = aws.mq.getInstanceTypeOfferings({
    engineType: "ACTIVEMQ",
    hostInstanceType: "mq.m5.large",
    storageType: "EBS",
});
```
```python
import pulumi
import pulumi_aws as aws

empty = aws.mq.get_instance_type_offerings()
engine = aws.mq.get_instance_type_offerings(engine_type="ACTIVEMQ")
storage = aws.mq.get_instance_type_offerings(storage_type="EBS")
instance = aws.mq.get_instance_type_offerings(host_instance_type="mq.m5.large")
all = aws.mq.get_instance_type_offerings(engine_type="ACTIVEMQ",
    host_instance_type="mq.m5.large",
    storage_type="EBS")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var empty = Aws.Mq.GetInstanceTypeOfferings.Invoke();

    var engine = Aws.Mq.GetInstanceTypeOfferings.Invoke(new()
    {
        EngineType = "ACTIVEMQ",
    });

    var storage = Aws.Mq.GetInstanceTypeOfferings.Invoke(new()
    {
        StorageType = "EBS",
    });

    var instance = Aws.Mq.GetInstanceTypeOfferings.Invoke(new()
    {
        HostInstanceType = "mq.m5.large",
    });

    var all = Aws.Mq.GetInstanceTypeOfferings.Invoke(new()
    {
        EngineType = "ACTIVEMQ",
        HostInstanceType = "mq.m5.large",
        StorageType = "EBS",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mq"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mq.GetInstanceTypeOfferings(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = mq.GetInstanceTypeOfferings(ctx, &mq.GetInstanceTypeOfferingsArgs{
			EngineType: pulumi.StringRef("ACTIVEMQ"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = mq.GetInstanceTypeOfferings(ctx, &mq.GetInstanceTypeOfferingsArgs{
			StorageType: pulumi.StringRef("EBS"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = mq.GetInstanceTypeOfferings(ctx, &mq.GetInstanceTypeOfferingsArgs{
			HostInstanceType: pulumi.StringRef("mq.m5.large"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = mq.GetInstanceTypeOfferings(ctx, &mq.GetInstanceTypeOfferingsArgs{
			EngineType:       pulumi.StringRef("ACTIVEMQ"),
			HostInstanceType: pulumi.StringRef("mq.m5.large"),
			StorageType:      pulumi.StringRef("EBS"),
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
import com.pulumi.aws.mq.inputs.GetInstanceTypeOfferingsArgs;
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
        final var empty = MqFunctions.getInstanceTypeOfferings();

        final var engine = MqFunctions.getInstanceTypeOfferings(GetInstanceTypeOfferingsArgs.builder()
            .engineType("ACTIVEMQ")
            .build());

        final var storage = MqFunctions.getInstanceTypeOfferings(GetInstanceTypeOfferingsArgs.builder()
            .storageType("EBS")
            .build());

        final var instance = MqFunctions.getInstanceTypeOfferings(GetInstanceTypeOfferingsArgs.builder()
            .hostInstanceType("mq.m5.large")
            .build());

        final var all = MqFunctions.getInstanceTypeOfferings(GetInstanceTypeOfferingsArgs.builder()
            .engineType("ACTIVEMQ")
            .hostInstanceType("mq.m5.large")
            .storageType("EBS")
            .build());

    }
}
```
```yaml
variables:
  empty:
    fn::invoke:
      Function: aws:mq:getInstanceTypeOfferings
      Arguments: {}
  engine:
    fn::invoke:
      Function: aws:mq:getInstanceTypeOfferings
      Arguments:
        engineType: ACTIVEMQ
  storage:
    fn::invoke:
      Function: aws:mq:getInstanceTypeOfferings
      Arguments:
        storageType: EBS
  instance:
    fn::invoke:
      Function: aws:mq:getInstanceTypeOfferings
      Arguments:
        hostInstanceType: mq.m5.large
  all:
    fn::invoke:
      Function: aws:mq:getInstanceTypeOfferings
      Arguments:
        engineType: ACTIVEMQ
        hostInstanceType: mq.m5.large
        storageType: EBS
```
{{% /example %}}
{{% /examples %}}