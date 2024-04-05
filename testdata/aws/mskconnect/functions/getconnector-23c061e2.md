Get information on an Amazon MSK Connect Connector.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.mskconnect.getConnector({
    name: "example-mskconnector",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mskconnect.get_connector(name="example-mskconnector")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.MskConnect.GetConnector.Invoke(new()
    {
        Name = "example-mskconnector",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mskconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := mskconnect.LookupConnector(ctx, &mskconnect.LookupConnectorArgs{
			Name: "example-mskconnector",
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
import com.pulumi.aws.mskconnect.MskconnectFunctions;
import com.pulumi.aws.mskconnect.inputs.GetConnectorArgs;
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
        final var example = MskconnectFunctions.getConnector(GetConnectorArgs.builder()
            .name("example-mskconnector")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:mskconnect:getConnector
      Arguments:
        name: example-mskconnector
```
{{% /example %}}
{{% /examples %}}