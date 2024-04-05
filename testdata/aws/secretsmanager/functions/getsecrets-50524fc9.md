Use this data source to get the ARNs and names of Secrets Manager secrets matching the specified criteria.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.secretsmanager.getSecrets({
    filters: [{
        name: "name",
        values: ["example"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.secretsmanager.get_secrets(filters=[aws.secretsmanager.GetSecretsFilterArgs(
    name="name",
    values=["example"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SecretsManager.GetSecrets.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.SecretsManager.Inputs.GetSecretsFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "example",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.GetSecrets(ctx, &secretsmanager.GetSecretsArgs{
			Filters: []secretsmanager.GetSecretsFilter{
				{
					Name: "name",
					Values: []string{
						"example",
					},
				},
			},
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
import com.pulumi.aws.secretsmanager.SecretsmanagerFunctions;
import com.pulumi.aws.secretsmanager.inputs.GetSecretsArgs;
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
        final var example = SecretsmanagerFunctions.getSecrets(GetSecretsArgs.builder()
            .filters(GetSecretsFilterArgs.builder()
                .name("name")
                .values("example")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:secretsmanager:getSecrets
      Arguments:
        filters:
          - name: name
            values:
              - example
```
{{% /example %}}
{{% /examples %}}