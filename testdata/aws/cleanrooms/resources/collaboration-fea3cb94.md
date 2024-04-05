Provides a AWS Clean Rooms collaboration.  All members included in the definition will be invited to
join the collaboration and can create memberships.

{{% examples %}}
## Example Usage
{{% example %}}
### Collaboration with tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCollaboration = new aws.cleanrooms.Collaboration("testCollaboration", {
    creatorDisplayName: "Creator ",
    creatorMemberAbilities: [
        "CAN_QUERY",
        "CAN_RECEIVE_RESULTS",
    ],
    dataEncryptionMetadata: {
        allowClearText: true,
        allowDuplicates: true,
        allowJoinsOnColumnsWithDifferentNames: true,
        preserveNulls: false,
    },
    description: "I made this collaboration with Pulumi!",
    members: [{
        accountId: "123456789012",
        displayName: "Other member",
        memberAbilities: [],
    }],
    queryLogStatus: "DISABLED",
    tags: {
        Project: "Pulumi",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_collaboration = aws.cleanrooms.Collaboration("testCollaboration",
    creator_display_name="Creator ",
    creator_member_abilities=[
        "CAN_QUERY",
        "CAN_RECEIVE_RESULTS",
    ],
    data_encryption_metadata=aws.cleanrooms.CollaborationDataEncryptionMetadataArgs(
        allow_clear_text=True,
        allow_duplicates=True,
        allow_joins_on_columns_with_different_names=True,
        preserve_nulls=False,
    ),
    description="I made this collaboration with Pulumi!",
    members=[aws.cleanrooms.CollaborationMemberArgs(
        account_id="123456789012",
        display_name="Other member",
        member_abilities=[],
    )],
    query_log_status="DISABLED",
    tags={
        "Project": "Pulumi",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCollaboration = new Aws.CleanRooms.Collaboration("testCollaboration", new()
    {
        CreatorDisplayName = "Creator ",
        CreatorMemberAbilities = new[]
        {
            "CAN_QUERY",
            "CAN_RECEIVE_RESULTS",
        },
        DataEncryptionMetadata = new Aws.CleanRooms.Inputs.CollaborationDataEncryptionMetadataArgs
        {
            AllowClearText = true,
            AllowDuplicates = true,
            AllowJoinsOnColumnsWithDifferentNames = true,
            PreserveNulls = false,
        },
        Description = "I made this collaboration with Pulumi!",
        Members = new[]
        {
            new Aws.CleanRooms.Inputs.CollaborationMemberArgs
            {
                AccountId = "123456789012",
                DisplayName = "Other member",
                MemberAbilities = new() { },
            },
        },
        QueryLogStatus = "DISABLED",
        Tags = 
        {
            { "Project", "Pulumi" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cleanrooms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cleanrooms.NewCollaboration(ctx, "testCollaboration", &cleanrooms.CollaborationArgs{
			CreatorDisplayName: pulumi.String("Creator "),
			CreatorMemberAbilities: pulumi.StringArray{
				pulumi.String("CAN_QUERY"),
				pulumi.String("CAN_RECEIVE_RESULTS"),
			},
			DataEncryptionMetadata: &cleanrooms.CollaborationDataEncryptionMetadataArgs{
				AllowClearText:                        pulumi.Bool(true),
				AllowDuplicates:                       pulumi.Bool(true),
				AllowJoinsOnColumnsWithDifferentNames: pulumi.Bool(true),
				PreserveNulls:                         pulumi.Bool(false),
			},
			Description: pulumi.String("I made this collaboration with Pulumi!"),
			Members: cleanrooms.CollaborationMemberArray{
				&cleanrooms.CollaborationMemberArgs{
					AccountId:       pulumi.String("123456789012"),
					DisplayName:     pulumi.String("Other member"),
					MemberAbilities: pulumi.StringArray{},
				},
			},
			QueryLogStatus: pulumi.String("DISABLED"),
			Tags: pulumi.StringMap{
				"Project": pulumi.String("Pulumi"),
			},
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
import com.pulumi.aws.cleanrooms.Collaboration;
import com.pulumi.aws.cleanrooms.CollaborationArgs;
import com.pulumi.aws.cleanrooms.inputs.CollaborationDataEncryptionMetadataArgs;
import com.pulumi.aws.cleanrooms.inputs.CollaborationMemberArgs;
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
        var testCollaboration = new Collaboration("testCollaboration", CollaborationArgs.builder()        
            .creatorDisplayName("Creator ")
            .creatorMemberAbilities(            
                "CAN_QUERY",
                "CAN_RECEIVE_RESULTS")
            .dataEncryptionMetadata(CollaborationDataEncryptionMetadataArgs.builder()
                .allowClearText(true)
                .allowDuplicates(true)
                .allowJoinsOnColumnsWithDifferentNames(true)
                .preserveNulls(false)
                .build())
            .description("I made this collaboration with Pulumi!")
            .members(CollaborationMemberArgs.builder()
                .accountId(123456789012)
                .displayName("Other member")
                .memberAbilities()
                .build())
            .queryLogStatus("DISABLED")
            .tags(Map.of("Project", "Pulumi"))
            .build());

    }
}
```
```yaml
resources:
  testCollaboration:
    type: aws:cleanrooms:Collaboration
    properties:
      creatorDisplayName: 'Creator '
      creatorMemberAbilities:
        - CAN_QUERY
        - CAN_RECEIVE_RESULTS
      dataEncryptionMetadata:
        allowClearText: true
        allowDuplicates: true
        allowJoinsOnColumnsWithDifferentNames: true
        preserveNulls: false
      description: I made this collaboration with Pulumi!
      members:
        - accountId: 1.23456789012e+11
          displayName: Other member
          memberAbilities: []
      queryLogStatus: DISABLED
      tags:
        Project: Pulumi
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_cleanrooms_collaboration` using the `id`. For example:

```sh
 $ pulumi import aws:cleanrooms/collaboration:Collaboration collaboration 1234abcd-12ab-34cd-56ef-1234567890ab
```
 