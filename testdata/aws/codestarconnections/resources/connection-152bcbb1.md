Provides a CodeStar Connection.

> **NOTE:** The `aws.codestarconnections.Connection` resource is created in the state `PENDING`. Authentication with the connection provider must be completed in the AWS Console. See the [AWS documentation](https://docs.aws.amazon.com/dtconsole/latest/userguide/connections-update.html) for details.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConnection = new aws.codestarconnections.Connection("exampleConnection", {providerType: "Bitbucket"});
const examplePipeline = new aws.codepipeline.Pipeline("examplePipeline", {
    roleArn: aws_iam_role.codepipeline_role.arn,
    artifactStores: [{}],
    stages: [
        {
            name: "Source",
            actions: [{
                name: "Source",
                category: "Source",
                owner: "AWS",
                provider: "CodeStarSourceConnection",
                version: "1",
                outputArtifacts: ["source_output"],
                configuration: {
                    ConnectionArn: exampleConnection.arn,
                    FullRepositoryId: "my-organization/test",
                    BranchName: "main",
                },
            }],
        },
        {
            name: "Build",
            actions: [{}],
        },
        {
            name: "Deploy",
            actions: [{}],
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_connection = aws.codestarconnections.Connection("exampleConnection", provider_type="Bitbucket")
example_pipeline = aws.codepipeline.Pipeline("examplePipeline",
    role_arn=aws_iam_role["codepipeline_role"]["arn"],
    artifact_stores=[aws.codepipeline.PipelineArtifactStoreArgs()],
    stages=[
        aws.codepipeline.PipelineStageArgs(
            name="Source",
            actions=[aws.codepipeline.PipelineStageActionArgs(
                name="Source",
                category="Source",
                owner="AWS",
                provider="CodeStarSourceConnection",
                version="1",
                output_artifacts=["source_output"],
                configuration={
                    "ConnectionArn": example_connection.arn,
                    "FullRepositoryId": "my-organization/test",
                    "BranchName": "main",
                },
            )],
        ),
        aws.codepipeline.PipelineStageArgs(
            name="Build",
            actions=[aws.codepipeline.PipelineStageActionArgs()],
        ),
        aws.codepipeline.PipelineStageArgs(
            name="Deploy",
            actions=[aws.codepipeline.PipelineStageActionArgs()],
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConnection = new Aws.CodeStarConnections.Connection("exampleConnection", new()
    {
        ProviderType = "Bitbucket",
    });

    var examplePipeline = new Aws.CodePipeline.Pipeline("examplePipeline", new()
    {
        RoleArn = aws_iam_role.Codepipeline_role.Arn,
        ArtifactStores = new[]
        {
            null,
        },
        Stages = new[]
        {
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Source",
                Actions = new[]
                {
                    new Aws.CodePipeline.Inputs.PipelineStageActionArgs
                    {
                        Name = "Source",
                        Category = "Source",
                        Owner = "AWS",
                        Provider = "CodeStarSourceConnection",
                        Version = "1",
                        OutputArtifacts = new[]
                        {
                            "source_output",
                        },
                        Configuration = 
                        {
                            { "ConnectionArn", exampleConnection.Arn },
                            { "FullRepositoryId", "my-organization/test" },
                            { "BranchName", "main" },
                        },
                    },
                },
            },
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Build",
                Actions = new[]
                {
                    null,
                },
            },
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Deploy",
                Actions = new[]
                {
                    null,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codepipeline"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codestarconnections"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleConnection, err := codestarconnections.NewConnection(ctx, "exampleConnection", &codestarconnections.ConnectionArgs{
			ProviderType: pulumi.String("Bitbucket"),
		})
		if err != nil {
			return err
		}
		_, err = codepipeline.NewPipeline(ctx, "examplePipeline", &codepipeline.PipelineArgs{
			RoleArn: pulumi.Any(aws_iam_role.Codepipeline_role.Arn),
			ArtifactStores: codepipeline.PipelineArtifactStoreArray{
				nil,
			},
			Stages: codepipeline.PipelineStageArray{
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Source"),
					Actions: codepipeline.PipelineStageActionArray{
						&codepipeline.PipelineStageActionArgs{
							Name:     pulumi.String("Source"),
							Category: pulumi.String("Source"),
							Owner:    pulumi.String("AWS"),
							Provider: pulumi.String("CodeStarSourceConnection"),
							Version:  pulumi.String("1"),
							OutputArtifacts: pulumi.StringArray{
								pulumi.String("source_output"),
							},
							Configuration: pulumi.StringMap{
								"ConnectionArn":    exampleConnection.Arn,
								"FullRepositoryId": pulumi.String("my-organization/test"),
								"BranchName":       pulumi.String("main"),
							},
						},
					},
				},
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Build"),
					Actions: codepipeline.PipelineStageActionArray{
						nil,
					},
				},
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Deploy"),
					Actions: codepipeline.PipelineStageActionArray{
						nil,
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
import com.pulumi.aws.codestarconnections.Connection;
import com.pulumi.aws.codestarconnections.ConnectionArgs;
import com.pulumi.aws.codepipeline.Pipeline;
import com.pulumi.aws.codepipeline.PipelineArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineArtifactStoreArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineStageArgs;
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
        var exampleConnection = new Connection("exampleConnection", ConnectionArgs.builder()        
            .providerType("Bitbucket")
            .build());

        var examplePipeline = new Pipeline("examplePipeline", PipelineArgs.builder()        
            .roleArn(aws_iam_role.codepipeline_role().arn())
            .artifactStores()
            .stages(            
                PipelineStageArgs.builder()
                    .name("Source")
                    .actions(PipelineStageActionArgs.builder()
                        .name("Source")
                        .category("Source")
                        .owner("AWS")
                        .provider("CodeStarSourceConnection")
                        .version("1")
                        .outputArtifacts("source_output")
                        .configuration(Map.ofEntries(
                            Map.entry("ConnectionArn", exampleConnection.arn()),
                            Map.entry("FullRepositoryId", "my-organization/test"),
                            Map.entry("BranchName", "main")
                        ))
                        .build())
                    .build(),
                PipelineStageArgs.builder()
                    .name("Build")
                    .actions()
                    .build(),
                PipelineStageArgs.builder()
                    .name("Deploy")
                    .actions()
                    .build())
            .build());

    }
}
```
```yaml
resources:
  exampleConnection:
    type: aws:codestarconnections:Connection
    properties:
      providerType: Bitbucket
  examplePipeline:
    type: aws:codepipeline:Pipeline
    properties:
      roleArn: ${aws_iam_role.codepipeline_role.arn}
      artifactStores:
        - {}
      stages:
        - name: Source
          actions:
            - name: Source
              category: Source
              owner: AWS
              provider: CodeStarSourceConnection
              version: '1'
              outputArtifacts:
                - source_output
              configuration:
                ConnectionArn: ${exampleConnection.arn}
                FullRepositoryId: my-organization/test
                BranchName: main
        - name: Build
          actions:
            - {}
        - name: Deploy
          actions:
            - {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeStar connections using the ARN. For example:

```sh
 $ pulumi import aws:codestarconnections/connection:Connection test-connection arn:aws:codestar-connections:us-west-1:0123456789:connection/79d4d357-a2ee-41e4-b350-2fe39ae59448
```
 