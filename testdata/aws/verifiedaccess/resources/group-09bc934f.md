Resource for managing a Verified Access Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.verifiedaccess.Group("example", {verifiedaccessInstanceId: aws_verifiedaccess_instance.example.id});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.verifiedaccess.Group("example", verifiedaccess_instance_id=aws_verifiedaccess_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VerifiedAccess.Group("example", new()
    {
        VerifiedaccessInstanceId = aws_verifiedaccess_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := verifiedaccess.NewGroup(ctx, "example", &verifiedaccess.GroupArgs{
			VerifiedaccessInstanceId: pulumi.Any(aws_verifiedaccess_instance.Example.Id),
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
import com.pulumi.aws.verifiedaccess.Group;
import com.pulumi.aws.verifiedaccess.GroupArgs;
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
        var example = new Group("example", GroupArgs.builder()        
            .verifiedaccessInstanceId(aws_verifiedaccess_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:verifiedaccess:Group
    properties:
      verifiedaccessInstanceId: ${aws_verifiedaccess_instance.example.id}
```
{{% /example %}}
{{% /examples %}}