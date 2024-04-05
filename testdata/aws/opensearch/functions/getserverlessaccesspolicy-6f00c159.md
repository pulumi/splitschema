Data source for managing an AWS OpenSearch Serverless Access Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.opensearch.getServerlessAccessPolicy({
    name: aws_opensearchserverless_access_policy.example.name,
    type: aws_opensearchserverless_access_policy.example.type,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opensearch.get_serverless_access_policy(name=aws_opensearchserverless_access_policy["example"]["name"],
    type=aws_opensearchserverless_access_policy["example"]["type"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.OpenSearch.GetServerlessAccessPolicy.Invoke(new()
    {
        Name = aws_opensearchserverless_access_policy.Example.Name,
        Type = aws_opensearchserverless_access_policy.Example.Type,
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
		_, err := opensearch.LookupServerlessAccessPolicy(ctx, &opensearch.LookupServerlessAccessPolicyArgs{
			Name: aws_opensearchserverless_access_policy.Example.Name,
			Type: aws_opensearchserverless_access_policy.Example.Type,
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
import com.pulumi.aws.opensearch.inputs.GetServerlessAccessPolicyArgs;
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
        final var example = OpensearchFunctions.getServerlessAccessPolicy(GetServerlessAccessPolicyArgs.builder()
            .name(aws_opensearchserverless_access_policy.example().name())
            .type(aws_opensearchserverless_access_policy.example().type())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:opensearch:getServerlessAccessPolicy
      Arguments:
        name: ${aws_opensearchserverless_access_policy.example.name}
        type: ${aws_opensearchserverless_access_policy.example.type}
```
{{% /example %}}
{{% /examples %}}