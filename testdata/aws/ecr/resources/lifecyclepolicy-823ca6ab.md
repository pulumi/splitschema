Manages an ECR repository lifecycle policy.

> **NOTE:** Only one `aws.ecr.LifecyclePolicy` resource can be used with the same ECR repository. To apply multiple rules, they must be combined in the `policy` JSON.

> **NOTE:** The AWS ECR API seems to reorder rules based on `rulePriority`. If you define multiple rules that are not sorted in ascending `rulePriority` order in the this provider code, the resource will be flagged for recreation every deployment.

{{% examples %}}
## Example Usage
{{% example %}}
### Policy on untagged image

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ecr.Repository("foo", {});
const foopolicy = new aws.ecr.LifecyclePolicy("foopolicy", {
    repository: foo.name,
    policy: `{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Expire images older than 14 days",
            "selection": {
                "tagStatus": "untagged",
                "countType": "sinceImagePushed",
                "countUnit": "days",
                "countNumber": 14
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo")
foopolicy = aws.ecr.LifecyclePolicy("foopolicy",
    repository=foo.name,
    policy="""{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Expire images older than 14 days",
            "selection": {
                "tagStatus": "untagged",
                "countType": "sinceImagePushed",
                "countUnit": "days",
                "countNumber": 14
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ecr.Repository("foo");

    var foopolicy = new Aws.Ecr.LifecyclePolicy("foopolicy", new()
    {
        Repository = foo.Name,
        Policy = @"{
    ""rules"": [
        {
            ""rulePriority"": 1,
            ""description"": ""Expire images older than 14 days"",
            ""selection"": {
                ""tagStatus"": ""untagged"",
                ""countType"": ""sinceImagePushed"",
                ""countUnit"": ""days"",
                ""countNumber"": 14
            },
            ""action"": {
                ""type"": ""expire""
            }
        }
    ]
}
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ecr.NewRepository(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewLifecyclePolicy(ctx, "foopolicy", &ecr.LifecyclePolicyArgs{
			Repository: foo.Name,
			Policy: pulumi.Any(`{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Expire images older than 14 days",
            "selection": {
                "tagStatus": "untagged",
                "countType": "sinceImagePushed",
                "countUnit": "days",
                "countNumber": 14
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
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
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.ecr.LifecyclePolicy;
import com.pulumi.aws.ecr.LifecyclePolicyArgs;
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
        var foo = new Repository("foo");

        var foopolicy = new LifecyclePolicy("foopolicy", LifecyclePolicyArgs.builder()        
            .repository(foo.name())
            .policy("""
{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Expire images older than 14 days",
            "selection": {
                "tagStatus": "untagged",
                "countType": "sinceImagePushed",
                "countUnit": "days",
                "countNumber": 14
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
            """)
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ecr:Repository
  foopolicy:
    type: aws:ecr:LifecyclePolicy
    properties:
      repository: ${foo.name}
      policy: |
        {
            "rules": [
                {
                    "rulePriority": 1,
                    "description": "Expire images older than 14 days",
                    "selection": {
                        "tagStatus": "untagged",
                        "countType": "sinceImagePushed",
                        "countUnit": "days",
                        "countNumber": 14
                    },
                    "action": {
                        "type": "expire"
                    }
                }
            ]
        }
```
{{% /example %}}
{{% example %}}
### Policy on tagged image

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ecr.Repository("foo", {});
const foopolicy = new aws.ecr.LifecyclePolicy("foopolicy", {
    repository: foo.name,
    policy: `{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Keep last 30 images",
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["v"],
                "countType": "imageCountMoreThan",
                "countNumber": 30
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ecr.Repository("foo")
foopolicy = aws.ecr.LifecyclePolicy("foopolicy",
    repository=foo.name,
    policy="""{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Keep last 30 images",
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["v"],
                "countType": "imageCountMoreThan",
                "countNumber": 30
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ecr.Repository("foo");

    var foopolicy = new Aws.Ecr.LifecyclePolicy("foopolicy", new()
    {
        Repository = foo.Name,
        Policy = @"{
    ""rules"": [
        {
            ""rulePriority"": 1,
            ""description"": ""Keep last 30 images"",
            ""selection"": {
                ""tagStatus"": ""tagged"",
                ""tagPrefixList"": [""v""],
                ""countType"": ""imageCountMoreThan"",
                ""countNumber"": 30
            },
            ""action"": {
                ""type"": ""expire""
            }
        }
    ]
}
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ecr.NewRepository(ctx, "foo", nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewLifecyclePolicy(ctx, "foopolicy", &ecr.LifecyclePolicyArgs{
			Repository: foo.Name,
			Policy: pulumi.Any(`{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Keep last 30 images",
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["v"],
                "countType": "imageCountMoreThan",
                "countNumber": 30
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
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
import com.pulumi.aws.ecr.Repository;
import com.pulumi.aws.ecr.LifecyclePolicy;
import com.pulumi.aws.ecr.LifecyclePolicyArgs;
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
        var foo = new Repository("foo");

        var foopolicy = new LifecyclePolicy("foopolicy", LifecyclePolicyArgs.builder()        
            .repository(foo.name())
            .policy("""
{
    "rules": [
        {
            "rulePriority": 1,
            "description": "Keep last 30 images",
            "selection": {
                "tagStatus": "tagged",
                "tagPrefixList": ["v"],
                "countType": "imageCountMoreThan",
                "countNumber": 30
            },
            "action": {
                "type": "expire"
            }
        }
    ]
}
            """)
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ecr:Repository
  foopolicy:
    type: aws:ecr:LifecyclePolicy
    properties:
      repository: ${foo.name}
      policy: |
        {
            "rules": [
                {
                    "rulePriority": 1,
                    "description": "Keep last 30 images",
                    "selection": {
                        "tagStatus": "tagged",
                        "tagPrefixList": ["v"],
                        "countType": "imageCountMoreThan",
                        "countNumber": 30
                    },
                    "action": {
                        "type": "expire"
                    }
                }
            ]
        }
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECR Lifecycle Policy using the name of the repository. For example:

```sh
 $ pulumi import aws:ecr/lifecyclePolicy:LifecyclePolicy example tf-example
```
 