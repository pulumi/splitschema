Provides an Elastic Container Registry Pull Through Cache Rule.

More information about pull through cache rules, including the set of supported
upstream repositories, see [Using pull through cache rules](https://docs.aws.amazon.com/AmazonECR/latest/userguide/pull-through-cache.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ecr.PullThroughCacheRule("example", {
    ecrRepositoryPrefix: "ecr-public",
    upstreamRegistryUrl: "public.ecr.aws",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ecr.PullThroughCacheRule("example",
    ecr_repository_prefix="ecr-public",
    upstream_registry_url="public.ecr.aws")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ecr.PullThroughCacheRule("example", new()
    {
        EcrRepositoryPrefix = "ecr-public",
        UpstreamRegistryUrl = "public.ecr.aws",
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
		_, err := ecr.NewPullThroughCacheRule(ctx, "example", &ecr.PullThroughCacheRuleArgs{
			EcrRepositoryPrefix: pulumi.String("ecr-public"),
			UpstreamRegistryUrl: pulumi.String("public.ecr.aws"),
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
import com.pulumi.aws.ecr.PullThroughCacheRule;
import com.pulumi.aws.ecr.PullThroughCacheRuleArgs;
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
        var example = new PullThroughCacheRule("example", PullThroughCacheRuleArgs.builder()        
            .ecrRepositoryPrefix("ecr-public")
            .upstreamRegistryUrl("public.ecr.aws")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecr:PullThroughCacheRule
    properties:
      ecrRepositoryPrefix: ecr-public
      upstreamRegistryUrl: public.ecr.aws
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a pull-through cache rule using the `ecr_repository_prefix`. For example:

```sh
 $ pulumi import aws:ecr/pullThroughCacheRule:PullThroughCacheRule example ecr-public
```
 