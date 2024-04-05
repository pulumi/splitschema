Provides a resource to manage an [Amazon Detective Invitation Accepter](https://docs.aws.amazon.com/detective/latest/APIReference/API_AcceptInvitation.html). Ensure that the accepter is configured to use the AWS account you wish to _accept_ the invitation from the primary graph owner account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primaryGraph = new aws.detective.Graph("primaryGraph", {});
const primaryMember = new aws.detective.Member("primaryMember", {
    accountId: "ACCOUNT ID",
    emailAddress: "EMAIL",
    graphArn: primaryGraph.id,
    message: "Message of the invite",
});
const member = new aws.detective.InvitationAccepter("member", {graphArn: primaryGraph.graphArn}, {
    provider: "awsalternate",
    dependsOn: [primaryMember],
});
```
```python
import pulumi
import pulumi_aws as aws

primary_graph = aws.detective.Graph("primaryGraph")
primary_member = aws.detective.Member("primaryMember",
    account_id="ACCOUNT ID",
    email_address="EMAIL",
    graph_arn=primary_graph.id,
    message="Message of the invite")
member = aws.detective.InvitationAccepter("member", graph_arn=primary_graph.graph_arn,
opts=pulumi.ResourceOptions(provider="awsalternate",
    depends_on=[primary_member]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primaryGraph = new Aws.Detective.Graph("primaryGraph");

    var primaryMember = new Aws.Detective.Member("primaryMember", new()
    {
        AccountId = "ACCOUNT ID",
        EmailAddress = "EMAIL",
        GraphArn = primaryGraph.Id,
        Message = "Message of the invite",
    });

    var member = new Aws.Detective.InvitationAccepter("member", new()
    {
        GraphArn = primaryGraph.GraphArn,
    }, new CustomResourceOptions
    {
        Provider = "awsalternate",
        DependsOn = new[]
        {
            primaryMember,
        },
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
		primaryGraph, err := detective.NewGraph(ctx, "primaryGraph", nil)
		if err != nil {
			return err
		}
		primaryMember, err := detective.NewMember(ctx, "primaryMember", &detective.MemberArgs{
			AccountId:    pulumi.String("ACCOUNT ID"),
			EmailAddress: pulumi.String("EMAIL"),
			GraphArn:     primaryGraph.ID(),
			Message:      pulumi.String("Message of the invite"),
		})
		if err != nil {
			return err
		}
		_, err = detective.NewInvitationAccepter(ctx, "member", &detective.InvitationAccepterArgs{
			GraphArn: primaryGraph.GraphArn,
		}, pulumi.Provider("awsalternate"), pulumi.DependsOn([]pulumi.Resource{
			primaryMember,
		}))
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
import com.pulumi.aws.detective.InvitationAccepter;
import com.pulumi.aws.detective.InvitationAccepterArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var primaryGraph = new Graph("primaryGraph");

        var primaryMember = new Member("primaryMember", MemberArgs.builder()        
            .accountId("ACCOUNT ID")
            .emailAddress("EMAIL")
            .graphArn(primaryGraph.id())
            .message("Message of the invite")
            .build());

        var member = new InvitationAccepter("member", InvitationAccepterArgs.builder()        
            .graphArn(primaryGraph.graphArn())
            .build(), CustomResourceOptions.builder()
                .provider("awsalternate")
                .dependsOn(primaryMember)
                .build());

    }
}
```
```yaml
resources:
  primaryGraph:
    type: aws:detective:Graph
  primaryMember:
    type: aws:detective:Member
    properties:
      accountId: ACCOUNT ID
      emailAddress: EMAIL
      graphArn: ${primaryGraph.id}
      message: Message of the invite
  member:
    type: aws:detective:InvitationAccepter
    properties:
      graphArn: ${primaryGraph.graphArn}
    options:
      provider: awsalternate
      dependson:
        - ${primaryMember}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_detective_invitation_accepter` using the graph ARN. For example:

```sh
 $ pulumi import aws:detective/invitationAccepter:InvitationAccepter example arn:aws:detective:us-east-1:123456789101:graph:231684d34gh74g4bae1dbc7bd807d02d
```
 