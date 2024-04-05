Provides a SageMaker Notebook Instance resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ni = new aws.sagemaker.NotebookInstance("ni", {
    roleArn: aws_iam_role.role.arn,
    instanceType: "ml.t2.medium",
    tags: {
        Name: "foo",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

ni = aws.sagemaker.NotebookInstance("ni",
    role_arn=aws_iam_role["role"]["arn"],
    instance_type="ml.t2.medium",
    tags={
        "Name": "foo",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ni = new Aws.Sagemaker.NotebookInstance("ni", new()
    {
        RoleArn = aws_iam_role.Role.Arn,
        InstanceType = "ml.t2.medium",
        Tags = 
        {
            { "Name", "foo" },
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
		_, err := sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
			RoleArn:      pulumi.Any(aws_iam_role.Role.Arn),
			InstanceType: pulumi.String("ml.t2.medium"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo"),
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
import com.pulumi.aws.sagemaker.NotebookInstance;
import com.pulumi.aws.sagemaker.NotebookInstanceArgs;
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
        var ni = new NotebookInstance("ni", NotebookInstanceArgs.builder()        
            .roleArn(aws_iam_role.role().arn())
            .instanceType("ml.t2.medium")
            .tags(Map.of("Name", "foo"))
            .build());

    }
}
```
```yaml
resources:
  ni:
    type: aws:sagemaker:NotebookInstance
    properties:
      roleArn: ${aws_iam_role.role.arn}
      instanceType: ml.t2.medium
      tags:
        Name: foo
```
{{% /example %}}
{{% example %}}
### Code repository usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.CodeRepository("example", {
    codeRepositoryName: "my-notebook-instance-code-repo",
    gitConfig: {
        repositoryUrl: "https://github.com/github/docs.git",
    },
});
const ni = new aws.sagemaker.NotebookInstance("ni", {
    roleArn: aws_iam_role.role.arn,
    instanceType: "ml.t2.medium",
    defaultCodeRepository: example.codeRepositoryName,
    tags: {
        Name: "foo",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.CodeRepository("example",
    code_repository_name="my-notebook-instance-code-repo",
    git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
        repository_url="https://github.com/github/docs.git",
    ))
ni = aws.sagemaker.NotebookInstance("ni",
    role_arn=aws_iam_role["role"]["arn"],
    instance_type="ml.t2.medium",
    default_code_repository=example.code_repository_name,
    tags={
        "Name": "foo",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.CodeRepository("example", new()
    {
        CodeRepositoryName = "my-notebook-instance-code-repo",
        GitConfig = new Aws.Sagemaker.Inputs.CodeRepositoryGitConfigArgs
        {
            RepositoryUrl = "https://github.com/github/docs.git",
        },
    });

    var ni = new Aws.Sagemaker.NotebookInstance("ni", new()
    {
        RoleArn = aws_iam_role.Role.Arn,
        InstanceType = "ml.t2.medium",
        DefaultCodeRepository = example.CodeRepositoryName,
        Tags = 
        {
            { "Name", "foo" },
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
		example, err := sagemaker.NewCodeRepository(ctx, "example", &sagemaker.CodeRepositoryArgs{
			CodeRepositoryName: pulumi.String("my-notebook-instance-code-repo"),
			GitConfig: &sagemaker.CodeRepositoryGitConfigArgs{
				RepositoryUrl: pulumi.String("https://github.com/github/docs.git"),
			},
		})
		if err != nil {
			return err
		}
		_, err = sagemaker.NewNotebookInstance(ctx, "ni", &sagemaker.NotebookInstanceArgs{
			RoleArn:               pulumi.Any(aws_iam_role.Role.Arn),
			InstanceType:          pulumi.String("ml.t2.medium"),
			DefaultCodeRepository: example.CodeRepositoryName,
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo"),
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
import com.pulumi.aws.sagemaker.CodeRepository;
import com.pulumi.aws.sagemaker.CodeRepositoryArgs;
import com.pulumi.aws.sagemaker.inputs.CodeRepositoryGitConfigArgs;
import com.pulumi.aws.sagemaker.NotebookInstance;
import com.pulumi.aws.sagemaker.NotebookInstanceArgs;
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
        var example = new CodeRepository("example", CodeRepositoryArgs.builder()        
            .codeRepositoryName("my-notebook-instance-code-repo")
            .gitConfig(CodeRepositoryGitConfigArgs.builder()
                .repositoryUrl("https://github.com/github/docs.git")
                .build())
            .build());

        var ni = new NotebookInstance("ni", NotebookInstanceArgs.builder()        
            .roleArn(aws_iam_role.role().arn())
            .instanceType("ml.t2.medium")
            .defaultCodeRepository(example.codeRepositoryName())
            .tags(Map.of("Name", "foo"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:CodeRepository
    properties:
      codeRepositoryName: my-notebook-instance-code-repo
      gitConfig:
        repositoryUrl: https://github.com/github/docs.git
  ni:
    type: aws:sagemaker:NotebookInstance
    properties:
      roleArn: ${aws_iam_role.role.arn}
      instanceType: ml.t2.medium
      defaultCodeRepository: ${example.codeRepositoryName}
      tags:
        Name: foo
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Notebook Instances using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/notebookInstance:NotebookInstance test_notebook_instance my-notebook-instance
```
 