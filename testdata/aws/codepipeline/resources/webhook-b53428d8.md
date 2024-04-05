Provides a CodePipeline Webhook.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as github from "@pulumi/github";

const barPipeline = new aws.codepipeline.Pipeline("barPipeline", {
    roleArn: aws_iam_role.bar.arn,
    artifactStores: [{
        location: aws_s3_bucket.bar.bucket,
        type: "S3",
        encryptionKey: {
            id: data.aws_kms_alias.s3kmskey.arn,
            type: "KMS",
        },
    }],
    stages: [
        {
            name: "Source",
            actions: [{
                name: "Source",
                category: "Source",
                owner: "ThirdParty",
                provider: "GitHub",
                version: "1",
                outputArtifacts: ["test"],
                configuration: {
                    Owner: "my-organization",
                    Repo: "test",
                    Branch: "master",
                },
            }],
        },
        {
            name: "Build",
            actions: [{
                name: "Build",
                category: "Build",
                owner: "AWS",
                provider: "CodeBuild",
                inputArtifacts: ["test"],
                version: "1",
                configuration: {
                    ProjectName: "test",
                },
            }],
        },
    ],
});
const webhookSecret = "super-secret";
const barWebhook = new aws.codepipeline.Webhook("barWebhook", {
    authentication: "GITHUB_HMAC",
    targetAction: "Source",
    targetPipeline: barPipeline.name,
    authenticationConfiguration: {
        secretToken: webhookSecret,
    },
    filters: [{
        jsonPath: "$.ref",
        matchEquals: "refs/heads/{Branch}",
    }],
});
// Wire the CodePipeline webhook into a GitHub repository.
const barRepositoryWebhook = new github.RepositoryWebhook("barRepositoryWebhook", {
    repository: github_repository.repo.name,
    configuration: {
        url: barWebhook.url,
        contentType: "json",
        insecureSsl: true,
        secret: webhookSecret,
    },
    events: ["push"],
});
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_github as github

bar_pipeline = aws.codepipeline.Pipeline("barPipeline",
    role_arn=aws_iam_role["bar"]["arn"],
    artifact_stores=[aws.codepipeline.PipelineArtifactStoreArgs(
        location=aws_s3_bucket["bar"]["bucket"],
        type="S3",
        encryption_key=aws.codepipeline.PipelineArtifactStoreEncryptionKeyArgs(
            id=data["aws_kms_alias"]["s3kmskey"]["arn"],
            type="KMS",
        ),
    )],
    stages=[
        aws.codepipeline.PipelineStageArgs(
            name="Source",
            actions=[aws.codepipeline.PipelineStageActionArgs(
                name="Source",
                category="Source",
                owner="ThirdParty",
                provider="GitHub",
                version="1",
                output_artifacts=["test"],
                configuration={
                    "Owner": "my-organization",
                    "Repo": "test",
                    "Branch": "master",
                },
            )],
        ),
        aws.codepipeline.PipelineStageArgs(
            name="Build",
            actions=[aws.codepipeline.PipelineStageActionArgs(
                name="Build",
                category="Build",
                owner="AWS",
                provider="CodeBuild",
                input_artifacts=["test"],
                version="1",
                configuration={
                    "ProjectName": "test",
                },
            )],
        ),
    ])
webhook_secret = "super-secret"
bar_webhook = aws.codepipeline.Webhook("barWebhook",
    authentication="GITHUB_HMAC",
    target_action="Source",
    target_pipeline=bar_pipeline.name,
    authentication_configuration=aws.codepipeline.WebhookAuthenticationConfigurationArgs(
        secret_token=webhook_secret,
    ),
    filters=[aws.codepipeline.WebhookFilterArgs(
        json_path="$.ref",
        match_equals="refs/heads/{Branch}",
    )])
# Wire the CodePipeline webhook into a GitHub repository.
bar_repository_webhook = github.RepositoryWebhook("barRepositoryWebhook",
    repository=github_repository["repo"]["name"],
    configuration=github.RepositoryWebhookConfigurationArgs(
        url=bar_webhook.url,
        content_type="json",
        insecure_ssl=True,
        secret=webhook_secret,
    ),
    events=["push"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Github = Pulumi.Github;

return await Deployment.RunAsync(() => 
{
    var barPipeline = new Aws.CodePipeline.Pipeline("barPipeline", new()
    {
        RoleArn = aws_iam_role.Bar.Arn,
        ArtifactStores = new[]
        {
            new Aws.CodePipeline.Inputs.PipelineArtifactStoreArgs
            {
                Location = aws_s3_bucket.Bar.Bucket,
                Type = "S3",
                EncryptionKey = new Aws.CodePipeline.Inputs.PipelineArtifactStoreEncryptionKeyArgs
                {
                    Id = data.Aws_kms_alias.S3kmskey.Arn,
                    Type = "KMS",
                },
            },
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
                        Owner = "ThirdParty",
                        Provider = "GitHub",
                        Version = "1",
                        OutputArtifacts = new[]
                        {
                            "test",
                        },
                        Configuration = 
                        {
                            { "Owner", "my-organization" },
                            { "Repo", "test" },
                            { "Branch", "master" },
                        },
                    },
                },
            },
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Build",
                Actions = new[]
                {
                    new Aws.CodePipeline.Inputs.PipelineStageActionArgs
                    {
                        Name = "Build",
                        Category = "Build",
                        Owner = "AWS",
                        Provider = "CodeBuild",
                        InputArtifacts = new[]
                        {
                            "test",
                        },
                        Version = "1",
                        Configuration = 
                        {
                            { "ProjectName", "test" },
                        },
                    },
                },
            },
        },
    });

    var webhookSecret = "super-secret";

    var barWebhook = new Aws.CodePipeline.Webhook("barWebhook", new()
    {
        Authentication = "GITHUB_HMAC",
        TargetAction = "Source",
        TargetPipeline = barPipeline.Name,
        AuthenticationConfiguration = new Aws.CodePipeline.Inputs.WebhookAuthenticationConfigurationArgs
        {
            SecretToken = webhookSecret,
        },
        Filters = new[]
        {
            new Aws.CodePipeline.Inputs.WebhookFilterArgs
            {
                JsonPath = "$.ref",
                MatchEquals = "refs/heads/{Branch}",
            },
        },
    });

    // Wire the CodePipeline webhook into a GitHub repository.
    var barRepositoryWebhook = new Github.RepositoryWebhook("barRepositoryWebhook", new()
    {
        Repository = github_repository.Repo.Name,
        Configuration = new Github.Inputs.RepositoryWebhookConfigurationArgs
        {
            Url = barWebhook.Url,
            ContentType = "json",
            InsecureSsl = true,
            Secret = webhookSecret,
        },
        Events = new[]
        {
            "push",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codepipeline"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		barPipeline, err := codepipeline.NewPipeline(ctx, "barPipeline", &codepipeline.PipelineArgs{
			RoleArn: pulumi.Any(aws_iam_role.Bar.Arn),
			ArtifactStores: codepipeline.PipelineArtifactStoreArray{
				&codepipeline.PipelineArtifactStoreArgs{
					Location: pulumi.Any(aws_s3_bucket.Bar.Bucket),
					Type:     pulumi.String("S3"),
					EncryptionKey: &codepipeline.PipelineArtifactStoreEncryptionKeyArgs{
						Id:   pulumi.Any(data.Aws_kms_alias.S3kmskey.Arn),
						Type: pulumi.String("KMS"),
					},
				},
			},
			Stages: codepipeline.PipelineStageArray{
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Source"),
					Actions: codepipeline.PipelineStageActionArray{
						&codepipeline.PipelineStageActionArgs{
							Name:     pulumi.String("Source"),
							Category: pulumi.String("Source"),
							Owner:    pulumi.String("ThirdParty"),
							Provider: pulumi.String("GitHub"),
							Version:  pulumi.String("1"),
							OutputArtifacts: pulumi.StringArray{
								pulumi.String("test"),
							},
							Configuration: pulumi.StringMap{
								"Owner":  pulumi.String("my-organization"),
								"Repo":   pulumi.String("test"),
								"Branch": pulumi.String("master"),
							},
						},
					},
				},
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Build"),
					Actions: codepipeline.PipelineStageActionArray{
						&codepipeline.PipelineStageActionArgs{
							Name:     pulumi.String("Build"),
							Category: pulumi.String("Build"),
							Owner:    pulumi.String("AWS"),
							Provider: pulumi.String("CodeBuild"),
							InputArtifacts: pulumi.StringArray{
								pulumi.String("test"),
							},
							Version: pulumi.String("1"),
							Configuration: pulumi.StringMap{
								"ProjectName": pulumi.String("test"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		webhookSecret := "super-secret"
		barWebhook, err := codepipeline.NewWebhook(ctx, "barWebhook", &codepipeline.WebhookArgs{
			Authentication: pulumi.String("GITHUB_HMAC"),
			TargetAction:   pulumi.String("Source"),
			TargetPipeline: barPipeline.Name,
			AuthenticationConfiguration: &codepipeline.WebhookAuthenticationConfigurationArgs{
				SecretToken: pulumi.String(webhookSecret),
			},
			Filters: codepipeline.WebhookFilterArray{
				&codepipeline.WebhookFilterArgs{
					JsonPath:    pulumi.String("$.ref"),
					MatchEquals: pulumi.String("refs/heads/{Branch}"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = github.NewRepositoryWebhook(ctx, "barRepositoryWebhook", &github.RepositoryWebhookArgs{
			Repository: pulumi.Any(github_repository.Repo.Name),
			Configuration: &github.RepositoryWebhookConfigurationArgs{
				Url:         barWebhook.Url,
				ContentType: pulumi.String("json"),
				InsecureSsl: pulumi.Bool(true),
				Secret:      pulumi.String(webhookSecret),
			},
			Events: pulumi.StringArray{
				pulumi.String("push"),
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
import com.pulumi.aws.codepipeline.Pipeline;
import com.pulumi.aws.codepipeline.PipelineArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineArtifactStoreArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineArtifactStoreEncryptionKeyArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineStageArgs;
import com.pulumi.aws.codepipeline.Webhook;
import com.pulumi.aws.codepipeline.WebhookArgs;
import com.pulumi.aws.codepipeline.inputs.WebhookAuthenticationConfigurationArgs;
import com.pulumi.aws.codepipeline.inputs.WebhookFilterArgs;
import com.pulumi.github.RepositoryWebhook;
import com.pulumi.github.RepositoryWebhookArgs;
import com.pulumi.github.inputs.RepositoryWebhookConfigurationArgs;
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
        var barPipeline = new Pipeline("barPipeline", PipelineArgs.builder()        
            .roleArn(aws_iam_role.bar().arn())
            .artifactStores(PipelineArtifactStoreArgs.builder()
                .location(aws_s3_bucket.bar().bucket())
                .type("S3")
                .encryptionKey(PipelineArtifactStoreEncryptionKeyArgs.builder()
                    .id(data.aws_kms_alias().s3kmskey().arn())
                    .type("KMS")
                    .build())
                .build())
            .stages(            
                PipelineStageArgs.builder()
                    .name("Source")
                    .actions(PipelineStageActionArgs.builder()
                        .name("Source")
                        .category("Source")
                        .owner("ThirdParty")
                        .provider("GitHub")
                        .version("1")
                        .outputArtifacts("test")
                        .configuration(Map.ofEntries(
                            Map.entry("Owner", "my-organization"),
                            Map.entry("Repo", "test"),
                            Map.entry("Branch", "master")
                        ))
                        .build())
                    .build(),
                PipelineStageArgs.builder()
                    .name("Build")
                    .actions(PipelineStageActionArgs.builder()
                        .name("Build")
                        .category("Build")
                        .owner("AWS")
                        .provider("CodeBuild")
                        .inputArtifacts("test")
                        .version("1")
                        .configuration(Map.of("ProjectName", "test"))
                        .build())
                    .build())
            .build());

        final var webhookSecret = "super-secret";

        var barWebhook = new Webhook("barWebhook", WebhookArgs.builder()        
            .authentication("GITHUB_HMAC")
            .targetAction("Source")
            .targetPipeline(barPipeline.name())
            .authenticationConfiguration(WebhookAuthenticationConfigurationArgs.builder()
                .secretToken(webhookSecret)
                .build())
            .filters(WebhookFilterArgs.builder()
                .jsonPath("$.ref")
                .matchEquals("refs/heads/{Branch}")
                .build())
            .build());

        var barRepositoryWebhook = new RepositoryWebhook("barRepositoryWebhook", RepositoryWebhookArgs.builder()        
            .repository(github_repository.repo().name())
            .configuration(RepositoryWebhookConfigurationArgs.builder()
                .url(barWebhook.url())
                .contentType("json")
                .insecureSsl(true)
                .secret(webhookSecret)
                .build())
            .events("push")
            .build());

    }
}
```
```yaml
resources:
  barPipeline:
    type: aws:codepipeline:Pipeline
    properties:
      roleArn: ${aws_iam_role.bar.arn}
      artifactStores:
        - location: ${aws_s3_bucket.bar.bucket}
          type: S3
          encryptionKey:
            id: ${data.aws_kms_alias.s3kmskey.arn}
            type: KMS
      stages:
        - name: Source
          actions:
            - name: Source
              category: Source
              owner: ThirdParty
              provider: GitHub
              version: '1'
              outputArtifacts:
                - test
              configuration:
                Owner: my-organization
                Repo: test
                Branch: master
        - name: Build
          actions:
            - name: Build
              category: Build
              owner: AWS
              provider: CodeBuild
              inputArtifacts:
                - test
              version: '1'
              configuration:
                ProjectName: test
  barWebhook:
    type: aws:codepipeline:Webhook
    properties:
      authentication: GITHUB_HMAC
      targetAction: Source
      targetPipeline: ${barPipeline.name}
      authenticationConfiguration:
        secretToken: ${webhookSecret}
      filters:
        - jsonPath: $.ref
          matchEquals: refs/heads/{Branch}
  # Wire the CodePipeline webhook into a GitHub repository.
  barRepositoryWebhook:
    type: github:RepositoryWebhook
    properties:
      repository: ${github_repository.repo.name}
      configuration:
        url: ${barWebhook.url}
        contentType: json
        insecureSsl: true
        secret: ${webhookSecret}
      events:
        - push
variables:
  webhookSecret: super-secret
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodePipeline Webhooks using their ARN. For example:

```sh
 $ pulumi import aws:codepipeline/webhook:Webhook example arn:aws:codepipeline:us-west-2:123456789012:webhook:example
```
 