Data source for managing an AWS OpenSearch Serverless Lifecycle Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.opensearch.getServerlessLifecyclePolicy({
    name: "example-lifecycle-policy",
    type: "retention",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opensearch.get_serverless_lifecycle_policy(name="example-lifecycle-policy",
    type="retention")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.OpenSearch.GetServerlessLifecyclePolicy.Invoke(new()
    {
        Name = "example-lifecycle-policy",
        Type = "retention",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opensearch.LookupServerlessLifecyclePolicy(ctx, &opensearch.LookupServerlessLifecyclePolicyArgs{
			Name: "example-lifecycle-policy",
			Type: "retention",
		}, nil)
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
import com.pulumi.aws.opensearch.OpensearchFunctions;
import com.pulumi.aws.opensearch.inputs.GetServerlessLifecyclePolicyArgs;
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
        final var example = OpensearchFunctions.getServerlessLifecyclePolicy(GetServerlessLifecyclePolicyArgs.builder()
            .name("example-lifecycle-policy")
            .type("retention")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:opensearch:getServerlessLifecyclePolicy
      Arguments:
        name: example-lifecycle-policy
        type: retention
```
{{% /example %}}
{{% /examples %}}