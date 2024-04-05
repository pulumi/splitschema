Provides a CodeCommit Approval Rule Template Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codecommit.ApprovalRuleTemplate("example", {
    description: "This is an example approval rule template",
    content: JSON.stringify({
        Version: "2018-11-08",
        DestinationReferences: ["refs/heads/master"],
        Statements: [{
            Type: "Approvers",
            NumberOfApprovalsNeeded: 2,
            ApprovalPoolMembers: ["arn:aws:sts::123456789012:assumed-role/CodeCommitReview/*"],
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.codecommit.ApprovalRuleTemplate("example",
    description="This is an example approval rule template",
    content=json.dumps({
        "Version": "2018-11-08",
        "DestinationReferences": ["refs/heads/master"],
        "Statements": [{
            "Type": "Approvers",
            "NumberOfApprovalsNeeded": 2,
            "ApprovalPoolMembers": ["arn:aws:sts::123456789012:assumed-role/CodeCommitReview/*"],
        }],
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeCommit.ApprovalRuleTemplate("example", new()
    {
        Description = "This is an example approval rule template",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2018-11-08",
            ["DestinationReferences"] = new[]
            {
                "refs/heads/master",
            },
            ["Statements"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Type"] = "Approvers",
                    ["NumberOfApprovalsNeeded"] = 2,
                    ["ApprovalPoolMembers"] = new[]
                    {
                        "arn:aws:sts::123456789012:assumed-role/CodeCommitReview/*",
                    },
                },
            },
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2018-11-08",
			"DestinationReferences": []string{
				"refs/heads/master",
			},
			"Statements": []map[string]interface{}{
				map[string]interface{}{
					"Type":                    "Approvers",
					"NumberOfApprovalsNeeded": 2,
					"ApprovalPoolMembers": []string{
						"arn:aws:sts::123456789012:assumed-role/CodeCommitReview/*",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = codecommit.NewApprovalRuleTemplate(ctx, "example", &codecommit.ApprovalRuleTemplateArgs{
			Description: pulumi.String("This is an example approval rule template"),
			Content:     pulumi.String(json0),
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
import com.pulumi.aws.codecommit.ApprovalRuleTemplate;
import com.pulumi.aws.codecommit.ApprovalRuleTemplateArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new ApprovalRuleTemplate("example", ApprovalRuleTemplateArgs.builder()        
            .description("This is an example approval rule template")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2018-11-08"),
                    jsonProperty("DestinationReferences", jsonArray("refs/heads/master")),
                    jsonProperty("Statements", jsonArray(jsonObject(
                        jsonProperty("Type", "Approvers"),
                        jsonProperty("NumberOfApprovalsNeeded", 2),
                        jsonProperty("ApprovalPoolMembers", jsonArray("arn:aws:sts::123456789012:assumed-role/CodeCommitReview/*"))
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codecommit:ApprovalRuleTemplate
    properties:
      description: This is an example approval rule template
      content:
        fn::toJSON:
          Version: 2018-11-08
          DestinationReferences:
            - refs/heads/master
          Statements:
            - Type: Approvers
              NumberOfApprovalsNeeded: 2
              ApprovalPoolMembers:
                - arn:aws:sts::123456789012:assumed-role/CodeCommitReview/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeCommit approval rule templates using the `name`. For example:

```sh
 $ pulumi import aws:codecommit/approvalRuleTemplate:ApprovalRuleTemplate imported ExistingApprovalRuleTemplateName
```
 