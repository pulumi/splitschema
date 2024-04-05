Get information on a Amazon MSK Kafka Version

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const preferred = aws.msk.getKafkaVersion({
    preferredVersions: [
        "2.4.1.1",
        "2.4.1",
        "2.2.1",
    ],
});
const example = aws.msk.getKafkaVersion({
    version: "2.8.0",
});
```
```python
import pulumi
import pulumi_aws as aws

preferred = aws.msk.get_kafka_version(preferred_versions=[
    "2.4.1.1",
    "2.4.1",
    "2.2.1",
])
example = aws.msk.get_kafka_version(version="2.8.0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var preferred = Aws.Msk.GetKafkaVersion.Invoke(new()
    {
        PreferredVersions = new[]
        {
            "2.4.1.1",
            "2.4.1",
            "2.2.1",
        },
    });

    var example = Aws.Msk.GetKafkaVersion.Invoke(new()
    {
        Version = "2.8.0",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := msk.GetKafkaVersion(ctx, &msk.GetKafkaVersionArgs{
			PreferredVersions: []string{
				"2.4.1.1",
				"2.4.1",
				"2.2.1",
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = msk.GetKafkaVersion(ctx, &msk.GetKafkaVersionArgs{
			Version: pulumi.StringRef("2.8.0"),
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
import com.pulumi.aws.msk.MskFunctions;
import com.pulumi.aws.msk.inputs.GetKafkaVersionArgs;
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
        final var preferred = MskFunctions.getKafkaVersion(GetKafkaVersionArgs.builder()
            .preferredVersions(            
                "2.4.1.1",
                "2.4.1",
                "2.2.1")
            .build());

        final var example = MskFunctions.getKafkaVersion(GetKafkaVersionArgs.builder()
            .version("2.8.0")
            .build());

    }
}
```
```yaml
variables:
  preferred:
    fn::invoke:
      Function: aws:msk:getKafkaVersion
      Arguments:
        preferredVersions:
          - 2.4.1.1
          - 2.4.1
          - 2.2.1
  example:
    fn::invoke:
      Function: aws:msk:getKafkaVersion
      Arguments:
        version: 2.8.0
```
{{% /example %}}
{{% /examples %}}