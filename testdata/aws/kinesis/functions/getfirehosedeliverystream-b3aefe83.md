Use this data source to get information about a Kinesis Firehose Delivery Stream for use in other resources.

For more details, see the [Amazon Kinesis Firehose Documentation](https://aws.amazon.com/documentation/firehose/).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const stream = aws.kinesis.getFirehoseDeliveryStream({
    name: "stream-name",
});
```
```python
import pulumi
import pulumi_aws as aws

stream = aws.kinesis.get_firehose_delivery_stream(name="stream-name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var stream = Aws.Kinesis.GetFirehoseDeliveryStream.Invoke(new()
    {
        Name = "stream-name",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kinesis.LookupFirehoseDeliveryStream(ctx, &kinesis.LookupFirehoseDeliveryStreamArgs{
			Name: "stream-name",
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
import com.pulumi.aws.kinesis.KinesisFunctions;
import com.pulumi.aws.kinesis.inputs.GetFirehoseDeliveryStreamArgs;
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
        final var stream = KinesisFunctions.getFirehoseDeliveryStream(GetFirehoseDeliveryStreamArgs.builder()
            .name("stream-name")
            .build());

    }
}
```
```yaml
variables:
  stream:
    fn::invoke:
      Function: aws:kinesis:getFirehoseDeliveryStream
      Arguments:
        name: stream-name
```
{{% /example %}}
{{% /examples %}}