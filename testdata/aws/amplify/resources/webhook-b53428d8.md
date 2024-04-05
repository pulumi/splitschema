Provides an Amplify Webhook resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.amplify.App("example", {});
const masterBranch = new aws.amplify.Branch("masterBranch", {
    appId: example.id,
    branchName: "master",
});
const masterWebhook = new aws.amplify.Webhook("masterWebhook", {
    appId: example.id,
    branchName: masterBranch.branchName,
    description: "triggermaster",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.amplify.App("example")
master_branch = aws.amplify.Branch("masterBranch",
    app_id=example.id,
    branch_name="master")
master_webhook = aws.amplify.Webhook("masterWebhook",
    app_id=example.id,
    branch_name=master_branch.branch_name,
    description="triggermaster")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Amplify.App("example");

    var masterBranch = new Aws.Amplify.Branch("masterBranch", new()
    {
        AppId = example.Id,
        BranchName = "master",
    });

    var masterWebhook = new Aws.Amplify.Webhook("masterWebhook", new()
    {
        AppId = example.Id,
        BranchName = masterBranch.BranchName,
        Description = "triggermaster",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := amplify.NewApp(ctx, "example", nil)
		if err != nil {
			return err
		}
		masterBranch, err := amplify.NewBranch(ctx, "masterBranch", &amplify.BranchArgs{
			AppId:      example.ID(),
			BranchName: pulumi.String("master"),
		})
		if err != nil {
			return err
		}
		_, err = amplify.NewWebhook(ctx, "masterWebhook", &amplify.WebhookArgs{
			AppId:       example.ID(),
			BranchName:  masterBranch.BranchName,
			Description: pulumi.String("triggermaster"),
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
import com.pulumi.aws.amplify.App;
import com.pulumi.aws.amplify.Branch;
import com.pulumi.aws.amplify.BranchArgs;
import com.pulumi.aws.amplify.Webhook;
import com.pulumi.aws.amplify.WebhookArgs;
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
        var example = new App("example");

        var masterBranch = new Branch("masterBranch", BranchArgs.builder()        
            .appId(example.id())
            .branchName("master")
            .build());

        var masterWebhook = new Webhook("masterWebhook", WebhookArgs.builder()        
            .appId(example.id())
            .branchName(masterBranch.branchName())
            .description("triggermaster")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:amplify:App
  masterBranch:
    type: aws:amplify:Branch
    properties:
      appId: ${example.id}
      branchName: master
  masterWebhook:
    type: aws:amplify:Webhook
    properties:
      appId: ${example.id}
      branchName: ${masterBranch.branchName}
      description: triggermaster
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amplify webhook using a webhook ID. For example:

```sh
 $ pulumi import aws:amplify/webhook:Webhook master a26b22a0-748b-4b57-b9a0-ae7e601fe4b1
```
 