Manages an Amazon Managed Service for Prometheus (AMP) Rule Group Namespace

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const demoWorkspace = new aws.amp.Workspace("demoWorkspace", {});
const demoRuleGroupNamespace = new aws.amp.RuleGroupNamespace("demoRuleGroupNamespace", {
    workspaceId: demoWorkspace.id,
    data: `groups:
  - name: test
    rules:
    - record: metric:recording_rule
      expr: avg(rate(container_cpu_usage_seconds_total[5m]))
`,
});
```
```python
import pulumi
import pulumi_aws as aws

demo_workspace = aws.amp.Workspace("demoWorkspace")
demo_rule_group_namespace = aws.amp.RuleGroupNamespace("demoRuleGroupNamespace",
    workspace_id=demo_workspace.id,
    data="""groups:
  - name: test
    rules:
    - record: metric:recording_rule
      expr: avg(rate(container_cpu_usage_seconds_total[5m]))
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var demoWorkspace = new Aws.Amp.Workspace("demoWorkspace");

    var demoRuleGroupNamespace = new Aws.Amp.RuleGroupNamespace("demoRuleGroupNamespace", new()
    {
        WorkspaceId = demoWorkspace.Id,
        Data = @"groups:
  - name: test
    rules:
    - record: metric:recording_rule
      expr: avg(rate(container_cpu_usage_seconds_total[5m]))
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		demoWorkspace, err := amp.NewWorkspace(ctx, "demoWorkspace", nil)
		if err != nil {
			return err
		}
		_, err = amp.NewRuleGroupNamespace(ctx, "demoRuleGroupNamespace", &amp.RuleGroupNamespaceArgs{
			WorkspaceId: demoWorkspace.ID(),
			Data: pulumi.String(`groups:
  - name: test
    rules:
    - record: metric:recording_rule
      expr: avg(rate(container_cpu_usage_seconds_total[5m]))
`),
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
import com.pulumi.aws.amp.Workspace;
import com.pulumi.aws.amp.RuleGroupNamespace;
import com.pulumi.aws.amp.RuleGroupNamespaceArgs;
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
        var demoWorkspace = new Workspace("demoWorkspace");

        var demoRuleGroupNamespace = new RuleGroupNamespace("demoRuleGroupNamespace", RuleGroupNamespaceArgs.builder()        
            .workspaceId(demoWorkspace.id())
            .data("""
groups:
  - name: test
    rules:
    - record: metric:recording_rule
      expr: avg(rate(container_cpu_usage_seconds_total[5m]))
            """)
            .build());

    }
}
```
```yaml
resources:
  demoWorkspace:
    type: aws:amp:Workspace
  demoRuleGroupNamespace:
    type: aws:amp:RuleGroupNamespace
    properties:
      workspaceId: ${demoWorkspace.id}
      data: |
        groups:
          - name: test
            rules:
            - record: metric:recording_rule
              expr: avg(rate(container_cpu_usage_seconds_total[5m]))
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import the prometheus rule group namespace using the arn. For example:

```sh
 $ pulumi import aws:amp/ruleGroupNamespace:RuleGroupNamespace demo arn:aws:aps:us-west-2:123456789012:rulegroupsnamespace/IDstring/namespace_name
```
 