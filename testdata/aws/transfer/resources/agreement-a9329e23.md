Provides a AWS Transfer AS2 Agreement resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Agreement("example", {
    accessRole: aws_iam_role.test.arn,
    baseDirectory: "/DOC-EXAMPLE-BUCKET/home/mydirectory",
    description: "example",
    localProfileId: aws_transfer_profile.local.profile_id,
    partnerProfileId: aws_transfer_profile.partner.profile_id,
    serverId: aws_transfer_server.test.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Agreement("example",
    access_role=aws_iam_role["test"]["arn"],
    base_directory="/DOC-EXAMPLE-BUCKET/home/mydirectory",
    description="example",
    local_profile_id=aws_transfer_profile["local"]["profile_id"],
    partner_profile_id=aws_transfer_profile["partner"]["profile_id"],
    server_id=aws_transfer_server["test"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Agreement("example", new()
    {
        AccessRole = aws_iam_role.Test.Arn,
        BaseDirectory = "/DOC-EXAMPLE-BUCKET/home/mydirectory",
        Description = "example",
        LocalProfileId = aws_transfer_profile.Local.Profile_id,
        PartnerProfileId = aws_transfer_profile.Partner.Profile_id,
        ServerId = aws_transfer_server.Test.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewAgreement(ctx, "example", &transfer.AgreementArgs{
			AccessRole:       pulumi.Any(aws_iam_role.Test.Arn),
			BaseDirectory:    pulumi.String("/DOC-EXAMPLE-BUCKET/home/mydirectory"),
			Description:      pulumi.String("example"),
			LocalProfileId:   pulumi.Any(aws_transfer_profile.Local.Profile_id),
			PartnerProfileId: pulumi.Any(aws_transfer_profile.Partner.Profile_id),
			ServerId:         pulumi.Any(aws_transfer_server.Test.Id),
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
import com.pulumi.aws.transfer.Agreement;
import com.pulumi.aws.transfer.AgreementArgs;
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
        var example = new Agreement("example", AgreementArgs.builder()        
            .accessRole(aws_iam_role.test().arn())
            .baseDirectory("/DOC-EXAMPLE-BUCKET/home/mydirectory")
            .description("example")
            .localProfileId(aws_transfer_profile.local().profile_id())
            .partnerProfileId(aws_transfer_profile.partner().profile_id())
            .serverId(aws_transfer_server.test().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Agreement
    properties:
      accessRole: ${aws_iam_role.test.arn}
      baseDirectory: /DOC-EXAMPLE-BUCKET/home/mydirectory
      description: example
      localProfileId: ${aws_transfer_profile.local.profile_id}
      partnerProfileId: ${aws_transfer_profile.partner.profile_id}
      serverId: ${aws_transfer_server.test.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer AS2 Agreement using the `server_id/agreement_id`. For example:

```sh
 $ pulumi import aws:transfer/agreement:Agreement example s-4221a88afd5f4362a/a-4221a88afd5f4362a
```
 