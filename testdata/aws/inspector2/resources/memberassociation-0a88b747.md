Resource for associating accounts to existing Inspector instances.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.inspector2.MemberAssociation("example", {accountId: "123456789012"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.inspector2.MemberAssociation("example", account_id="123456789012")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Inspector2.MemberAssociation("example", new()
    {
        AccountId = "123456789012",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/inspector2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := inspector2.NewMemberAssociation(ctx, "example", &inspector2.MemberAssociationArgs{
			AccountId: pulumi.String("123456789012"),
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
import com.pulumi.aws.inspector2.MemberAssociation;
import com.pulumi.aws.inspector2.MemberAssociationArgs;
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
        var example = new MemberAssociation("example", MemberAssociationArgs.builder()        
            .accountId("123456789012")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:inspector2:MemberAssociation
    properties:
      accountId: '123456789012'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Inspector Member Association using the `account_id`. For example:

```sh
 $ pulumi import aws:inspector2/memberAssociation:MemberAssociation example 123456789012
```
 