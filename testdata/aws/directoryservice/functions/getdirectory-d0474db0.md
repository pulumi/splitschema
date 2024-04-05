Get attributes of AWS Directory Service directory (SimpleAD, Managed AD, AD Connector). It's especially useful to refer AWS Managed AD or on-premise AD in AD Connector configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.directoryservice.getDirectory({
    directoryId: aws_directory_service_directory.main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directoryservice.get_directory(directory_id=aws_directory_service_directory["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.DirectoryService.GetDirectory.Invoke(new()
    {
        DirectoryId = aws_directory_service_directory.Main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directoryservice.LookupDirectory(ctx, &directoryservice.LookupDirectoryArgs{
			DirectoryId: aws_directory_service_directory.Main.Id,
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
import com.pulumi.aws.directoryservice.DirectoryserviceFunctions;
import com.pulumi.aws.directoryservice.inputs.GetDirectoryArgs;
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
        final var example = DirectoryserviceFunctions.getDirectory(GetDirectoryArgs.builder()
            .directoryId(aws_directory_service_directory.main().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:directoryservice:getDirectory
      Arguments:
        directoryId: ${aws_directory_service_directory.main.id}
```
{{% /example %}}
{{% /examples %}}