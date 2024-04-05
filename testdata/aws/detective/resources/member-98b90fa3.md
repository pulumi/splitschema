Provides a resource to manage an [Amazon Detective Member](https://docs.aws.amazon.com/detective/latest/APIReference/API_CreateMembers.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGraph = new aws.detective.Graph("exampleGraph", {});
const exampleMember = new aws.detective.Member("exampleMember", {
    accountId: "AWS ACCOUNT ID",
    emailAddress: "EMAIL",
    graphArn: exampleGraph.id,
    message: "Message of the invitation",
    disableEmailNotification: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example_graph = aws.detective.Graph("exampleGraph")
example_member = aws.detective.Member("exampleMember",
    account_id="AWS ACCOUNT ID",
    email_address="EMAIL",
    graph_arn=example_graph.id,
    message="Message of the invitation",
    disable_email_notification=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGraph = new Aws.Detective.Graph("exampleGraph");

    var exampleMember = new Aws.Detective.Member("exampleMember", new()
    {
        AccountId = "AWS ACCOUNT ID",
        EmailAddress = "EMAIL",
        GraphArn = exampleGraph.Id,
        Message = "Message of the invitation",
        DisableEmailNotification = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/detective"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleGraph, err := detective.NewGraph(ctx, "exampleGraph", nil)
		if err != nil {
			return err
		}
		_, err = detective.NewMember(ctx, "exampleMember", &detective.MemberArgs{
			AccountId:                pulumi.String("AWS ACCOUNT ID"),
			EmailAddress:             pulumi.String("EMAIL"),
			GraphArn:                 exampleGraph.ID(),
			Message:                  pulumi.String("Message of the invitation"),
			DisableEmailNotification: pulumi.Bool(true),
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
import com.pulumi.aws.detective.Graph;
import com.pulumi.aws.detective.Member;
import com.pulumi.aws.detective.MemberArgs;
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
        var exampleGraph = new Graph("exampleGraph");

        var exampleMember = new Member("exampleMember", MemberArgs.builder()        
            .accountId("AWS ACCOUNT ID")
            .emailAddress("EMAIL")
            .graphArn(exampleGraph.id())
            .message("Message of the invitation")
            .disableEmailNotification(true)
            .build());

    }
}
```
```yaml
resources:
  exampleGraph:
    type: aws:detective:Graph
  exampleMember:
    type: aws:detective:Member
    properties:
      accountId: AWS ACCOUNT ID
      emailAddress: EMAIL
      graphArn: ${exampleGraph.id}
      message: Message of the invitation
      disableEmailNotification: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_detective_member` using the ARN of the graph followed by the account ID of the member account. For example:

```sh
 $ pulumi import aws:detective/member:Member example arn:aws:detective:us-east-1:123456789101:graph:231684d34gh74g4bae1dbc7bd807d02d/123456789012
```
 