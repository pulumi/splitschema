Resource for managing an AWS QuickSight Account Subscription.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const subscription = new aws.quicksight.AccountSubscription("subscription", {
    accountName: "quicksight-pulumi",
    authenticationMethod: "IAM_AND_QUICKSIGHT",
    edition: "ENTERPRISE",
    notificationEmail: "notification@email.com",
});
```
```python
import pulumi
import pulumi_aws as aws

subscription = aws.quicksight.AccountSubscription("subscription",
    account_name="quicksight-pulumi",
    authentication_method="IAM_AND_QUICKSIGHT",
    edition="ENTERPRISE",
    notification_email="notification@email.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var subscription = new Aws.Quicksight.AccountSubscription("subscription", new()
    {
        AccountName = "quicksight-pulumi",
        AuthenticationMethod = "IAM_AND_QUICKSIGHT",
        Edition = "ENTERPRISE",
        NotificationEmail = "notification@email.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewAccountSubscription(ctx, "subscription", &quicksight.AccountSubscriptionArgs{
			AccountName:          pulumi.String("quicksight-pulumi"),
			AuthenticationMethod: pulumi.String("IAM_AND_QUICKSIGHT"),
			Edition:              pulumi.String("ENTERPRISE"),
			NotificationEmail:    pulumi.String("notification@email.com"),
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
import com.pulumi.aws.quicksight.AccountSubscription;
import com.pulumi.aws.quicksight.AccountSubscriptionArgs;
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
        var subscription = new AccountSubscription("subscription", AccountSubscriptionArgs.builder()        
            .accountName("quicksight-pulumi")
            .authenticationMethod("IAM_AND_QUICKSIGHT")
            .edition("ENTERPRISE")
            .notificationEmail("notification@email.com")
            .build());

    }
}
```
```yaml
resources:
  subscription:
    type: aws:quicksight:AccountSubscription
    properties:
      accountName: quicksight-pulumi
      authenticationMethod: IAM_AND_QUICKSIGHT
      edition: ENTERPRISE
      notificationEmail: notification@email.com
```
{{% /example %}}
{{% /examples %}}

## Import

You cannot import this resource. 