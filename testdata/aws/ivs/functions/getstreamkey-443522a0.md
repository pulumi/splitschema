Data source for managing an AWS IVS (Interactive Video) Stream Key.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ivs.getStreamKey({
    channelArn: "arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ivs.get_stream_key(channel_arn="arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ivs.GetStreamKey.Invoke(new()
    {
        ChannelArn = "arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ivs.GetStreamKey(ctx, &ivs.GetStreamKeyArgs{
			ChannelArn: "arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5",
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
import com.pulumi.aws.ivs.IvsFunctions;
import com.pulumi.aws.ivs.inputs.GetStreamKeyArgs;
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
        final var example = IvsFunctions.getStreamKey(GetStreamKeyArgs.builder()
            .channelArn("arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ivs:getStreamKey
      Arguments:
        channelArn: arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5
```
{{% /example %}}
{{% /examples %}}