Data source for managing a Plan of an AWS SSM Contact.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ssmcontacts.getPlan({
    contactId: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ssmcontacts.get_plan(contact_id="arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.SsmContacts.GetPlan.Invoke(new()
    {
        ContactId = "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmcontacts"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmcontacts.LookupPlan(ctx, &ssmcontacts.LookupPlanArgs{
			ContactId: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
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
import com.pulumi.aws.ssmcontacts.SsmcontactsFunctions;
import com.pulumi.aws.ssmcontacts.inputs.GetPlanArgs;
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
        final var test = SsmcontactsFunctions.getPlan(GetPlanArgs.builder()
            .contactId("arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ssmcontacts:getPlan
      Arguments:
        contactId: arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias
```
{{% /example %}}
{{% /examples %}}