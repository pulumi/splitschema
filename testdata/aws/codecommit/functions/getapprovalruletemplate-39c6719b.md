Provides details about a specific CodeCommit Approval Rule Template.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.codecommit.getApprovalRuleTemplate({
    name: "MyExampleApprovalRuleTemplate",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codecommit.get_approval_rule_template(name="MyExampleApprovalRuleTemplate")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CodeCommit.GetApprovalRuleTemplate.Invoke(new()
    {
        Name = "MyExampleApprovalRuleTemplate",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codecommit.LookupApprovalRuleTemplate(ctx, &codecommit.LookupApprovalRuleTemplateArgs{
			Name: "MyExampleApprovalRuleTemplate",
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
import com.pulumi.aws.codecommit.CodecommitFunctions;
import com.pulumi.aws.codecommit.inputs.GetApprovalRuleTemplateArgs;
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
        final var example = CodecommitFunctions.getApprovalRuleTemplate(GetApprovalRuleTemplateArgs.builder()
            .name("MyExampleApprovalRuleTemplate")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:codecommit:getApprovalRuleTemplate
      Arguments:
        name: MyExampleApprovalRuleTemplate
```
{{% /example %}}
{{% /examples %}}