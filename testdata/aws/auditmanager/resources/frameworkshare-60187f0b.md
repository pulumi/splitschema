Resource for managing an AWS Audit Manager Framework Share.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.auditmanager.FrameworkShare("example", {
    destinationAccount: "012345678901",
    destinationRegion: "us-east-1",
    frameworkId: aws_auditmanager_framework.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.FrameworkShare("example",
    destination_account="012345678901",
    destination_region="us-east-1",
    framework_id=aws_auditmanager_framework["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Auditmanager.FrameworkShare("example", new()
    {
        DestinationAccount = "012345678901",
        DestinationRegion = "us-east-1",
        FrameworkId = aws_auditmanager_framework.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewFrameworkShare(ctx, "example", &auditmanager.FrameworkShareArgs{
			DestinationAccount: pulumi.String("012345678901"),
			DestinationRegion:  pulumi.String("us-east-1"),
			FrameworkId:        pulumi.Any(aws_auditmanager_framework.Example.Id),
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
import com.pulumi.aws.auditmanager.FrameworkShare;
import com.pulumi.aws.auditmanager.FrameworkShareArgs;
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
        var example = new FrameworkShare("example", FrameworkShareArgs.builder()        
            .destinationAccount("012345678901")
            .destinationRegion("us-east-1")
            .frameworkId(aws_auditmanager_framework.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:auditmanager:FrameworkShare
    properties:
      destinationAccount: '012345678901'
      destinationRegion: us-east-1
      frameworkId: ${aws_auditmanager_framework.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Framework Share using the `id`. For example:

```sh
 $ pulumi import aws:auditmanager/frameworkShare:FrameworkShare example abcdef-123456
```
 