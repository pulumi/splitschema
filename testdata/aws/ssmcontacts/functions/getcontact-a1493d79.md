Data source for managing an AWS SSM Contact.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ssmcontacts.getContact({
    arn: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssmcontacts.get_contact(arn="arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SsmContacts.GetContact.Invoke(new()
    {
        Arn = "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
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
		_, err := ssmcontacts.LookupContact(ctx, &ssmcontacts.LookupContactArgs{
			Arn: "arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias",
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
import com.pulumi.aws.ssmcontacts.inputs.GetContactArgs;
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
        final var example = SsmcontactsFunctions.getContact(GetContactArgs.builder()
            .arn("arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ssmcontacts:getContact
      Arguments:
        arn: arn:aws:ssm-contacts:us-west-2:123456789012:contact/contactalias
```
{{% /example %}}
{{% /examples %}}