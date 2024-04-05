Associates a CodeCommit Approval Rule Template with a Repository.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codecommit.ApprovalRuleTemplateAssociation("example", {
    approvalRuleTemplateName: aws_codecommit_approval_rule_template.example.name,
    repositoryName: aws_codecommit_repository.example.repository_name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codecommit.ApprovalRuleTemplateAssociation("example",
    approval_rule_template_name=aws_codecommit_approval_rule_template["example"]["name"],
    repository_name=aws_codecommit_repository["example"]["repository_name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeCommit.ApprovalRuleTemplateAssociation("example", new()
    {
        ApprovalRuleTemplateName = aws_codecommit_approval_rule_template.Example.Name,
        RepositoryName = aws_codecommit_repository.Example.Repository_name,
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
		_, err := codecommit.NewApprovalRuleTemplateAssociation(ctx, "example", &codecommit.ApprovalRuleTemplateAssociationArgs{
			ApprovalRuleTemplateName: pulumi.Any(aws_codecommit_approval_rule_template.Example.Name),
			RepositoryName:           pulumi.Any(aws_codecommit_repository.Example.Repository_name),
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
import com.pulumi.aws.codecommit.ApprovalRuleTemplateAssociation;
import com.pulumi.aws.codecommit.ApprovalRuleTemplateAssociationArgs;
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
        var example = new ApprovalRuleTemplateAssociation("example", ApprovalRuleTemplateAssociationArgs.builder()        
            .approvalRuleTemplateName(aws_codecommit_approval_rule_template.example().name())
            .repositoryName(aws_codecommit_repository.example().repository_name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codecommit:ApprovalRuleTemplateAssociation
    properties:
      approvalRuleTemplateName: ${aws_codecommit_approval_rule_template.example.name}
      repositoryName: ${aws_codecommit_repository.example.repository_name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeCommit approval rule template associations using the `approval_rule_template_name` and `repository_name` separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:codecommit/approvalRuleTemplateAssociation:ApprovalRuleTemplateAssociation example approver-rule-for-example,MyExampleRepo
```
 