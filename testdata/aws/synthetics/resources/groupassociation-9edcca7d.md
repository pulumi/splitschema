Provides a Synthetics Group Association resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.synthetics.GroupAssociation("example", {
    groupName: aws_synthetics_group.example.name,
    canaryArn: aws_synthetics_canary.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.synthetics.GroupAssociation("example",
    group_name=aws_synthetics_group["example"]["name"],
    canary_arn=aws_synthetics_canary["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Synthetics.GroupAssociation("example", new()
    {
        GroupName = aws_synthetics_group.Example.Name,
        CanaryArn = aws_synthetics_canary.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/synthetics"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synthetics.NewGroupAssociation(ctx, "example", &synthetics.GroupAssociationArgs{
			GroupName: pulumi.Any(aws_synthetics_group.Example.Name),
			CanaryArn: pulumi.Any(aws_synthetics_canary.Example.Arn),
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
import com.pulumi.aws.synthetics.GroupAssociation;
import com.pulumi.aws.synthetics.GroupAssociationArgs;
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
        var example = new GroupAssociation("example", GroupAssociationArgs.builder()        
            .groupName(aws_synthetics_group.example().name())
            .canaryArn(aws_synthetics_canary.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:synthetics:GroupAssociation
    properties:
      groupName: ${aws_synthetics_group.example.name}
      canaryArn: ${aws_synthetics_canary.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Synthetics Group Association using the `canary_arn,group_name`. For example:

```sh
 $ pulumi import aws:synthetics/groupAssociation:GroupAssociation example arn:aws:synthetics:us-west-2:123456789012:canary:tf-acc-test-abcd1234,examplename
```
 