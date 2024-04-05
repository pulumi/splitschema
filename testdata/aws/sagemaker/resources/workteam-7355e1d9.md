Provides a SageMaker Workteam resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Cognito Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Workteam("example", {
    workteamName: "example",
    workforceName: aws_sagemaker_workforce.example.id,
    description: "example",
    memberDefinitions: [{
        cognitoMemberDefinition: {
            clientId: aws_cognito_user_pool_client.example.id,
            userPool: aws_cognito_user_pool_domain.example.user_pool_id,
            userGroup: aws_cognito_user_group.example.id,
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Workteam("example",
    workteam_name="example",
    workforce_name=aws_sagemaker_workforce["example"]["id"],
    description="example",
    member_definitions=[aws.sagemaker.WorkteamMemberDefinitionArgs(
        cognito_member_definition=aws.sagemaker.WorkteamMemberDefinitionCognitoMemberDefinitionArgs(
            client_id=aws_cognito_user_pool_client["example"]["id"],
            user_pool=aws_cognito_user_pool_domain["example"]["user_pool_id"],
            user_group=aws_cognito_user_group["example"]["id"],
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Workteam("example", new()
    {
        WorkteamName = "example",
        WorkforceName = aws_sagemaker_workforce.Example.Id,
        Description = "example",
        MemberDefinitions = new[]
        {
            new Aws.Sagemaker.Inputs.WorkteamMemberDefinitionArgs
            {
                CognitoMemberDefinition = new Aws.Sagemaker.Inputs.WorkteamMemberDefinitionCognitoMemberDefinitionArgs
                {
                    ClientId = aws_cognito_user_pool_client.Example.Id,
                    UserPool = aws_cognito_user_pool_domain.Example.User_pool_id,
                    UserGroup = aws_cognito_user_group.Example.Id,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewWorkteam(ctx, "example", &sagemaker.WorkteamArgs{
			WorkteamName:  pulumi.String("example"),
			WorkforceName: pulumi.Any(aws_sagemaker_workforce.Example.Id),
			Description:   pulumi.String("example"),
			MemberDefinitions: sagemaker.WorkteamMemberDefinitionArray{
				&sagemaker.WorkteamMemberDefinitionArgs{
					CognitoMemberDefinition: &sagemaker.WorkteamMemberDefinitionCognitoMemberDefinitionArgs{
						ClientId:  pulumi.Any(aws_cognito_user_pool_client.Example.Id),
						UserPool:  pulumi.Any(aws_cognito_user_pool_domain.Example.User_pool_id),
						UserGroup: pulumi.Any(aws_cognito_user_group.Example.Id),
					},
				},
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
import com.pulumi.aws.sagemaker.Workteam;
import com.pulumi.aws.sagemaker.WorkteamArgs;
import com.pulumi.aws.sagemaker.inputs.WorkteamMemberDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.WorkteamMemberDefinitionCognitoMemberDefinitionArgs;
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
        var example = new Workteam("example", WorkteamArgs.builder()        
            .workteamName("example")
            .workforceName(aws_sagemaker_workforce.example().id())
            .description("example")
            .memberDefinitions(WorkteamMemberDefinitionArgs.builder()
                .cognitoMemberDefinition(WorkteamMemberDefinitionCognitoMemberDefinitionArgs.builder()
                    .clientId(aws_cognito_user_pool_client.example().id())
                    .userPool(aws_cognito_user_pool_domain.example().user_pool_id())
                    .userGroup(aws_cognito_user_group.example().id())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Workteam
    properties:
      workteamName: example
      workforceName: ${aws_sagemaker_workforce.example.id}
      description: example
      memberDefinitions:
        - cognitoMemberDefinition:
            clientId: ${aws_cognito_user_pool_client.example.id}
            userPool: ${aws_cognito_user_pool_domain.example.user_pool_id}
            userGroup: ${aws_cognito_user_group.example.id}
```
{{% /example %}}
{{% example %}}
### Oidc Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Workteam("example", {
    workteamName: "example",
    workforceName: aws_sagemaker_workforce.example.id,
    description: "example",
    memberDefinitions: [{
        oidcMemberDefinition: {
            groups: ["example"],
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Workteam("example",
    workteam_name="example",
    workforce_name=aws_sagemaker_workforce["example"]["id"],
    description="example",
    member_definitions=[aws.sagemaker.WorkteamMemberDefinitionArgs(
        oidc_member_definition=aws.sagemaker.WorkteamMemberDefinitionOidcMemberDefinitionArgs(
            groups=["example"],
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Workteam("example", new()
    {
        WorkteamName = "example",
        WorkforceName = aws_sagemaker_workforce.Example.Id,
        Description = "example",
        MemberDefinitions = new[]
        {
            new Aws.Sagemaker.Inputs.WorkteamMemberDefinitionArgs
            {
                OidcMemberDefinition = new Aws.Sagemaker.Inputs.WorkteamMemberDefinitionOidcMemberDefinitionArgs
                {
                    Groups = new[]
                    {
                        "example",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewWorkteam(ctx, "example", &sagemaker.WorkteamArgs{
			WorkteamName:  pulumi.String("example"),
			WorkforceName: pulumi.Any(aws_sagemaker_workforce.Example.Id),
			Description:   pulumi.String("example"),
			MemberDefinitions: sagemaker.WorkteamMemberDefinitionArray{
				&sagemaker.WorkteamMemberDefinitionArgs{
					OidcMemberDefinition: &sagemaker.WorkteamMemberDefinitionOidcMemberDefinitionArgs{
						Groups: pulumi.StringArray{
							pulumi.String("example"),
						},
					},
				},
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
import com.pulumi.aws.sagemaker.Workteam;
import com.pulumi.aws.sagemaker.WorkteamArgs;
import com.pulumi.aws.sagemaker.inputs.WorkteamMemberDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.WorkteamMemberDefinitionOidcMemberDefinitionArgs;
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
        var example = new Workteam("example", WorkteamArgs.builder()        
            .workteamName("example")
            .workforceName(aws_sagemaker_workforce.example().id())
            .description("example")
            .memberDefinitions(WorkteamMemberDefinitionArgs.builder()
                .oidcMemberDefinition(WorkteamMemberDefinitionOidcMemberDefinitionArgs.builder()
                    .groups("example")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Workteam
    properties:
      workteamName: example
      workforceName: ${aws_sagemaker_workforce.example.id}
      description: example
      memberDefinitions:
        - oidcMemberDefinition:
            groups:
              - example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Workteams using the `workteam_name`. For example:

```sh
 $ pulumi import aws:sagemaker/workteam:Workteam example example
```
 