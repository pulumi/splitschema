Attaches an EC2 instance to an Elastic Load Balancer (ELB). For attaching resources with Application Load Balancer (ALB) or Network Load Balancer (NLB), see the `aws.lb.TargetGroupAttachment` resource.

> **NOTE on ELB Instances and ELB Attachments:** This provider currently provides
both a standalone ELB Attachment resource (describing an instance attached to
an ELB), and an Elastic Load Balancer resource with
`instances` defined in-line. At this time you cannot use an ELB with in-line
instances in conjunction with an ELB Attachment resource. Doing so will cause a
conflict and will overwrite attachments.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new load balancer attachment
const baz = new aws.elb.Attachment("baz", {
    elb: aws_elb.bar.id,
    instance: aws_instance.foo.id,
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new load balancer attachment
baz = aws.elb.Attachment("baz",
    elb=aws_elb["bar"]["id"],
    instance=aws_instance["foo"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new load balancer attachment
    var baz = new Aws.Elb.Attachment("baz", new()
    {
        Elb = aws_elb.Bar.Id,
        Instance = aws_instance.Foo.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elb.NewAttachment(ctx, "baz", &elb.AttachmentArgs{
			Elb:      pulumi.Any(aws_elb.Bar.Id),
			Instance: pulumi.Any(aws_instance.Foo.Id),
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
import com.pulumi.aws.elb.Attachment;
import com.pulumi.aws.elb.AttachmentArgs;
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
        var baz = new Attachment("baz", AttachmentArgs.builder()        
            .elb(aws_elb.bar().id())
            .instance(aws_instance.foo().id())
            .build());

    }
}
```
```yaml
resources:
  # Create a new load balancer attachment
  baz:
    type: aws:elb:Attachment
    properties:
      elb: ${aws_elb.bar.id}
      instance: ${aws_instance.foo.id}
```
{{% /example %}}
{{% /examples %}}