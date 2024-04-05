Get information on an Amazon MSK VPC Connection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.msk.getVpcConnection({
    arn: aws_msk_vpc_connection.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.msk.get_vpc_connection(arn=aws_msk_vpc_connection["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Msk.GetVpcConnection.Invoke(new()
    {
        Arn = aws_msk_vpc_connection.Example.Arn,
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
		_, err := msk.LookupVpcConnection(ctx, &msk.LookupVpcConnectionArgs{
			Arn: aws_msk_vpc_connection.Example.Arn,
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
import com.pulumi.aws.msk.inputs.GetVpcConnectionArgs;
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
        final var example = MskFunctions.getVpcConnection(GetVpcConnectionArgs.builder()
            .arn(aws_msk_vpc_connection.example().arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:msk:getVpcConnection
      Arguments:
        arn: ${aws_msk_vpc_connection.example.arn}
```
{{% /example %}}
{{% /examples %}}