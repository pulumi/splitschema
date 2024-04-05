Use this data source to get information about a specific EC2 Key Pair.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how to get a EC2 Key Pair including the public key material from its name.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getKeyPair({
    keyName: "test",
    includePublicKey: true,
    filters: [{
        name: "tag:Component",
        values: ["web"],
    }],
});
export const fingerprint = example.then(example => example.fingerprint);
export const name = example.then(example => example.keyName);
export const id = example.then(example => example.id);
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_key_pair(key_name="test",
    include_public_key=True,
    filters=[aws.ec2.GetKeyPairFilterArgs(
        name="tag:Component",
        values=["web"],
    )])
pulumi.export("fingerprint", example.fingerprint)
pulumi.export("name", example.key_name)
pulumi.export("id", example.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetKeyPair.Invoke(new()
    {
        KeyName = "test",
        IncludePublicKey = true,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetKeyPairFilterInputArgs
            {
                Name = "tag:Component",
                Values = new[]
                {
                    "web",
                },
            },
        },
    });

    return new Dictionary<string, object?>
    {
        ["fingerprint"] = example.Apply(getKeyPairResult => getKeyPairResult.Fingerprint),
        ["name"] = example.Apply(getKeyPairResult => getKeyPairResult.KeyName),
        ["id"] = example.Apply(getKeyPairResult => getKeyPairResult.Id),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := ec2.LookupKeyPair(ctx, &ec2.LookupKeyPairArgs{
			KeyName:          pulumi.StringRef("test"),
			IncludePublicKey: pulumi.BoolRef(true),
			Filters: []ec2.GetKeyPairFilter{
				{
					Name: "tag:Component",
					Values: []string{
						"web",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("fingerprint", example.Fingerprint)
		ctx.Export("name", example.KeyName)
		ctx.Export("id", example.Id)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetKeyPairArgs;
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
        final var example = Ec2Functions.getKeyPair(GetKeyPairArgs.builder()
            .keyName("test")
            .includePublicKey(true)
            .filters(GetKeyPairFilterArgs.builder()
                .name("tag:Component")
                .values("web")
                .build())
            .build());

        ctx.export("fingerprint", example.applyValue(getKeyPairResult -> getKeyPairResult.fingerprint()));
        ctx.export("name", example.applyValue(getKeyPairResult -> getKeyPairResult.keyName()));
        ctx.export("id", example.applyValue(getKeyPairResult -> getKeyPairResult.id()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getKeyPair
      Arguments:
        keyName: test
        includePublicKey: true
        filters:
          - name: tag:Component
            values:
              - web
outputs:
  fingerprint: ${example.fingerprint}
  name: ${example.keyName}
  id: ${example.id}
```
{{% /example %}}
{{% /examples %}}