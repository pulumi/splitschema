Provides the ability to register a target with an AWS VPC Lattice Target Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.TargetGroupAttachment("example", {
    targetGroupIdentifier: aws_vpclattice_target_group.example.id,
    target: {
        id: aws_lb.example.arn,
        port: 80,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.TargetGroupAttachment("example",
    target_group_identifier=aws_vpclattice_target_group["example"]["id"],
    target=aws.vpclattice.TargetGroupAttachmentTargetArgs(
        id=aws_lb["example"]["arn"],
        port=80,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.TargetGroupAttachment("example", new()
    {
        TargetGroupIdentifier = aws_vpclattice_target_group.Example.Id,
        Target = new Aws.VpcLattice.Inputs.TargetGroupAttachmentTargetArgs
        {
            Id = aws_lb.Example.Arn,
            Port = 80,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewTargetGroupAttachment(ctx, "example", &vpclattice.TargetGroupAttachmentArgs{
			TargetGroupIdentifier: pulumi.Any(aws_vpclattice_target_group.Example.Id),
			Target: &vpclattice.TargetGroupAttachmentTargetArgs{
				Id:   pulumi.Any(aws_lb.Example.Arn),
				Port: pulumi.Int(80),
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
import com.pulumi.aws.vpclattice.TargetGroupAttachment;
import com.pulumi.aws.vpclattice.TargetGroupAttachmentArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupAttachmentTargetArgs;
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
        var example = new TargetGroupAttachment("example", TargetGroupAttachmentArgs.builder()        
            .targetGroupIdentifier(aws_vpclattice_target_group.example().id())
            .target(TargetGroupAttachmentTargetArgs.builder()
                .id(aws_lb.example().arn())
                .port(80)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:TargetGroupAttachment
    properties:
      targetGroupIdentifier: ${aws_vpclattice_target_group.example.id}
      target:
        id: ${aws_lb.example.arn}
        port: 80
```
{{% /example %}}
{{% /examples %}}