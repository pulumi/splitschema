Provides a resource to manage an [AWS Macie Account](https://docs.aws.amazon.com/macie/latest/APIReference/macie.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.macie2.Account("test", {
    findingPublishingFrequency: "FIFTEEN_MINUTES",
    status: "ENABLED",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.macie2.Account("test",
    finding_publishing_frequency="FIFTEEN_MINUTES",
    status="ENABLED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Macie2.Account("test", new()
    {
        FindingPublishingFrequency = "FIFTEEN_MINUTES",
        Status = "ENABLED",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := macie2.NewAccount(ctx, "test", &macie2.AccountArgs{
			FindingPublishingFrequency: pulumi.String("FIFTEEN_MINUTES"),
			Status:                     pulumi.String("ENABLED"),
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
import com.pulumi.aws.macie2.Account;
import com.pulumi.aws.macie2.AccountArgs;
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
        var test = new Account("test", AccountArgs.builder()        
            .findingPublishingFrequency("FIFTEEN_MINUTES")
            .status("ENABLED")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:macie2:Account
    properties:
      findingPublishingFrequency: FIFTEEN_MINUTES
      status: ENABLED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_account` using the id. For example:

```sh
 $ pulumi import aws:macie2/account:Account example abcd1
```
 