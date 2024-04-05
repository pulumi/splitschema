Manages a CodeBuild webhook, which is an endpoint accepted by the CodeBuild service to trigger builds from source code repositories. Depending on the source type of the CodeBuild project, the CodeBuild service may also automatically create and delete the actual repository webhook as well.

{{% examples %}}
## Example Usage
{{% example %}}
### Bitbucket and GitHub

When working with [Bitbucket](https://bitbucket.org) and [GitHub](https://github.com) source CodeBuild webhooks, the CodeBuild service will automatically create (on `aws.codebuild.Webhook` resource creation) and delete (on `aws.codebuild.Webhook` resource deletion) the Bitbucket/GitHub repository webhook using its granted OAuth permissions. This behavior cannot be controlled by this provider.

> **Note:** The AWS account that this provider uses to create this resource *must* have authorized CodeBuild to access Bitbucket/GitHub's OAuth API in each applicable region. This is a manual step that must be done *before* creating webhooks with this resource. If OAuth is not configured, AWS will return an error similar to `ResourceNotFoundException: Could not find access token for server type github`. More information can be found in the CodeBuild User Guide for [Bitbucket](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-bitbucket-pull-request.html) and [GitHub](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-pull-request.html).

> **Note:** Further managing the automatically created Bitbucket/GitHub webhook with the `bitbucket_hook`/`github_repository_webhook` resource is only possible with importing that resource after creation of the `aws.codebuild.Webhook` resource. The CodeBuild API does not ever provide the `secret` attribute for the `aws.codebuild.Webhook` resource in this scenario.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codebuild.Webhook("example", {
    projectName: aws_codebuild_project.example.name,
    buildType: "BUILD",
    filterGroups: [{
        filters: [
            {
                type: "EVENT",
                pattern: "PUSH",
            },
            {
                type: "BASE_REF",
                pattern: "master",
            },
        ],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codebuild.Webhook("example",
    project_name=aws_codebuild_project["example"]["name"],
    build_type="BUILD",
    filter_groups=[aws.codebuild.WebhookFilterGroupArgs(
        filters=[
            aws.codebuild.WebhookFilterGroupFilterArgs(
                type="EVENT",
                pattern="PUSH",
            ),
            aws.codebuild.WebhookFilterGroupFilterArgs(
                type="BASE_REF",
                pattern="master",
            ),
        ],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeBuild.Webhook("example", new()
    {
        ProjectName = aws_codebuild_project.Example.Name,
        BuildType = "BUILD",
        FilterGroups = new[]
        {
            new Aws.CodeBuild.Inputs.WebhookFilterGroupArgs
            {
                Filters = new[]
                {
                    new Aws.CodeBuild.Inputs.WebhookFilterGroupFilterArgs
                    {
                        Type = "EVENT",
                        Pattern = "PUSH",
                    },
                    new Aws.CodeBuild.Inputs.WebhookFilterGroupFilterArgs
                    {
                        Type = "BASE_REF",
                        Pattern = "master",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codebuild.NewWebhook(ctx, "example", &codebuild.WebhookArgs{
			ProjectName: pulumi.Any(aws_codebuild_project.Example.Name),
			BuildType:   pulumi.String("BUILD"),
			FilterGroups: codebuild.WebhookFilterGroupArray{
				&codebuild.WebhookFilterGroupArgs{
					Filters: codebuild.WebhookFilterGroupFilterArray{
						&codebuild.WebhookFilterGroupFilterArgs{
							Type:    pulumi.String("EVENT"),
							Pattern: pulumi.String("PUSH"),
						},
						&codebuild.WebhookFilterGroupFilterArgs{
							Type:    pulumi.String("BASE_REF"),
							Pattern: pulumi.String("master"),
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
import com.pulumi.aws.codebuild.Webhook;
import com.pulumi.aws.codebuild.WebhookArgs;
import com.pulumi.aws.codebuild.inputs.WebhookFilterGroupArgs;
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
        var example = new Webhook("example", WebhookArgs.builder()        
            .projectName(aws_codebuild_project.example().name())
            .buildType("BUILD")
            .filterGroups(WebhookFilterGroupArgs.builder()
                .filters(                
                    WebhookFilterGroupFilterArgs.builder()
                        .type("EVENT")
                        .pattern("PUSH")
                        .build(),
                    WebhookFilterGroupFilterArgs.builder()
                        .type("BASE_REF")
                        .pattern("master")
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codebuild:Webhook
    properties:
      projectName: ${aws_codebuild_project.example.name}
      buildType: BUILD
      filterGroups:
        - filters:
            - type: EVENT
              pattern: PUSH
            - type: BASE_REF
              pattern: master
```
{{% /example %}}
{{% example %}}
### GitHub Enterprise

When working with [GitHub Enterprise](https://enterprise.github.com/) source CodeBuild webhooks, the GHE repository webhook must be separately managed (e.g., manually or with the `github_repository_webhook` resource).

More information creating webhooks with GitHub Enterprise can be found in the [CodeBuild User Guide](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-enterprise.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as github from "@pulumi/github";

const exampleWebhook = new aws.codebuild.Webhook("exampleWebhook", {projectName: aws_codebuild_project.example.name});
const exampleRepositoryWebhook = new github.RepositoryWebhook("exampleRepositoryWebhook", {
    active: true,
    events: ["push"],
    repository: github_repository.example.name,
    configuration: {
        url: exampleWebhook.payloadUrl,
        secret: exampleWebhook.secret,
        contentType: "json",
        insecureSsl: false,
    },
});
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_github as github

example_webhook = aws.codebuild.Webhook("exampleWebhook", project_name=aws_codebuild_project["example"]["name"])
example_repository_webhook = github.RepositoryWebhook("exampleRepositoryWebhook",
    active=True,
    events=["push"],
    repository=github_repository["example"]["name"],
    configuration=github.RepositoryWebhookConfigurationArgs(
        url=example_webhook.payload_url,
        secret=example_webhook.secret,
        content_type="json",
        insecure_ssl=False,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Github = Pulumi.Github;

return await Deployment.RunAsync(() => 
{
    var exampleWebhook = new Aws.CodeBuild.Webhook("exampleWebhook", new()
    {
        ProjectName = aws_codebuild_project.Example.Name,
    });

    var exampleRepositoryWebhook = new Github.RepositoryWebhook("exampleRepositoryWebhook", new()
    {
        Active = true,
        Events = new[]
        {
            "push",
        },
        Repository = github_repository.Example.Name,
        Configuration = new Github.Inputs.RepositoryWebhookConfigurationArgs
        {
            Url = exampleWebhook.PayloadUrl,
            Secret = exampleWebhook.Secret,
            ContentType = "json",
            InsecureSsl = false,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
	"github.com/pulumi/pulumi-github/sdk/v5/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleWebhook, err := codebuild.NewWebhook(ctx, "exampleWebhook", &codebuild.WebhookArgs{
			ProjectName: pulumi.Any(aws_codebuild_project.Example.Name),
		})
		if err != nil {
			return err
		}
		_, err = github.NewRepositoryWebhook(ctx, "exampleRepositoryWebhook", &github.RepositoryWebhookArgs{
			Active: pulumi.Bool(true),
			Events: pulumi.StringArray{
				pulumi.String("push"),
			},
			Repository: pulumi.Any(github_repository.Example.Name),
			Configuration: &github.RepositoryWebhookConfigurationArgs{
				Url:         exampleWebhook.PayloadUrl,
				Secret:      exampleWebhook.Secret,
				ContentType: pulumi.String("json"),
				InsecureSsl: pulumi.Bool(false),
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
import com.pulumi.aws.codebuild.Webhook;
import com.pulumi.aws.codebuild.WebhookArgs;
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
        var exampleWebhook = new Webhook("exampleWebhook", WebhookArgs.builder()        
            .projectName(aws_codebuild_project.example().name())
            .build());

        var exampleRepositoryWebhook = new RepositoryWebhook("exampleRepositoryWebhook", RepositoryWebhookArgs.builder()        
            .active(true)
            .events("push")
            .repository(github_repository.example().name())
            .configuration(RepositoryWebhookConfigurationArgs.builder()
                .url(exampleWebhook.payloadUrl())
                .secret(exampleWebhook.secret())
                .contentType("json")
                .insecureSsl(false)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleWebhook:
    type: aws:codebuild:Webhook
    properties:
      projectName: ${aws_codebuild_project.example.name}
  exampleRepositoryWebhook:
    type: github:RepositoryWebhook
    properties:
      active: true
      events:
        - push
      repository: ${github_repository.example.name}
      configuration:
        url: ${exampleWebhook.payloadUrl}
        secret: ${exampleWebhook.secret}
        contentType: json
        insecureSsl: false
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeBuild Webhooks using the CodeBuild Project name. For example:

```sh
 $ pulumi import aws:codebuild/webhook:Webhook example MyProjectName
```
 