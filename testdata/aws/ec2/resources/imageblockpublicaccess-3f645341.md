Provides a regional public access block for AMIs. This prevents AMIs from being made publicly accessible.
If you already have public AMIs, they will remain publicly available.

> **NOTE:** Deleting this resource does not change the block public access value, the resource in simply removed from state instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Prevent making AMIs publicly accessible in the region and account for which the provider is configured
const test = new aws.ec2.ImageBlockPublicAccess("test", {state: "block-new-sharing"});
```
```python
import pulumi
import pulumi_aws as aws

# Prevent making AMIs publicly accessible in the region and account for which the provider is configured
test = aws.ec2.ImageBlockPublicAccess("test", state="block-new-sharing")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Prevent making AMIs publicly accessible in the region and account for which the provider is configured
    var test = new Aws.Ec2.ImageBlockPublicAccess("test", new()
    {
        State = "block-new-sharing",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewImageBlockPublicAccess(ctx, "test", &ec2.ImageBlockPublicAccessArgs{
			State: pulumi.String("block-new-sharing"),
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
import com.pulumi.aws.ec2.ImageBlockPublicAccess;
import com.pulumi.aws.ec2.ImageBlockPublicAccessArgs;
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
        var test = new ImageBlockPublicAccess("test", ImageBlockPublicAccessArgs.builder()        
            .state("block-new-sharing")
            .build());

    }
}
```
```yaml
resources:
  # Prevent making AMIs publicly accessible in the region and account for which the provider is configured
  test:
    type: aws:ec2:ImageBlockPublicAccess
    properties:
      state: block-new-sharing
```
{{% /example %}}
{{% /examples %}}

## Import

You cannot import this resource. 