Provides a CodeArtifact Repository Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {description: "domain key"});
const exampleDomain = new aws.codeartifact.Domain("exampleDomain", {
    domain: "example",
    encryptionKey: exampleKey.arn,
});
const test = new aws.codeartifact.Repository("test", {
    repository: "example",
    domain: exampleDomain.domain,
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey", description="domain key")
example_domain = aws.codeartifact.Domain("exampleDomain",
    domain="example",
    encryption_key=example_key.arn)
test = aws.codeartifact.Repository("test",
    repository="example",
    domain=example_domain.domain)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "domain key",
    });

    var exampleDomain = new Aws.CodeArtifact.Domain("exampleDomain", new()
    {
        DomainName = "example",
        EncryptionKey = exampleKey.Arn,
    });

    var test = new Aws.CodeArtifact.Repository("test", new()
    {
        RepositoryName = "example",
        Domain = exampleDomain.DomainName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeartifact"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description: pulumi.String("domain key"),
		})
		if err != nil {
			return err
		}
		exampleDomain, err := codeartifact.NewDomain(ctx, "exampleDomain", &codeartifact.DomainArgs{
			Domain:        pulumi.String("example"),
			EncryptionKey: exampleKey.Arn,
		})
		if err != nil {
			return err
		}
		_, err = codeartifact.NewRepository(ctx, "test", &codeartifact.RepositoryArgs{
			Repository: pulumi.String("example"),
			Domain:     exampleDomain.Domain,
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.codeartifact.Domain;
import com.pulumi.aws.codeartifact.DomainArgs;
import com.pulumi.aws.codeartifact.Repository;
import com.pulumi.aws.codeartifact.RepositoryArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("domain key")
            .build());

        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .domain("example")
            .encryptionKey(exampleKey.arn())
            .build());

        var test = new Repository("test", RepositoryArgs.builder()        
            .repository("example")
            .domain(exampleDomain.domain())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: domain key
  exampleDomain:
    type: aws:codeartifact:Domain
    properties:
      domain: example
      encryptionKey: ${exampleKey.arn}
  test:
    type: aws:codeartifact:Repository
    properties:
      repository: example
      domain: ${exampleDomain.domain}
```

{{% /example %}}
{{% example %}}
### With Upstream Repository

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const upstream = new aws.codeartifact.Repository("upstream", {
    repository: "upstream",
    domain: aws_codeartifact_domain.test.domain,
});
const test = new aws.codeartifact.Repository("test", {
    repository: "example",
    domain: aws_codeartifact_domain.example.domain,
    upstreams: [{
        repositoryName: upstream.repository,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

upstream = aws.codeartifact.Repository("upstream",
    repository="upstream",
    domain=aws_codeartifact_domain["test"]["domain"])
test = aws.codeartifact.Repository("test",
    repository="example",
    domain=aws_codeartifact_domain["example"]["domain"],
    upstreams=[aws.codeartifact.RepositoryUpstreamArgs(
        repository_name=upstream.repository,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var upstream = new Aws.CodeArtifact.Repository("upstream", new()
    {
        RepositoryName = "upstream",
        Domain = aws_codeartifact_domain.Test.Domain,
    });

    var test = new Aws.CodeArtifact.Repository("test", new()
    {
        RepositoryName = "example",
        Domain = aws_codeartifact_domain.Example.Domain,
        Upstreams = new[]
        {
            new Aws.CodeArtifact.Inputs.RepositoryUpstreamArgs
            {
                RepositoryName = upstream.RepositoryName,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeartifact"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		upstream, err := codeartifact.NewRepository(ctx, "upstream", &codeartifact.RepositoryArgs{
			Repository: pulumi.String("upstream"),
			Domain:     pulumi.Any(aws_codeartifact_domain.Test.Domain),
		})
		if err != nil {
			return err
		}
		_, err = codeartifact.NewRepository(ctx, "test", &codeartifact.RepositoryArgs{
			Repository: pulumi.String("example"),
			Domain:     pulumi.Any(aws_codeartifact_domain.Example.Domain),
			Upstreams: codeartifact.RepositoryUpstreamArray{
				&codeartifact.RepositoryUpstreamArgs{
					RepositoryName: upstream.Repository,
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
import com.pulumi.aws.codeartifact.Repository;
import com.pulumi.aws.codeartifact.RepositoryArgs;
import com.pulumi.aws.codeartifact.inputs.RepositoryUpstreamArgs;
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
        var upstream = new Repository("upstream", RepositoryArgs.builder()        
            .repository("upstream")
            .domain(aws_codeartifact_domain.test().domain())
            .build());

        var test = new Repository("test", RepositoryArgs.builder()        
            .repository("example")
            .domain(aws_codeartifact_domain.example().domain())
            .upstreams(RepositoryUpstreamArgs.builder()
                .repositoryName(upstream.repository())
                .build())
            .build());

    }
}
```
```yaml
resources:
  upstream:
    type: aws:codeartifact:Repository
    properties:
      repository: upstream
      domain: ${aws_codeartifact_domain.test.domain}
  test:
    type: aws:codeartifact:Repository
    properties:
      repository: example
      domain: ${aws_codeartifact_domain.example.domain}
      upstreams:
        - repositoryName: ${upstream.repository}
```

{{% /example %}}
{{% example %}}
### With External Connection

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const upstream = new aws.codeartifact.Repository("upstream", {
    repository: "upstream",
    domain: aws_codeartifact_domain.test.domain,
});
const test = new aws.codeartifact.Repository("test", {
    repository: "example",
    domain: aws_codeartifact_domain.example.domain,
    externalConnections: {
        externalConnectionName: "public:npmjs",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

upstream = aws.codeartifact.Repository("upstream",
    repository="upstream",
    domain=aws_codeartifact_domain["test"]["domain"])
test = aws.codeartifact.Repository("test",
    repository="example",
    domain=aws_codeartifact_domain["example"]["domain"],
    external_connections=aws.codeartifact.RepositoryExternalConnectionsArgs(
        external_connection_name="public:npmjs",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var upstream = new Aws.CodeArtifact.Repository("upstream", new()
    {
        RepositoryName = "upstream",
        Domain = aws_codeartifact_domain.Test.Domain,
    });

    var test = new Aws.CodeArtifact.Repository("test", new()
    {
        RepositoryName = "example",
        Domain = aws_codeartifact_domain.Example.Domain,
        ExternalConnections = new Aws.CodeArtifact.Inputs.RepositoryExternalConnectionsArgs
        {
            ExternalConnectionName = "public:npmjs",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeartifact"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codeartifact.NewRepository(ctx, "upstream", &codeartifact.RepositoryArgs{
			Repository: pulumi.String("upstream"),
			Domain:     pulumi.Any(aws_codeartifact_domain.Test.Domain),
		})
		if err != nil {
			return err
		}
		_, err = codeartifact.NewRepository(ctx, "test", &codeartifact.RepositoryArgs{
			Repository: pulumi.String("example"),
			Domain:     pulumi.Any(aws_codeartifact_domain.Example.Domain),
			ExternalConnections: &codeartifact.RepositoryExternalConnectionsArgs{
				ExternalConnectionName: pulumi.String("public:npmjs"),
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
import com.pulumi.aws.codeartifact.Repository;
import com.pulumi.aws.codeartifact.RepositoryArgs;
import com.pulumi.aws.codeartifact.inputs.RepositoryExternalConnectionsArgs;
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
        var upstream = new Repository("upstream", RepositoryArgs.builder()        
            .repository("upstream")
            .domain(aws_codeartifact_domain.test().domain())
            .build());

        var test = new Repository("test", RepositoryArgs.builder()        
            .repository("example")
            .domain(aws_codeartifact_domain.example().domain())
            .externalConnections(RepositoryExternalConnectionsArgs.builder()
                .externalConnectionName("public:npmjs")
                .build())
            .build());

    }
}
```
```yaml
resources:
  upstream:
    type: aws:codeartifact:Repository
    properties:
      repository: upstream
      domain: ${aws_codeartifact_domain.test.domain}
  test:
    type: aws:codeartifact:Repository
    properties:
      repository: example
      domain: ${aws_codeartifact_domain.example.domain}
      externalConnections:
        externalConnectionName: public:npmjs
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeArtifact Repository using the CodeArtifact Repository ARN. For example:

```sh
 $ pulumi import aws:codeartifact/repository:Repository example arn:aws:codeartifact:us-west-2:012345678912:repository/tf-acc-test-6968272603913957763/tf-acc-test-6968272603913957763
```
 